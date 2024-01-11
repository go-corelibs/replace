// Copyright (c) 2024  The Go-CoreLibs Authors
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

// Package replace provides text replacement utilities
package replace

import (
	"regexp"
	"strings"
	"unicode"
)

// Regex counts the number of matches and runs ReplaceAllString on the given
// contents, if there are no matches, `modified` will be the same as `contents`
func Regex(search *regexp.Regexp, replace, contents string) (modified string, count int) {
	if search != nil {
		m := search.FindAllString(contents, -1)
		if count = len(m); count > 0 {
			modified = search.ReplaceAllString(contents, replace)
			return
		}
	}
	modified = contents
	return
}

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

// StringPreserve is like StringInsensitive except that StringPreserve attempts
// to preserve the per-instance case when doing replacements. If `search` or
// `replace` have any spaces, StringPreserve is wraps String.
//
// To preserve the per-instance cases, each instance has DetectCase run and
// uses Case.Apply to derive the `replace` value actually used
//
// See the Case constants for the list of string cases supported
func StringPreserve(search, replace, contents string) (modified string, count int) {
	// See: https://stackoverflow.com/questions/31348919/case-insensitive-string-replace-in-go/76512289

	if search != "" && search != replace {
		if strings.ContainsFunc(search+replace, func(r rune) bool {
			return unicode.IsSpace(r)
		}) {
			modified, count = String(search, replace, contents)
			return
		}

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
				c := DetectCase(found)
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

func RegexPreserve(search *regexp.Regexp, replace, contents string) (modified string, count int) {
	if search != nil {

		if strings.ContainsFunc(search.String()+replace, func(r rune) bool {
			return unicode.IsSpace(r)
		}) {
			modified, count = Regex(search, replace, contents)
			return
		}

		m := search.FindAllString(contents, -1)
		if count = len(m); count > 0 {

			var buffer strings.Builder

			var start int
			for i := 0; i < count; i++ {
				j := start
				// find next index
				j += strings.Index(contents[start:], m[i])
				// write non-match contents
				buffer.WriteString(contents[start:j])
				// derive replacement value
				c := DetectCase(m[i])
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
