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

// AchInstruction struct for AchInstruction
type AchInstruction struct {
	Request OutgoingAchRequest `json:"request"`
	Type string `json:"type"`
}

// NewAchInstruction instantiates a new AchInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchInstruction(request OutgoingAchRequest, type_ string) *AchInstruction {
	this := AchInstruction{}
	this.Request = request
	this.Type = type_
	return &this
}

// NewAchInstructionWithDefaults instantiates a new AchInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchInstructionWithDefaults() *AchInstruction {
	this := AchInstruction{}
	return &this
}

// GetRequest returns the Request field value
func (o *AchInstruction) GetRequest() OutgoingAchRequest {
	if o == nil {
		var ret OutgoingAchRequest
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *AchInstruction) GetRequestOk() (*OutgoingAchRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *AchInstruction) SetRequest(v OutgoingAchRequest) {
	o.Request = v
}

// GetType returns the Type field value
func (o *AchInstruction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AchInstruction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AchInstruction) SetType(v string) {
	o.Type = v
}

func (o AchInstruction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request"] = o.Request
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAchInstruction struct {
	value *AchInstruction
	isSet bool
}

func (v NullableAchInstruction) Get() *AchInstruction {
	return v.value
}

func (v *NullableAchInstruction) Set(val *AchInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableAchInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableAchInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchInstruction(val *AchInstruction) *NullableAchInstruction {
	return &NullableAchInstruction{value: val, isSet: true}
}

func (v NullableAchInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


