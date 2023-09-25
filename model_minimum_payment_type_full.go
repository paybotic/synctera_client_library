/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MinimumPaymentTypeFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinimumPaymentTypeFull{}

// MinimumPaymentTypeFull struct for MinimumPaymentTypeFull
type MinimumPaymentTypeFull struct {
	Type MinimumPaymentType `json:"type"`
}

// NewMinimumPaymentTypeFull instantiates a new MinimumPaymentTypeFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimumPaymentTypeFull(type_ MinimumPaymentType) *MinimumPaymentTypeFull {
	this := MinimumPaymentTypeFull{}
	this.Type = type_
	return &this
}

// NewMinimumPaymentTypeFullWithDefaults instantiates a new MinimumPaymentTypeFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimumPaymentTypeFullWithDefaults() *MinimumPaymentTypeFull {
	this := MinimumPaymentTypeFull{}
	return &this
}

// GetType returns the Type field value
func (o *MinimumPaymentTypeFull) GetType() MinimumPaymentType {
	if o == nil {
		var ret MinimumPaymentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MinimumPaymentTypeFull) GetTypeOk() (*MinimumPaymentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MinimumPaymentTypeFull) SetType(v MinimumPaymentType) {
	o.Type = v
}

func (o MinimumPaymentTypeFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinimumPaymentTypeFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMinimumPaymentTypeFull struct {
	value *MinimumPaymentTypeFull
	isSet bool
}

func (v NullableMinimumPaymentTypeFull) Get() *MinimumPaymentTypeFull {
	return v.value
}

func (v *NullableMinimumPaymentTypeFull) Set(val *MinimumPaymentTypeFull) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimumPaymentTypeFull) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimumPaymentTypeFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimumPaymentTypeFull(val *MinimumPaymentTypeFull) *NullableMinimumPaymentTypeFull {
	return &NullableMinimumPaymentTypeFull{value: val, isSet: true}
}

func (v NullableMinimumPaymentTypeFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimumPaymentTypeFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


