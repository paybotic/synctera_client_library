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

// SpendControlRollingWindowDaysAllOf struct for SpendControlRollingWindowDaysAllOf
type SpendControlRollingWindowDaysAllOf struct {
	// The number of days to define a rolling window for a spend control
	Days int32 `json:"days"`
}

// NewSpendControlRollingWindowDaysAllOf instantiates a new SpendControlRollingWindowDaysAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendControlRollingWindowDaysAllOf(days int32) *SpendControlRollingWindowDaysAllOf {
	this := SpendControlRollingWindowDaysAllOf{}
	this.Days = days
	return &this
}

// NewSpendControlRollingWindowDaysAllOfWithDefaults instantiates a new SpendControlRollingWindowDaysAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendControlRollingWindowDaysAllOfWithDefaults() *SpendControlRollingWindowDaysAllOf {
	this := SpendControlRollingWindowDaysAllOf{}
	return &this
}

// GetDays returns the Days field value
func (o *SpendControlRollingWindowDaysAllOf) GetDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *SpendControlRollingWindowDaysAllOf) GetDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Days, true
}

// SetDays sets field value
func (o *SpendControlRollingWindowDaysAllOf) SetDays(v int32) {
	o.Days = v
}

func (o SpendControlRollingWindowDaysAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["days"] = o.Days
	}
	return json.Marshal(toSerialize)
}

type NullableSpendControlRollingWindowDaysAllOf struct {
	value *SpendControlRollingWindowDaysAllOf
	isSet bool
}

func (v NullableSpendControlRollingWindowDaysAllOf) Get() *SpendControlRollingWindowDaysAllOf {
	return v.value
}

func (v *NullableSpendControlRollingWindowDaysAllOf) Set(val *SpendControlRollingWindowDaysAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendControlRollingWindowDaysAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendControlRollingWindowDaysAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendControlRollingWindowDaysAllOf(val *SpendControlRollingWindowDaysAllOf) *NullableSpendControlRollingWindowDaysAllOf {
	return &NullableSpendControlRollingWindowDaysAllOf{value: val, isSet: true}
}

func (v NullableSpendControlRollingWindowDaysAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendControlRollingWindowDaysAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
