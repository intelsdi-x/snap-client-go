package snap

import "github.com/intelsdi-x/snap-client-go/client/operations"

// NewGetPluginParams defines the input for getting a plugin.
// The required parameters are plugin type, name and version.
func NewGetPluginParams() *operations.GetPluginParams {
	return operations.NewGetPluginParams()
}

// NewGetPluginsParams defines the input for getting a list of loaded plugins.
// No specific parameter is required.
func NewGetPluginsParams() *operations.GetPluginsParams {
	return operations.NewGetPluginsParams()
}

// NewLoadPluginParams defines the input for loading a plugin.
// To be loaded plugin binary is required.
func NewLoadPluginParams() *operations.LoadPluginParams {
	return operations.NewLoadPluginParams()
}

// NewUnloadPluginParams defines the input for unload a plugin.
// It requires the plugin type, name and version.
func NewUnloadPluginParams() *operations.UnloadPluginParams {
	return operations.NewUnloadPluginParams()
}

// NewGetPluginConfigItemParams defines the input for getting a plugin config.
// It requires a plugin type, name and version.
func NewGetPluginConfigItemParams() *operations.GetPluginConfigItemParams {
	return operations.NewGetPluginConfigItemParams()
}

// NewSetPluginConfigItemParams defines the input for setting a plugin config.
// It requires a plugin type, name, version and the config.
func NewSetPluginConfigItemParams() *operations.SetPluginConfigItemParams {
	return operations.NewSetPluginConfigItemParams()
}

// NewDeletePluginConfigItemParams defines the input for deleting a plugin config.
// It requires a plugin type, name, version and the config keys.
func NewDeletePluginConfigItemParams() *operations.DeletePluginConfigItemParams {
	return operations.NewDeletePluginConfigItemParams()
}

// NewGetMetricsParams defines the input for getting  a list of metrics.
// No specific paramter is required.
func NewGetMetricsParams() *operations.GetMetricsParams {
	return operations.NewGetMetricsParams()
}

// NewAddTaskParams defines the input for creating a task.
// Supplying a task manifest or workflow JSON string as the input.
func NewAddTaskParams() *operations.AddTaskParams {
	return operations.NewAddTaskParams()
}

// NewGetTasksParams defines the input for getting a list of tasks.
// No specific paramter is required.
func NewGetTasksParams() *operations.GetTasksParams {
	return operations.NewGetTasksParams()
}

// NewGetTaskParams defines the input for getting a task by its id.
func NewGetTaskParams() *operations.GetTaskParams {
	return operations.NewGetTaskParams()
}

// NewUpdateTaskStateParams defines the input for updating the state of a task.
// Two parameters are required. They're taskid and action.
func NewUpdateTaskStateParams() *operations.UpdateTaskStateParams {
	return operations.NewUpdateTaskStateParams()
}

// NewRemoveTaskParams defines the input for removign a task. Note that only a stopped
// task may be removed. The required paramter is the task id.
func NewRemoveTaskParams() *operations.RemoveTaskParams {
	return operations.NewRemoveTaskParams()
}
