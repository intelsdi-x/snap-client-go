package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLoadPluginParams creates a new LoadPluginParams object
// with the default values initialized.
func NewLoadPluginParams() *LoadPluginParams {
	var ()
	return &LoadPluginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadPluginParamsWithTimeout creates a new LoadPluginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadPluginParamsWithTimeout(timeout time.Duration) *LoadPluginParams {
	var ()
	return &LoadPluginParams{

		timeout: timeout,
	}
}

// NewLoadPluginParamsWithContext creates a new LoadPluginParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadPluginParamsWithContext(ctx context.Context) *LoadPluginParams {
	var ()
	return &LoadPluginParams{

		Context: ctx,
	}
}

// NewLoadPluginParamsWithHTTPClient creates a new LoadPluginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadPluginParamsWithHTTPClient(client *http.Client) *LoadPluginParams {
	var ()
	return &LoadPluginParams{
		HTTPClient: client,
	}
}

/*LoadPluginParams contains all the parameters to send to the API endpoint
for the load plugin operation typically these are written to a http.Request
*/
type LoadPluginParams struct {

	/*PluginData
	  loads a plugin.

	*/
	PluginData *os.File

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load plugin params
func (o *LoadPluginParams) WithTimeout(timeout time.Duration) *LoadPluginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load plugin params
func (o *LoadPluginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load plugin params
func (o *LoadPluginParams) WithContext(ctx context.Context) *LoadPluginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load plugin params
func (o *LoadPluginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load plugin params
func (o *LoadPluginParams) WithHTTPClient(client *http.Client) *LoadPluginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load plugin params
func (o *LoadPluginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPluginData adds the pluginData to the load plugin params
func (o *LoadPluginParams) WithPluginData(pluginData *os.File) *LoadPluginParams {
	o.SetPluginData(pluginData)
	return o
}

// SetPluginData adds the pluginData to the load plugin params
func (o *LoadPluginParams) SetPluginData(pluginData *os.File) {
	o.PluginData = pluginData
}

// WriteToRequest writes these params to a swagger request
func (o *LoadPluginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.PluginData != nil {

		if o.PluginData != nil {

			// form file param pluginData
			if err := r.SetFileParam("pluginData", o.PluginData); err != nil {
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
