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

// InternalAccountType type associated with the internal account.
type InternalAccountType string

// List of internal_account_type
const (
	INTERNALACCOUNTTYPE_ACH_SETTLEMENT InternalAccountType = "ACH_SETTLEMENT"
	INTERNALACCOUNTTYPE_ACH_SUSPENSE InternalAccountType = "ACH_SUSPENSE"
	INTERNALACCOUNTTYPE_CARD_SETTLEMENT InternalAccountType = "CARD_SETTLEMENT"
	INTERNALACCOUNTTYPE_INTEREST_PAYOUT InternalAccountType = "INTEREST_PAYOUT"
	INTERNALACCOUNTTYPE_WIRE_SETTLEMENT InternalAccountType = "WIRE_SETTLEMENT"
	INTERNALACCOUNTTYPE_WIRE_SUSPENSE InternalAccountType = "WIRE_SUSPENSE"
	INTERNALACCOUNTTYPE_CHECK_SETTLEMENT InternalAccountType = "CHECK_SETTLEMENT"
	INTERNALACCOUNTTYPE_CASH_SETTLEMENT InternalAccountType = "CASH_SETTLEMENT"
	INTERNALACCOUNTTYPE_CASH_SUSPENSE InternalAccountType = "CASH_SUSPENSE"
	INTERNALACCOUNTTYPE_NETWORK_ADJUSTMENT InternalAccountType = "NETWORK_ADJUSTMENT"
	INTERNALACCOUNTTYPE_FEES InternalAccountType = "FEES"
	INTERNALACCOUNTTYPE_REWARDS InternalAccountType = "REWARDS"
)

// All allowed values of InternalAccountType enum
var AllowedInternalAccountTypeEnumValues = []InternalAccountType{
	"ACH_SETTLEMENT",
	"ACH_SUSPENSE",
	"CARD_SETTLEMENT",
	"INTEREST_PAYOUT",
	"WIRE_SETTLEMENT",
	"WIRE_SUSPENSE",
	"CHECK_SETTLEMENT",
	"CASH_SETTLEMENT",
	"CASH_SUSPENSE",
	"NETWORK_ADJUSTMENT",
	"FEES",
	"REWARDS",
}

func (v *InternalAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InternalAccountType(value)
	for _, existing := range AllowedInternalAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InternalAccountType", value)
}

// NewInternalAccountTypeFromValue returns a pointer to a valid InternalAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInternalAccountTypeFromValue(v string) (*InternalAccountType, error) {
	ev := InternalAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InternalAccountType: valid values are %v", v, AllowedInternalAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InternalAccountType) IsValid() bool {
	for _, existing := range AllowedInternalAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to internal_account_type value
func (v InternalAccountType) Ptr() *InternalAccountType {
	return &v
}

type NullableInternalAccountType struct {
	value *InternalAccountType
	isSet bool
}

func (v NullableInternalAccountType) Get() *InternalAccountType {
	return v.value
}

func (v *NullableInternalAccountType) Set(val *InternalAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalAccountType(val *InternalAccountType) *NullableInternalAccountType {
	return &NullableInternalAccountType{value: val, isSet: true}
}

func (v NullableInternalAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

