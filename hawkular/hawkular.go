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
	"crypto/tls"
	"fmt"

	"github.com/golang/glog"
	"github.com/hawkular/hawkular-client-go/metrics"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

var (
	// GClient is the global Hawkular client
	GClient = client{}
)

// HPublisher the struct type of Hawkular publisher
type HPublisher struct {
}

// NewHPublisher returns a Kawkular publisher instance
func NewHPublisher() HPublisher {
	return HPublisher{}
}

// GetConfigPolicy returns an instance of publisher config policy
func (h HPublisher) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	err := policy.AddNewStringRule([]string{}, "server", true)
	if err != nil {
		return plugin.ConfigPolicy{}, err
	}

	err = policy.AddNewStringRule([]string{}, "scheme", false, plugin.SetDefaultString("http"))
	if err != nil {
		return plugin.ConfigPolicy{}, err
	}

	err = policy.AddNewBoolRule([]string{}, "insecureSkipVerify", false, plugin.SetDefaultBool(true))
	if err != nil {
		return plugin.ConfigPolicy{}, err
	}

	err = policy.AddNewIntRule([]string{}, "port", false, plugin.SetDefaultInt(8080))
	if err != nil {
		return plugin.ConfigPolicy{}, err
	}

	err = policy.AddNewStringRule([]string{}, "user", false, plugin.SetDefaultString("jdoe"))
	if err != nil {
		return plugin.ConfigPolicy{}, err
	}

	err = policy.AddNewStringRule([]string{}, "password", false, plugin.SetDefaultString("password"))
	if err != nil {
		return plugin.ConfigPolicy{}, err
	}

	err = policy.AddNewStringRule([]string{}, "tenant", false, plugin.SetDefaultString("snap"))
	if err != nil {
		return plugin.ConfigPolicy{}, err
	}

	err = policy.AddNewIntRule([]string{}, "concurrency", false, plugin.SetDefaultInt(10))
	if err != nil {
		return plugin.ConfigPolicy{}, err
	}
	return *policy, nil
}

// Publish publishes Snap metrics to Kawkular
func (h HPublisher) Publish(mts []plugin.Metric, cfg plugin.Config) error {
	if GClient.hclient == nil {
		param := getParameters(cfg)
		GClient = newClient(param)
		GClient.needMetricDef = true
	}

	err := GClient.toHawkular(mts)
	if err != nil {
		return err
	}

	if GClient.needMetricDef {
		GClient.needMetricDef = false
	}
	return nil
}

func getParameters(cfg plugin.Config) metrics.Parameters {
	server, err := cfg.GetString("server")
	if err != nil {
		glog.Fatalf("No server defined: %s", err)
	}

	scheme, err := cfg.GetString("scheme")
	if err != nil {
		scheme = "http"
	}

	insecureSkipVerify, err := cfg.GetBool("insecureSkipVerify")
	if err != nil {
		insecureSkipVerify = true
	}

	port, err := cfg.GetInt("port")
	if err != nil {
		glog.Infof("No port defined: %s", err)
	}
	user, err := cfg.GetString("user")
	if err != nil {
		glog.Infof("No user defined: %s", err)
	}

	pw, err := cfg.GetString("password")
	if err != nil {
		glog.Infof("No password defined: %s", err)
	}

	tenant, err := cfg.GetString("tenant")
	if err != nil {
		glog.Infof("No tenant defined: %s", err)
	}

	concur, err := cfg.GetInt("concurrency")
	if err != nil {
		glog.Infof("No concurrency defined: %s", err)
	}

	tlsC := &tls.Config{InsecureSkipVerify: insecureSkipVerify}

	p := metrics.Parameters{
		Tenant:      tenant,
		Url:         fmt.Sprintf(scheme+"://%s:%d", server, port),
		Username:    user,
		Password:    pw,
		Concurrency: int(concur),
		TLSConfig:   tlsC,
	}
	return p
}
