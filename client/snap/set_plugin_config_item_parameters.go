package snap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSetPluginConfigItemParams creates a new SetPluginConfigItemParams object
// with the default values initialized.
func NewSetPluginConfigItemParams() *SetPluginConfigItemParams {
	var ()
	return &SetPluginConfigItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetPluginConfigItemParamsWithTimeout creates a new SetPluginConfigItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetPluginConfigItemParamsWithTimeout(timeout time.Duration) *SetPluginConfigItemParams {
	var ()
	return &SetPluginConfigItemParams{

		timeout: timeout,
	}
}

// NewSetPluginConfigItemParamsWithContext creates a new SetPluginConfigItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetPluginConfigItemParamsWithContext(ctx context.Context) *SetPluginConfigItemParams {
	var ()
	return &SetPluginConfigItemParams{

		Context: ctx,
	}
}

// NewSetPluginConfigItemParamsWithHTTPClient creates a new SetPluginConfigItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetPluginConfigItemParamsWithHTTPClient(client *http.Client) *SetPluginConfigItemParams {
	var ()
	return &SetPluginConfigItemParams{
		HTTPClient: client,
	}
}

/*SetPluginConfigItemParams contains all the parameters to send to the API endpoint
for the set plugin config item operation typically these are written to a http.Request
*/
type SetPluginConfigItemParams struct {

	/*Config*/
	Config *string
	/*Pname*/
	Pname string
	/*Ptype*/
	Ptype string
	/*Pversion*/
	Pversion int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set plugin config item params
func (o *SetPluginConfigItemParams) WithTimeout(timeout time.Duration) *SetPluginConfigItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set plugin config item params
func (o *SetPluginConfigItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set plugin config item params
func (o *SetPluginConfigItemParams) WithContext(ctx context.Context) *SetPluginConfigItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set plugin config item params
func (o *SetPluginConfigItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set plugin config item params
func (o *SetPluginConfigItemParams) WithHTTPClient(client *http.Client) *SetPluginConfigItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set plugin config item params
func (o *SetPluginConfigItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the set plugin config item params
func (o *SetPluginConfigItemParams) WithConfig(config *string) *SetPluginConfigItemParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the set plugin config item params
func (o *SetPluginConfigItemParams) SetConfig(config *string) {
	o.Config = config
}

// WithPname adds the pname to the set plugin config item params
func (o *SetPluginConfigItemParams) WithPname(pname string) *SetPluginConfigItemParams {
	o.SetPname(pname)
	return o
}

// SetPname adds the pname to the set plugin config item params
func (o *SetPluginConfigItemParams) SetPname(pname string) {
	o.Pname = pname
}

// WithPtype adds the ptype to the set plugin config item params
func (o *SetPluginConfigItemParams) WithPtype(ptype string) *SetPluginConfigItemParams {
	o.SetPtype(ptype)
	return o
}

// SetPtype adds the ptype to the set plugin config item params
func (o *SetPluginConfigItemParams) SetPtype(ptype string) {
	o.Ptype = ptype
}

// WithPversion adds the pversion to the set plugin config item params
func (o *SetPluginConfigItemParams) WithPversion(pversion int64) *SetPluginConfigItemParams {
	o.SetPversion(pversion)
	return o
}

// SetPversion adds the pversion to the set plugin config item params
func (o *SetPluginConfigItemParams) SetPversion(pversion int64) {
	o.Pversion = pversion
}

// WriteToRequest writes these params to a swagger request
func (o *SetPluginConfigItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Config != nil {

		// form param config
		var frConfig string
		if o.Config != nil {
			frConfig = *o.Config
		}
		fConfig := frConfig
		if fConfig != "" {
			if err := r.SetFormParam("config", fConfig); err != nil {
				return err
			}
		}

	}

	// path param pname
	if err := r.SetPathParam("pname", o.Pname); err != nil {
		return err
	}

	// path param ptype
	if err := r.SetPathParam("ptype", o.Ptype); err != nil {
		return err
	}

	// path param pversion
	if err := r.SetPathParam("pversion", swag.FormatInt64(o.Pversion)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}