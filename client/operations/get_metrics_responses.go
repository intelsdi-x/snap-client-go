package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intelsdi-x/snap-client-go/models"
)

// GetMetricsReader is a Reader for the GetMetrics structure.
type GetMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMetricsOK creates a GetMetricsOK with default headers values
func NewGetMetricsOK() *GetMetricsOK {
	return &GetMetricsOK{}
}

/*GetMetricsOK handles this case with default header values.

MetricsResp is the representation of metric operation response.
*/
type GetMetricsOK struct {
	Payload GetMetricsOKBody
}

func (o *GetMetricsOK) Error() string {
	return fmt.Sprintf("[GET /metrics][%d] getMetricsOK  %+v", 200, o.Payload)
}

func (o *GetMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsNotFound creates a GetMetricsNotFound with default headers values
func NewGetMetricsNotFound() *GetMetricsNotFound {
	return &GetMetricsNotFound{}
}

/*GetMetricsNotFound handles this case with default header values.

Error unsuccessful generic response to a failed API call
*/
type GetMetricsNotFound struct {
	Payload GetMetricsNotFoundBody
}

func (o *GetMetricsNotFound) Error() string {
	return fmt.Sprintf("[GET /metrics][%d] getMetricsNotFound  %+v", 404, o.Payload)
}

func (o *GetMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsInternalServerError creates a GetMetricsInternalServerError with default headers values
func NewGetMetricsInternalServerError() *GetMetricsInternalServerError {
	return &GetMetricsInternalServerError{}
}

/*GetMetricsInternalServerError handles this case with default header values.

Error unsuccessful generic response to a failed API call
*/
type GetMetricsInternalServerError struct {
	Payload GetMetricsInternalServerErrorBody
}

func (o *GetMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /metrics][%d] getMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMetricsInternalServerErrorBody get metrics internal server error body
swagger:model GetMetricsInternalServerErrorBody
*/
type GetMetricsInternalServerErrorBody map[string]string

// Validate validates this get metrics internal server error body
func (o GetMetricsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if swag.IsZero(o) { // not required
		return nil
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

/*GetMetricsNotFoundBody get metrics not found body
swagger:model GetMetricsNotFoundBody
*/
type GetMetricsNotFoundBody map[string]string

// Validate validates this get metrics not found body
func (o GetMetricsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if swag.IsZero(o) { // not required
		return nil
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

/*GetMetricsOKBody get metrics o k body
swagger:model GetMetricsOKBody
*/
type GetMetricsOKBody struct {

	// metrics
	// Required: true
	Metrics []*models.Metric `json:"Metrics"`
}

// Validate validates this get metrics o k body
func (o *GetMetricsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetrics(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMetricsOKBody) validateMetrics(formats strfmt.Registry) error {

	if err := validate.Required("getMetricsOK"+"."+"metrics", "body", o.Metrics); err != nil {
		return err
	}

	for i := 0; i < len(o.Metrics); i++ {

		if swag.IsZero(o.Metrics[i]) { // not required
			continue
		}

		if o.Metrics[i] != nil {

			if err := o.Metrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getMetricsOK" + "." + "metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
