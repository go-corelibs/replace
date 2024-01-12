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
	"os"

	"github.com/go-corelibs/diff"
)

// ProcessFile is a convenience function which calls os.ReadFile on the
// target and runs the given `fn` with the string contents and creates a
// Diff of the changes between the original and the modified output
func ProcessFile(target string, fn func(original string) (modified string, count int)) (original, modified string, count int, delta *diff.Diff, err error) {
	var data []byte
	if data, err = os.ReadFile(target); err == nil {
		original = string(data)
		modified, count = fn(original)
		delta = diff.New(target, original, modified)
	}
	return
}
