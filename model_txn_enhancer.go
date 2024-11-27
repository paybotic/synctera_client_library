/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TxnEnhancer Whether to use a transaction enhancer and/or which vendor to use. Enhancer is a third party service that provides additional data for card transactions. MX is included by default.
type TxnEnhancer string

// List of txn_enhancer
const (
	TXNENHANCER_MX   TxnEnhancer = "MX"
	TXNENHANCER_NONE TxnEnhancer = "NONE"
)

// All allowed values of TxnEnhancer enum
var AllowedTxnEnhancerEnumValues = []TxnEnhancer{
	"MX",
	"NONE",
}

func (v *TxnEnhancer) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TxnEnhancer(value)
	for _, existing := range AllowedTxnEnhancerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TxnEnhancer", value)
}

// NewTxnEnhancerFromValue returns a pointer to a valid TxnEnhancer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTxnEnhancerFromValue(v string) (*TxnEnhancer, error) {
	ev := TxnEnhancer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TxnEnhancer: valid values are %v", v, AllowedTxnEnhancerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TxnEnhancer) IsValid() bool {
	for _, existing := range AllowedTxnEnhancerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to txn_enhancer value
func (v TxnEnhancer) Ptr() *TxnEnhancer {
	return &v
}

type NullableTxnEnhancer struct {
	value *TxnEnhancer
	isSet bool
}

func (v NullableTxnEnhancer) Get() *TxnEnhancer {
	return v.value
}

func (v *NullableTxnEnhancer) Set(val *TxnEnhancer) {
	v.value = val
	v.isSet = true
}

func (v NullableTxnEnhancer) IsSet() bool {
	return v.isSet
}

func (v *NullableTxnEnhancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxnEnhancer(val *TxnEnhancer) *NullableTxnEnhancer {
	return &NullableTxnEnhancer{value: val, isSet: true}
}

func (v NullableTxnEnhancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxnEnhancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
