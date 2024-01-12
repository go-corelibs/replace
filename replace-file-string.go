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
	"github.com/go-corelibs/diff"
)

// StringFile uses String to ProcessFile
func StringFile(search, replace, target string) (original, modified string, count int, delta *diff.Diff, err error) {
	original, modified, count, delta, err = ProcessFile(target, func(original string) (modified string, count int) {
		modified, count = String(search, replace, original)
		return
	})
	return
}

// StringInsensitiveFile uses StringInsensitive to ProcessFile
func StringInsensitiveFile(search, replace, target string) (original, modified string, count int, delta *diff.Diff, err error) {
	original, modified, count, delta, err = ProcessFile(target, func(original string) (modified string, count int) {
		modified, count = StringInsensitive(search, replace, original)
		return
	})
	return
}

// StringPreserveFile uses StringPreserve to ProcessFile
func StringPreserveFile(search, replace, target string) (original, modified string, count int, delta *diff.Diff, err error) {
	original, modified, count, delta, err = ProcessFile(target, func(original string) (modified string, count int) {
		modified, count = StringPreserve(search, replace, original)
		return
	})
	return
}
