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

func TestVars(t *testing.T) {
	Convey("Empty input", t, func() {
		So(Vars("", map[string]string{}), ShouldEqual, "")
	})

	Convey("Plain input", t, func() {
		So(Vars("Plain text input", map[string]string{}), ShouldEqual, "Plain text input")
	})

	Convey("Plain vars", t, func() {
		So(Vars("Plain $text input", map[string]string{"text": "string"}), ShouldEqual, "Plain string input")
		So(Vars("Plain $ text input", map[string]string{"text": "string"}), ShouldEqual, "Plain $ text input")
		So(Vars("Plain $_text input", map[string]string{"_text": "string"}), ShouldEqual, "Plain string input")
		So(Vars("Plain $_text$ input", map[string]string{"_text": "string"}), ShouldEqual, "Plain string$ input")
	})

	Convey("Brace vars", t, func() {
		So(Vars("Plain ${text} input", map[string]string{"text": "string"}), ShouldEqual, "Plain string input")
		So(Vars("Plain ${ text} input", map[string]string{"text": "string"}), ShouldEqual, "Plain ${ text} input")
		So(Vars("Plain $_text input", map[string]string{"_text": "string"}), ShouldEqual, "Plain string input")
		So(Vars("Plain $_text$ input", map[string]string{"_text": "string"}), ShouldEqual, "Plain string$ input")
		So(Vars("Plain ${_text$ input", map[string]string{"_text": "string"}), ShouldEqual, "Plain ${_text$ input")
	})

	Convey("Corner cases", t, func() {
		So(Vars("Plain ${text} input $text", map[string]string{"text": "string"}), ShouldEqual, "Plain string input string")
		So(Vars("$text plain ${text} input $text", map[string]string{"text": "string"}), ShouldEqual, "string plain string input string")
		So(Vars("$text", map[string]string{"text": "string"}), ShouldEqual, "string")
		So(Vars("${text}", map[string]string{"text": "string"}), ShouldEqual, "string")
	})
}
