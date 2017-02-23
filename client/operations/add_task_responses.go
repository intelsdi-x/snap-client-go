package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intelsdi-x/snap-client-go/models"
)

// AddTaskReader is a Reader for the AddTask structure.
type AddTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddTaskCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewAddTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddTaskCreated creates a AddTaskCreated with default headers values
func NewAddTaskCreated() *AddTaskCreated {
	return &AddTaskCreated{}
}

/*AddTaskCreated handles this case with default header values.

TasksResp returns a list of created tasks.
*/
type AddTaskCreated struct {
	Payload AddTaskCreatedBody
}

func (o *AddTaskCreated) Error() string {
	return fmt.Sprintf("[POST /tasks][%d] addTaskCreated  %+v", 201, o.Payload)
}

func (o *AddTaskCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTaskInternalServerError creates a AddTaskInternalServerError with default headers values
func NewAddTaskInternalServerError() *AddTaskInternalServerError {
	return &AddTaskInternalServerError{}
}

/*AddTaskInternalServerError handles this case with default header values.

Error unsuccessful generic response to a failed API call
*/
type AddTaskInternalServerError struct {
	Payload AddTaskInternalServerErrorBody
}

func (o *AddTaskInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tasks][%d] addTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *AddTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddTaskCreatedBody add task created body
swagger:model AddTaskCreatedBody
*/
type AddTaskCreatedBody struct {

	// tasks
	// Required: true
	Tasks models.Tasks `json:"Tasks"`
}

// Validate validates this add task created body
func (o *AddTaskCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTasks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddTaskCreatedBody) validateTasks(formats strfmt.Registry) error {

	if err := validate.Required("addTaskCreated"+"."+"tasks", "body", o.Tasks); err != nil {
		return err
	}

	return nil
}

/*AddTaskInternalServerErrorBody add task internal server error body
swagger:model AddTaskInternalServerErrorBody
*/
type AddTaskInternalServerErrorBody map[string]string

// Validate validates this add task internal server error body
func (o AddTaskInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if swag.IsZero(o) { // not required
		return nil
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
