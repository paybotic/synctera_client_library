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

// AccountProductList struct for AccountProductList
type AccountProductList struct {
	// Array of account products
	AccountProducts []AccountProduct `json:"account_products"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewAccountProductList instantiates a new AccountProductList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountProductList(accountProducts []AccountProduct) *AccountProductList {
	this := AccountProductList{}
	this.AccountProducts = accountProducts
	return &this
}

// NewAccountProductListWithDefaults instantiates a new AccountProductList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountProductListWithDefaults() *AccountProductList {
	this := AccountProductList{}
	return &this
}

// GetAccountProducts returns the AccountProducts field value
func (o *AccountProductList) GetAccountProducts() []AccountProduct {
	if o == nil {
		var ret []AccountProduct
		return ret
	}

	return o.AccountProducts
}

// GetAccountProductsOk returns a tuple with the AccountProducts field value
// and a boolean to check if the value has been set.
func (o *AccountProductList) GetAccountProductsOk() ([]AccountProduct, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountProducts, true
}

// SetAccountProducts sets field value
func (o *AccountProductList) SetAccountProducts(v []AccountProduct) {
	o.AccountProducts = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *AccountProductList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountProductList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *AccountProductList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *AccountProductList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o AccountProductList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_products"] = o.AccountProducts
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableAccountProductList struct {
	value *AccountProductList
	isSet bool
}

func (v NullableAccountProductList) Get() *AccountProductList {
	return v.value
}

func (v *NullableAccountProductList) Set(val *AccountProductList) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountProductList) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountProductList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountProductList(val *AccountProductList) *NullableAccountProductList {
	return &NullableAccountProductList{value: val, isSet: true}
}

func (v NullableAccountProductList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountProductList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


