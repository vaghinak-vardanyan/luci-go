// Copyright 2022 The LUCI Authors.
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

package datastoreutil

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"go.chromium.org/luci/bisection/model"
	pb "go.chromium.org/luci/bisection/proto"

	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
)

func TestCountLatestRevertsCreated(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	datastore.GetTestable(c).AddIndexes(
		&datastore.IndexDefinition{
			Kind: "Suspect",
			SortBy: []datastore.IndexColumn{
				{
					Property: "is_revert_created",
				},
				{
					Property: "revert_create_time",
				},
			},
		},
	)
	datastore.GetTestable(c).CatchupIndexes()

	// Set test clock
	cl := testclock.New(testclock.TestTimeUTC)
	c = clock.Set(c, cl)

	Convey("No suspects at all", t, func() {
		count, err := CountLatestRevertsCreated(c, 24)
		So(err, ShouldBeNil)
		So(count, ShouldEqual, 0)
	})

	Convey("Count of recent created reverts", t, func() {
		// Set up suspects
		suspect1 := &model.Suspect{}
		suspect2 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				RevertURL:        "https://chromium-test-review.googlesource.com/c/test/project/+/100000",
				IsRevertCreated:  true,
				RevertCreateTime: clock.Now(c),
			},
		}
		suspect3 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				RevertURL:        "https://chromium-test-review.googlesource.com/c/test/project/+/100001",
				IsRevertCreated:  true,
				RevertCreateTime: clock.Now(c).Add(-time.Hour * 72),
			},
		}
		suspect4 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				RevertURL:        "https://chromium-test-review.googlesource.com/c/test/project/+/100002",
				IsRevertCreated:  true,
				RevertCreateTime: clock.Now(c).Add(-time.Hour * 4),
			},
		}
		suspect5 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				RevertURL:        "https://chromium-test-review.googlesource.com/c/test/project/+/100003",
				IsRevertCreated:  true,
				RevertCreateTime: clock.Now(c).Add(-time.Minute * 10),
			},
		}
		suspect6 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				RevertURL:        "https://chromium-test-review.googlesource.com/c/test/project/+/100004",
				IsRevertCreated:  true,
				RevertCreateTime: clock.Now(c).Add(-time.Hour * 24),
			},
		}
		suspect7 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				RevertURL:        "",
				IsRevertCreated:  false,
				RevertCreateTime: clock.Now(c).Add(-time.Minute * 10),
			},
		}
		So(datastore.Put(c, suspect1), ShouldBeNil)
		So(datastore.Put(c, suspect2), ShouldBeNil)
		So(datastore.Put(c, suspect3), ShouldBeNil)
		So(datastore.Put(c, suspect4), ShouldBeNil)
		So(datastore.Put(c, suspect5), ShouldBeNil)
		So(datastore.Put(c, suspect6), ShouldBeNil)
		So(datastore.Put(c, suspect7), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		count, err := CountLatestRevertsCreated(c, 24)
		So(err, ShouldBeNil)
		So(count, ShouldEqual, 3)
	})
}

func TestCountLatestRevertsCommitted(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	datastore.GetTestable(c).AddIndexes(
		&datastore.IndexDefinition{
			Kind: "Suspect",
			SortBy: []datastore.IndexColumn{
				{
					Property: "is_revert_committed",
				},
				{
					Property: "revert_commit_time",
				},
			},
		},
	)
	datastore.GetTestable(c).CatchupIndexes()

	// Set test clock
	cl := testclock.New(testclock.TestTimeUTC)
	c = clock.Set(c, cl)

	Convey("No suspects at all", t, func() {
		count, err := CountLatestRevertsCommitted(c, 24)
		So(err, ShouldBeNil)
		So(count, ShouldEqual, 0)
	})

	Convey("Count of recent committed reverts", t, func() {
		// Set up suspects
		suspect1 := &model.Suspect{}
		suspect2 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				IsRevertCommitted: true,
				RevertCommitTime:  clock.Now(c),
			},
		}
		suspect3 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				IsRevertCommitted: true,
				RevertCommitTime:  clock.Now(c).Add(-time.Hour * 72),
			},
		}
		suspect4 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				IsRevertCommitted: true,
				RevertCommitTime:  clock.Now(c).Add(-time.Hour * 4),
			},
		}
		suspect5 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				IsRevertCommitted: true,
				RevertCommitTime:  clock.Now(c).Add(-time.Minute * 10),
			},
		}
		suspect6 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				IsRevertCommitted: true,
				RevertCommitTime:  clock.Now(c).Add(-time.Hour * 24),
			},
		}
		suspect7 := &model.Suspect{
			RevertDetails: model.RevertDetails{
				IsRevertCommitted: false,
				RevertCommitTime:  clock.Now(c).Add(-time.Minute * 10),
			},
		}
		So(datastore.Put(c, suspect1), ShouldBeNil)
		So(datastore.Put(c, suspect2), ShouldBeNil)
		So(datastore.Put(c, suspect3), ShouldBeNil)
		So(datastore.Put(c, suspect4), ShouldBeNil)
		So(datastore.Put(c, suspect5), ShouldBeNil)
		So(datastore.Put(c, suspect6), ShouldBeNil)
		So(datastore.Put(c, suspect7), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		count, err := CountLatestRevertsCommitted(c, 24)
		So(err, ShouldBeNil)
		So(count, ShouldEqual, 3)
	})
}

func TestGetAssociatedBuildID(t *testing.T) {
	ctx := memory.Use(context.Background())

	Convey("Associated failed build ID for heuristic suspect", t, func() {
		failedBuild := &model.LuciFailedBuild{
			Id: 88128398584903,
			LuciBuild: model.LuciBuild{
				BuildId:     88128398584903,
				Project:     "chromium",
				Bucket:      "ci",
				Builder:     "android",
				BuildNumber: 123,
			},
			BuildFailureType: pb.BuildFailureType_COMPILE,
		}
		So(datastore.Put(ctx, failedBuild), ShouldBeNil)
		datastore.GetTestable(ctx).CatchupIndexes()
		compileFailure := &model.CompileFailure{
			Build: datastore.KeyForObj(ctx, failedBuild),
		}
		So(datastore.Put(ctx, compileFailure), ShouldBeNil)
		datastore.GetTestable(ctx).CatchupIndexes()
		analysis := &model.CompileFailureAnalysis{
			Id:             444,
			CompileFailure: datastore.KeyForObj(ctx, compileFailure),
		}
		So(datastore.Put(ctx, analysis), ShouldBeNil)
		datastore.GetTestable(ctx).CatchupIndexes()
		heuristicAnalysis := &model.CompileHeuristicAnalysis{
			ParentAnalysis: datastore.KeyForObj(ctx, analysis),
		}
		So(datastore.Put(ctx, heuristicAnalysis), ShouldBeNil)
		datastore.GetTestable(ctx).CatchupIndexes()

		heuristicSuspect := &model.Suspect{
			Id:             1,
			Type:           model.SuspectType_Heuristic,
			Score:          10,
			ParentAnalysis: datastore.KeyForObj(ctx, heuristicAnalysis),
			GitilesCommit: buildbucketpb.GitilesCommit{
				Host:    "test.googlesource.com",
				Project: "chromium/test",
				Id:      "12ab34cd56ef",
			},
			ReviewUrl:          "https://test-review.googlesource.com/c/chromium/test/+/876543",
			VerificationStatus: model.SuspectVerificationStatus_UnderVerification,
		}
		So(datastore.Put(ctx, heuristicSuspect), ShouldBeNil)
		datastore.GetTestable(ctx).CatchupIndexes()

		bbid, err := GetAssociatedBuildID(ctx, heuristicSuspect)
		So(err, ShouldBeNil)
		So(bbid, ShouldEqual, 88128398584903)
	})
}