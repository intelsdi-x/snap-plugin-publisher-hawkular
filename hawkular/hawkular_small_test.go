// +build small

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2017 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package hawkular

import (
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHawkularPlugin(t *testing.T) {
	hp := NewHPublisher()

	Convey("Create HawkularPublisher", t, func() {
		Convey("So hp should not be nil", func() {
			So(hp, ShouldNotBeNil)
		})
		Convey("So hp should be of HPublisher type", func() {
			So(hp, ShouldHaveSameTypeAs, HPublisher{})
		})
	})

	Convey("Get Config Policy", t, func() {
		cp, err := hp.GetConfigPolicy()
		Convey("So config should not be empty", func() {
			So(cp, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
	})

	Convey("Get Parameters", t, func() {
		p := getParameters(plugin.Config{
			"tenant":      "snap",
			"user":        "jdoe",
			"password":    "password",
			"concurrency": int64(12),
			"server":      "localhost",
			"port":        int64(8080),
		})

		Convey("check url", func() {
			So(p.Url, ShouldEqual, "http://localhost:8080")
		})

		Convey("check user", func() {
			So(p.Username, ShouldEqual, "jdoe")
		})

		Convey("check password", func() {
			So(p.Password, ShouldEqual, "password")
		})

		Convey("check tenant", func() {
			So(p.Tenant, ShouldEqual, "snap")
		})

		Convey("check concurrency", func() {
			So(p.Concurrency, ShouldEqual, 12)
		})
	})
}
