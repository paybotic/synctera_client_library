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

// WaitlistsListAllOf struct for WaitlistsListAllOf
type WaitlistsListAllOf struct {
	// Array of Waitlists
	Waitlists []Waitlist `json:"waitlists"`
}

// NewWaitlistsListAllOf instantiates a new WaitlistsListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaitlistsListAllOf(waitlists []Waitlist) *WaitlistsListAllOf {
	this := WaitlistsListAllOf{}
	this.Waitlists = waitlists
	return &this
}

// NewWaitlistsListAllOfWithDefaults instantiates a new WaitlistsListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaitlistsListAllOfWithDefaults() *WaitlistsListAllOf {
	this := WaitlistsListAllOf{}
	return &this
}

// GetWaitlists returns the Waitlists field value
func (o *WaitlistsListAllOf) GetWaitlists() []Waitlist {
	if o == nil {
		var ret []Waitlist
		return ret
	}

	return o.Waitlists
}

// GetWaitlistsOk returns a tuple with the Waitlists field value
// and a boolean to check if the value has been set.
func (o *WaitlistsListAllOf) GetWaitlistsOk() ([]Waitlist, bool) {
	if o == nil {
		return nil, false
	}
	return o.Waitlists, true
}

// SetWaitlists sets field value
func (o *WaitlistsListAllOf) SetWaitlists(v []Waitlist) {
	o.Waitlists = v
}

func (o WaitlistsListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["waitlists"] = o.Waitlists
	}
	return json.Marshal(toSerialize)
}

type NullableWaitlistsListAllOf struct {
	value *WaitlistsListAllOf
	isSet bool
}

func (v NullableWaitlistsListAllOf) Get() *WaitlistsListAllOf {
	return v.value
}

func (v *NullableWaitlistsListAllOf) Set(val *WaitlistsListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWaitlistsListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWaitlistsListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaitlistsListAllOf(val *WaitlistsListAllOf) *NullableWaitlistsListAllOf {
	return &NullableWaitlistsListAllOf{value: val, isSet: true}
}

func (v NullableWaitlistsListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaitlistsListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


