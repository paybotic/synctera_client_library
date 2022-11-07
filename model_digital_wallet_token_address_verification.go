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

// DigitalWalletTokenAddressVerification struct for DigitalWalletTokenAddressVerification
type DigitalWalletTokenAddressVerification struct {
	Validate *bool `json:"validate,omitempty"`
}

// NewDigitalWalletTokenAddressVerification instantiates a new DigitalWalletTokenAddressVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletTokenAddressVerification() *DigitalWalletTokenAddressVerification {
	this := DigitalWalletTokenAddressVerification{}
	return &this
}

// NewDigitalWalletTokenAddressVerificationWithDefaults instantiates a new DigitalWalletTokenAddressVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletTokenAddressVerificationWithDefaults() *DigitalWalletTokenAddressVerification {
	this := DigitalWalletTokenAddressVerification{}
	return &this
}

// GetValidate returns the Validate field value if set, zero value otherwise.
func (o *DigitalWalletTokenAddressVerification) GetValidate() bool {
	if o == nil || o.Validate == nil {
		var ret bool
		return ret
	}
	return *o.Validate
}

// GetValidateOk returns a tuple with the Validate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenAddressVerification) GetValidateOk() (*bool, bool) {
	if o == nil || o.Validate == nil {
		return nil, false
	}
	return o.Validate, true
}

// HasValidate returns a boolean if a field has been set.
func (o *DigitalWalletTokenAddressVerification) HasValidate() bool {
	if o != nil && o.Validate != nil {
		return true
	}

	return false
}

// SetValidate gets a reference to the given bool and assigns it to the Validate field.
func (o *DigitalWalletTokenAddressVerification) SetValidate(v bool) {
	o.Validate = &v
}

func (o DigitalWalletTokenAddressVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Validate != nil {
		toSerialize["validate"] = o.Validate
	}
	return json.Marshal(toSerialize)
}

type NullableDigitalWalletTokenAddressVerification struct {
	value *DigitalWalletTokenAddressVerification
	isSet bool
}

func (v NullableDigitalWalletTokenAddressVerification) Get() *DigitalWalletTokenAddressVerification {
	return v.value
}

func (v *NullableDigitalWalletTokenAddressVerification) Set(val *DigitalWalletTokenAddressVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletTokenAddressVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletTokenAddressVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletTokenAddressVerification(val *DigitalWalletTokenAddressVerification) *NullableDigitalWalletTokenAddressVerification {
	return &NullableDigitalWalletTokenAddressVerification{value: val, isSet: true}
}

func (v NullableDigitalWalletTokenAddressVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletTokenAddressVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


