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

// CardProductType The type of the card product
type CardProductType string

// List of card_product_type
const (
	CARDPRODUCTTYPE_DEBIT CardProductType = "DEBIT"
	CARDPRODUCTTYPE_CREDIT CardProductType = "CREDIT"
	CARDPRODUCTTYPE_PREPAID CardProductType = "PREPAID"
)

// All allowed values of CardProductType enum
var AllowedCardProductTypeEnumValues = []CardProductType{
	"DEBIT",
	"CREDIT",
	"PREPAID",
}

func (v *CardProductType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardProductType(value)
	for _, existing := range AllowedCardProductTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardProductType", value)
}

// NewCardProductTypeFromValue returns a pointer to a valid CardProductType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardProductTypeFromValue(v string) (*CardProductType, error) {
	ev := CardProductType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardProductType: valid values are %v", v, AllowedCardProductTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardProductType) IsValid() bool {
	for _, existing := range AllowedCardProductTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to card_product_type value
func (v CardProductType) Ptr() *CardProductType {
	return &v
}

type NullableCardProductType struct {
	value *CardProductType
	isSet bool
}

func (v NullableCardProductType) Get() *CardProductType {
	return v.value
}

func (v *NullableCardProductType) Set(val *CardProductType) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProductType) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProductType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProductType(val *CardProductType) *NullableCardProductType {
	return &NullableCardProductType{value: val, isSet: true}
}

func (v NullableCardProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProductType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

