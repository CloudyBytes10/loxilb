// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BGPNeighGetEntry b g p neigh get entry
//
// swagger:model BGPNeighGetEntry
type BGPNeighGetEntry struct {

	// BGP Neighbor IP address
	IPAddress string `json:"ipAddress,omitempty"`

	// Remote AS number
	RemoteAs int64 `json:"remoteAs,omitempty"`

	// Current state
	State string `json:"state,omitempty"`

	// Current uptime
	Updowntime string `json:"updowntime,omitempty"`
}

// Validate validates this b g p neigh get entry
func (m *BGPNeighGetEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this b g p neigh get entry based on context it is used
func (m *BGPNeighGetEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BGPNeighGetEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BGPNeighGetEntry) UnmarshalBinary(b []byte) error {
	var res BGPNeighGetEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
