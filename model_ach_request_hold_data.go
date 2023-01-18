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

// AchRequestHoldData struct for AchRequestHoldData
type AchRequestHoldData struct {
	Amount int32 `json:"amount"`
	Duration int32 `json:"duration"`
}

// NewAchRequestHoldData instantiates a new AchRequestHoldData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchRequestHoldData(amount int32, duration int32) *AchRequestHoldData {
	this := AchRequestHoldData{}
	this.Amount = amount
	this.Duration = duration
	return &this
}

// NewAchRequestHoldDataWithDefaults instantiates a new AchRequestHoldData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchRequestHoldDataWithDefaults() *AchRequestHoldData {
	this := AchRequestHoldData{}
	return &this
}

// GetAmount returns the Amount field value
func (o *AchRequestHoldData) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AchRequestHoldData) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AchRequestHoldData) SetAmount(v int32) {
	o.Amount = v
}

// GetDuration returns the Duration field value
func (o *AchRequestHoldData) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *AchRequestHoldData) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *AchRequestHoldData) SetDuration(v int32) {
	o.Duration = v
}

func (o AchRequestHoldData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	return json.Marshal(toSerialize)
}

type NullableAchRequestHoldData struct {
	value *AchRequestHoldData
	isSet bool
}

func (v NullableAchRequestHoldData) Get() *AchRequestHoldData {
	return v.value
}

func (v *NullableAchRequestHoldData) Set(val *AchRequestHoldData) {
	v.value = val
	v.isSet = true
}

func (v NullableAchRequestHoldData) IsSet() bool {
	return v.isSet
}

func (v *NullableAchRequestHoldData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchRequestHoldData(val *AchRequestHoldData) *NullableAchRequestHoldData {
	return &NullableAchRequestHoldData{value: val, isSet: true}
}

func (v NullableAchRequestHoldData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchRequestHoldData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


