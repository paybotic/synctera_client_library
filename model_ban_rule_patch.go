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

// BanRulePatch struct for BanRulePatch
type BanRulePatch struct {
	Status *BanRuleStatus `json:"status,omitempty"`
}

// NewBanRulePatch instantiates a new BanRulePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanRulePatch() *BanRulePatch {
	this := BanRulePatch{}
	return &this
}

// NewBanRulePatchWithDefaults instantiates a new BanRulePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanRulePatchWithDefaults() *BanRulePatch {
	this := BanRulePatch{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BanRulePatch) GetStatus() BanRuleStatus {
	if o == nil || o.Status == nil {
		var ret BanRuleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BanRulePatch) GetStatusOk() (*BanRuleStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BanRulePatch) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BanRuleStatus and assigns it to the Status field.
func (o *BanRulePatch) SetStatus(v BanRuleStatus) {
	o.Status = &v
}

func (o BanRulePatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBanRulePatch struct {
	value *BanRulePatch
	isSet bool
}

func (v NullableBanRulePatch) Get() *BanRulePatch {
	return v.value
}

func (v *NullableBanRulePatch) Set(val *BanRulePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableBanRulePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableBanRulePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanRulePatch(val *BanRulePatch) *NullableBanRulePatch {
	return &NullableBanRulePatch{value: val, isSet: true}
}

func (v NullableBanRulePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanRulePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


