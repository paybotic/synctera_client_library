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

// BypassableErrors1 the model 'BypassableErrors1'
type BypassableErrors1 string

// List of bypassable_errors1
const (
	BYPASSABLEERRORS1_ACCOUNT_CLOSED                         BypassableErrors1 = "ACCOUNT_CLOSED"
	BYPASSABLEERRORS1_ACCOUNT_OWNER_WATCHLIST                BypassableErrors1 = "ACCOUNT_OWNER_WATCHLIST"
	BYPASSABLEERRORS1_BALANCE_VIOLATION                      BypassableErrors1 = "BALANCE_VIOLATION"
	BYPASSABLEERRORS1_IMPROPER_ACCOUNT_CUSTOMER_RELATIONSHIP BypassableErrors1 = "IMPROPER_ACCOUNT_CUSTOMER_RELATIONSHIP"
	BYPASSABLEERRORS1_IMPROPER_ACCOUNT_STATUS                BypassableErrors1 = "IMPROPER_ACCOUNT_STATUS"
	BYPASSABLEERRORS1_IMPROPER_CUSTOMER_STATUS               BypassableErrors1 = "IMPROPER_CUSTOMER_STATUS"
	BYPASSABLEERRORS1_IMPROPER_CUSTOMER_VERIFICATION_STATUS  BypassableErrors1 = "IMPROPER_CUSTOMER_VERIFICATION_STATUS"
	BYPASSABLEERRORS1_IMPROPER_PARTNER_LIFECYCLE_STATUS      BypassableErrors1 = "IMPROPER_PARTNER_LIFECYCLE_STATUS"
	BYPASSABLEERRORS1_IMPROPER_PARTNER_VERIFICATION_STATUS   BypassableErrors1 = "IMPROPER_PARTNER_VERIFICATION_STATUS"
	BYPASSABLEERRORS1_INSUFFICIENT_FUNDS                     BypassableErrors1 = "INSUFFICIENT_FUNDS"
	BYPASSABLEERRORS1_SPEND_CONTROL_VIOLATION                BypassableErrors1 = "SPEND_CONTROL_VIOLATION"
	BYPASSABLEERRORS1_SUSPECTED_FRAUD                        BypassableErrors1 = "SUSPECTED_FRAUD"
)

// All allowed values of BypassableErrors1 enum
var AllowedBypassableErrors1EnumValues = []BypassableErrors1{
	"ACCOUNT_CLOSED",
	"ACCOUNT_OWNER_WATCHLIST",
	"BALANCE_VIOLATION",
	"IMPROPER_ACCOUNT_CUSTOMER_RELATIONSHIP",
	"IMPROPER_ACCOUNT_STATUS",
	"IMPROPER_CUSTOMER_STATUS",
	"IMPROPER_CUSTOMER_VERIFICATION_STATUS",
	"IMPROPER_PARTNER_LIFECYCLE_STATUS",
	"IMPROPER_PARTNER_VERIFICATION_STATUS",
	"INSUFFICIENT_FUNDS",
	"SPEND_CONTROL_VIOLATION",
	"SUSPECTED_FRAUD",
}

func (v *BypassableErrors1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BypassableErrors1(value)
	for _, existing := range AllowedBypassableErrors1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BypassableErrors1", value)
}

// NewBypassableErrors1FromValue returns a pointer to a valid BypassableErrors1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBypassableErrors1FromValue(v string) (*BypassableErrors1, error) {
	ev := BypassableErrors1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BypassableErrors1: valid values are %v", v, AllowedBypassableErrors1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BypassableErrors1) IsValid() bool {
	for _, existing := range AllowedBypassableErrors1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bypassable_errors1 value
func (v BypassableErrors1) Ptr() *BypassableErrors1 {
	return &v
}

type NullableBypassableErrors1 struct {
	value *BypassableErrors1
	isSet bool
}

func (v NullableBypassableErrors1) Get() *BypassableErrors1 {
	return v.value
}

func (v *NullableBypassableErrors1) Set(val *BypassableErrors1) {
	v.value = val
	v.isSet = true
}

func (v NullableBypassableErrors1) IsSet() bool {
	return v.isSet
}

func (v *NullableBypassableErrors1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBypassableErrors1(val *BypassableErrors1) *NullableBypassableErrors1 {
	return &NullableBypassableErrors1{value: val, isSet: true}
}

func (v NullableBypassableErrors1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBypassableErrors1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
