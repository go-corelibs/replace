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
	"unicode"
)

// Vars searches through `input` for variables in the form of `$Name` or
// `${Name}` and replaces them with the corresponding `replacements` value.
// Missing keys are replaced with empty strings
//
// Vars follows the POSIX [3.235 Name] definition:
//
//	a word consisting solely of underscores, digits, and alphabetics from
//	the portable character set. The first character of a name is not a digit
//
// [3.235 Name]: https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap03.html#tag_03_231
func Vars(input string, replacements map[string]string) (expanded string) {

	s := struct {
		opened bool
		braced bool
		key    string
		source string
	}{}

	for _, r := range input {
		char := string(r)

		if s.opened {
			// dollar sign was previously found
			s.source += char // always track the original text
			keyLen := len(s.key)

			if s.braced {
				// dollar sign was followed by an opening brace, must have a
				// closing brace in order for this to be an actual variable
				// instance
				if keyLen == 0 {
					if isVarFirstCharacter(r) {
						// valid first variable name character detected
						s.key += char
						continue
					}
					// not a valid variable name
					expanded += s.source
				} else if isVarOtherCharacter(r) {
					// valid subsequent variable name character detected
					s.key += char
					continue
				} else if r == '}' {
					// closing brace detected
					if v, ok := replacements[s.key]; ok {
						// replacement value ready
						expanded += v
					}
				} else {
					// this is not a valid variable instance
					expanded += s.source
				}
				// cleanup the current state
				s.opened = false
				s.braced = false
				s.key = ""
				continue
			}

			if keyLen == 0 {
				if r == '{' {
					s.braced = true
					continue
				} else if isVarFirstCharacter(r) {
					// valid first variable name character detected
					s.key += char
					continue
				}
				// this is not a valid variable instance
				expanded += s.source
			} else if isVarOtherCharacter(r) {
				// valid subsequent variable name character detected
				s.key += char
				continue
			}

			if v, ok := replacements[s.key]; ok {
				// replacement value ready
				expanded += v
				expanded += char
			}

			// cleanup the current state
			s.opened = false
			s.braced = false
			s.key = ""
			continue
		}

		if s.opened = r == '$'; s.opened {
			// variable opening detected
			s.source += char
			continue
		}

		// actual content
		expanded += char
	}

	return
}

func isVarFirstCharacter(r rune) (valid bool) {
	valid = r == '_' || unicode.IsLower(r)
	return
}

func isVarOtherCharacter(r rune) (valid bool) {
	valid = r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
	return
}
