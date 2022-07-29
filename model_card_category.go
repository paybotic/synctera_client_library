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

// CardCategory The category of the card
type CardCategory string

// List of card_category
const (
	CARDCATEGORY_CONSUMER CardCategory = "CONSUMER"
	CARDCATEGORY_COMMERCIAL CardCategory = "COMMERCIAL"
)

// All allowed values of CardCategory enum
var AllowedCardCategoryEnumValues = []CardCategory{
	"CONSUMER",
	"COMMERCIAL",
}

func (v *CardCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardCategory(value)
	for _, existing := range AllowedCardCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardCategory", value)
}

// NewCardCategoryFromValue returns a pointer to a valid CardCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardCategoryFromValue(v string) (*CardCategory, error) {
	ev := CardCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardCategory: valid values are %v", v, AllowedCardCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardCategory) IsValid() bool {
	for _, existing := range AllowedCardCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to card_category value
func (v CardCategory) Ptr() *CardCategory {
	return &v
}

type NullableCardCategory struct {
	value *CardCategory
	isSet bool
}

func (v NullableCardCategory) Get() *CardCategory {
	return v.value
}

func (v *NullableCardCategory) Set(val *CardCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCardCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCardCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardCategory(val *CardCategory) *NullableCardCategory {
	return &NullableCardCategory{value: val, isSet: true}
}

func (v NullableCardCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

