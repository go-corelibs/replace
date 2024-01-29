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
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var (
	rxNil = (*regexp.Regexp)(nil)
)

func TestWorkMakers(t *testing.T) {
	convey.Convey("MakeRegexp", t, func() {
		rx, err := MakeRegexp(`[nope`, false, false, false)
		convey.So(err, convey.ShouldNotEqual, nil)
		convey.So(rx, convey.ShouldEqual, rxNil)
		rx, err = MakeRegexp(`true`, false, false, false)
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(rx, convey.ShouldNotEqual, rxNil)
		convey.So(rx.String(), convey.ShouldEqual, `true`)
		rx, err = MakeRegexp(`true`, true, false, false)
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(rx, convey.ShouldNotEqual, rxNil)
		convey.So(rx.String(), convey.ShouldEqual, `(?m)true`)
		rx, err = MakeRegexp(`true`, true, true, false)
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(rx, convey.ShouldNotEqual, rxNil)
		convey.So(rx.String(), convey.ShouldEqual, `(?ms)true`)
		rx, err = MakeRegexp(`true`, true, true, true)
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(rx, convey.ShouldNotEqual, rxNil)
		convey.So(rx.String(), convey.ShouldEqual, `(?msi)true`)
	})
}
