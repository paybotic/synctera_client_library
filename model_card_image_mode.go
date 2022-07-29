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

// CardImageMode The image mode of a card product. If the card product supports custom card images, this value determines how the images will be handled during card issuance. REQUIRED means that cards issued with this card product must have the ID of an image that has been uploaded. Note that the image does not necessarily have to have been approved prior to the issuance request. This mode is currently disabled. REQUIRED_APPROVED_FIRST means that cards issued with this card product must have the ID of an image that has been uploaded and approved. 
type CardImageMode string

// List of card_image_mode
const (
	CARDIMAGEMODE_REQUIRED CardImageMode = "REQUIRED"
	CARDIMAGEMODE_REQUIRED_APPROVED_FIRST CardImageMode = "REQUIRED_APPROVED_FIRST"
)

// All allowed values of CardImageMode enum
var AllowedCardImageModeEnumValues = []CardImageMode{
	"REQUIRED",
	"REQUIRED_APPROVED_FIRST",
}

func (v *CardImageMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardImageMode(value)
	for _, existing := range AllowedCardImageModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardImageMode", value)
}

// NewCardImageModeFromValue returns a pointer to a valid CardImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardImageModeFromValue(v string) (*CardImageMode, error) {
	ev := CardImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardImageMode: valid values are %v", v, AllowedCardImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardImageMode) IsValid() bool {
	for _, existing := range AllowedCardImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to card_image_mode value
func (v CardImageMode) Ptr() *CardImageMode {
	return &v
}

type NullableCardImageMode struct {
	value *CardImageMode
	isSet bool
}

func (v NullableCardImageMode) Get() *CardImageMode {
	return v.value
}

func (v *NullableCardImageMode) Set(val *CardImageMode) {
	v.value = val
	v.isSet = true
}

func (v NullableCardImageMode) IsSet() bool {
	return v.isSet
}

func (v *NullableCardImageMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardImageMode(val *CardImageMode) *NullableCardImageMode {
	return &NullableCardImageMode{value: val, isSet: true}
}

func (v NullableCardImageMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardImageMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

