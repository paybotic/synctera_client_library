/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DigitalWalletTokenAddressVerification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalWalletTokenAddressVerification{}

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
	if o == nil || IsNil(o.Validate) {
		var ret bool
		return ret
	}
	return *o.Validate
}

// GetValidateOk returns a tuple with the Validate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenAddressVerification) GetValidateOk() (*bool, bool) {
	if o == nil || IsNil(o.Validate) {
		return nil, false
	}
	return o.Validate, true
}

// HasValidate returns a boolean if a field has been set.
func (o *DigitalWalletTokenAddressVerification) HasValidate() bool {
	if o != nil && !IsNil(o.Validate) {
		return true
	}

	return false
}

// SetValidate gets a reference to the given bool and assigns it to the Validate field.
func (o *DigitalWalletTokenAddressVerification) SetValidate(v bool) {
	o.Validate = &v
}

func (o DigitalWalletTokenAddressVerification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalWalletTokenAddressVerification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Validate) {
		toSerialize["validate"] = o.Validate
	}
	return toSerialize, nil
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
