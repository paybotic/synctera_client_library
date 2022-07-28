/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DebitNetworkResponseList struct for DebitNetworkResponseList
type DebitNetworkResponseList struct {
	// Array of debit networks
	DebitNetworks []DebitNetworkResponse `json:"debit_networks"`
}

// NewDebitNetworkResponseList instantiates a new DebitNetworkResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebitNetworkResponseList(debitNetworks []DebitNetworkResponse) *DebitNetworkResponseList {
	this := DebitNetworkResponseList{}
	this.DebitNetworks = debitNetworks
	return &this
}

// NewDebitNetworkResponseListWithDefaults instantiates a new DebitNetworkResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebitNetworkResponseListWithDefaults() *DebitNetworkResponseList {
	this := DebitNetworkResponseList{}
	return &this
}

// GetDebitNetworks returns the DebitNetworks field value
func (o *DebitNetworkResponseList) GetDebitNetworks() []DebitNetworkResponse {
	if o == nil {
		var ret []DebitNetworkResponse
		return ret
	}

	return o.DebitNetworks
}

// GetDebitNetworksOk returns a tuple with the DebitNetworks field value
// and a boolean to check if the value has been set.
func (o *DebitNetworkResponseList) GetDebitNetworksOk() ([]DebitNetworkResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebitNetworks, true
}

// SetDebitNetworks sets field value
func (o *DebitNetworkResponseList) SetDebitNetworks(v []DebitNetworkResponse) {
	o.DebitNetworks = v
}

func (o DebitNetworkResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["debit_networks"] = o.DebitNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableDebitNetworkResponseList struct {
	value *DebitNetworkResponseList
	isSet bool
}

func (v NullableDebitNetworkResponseList) Get() *DebitNetworkResponseList {
	return v.value
}

func (v *NullableDebitNetworkResponseList) Set(val *DebitNetworkResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableDebitNetworkResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableDebitNetworkResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebitNetworkResponseList(val *DebitNetworkResponseList) *NullableDebitNetworkResponseList {
	return &NullableDebitNetworkResponseList{value: val, isSet: true}
}

func (v NullableDebitNetworkResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebitNetworkResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


