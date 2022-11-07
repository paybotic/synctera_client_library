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

// VirtualCardResponseAllOf struct for VirtualCardResponseAllOf
type VirtualCardResponseAllOf struct {
	// The bin number
	Bin *string `json:"bin,omitempty"`
	CardBrand CardBrand `json:"card_brand"`
}

// NewVirtualCardResponseAllOf instantiates a new VirtualCardResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualCardResponseAllOf(cardBrand CardBrand) *VirtualCardResponseAllOf {
	this := VirtualCardResponseAllOf{}
	this.CardBrand = cardBrand
	return &this
}

// NewVirtualCardResponseAllOfWithDefaults instantiates a new VirtualCardResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualCardResponseAllOfWithDefaults() *VirtualCardResponseAllOf {
	this := VirtualCardResponseAllOf{}
	return &this
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *VirtualCardResponseAllOf) GetBin() string {
	if o == nil || o.Bin == nil {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponseAllOf) GetBinOk() (*string, bool) {
	if o == nil || o.Bin == nil {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *VirtualCardResponseAllOf) HasBin() bool {
	if o != nil && o.Bin != nil {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *VirtualCardResponseAllOf) SetBin(v string) {
	o.Bin = &v
}

// GetCardBrand returns the CardBrand field value
func (o *VirtualCardResponseAllOf) GetCardBrand() CardBrand {
	if o == nil {
		var ret CardBrand
		return ret
	}

	return o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponseAllOf) GetCardBrandOk() (*CardBrand, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardBrand, true
}

// SetCardBrand sets field value
func (o *VirtualCardResponseAllOf) SetCardBrand(v CardBrand) {
	o.CardBrand = v
}

func (o VirtualCardResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bin != nil {
		toSerialize["bin"] = o.Bin
	}
	if true {
		toSerialize["card_brand"] = o.CardBrand
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualCardResponseAllOf struct {
	value *VirtualCardResponseAllOf
	isSet bool
}

func (v NullableVirtualCardResponseAllOf) Get() *VirtualCardResponseAllOf {
	return v.value
}

func (v *NullableVirtualCardResponseAllOf) Set(val *VirtualCardResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualCardResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualCardResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualCardResponseAllOf(val *VirtualCardResponseAllOf) *NullableVirtualCardResponseAllOf {
	return &NullableVirtualCardResponseAllOf{value: val, isSet: true}
}

func (v NullableVirtualCardResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualCardResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


