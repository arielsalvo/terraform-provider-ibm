// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceAction instanceAction
// swagger:model instance_action
type InstanceAction struct {

	// The date and time that the action was completed
	// Format: date-time
	CompletedAt strfmt.DateTime `json:"completed_at,omitempty"`

	// The date and time that the action was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The URL for this action
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The identifier for this action
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The date and time that the action was started
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// The current status of this action
	// Enum: [pending running failed completed]
	Status string `json:"status,omitempty"`

	// The type of action
	// Enum: [start stop reboot reset]
	Type string `json:"type,omitempty"`
}

// Validate validates this instance action
func (m *InstanceAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceAction) validateCompletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InstanceAction) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InstanceAction) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *InstanceAction) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InstanceAction) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var instanceActionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","running","failed","completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceActionTypeStatusPropEnum = append(instanceActionTypeStatusPropEnum, v)
	}
}

const (

	// InstanceActionStatusPending captures enum value "pending"
	InstanceActionStatusPending string = "pending"

	// InstanceActionStatusRunning captures enum value "running"
	InstanceActionStatusRunning string = "running"

	// InstanceActionStatusFailed captures enum value "failed"
	InstanceActionStatusFailed string = "failed"

	// InstanceActionStatusCompleted captures enum value "completed"
	InstanceActionStatusCompleted string = "completed"
)

// prop value enum
func (m *InstanceAction) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceActionTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceAction) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var instanceActionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["start","stop","reboot","reset"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceActionTypeTypePropEnum = append(instanceActionTypeTypePropEnum, v)
	}
}

const (

	// InstanceActionTypeStart captures enum value "start"
	InstanceActionTypeStart string = "start"

	// InstanceActionTypeStop captures enum value "stop"
	InstanceActionTypeStop string = "stop"

	// InstanceActionTypeReboot captures enum value "reboot"
	InstanceActionTypeReboot string = "reboot"

	// InstanceActionTypeReset captures enum value "reset"
	InstanceActionTypeReset string = "reset"
)

// prop value enum
func (m *InstanceAction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceActionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceAction) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceAction) UnmarshalBinary(b []byte) error {
	var res InstanceAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
