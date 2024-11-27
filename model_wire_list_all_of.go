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

// WireListAllOf struct for WireListAllOf
type WireListAllOf struct {
	// Array of wires
	Wires []Wire `json:"wires"`
}

// NewWireListAllOf instantiates a new WireListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireListAllOf(wires []Wire) *WireListAllOf {
	this := WireListAllOf{}
	this.Wires = wires
	return &this
}

// NewWireListAllOfWithDefaults instantiates a new WireListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireListAllOfWithDefaults() *WireListAllOf {
	this := WireListAllOf{}
	return &this
}

// GetWires returns the Wires field value
func (o *WireListAllOf) GetWires() []Wire {
	if o == nil {
		var ret []Wire
		return ret
	}

	return o.Wires
}

// GetWiresOk returns a tuple with the Wires field value
// and a boolean to check if the value has been set.
func (o *WireListAllOf) GetWiresOk() ([]Wire, bool) {
	if o == nil {
		return nil, false
	}
	return o.Wires, true
}

// SetWires sets field value
func (o *WireListAllOf) SetWires(v []Wire) {
	o.Wires = v
}

func (o WireListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wires"] = o.Wires
	}
	return json.Marshal(toSerialize)
}

type NullableWireListAllOf struct {
	value *WireListAllOf
	isSet bool
}

func (v NullableWireListAllOf) Get() *WireListAllOf {
	return v.value
}

func (v *NullableWireListAllOf) Set(val *WireListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWireListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWireListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireListAllOf(val *WireListAllOf) *NullableWireListAllOf {
	return &NullableWireListAllOf{value: val, isSet: true}
}

func (v NullableWireListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
