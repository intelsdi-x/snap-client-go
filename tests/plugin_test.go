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
	"os"
	"testing"

	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/intelsdi-x/snap-client-go/client/plugins"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPlugins(t *testing.T) {
	c := getSnapClient()

	Convey("Testing GetPlugins", t, func() {
		Convey("Test get a list of plugins", func() {
			params := plugins.NewGetPluginsParams()

			resp, err := c.Plugins.GetPlugins(params, openapiclient.BasicAuth("snap", "snap"))
			So(err, ShouldBeNil)
			So(resp.Payload.Plugins, ShouldNotBeNil)
		})

		Convey("Test get a giving plugin", func() {
			params := plugins.NewGetPluginParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(2))

			resp, err := c.Plugins.GetPlugin(params, nil)
			So(err, ShouldBeNil)
			So(resp.Payload.Name, ShouldEqual, "mock")
		})

		Convey("Test get a non-existing plugin", func() {
			params := plugins.NewGetPluginParams()
			params.SetPname("mock1")
			params.SetPtype("collector")
			params.SetPversion(int64(1))

			resp, err := c.Plugins.GetPlugin(params, nil)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})
	})
}

func TestUnloadPlugin(t *testing.T) {
	c := getSnapClient()

	Convey("Testing Unload a plugin", t, func() {
		Convey("Test unload an existing plugin", func() {
			params := plugins.NewUnloadPluginParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(1))

			resp, err := c.Plugins.UnloadPlugin(params, openapiclient.BasicAuth("snap", "snap"))
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Test unload a non-existing plugin", func() {
			params := plugins.NewUnloadPluginParams()
			params.SetPname("mock3")
			params.SetPtype("collector")
			params.SetPversion(int64(1))

			resp, err := c.Plugins.UnloadPlugin(params, nil)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})
	})
}

func TestLoadPlugin(t *testing.T) {
	c := getSnapClient()

	Convey("Testing load a plugin", t, func() {
		Convey("Test load an existing plugin", func() {
			params := plugins.NewLoadPluginParams()
			f, err := os.Open("/tmp/snap-plugin-collector-mock1")
			if err != nil {
				t.Fatalf("\nNo plugin to load: %v", err)
			}
			defer f.Close()

			params.SetPluginData(f)
			resp, err := c.Plugins.LoadPlugin(params, openapiclient.BasicAuth("snap", "snap"))
			So(err, ShouldBeNil)
			So(resp.Payload.Name, ShouldEqual, "mock")
			So(resp.Payload.Version, ShouldEqual, 1)
		})
	})
}
