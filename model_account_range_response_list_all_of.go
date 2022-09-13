/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AccountRangeResponseListAllOf struct for AccountRangeResponseListAllOf
type AccountRangeResponseListAllOf struct {
	// Array of Account Ranges
	AccountRanges []AccountRangeResponse `json:"account_ranges"`
}

// NewAccountRangeResponseListAllOf instantiates a new AccountRangeResponseListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRangeResponseListAllOf(accountRanges []AccountRangeResponse) *AccountRangeResponseListAllOf {
	this := AccountRangeResponseListAllOf{}
	this.AccountRanges = accountRanges
	return &this
}

// NewAccountRangeResponseListAllOfWithDefaults instantiates a new AccountRangeResponseListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRangeResponseListAllOfWithDefaults() *AccountRangeResponseListAllOf {
	this := AccountRangeResponseListAllOf{}
	return &this
}

// GetAccountRanges returns the AccountRanges field value
func (o *AccountRangeResponseListAllOf) GetAccountRanges() []AccountRangeResponse {
	if o == nil {
		var ret []AccountRangeResponse
		return ret
	}

	return o.AccountRanges
}

// GetAccountRangesOk returns a tuple with the AccountRanges field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponseListAllOf) GetAccountRangesOk() ([]AccountRangeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountRanges, true
}

// SetAccountRanges sets field value
func (o *AccountRangeResponseListAllOf) SetAccountRanges(v []AccountRangeResponse) {
	o.AccountRanges = v
}

func (o AccountRangeResponseListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_ranges"] = o.AccountRanges
	}
	return json.Marshal(toSerialize)
}

type NullableAccountRangeResponseListAllOf struct {
	value *AccountRangeResponseListAllOf
	isSet bool
}

func (v NullableAccountRangeResponseListAllOf) Get() *AccountRangeResponseListAllOf {
	return v.value
}

func (v *NullableAccountRangeResponseListAllOf) Set(val *AccountRangeResponseListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRangeResponseListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRangeResponseListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRangeResponseListAllOf(val *AccountRangeResponseListAllOf) *NullableAccountRangeResponseListAllOf {
	return &NullableAccountRangeResponseListAllOf{value: val, isSet: true}
}

func (v NullableAccountRangeResponseListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRangeResponseListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


