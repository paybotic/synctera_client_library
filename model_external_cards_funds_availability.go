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

// ExternalCardsFundsAvailability Estimated timeframe of funds availability  Value | Description --- | --- NOW | Within 30 minutes NEXT | Within the next business day FEW | Within a few business days
type ExternalCardsFundsAvailability string

// List of external_cards_funds_availability
const (
	EXTERNALCARDSFUNDSAVAILABILITY_FEW  ExternalCardsFundsAvailability = "FEW"
	EXTERNALCARDSFUNDSAVAILABILITY_NEXT ExternalCardsFundsAvailability = "NEXT"
	EXTERNALCARDSFUNDSAVAILABILITY_NOW  ExternalCardsFundsAvailability = "NOW"
)

// All allowed values of ExternalCardsFundsAvailability enum
var AllowedExternalCardsFundsAvailabilityEnumValues = []ExternalCardsFundsAvailability{
	"FEW",
	"NEXT",
	"NOW",
}

func (v *ExternalCardsFundsAvailability) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExternalCardsFundsAvailability(value)
	for _, existing := range AllowedExternalCardsFundsAvailabilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExternalCardsFundsAvailability", value)
}

// NewExternalCardsFundsAvailabilityFromValue returns a pointer to a valid ExternalCardsFundsAvailability
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExternalCardsFundsAvailabilityFromValue(v string) (*ExternalCardsFundsAvailability, error) {
	ev := ExternalCardsFundsAvailability(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExternalCardsFundsAvailability: valid values are %v", v, AllowedExternalCardsFundsAvailabilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExternalCardsFundsAvailability) IsValid() bool {
	for _, existing := range AllowedExternalCardsFundsAvailabilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to external_cards_funds_availability value
func (v ExternalCardsFundsAvailability) Ptr() *ExternalCardsFundsAvailability {
	return &v
}

type NullableExternalCardsFundsAvailability struct {
	value *ExternalCardsFundsAvailability
	isSet bool
}

func (v NullableExternalCardsFundsAvailability) Get() *ExternalCardsFundsAvailability {
	return v.value
}

func (v *NullableExternalCardsFundsAvailability) Set(val *ExternalCardsFundsAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalCardsFundsAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalCardsFundsAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalCardsFundsAvailability(val *ExternalCardsFundsAvailability) *NullableExternalCardsFundsAvailability {
	return &NullableExternalCardsFundsAvailability{value: val, isSet: true}
}

func (v NullableExternalCardsFundsAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalCardsFundsAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
