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
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/go-corelibs/globs"
	"github.com/go-corelibs/path"
)

var (
	ErrLargeFile    = errors.New("large file")
	ErrBinaryFile   = errors.New("binary file")
	ErrTooManyFiles = fmt.Errorf("too many files")
)

var (
	MaxFileSize  = int64(math.Round(1024.0 * 1024.0 * 5.0))
	MaxFileCount = 1000000
)

// FindAllMatcherFn is the function signature for custom matching of content
type FindAllMatcherFn func(data []byte) (matched bool)

// FindAllMatchingFn is the function signature for custom tracking of the
// finding process
type FindAllMatchingFn func(file string, matched bool, err error)

// IsIncluded returns true if the given `input` string is included or is not
// explicitly excluded, or if `include` and `exclude` are nil or empty
func IsIncluded(include, exclude globs.Globs, input string) (included bool) {
	eCount, iCount := len(exclude), len(include)

	if included = eCount == 0 && iCount == 0; included {
		// no constraints present, input allowed
		return
	}

	if eCount > 0 && exclude.Match(input) {
		// excluded
		return
	}

	// not excluded, check if includes are constrained
	if included = iCount == 0; included {
		// and no include constraints
		return
	}
	included = include.Match(input)
	return
}

// FindAllIncluded walks the given target paths, looking for unique IsIncluded
// files
func FindAllIncluded(targets []string, includeHidden, noLimit, binAsText, recurse bool, include, exclude globs.Globs) (found []string) {
	unique := make(map[string]struct{})
	check := func(target string) (allowed bool) {
		if _, present := unique[target]; present {
			return
		} else if !includeHidden && path.IsHidden(target) {
			return
		}
		allowed = IsIncluded(include, exclude, target)
		unique[target] = struct{}{} // don't check this target again
		return
	}
	for _, target := range targets {
		if path.IsFile(target) {
			// process file path
			if check(target) {
				found = append(found, target)
			}
		} else if recurse && path.IsDir(target) {
			// process dir path
			files, _ := path.ListFiles(target, includeHidden)
			for _, file := range files {
				if check(file) {
					found = append(found, file)
				}
			}
			dirs, _ := path.ListDirs(target, includeHidden)
			more := FindAllIncluded(dirs, includeHidden, noLimit, binAsText, recurse, include, exclude)
			found = append(found, more...)
		}
	}
	return
}

// FindAllMatcher uses FindAllIncluded to derive a list of `files` and
// a list of `matches` (using the `matcher` func), returning ErrTooManyFiles
// if the total number of files exceeds the MaxFileCount. While performing the
// find process, calls the given `fn` to report progress and the state of each
// file processed
func FindAllMatcher(targets []string, includeHidden, noLimit, binAsText, recurse bool, include, exclude globs.Globs, fn FindAllMatchingFn, matcher FindAllMatcherFn) (files, matches []string, err error) {
	if fn == nil {
		fn = func(file string, matched bool, err error) {}
	}
	for _, target := range FindAllIncluded(targets, includeHidden, noLimit, binAsText, recurse, include, exclude) {
		files = append(files, target)
		if len(files) > MaxFileCount {
			err = ErrTooManyFiles
			return
		}
		var ee error
		var data []byte
		var matched bool
		if !noLimit && path.FileSize(target) > MaxFileSize {
			ee = ErrLargeFile
		} else if !binAsText && !path.IsPlainText(target) {
			ee = ErrBinaryFile
		} else if data, ee = os.ReadFile(target); ee == nil {
			if matched = matcher(data); matched {
				matches = append(matches, target)
			}
		}
		fn(target, matched, ee)
	}
	return
}

// FindAllMatchingRegexp is a wrapper around FindAllMatcher with a custom
// matcher func which uses `search.Match` to filter the `matches` list
func FindAllMatchingRegexp(search *regexp.Regexp, targets []string, includeHidden, noLimit, binAsText, recurse bool, include, exclude globs.Globs, fn FindAllMatchingFn) (files, matches []string, err error) {
	files, matches, err = FindAllMatcher(targets, includeHidden, noLimit, binAsText, recurse, include, exclude, fn, func(data []byte) (matched bool) {
		matched = search.Match(data)
		return
	})
	return
}

// FindAllMatchingRegexpLines is a wrapper around FindAllMatcher with a custom
// matcher func which uses `search.Match` on each line of each file to filter
// the `matches` list
func FindAllMatchingRegexpLines(search *regexp.Regexp, targets []string, includeHidden, noLimit, binAsText, recurse bool, include, exclude globs.Globs, fn FindAllMatchingFn) (files, matches []string, err error) {
	files, matches, err = FindAllMatcher(targets, includeHidden, noLimit, binAsText, recurse, include, exclude, fn, func(data []byte) (matched bool) {
		lines := strings.Split(string(data), "\n")
		last := len(lines) - 1
		for idx, line := range lines {
			if idx < last {
				line += "\n"
			}
			if matched = search.MatchString(line); matched {
				return
			}
		}
		return
	})
	return
}

// FindAllMatchingString is a wrapper around FindAllMatcher with a custom
// matcher func which uses [strings.Contains] to filter the `matches` list
func FindAllMatchingString(search string, targets []string, includeHidden, noLimit, binAsText, recurse bool, include, exclude globs.Globs, fn FindAllMatchingFn) (files, matches []string, err error) {
	files, matches, err = FindAllMatcher(targets, includeHidden, noLimit, binAsText, recurse, include, exclude, fn, func(data []byte) (matched bool) {
		matched = strings.Contains(string(data), search)
		return
	})
	return
}

// FindAllMatchingStringInsensitive is a wrapper around FindAllMatcher with a
// custom matcher func which uses [strings.Contains], in a case-insensitive
// way, to filter the `matches` list
func FindAllMatchingStringInsensitive(search string, targets []string, includeHidden, noLimit, binAsText, recurse bool, include, exclude globs.Globs, fn FindAllMatchingFn) (files, matches []string, err error) {
	files, matches, err = FindAllMatcher(targets, includeHidden, noLimit, binAsText, recurse, include, exclude, fn, func(data []byte) (matched bool) {
		matched = strings.Contains(strings.ToLower(string(data)), strings.ToLower(search))
		return
	})
	return
}
