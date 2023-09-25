/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BypassableErrors the model 'BypassableErrors'
type BypassableErrors string

// List of bypassable_errors
const (
	BYPASSABLEERRORS_ACCOUNT_OWNER_WATCHLIST BypassableErrors = "ACCOUNT_OWNER_WATCHLIST"
	BYPASSABLEERRORS_DURATION_CONFLICT BypassableErrors = "DURATION_CONFLICT"
	BYPASSABLEERRORS_EXTERNAL_ACCOUNT_NOT_ACTIVE BypassableErrors = "EXTERNAL_ACCOUNT_NOT_ACTIVE"
	BYPASSABLEERRORS_EXTERNAL_ACCOUNT_NOT_VERIFIED BypassableErrors = "EXTERNAL_ACCOUNT_NOT_VERIFIED"
	BYPASSABLEERRORS_EXTERNAL_ACH_DISABLED_FOR_TENANT BypassableErrors = "EXTERNAL_ACH_DISABLED_FOR_TENANT"
	BYPASSABLEERRORS_FRAUD_DETECTED BypassableErrors = "FRAUD_DETECTED"
	BYPASSABLEERRORS_IMPROPER_ACCOUNT_CUSTOMER_RELATIONSHIP BypassableErrors = "IMPROPER_ACCOUNT_CUSTOMER_RELATIONSHIP"
	BYPASSABLEERRORS_IMPROPER_ACCOUNT_STATUS BypassableErrors = "IMPROPER_ACCOUNT_STATUS"
	BYPASSABLEERRORS_IMPROPER_CUSTOMER_STATUS BypassableErrors = "IMPROPER_CUSTOMER_STATUS"
	BYPASSABLEERRORS_IMPROPER_CUSTOMER_VERIFICATION_STATUS BypassableErrors = "IMPROPER_CUSTOMER_VERIFICATION_STATUS"
	BYPASSABLEERRORS_IMPROPER_PARTNER_LIFECYCLE_STATUS BypassableErrors = "IMPROPER_PARTNER_LIFECYCLE_STATUS"
	BYPASSABLEERRORS_IMPROPER_PARTNER_VERIFICATION_STATUS BypassableErrors = "IMPROPER_PARTNER_VERIFICATION_STATUS"
	BYPASSABLEERRORS_PARTIAL_HOLDS_DISABLED BypassableErrors = "PARTIAL_HOLDS_DISABLED"
)

// All allowed values of BypassableErrors enum
var AllowedBypassableErrorsEnumValues = []BypassableErrors{
	"ACCOUNT_OWNER_WATCHLIST",
	"DURATION_CONFLICT",
	"EXTERNAL_ACCOUNT_NOT_ACTIVE",
	"EXTERNAL_ACCOUNT_NOT_VERIFIED",
	"EXTERNAL_ACH_DISABLED_FOR_TENANT",
	"FRAUD_DETECTED",
	"IMPROPER_ACCOUNT_CUSTOMER_RELATIONSHIP",
	"IMPROPER_ACCOUNT_STATUS",
	"IMPROPER_CUSTOMER_STATUS",
	"IMPROPER_CUSTOMER_VERIFICATION_STATUS",
	"IMPROPER_PARTNER_LIFECYCLE_STATUS",
	"IMPROPER_PARTNER_VERIFICATION_STATUS",
	"PARTIAL_HOLDS_DISABLED",
}

func (v *BypassableErrors) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BypassableErrors(value)
	for _, existing := range AllowedBypassableErrorsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BypassableErrors", value)
}

// NewBypassableErrorsFromValue returns a pointer to a valid BypassableErrors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBypassableErrorsFromValue(v string) (*BypassableErrors, error) {
	ev := BypassableErrors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BypassableErrors: valid values are %v", v, AllowedBypassableErrorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BypassableErrors) IsValid() bool {
	for _, existing := range AllowedBypassableErrorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bypassable_errors value
func (v BypassableErrors) Ptr() *BypassableErrors {
	return &v
}

type NullableBypassableErrors struct {
	value *BypassableErrors
	isSet bool
}

func (v NullableBypassableErrors) Get() *BypassableErrors {
	return v.value
}

func (v *NullableBypassableErrors) Set(val *BypassableErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableBypassableErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableBypassableErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBypassableErrors(val *BypassableErrors) *NullableBypassableErrors {
	return &NullableBypassableErrors{value: val, isSet: true}
}

func (v NullableBypassableErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBypassableErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

