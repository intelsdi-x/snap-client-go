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

	"github.com/intelsdi-x/snap-client-go/client/operations"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPlugins(t *testing.T) {
	op := getOperationClient(getHost(), snapBasePath, snapScheme)

	Convey("Testing GetPlugins", t, func() {
		Convey("Test get a list of plugins", func() {
			params := operations.NewGetPluginsParams()

			resp, err := op.GetPlugins(params)
			So(err, ShouldBeNil)
			So(resp.Payload.Plugins, ShouldNotBeNil)
		})

		Convey("Test get a giving plugin", func() {
			params := operations.NewGetPluginParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(2))

			resp, err := op.GetPlugin(params)
			So(err, ShouldBeNil)
			So(resp.Payload.Name, ShouldNotBeNil)
		})

		Convey("Test get a non-existing plugin", func() {
			params := operations.NewGetPluginParams()
			params.SetPname("mock1")
			params.SetPtype("collector")
			params.SetPversion(int64(1))

			resp, err := op.GetPlugin(params)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})
	})
}

func TestUnloadPlugin(t *testing.T) {
	op := getOperationClient(getHost(), snapBasePath, snapScheme)

	Convey("Testing Unload a plugin", t, func() {
		Convey("Test unload an existing plugin", func() {
			params := operations.NewUnloadPluginParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(1))

			resp, ct, err := op.UnloadPlugin(params)
			So(err, ShouldBeNil)
			So(resp, ShouldBeNil)
			So(ct, ShouldNotBeNil)
		})

		Convey("Test unload a non-existing plugin", func() {
			params := operations.NewUnloadPluginParams()
			params.SetPname("mock3")
			params.SetPtype("collector")
			params.SetPversion(int64(1))

			ok, resp, err := op.UnloadPlugin(params)
			So(err, ShouldNotBeNil)
			So(ok, ShouldBeNil)
			So(resp, ShouldBeNil)
		})
	})
}

func TestLoadPlugin(t *testing.T) {
	op := getOperationClient(getHost(), snapBasePath, snapScheme)

	Convey("Testing load a plugin", t, func() {
		Convey("Test load an existing plugin", func() {
			params := operations.NewLoadPluginParams()
			f, err := os.Open("/tmp/snap-plugin-collector-mock1")
			if err != nil {
				t.Fatalf("\nNo plugin to load: %v", err)
			}
			defer f.Close()

			params.SetPluginData(f)
			ok, resp, err := op.LoadPlugin(params)
			So(err, ShouldBeNil)
			So(ok, ShouldBeNil)
			So(resp.Payload.Name, ShouldEqual, "mock")
			So(resp.Payload.Version, ShouldEqual, 1)
		})
	})
}
