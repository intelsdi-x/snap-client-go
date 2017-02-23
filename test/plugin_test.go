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
	"os"
	"testing"

	"github.com/intelsdi-x/snap-client-go/client/snap"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPlugins(t *testing.T) {
	c := getSnapClient()

	Convey("Testing GetPlugins", t, func() {
		Convey("Test get a list of plugins", func() {
			params := snap.NewGetPluginsParams()

			resp, err := c.Snap.GetPlugins(params)
			So(err, ShouldBeNil)
			So(resp.Payload.Plugins, ShouldNotBeNil)
		})

		Convey("Test get a giving plugin", func() {
			params := snap.NewGetPluginParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(2))

			resp, err := c.Snap.GetPlugin(params)
			So(err, ShouldBeNil)
			So(resp.Payload.Name, ShouldEqual, "mock")
		})

		Convey("Test get a non-existing plugin", func() {
			params := snap.NewGetPluginParams()
			params.SetPname("mock1")
			params.SetPtype("collector")
			params.SetPversion(int64(1))

			resp, err := c.Snap.GetPlugin(params)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})
	})
}

func TestUnloadPlugin(t *testing.T) {
	c := getSnapClient()

	Convey("Testing Unload a plugin", t, func() {
		Convey("Test unload an existing plugin", func() {
			params := snap.NewUnloadPluginParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(1))

			resp, err := c.Snap.UnloadPlugin(params)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Test unload a non-existing plugin", func() {
			params := snap.NewUnloadPluginParams()
			params.SetPname("mock3")
			params.SetPtype("collector")
			params.SetPversion(int64(1))

			resp, err := c.Snap.UnloadPlugin(params)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})
	})
}

func TestLoadPlugin(t *testing.T) {
	c := getSnapClient()

	Convey("Testing load a plugin", t, func() {
		Convey("Test load an existing plugin", func() {
			params := snap.NewLoadPluginParams()
			f, err := os.Open("/tmp/snap-plugin-collector-mock1")
			if err != nil {
				t.Fatalf("\nNo plugin to load: %v", err)
			}
			defer f.Close()

			params.SetPluginData(f)
			resp, err := c.Snap.LoadPlugin(params)
			So(err, ShouldBeNil)
			So(resp.Payload.Name, ShouldEqual, "mock")
			So(resp.Payload.Version, ShouldEqual, 1)
		})
	})
}
