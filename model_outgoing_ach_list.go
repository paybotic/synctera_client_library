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

// OutgoingAchList struct for OutgoingAchList
type OutgoingAchList struct {
	// Array of outgoing ACH transactions.
	Transactions []OutgoingAch `json:"transactions"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewOutgoingAchList instantiates a new OutgoingAchList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutgoingAchList(transactions []OutgoingAch) *OutgoingAchList {
	this := OutgoingAchList{}
	this.Transactions = transactions
	return &this
}

// NewOutgoingAchListWithDefaults instantiates a new OutgoingAchList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutgoingAchListWithDefaults() *OutgoingAchList {
	this := OutgoingAchList{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *OutgoingAchList) GetTransactions() []OutgoingAch {
	if o == nil {
		var ret []OutgoingAch
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *OutgoingAchList) GetTransactionsOk() ([]OutgoingAch, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *OutgoingAchList) SetTransactions(v []OutgoingAch) {
	o.Transactions = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *OutgoingAchList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAchList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *OutgoingAchList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *OutgoingAchList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o OutgoingAchList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableOutgoingAchList struct {
	value *OutgoingAchList
	isSet bool
}

func (v NullableOutgoingAchList) Get() *OutgoingAchList {
	return v.value
}

func (v *NullableOutgoingAchList) Set(val *OutgoingAchList) {
	v.value = val
	v.isSet = true
}

func (v NullableOutgoingAchList) IsSet() bool {
	return v.isSet
}

func (v *NullableOutgoingAchList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutgoingAchList(val *OutgoingAchList) *NullableOutgoingAchList {
	return &NullableOutgoingAchList{value: val, isSet: true}
}

func (v NullableOutgoingAchList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutgoingAchList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


