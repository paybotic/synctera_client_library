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

// SavingsSummary A summary of the accrued interest for the saving account in the current period
type SavingsSummary struct {
	Apy *float32 `json:"apy,omitempty"`
}

// NewSavingsSummary instantiates a new SavingsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavingsSummary() *SavingsSummary {
	this := SavingsSummary{}
	return &this
}

// NewSavingsSummaryWithDefaults instantiates a new SavingsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavingsSummaryWithDefaults() *SavingsSummary {
	this := SavingsSummary{}
	return &this
}

// GetApy returns the Apy field value if set, zero value otherwise.
func (o *SavingsSummary) GetApy() float32 {
	if o == nil || o.Apy == nil {
		var ret float32
		return ret
	}
	return *o.Apy
}

// GetApyOk returns a tuple with the Apy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavingsSummary) GetApyOk() (*float32, bool) {
	if o == nil || o.Apy == nil {
		return nil, false
	}
	return o.Apy, true
}

// HasApy returns a boolean if a field has been set.
func (o *SavingsSummary) HasApy() bool {
	if o != nil && o.Apy != nil {
		return true
	}

	return false
}

// SetApy gets a reference to the given float32 and assigns it to the Apy field.
func (o *SavingsSummary) SetApy(v float32) {
	o.Apy = &v
}

func (o SavingsSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apy != nil {
		toSerialize["apy"] = o.Apy
	}
	return json.Marshal(toSerialize)
}

type NullableSavingsSummary struct {
	value *SavingsSummary
	isSet bool
}

func (v NullableSavingsSummary) Get() *SavingsSummary {
	return v.value
}

func (v *NullableSavingsSummary) Set(val *SavingsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableSavingsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSavingsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavingsSummary(val *SavingsSummary) *NullableSavingsSummary {
	return &NullableSavingsSummary{value: val, isSet: true}
}

func (v NullableSavingsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavingsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


