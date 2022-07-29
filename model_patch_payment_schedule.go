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

// PatchPaymentSchedule Patch request for payment schedule
type PatchPaymentSchedule struct {
	// Target payment schedule status
	Status *string `json:"status,omitempty"`
}

// NewPatchPaymentSchedule instantiates a new PatchPaymentSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPaymentSchedule() *PatchPaymentSchedule {
	this := PatchPaymentSchedule{}
	return &this
}

// NewPatchPaymentScheduleWithDefaults instantiates a new PatchPaymentSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPaymentScheduleWithDefaults() *PatchPaymentSchedule {
	this := PatchPaymentSchedule{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchPaymentSchedule) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPaymentSchedule) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchPaymentSchedule) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchPaymentSchedule) SetStatus(v string) {
	o.Status = &v
}

func (o PatchPaymentSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullablePatchPaymentSchedule struct {
	value *PatchPaymentSchedule
	isSet bool
}

func (v NullablePatchPaymentSchedule) Get() *PatchPaymentSchedule {
	return v.value
}

func (v *NullablePatchPaymentSchedule) Set(val *PatchPaymentSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPaymentSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPaymentSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPaymentSchedule(val *PatchPaymentSchedule) *NullablePatchPaymentSchedule {
	return &NullablePatchPaymentSchedule{value: val, isSet: true}
}

func (v NullablePatchPaymentSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPaymentSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


