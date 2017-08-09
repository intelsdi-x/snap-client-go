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
// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intelsdi-x/snap-client-go/models"
)

// RemoveTaskReader is a Reader for the RemoveTask structure.
type RemoveTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRemoveTaskNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRemoveTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRemoveTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveTaskNoContent creates a RemoveTaskNoContent with default headers values
func NewRemoveTaskNoContent() *RemoveTaskNoContent {
	return &RemoveTaskNoContent{}
}

/*RemoveTaskNoContent handles this case with default header values.

TaskResponse returns a task.
*/
type RemoveTaskNoContent struct {
	Payload *models.Task
}

func (o *RemoveTaskNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] removeTaskNoContent  %+v", 204, o.Payload)
}

func (o *RemoveTaskNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTaskUnauthorized creates a RemoveTaskUnauthorized with default headers values
func NewRemoveTaskUnauthorized() *RemoveTaskUnauthorized {
	return &RemoveTaskUnauthorized{}
}

/*RemoveTaskUnauthorized handles this case with default header values.

UnauthResponse returns Unauthorized error struct message.
*/
type RemoveTaskUnauthorized struct {
	Payload *models.UnauthError
}

func (o *RemoveTaskUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] removeTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTaskNotFound creates a RemoveTaskNotFound with default headers values
func NewRemoveTaskNotFound() *RemoveTaskNotFound {
	return &RemoveTaskNotFound{}
}

/*RemoveTaskNotFound handles this case with default header values.

ErrorResponse represents the Snap error response type.

It includes an error message and a map of fields.
*/
type RemoveTaskNotFound struct {
	Payload *models.Error
}

func (o *RemoveTaskNotFound) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] removeTaskNotFound  %+v", 404, o.Payload)
}

func (o *RemoveTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTaskInternalServerError creates a RemoveTaskInternalServerError with default headers values
func NewRemoveTaskInternalServerError() *RemoveTaskInternalServerError {
	return &RemoveTaskInternalServerError{}
}

/*RemoveTaskInternalServerError handles this case with default header values.

TaskErrorResponse returns removing a task error.
*/
type RemoveTaskInternalServerError struct {
}

func (o *RemoveTaskInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] removeTaskInternalServerError ", 500)
}

func (o *RemoveTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
