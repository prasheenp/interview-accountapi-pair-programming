// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ValidationSource validation source
// swagger:model ValidationSource
type ValidationSource string

const (
	// ValidationSourcePaymentAPI captures enum value "payment_api"
	ValidationSourcePaymentAPI ValidationSource = "payment_api"
	// ValidationSourcePayportInterface captures enum value "payport_interface"
	ValidationSourcePayportInterface ValidationSource = "payport_interface"
	// ValidationSourceStarlingGateway captures enum value "starling_gateway"
	ValidationSourceStarlingGateway ValidationSource = "starling_gateway"
	// ValidationSourceBacsGateway captures enum value "bacs_gateway"
	ValidationSourceBacsGateway ValidationSource = "bacs_gateway"
	// ValidationSourceSepainstantGateway captures enum value "sepainstant_gateway"
	ValidationSourceSepainstantGateway ValidationSource = "sepainstant_gateway"
)

// for schema
var validationSourceEnum []interface{}

func init() {
	var res []ValidationSource
	if err := json.Unmarshal([]byte(`["payment_api","payport_interface","starling_gateway","bacs_gateway","sepainstant_gateway"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		validationSourceEnum = append(validationSourceEnum, v)
	}
}

func (m ValidationSource) validateValidationSourceEnum(path, location string, value ValidationSource) error {
	if err := validate.Enum(path, location, value, validationSourceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this validation source
func (m ValidationSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateValidationSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}