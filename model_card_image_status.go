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

// CardImageStatus The status of a custom card image
type CardImageStatus string

// List of card_image_status
const (
	CARDIMAGESTATUS_NOT_UPLOADED CardImageStatus = "NOT_UPLOADED"
	CARDIMAGESTATUS_UNREVIEWED CardImageStatus = "UNREVIEWED"
	CARDIMAGESTATUS_APPROVED CardImageStatus = "APPROVED"
)

// All allowed values of CardImageStatus enum
var AllowedCardImageStatusEnumValues = []CardImageStatus{
	"NOT_UPLOADED",
	"UNREVIEWED",
	"APPROVED",
	"REJECTED",
}

func (v *CardImageStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardImageStatus(value)
	for _, existing := range AllowedCardImageStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardImageStatus", value)
}

// NewCardImageStatusFromValue returns a pointer to a valid CardImageStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardImageStatusFromValue(v string) (*CardImageStatus, error) {
	ev := CardImageStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardImageStatus: valid values are %v", v, AllowedCardImageStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardImageStatus) IsValid() bool {
	for _, existing := range AllowedCardImageStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to card_image_status value
func (v CardImageStatus) Ptr() *CardImageStatus {
	return &v
}

type NullableCardImageStatus struct {
	value *CardImageStatus
	isSet bool
}

func (v NullableCardImageStatus) Get() *CardImageStatus {
	return v.value
}

func (v *NullableCardImageStatus) Set(val *CardImageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCardImageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCardImageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardImageStatus(val *CardImageStatus) *NullableCardImageStatus {
	return &NullableCardImageStatus{value: val, isSet: true}
}

func (v NullableCardImageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardImageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

