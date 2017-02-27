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
	"fmt"
	"net/http"
	"testing"

	"github.com/intelsdi-x/snap-client-go/client/operations"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTask(t *testing.T) {
	op := getOperationClient(getHost(), snapBasePath, snapScheme)

	Convey("Testing AddTask", t, func() {
		Convey("Test add a task", func() {
			params := operations.NewAddTaskParams()
			params.SetTask(`{"version":1,"schedule":{"type":"simple","interval":"15s"},"workflow":{"collect":{"metrics":{"/intel/mock/foo":{},"/intel/mock/bar":{},"/intel/mock/*/baz":{}},"config":{"/intel/mock":{"user":"root","password":"secret"}},"process":null,"publish":[{"plugin_name":"file","config":{"file":"/tmp/collected_swagger"}}]}}}`)

			resp, err := op.AddTask(params)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldNotBeNil)
		})
	})
}

func TestTask(t *testing.T) {
	op := getOperationClient(getHost(), snapBasePath, snapScheme)

	var id string
	Convey("Testing Tasks", t, func() {
		Convey("Test get a list of tasks", func() {
			params := operations.NewGetTasksParams()

			resp, err := op.GetTasks(params)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldNotBeNil)
			So(len(resp.Payload.Tasks), ShouldBeGreaterThan, 0)
			id = resp.Payload.Tasks[0].ID
		})

		Convey("Test get a task", func() {
			params := operations.NewGetTaskParams()
			params.SetID(id)

			resp, err := op.GetTask(params)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldNotBeNil)
			So(resp.Payload.ID, ShouldEqual, id)
		})

		Convey("Test start a task", func() {
			params := operations.NewUpdateTaskStateParams()
			params.SetID(id)
			params.SetAction("start")

			resp, err := op.UpdateTaskState(params)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Test watch a task", func() {
			resp, err := http.Get(fmt.Sprintf("http://%s/%s/tasks/:%s/watch", getHost(), snapBasePath, id))
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Test stop a task", func() {
			params := operations.NewUpdateTaskStateParams()
			params.SetID(id)
			params.SetAction("stop")

			resp, err := op.UpdateTaskState(params)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Test delete a task", func() {
			params := operations.NewRemoveTaskParams()
			params.SetID(id)

			_, err := op.RemoveTask(params)
			So(err, ShouldBeNil)
		})
	})
}
