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

	"github.com/go-corelibs/diff"
)

// RegexFile uses Regex to ProcessFile
func RegexFile(search *regexp.Regexp, replace, target string) (original, modified string, count int, delta *diff.Diff, err error) {
	original, modified, count, delta, err = ProcessFile(target, func(original string) (modified string, count int) {
		modified, count = Regex(search, replace, original)
		return
	})
	return
}

// RegexLinesFile uses RegexLines to ProcessFile
func RegexLinesFile(search *regexp.Regexp, replace, target string) (original, modified string, count int, delta *diff.Diff, err error) {
	original, modified, count, delta, err = ProcessFile(target, func(original string) (modified string, count int) {
		modified, count = RegexLines(search, replace, original)
		return
	})
	return
}

// RegexPreserveFile uses StringPreserve to ProcessFile
func RegexPreserveFile(search *regexp.Regexp, replace, target string) (original, modified string, count int, delta *diff.Diff, err error) {
	original, modified, count, delta, err = ProcessFile(target, func(original string) (modified string, count int) {
		modified, count = RegexPreserve(search, replace, original)
		return
	})
	return
}
