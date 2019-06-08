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

// UserTeamParams user team params
// swagger:model user_team_params
type UserTeamParams struct {

	// perm
	// Required: true
	// Enum: [user admin owner]
	Perm *string `json:"perm"`

	// team
	// Required: true
	Team *string `json:"team"`
}

// Validate validates this user team params
func (m *UserTeamParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userTeamParamsTypePermPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","admin","owner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userTeamParamsTypePermPropEnum = append(userTeamParamsTypePermPropEnum, v)
	}
}

const (

	// UserTeamParamsPermUser captures enum value "user"
	UserTeamParamsPermUser string = "user"

	// UserTeamParamsPermAdmin captures enum value "admin"
	UserTeamParamsPermAdmin string = "admin"

	// UserTeamParamsPermOwner captures enum value "owner"
	UserTeamParamsPermOwner string = "owner"
)

// prop value enum
func (m *UserTeamParams) validatePermEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userTeamParamsTypePermPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserTeamParams) validatePerm(formats strfmt.Registry) error {

	if err := validate.Required("perm", "body", m.Perm); err != nil {
		return err
	}

	// value enum
	if err := m.validatePermEnum("perm", "body", *m.Perm); err != nil {
		return err
	}

	return nil
}

func (m *UserTeamParams) validateTeam(formats strfmt.Registry) error {

	if err := validate.Required("team", "body", m.Team); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserTeamParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserTeamParams) UnmarshalBinary(b []byte) error {
	var res UserTeamParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
