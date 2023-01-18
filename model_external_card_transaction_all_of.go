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

// ExternalCardTransactionAllOf struct for ExternalCardTransactionAllOf
type ExternalCardTransactionAllOf struct {
	ExternalCardTransaction CardTransactionData `json:"external_card_transaction"`
	Subtype ExternalCardTransactionSubtypes `json:"subtype"`
}

// NewExternalCardTransactionAllOf instantiates a new ExternalCardTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalCardTransactionAllOf(externalCardTransaction CardTransactionData, subtype ExternalCardTransactionSubtypes) *ExternalCardTransactionAllOf {
	this := ExternalCardTransactionAllOf{}
	this.ExternalCardTransaction = externalCardTransaction
	this.Subtype = subtype
	return &this
}

// NewExternalCardTransactionAllOfWithDefaults instantiates a new ExternalCardTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalCardTransactionAllOfWithDefaults() *ExternalCardTransactionAllOf {
	this := ExternalCardTransactionAllOf{}
	return &this
}

// GetExternalCardTransaction returns the ExternalCardTransaction field value
func (o *ExternalCardTransactionAllOf) GetExternalCardTransaction() CardTransactionData {
	if o == nil {
		var ret CardTransactionData
		return ret
	}

	return o.ExternalCardTransaction
}

// GetExternalCardTransactionOk returns a tuple with the ExternalCardTransaction field value
// and a boolean to check if the value has been set.
func (o *ExternalCardTransactionAllOf) GetExternalCardTransactionOk() (*CardTransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCardTransaction, true
}

// SetExternalCardTransaction sets field value
func (o *ExternalCardTransactionAllOf) SetExternalCardTransaction(v CardTransactionData) {
	o.ExternalCardTransaction = v
}

// GetSubtype returns the Subtype field value
func (o *ExternalCardTransactionAllOf) GetSubtype() ExternalCardTransactionSubtypes {
	if o == nil {
		var ret ExternalCardTransactionSubtypes
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *ExternalCardTransactionAllOf) GetSubtypeOk() (*ExternalCardTransactionSubtypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *ExternalCardTransactionAllOf) SetSubtype(v ExternalCardTransactionSubtypes) {
	o.Subtype = v
}

func (o ExternalCardTransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["external_card_transaction"] = o.ExternalCardTransaction
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableExternalCardTransactionAllOf struct {
	value *ExternalCardTransactionAllOf
	isSet bool
}

func (v NullableExternalCardTransactionAllOf) Get() *ExternalCardTransactionAllOf {
	return v.value
}

func (v *NullableExternalCardTransactionAllOf) Set(val *ExternalCardTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalCardTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalCardTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalCardTransactionAllOf(val *ExternalCardTransactionAllOf) *NullableExternalCardTransactionAllOf {
	return &NullableExternalCardTransactionAllOf{value: val, isSet: true}
}

func (v NullableExternalCardTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalCardTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

