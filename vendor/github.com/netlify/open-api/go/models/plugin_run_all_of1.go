// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PluginRunAllOf1 plugin run all of1
//
// swagger:model pluginRunAllOf1
type PluginRunAllOf1 struct {

	// deploy id
	DeployID string `json:"deploy_id,omitempty"`
}

// Validate validates this plugin run all of1
func (m *PluginRunAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginRunAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginRunAllOf1) UnmarshalBinary(b []byte) error {
	var res PluginRunAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
