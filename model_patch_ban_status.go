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

// PatchBanStatus struct for PatchBanStatus
type PatchBanStatus struct {
	BanStatus BanStatus `json:"ban_status"`
}

// NewPatchBanStatus instantiates a new PatchBanStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchBanStatus(banStatus BanStatus) *PatchBanStatus {
	this := PatchBanStatus{}
	this.BanStatus = banStatus
	return &this
}

// NewPatchBanStatusWithDefaults instantiates a new PatchBanStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchBanStatusWithDefaults() *PatchBanStatus {
	this := PatchBanStatus{}
	return &this
}

// GetBanStatus returns the BanStatus field value
func (o *PatchBanStatus) GetBanStatus() BanStatus {
	if o == nil {
		var ret BanStatus
		return ret
	}

	return o.BanStatus
}

// GetBanStatusOk returns a tuple with the BanStatus field value
// and a boolean to check if the value has been set.
func (o *PatchBanStatus) GetBanStatusOk() (*BanStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BanStatus, true
}

// SetBanStatus sets field value
func (o *PatchBanStatus) SetBanStatus(v BanStatus) {
	o.BanStatus = v
}

func (o PatchBanStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ban_status"] = o.BanStatus
	}
	return json.Marshal(toSerialize)
}

type NullablePatchBanStatus struct {
	value *PatchBanStatus
	isSet bool
}

func (v NullablePatchBanStatus) Get() *PatchBanStatus {
	return v.value
}

func (v *NullablePatchBanStatus) Set(val *PatchBanStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBanStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBanStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBanStatus(val *PatchBanStatus) *NullablePatchBanStatus {
	return &NullablePatchBanStatus{value: val, isSet: true}
}

func (v NullablePatchBanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBanStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
