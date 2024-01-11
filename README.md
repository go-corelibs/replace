[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/replace)
[![codecov](https://codecov.io/gh/go-corelibs/replace/graph/badge.svg?token=VYMEpT55Gx)](https://codecov.io/gh/go-corelibs/replace)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/replace)](https://goreportcard.com/report/github.com/go-corelibs/replace)

# replace - text replacement utilities

# Installation

``` shell
> go get github.com/go-corelibs/replace@latest
```

# Examples

## String, StringInsensitive, StringPreserve

``` go
var contents := `First line of text says something.
Text on the second line says more.`

func main() {
    // replace "text" (case-sensitively) with "this"
    modified, count := replace.String("text", "this", contents)
    // count == 1
    // modified == "First line of this says something.\nText on the second line says more."

    // replace "text" (case-insensitively) with "this"
    modified, count := replace.StringInsensitive("text", "this", contents)
    // count == 2
    // modified == "First line of this says something.\nthis on the second line says more."

    // replace "text" (case-preserved) with "this"
    modified, count := replace.StringPreserve("text", "this", contents)
    // count == 2
    // modified == "First line of this says something.\nThis on the second line says more."
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
