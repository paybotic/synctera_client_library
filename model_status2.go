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

// Status2 Status of the person. One of the following: * `ACTIVE` – is an integrator defined status.  Integrators should set a person to active if they believe the person to be qualified for conducting business.  Synctera will combine this status with other statuses such a verification to determine if the person is eligible for specific actions such as initiating transactions or issuing a card. * `DECEASED` – person is deceased. * `DENIED` – customer was turned down. * `DORMANT` – person is no longer active. * `ESCHEAT` – person's assets are abandoned and are property of the state. * `FROZEN` – person's actions are blocked for security, legal, or other reasons. * `INACTIVE` – an inactive status indicating that the person is no longer active. * `PROSPECT` – a potential customer, used for information-gathering and disclosures. * `SANCTION` – person is on a sanctions list and should be carefully monitored. 
type Status2 string

// List of status2
const (
	STATUS2_ACTIVE Status2 = "ACTIVE"
	STATUS2_DECEASED Status2 = "DECEASED"
	STATUS2_DENIED Status2 = "DENIED"
	STATUS2_DORMANT Status2 = "DORMANT"
	STATUS2_ESCHEAT Status2 = "ESCHEAT"
	STATUS2_FROZEN Status2 = "FROZEN"
	STATUS2_INACTIVE Status2 = "INACTIVE"
	STATUS2_PROSPECT Status2 = "PROSPECT"
	STATUS2_SANCTION Status2 = "SANCTION"
)

// All allowed values of Status2 enum
var AllowedStatus2EnumValues = []Status2{
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

func (v *Status2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Status2(value)
	for _, existing := range AllowedStatus2EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Status2", value)
}

// NewStatus2FromValue returns a pointer to a valid Status2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatus2FromValue(v string) (*Status2, error) {
	ev := Status2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Status2: valid values are %v", v, AllowedStatus2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Status2) IsValid() bool {
	for _, existing := range AllowedStatus2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to status2 value
func (v Status2) Ptr() *Status2 {
	return &v
}

type NullableStatus2 struct {
	value *Status2
	isSet bool
}

func (v NullableStatus2) Get() *Status2 {
	return v.value
}

func (v *NullableStatus2) Set(val *Status2) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus2) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus2(val *Status2) *NullableStatus2 {
	return &NullableStatus2{value: val, isSet: true}
}

func (v NullableStatus2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
