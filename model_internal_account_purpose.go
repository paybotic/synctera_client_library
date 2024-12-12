/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// InternalAccountPurpose The purpose of the internal account. On creation, the default is PROFIT_AND_LOSS.
type InternalAccountPurpose string

// List of internal_account_purpose
const (
	INTERNALACCOUNTPURPOSE_CORE            InternalAccountPurpose = "CORE"
	INTERNALACCOUNTPURPOSE_PROFIT_AND_LOSS InternalAccountPurpose = "PROFIT_AND_LOSS"
	INTERNALACCOUNTPURPOSE_RESERVE         InternalAccountPurpose = "RESERVE"
	INTERNALACCOUNTPURPOSE_SETTLEMENT      InternalAccountPurpose = "SETTLEMENT"
	INTERNALACCOUNTPURPOSE_SUSPENSE        InternalAccountPurpose = "SUSPENSE"
	INTERNALACCOUNTPURPOSE_TREASURY        InternalAccountPurpose = "TREASURY"
)

// All allowed values of InternalAccountPurpose enum
var AllowedInternalAccountPurposeEnumValues = []InternalAccountPurpose{
	"CORE",
	"PROFIT_AND_LOSS",
	"RESERVE",
	"SETTLEMENT",
	"SUSPENSE",
	"TREASURY",
}

func (v *InternalAccountPurpose) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InternalAccountPurpose(value)
	for _, existing := range AllowedInternalAccountPurposeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InternalAccountPurpose", value)
}

// NewInternalAccountPurposeFromValue returns a pointer to a valid InternalAccountPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInternalAccountPurposeFromValue(v string) (*InternalAccountPurpose, error) {
	ev := InternalAccountPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InternalAccountPurpose: valid values are %v", v, AllowedInternalAccountPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InternalAccountPurpose) IsValid() bool {
	for _, existing := range AllowedInternalAccountPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to internal_account_purpose value
func (v InternalAccountPurpose) Ptr() *InternalAccountPurpose {
	return &v
}

type NullableInternalAccountPurpose struct {
	value *InternalAccountPurpose
	isSet bool
}

func (v NullableInternalAccountPurpose) Get() *InternalAccountPurpose {
	return v.value
}

func (v *NullableInternalAccountPurpose) Set(val *InternalAccountPurpose) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalAccountPurpose) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalAccountPurpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalAccountPurpose(val *InternalAccountPurpose) *NullableInternalAccountPurpose {
	return &NullableInternalAccountPurpose{value: val, isSet: true}
}

func (v NullableInternalAccountPurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalAccountPurpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
