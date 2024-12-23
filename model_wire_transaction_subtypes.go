/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// WireTransactionSubtypes The set of valid WIRE transaction subtypes
type WireTransactionSubtypes string

// List of wire_transaction_subtypes
const (
	WIRETRANSACTIONSUBTYPES_ORIGINATED WireTransactionSubtypes = "ORIGINATED"
	WIRETRANSACTIONSUBTYPES_RECEIVED   WireTransactionSubtypes = "RECEIVED"
)

// All allowed values of WireTransactionSubtypes enum
var AllowedWireTransactionSubtypesEnumValues = []WireTransactionSubtypes{
	"ORIGINATED",
	"RECEIVED",
}

func (v *WireTransactionSubtypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WireTransactionSubtypes(value)
	for _, existing := range AllowedWireTransactionSubtypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WireTransactionSubtypes", value)
}

// NewWireTransactionSubtypesFromValue returns a pointer to a valid WireTransactionSubtypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWireTransactionSubtypesFromValue(v string) (*WireTransactionSubtypes, error) {
	ev := WireTransactionSubtypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WireTransactionSubtypes: valid values are %v", v, AllowedWireTransactionSubtypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WireTransactionSubtypes) IsValid() bool {
	for _, existing := range AllowedWireTransactionSubtypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wire_transaction_subtypes value
func (v WireTransactionSubtypes) Ptr() *WireTransactionSubtypes {
	return &v
}

type NullableWireTransactionSubtypes struct {
	value *WireTransactionSubtypes
	isSet bool
}

func (v NullableWireTransactionSubtypes) Get() *WireTransactionSubtypes {
	return v.value
}

func (v *NullableWireTransactionSubtypes) Set(val *WireTransactionSubtypes) {
	v.value = val
	v.isSet = true
}

func (v NullableWireTransactionSubtypes) IsSet() bool {
	return v.isSet
}

func (v *NullableWireTransactionSubtypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireTransactionSubtypes(val *WireTransactionSubtypes) *NullableWireTransactionSubtypes {
	return &NullableWireTransactionSubtypes{value: val, isSet: true}
}

func (v NullableWireTransactionSubtypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireTransactionSubtypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
