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

// GetTaskReader is a Reader for the GetTask structure.
type GetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTaskOK creates a GetTaskOK with default headers values
func NewGetTaskOK() *GetTaskOK {
	return &GetTaskOK{}
}

/*GetTaskOK handles this case with default header values.

TaskResp returns the giving task.
*/
type GetTaskOK struct {
	Payload *models.Task
}

func (o *GetTaskOK) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
