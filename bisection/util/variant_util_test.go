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

package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	pb "go.chromium.org/luci/bisection/proto/v1"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestVariantUtil(t *testing.T) {
	Convey("VariantPB", t, func() {
		Convey("no error", func() {
			variant, err := VariantPB(`{"builder": "testbuilder"}`)
			So(err, ShouldBeNil)
			So(variant, ShouldResembleProto, &pb.Variant{
				Def: map[string]string{"builder": "testbuilder"},
			})
		})
		Convey("error", func() {
			variant, err := VariantPB("invalid json")
			So(err, ShouldErrLike, "invalid")
			So(variant, ShouldBeNil)
		})
	})

	Convey("VariantToStrings", t, func() {
		variants := VariantToStrings(&pb.Variant{
			Def: map[string]string{"builder": "testbuilder", "os": "mac"},
		})
		So(variants, ShouldResemble, []string{"builder:testbuilder", "os:mac"})
	})
}