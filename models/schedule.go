package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Schedule Schedule defines a scheduler.
// swagger:model Schedule
type Schedule struct {

	// count
	Count uint64 `json:"count,omitempty"`

	// interval
	// Required: true
	Interval *string `json:"interval"`

	// start timestamp
	StartTimestamp interface{} `json:"start_timestamp,omitempty"`

	// stop timestamp
	StopTimestamp interface{} `json:"stop_timestamp,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this schedule
func (m *Schedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterval(formats); err != nil {
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

func (m *Schedule) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

var scheduleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["simple"," windowed"," streaming"," cron"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleTypeTypePropEnum = append(scheduleTypeTypePropEnum, v)
	}
}

const (
	// ScheduleTypeSimple captures enum value "simple"
	ScheduleTypeSimple string = "simple"
	// ScheduleTypeNrWindowed captures enum value " windowed"
	ScheduleTypeNrWindowed string = " windowed"
	// ScheduleTypeNrStreaming captures enum value " streaming"
	ScheduleTypeNrStreaming string = " streaming"
	// ScheduleTypeNrCron captures enum value " cron"
	ScheduleTypeNrCron string = " cron"
)

// prop value enum
func (m *Schedule) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, scheduleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Schedule) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}
