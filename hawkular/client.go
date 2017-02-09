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
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/hawkular/hawkular-client-go/metrics"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

type client struct {
	param         metrics.Parameters
	hclient       *metrics.Client
	needMetricDef bool
}

func newClient(p metrics.Parameters) client {
	c, err := metrics.NewHawkularClient(p)
	if err != nil {
		glog.Fatalln("Initiating hawkular client error: %s", err)
	}
	return client{param: p, hclient: c}
}

func (c client) toHawkular(mts []plugin.Metric) error {
	mhs := []metrics.MetricHeader{}

	for _, m := range mts {
		dp, ty, err := getDataAndType(m)
		if err != nil {
			return err
		}

		if GClient.needMetricDef {
			err := c.createMetricDefinition(ty, m.Namespace.Strings(), m.Tags)
			if err != nil {
				return err
			}
		}

		header := metrics.MetricHeader{
			ID:   strings.Join(m.Namespace.Strings(), "."),
			Data: []metrics.Datapoint{dp},
			Type: ty,
		}
		mhs = append(mhs, header)
		c.hclient.UpdateTags(ty, strings.Join(m.Namespace.Strings(), "."), m.Tags)
	}

	if len(mhs) > 0 {
		err := c.hclient.Write(mhs)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c client) createMetricDefinition(t metrics.MetricType, ns []string, tags map[string]string) error {
	md := metrics.MetricDefinition{
		Type: t,
		ID:   strings.Join(ns, "."),
		Tags: tags,
	}

	ok, err := c.hclient.Create(md)
	if !ok && err != nil {
		return err
	}
	return nil
}

func getDataAndType(m plugin.Metric) (metrics.Datapoint, metrics.MetricType, error) {
	var dp metrics.Datapoint
	f64, err := metrics.ConvertToFloat64(m.Data)
	if err != nil {
		switch ty := m.Data.(type) {
		case string:
			dp = metrics.Datapoint{Value: m.Data, Timestamp: m.Timestamp}
			return dp, metrics.String, nil
		case bool:
			dp = metrics.Datapoint{Value: strconv.FormatBool(m.Data.(bool)), Timestamp: m.Timestamp}
			return dp, metrics.String, nil
		default:
			glog.Warning("Metric type %v is not supported.", ty)
			return metrics.Datapoint{}, metrics.Generic, fmt.Errorf("Metric type %v is not supported.", ty)
		}
	}
	dp = metrics.Datapoint{Value: f64, Timestamp: m.Timestamp}
	return dp, metrics.Gauge, nil
}
