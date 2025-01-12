// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NutanixCSIConfig NutanixCSIConfig contains credentials and the endpoint for the Nutanix Prism Element to which the CSI driver connects.
//
// swagger:model NutanixCSIConfig
type NutanixCSIConfig struct {

	// Prism Element Endpoint to access Nutanix Prism Element for csi driver
	Endpoint string `json:"endpoint,omitempty"`

	// Prism Element Password for csi driver
	Password string `json:"password,omitempty"`

	// Optional: Port to use when connecting to the Nutanix Prism Element endpoint (defaults to 9440)
	// +optional
	Port int32 `json:"port,omitempty"`

	// Prism Element Username for csi driver
	Username string `json:"username,omitempty"`
}

// Validate validates this nutanix c s i config
func (m *NutanixCSIConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nutanix c s i config based on context it is used
func (m *NutanixCSIConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NutanixCSIConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NutanixCSIConfig) UnmarshalBinary(b []byte) error {
	var res NutanixCSIConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
