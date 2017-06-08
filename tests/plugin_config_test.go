// +build medium

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

package tests

import (
	"encoding/json"
	"testing"

	"github.com/intelsdi-x/snap-client-go/client/plugins"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPluginConfigItems(t *testing.T) {
	c := getSnapClient()

	Convey("Testing GetPluginConfigItems", t, func() {
		Convey("Test get the config without parameters", func() {
			params := plugins.NewGetPluginConfigItemParams()

			resp, err := c.Plugins.GetPluginConfigItem(params, nil)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Test get the config with parameters", func() {
			params := plugins.NewGetPluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(2))

			resp, err := c.Plugins.GetPluginConfigItem(params, nil)
			So(resp.Payload, ShouldResemble, map[string]interface{}{})
			So(err, ShouldBeNil)
		})
	})
}

func TestUpdatePluginConfigItems(t *testing.T) {
	c := getSnapClient()

	Convey("Testing UpdatePluginConfigItems", t, func() {
		Convey("Test set the plugin config without parameters", func() {
			params := plugins.NewSetPluginConfigItemParams()

			resp, err := c.Plugins.SetPluginConfigItem(params, nil)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Test set the plugin config without a config", func() {
			params := plugins.NewSetPluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(2))

			resp, err := c.Plugins.SetPluginConfigItem(params, nil)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Test set the plugin config with a config", func() {
			params := plugins.NewSetPluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(1))
			cfg := map[string]interface{}{"user": "jean", "someint": 1234567, "somefloat": 3.1418, "somebool": false}
			params.SetConfig(cfg)

			resp, err := c.Plugins.SetPluginConfigItem(params, nil)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldResemble, map[string]interface{}{"somebool": false, "somefloat": json.Number("3.1418"), "someint": json.Number("1234567"), "user": "jean"})
		})
	})
}

func TestDeletePluginConfigItems(t *testing.T) {
	c := getSnapClient()

	Convey("Testing DeletePluginConfigItems", t, func() {
		Convey("Test delete the plugin config without parameters", func() {
			params := plugins.NewDeletePluginConfigItemParams()

			rep, err := c.Plugins.DeletePluginConfigItem(params, nil)
			So(err, ShouldNotBeNil)
			So(rep, ShouldBeNil)
		})

		Convey("Test delete the plugin config with parameters", func() {
			params := plugins.NewDeletePluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(1))
			cfg := []string{"somefloat"}
			params.SetConfig(cfg)

			resp, err := c.Plugins.DeletePluginConfigItem(params, nil)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldNotBeNil)
		})
	})
}
