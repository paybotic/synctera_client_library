/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ApplicationType1 Type of Credit Application
type ApplicationType1 string

// List of application_type1
const (
	APPLICATIONTYPE1_LINE_OF_CREDIT ApplicationType1 = "LINE_OF_CREDIT"
)

// All allowed values of ApplicationType1 enum
var AllowedApplicationType1EnumValues = []ApplicationType1{
	"LINE_OF_CREDIT",
}

func (v *ApplicationType1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplicationType1(value)
	for _, existing := range AllowedApplicationType1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplicationType1", value)
}

// NewApplicationType1FromValue returns a pointer to a valid ApplicationType1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplicationType1FromValue(v string) (*ApplicationType1, error) {
	ev := ApplicationType1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplicationType1: valid values are %v", v, AllowedApplicationType1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplicationType1) IsValid() bool {
	for _, existing := range AllowedApplicationType1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to application_type1 value
func (v ApplicationType1) Ptr() *ApplicationType1 {
	return &v
}

type NullableApplicationType1 struct {
	value *ApplicationType1
	isSet bool
}

func (v NullableApplicationType1) Get() *ApplicationType1 {
	return v.value
}

func (v *NullableApplicationType1) Set(val *ApplicationType1) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationType1) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationType1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationType1(val *ApplicationType1) *NullableApplicationType1 {
	return &NullableApplicationType1{value: val, isSet: true}
}

func (v NullableApplicationType1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationType1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
