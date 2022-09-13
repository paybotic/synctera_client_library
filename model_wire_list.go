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

// WireList struct for WireList
type WireList struct {
	// Array of wires
	Wires []Wire `json:"wires"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewWireList instantiates a new WireList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireList(wires []Wire) *WireList {
	this := WireList{}
	this.Wires = wires
	return &this
}

// NewWireListWithDefaults instantiates a new WireList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireListWithDefaults() *WireList {
	this := WireList{}
	return &this
}

// GetWires returns the Wires field value
func (o *WireList) GetWires() []Wire {
	if o == nil {
		var ret []Wire
		return ret
	}

	return o.Wires
}

// GetWiresOk returns a tuple with the Wires field value
// and a boolean to check if the value has been set.
func (o *WireList) GetWiresOk() ([]Wire, bool) {
	if o == nil {
		return nil, false
	}
	return o.Wires, true
}

// SetWires sets field value
func (o *WireList) SetWires(v []Wire) {
	o.Wires = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *WireList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *WireList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *WireList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o WireList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wires"] = o.Wires
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableWireList struct {
	value *WireList
	isSet bool
}

func (v NullableWireList) Get() *WireList {
	return v.value
}

func (v *NullableWireList) Set(val *WireList) {
	v.value = val
	v.isSet = true
}

func (v NullableWireList) IsSet() bool {
	return v.isSet
}

func (v *NullableWireList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireList(val *WireList) *NullableWireList {
	return &NullableWireList{value: val, isSet: true}
}

func (v NullableWireList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


