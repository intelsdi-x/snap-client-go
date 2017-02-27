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
	"testing"

	"github.com/intelsdi-x/snap-client-go/client/operations"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPluginConfigItems(t *testing.T) {
	op := getOperationClient(getHost(), snapBasePath, snapScheme)

	Convey("Testing GetPluginConfigItems", t, func() {
		Convey("Test get the config without parameters", func() {
			params := operations.NewGetPluginConfigItemParams()

			resp, err := op.GetPluginConfigItem(params)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Test get the config with parameters", func() {
			params := operations.NewGetPluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(2))

			resp, err := op.GetPluginConfigItem(params)
			So(resp, ShouldResemble, &operations.GetPluginConfigItemOK{})
			So(err, ShouldBeNil)
		})
	})
}

func TestUpdatePluginConfigItems(t *testing.T) {
	op := getOperationClient(getHost(), snapBasePath, snapScheme)

	Convey("Testing UpdatePluginConfigItems", t, func() {
		Convey("Test set the plugin config without parameters", func() {
			params := operations.NewSetPluginConfigItemParams()

			resp, err := op.SetPluginConfigItem(params)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Test set the plugin config without a config", func() {
			params := operations.NewSetPluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(2))

			resp, err := op.SetPluginConfigItem(params)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Test set the plugin config with a config", func() {
			params := operations.NewSetPluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(1))
			cfg := `{"user":"jean","someint":1234567,"somefloat":3.1418,"somebool":false}`
			params.SetConfig(&cfg)

			resp, err := op.SetPluginConfigItem(params)
			So(err, ShouldBeNil)
			So(resp, ShouldResemble, &operations.SetPluginConfigItemOK{})
		})
	})
}

func TestDeletePluginConfigItems(t *testing.T) {
	op := getOperationClient(getHost(), snapBasePath, snapScheme)

	Convey("Testing DeletePluginConfigItems", t, func() {
		Convey("Test delete the plugin config without parameters", func() {
			params := operations.NewDeletePluginConfigItemParams()

			rep, err := op.DeletePluginConfigItem(params)
			So(err, ShouldNotBeNil)
			So(rep, ShouldBeNil)
		})

		Convey("Test delete the plugin config with parameters", func() {
			params := operations.NewDeletePluginConfigItemParams()
			params.SetPname("mock")
			params.SetPtype("collector")
			params.SetPversion(int64(1))
			params.SetConfig([]string{"somefloat"})

			resp, err := op.DeletePluginConfigItem(params)
			So(err, ShouldBeNil)
			So(resp, ShouldResemble, &operations.DeletePluginConfigItemOK{})
		})
	})
}
