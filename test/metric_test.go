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

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/intelsdi-x/snap-client-go/client"
	"github.com/intelsdi-x/snap-client-go/client/operations"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	snapScheme   = "http"
	snapBasePath = "v2"
)

func TestGetMetrics(t *testing.T) {
	op := getOperationClient(getHost(), snapBasePath, snapScheme)

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

func getOperationClient(host, basePath, scheme string) *operations.Client {
	transport := httptransport.New(host, basePath, []string{scheme})
	tc := client.New(transport, strfmt.Default)

	return tc.Operations
}

func getHost() string {
	host := os.Getenv("SNAP_CLIENT_GO_HOST") + ":8181"
	if host == ":8181" {
		host = "127.0.0.1:8181"
	}
	return host
}
