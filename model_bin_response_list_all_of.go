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

// BinResponseListAllOf struct for BinResponseListAllOf
type BinResponseListAllOf struct {
	// Array of Bins
	Bins []BinResponse `json:"bins"`
}

// NewBinResponseListAllOf instantiates a new BinResponseListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinResponseListAllOf(bins []BinResponse) *BinResponseListAllOf {
	this := BinResponseListAllOf{}
	this.Bins = bins
	return &this
}

// NewBinResponseListAllOfWithDefaults instantiates a new BinResponseListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinResponseListAllOfWithDefaults() *BinResponseListAllOf {
	this := BinResponseListAllOf{}
	return &this
}

// GetBins returns the Bins field value
func (o *BinResponseListAllOf) GetBins() []BinResponse {
	if o == nil {
		var ret []BinResponse
		return ret
	}

	return o.Bins
}

// GetBinsOk returns a tuple with the Bins field value
// and a boolean to check if the value has been set.
func (o *BinResponseListAllOf) GetBinsOk() ([]BinResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bins, true
}

// SetBins sets field value
func (o *BinResponseListAllOf) SetBins(v []BinResponse) {
	o.Bins = v
}

func (o BinResponseListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bins"] = o.Bins
	}
	return json.Marshal(toSerialize)
}

type NullableBinResponseListAllOf struct {
	value *BinResponseListAllOf
	isSet bool
}

func (v NullableBinResponseListAllOf) Get() *BinResponseListAllOf {
	return v.value
}

func (v *NullableBinResponseListAllOf) Set(val *BinResponseListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBinResponseListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBinResponseListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinResponseListAllOf(val *BinResponseListAllOf) *NullableBinResponseListAllOf {
	return &NullableBinResponseListAllOf{value: val, isSet: true}
}

func (v NullableBinResponseListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinResponseListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
