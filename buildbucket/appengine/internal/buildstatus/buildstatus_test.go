// Copyright 2023 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package buildstatus

import (
	"context"
	"testing"

	"go.chromium.org/luci/gae/impl/memory"

	"go.chromium.org/luci/gae/filter/txndefer"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/buildbucket/appengine/model"
	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func update(ctx context.Context, u *Updater) (*model.BuildStatus, error) {
	var bs *model.BuildStatus
	txErr := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		var err error
		bs, err = u.Do(ctx)
		return err
	}, nil)
	return bs, txErr
}
func TestUpdate(t *testing.T) {
	t.Parallel()
	Convey("Update", t, func() {
		ctx := memory.Use(context.Background())
		ctx = txndefer.FilterRDS(ctx)
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)

		Convey("fail", func() {

			Convey("not in transaction", func() {
				u := &Updater{}
				_, err := u.Do(ctx)
				So(err, ShouldErrLike, "must update build status in a transaction")
			})

			Convey("update ended build", func() {
				b := &model.Build{
					ID: 1,
					Proto: &pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Status: pb.Status_SUCCESS,
					},
					Status: pb.Status_SUCCESS,
				}
				So(datastore.Put(ctx, b), ShouldBeNil)
				u := &Updater{
					Build:       b,
					BuildStatus: pb.Status_SUCCESS,
				}
				_, err := update(ctx, u)
				So(err, ShouldErrLike, "cannot update status for an ended build")
			})

			Convey("output status and task status", func() {
				b := &model.Build{
					ID: 1,
					Proto: &pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Status: pb.Status_SCHEDULED,
					},
					Status: pb.Status_SCHEDULED,
				}
				So(datastore.Put(ctx, b), ShouldBeNil)
				u := &Updater{
					Build:        b,
					OutputStatus: pb.Status_SUCCESS,
					TaskStatus:   pb.Status_SUCCESS,
				}
				_, err := update(ctx, u)
				So(err, ShouldErrLike, "impossible: update build output status and task status at the same time")
			})

			Convey("nothing is provided to update", func() {
				b := &model.Build{
					ID: 1,
					Proto: &pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Status: pb.Status_SCHEDULED,
					},
					Status: pb.Status_SCHEDULED,
				}
				So(datastore.Put(ctx, b), ShouldBeNil)
				u := &Updater{
					Build: b,
				}
				_, err := update(ctx, u)
				So(err, ShouldErrLike, "cannot set a build status to UNSPECIFIED")
			})

			Convey("BuildStatus not found", func() {
				b := &model.Build{
					ID: 1,
					Proto: &pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Status: pb.Status_SCHEDULED,
					},
					Status: pb.Status_SCHEDULED,
				}
				So(datastore.Put(ctx, b), ShouldBeNil)
				u := &Updater{
					Build:       b,
					BuildStatus: pb.Status_SUCCESS,
				}
				_, err := update(ctx, u)
				So(err, ShouldErrLike, "not found")
			})
		})

		Convey("pass", func() {

			b := &model.Build{
				ID: 87654321,
				Proto: &pb.Build{
					Id: 87654321,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Status: pb.Status_SCHEDULED,
				},
				Status: pb.Status_SCHEDULED,
			}
			bk := datastore.KeyForObj(ctx, b)
			bs := &model.BuildStatus{
				Build:  bk,
				Status: pb.Status_SCHEDULED,
			}
			So(datastore.Put(ctx, b, bs), ShouldBeNil)
			updatedStatus := b.Proto.Status
			u := &Updater{
				Build: b,
				PostProcess: func(c context.Context, bld *model.Build) error {
					updatedStatus = bld.Proto.Status
					return nil
				},
			}

			Convey("direct update on build status ignore sub status", func() {
				u.BuildStatus = pb.Status_STARTED
				u.OutputStatus = pb.Status_SUCCESS // only for test, impossible in practice
				bs, err := update(ctx, u)
				So(err, ShouldBeNil)
				So(bs.Status, ShouldEqual, pb.Status_STARTED)
				So(updatedStatus, ShouldEqual, pb.Status_STARTED)
			})

			Convey("update output status", func() {
				Convey("start, so build status is updated", func() {
					u.OutputStatus = pb.Status_STARTED
					bs, err := update(ctx, u)
					So(err, ShouldBeNil)
					So(bs.Status, ShouldEqual, pb.Status_STARTED)
					So(updatedStatus, ShouldEqual, pb.Status_STARTED)
				})

				Convey("end, so build status is unchanged", func() {
					u.OutputStatus = pb.Status_SUCCESS
					bs, err := update(ctx, u)
					So(err, ShouldBeNil)
					So(bs, ShouldBeNil)
					So(updatedStatus, ShouldEqual, pb.Status_SCHEDULED)
				})
			})

			Convey("update task status", func() {
				Convey("end, so build status is updated", func() {
					u.TaskStatus = pb.Status_SUCCESS
					bs, err := update(ctx, u)
					So(err, ShouldBeNil)
					So(bs.Status, ShouldEqual, pb.Status_SUCCESS)
					So(updatedStatus, ShouldEqual, pb.Status_SUCCESS)
				})

				Convey("start, so build status is unchanged", func() {
					u.TaskStatus = pb.Status_STARTED
					bs, err := update(ctx, u)
					So(err, ShouldBeNil)
					So(bs, ShouldBeNil)
					So(updatedStatus, ShouldEqual, pb.Status_SCHEDULED)
				})
			})
		})
	})
}