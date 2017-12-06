// Code generated by go-swagger; DO NOT EDIT.

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

// CloudDetails cloud details
// swagger:model CloudDetails
type CloudDetails struct {

	// bmc
	Bmc *CloudDetailsBmc `json:"bmc,omitempty"`
}

// Validate validates this cloud details
func (m *CloudDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBmc(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudDetails) validateBmc(formats strfmt.Registry) error {

	if swag.IsZero(m.Bmc) { // not required
		return nil
	}

	if m.Bmc != nil {

		if err := m.Bmc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bmc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudDetails) UnmarshalBinary(b []byte) error {
	var res CloudDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudDetailsBmc cloud details bmc
// swagger:model CloudDetailsBmc
type CloudDetailsBmc struct {

	// lb shape
	LbShape string `json:"lbShape,omitempty"`

	// lb type
	LbType string `json:"lbType,omitempty"`

	// node Id
	NodeID string `json:"nodeId,omitempty"`
}

// Validate validates this cloud details bmc
func (m *CloudDetailsBmc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLbType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cloudDetailsBmcTypeLbTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["node-port","external"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudDetailsBmcTypeLbTypePropEnum = append(cloudDetailsBmcTypeLbTypePropEnum, v)
	}
}

const (
	// CloudDetailsBmcLbTypeNodePort captures enum value "node-port"
	CloudDetailsBmcLbTypeNodePort string = "node-port"
	// CloudDetailsBmcLbTypeExternal captures enum value "external"
	CloudDetailsBmcLbTypeExternal string = "external"
)

// prop value enum
func (m *CloudDetailsBmc) validateLbTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudDetailsBmcTypeLbTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudDetailsBmc) validateLbType(formats strfmt.Registry) error {

	if swag.IsZero(m.LbType) { // not required
		return nil
	}

	// value enum
	if err := m.validateLbTypeEnum("bmc"+"."+"lbType", "body", m.LbType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudDetailsBmc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudDetailsBmc) UnmarshalBinary(b []byte) error {
	var res CloudDetailsBmc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}