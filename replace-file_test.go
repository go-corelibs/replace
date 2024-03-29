// Copyright (c) 2024  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package replace

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	gProcessTestGoSrc = "replace-file_test.go"
	gTestingTestMd    = "_testing/test.md"
)

func TestReplaceFile(t *testing.T) {
	Convey("ProcessFile", t, func() {
		original, modified, count, diff, err := ProcessFile(gProcessTestGoSrc, func(original string) (modified string, count int) {
			modified = original
			count = 0
			return
		})
		So(err, ShouldEqual, nil)
		So(modified, ShouldEqual, original)
		So(count, ShouldEqual, 0)
		So(diff.Len(), ShouldEqual, 0)
		_, _, _, _, err = ProcessFile(gProcessTestGoSrc+".nope", func(original string) (modified string, count int) {
			modified = original
			count = 0
			return
		})
		So(err, ShouldNotEqual, nil)
	})

}
