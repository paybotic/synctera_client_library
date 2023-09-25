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

// InternationalWireStatus The current status of the transfer.
type InternationalWireStatus string

// List of international_wire_status
const (
	INTERNATIONALWIRESTATUS_CANCELED InternationalWireStatus = "CANCELED"
	INTERNATIONALWIRESTATUS_COMPLETED InternationalWireStatus = "COMPLETED"
	INTERNATIONALWIRESTATUS_CONFIRMED InternationalWireStatus = "CONFIRMED"
	INTERNATIONALWIRESTATUS_DECLINED InternationalWireStatus = "DECLINED"
	INTERNATIONALWIRESTATUS_PENDING InternationalWireStatus = "PENDING"
)

// All allowed values of InternationalWireStatus enum
var AllowedInternationalWireStatusEnumValues = []InternationalWireStatus{
	"CANCELED",
	"COMPLETED",
	"CONFIRMED",
	"DECLINED",
	"PENDING",
}

func (v *InternationalWireStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InternationalWireStatus(value)
	for _, existing := range AllowedInternationalWireStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InternationalWireStatus", value)
}

// NewInternationalWireStatusFromValue returns a pointer to a valid InternationalWireStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInternationalWireStatusFromValue(v string) (*InternationalWireStatus, error) {
	ev := InternationalWireStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InternationalWireStatus: valid values are %v", v, AllowedInternationalWireStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InternationalWireStatus) IsValid() bool {
	for _, existing := range AllowedInternationalWireStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to international_wire_status value
func (v InternationalWireStatus) Ptr() *InternationalWireStatus {
	return &v
}

type NullableInternationalWireStatus struct {
	value *InternationalWireStatus
	isSet bool
}

func (v NullableInternationalWireStatus) Get() *InternationalWireStatus {
	return v.value
}

func (v *NullableInternationalWireStatus) Set(val *InternationalWireStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInternationalWireStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInternationalWireStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternationalWireStatus(val *InternationalWireStatus) *NullableInternationalWireStatus {
	return &NullableInternationalWireStatus{value: val, isSet: true}
}

func (v NullableInternationalWireStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternationalWireStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

