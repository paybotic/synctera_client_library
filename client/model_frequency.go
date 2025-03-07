/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// Frequency The frequency of the expense
type Frequency string

// List of frequency
const (
	FREQUENCY_BI_WEEKLY    Frequency = "BI_WEEKLY"
	FREQUENCY_DAILY        Frequency = "DAILY"
	FREQUENCY_MONTHLY      Frequency = "MONTHLY"
	FREQUENCY_ONE_TIME     Frequency = "ONE_TIME"
	FREQUENCY_OTHER        Frequency = "OTHER"
	FREQUENCY_QUARTERLY    Frequency = "QUARTERLY"
	FREQUENCY_SEMI_MONTHLY Frequency = "SEMI_MONTHLY"
	FREQUENCY_WEEKLY       Frequency = "WEEKLY"
	FREQUENCY_YEARLY       Frequency = "YEARLY"
)

// All allowed values of Frequency enum
var AllowedFrequencyEnumValues = []Frequency{
	"BI_WEEKLY",
	"DAILY",
	"MONTHLY",
	"ONE_TIME",
	"OTHER",
	"QUARTERLY",
	"SEMI_MONTHLY",
	"WEEKLY",
	"YEARLY",
}

func (v *Frequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Frequency(value)
	for _, existing := range AllowedFrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Frequency", value)
}

// NewFrequencyFromValue returns a pointer to a valid Frequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFrequencyFromValue(v string) (*Frequency, error) {
	ev := Frequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Frequency: valid values are %v", v, AllowedFrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Frequency) IsValid() bool {
	for _, existing := range AllowedFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to frequency value
func (v Frequency) Ptr() *Frequency {
	return &v
}

type NullableFrequency struct {
	value *Frequency
	isSet bool
}

func (v NullableFrequency) Get() *Frequency {
	return v.value
}

func (v *NullableFrequency) Set(val *Frequency) {
	v.value = val
	v.isSet = true
}

func (v NullableFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrequency(val *Frequency) *NullableFrequency {
	return &NullableFrequency{value: val, isSet: true}
}

func (v NullableFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
