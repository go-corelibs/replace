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
	"strings"

	"github.com/go-corelibs/strcases"
)

// Regex counts the number of matches and runs ReplaceAllString on the given
// contents, if there are no matches, `modified` will be the same as `contents`
func Regex(search *regexp.Regexp, replace, contents string) (modified string, count int) {
	if search != nil {
		if count = len(search.FindAllString(contents, -1)); count > 0 {
			modified = search.ReplaceAllString(contents, replace)
			return
		}
	}
	modified = contents
	return
}

// RegexLines is like Regex with the exception that the contents are split
// into a list of lines and `search` is applied to each line individually
func RegexLines(search *regexp.Regexp, replace, contents string) (modified string, count int) {
	if search != nil {
		var replaced []string
		lines := strings.Split(contents, "\n")
		last := len(lines) - 1
		for idx, line := range lines {
			if idx < last {
				line += "\n"
			}
			if num := len(search.FindAllString(line, -1)); num > 0 {
				count += num
				replaced = append(replaced, search.ReplaceAllString(line, replace))
			} else {
				replaced = append(replaced, line)
			}
		}
		if len(replaced) > 0 {
			modified = strings.Join(replaced, "")
			return
		}
	}
	modified = contents
	return
}

// RegexPreserve is similar to StringPreserve except that it works with
// regular expressions to perform the search and replacement process.
//
// While StringPreserve can easily detect un-case-detectable inputs, due to
// the variable nature of regular expressions it is assumed that the developer
// using RegexPreserve is confident that the `search` and `replace` arguments
// result in case-detectable string replacements.
func RegexPreserve(search *regexp.Regexp, replace, contents string) (modified string, count int) {
	if search != nil {

		m := search.FindAllString(contents, -1)
		if count = len(m); count > 0 {

			d := strcases.NewCaseDetector()
			var buffer strings.Builder

			var start int
			for i := 0; i < count; i++ {
				j := start
				// find next index
				j += strings.Index(contents[start:], m[i])
				// write non-match contents
				buffer.WriteString(contents[start:j])
				// derive replacement value
				c := d.Detect(m[i])
				// replacement my contain regex goodness, so must
				// call search.ReplaceAllString to get the corrent
				// results
				replaced := search.ReplaceAllString(m[i], replace)
				// write modified replacement
				buffer.WriteString(c.Apply(replaced))
				// move the start point
				start = j + len(m[i])
			}
			buffer.WriteString(contents[start:])
			modified = buffer.String()

			return
		}

	}
	modified = contents
	return
}
