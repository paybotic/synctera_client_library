/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OutgoingAchListAllOf struct for OutgoingAchListAllOf
type OutgoingAchListAllOf struct {
	// Array of sent ACH transactions.
	Transactions []OutgoingAch `json:"transactions"`
}

// NewOutgoingAchListAllOf instantiates a new OutgoingAchListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutgoingAchListAllOf(transactions []OutgoingAch) *OutgoingAchListAllOf {
	this := OutgoingAchListAllOf{}
	this.Transactions = transactions
	return &this
}

// NewOutgoingAchListAllOfWithDefaults instantiates a new OutgoingAchListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutgoingAchListAllOfWithDefaults() *OutgoingAchListAllOf {
	this := OutgoingAchListAllOf{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *OutgoingAchListAllOf) GetTransactions() []OutgoingAch {
	if o == nil {
		var ret []OutgoingAch
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *OutgoingAchListAllOf) GetTransactionsOk() ([]OutgoingAch, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *OutgoingAchListAllOf) SetTransactions(v []OutgoingAch) {
	o.Transactions = v
}

func (o OutgoingAchListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableOutgoingAchListAllOf struct {
	value *OutgoingAchListAllOf
	isSet bool
}

func (v NullableOutgoingAchListAllOf) Get() *OutgoingAchListAllOf {
	return v.value
}

func (v *NullableOutgoingAchListAllOf) Set(val *OutgoingAchListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOutgoingAchListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOutgoingAchListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutgoingAchListAllOf(val *OutgoingAchListAllOf) *NullableOutgoingAchListAllOf {
	return &NullableOutgoingAchListAllOf{value: val, isSet: true}
}

func (v NullableOutgoingAchListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutgoingAchListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


