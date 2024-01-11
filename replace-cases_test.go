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

func TestCases(t *testing.T) {
	Convey("ProfileCase", t, func() {
		lower, upper, space, kebab, underscore := ProfileCase(`lU -_`)
		So(lower, ShouldEqual, true)
		So(upper, ShouldEqual, true)
		So(space, ShouldEqual, true)
		So(kebab, ShouldEqual, true)
		So(underscore, ShouldEqual, true)
		lower, upper, space, kebab, underscore = ProfileCase(`!.;'"/`)
		So(lower, ShouldEqual, false)
		So(upper, ShouldEqual, false)
		So(space, ShouldEqual, false)
		So(kebab, ShouldEqual, false)
		So(underscore, ShouldEqual, false)
	})

	Convey("DetectCase", t, func() {
		So(DetectCase("has space"), ShouldEqual, UnknownCase)
		So(DetectCase("lower"), ShouldEqual, LowerCase)
		So(DetectCase("UPPER"), ShouldEqual, UpperCase)
		So(DetectCase("CamelCase"), ShouldEqual, CamelCase)
		So(DetectCase("lowerCamelCase"), ShouldEqual, LowerCamelCase)
		So(DetectCase("kebab-case"), ShouldEqual, KebabCase)
		So(DetectCase("SCREAMING-SNAKE-CASE"), ShouldEqual, ScreamingKebabCase)
		So(DetectCase("snake_case"), ShouldEqual, SnakeCase)
		So(DetectCase("SCREAMING_SNAKE_CASE"), ShouldEqual, ScreamingSnakeCase)
	})

	Convey("Case.String", t, func() {
		So(UnknownCase.String(), ShouldEqual, "")
		So(LowerCase.String(), ShouldEqual, "lower")
		So(UpperCase.String(), ShouldEqual, "UPPER")
		So(CamelCase.String(), ShouldEqual, "CamelCase")
		So(LowerCamelCase.String(), ShouldEqual, "lowerCamelCase")
		So(KebabCase.String(), ShouldEqual, "kebab-case")
		So(ScreamingKebabCase.String(), ShouldEqual, "SCREAMING-KEBAB-CASE")
		So(SnakeCase.String(), ShouldEqual, "snake_case")
		So(ScreamingSnakeCase.String(), ShouldEqual, "SCREAMING_SNAKE_CASE")
	})

	Convey("Case.Apply", t, func() {
		So(UnknownCase.Apply("has space"), ShouldEqual, "has space")
		So(LowerCase.Apply("Lower"), ShouldEqual, "lower")
		So(UpperCase.Apply("uPPER"), ShouldEqual, "UPPER")
		So(CamelCase.Apply("camelCase"), ShouldEqual, "CamelCase")
		So(LowerCamelCase.Apply("LowerCamelCase"), ShouldEqual, "lowerCamelCase")
		So(KebabCase.Apply("kebab_case"), ShouldEqual, "kebab-case")
		So(ScreamingKebabCase.Apply("SCREAMING_KEBAB_CASE"), ShouldEqual, "SCREAMING-KEBAB-CASE")
		So(SnakeCase.Apply("snake-case"), ShouldEqual, "snake_case")
		So(ScreamingSnakeCase.Apply("SCREAMING-SNAKE-CASE"), ShouldEqual, "SCREAMING_SNAKE_CASE")
	})
}
