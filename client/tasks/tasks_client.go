package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddTask adds

A string representation of Snap task manifest is required.
*/
func (a *Client) AddTask(params *AddTaskParams) (*AddTaskCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addTask",
		Method:             "POST",
		PathPattern:        "/tasks",
		ProducesMediaTypes: []string{"application/json", "application/x-protobuf"},
		ConsumesMediaTypes: []string{"application/json", "application/x-protobuf"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddTaskCreated), nil

}

/*
GetTask gets

The task ID is required.
*/
func (a *Client) GetTask(params *GetTaskParams) (*GetTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTask",
		Method:             "GET",
		PathPattern:        "/tasks/{id}",
		ProducesMediaTypes: []string{"application/json", "application/x-protobuf"},
		ConsumesMediaTypes: []string{"application/json", "application/x-protobuf"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTaskOK), nil

}

/*
GetTasks gets all

An empty list returns if no tasks exist.
*/
func (a *Client) GetTasks(params *GetTasksParams) (*GetTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTasks",
		Method:             "GET",
		PathPattern:        "/tasks",
		ProducesMediaTypes: []string{"application/json", "application/x-protobuf"},
		ConsumesMediaTypes: []string{"application/json", "application/x-protobuf"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTasksOK), nil

}

/*
RemoveTask removes

The task ID is required.
*/
func (a *Client) RemoveTask(params *RemoveTaskParams) (*RemoveTaskNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeTask",
		Method:             "DELETE",
		PathPattern:        "/tasks/{id}",
		ProducesMediaTypes: []string{"application/json", "application/x-protobuf"},
		ConsumesMediaTypes: []string{"application/json", "application/x-protobuf"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveTaskNoContent), nil

}

/*
UpdateTaskState enables start stop

The task ID is required.
*/
func (a *Client) UpdateTaskState(params *UpdateTaskStateParams) (*UpdateTaskStateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTaskStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateTaskState",
		Method:             "PUT",
		PathPattern:        "/tasks/{id}",
		ProducesMediaTypes: []string{"application/json", "application/x-protobuf"},
		ConsumesMediaTypes: []string{"application/json", "application/x-protobuf"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateTaskStateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateTaskStateNoContent), nil

}

/*
WatchTask watches

The task ID is required.
*/
func (a *Client) WatchTask(params *WatchTaskParams) (*WatchTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchTask",
		Method:             "GET",
		PathPattern:        "/tasks/{id}/watch",
		ProducesMediaTypes: []string{"application/json", "application/x-protobuf", "text/event-stream"},
		ConsumesMediaTypes: []string{"application/json", "application/x-protobuf", "text/event-stream"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WatchTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WatchTaskOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
