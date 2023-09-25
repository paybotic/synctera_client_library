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

// PaymentType The type of payment to which a spend control applies. If this is not set, the spend control applies to all spending, regardless of payment type. 
type PaymentType string

// List of payment_type
const (
	PAYMENTTYPE_ACH PaymentType = "ACH"
	PAYMENTTYPE_CARD PaymentType = "CARD"
	PAYMENTTYPE_CASH PaymentType = "CASH"
	PAYMENTTYPE_CHECK PaymentType = "CHECK"
	PAYMENTTYPE_EFT_CA PaymentType = "EFT_CA"
	PAYMENTTYPE_EXTERNAL_CARD PaymentType = "EXTERNAL_CARD"
	PAYMENTTYPE_INTERNAL_TRANSFER PaymentType = "INTERNAL_TRANSFER"
	PAYMENTTYPE_WIRE PaymentType = "WIRE"
)

// All allowed values of PaymentType enum
var AllowedPaymentTypeEnumValues = []PaymentType{
	"ACH",
	"CARD",
	"CASH",
	"CHECK",
	"EFT_CA",
	"EXTERNAL_CARD",
	"INTERNAL_TRANSFER",
	"WIRE",
}

func (v *PaymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentType(value)
	for _, existing := range AllowedPaymentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentType", value)
}

// NewPaymentTypeFromValue returns a pointer to a valid PaymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentTypeFromValue(v string) (*PaymentType, error) {
	ev := PaymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentType: valid values are %v", v, AllowedPaymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentType) IsValid() bool {
	for _, existing := range AllowedPaymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to payment_type value
func (v PaymentType) Ptr() *PaymentType {
	return &v
}

type NullablePaymentType struct {
	value *PaymentType
	isSet bool
}

func (v NullablePaymentType) Get() *PaymentType {
	return v.value
}

func (v *NullablePaymentType) Set(val *PaymentType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentType(val *PaymentType) *NullablePaymentType {
	return &NullablePaymentType{value: val, isSet: true}
}

func (v NullablePaymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

