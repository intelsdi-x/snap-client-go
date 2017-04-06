package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProcessWorkflowMapNode process workflow map node
// swagger:model ProcessWorkflowMapNode
type ProcessWorkflowMapNode struct {

	// config
	Config map[string]interface{} `json:"config,omitempty"`

	// plugin name
	// Required: true
	PluginName *string `json:"plugin_name"`

	// plugin version
	PluginVersion int64 `json:"plugin_version,omitempty"`

	// process
	Process []*ProcessWorkflowMapNode `json:"process"`

	// publish
	Publish []*PublishWorkflowMapNode `json:"publish"`

	// target
	Target string `json:"target,omitempty"`
}

// Validate validates this process workflow map node
func (m *ProcessWorkflowMapNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePluginName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublish(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessWorkflowMapNode) validatePluginName(formats strfmt.Registry) error {

	if err := validate.Required("plugin_name", "body", m.PluginName); err != nil {
		return err
	}

	return nil
}

func (m *ProcessWorkflowMapNode) validateProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.Process) { // not required
		return nil
	}

	for i := 0; i < len(m.Process); i++ {

		if swag.IsZero(m.Process[i]) { // not required
			continue
		}

		if m.Process[i] != nil {

			if err := m.Process[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("process" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProcessWorkflowMapNode) validatePublish(formats strfmt.Registry) error {

	if swag.IsZero(m.Publish) { // not required
		return nil
	}

	for i := 0; i < len(m.Publish); i++ {

		if swag.IsZero(m.Publish[i]) { // not required
			continue
		}

		if m.Publish[i] != nil {

			if err := m.Publish[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("publish" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
