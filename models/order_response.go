// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderResponse order response
//
// swagger:model OrderResponse
type OrderResponse struct {

	// cart Id
	CartID string `json:"cartId,omitempty"`

	// totals
	Totals *OrderResponseTotals `json:"totals,omitempty"`
}

// Validate validates this order response
func (m *OrderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTotals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderResponse) validateTotals(formats strfmt.Registry) error {

	if swag.IsZero(m.Totals) { // not required
		return nil
	}

	if m.Totals != nil {
		if err := m.Totals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("totals")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderResponse) UnmarshalBinary(b []byte) error {
	var res OrderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OrderResponseTotals order response totals
//
// swagger:model OrderResponseTotals
type OrderResponseTotals struct {

	// discounts
	Discounts float64 `json:"discounts,omitempty"`

	// order
	Order float64 `json:"order,omitempty"`

	// products
	Products float64 `json:"products,omitempty"`

	// shipping
	Shipping float64 `json:"shipping,omitempty"`
}

// Validate validates this order response totals
func (m *OrderResponseTotals) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderResponseTotals) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderResponseTotals) UnmarshalBinary(b []byte) error {
	var res OrderResponseTotals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}