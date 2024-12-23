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

// WireTransactionAllOf struct for WireTransactionAllOf
type WireTransactionAllOf struct {
	Subtype         WireTransactionSubtypes `json:"subtype"`
	WireTransaction WireTransactionData     `json:"wire_transaction"`
}

// NewWireTransactionAllOf instantiates a new WireTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireTransactionAllOf(subtype WireTransactionSubtypes, wireTransaction WireTransactionData) *WireTransactionAllOf {
	this := WireTransactionAllOf{}
	this.Subtype = subtype
	this.WireTransaction = wireTransaction
	return &this
}

// NewWireTransactionAllOfWithDefaults instantiates a new WireTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireTransactionAllOfWithDefaults() *WireTransactionAllOf {
	this := WireTransactionAllOf{}
	return &this
}

// GetSubtype returns the Subtype field value
func (o *WireTransactionAllOf) GetSubtype() WireTransactionSubtypes {
	if o == nil {
		var ret WireTransactionSubtypes
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *WireTransactionAllOf) GetSubtypeOk() (*WireTransactionSubtypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *WireTransactionAllOf) SetSubtype(v WireTransactionSubtypes) {
	o.Subtype = v
}

// GetWireTransaction returns the WireTransaction field value
func (o *WireTransactionAllOf) GetWireTransaction() WireTransactionData {
	if o == nil {
		var ret WireTransactionData
		return ret
	}

	return o.WireTransaction
}

// GetWireTransactionOk returns a tuple with the WireTransaction field value
// and a boolean to check if the value has been set.
func (o *WireTransactionAllOf) GetWireTransactionOk() (*WireTransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WireTransaction, true
}

// SetWireTransaction sets field value
func (o *WireTransactionAllOf) SetWireTransaction(v WireTransactionData) {
	o.WireTransaction = v
}

func (o WireTransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["wire_transaction"] = o.WireTransaction
	}
	return json.Marshal(toSerialize)
}

type NullableWireTransactionAllOf struct {
	value *WireTransactionAllOf
	isSet bool
}

func (v NullableWireTransactionAllOf) Get() *WireTransactionAllOf {
	return v.value
}

func (v *NullableWireTransactionAllOf) Set(val *WireTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWireTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWireTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireTransactionAllOf(val *WireTransactionAllOf) *NullableWireTransactionAllOf {
	return &NullableWireTransactionAllOf{value: val, isSet: true}
}

func (v NullableWireTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
