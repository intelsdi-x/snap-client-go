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

package integration

import (
	"testing"

	"github.com/intelsdi-x/snap/client"
	"github.com/intelsdi-x/snap/client/operations"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetMetrics(t *testing.T) {
	op := client.Default.Operations

	Convey("Testing GetMetrics", t, func() {
		Convey("Test get a list of metrics", func() {
			params := operations.NewGetMetricsParams()

			resp, err := op.GetMetrics(params)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldNotBeNil)
			So(len(resp.Payload.Metrics), ShouldBeGreaterThan, 0)
		})
	})
}
