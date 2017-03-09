# snap-client-go

## Overview

This repo has [Snap](https://github.com/intelsdi-x/snap) restful Go clients generated by [go-swagger](https://github.com/go-swagger/go-swagger). You can run the following command to view available Snap APIs locally.
> swagger serve swagger.json --host=127.0.0.1

You can browse and test Snap APIs interactively by leveraging tools like [swagger-ui](https://github.com/swagger-api/swagger-ui) or [apistudio.io](http://apistudio.io/).

If any change for swagger.json specification, running the command `make` to generate a new client whenever the specification changes.

## Creating client

Creating a default client by using an empty param object. If supplying the different host info, API version, or scheme, it creates a custom client.

### Default client

```sh
  c := snap.New(snap.ClientParams{})
```

### Custom client

```sh
  c := snap.New(snap.ClientParam{URL: "example.com:80", APIVersion: "v2", Scheme: "http"})
```

## Operations

For operations, simply do following with appropriate parameters required by the specification. Let's assume `c` is the client created and examples:

### Get plugins

```sh
getPlugins, err := c.GetPlugins(operations.NewGetPluginsParams())
```

### Get a plugin

```sh
params := operations.NewGetPluginParams()
params.SetPtype("collector")
params.SetPname("mock")
params.SetPversion(int64(1))

getPlugin, err := c.GetPlugin(params)
```

### Load a plugin

```sh
params := operations.NewLoadPluginParams()
params.SetPluginData(<*os.file>)

loadPlugin, err := c.LoadPlugin(params)
```

### Unload a plugin

```sh
params := operations.NewUnloadPluginParams()
params.SetPtype("collector")
params.SetPname("mock")
params.SetPversion(int64(1))

unloadPlugin, err := c.UnloadPlugin(params)
```

### Get a plugin config

```sh
params := operations.NewGetPluginConfigItemParams()
params.SetPname("mock")
params.SetPtype("collector")
params.SetPversion(int64(1))

getPluginConfig, err := c.GetPluginConfigItem(params)
```

### Update a plugin config

```sh
cfg := `{"user":"jean","someint":1234567,"somefloat":3.1418,"somebool":false}`

params := operations.NewSetPluginConfigItemParams()
params.SetPname("mock")
params.SetPtype("collector")
params.SetPversion(int64(1))
params.SetConfig(&cfg)

setPluginConfig, err := c.SetPluginConfigItem(params)
```

### Delete a plugin config

```sh
params := operations.NewDeletePluginConfigItemParams()
params.SetPname("mock")
params.SetPtype("collector")
params.SetPversion(int64(1))
params.SetConfig([]string{"somefloat", "someint"})

deletePluginConfig, err := c.DeletePluginConfigItem(params)
```

### Get metrics

```sh
params := operations.NewGetMetricsParams()

getMetrics, err := c.GetMetrics(params)
```

### Get metrics giving a namespace

```sh
ns := "/intel/mock/bar"
params := operations.NewGetMetricsParams()
params.SetNs(&ns)

getMetrics, err := c.GetMetrics(params)
```

### Get a metric

```sh
ns := "/intel/mock/bar"
ver := int64(1)
params := operations.NewGetMetricsParams()
params.SetNs(&ns)
params.SetVer(&ver)

getMetric, err := c.GetMetrics(params)
```

### Get tasks

```sh
params := operations.NewGetTasksParams()

getTasks, err := c.GetTasks(params)
```

### Get a task

```sh
params := operations.NewGetTaskParams()
params.SetID(id)

getTask, err := c.GetTask(params)
```

### Create a task

```sh
params := operations.NewAddTaskParams()
params.SetTask(<task manifest> || <workflow manifest>)

createTask, err := c.AddTask(params)
```

### Start a task

```sh
params := operations.NewUpdateTaskStateParams()
params.SetID(<task id>)
params.SetAction("start")

startTask, err := c.UpdateTaskState(params)
```

### Stop a task

```sh
params := operations.NewUpdateTaskStateParams()
params.SetID(<task id>)
params.SetAction("stop")

stopTask, err := c.UpdateTaskState(params)
```

### Enable a task

```sh
params := operations.NewUpdateTaskStateParams()
params.SetID(<task id>)
params.SetAction("enable")

enableTask, err := c.UpdateTaskState(params)
```

### Remove a task

```sh
params := operations.NewRemoveTaskParams()
params.SetID(<task id>)

removeTask, err := c.RemoveTask(params)
```

## Running tests

The package `test` contains integration tests that verify the generated client library if it is working appropriately against the actual behavior of the Snap APIs. Tests will fail if any incompatible change in the API.

For your convenience, the Docker and docker-compose based medium test is available for you to try out.  At the root package, you can run:  

```sh
make test-medium
```

## Security Disclosure

The Snap team takes security very seriously. If you have any issue regarding security, please notify us by sending an email to snap-security@intel.com
and not by creating a GitHub issue. We will follow up with you promptly with more information and a plan for remediation.

## License
Snap Client Go is Open Source software released under the [Apache 2.0 License](LICENSE).

## Thank You
And **thank you!** Your contribution, through code and participation, is incredibly important to us.