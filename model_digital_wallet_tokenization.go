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

// DigitalWalletTokenization struct for DigitalWalletTokenization
type DigitalWalletTokenization struct {
	ProvisioningControls *ProvisioningControls `json:"provisioning_controls,omitempty"`
}

// NewDigitalWalletTokenization instantiates a new DigitalWalletTokenization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletTokenization() *DigitalWalletTokenization {
	this := DigitalWalletTokenization{}
	return &this
}

// NewDigitalWalletTokenizationWithDefaults instantiates a new DigitalWalletTokenization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletTokenizationWithDefaults() *DigitalWalletTokenization {
	this := DigitalWalletTokenization{}
	return &this
}

// GetProvisioningControls returns the ProvisioningControls field value if set, zero value otherwise.
func (o *DigitalWalletTokenization) GetProvisioningControls() ProvisioningControls {
	if o == nil || o.ProvisioningControls == nil {
		var ret ProvisioningControls
		return ret
	}
	return *o.ProvisioningControls
}

// GetProvisioningControlsOk returns a tuple with the ProvisioningControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenization) GetProvisioningControlsOk() (*ProvisioningControls, bool) {
	if o == nil || o.ProvisioningControls == nil {
		return nil, false
	}
	return o.ProvisioningControls, true
}

// HasProvisioningControls returns a boolean if a field has been set.
func (o *DigitalWalletTokenization) HasProvisioningControls() bool {
	if o != nil && o.ProvisioningControls != nil {
		return true
	}

	return false
}

// SetProvisioningControls gets a reference to the given ProvisioningControls and assigns it to the ProvisioningControls field.
func (o *DigitalWalletTokenization) SetProvisioningControls(v ProvisioningControls) {
	o.ProvisioningControls = &v
}

func (o DigitalWalletTokenization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProvisioningControls != nil {
		toSerialize["provisioning_controls"] = o.ProvisioningControls
	}
	return json.Marshal(toSerialize)
}

type NullableDigitalWalletTokenization struct {
	value *DigitalWalletTokenization
	isSet bool
}

func (v NullableDigitalWalletTokenization) Get() *DigitalWalletTokenization {
	return v.value
}

func (v *NullableDigitalWalletTokenization) Set(val *DigitalWalletTokenization) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletTokenization) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletTokenization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletTokenization(val *DigitalWalletTokenization) *NullableDigitalWalletTokenization {
	return &NullableDigitalWalletTokenization{value: val, isSet: true}
}

func (v NullableDigitalWalletTokenization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletTokenization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


