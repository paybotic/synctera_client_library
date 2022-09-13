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

// ChangeChannel How the change was submitted
type ChangeChannel string

// List of change_channel
const (
	CHANGECHANNEL_API ChangeChannel = "API"
	CHANGECHANNEL_ADMIN ChangeChannel = "ADMIN"
	CHANGECHANNEL_SYSTEM ChangeChannel = "SYSTEM"
	CHANGECHANNEL_FRAUD ChangeChannel = "FRAUD"
)

// All allowed values of ChangeChannel enum
var AllowedChangeChannelEnumValues = []ChangeChannel{
	"API",
	"ADMIN",
	"SYSTEM",
	"FRAUD",
}

func (v *ChangeChannel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChangeChannel(value)
	for _, existing := range AllowedChangeChannelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChangeChannel", value)
}

// NewChangeChannelFromValue returns a pointer to a valid ChangeChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeChannelFromValue(v string) (*ChangeChannel, error) {
	ev := ChangeChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeChannel: valid values are %v", v, AllowedChangeChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeChannel) IsValid() bool {
	for _, existing := range AllowedChangeChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to change_channel value
func (v ChangeChannel) Ptr() *ChangeChannel {
	return &v
}

type NullableChangeChannel struct {
	value *ChangeChannel
	isSet bool
}

func (v NullableChangeChannel) Get() *ChangeChannel {
	return v.value
}

func (v *NullableChangeChannel) Set(val *ChangeChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeChannel(val *ChangeChannel) *NullableChangeChannel {
	return &NullableChangeChannel{value: val, isSet: true}
}

func (v NullableChangeChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

