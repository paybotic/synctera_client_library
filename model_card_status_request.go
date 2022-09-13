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

// CardStatusRequest The status indicating the card lifecycle state
type CardStatusRequest string

// List of card_status_request
const (
	CARDSTATUSREQUEST_ACTIVE CardStatusRequest = "ACTIVE"
	CARDSTATUSREQUEST_SUSPENDED CardStatusRequest = "SUSPENDED"
	CARDSTATUSREQUEST_TERMINATED CardStatusRequest = "TERMINATED"
)

// All allowed values of CardStatusRequest enum
var AllowedCardStatusRequestEnumValues = []CardStatusRequest{
	"ACTIVE",
	"SUSPENDED",
	"TERMINATED",
}

func (v *CardStatusRequest) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardStatusRequest(value)
	for _, existing := range AllowedCardStatusRequestEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardStatusRequest", value)
}

// NewCardStatusRequestFromValue returns a pointer to a valid CardStatusRequest
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardStatusRequestFromValue(v string) (*CardStatusRequest, error) {
	ev := CardStatusRequest(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardStatusRequest: valid values are %v", v, AllowedCardStatusRequestEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardStatusRequest) IsValid() bool {
	for _, existing := range AllowedCardStatusRequestEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to card_status_request value
func (v CardStatusRequest) Ptr() *CardStatusRequest {
	return &v
}

type NullableCardStatusRequest struct {
	value *CardStatusRequest
	isSet bool
}

func (v NullableCardStatusRequest) Get() *CardStatusRequest {
	return v.value
}

func (v *NullableCardStatusRequest) Set(val *CardStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardStatusRequest(val *CardStatusRequest) *NullableCardStatusRequest {
	return &NullableCardStatusRequest{value: val, isSet: true}
}

func (v NullableCardStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

