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
	"fmt"
	"net/http"
	"testing"

	"github.com/intelsdi-x/snap-client-go/client/tasks"
	"github.com/intelsdi-x/snap-client-go/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTask(t *testing.T) {
	c := getSnapClient()

	Convey("Testing AddTask", t, func() {
		Convey("Test add a task", func() {
			params := tasks.NewAddTaskParams()
			tsk := `{"version":1,"schedule":{"type":"simple","interval":"15s"},"workflow":{"collect":{"metrics":{"/intel/mock/foo":{},"/intel/mock/bar":{},"/intel/mock/*/baz":{}},"config":{"/intel/mock":{"user":"root","password":"secret"}},"process":null,"publish":[{"plugin_name":"file","config":{"file":"/tmp/collected_swagger"}}]}}}`
			var task models.Task
			err := json.Unmarshal([]byte(tsk), &task)
			So(err, ShouldBeNil)
			params.SetTask(&task)

			resp, err := c.Tasks.AddTask(params)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldNotBeNil)
		})
	})
}

func TestTask(t *testing.T) {
	c := getSnapClient()

	var id string
	Convey("Testing Tasks", t, func() {
		Convey("Test get a list of tasks", func() {
			params := tasks.NewGetTasksParams()

			resp, err := c.Tasks.GetTasks(params)
			So(err, ShouldBeNil)
			So(resp.Payload, ShouldNotBeNil)
			So(len(resp.Payload.Tasks), ShouldBeGreaterThan, 0)
			id = resp.Payload.Tasks[0].ID

			Convey("Test get a task", func() {
				params := tasks.NewGetTaskParams()
				params.SetID(id)

				resp, err := c.Tasks.GetTask(params)
				So(err, ShouldBeNil)
				So(resp.Payload, ShouldNotBeNil)
				So(resp.Payload.ID, ShouldEqual, id)
			})

			Convey("Test start a task", func() {
				params := tasks.NewUpdateTaskStateParams()
				params.SetID(id)
				params.SetAction("start")

				resp, err := c.Tasks.UpdateTaskState(params)
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})

			Convey("Test watch a task", func() {
				resp, err := http.Get(fmt.Sprintf("http://%s/%s/tasks/:%s/watch", getHost(), snapBasePath, id))
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})

			Convey("Test stop a task", func() {
				params := tasks.NewUpdateTaskStateParams()
				params.SetID(id)
				params.SetAction("stop")

				resp, err := c.Tasks.UpdateTaskState(params)
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})

			Convey("Test delete a task", func() {
				params := tasks.NewRemoveTaskParams()
				params.SetID(id)

				_, err := c.Tasks.RemoveTask(params)
				So(err, ShouldBeNil)
			})
		})
	})
}
