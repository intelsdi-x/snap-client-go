package operations

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

TaskResp returns the giving task.
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

// NewRemoveTaskNotFound creates a RemoveTaskNotFound with default headers values
func NewRemoveTaskNotFound() *RemoveTaskNotFound {
	return &RemoveTaskNotFound{}
}

/*RemoveTaskNotFound handles this case with default header values.

Error unsuccessful generic response to a failed API call
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

RemoveTaskError unsuccessful generic response to a failed API call
*/
type RemoveTaskInternalServerError struct {
}

func (o *RemoveTaskInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] removeTaskInternalServerError ", 500)
}

func (o *RemoveTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
