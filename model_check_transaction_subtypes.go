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

// CheckTransactionSubtypes The set of valid CHECK transaction subtypes
type CheckTransactionSubtypes string

// List of check_transaction_subtypes
const (
	CHECKTRANSACTIONSUBTYPES_DEPOSIT CheckTransactionSubtypes = "MOBILE_DEPOSIT"
	CHECKTRANSACTIONSUBTYPES_DEPOSIT_REVERSAL CheckTransactionSubtypes = "MOBILE_DEPOSIT_REVERSAL"
	CHECKTRANSACTIONSUBTYPES_DEPOSIT_RETURN CheckTransactionSubtypes = "MOBILE_DEPOSIT_RETURN"
	CHECKTRANSACTIONSUBTYPES_DEPOSIT_RETURN_REVERSAL CheckTransactionSubtypes = "MOBILE_DEPOSIT_RETURN_REVERSAL"
)

// All allowed values of CheckTransactionSubtypes enum
var AllowedCheckTransactionSubtypesEnumValues = []CheckTransactionSubtypes{
	"MOBILE_DEPOSIT",
	"MOBILE_DEPOSIT_REVERSAL",
	"MOBILE_DEPOSIT_RETURN",
	"MOBILE_DEPOSIT_RETURN_REVERSAL",
}

func (v *CheckTransactionSubtypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CheckTransactionSubtypes(value)
	for _, existing := range AllowedCheckTransactionSubtypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CheckTransactionSubtypes", value)
}

// NewCheckTransactionSubtypesFromValue returns a pointer to a valid CheckTransactionSubtypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCheckTransactionSubtypesFromValue(v string) (*CheckTransactionSubtypes, error) {
	ev := CheckTransactionSubtypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CheckTransactionSubtypes: valid values are %v", v, AllowedCheckTransactionSubtypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CheckTransactionSubtypes) IsValid() bool {
	for _, existing := range AllowedCheckTransactionSubtypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to check_transaction_subtypes value
func (v CheckTransactionSubtypes) Ptr() *CheckTransactionSubtypes {
	return &v
}

type NullableCheckTransactionSubtypes struct {
	value *CheckTransactionSubtypes
	isSet bool
}

func (v NullableCheckTransactionSubtypes) Get() *CheckTransactionSubtypes {
	return v.value
}

func (v *NullableCheckTransactionSubtypes) Set(val *CheckTransactionSubtypes) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckTransactionSubtypes) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckTransactionSubtypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckTransactionSubtypes(val *CheckTransactionSubtypes) *NullableCheckTransactionSubtypes {
	return &NullableCheckTransactionSubtypes{value: val, isSet: true}
}

func (v NullableCheckTransactionSubtypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckTransactionSubtypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
