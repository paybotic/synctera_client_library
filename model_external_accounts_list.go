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

// ExternalAccountsList struct for ExternalAccountsList
type ExternalAccountsList struct {
	// Array of external accounts
	ExternalAccounts []ExternalAccount `json:"external_accounts"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewExternalAccountsList instantiates a new ExternalAccountsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccountsList(externalAccounts []ExternalAccount) *ExternalAccountsList {
	this := ExternalAccountsList{}
	this.ExternalAccounts = externalAccounts
	return &this
}

// NewExternalAccountsListWithDefaults instantiates a new ExternalAccountsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountsListWithDefaults() *ExternalAccountsList {
	this := ExternalAccountsList{}
	return &this
}

// GetExternalAccounts returns the ExternalAccounts field value
func (o *ExternalAccountsList) GetExternalAccounts() []ExternalAccount {
	if o == nil {
		var ret []ExternalAccount
		return ret
	}

	return o.ExternalAccounts
}

// GetExternalAccountsOk returns a tuple with the ExternalAccounts field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountsList) GetExternalAccountsOk() ([]ExternalAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalAccounts, true
}

// SetExternalAccounts sets field value
func (o *ExternalAccountsList) SetExternalAccounts(v []ExternalAccount) {
	o.ExternalAccounts = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ExternalAccountsList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountsList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ExternalAccountsList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ExternalAccountsList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ExternalAccountsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["external_accounts"] = o.ExternalAccounts
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableExternalAccountsList struct {
	value *ExternalAccountsList
	isSet bool
}

func (v NullableExternalAccountsList) Get() *ExternalAccountsList {
	return v.value
}

func (v *NullableExternalAccountsList) Set(val *ExternalAccountsList) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccountsList) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccountsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccountsList(val *ExternalAccountsList) *NullableExternalAccountsList {
	return &NullableExternalAccountsList{value: val, isSet: true}
}

func (v NullableExternalAccountsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccountsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


