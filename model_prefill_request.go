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

// PrefillRequest struct for PrefillRequest
type PrefillRequest struct {
	// If true, the person's SSN was successfully populated.
	SsnFilled *bool `json:"ssn_filled,omitempty"`
	// Last four digits of person's Social Security number (SSN).
	SsnLast4 *string `json:"ssn_last4,omitempty"`
}

// NewPrefillRequest instantiates a new PrefillRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrefillRequest() *PrefillRequest {
	this := PrefillRequest{}
	return &this
}

// NewPrefillRequestWithDefaults instantiates a new PrefillRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrefillRequestWithDefaults() *PrefillRequest {
	this := PrefillRequest{}
	return &this
}

// GetSsnFilled returns the SsnFilled field value if set, zero value otherwise.
func (o *PrefillRequest) GetSsnFilled() bool {
	if o == nil || o.SsnFilled == nil {
		var ret bool
		return ret
	}
	return *o.SsnFilled
}

// GetSsnFilledOk returns a tuple with the SsnFilled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefillRequest) GetSsnFilledOk() (*bool, bool) {
	if o == nil || o.SsnFilled == nil {
		return nil, false
	}
	return o.SsnFilled, true
}

// HasSsnFilled returns a boolean if a field has been set.
func (o *PrefillRequest) HasSsnFilled() bool {
	if o != nil && o.SsnFilled != nil {
		return true
	}

	return false
}

// SetSsnFilled gets a reference to the given bool and assigns it to the SsnFilled field.
func (o *PrefillRequest) SetSsnFilled(v bool) {
	o.SsnFilled = &v
}

// GetSsnLast4 returns the SsnLast4 field value if set, zero value otherwise.
func (o *PrefillRequest) GetSsnLast4() string {
	if o == nil || o.SsnLast4 == nil {
		var ret string
		return ret
	}
	return *o.SsnLast4
}

// GetSsnLast4Ok returns a tuple with the SsnLast4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrefillRequest) GetSsnLast4Ok() (*string, bool) {
	if o == nil || o.SsnLast4 == nil {
		return nil, false
	}
	return o.SsnLast4, true
}

// HasSsnLast4 returns a boolean if a field has been set.
func (o *PrefillRequest) HasSsnLast4() bool {
	if o != nil && o.SsnLast4 != nil {
		return true
	}

	return false
}

// SetSsnLast4 gets a reference to the given string and assigns it to the SsnLast4 field.
func (o *PrefillRequest) SetSsnLast4(v string) {
	o.SsnLast4 = &v
}

func (o PrefillRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SsnFilled != nil {
		toSerialize["ssn_filled"] = o.SsnFilled
	}
	if o.SsnLast4 != nil {
		toSerialize["ssn_last4"] = o.SsnLast4
	}
	return json.Marshal(toSerialize)
}

type NullablePrefillRequest struct {
	value *PrefillRequest
	isSet bool
}

func (v NullablePrefillRequest) Get() *PrefillRequest {
	return v.value
}

func (v *NullablePrefillRequest) Set(val *PrefillRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePrefillRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePrefillRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrefillRequest(val *PrefillRequest) *NullablePrefillRequest {
	return &NullablePrefillRequest{value: val, isSet: true}
}

func (v NullablePrefillRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrefillRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
