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

// CardProgramResponseListAllOf struct for CardProgramResponseListAllOf
type CardProgramResponseListAllOf struct {
	// Array of Card Programs
	Programs []CardProgramResponse `json:"programs"`
}

// NewCardProgramResponseListAllOf instantiates a new CardProgramResponseListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProgramResponseListAllOf(programs []CardProgramResponse) *CardProgramResponseListAllOf {
	this := CardProgramResponseListAllOf{}
	this.Programs = programs
	return &this
}

// NewCardProgramResponseListAllOfWithDefaults instantiates a new CardProgramResponseListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProgramResponseListAllOfWithDefaults() *CardProgramResponseListAllOf {
	this := CardProgramResponseListAllOf{}
	return &this
}

// GetPrograms returns the Programs field value
func (o *CardProgramResponseListAllOf) GetPrograms() []CardProgramResponse {
	if o == nil {
		var ret []CardProgramResponse
		return ret
	}

	return o.Programs
}

// GetProgramsOk returns a tuple with the Programs field value
// and a boolean to check if the value has been set.
func (o *CardProgramResponseListAllOf) GetProgramsOk() ([]CardProgramResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Programs, true
}

// SetPrograms sets field value
func (o *CardProgramResponseListAllOf) SetPrograms(v []CardProgramResponse) {
	o.Programs = v
}

func (o CardProgramResponseListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["programs"] = o.Programs
	}
	return json.Marshal(toSerialize)
}

type NullableCardProgramResponseListAllOf struct {
	value *CardProgramResponseListAllOf
	isSet bool
}

func (v NullableCardProgramResponseListAllOf) Get() *CardProgramResponseListAllOf {
	return v.value
}

func (v *NullableCardProgramResponseListAllOf) Set(val *CardProgramResponseListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProgramResponseListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProgramResponseListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProgramResponseListAllOf(val *CardProgramResponseListAllOf) *NullableCardProgramResponseListAllOf {
	return &NullableCardProgramResponseListAllOf{value: val, isSet: true}
}

func (v NullableCardProgramResponseListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProgramResponseListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
