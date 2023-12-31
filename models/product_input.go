// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProductInput product input
//
// swagger:model ProductInput
type ProductInput struct {

	// product id
	ProductID string `json:"productId,omitempty"`

	// quantity
	Quantity float64 `json:"quantity,omitempty"`
}

// Validate validates this product input
func (m *ProductInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProductInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductInput) UnmarshalBinary(b []byte) error {
	var res ProductInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
