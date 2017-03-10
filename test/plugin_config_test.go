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

package test

import (
	"encoding/json"
	"testing"

	"github.com/intelsdi-x/snap-client-go/snap"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPluginConfigItems(t *testing.T) {
	c := snap.New(snap.ClientParams{URL: getUrl()})

	Convey("Testing GetPluginConfigItems", t, func() {
		Convey("Test get the config without parameters", func() {
			params := snap.NewGetPluginConfigItemParams()

			resp, err := c.GetPluginConfigItem(params)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Test get the config with parameters", func() {
			params := snap.NewGetPluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(2))

			resp, err := c.GetPluginConfigItem(params)
			So(resp.Payload, ShouldResemble, map[string]interface{}{"config": map[string]interface{}{}})
			So(err, ShouldBeNil)
		})
	})
}

func TestUpdatePluginConfigItems(t *testing.T) {
	c := snap.New(snap.ClientParams{URL: getUrl()})

	Convey("Testing UpdatePluginConfigItems", t, func() {
		Convey("Test set the plugin config without parameters", func() {
			params := snap.NewSetPluginConfigItemParams()

			resp, err := c.SetPluginConfigItem(params)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Test set the plugin config without a config", func() {
			params := snap.NewSetPluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(2))

			resp, err := c.SetPluginConfigItem(params)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Test set the plugin config with a config", func() {
			params := snap.NewSetPluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(1))
			cfg := `{"user":"jean","someint":1234567,"somefloat":3.1418,"somebool":false}`
			params.SetConfig(&cfg)

			resp, err := c.SetPluginConfigItem(params)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldResemble, map[string]interface{}{"config": map[string]interface{}{"somebool": false, "somefloat": json.Number("3.1418"), "someint": json.Number("1234567"), "user": "jean"}})
		})
	})
}

func TestDeletePluginConfigItems(t *testing.T) {
	c := snap.New(snap.ClientParams{URL: getUrl()})

	Convey("Testing DeletePluginConfigItems", t, func() {
		Convey("Test delete the plugin config without parameters", func() {
			params := snap.NewDeletePluginConfigItemParams()

			rep, err := c.DeletePluginConfigItem(params)
			So(err, ShouldNotBeNil)
			So(rep, ShouldBeNil)
		})

		Convey("Test delete the plugin config with parameters", func() {
			params := snap.NewDeletePluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(1))
			params.SetConfig([]string{"somefloat"})

			resp, err := c.DeletePluginConfigItem(params)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldNotBeNil)
		})
	})
}
