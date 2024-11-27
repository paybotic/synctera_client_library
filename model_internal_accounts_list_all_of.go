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

// InternalAccountsListAllOf struct for InternalAccountsListAllOf
type InternalAccountsListAllOf struct {
	// Array of internal accounts
	InternalAccounts []InternalAccount `json:"internal_accounts"`
}

// NewInternalAccountsListAllOf instantiates a new InternalAccountsListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalAccountsListAllOf(internalAccounts []InternalAccount) *InternalAccountsListAllOf {
	this := InternalAccountsListAllOf{}
	this.InternalAccounts = internalAccounts
	return &this
}

// NewInternalAccountsListAllOfWithDefaults instantiates a new InternalAccountsListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalAccountsListAllOfWithDefaults() *InternalAccountsListAllOf {
	this := InternalAccountsListAllOf{}
	return &this
}

// GetInternalAccounts returns the InternalAccounts field value
func (o *InternalAccountsListAllOf) GetInternalAccounts() []InternalAccount {
	if o == nil {
		var ret []InternalAccount
		return ret
	}

	return o.InternalAccounts
}

// GetInternalAccountsOk returns a tuple with the InternalAccounts field value
// and a boolean to check if the value has been set.
func (o *InternalAccountsListAllOf) GetInternalAccountsOk() ([]InternalAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalAccounts, true
}

// SetInternalAccounts sets field value
func (o *InternalAccountsListAllOf) SetInternalAccounts(v []InternalAccount) {
	o.InternalAccounts = v
}

func (o InternalAccountsListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["internal_accounts"] = o.InternalAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableInternalAccountsListAllOf struct {
	value *InternalAccountsListAllOf
	isSet bool
}

func (v NullableInternalAccountsListAllOf) Get() *InternalAccountsListAllOf {
	return v.value
}

func (v *NullableInternalAccountsListAllOf) Set(val *InternalAccountsListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalAccountsListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalAccountsListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalAccountsListAllOf(val *InternalAccountsListAllOf) *NullableInternalAccountsListAllOf {
	return &NullableInternalAccountsListAllOf{value: val, isSet: true}
}

func (v NullableInternalAccountsListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalAccountsListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
