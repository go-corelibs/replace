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

	"github.com/go-corelibs/maps"
)

const (
	tRegexOriginal = `
One one ONE
Two two TWO
`
	tRegexOnlyOne = `
One two ONE
Two two TWO
`
	tRegexOnceOnes = `
two ONE
Two two TWO
`
	tRegexReplaced = `
two two two
Two two TWO
`
	tRegexPreserve = `
Two two TWO
Two two TWO
`
)

type tRegexTest struct {
	pattern *regexp.Regexp
	replace string
	matches int
	expects string
}

var (
	tRegexTestData = map[string]struct {
		fn   func(search *regexp.Regexp, replace, contents string) (modified string, count int)
		data map[string]tRegexTest
	}{
		"Regex": {
			fn: Regex,
			data: map[string]tRegexTest{
				"only one replaced": {regexp.MustCompile(`one`), `two`, 1, tRegexOnlyOne},
				"any one replaced":  {regexp.MustCompile(`(?i)one`), `two`, 3, tRegexReplaced},
				"none replaced":     {regexp.MustCompile(`(?i)none`), `two`, 0, tRegexOriginal},
				"nil search":        {nil, `two`, 0, tRegexOriginal},
			},
		},

		"RegexLines": {
			fn: RegexLines,
			data: map[string]tRegexTest{
				"only one replaced": {regexp.MustCompile(`one`), `two`, 1, tRegexOnlyOne},
				"any one replaced":  {regexp.MustCompile(`(?i)one`), `two`, 3, tRegexReplaced},
				"none replaced":     {regexp.MustCompile(`(?i)none`), `two`, 0, tRegexOriginal},
				"nil search":        {nil, `two`, 0, tRegexOriginal},
			},
		},

		"RegexPreserve": {
			fn: RegexPreserve,
			data: map[string]tRegexTest{
				"any one replaced":  {regexp.MustCompile(`(?i)one`), `two`, 3, tRegexPreserve},
				"same one replaced": {regexp.MustCompile(`(?i)one`), `one`, 3, tRegexOriginal},
				"ones replaced":     {regexp.MustCompile(`One one`), `two`, 1, tRegexOnceOnes},
				"nil replaced":      {nil, `one`, 0, tRegexOriginal},
			},
		},
	}
)

func TestRegex(t *testing.T) {

	for _, name := range maps.SortedKeys(tRegexTestData) {
		group := tRegexTestData[name]
		Convey(name, t, func() {
			for _, key := range maps.SortedKeys(group.data) {
				test := group.data[key]
				Convey(key, func() {
					modified, count := group.fn(test.pattern, test.replace, tRegexOriginal)
					So(count, ShouldEqual, test.matches)
					So(modified, ShouldEqual, test.expects)
				})
			}
		})
	}

}
