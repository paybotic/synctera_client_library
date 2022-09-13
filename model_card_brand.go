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

// CardBrand The brand of a card product
type CardBrand string

// List of card_brand
const (
	CARDBRAND_MASTERCARD CardBrand = "MASTERCARD"
	CARDBRAND_VISA CardBrand = "VISA"
)

// All allowed values of CardBrand enum
var AllowedCardBrandEnumValues = []CardBrand{
	"MASTERCARD",
	"VISA",
}

func (v *CardBrand) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardBrand(value)
	for _, existing := range AllowedCardBrandEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardBrand", value)
}

// NewCardBrandFromValue returns a pointer to a valid CardBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardBrandFromValue(v string) (*CardBrand, error) {
	ev := CardBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardBrand: valid values are %v", v, AllowedCardBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardBrand) IsValid() bool {
	for _, existing := range AllowedCardBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to card_brand value
func (v CardBrand) Ptr() *CardBrand {
	return &v
}

type NullableCardBrand struct {
	value *CardBrand
	isSet bool
}

func (v NullableCardBrand) Get() *CardBrand {
	return v.value
}

func (v *NullableCardBrand) Set(val *CardBrand) {
	v.value = val
	v.isSet = true
}

func (v NullableCardBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableCardBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardBrand(val *CardBrand) *NullableCardBrand {
	return &NullableCardBrand{value: val, isSet: true}
}

func (v NullableCardBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

