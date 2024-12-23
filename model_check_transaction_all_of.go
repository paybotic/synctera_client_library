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

// CheckTransactionAllOf struct for CheckTransactionAllOf
type CheckTransactionAllOf struct {
	CheckTransaction CheckTransactionData     `json:"check_transaction"`
	Subtype          CheckTransactionSubtypes `json:"subtype"`
}

// NewCheckTransactionAllOf instantiates a new CheckTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckTransactionAllOf(checkTransaction CheckTransactionData, subtype CheckTransactionSubtypes) *CheckTransactionAllOf {
	this := CheckTransactionAllOf{}
	this.CheckTransaction = checkTransaction
	this.Subtype = subtype
	return &this
}

// NewCheckTransactionAllOfWithDefaults instantiates a new CheckTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckTransactionAllOfWithDefaults() *CheckTransactionAllOf {
	this := CheckTransactionAllOf{}
	return &this
}

// GetCheckTransaction returns the CheckTransaction field value
func (o *CheckTransactionAllOf) GetCheckTransaction() CheckTransactionData {
	if o == nil {
		var ret CheckTransactionData
		return ret
	}

	return o.CheckTransaction
}

// GetCheckTransactionOk returns a tuple with the CheckTransaction field value
// and a boolean to check if the value has been set.
func (o *CheckTransactionAllOf) GetCheckTransactionOk() (*CheckTransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckTransaction, true
}

// SetCheckTransaction sets field value
func (o *CheckTransactionAllOf) SetCheckTransaction(v CheckTransactionData) {
	o.CheckTransaction = v
}

// GetSubtype returns the Subtype field value
func (o *CheckTransactionAllOf) GetSubtype() CheckTransactionSubtypes {
	if o == nil {
		var ret CheckTransactionSubtypes
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *CheckTransactionAllOf) GetSubtypeOk() (*CheckTransactionSubtypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *CheckTransactionAllOf) SetSubtype(v CheckTransactionSubtypes) {
	o.Subtype = v
}

func (o CheckTransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["check_transaction"] = o.CheckTransaction
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableCheckTransactionAllOf struct {
	value *CheckTransactionAllOf
	isSet bool
}

func (v NullableCheckTransactionAllOf) Get() *CheckTransactionAllOf {
	return v.value
}

func (v *NullableCheckTransactionAllOf) Set(val *CheckTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckTransactionAllOf(val *CheckTransactionAllOf) *NullableCheckTransactionAllOf {
	return &NullableCheckTransactionAllOf{value: val, isSet: true}
}

func (v NullableCheckTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
