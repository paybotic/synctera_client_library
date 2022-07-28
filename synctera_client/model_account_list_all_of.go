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

// AccountListAllOf struct for AccountListAllOf
type AccountListAllOf struct {
	// Array of Accounts
	Accounts []AccountGenericResponse `json:"accounts"`
}

// NewAccountListAllOf instantiates a new AccountListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountListAllOf(accounts []AccountGenericResponse) *AccountListAllOf {
	this := AccountListAllOf{}
	this.Accounts = accounts
	return &this
}

// NewAccountListAllOfWithDefaults instantiates a new AccountListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountListAllOfWithDefaults() *AccountListAllOf {
	this := AccountListAllOf{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *AccountListAllOf) GetAccounts() []AccountGenericResponse {
	if o == nil {
		var ret []AccountGenericResponse
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *AccountListAllOf) GetAccountsOk() ([]AccountGenericResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *AccountListAllOf) SetAccounts(v []AccountGenericResponse) {
	o.Accounts = v
}

func (o AccountListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableAccountListAllOf struct {
	value *AccountListAllOf
	isSet bool
}

func (v NullableAccountListAllOf) Get() *AccountListAllOf {
	return v.value
}

func (v *NullableAccountListAllOf) Set(val *AccountListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountListAllOf(val *AccountListAllOf) *NullableAccountListAllOf {
	return &NullableAccountListAllOf{value: val, isSet: true}
}

func (v NullableAccountListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


