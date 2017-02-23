package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RuleTable RuleTable defines a named property characteristics.
// swagger:model RuleTable
type RuleTable struct {

	// default
	Default interface{} `json:"default,omitempty"`

	// maximum
	Maximum interface{} `json:"maximum,omitempty"`

	// minimum
	Minimum interface{} `json:"minimum,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// required
	// Required: true
	Required *bool `json:"required"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this rule table
func (m *RuleTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRequired(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleTable) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RuleTable) validateRequired(formats strfmt.Registry) error {

	if err := validate.Required("required", "body", m.Required); err != nil {
		return err
	}

	return nil
}

func (m *RuleTable) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
