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

// ExternalCardStatus Status of an External Card
type ExternalCardStatus string

// List of external_card_status
const (
	EXTERNALCARDSTATUS_ACTIVE    ExternalCardStatus = "ACTIVE"
	EXTERNALCARDSTATUS_SUSPENDED ExternalCardStatus = "SUSPENDED"
)

// All allowed values of ExternalCardStatus enum
var AllowedExternalCardStatusEnumValues = []ExternalCardStatus{
	"ACTIVE",
	"SUSPENDED",
}

func (v *ExternalCardStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExternalCardStatus(value)
	for _, existing := range AllowedExternalCardStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExternalCardStatus", value)
}

// NewExternalCardStatusFromValue returns a pointer to a valid ExternalCardStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExternalCardStatusFromValue(v string) (*ExternalCardStatus, error) {
	ev := ExternalCardStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExternalCardStatus: valid values are %v", v, AllowedExternalCardStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExternalCardStatus) IsValid() bool {
	for _, existing := range AllowedExternalCardStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to external_card_status value
func (v ExternalCardStatus) Ptr() *ExternalCardStatus {
	return &v
}

type NullableExternalCardStatus struct {
	value *ExternalCardStatus
	isSet bool
}

func (v NullableExternalCardStatus) Get() *ExternalCardStatus {
	return v.value
}

func (v *NullableExternalCardStatus) Set(val *ExternalCardStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalCardStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalCardStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalCardStatus(val *ExternalCardStatus) *NullableExternalCardStatus {
	return &NullableExternalCardStatus{value: val, isSet: true}
}

func (v NullableExternalCardStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalCardStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}