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

// GetPluginReader is a Reader for the GetPlugin structure.
type GetPluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPluginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPluginOK creates a GetPluginOK with default headers values
func NewGetPluginOK() *GetPluginOK {
	return &GetPluginOK{}
}

/*GetPluginOK handles this case with default header values.

PluginResponse represents the response from plugin operations.
*/
type GetPluginOK struct {
	Payload *models.Plugin
}

func (o *GetPluginOK) Error() string {
	return fmt.Sprintf("[GET /plugins/{ptype}/{pname}/{pversion}][%d] getPluginOK  %+v", 200, o.Payload)
}

func (o *GetPluginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Plugin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
