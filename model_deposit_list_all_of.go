/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DepositListAllOf struct for DepositListAllOf
type DepositListAllOf struct {
	// Array of  Remote Check Deposits
	Deposits []Deposit `json:"deposits"`
}

// NewDepositListAllOf instantiates a new DepositListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositListAllOf(deposits []Deposit) *DepositListAllOf {
	this := DepositListAllOf{}
	this.Deposits = deposits
	return &this
}

// NewDepositListAllOfWithDefaults instantiates a new DepositListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositListAllOfWithDefaults() *DepositListAllOf {
	this := DepositListAllOf{}
	return &this
}

// GetDeposits returns the Deposits field value
func (o *DepositListAllOf) GetDeposits() []Deposit {
	if o == nil {
		var ret []Deposit
		return ret
	}

	return o.Deposits
}

// GetDepositsOk returns a tuple with the Deposits field value
// and a boolean to check if the value has been set.
func (o *DepositListAllOf) GetDepositsOk() ([]Deposit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deposits, true
}

// SetDeposits sets field value
func (o *DepositListAllOf) SetDeposits(v []Deposit) {
	o.Deposits = v
}

func (o DepositListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deposits"] = o.Deposits
	}
	return json.Marshal(toSerialize)
}

type NullableDepositListAllOf struct {
	value *DepositListAllOf
	isSet bool
}

func (v NullableDepositListAllOf) Get() *DepositListAllOf {
	return v.value
}

func (v *NullableDepositListAllOf) Set(val *DepositListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositListAllOf(val *DepositListAllOf) *NullableDepositListAllOf {
	return &NullableDepositListAllOf{value: val, isSet: true}
}

func (v NullableDepositListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
