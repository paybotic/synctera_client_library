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

// Status1 Status of the person. One of the following: * `ACTIVE` – is an integrator defined status.  Integrators should set a person to active if they believe the person to be qualified for conducting business.  Synctera will combine this status with other statuses such a verification to determine if the person is eligible for specific actions such as initiating transactions or issuing a card. * `DECEASED` – person is deceased. * `DENIED` – customer was turned down. * `DORMANT` – person is no longer active. * `ESCHEAT` – person's assets are abandoned and are property of the state. * `FROZEN` – person's actions are blocked for security, legal, or other reasons. * `INACTIVE` – an inactive status indicating that the person is no longer active. * `PROSPECT` – a potential customer, used for information-gathering and disclosures. * `SANCTION` – person is on a sanctions list and should be carefully monitored. 
type Status1 string

// List of status1
const (
	STATUS1_ACTIVE Status1 = "ACTIVE"
	STATUS1_DECEASED Status1 = "DECEASED"
	STATUS1_DENIED Status1 = "DENIED"
	STATUS1_DORMANT Status1 = "DORMANT"
	STATUS1_ESCHEAT Status1 = "ESCHEAT"
	STATUS1_FROZEN Status1 = "FROZEN"
	STATUS1_INACTIVE Status1 = "INACTIVE"
	STATUS1_PROSPECT Status1 = "PROSPECT"
	STATUS1_SANCTION Status1 = "SANCTION"
)

// All allowed values of Status1 enum
var AllowedStatus1EnumValues = []Status1{
	"ACTIVE",
	"DECEASED",
	"DENIED",
	"DORMANT",
	"ESCHEAT",
	"FROZEN",
	"INACTIVE",
	"PROSPECT",
	"SANCTION",
}

func (v *Status1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Status1(value)
	for _, existing := range AllowedStatus1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Status1", value)
}

// NewStatus1FromValue returns a pointer to a valid Status1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatus1FromValue(v string) (*Status1, error) {
	ev := Status1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Status1: valid values are %v", v, AllowedStatus1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Status1) IsValid() bool {
	for _, existing := range AllowedStatus1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to status1 value
func (v Status1) Ptr() *Status1 {
	return &v
}

type NullableStatus1 struct {
	value *Status1
	isSet bool
}

func (v NullableStatus1) Get() *Status1 {
	return v.value
}

func (v *NullableStatus1) Set(val *Status1) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus1) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus1(val *Status1) *NullableStatus1 {
	return &NullableStatus1{value: val, isSet: true}
}

func (v NullableStatus1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

