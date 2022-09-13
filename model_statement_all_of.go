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

// StatementAllOf struct for StatementAllOf
type StatementAllOf struct {
	SavingsSummary *SavingsSummary `json:"savings_summary,omitempty"`
}

// NewStatementAllOf instantiates a new StatementAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementAllOf() *StatementAllOf {
	this := StatementAllOf{}
	return &this
}

// NewStatementAllOfWithDefaults instantiates a new StatementAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementAllOfWithDefaults() *StatementAllOf {
	this := StatementAllOf{}
	return &this
}

// GetSavingsSummary returns the SavingsSummary field value if set, zero value otherwise.
func (o *StatementAllOf) GetSavingsSummary() SavingsSummary {
	if o == nil || o.SavingsSummary == nil {
		var ret SavingsSummary
		return ret
	}
	return *o.SavingsSummary
}

// GetSavingsSummaryOk returns a tuple with the SavingsSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementAllOf) GetSavingsSummaryOk() (*SavingsSummary, bool) {
	if o == nil || o.SavingsSummary == nil {
		return nil, false
	}
	return o.SavingsSummary, true
}

// HasSavingsSummary returns a boolean if a field has been set.
func (o *StatementAllOf) HasSavingsSummary() bool {
	if o != nil && o.SavingsSummary != nil {
		return true
	}

	return false
}

// SetSavingsSummary gets a reference to the given SavingsSummary and assigns it to the SavingsSummary field.
func (o *StatementAllOf) SetSavingsSummary(v SavingsSummary) {
	o.SavingsSummary = &v
}

func (o StatementAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SavingsSummary != nil {
		toSerialize["savings_summary"] = o.SavingsSummary
	}
	return json.Marshal(toSerialize)
}

type NullableStatementAllOf struct {
	value *StatementAllOf
	isSet bool
}

func (v NullableStatementAllOf) Get() *StatementAllOf {
	return v.value
}

func (v *NullableStatementAllOf) Set(val *StatementAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementAllOf(val *StatementAllOf) *NullableStatementAllOf {
	return &NullableStatementAllOf{value: val, isSet: true}
}

func (v NullableStatementAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


