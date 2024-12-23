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

// WaitlistsList struct for WaitlistsList
type WaitlistsList struct {
	// Array of Waitlists
	Waitlists []Waitlist `json:"waitlists"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewWaitlistsList instantiates a new WaitlistsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaitlistsList(waitlists []Waitlist) *WaitlistsList {
	this := WaitlistsList{}
	this.Waitlists = waitlists
	return &this
}

// NewWaitlistsListWithDefaults instantiates a new WaitlistsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaitlistsListWithDefaults() *WaitlistsList {
	this := WaitlistsList{}
	return &this
}

// GetWaitlists returns the Waitlists field value
func (o *WaitlistsList) GetWaitlists() []Waitlist {
	if o == nil {
		var ret []Waitlist
		return ret
	}

	return o.Waitlists
}

// GetWaitlistsOk returns a tuple with the Waitlists field value
// and a boolean to check if the value has been set.
func (o *WaitlistsList) GetWaitlistsOk() ([]Waitlist, bool) {
	if o == nil {
		return nil, false
	}
	return o.Waitlists, true
}

// SetWaitlists sets field value
func (o *WaitlistsList) SetWaitlists(v []Waitlist) {
	o.Waitlists = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *WaitlistsList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistsList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *WaitlistsList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *WaitlistsList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o WaitlistsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["waitlists"] = o.Waitlists
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableWaitlistsList struct {
	value *WaitlistsList
	isSet bool
}

func (v NullableWaitlistsList) Get() *WaitlistsList {
	return v.value
}

func (v *NullableWaitlistsList) Set(val *WaitlistsList) {
	v.value = val
	v.isSet = true
}

func (v NullableWaitlistsList) IsSet() bool {
	return v.isSet
}

func (v *NullableWaitlistsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaitlistsList(val *WaitlistsList) *NullableWaitlistsList {
	return &NullableWaitlistsList{value: val, isSet: true}
}

func (v NullableWaitlistsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaitlistsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
