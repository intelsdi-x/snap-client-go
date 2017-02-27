package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// MetricInfo MetricInfo defines metric version.
// swagger:model MetricInfo
type MetricInfo struct {

	// version
	Version int64 `json:"Version_,omitempty"`
}

// Validate validates this metric info
func (m *MetricInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
