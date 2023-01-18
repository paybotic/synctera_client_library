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

// DepositPatchRequest Approve or reject deposits
type DepositPatchRequest struct {
	// The status of the deposit
	Status *string `json:"status,omitempty"`
}

// NewDepositPatchRequest instantiates a new DepositPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositPatchRequest() *DepositPatchRequest {
	this := DepositPatchRequest{}
	return &this
}

// NewDepositPatchRequestWithDefaults instantiates a new DepositPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositPatchRequestWithDefaults() *DepositPatchRequest {
	this := DepositPatchRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DepositPatchRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPatchRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DepositPatchRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DepositPatchRequest) SetStatus(v string) {
	o.Status = &v
}

func (o DepositPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDepositPatchRequest struct {
	value *DepositPatchRequest
	isSet bool
}

func (v NullableDepositPatchRequest) Get() *DepositPatchRequest {
	return v.value
}

func (v *NullableDepositPatchRequest) Set(val *DepositPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositPatchRequest(val *DepositPatchRequest) *NullableDepositPatchRequest {
	return &NullableDepositPatchRequest{value: val, isSet: true}
}

func (v NullableDepositPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


