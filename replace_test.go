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

const (
	wrTestOriginal = `One one ONE`
	wrTestReplaced = `two two two`
	wrTestPreserve = `Two two TWO`
)

func TestReplace(t *testing.T) {
	Convey("Regex", t, func() {
		modified, count := Regex(
			regexp.MustCompile(`one`),
			`two`,
			wrTestOriginal,
		)
		So(count, ShouldEqual, 1)
		So(modified, ShouldEqual, `One two ONE`)
		modified, count = Regex(
			regexp.MustCompile(`(?i)one`),
			`two`,
			wrTestOriginal,
		)
		So(count, ShouldEqual, 3)
		So(modified, ShouldEqual, wrTestReplaced)
		modified, count = Regex(
			regexp.MustCompile(`(?i)none`),
			`two`,
			wrTestOriginal,
		)
		So(count, ShouldEqual, 0)
		So(modified, ShouldEqual, wrTestOriginal)
		modified, count = Regex(
			nil,
			`two`,
			wrTestOriginal,
		)
		So(count, ShouldEqual, 0)
		So(modified, ShouldEqual, wrTestOriginal)
	})

	Convey("String", t, func() {
		modified, count := String(`one`, `two`, wrTestOriginal)
		So(count, ShouldEqual, 1)
		So(modified, ShouldEqual, `One two ONE`)
		modified, count = String(`one`, `one`, wrTestOriginal)
		So(count, ShouldEqual, 0)
		So(modified, ShouldEqual, wrTestOriginal)
		modified, count = String(``, `two`, wrTestOriginal)
		So(count, ShouldEqual, 0)
		So(modified, ShouldEqual, `One one ONE`)
		modified, count = String(`none`, `two`, wrTestOriginal)
		So(count, ShouldEqual, 0)
		So(modified, ShouldEqual, `One one ONE`)
	})

	Convey("StringInsensitive", t, func() {
		modified, count := StringInsensitive(`one`, `two`, wrTestOriginal)
		So(count, ShouldEqual, 3)
		So(modified, ShouldEqual, wrTestReplaced)
		modified, count = StringInsensitive(`one`, `one`, wrTestOriginal)
		So(count, ShouldEqual, 0)
		So(modified, ShouldEqual, wrTestOriginal)
		modified, count = StringInsensitive(``, `two`, wrTestOriginal)
		So(count, ShouldEqual, 0)
		So(modified, ShouldEqual, wrTestOriginal)
		modified, count = StringInsensitive(`none`, `two`, wrTestOriginal)
		So(count, ShouldEqual, 0)
		So(modified, ShouldEqual, wrTestOriginal)
	})

	Convey("StringPreserve", t, func() {
		modified, count := StringPreserve(`one`, `one`, wrTestOriginal)
		So(count, ShouldEqual, 0)
		So(modified, ShouldEqual, wrTestOriginal)
		modified, count = StringPreserve(`One one`, `two`, wrTestOriginal)
		So(count, ShouldEqual, 1)
		So(modified, ShouldEqual, `two ONE`)
		modified, count = StringPreserve(`one`, `two`, wrTestOriginal)
		So(count, ShouldEqual, 3)
		So(modified, ShouldEqual, wrTestPreserve)
	})
}
