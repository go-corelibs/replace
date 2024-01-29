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
	"os"
	"regexp"
	"strings"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/go-corelibs/globs"
)

func TestFinders(t *testing.T) {
	Convey("IsIncluded", t, func() {
		var err error
		var include, exclude globs.Globs
		So(IsIncluded(nil, nil, "/test"), ShouldEqual, true)
		include, err = globs.Parse("*.txt")
		So(err, ShouldEqual, nil)
		So(IsIncluded(include, nil, "_testing/test.txt"), ShouldEqual, true)
		exclude, err = globs.Parse("*.md")
		So(err, ShouldEqual, nil)
		So(IsIncluded(nil, exclude, "_testing/test.txt"), ShouldEqual, true)
		exclude, err = globs.Parse("*.txt")
		So(err, ShouldEqual, nil)
		So(IsIncluded(nil, exclude, "_testing/test.txt"), ShouldEqual, false)
	})

	Convey("FindAllIncluded", t, func() {
		files := []string{"_testing/.hello-world.hidden", "_testing/test.txt", "_testing/test.txt"}
		results := FindAllIncluded(files, false, false, false, false, nil, nil)
		So(len(results), ShouldEqual, 1)
		files = []string{"_testing"}
		results = FindAllIncluded(files, false, false, false, true, nil, nil)
		So(len(results), ShouldEqual, 2)
	})

	Convey("FindAllMatchingRegexp", t, func() {
		targets := []string{"_testing"}
		rx := regexp.MustCompile(`(?ms)te[sx]t`)
		found, matches, err := FindAllMatchingRegexp(rx, targets, false, false, false, true, nil, nil, nil)
		So(err, ShouldBeNil)
		So(len(found), ShouldEqual, 2)
		So(len(matches), ShouldEqual, 2)
	})

	Convey("FindAllMatchingRegexpLines", t, func() {
		targets := []string{"_testing"}
		rx := regexp.MustCompile(`(?ms)test`)
		found, matches, err := FindAllMatchingRegexpLines(rx, targets, false, false, false, true, nil, nil, nil)
		So(err, ShouldBeNil)
		So(len(found), ShouldEqual, 2)
		So(len(matches), ShouldEqual, 1)
	})

	Convey("FindAllMatchingString", t, func() {
		targets := []string{"_testing"}
		found, matches, err := FindAllMatchingString("test", targets, false, false, false, true, nil, nil, nil)
		So(err, ShouldBeNil)
		So(len(found), ShouldEqual, 2)
		So(len(matches), ShouldEqual, 1)
	})

	Convey("FindAllMatchingStringInsensitive", t, func() {
		targets := []string{"_testing"}
		found, matches, err := FindAllMatchingStringInsensitive("test", targets, false, false, false, true, nil, nil, nil)
		So(err, ShouldBeNil)
		So(len(found), ShouldEqual, 2)
		So(len(matches), ShouldEqual, 2)
	})

	Convey("FindAllMatcher", t, func() {
		targets := []string{"_testing"}
		origMaxFileSize := MaxFileSize
		origMaxFileCount := MaxFileCount
		m := &sync.Mutex{}

		Convey("too many files error", func(c C) {
			m.Lock()
			defer func() {
				MaxFileCount = origMaxFileCount
				m.Unlock()
			}()
			MaxFileCount = 1
			found, matches, err := FindAllMatcher(targets, false, false, false, true, nil, nil, nil, func(data []byte) (matched bool) {
				matched = strings.Contains(string(data), "test")
				return
			})
			So(err, ShouldEqual, ErrTooManyFiles)
			So(len(found), ShouldEqual, 2)
			So(len(matches), ShouldEqual, 0)
		})

		Convey("large file error", func(c C) {
			m.Lock()
			defer func() {
				MaxFileSize = origMaxFileSize
				m.Unlock()
			}()
			MaxFileSize = 1
			found, matches, err := FindAllMatcher(
				targets, false, false, false, true, nil, nil,
				func(file string, matched bool, err error) {
					c.So(err, ShouldEqual, ErrLargeFile)
					return
				},
				func(data []byte) (matched bool) {
					matched = strings.Contains(string(data), "test")
					return
				},
			)
			So(err, ShouldBeNil)
			So(len(found), ShouldEqual, 2)
			So(len(matches), ShouldEqual, 0)
		})

		Convey("binary file error", func(c C) {
			m.Lock()
			defer m.Unlock()
			MaxFileSize = origMaxFileSize
			MaxFileCount = origMaxFileCount
			thisBin, err0 := os.Executable()
			So(err0, ShouldEqual, nil)
			binaryTargets := []string{thisBin}
			found, matches, err := FindAllMatcher(
				binaryTargets, false, false, false, false, nil, nil,
				func(file string, matched bool, err error) {
					c.So(err, ShouldEqual, ErrBinaryFile)
					return
				},
				func(data []byte) (matched bool) {
					matched = strings.Contains(string(data), "test")
					return
				},
			)
			So(err, ShouldBeNil)
			So(len(found), ShouldEqual, 1)
			So(len(matches), ShouldEqual, 0)
		})

	})

}
