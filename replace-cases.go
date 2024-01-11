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
	"unicode"

	"github.com/iancoleman/strcase"
)

// Case is a simple type for indicating a detected string case
type Case uint8

const (
	UnknownCase Case = iota
	LowerCase
	UpperCase
	CamelCase
	LowerCamelCase
	KebabCase
	ScreamingKebabCase
	SnakeCase
	ScreamingSnakeCase
)

// ProfileCase scans through each rune in the given input and checks for a
// number of interesting hints about the string case of the input
func ProfileCase(input string) (lower, upper, space, kebab, underscore bool) {
	for _, c := range input {
		if !lower && unicode.IsLower(c) {
			lower = true
		} else if !upper && unicode.IsUpper(c) {
			upper = true
		} else if !space && unicode.IsSpace(c) {
			space = true
		} else if !kebab && c == '-' {
			kebab = true
		} else if !underscore && c == '_' {
			underscore = true
		}
	}
	return
}

// DetectCase uses ProfileCase and some extra efforts to reliably discern the
// obvious string case of the given input
func DetectCase(input string) (detected Case) {
	lower, upper, space, kebab, underscore := ProfileCase(input)

	if space {
		detected = UnknownCase
		return // early out
	} else if kebab {
		// check for kebab things
		if input == strcase.ToKebab(input) {
			detected = KebabCase
			return
		} else if input == strcase.ToScreamingKebab(input) {
			detected = ScreamingKebabCase
			return
		}
	} else if underscore {
		// check for snake things
		if input == strcase.ToSnake(input) {
			detected = SnakeCase
			return
		} else if input == strcase.ToScreamingSnake(input) {
			detected = ScreamingSnakeCase
			return
		}
	} else if lower && upper {
		// check for camel things
		if input == strcase.ToCamel(input) {
			detected = CamelCase
			return
		} else if input == strcase.ToLowerCamel(input) {
			detected = LowerCamelCase
			return
		}
	}

	// check for default things
	if input == strings.ToLower(input) {
		detected = LowerCase
	} else if input == strings.ToUpper(input) {
		detected = UpperCase
	}
	return
}

// String returns the name of the string case, in that string cases' case
func (c Case) String() (name string) {
	switch c {
	case LowerCase:
		return "lower"
	case UpperCase:
		return "UPPER"
	case CamelCase:
		return "CamelCase"
	case LowerCamelCase:
		return "lowerCamelCase"
	case KebabCase:
		return "kebab-case"
	case ScreamingKebabCase:
		return "SCREAMING-KEBAB-CASE"
	case SnakeCase:
		return "snake_case"
	case ScreamingSnakeCase:
		return "SCREAMING_SNAKE_CASE"
	default:
		name = ""
	}
	return
}

// Apply returns the input text with the detected Case applied. For example:
//
//	CamelCase.Apply("kebab-thing") == "KebabThing"
//	KebabCase.Apply("CamelCase") == "camel-case"
func (c Case) Apply(input string) (modified string) {
	switch c {
	case LowerCase:
		modified = strings.ToLower(input)
	case UpperCase:
		modified = strings.ToUpper(input)
	case CamelCase:
		modified = strcase.ToCamel(input)
	case LowerCamelCase:
		modified = strcase.ToLowerCamel(input)
	case KebabCase:
		modified = strcase.ToKebab(input)
	case ScreamingKebabCase:
		modified = strcase.ToScreamingKebab(input)
	case SnakeCase:
		modified = strcase.ToSnake(input)
	case ScreamingSnakeCase:
		modified = strcase.ToScreamingSnake(input)
	case UnknownCase:
		fallthrough
	default:
		modified = input
	}
	return
}
