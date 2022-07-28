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

// AccountRangeResponseList struct for AccountRangeResponseList
type AccountRangeResponseList struct {
	// Array of Account Ranges
	AccountRanges []AccountRangeResponse `json:"account_ranges"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewAccountRangeResponseList instantiates a new AccountRangeResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRangeResponseList(accountRanges []AccountRangeResponse) *AccountRangeResponseList {
	this := AccountRangeResponseList{}
	this.AccountRanges = accountRanges
	return &this
}

// NewAccountRangeResponseListWithDefaults instantiates a new AccountRangeResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRangeResponseListWithDefaults() *AccountRangeResponseList {
	this := AccountRangeResponseList{}
	return &this
}

// GetAccountRanges returns the AccountRanges field value
func (o *AccountRangeResponseList) GetAccountRanges() []AccountRangeResponse {
	if o == nil {
		var ret []AccountRangeResponse
		return ret
	}

	return o.AccountRanges
}

// GetAccountRangesOk returns a tuple with the AccountRanges field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponseList) GetAccountRangesOk() ([]AccountRangeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountRanges, true
}

// SetAccountRanges sets field value
func (o *AccountRangeResponseList) SetAccountRanges(v []AccountRangeResponse) {
	o.AccountRanges = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *AccountRangeResponseList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRangeResponseList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *AccountRangeResponseList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *AccountRangeResponseList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o AccountRangeResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_ranges"] = o.AccountRanges
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableAccountRangeResponseList struct {
	value *AccountRangeResponseList
	isSet bool
}

func (v NullableAccountRangeResponseList) Get() *AccountRangeResponseList {
	return v.value
}

func (v *NullableAccountRangeResponseList) Set(val *AccountRangeResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRangeResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRangeResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRangeResponseList(val *AccountRangeResponseList) *NullableAccountRangeResponseList {
	return &NullableAccountRangeResponseList{value: val, isSet: true}
}

func (v NullableAccountRangeResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRangeResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


