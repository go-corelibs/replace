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
	"strings"

	"github.com/go-corelibs/strcases"
)

// String counts the number of matches and runs strings.ReplaceAll on the
// given contents, if there are no matches, `modified` will be the same as
// `contents`
func String(search, replace, contents string) (modified string, count int) {
	if search != "" && search != replace {
		if count = strings.Count(contents, search); count > 0 {
			modified = strings.ReplaceAll(contents, search, replace)
			return
		}
	}
	modified = contents
	return
}

// StringInsensitive counts the number of case-insensitive matches and
// replaces each instance with the `replace` value, if there are no matches
// `modified` will be the same as `contents`
func StringInsensitive(search, replace, contents string) (modified string, count int) {
	// See: https://stackoverflow.com/questions/31348919/case-insensitive-string-replace-in-go/76512289

	if search != "" && search != replace {
		lowerContents := strings.ToLower(contents)
		lowerSearch := strings.ToLower(search)

		if count = strings.Count(lowerContents, lowerSearch); count > 0 {
			var buffer strings.Builder
			buffer.Grow(len(contents) + (count * (len(replace) - len(search))))

			var start int
			for i := 0; i < count; i++ {
				j := start
				j += strings.Index(lowerContents[start:], lowerSearch)
				buffer.WriteString(contents[start:j])
				buffer.WriteString(replace)
				start = j + len(search)
			}
			buffer.WriteString(contents[start:])
			modified = buffer.String()
			return
		}

	}

	modified = contents
	return
}

// StringPreserve attempts to preserve the case of per-replacement matches.
// Not all strings can have their cases easily detected and StringPreserve
// uses CanPreserve to check if the `search` and `replace` arguments are
// simple enough for DetectCase to discern the case patterns. If
// StringPreserve can't preserve, it defaults to returning the results of
// calling String with the same arguments given to StringPreserve.
//
// To preserve the per-instance cases, each instance has DetectCase run and
// uses Case.Apply to derive the `replace` value actually used.
//
// See the Case constants for the list of string cases supported.
func StringPreserve(search, replace, contents string) (modified string, count int) {
	// See: https://stackoverflow.com/questions/31348919/case-insensitive-string-replace-in-go/76512289

	if search != "" && search != replace {
		if !strcases.CanPreserve(search + replace) {
			modified, count = String(search, replace, contents)
			return
		}

		d := strcases.NewCaseDetector()
		lowerContents := strings.ToLower(contents)
		lowerSearch := strings.ToLower(search)
		searchSize := len(search)

		if count = strings.Count(lowerContents, lowerSearch); count > 0 {
			var buffer strings.Builder
			buffer.Grow(len(contents) + (count * (len(replace) - searchSize)))

			var start int
			for i := 0; i < count; i++ {
				j := start
				// find next index
				j += strings.Index(lowerContents[start:], lowerSearch)
				// write non-match contents
				buffer.WriteString(contents[start:j])
				// derive replacement value
				found := contents[j : j+searchSize]
				c := d.Detect(found)
				// write modified replacement
				buffer.WriteString(c.Apply(replace))
				// move the start point
				start = j + searchSize
			}
			buffer.WriteString(contents[start:])
			modified = buffer.String()
			return
		}

	}

	modified = contents
	return
}
