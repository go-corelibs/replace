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

	"github.com/go-corelibs/maps"
)

const (
	tStringOriginal = `
One one ONE
Two two TWO
`
	tStringOnlyOne = `
One two ONE
Two two TWO
`
	tStringOnceOnes = `
two ONE
Two two TWO
`
	tStringReplaced = `
two two two
Two two TWO
`
	tStringPreserve = `
Two two TWO
Two two TWO
`
)

type tStringTest struct {
	pattern string
	replace string
	matches int
	expects string
}

var (
	tStringTestData = map[string]struct {
		fn   func(search string, replace, contents string) (modified string, count int)
		data map[string]tStringTest
	}{
		"String": {
			fn: String,
			data: map[string]tStringTest{
				"only one replaced": {`one`, `two`, 1, tStringOnlyOne},
				"none replaced":     {`none`, `two`, 0, tStringOriginal},
				"same search":       {"two", `two`, 0, tStringOriginal},
				"nil search":        {"", `two`, 0, tStringOriginal},
			},
		},

		"StringInsensitive": {
			fn: StringInsensitive,
			data: map[string]tStringTest{
				"any one replaced": {`one`, `two`, 3, tStringReplaced},
				"none replaced":    {`none`, `two`, 0, tStringOriginal},
				"nil search":       {"", `two`, 0, tStringOriginal},
			},
		},

		"StringPreserve": {
			fn: StringPreserve,
			data: map[string]tStringTest{
				"any one replaced":  {`one`, `two`, 3, tStringPreserve},
				"same one replaced": {`one`, `one`, 0, tStringOriginal},
				"ones replaced":     {`One one`, `two`, 1, tStringOnceOnes},
				"nil replaced":      {"", `one`, 0, tStringOriginal},
			},
		},
	}
)

func TestString(t *testing.T) {

	for _, name := range maps.SortedKeys(tStringTestData) {
		group := tStringTestData[name]
		Convey(name, t, func() {
			for _, key := range maps.SortedKeys(group.data) {
				test := group.data[key]
				Convey(key, func() {
					modified, count := group.fn(test.pattern, test.replace, tStringOriginal)
					So(count, ShouldEqual, test.matches)
					So(modified, ShouldEqual, test.expects)
				})
			}
		})
	}

}
