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

// SpendControlResponseAllOf struct for SpendControlResponseAllOf
type SpendControlResponseAllOf struct {
	// A count of how many accounts are using this spend control
	NumberOfRelatedAccounts int32 `json:"number_of_related_accounts"`
}

// NewSpendControlResponseAllOf instantiates a new SpendControlResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendControlResponseAllOf(numberOfRelatedAccounts int32) *SpendControlResponseAllOf {
	this := SpendControlResponseAllOf{}
	this.NumberOfRelatedAccounts = numberOfRelatedAccounts
	return &this
}

// NewSpendControlResponseAllOfWithDefaults instantiates a new SpendControlResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendControlResponseAllOfWithDefaults() *SpendControlResponseAllOf {
	this := SpendControlResponseAllOf{}
	return &this
}

// GetNumberOfRelatedAccounts returns the NumberOfRelatedAccounts field value
func (o *SpendControlResponseAllOf) GetNumberOfRelatedAccounts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfRelatedAccounts
}

// GetNumberOfRelatedAccountsOk returns a tuple with the NumberOfRelatedAccounts field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponseAllOf) GetNumberOfRelatedAccountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfRelatedAccounts, true
}

// SetNumberOfRelatedAccounts sets field value
func (o *SpendControlResponseAllOf) SetNumberOfRelatedAccounts(v int32) {
	o.NumberOfRelatedAccounts = v
}

func (o SpendControlResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["number_of_related_accounts"] = o.NumberOfRelatedAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableSpendControlResponseAllOf struct {
	value *SpendControlResponseAllOf
	isSet bool
}

func (v NullableSpendControlResponseAllOf) Get() *SpendControlResponseAllOf {
	return v.value
}

func (v *NullableSpendControlResponseAllOf) Set(val *SpendControlResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendControlResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendControlResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendControlResponseAllOf(val *SpendControlResponseAllOf) *NullableSpendControlResponseAllOf {
	return &NullableSpendControlResponseAllOf{value: val, isSet: true}
}

func (v NullableSpendControlResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendControlResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
