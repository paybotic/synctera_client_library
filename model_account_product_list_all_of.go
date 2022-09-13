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

// AccountProductListAllOf struct for AccountProductListAllOf
type AccountProductListAllOf struct {
	// Array of account products
	AccountProducts []AccountProduct `json:"account_products"`
}

// NewAccountProductListAllOf instantiates a new AccountProductListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountProductListAllOf(accountProducts []AccountProduct) *AccountProductListAllOf {
	this := AccountProductListAllOf{}
	this.AccountProducts = accountProducts
	return &this
}

// NewAccountProductListAllOfWithDefaults instantiates a new AccountProductListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountProductListAllOfWithDefaults() *AccountProductListAllOf {
	this := AccountProductListAllOf{}
	return &this
}

// GetAccountProducts returns the AccountProducts field value
func (o *AccountProductListAllOf) GetAccountProducts() []AccountProduct {
	if o == nil {
		var ret []AccountProduct
		return ret
	}

	return o.AccountProducts
}

// GetAccountProductsOk returns a tuple with the AccountProducts field value
// and a boolean to check if the value has been set.
func (o *AccountProductListAllOf) GetAccountProductsOk() ([]AccountProduct, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountProducts, true
}

// SetAccountProducts sets field value
func (o *AccountProductListAllOf) SetAccountProducts(v []AccountProduct) {
	o.AccountProducts = v
}

func (o AccountProductListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_products"] = o.AccountProducts
	}
	return json.Marshal(toSerialize)
}

type NullableAccountProductListAllOf struct {
	value *AccountProductListAllOf
	isSet bool
}

func (v NullableAccountProductListAllOf) Get() *AccountProductListAllOf {
	return v.value
}

func (v *NullableAccountProductListAllOf) Set(val *AccountProductListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountProductListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountProductListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountProductListAllOf(val *AccountProductListAllOf) *NullableAccountProductListAllOf {
	return &NullableAccountProductListAllOf{value: val, isSet: true}
}

func (v NullableAccountProductListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountProductListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


