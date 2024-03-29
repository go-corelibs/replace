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
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegexFile(t *testing.T) {
	Convey("RegexFile", t, func() {
		original, modified, count, diff, err := RegexFile(regexp.MustCompile(`(?i)the`), `THE`, gTestingTestMd)
		So(err, ShouldEqual, nil)
		So(modified, ShouldNotEqual, original)
		So(count, ShouldEqual, 4)
		So(diff.Len(), ShouldEqual, 6)
	})

	Convey("RegexLinesFile", t, func() {
		original, modified, count, diff, err := RegexLinesFile(regexp.MustCompile(`(?i)the`), `THE`, gTestingTestMd)
		So(err, ShouldEqual, nil)
		So(modified, ShouldNotEqual, original)
		So(count, ShouldEqual, 4)
		So(diff.Len(), ShouldEqual, 6)
	})

	Convey("RegexPreserveFile", t, func() {
		original, modified, count, diff, err := RegexPreserveFile(regexp.MustCompile(`(?i)the`), "this", gTestingTestMd)
		So(err, ShouldEqual, nil)
		So(modified, ShouldNotEqual, original)
		So(count, ShouldEqual, 4)
		So(diff.Len(), ShouldEqual, 6)
	})
}
