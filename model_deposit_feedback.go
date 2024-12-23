/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DepositFeedback Approve or reject deposit transactions
type DepositFeedback struct {
	// The value of the feedback for the deposit transaction
	Feedback *string `json:"feedback,omitempty"`
}

// NewDepositFeedback instantiates a new DepositFeedback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositFeedback() *DepositFeedback {
	this := DepositFeedback{}
	return &this
}

// NewDepositFeedbackWithDefaults instantiates a new DepositFeedback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositFeedbackWithDefaults() *DepositFeedback {
	this := DepositFeedback{}
	return &this
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *DepositFeedback) GetFeedback() string {
	if o == nil || o.Feedback == nil {
		var ret string
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositFeedback) GetFeedbackOk() (*string, bool) {
	if o == nil || o.Feedback == nil {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *DepositFeedback) HasFeedback() bool {
	if o != nil && o.Feedback != nil {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given string and assigns it to the Feedback field.
func (o *DepositFeedback) SetFeedback(v string) {
	o.Feedback = &v
}

func (o DepositFeedback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Feedback != nil {
		toSerialize["feedback"] = o.Feedback
	}
	return json.Marshal(toSerialize)
}

type NullableDepositFeedback struct {
	value *DepositFeedback
	isSet bool
}

func (v NullableDepositFeedback) Get() *DepositFeedback {
	return v.value
}

func (v *NullableDepositFeedback) Set(val *DepositFeedback) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositFeedback) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositFeedback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositFeedback(val *DepositFeedback) *NullableDepositFeedback {
	return &NullableDepositFeedback{value: val, isSet: true}
}

func (v NullableDepositFeedback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositFeedback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
