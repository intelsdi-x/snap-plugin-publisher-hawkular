// +build medium

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2017 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file excfpt in compliance with the License.
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
	"os"
	"testing"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKawkularPublisher(t *testing.T) {
	kp := NewHPublisher()

	host := os.Getenv("SNAP_HAWKULAR_HOST")
	cfg := plugin.Config{
		"tenant":   "snap",
		"server":   host,
		"port":     int64(8080),
		"user":     "jdoe",
		"password": "password",
	}

	Convey("Test Kawkular publisher", t, func() {
		Convey("Publish float 1", func() {
			metrics := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("x", "y", "z"),
					Config:    map[string]interface{}{"pw": "123aB"},
					Data:      5.7,
					Unit:      "float",
					Timestamp: time.Now(),
					Tags:      map[string]string{"y": "2"},
				},
			}
			err := kp.Publish(metrics, cfg)
			So(err, ShouldBeNil)
		})
		Convey("Publish float 2", func() {
			metrics := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("x", "x", "x"),
					Config:    map[string]interface{}{"pw": "abc123"},
					Data:      3.3,
					Tags:      map[string]string{"hello": "world"},
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			err := kp.Publish(metrics, cfg)
			So(err, ShouldBeNil)
		})
		Convey("Publish float 3", func() {
			metrics := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("y", "y", "3"),
					Config:    map[string]interface{}{"pw": "abc123"},
					Data:      3.4,
					Tags:      map[string]string{"tag2": "two"},
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			err := kp.Publish(metrics, cfg)
			So(err, ShouldBeNil)
		})
		Convey("Publish integer", func() {
			metrics := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("y", "y", "y"),
					Config:    map[string]interface{}{"pw": "abc1234"},
					Data:      1200,
					Tags:      map[string]string{"test": "test"},
					Unit:      "integer",
					Timestamp: time.Now(),
				},
			}
			err := kp.Publish(metrics, cfg)
			So(err, ShouldBeNil)
		})
		Convey("Publish a bool", func() {
			metrics := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("z", "z", "z"),
					Config:    map[string]interface{}{"pw": "abc123"},
					Data:      true,
					Tags:      map[string]string{"tag3": "three"},
					Unit:      "bool",
					Timestamp: time.Now(),
				},
			}
			err := kp.Publish(metrics, cfg)
			So(err, ShouldBeNil)
		})
		Convey("Publish a string", func() {
			metrics := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("o", "p", "q"),
					Config:    map[string]interface{}{"pw": "abc123"},
					Data:      "hawkular",
					Tags:      map[string]string{"tag3": "three"},
					Unit:      "string",
					Timestamp: time.Now(),
				},
			}
			err := kp.Publish(metrics, cfg)
			So(err, ShouldBeNil)
		})
	})
}
