/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PhysicalCardFormat The format of a physical card product
type PhysicalCardFormat string

// List of physical_card_format
const (
	MAGNETIC_STRIPE PhysicalCardFormat = "MAGNETIC_STRIPE"
	CHIP PhysicalCardFormat = "CHIP"
	CONTACT PhysicalCardFormat = "CONTACT"
	CONTACTLESS PhysicalCardFormat = "CONTACTLESS"
	PHYSICAL_COMBO PhysicalCardFormat = "PHYSICAL_COMBO"
)

// All allowed values of PhysicalCardFormat enum
var AllowedPhysicalCardFormatEnumValues = []PhysicalCardFormat{
	"MAGNETIC_STRIPE",
	"CHIP",
	"CONTACT",
	"CONTACTLESS",
	"PHYSICAL_COMBO",
}

func (v *PhysicalCardFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PhysicalCardFormat(value)
	for _, existing := range AllowedPhysicalCardFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PhysicalCardFormat", value)
}

// NewPhysicalCardFormatFromValue returns a pointer to a valid PhysicalCardFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPhysicalCardFormatFromValue(v string) (*PhysicalCardFormat, error) {
	ev := PhysicalCardFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PhysicalCardFormat: valid values are %v", v, AllowedPhysicalCardFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PhysicalCardFormat) IsValid() bool {
	for _, existing := range AllowedPhysicalCardFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to physical_card_format value
func (v PhysicalCardFormat) Ptr() *PhysicalCardFormat {
	return &v
}

type NullablePhysicalCardFormat struct {
	value *PhysicalCardFormat
	isSet bool
}

func (v NullablePhysicalCardFormat) Get() *PhysicalCardFormat {
	return v.value
}

func (v *NullablePhysicalCardFormat) Set(val *PhysicalCardFormat) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalCardFormat) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalCardFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalCardFormat(val *PhysicalCardFormat) *NullablePhysicalCardFormat {
	return &NullablePhysicalCardFormat{value: val, isSet: true}
}

func (v NullablePhysicalCardFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalCardFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

