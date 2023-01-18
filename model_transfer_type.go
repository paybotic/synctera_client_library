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

// TransferType Type of operation for transfer.  Type | Description --- | --- PUSH | Push fund to the external card from the account_id PULL | Pull funds from the external card to the account_id. PULL_REVERSAL | Reversed pull funds transfer from an external card. 
type TransferType string

// List of transfer_type
const (
	TRANSFERTYPE_PULL TransferType = "PULL"
	TRANSFERTYPE_PUSH TransferType = "PUSH"
	TRANSFERTYPE_PULL_REVERSAL TransferType = "PULL_REVERSAL"
)

// All allowed values of TransferType enum
var AllowedTransferTypeEnumValues = []TransferType{
	"PULL",
	"PUSH",
	"PULL_REVERSAL",
}

func (v *TransferType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferType(value)
	for _, existing := range AllowedTransferTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferType", value)
}

// NewTransferTypeFromValue returns a pointer to a valid TransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferTypeFromValue(v string) (*TransferType, error) {
	ev := TransferType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferType: valid values are %v", v, AllowedTransferTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferType) IsValid() bool {
	for _, existing := range AllowedTransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to transfer_type value
func (v TransferType) Ptr() *TransferType {
	return &v
}

type NullableTransferType struct {
	value *TransferType
	isSet bool
}

func (v NullableTransferType) Get() *TransferType {
	return v.value
}

func (v *NullableTransferType) Set(val *TransferType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferType(val *TransferType) *NullableTransferType {
	return &NullableTransferType{value: val, isSet: true}
}

func (v NullableTransferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

