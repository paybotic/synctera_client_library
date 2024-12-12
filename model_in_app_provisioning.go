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

// checks if the InAppProvisioning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppProvisioning{}

// InAppProvisioning struct for InAppProvisioning
type InAppProvisioning struct {
	AddressVerification *DigitalWalletTokenAddressVerification `json:"address_verification,omitempty"`
	Enabled             *bool                                  `json:"enabled,omitempty"`
}

// NewInAppProvisioning instantiates a new InAppProvisioning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppProvisioning() *InAppProvisioning {
	this := InAppProvisioning{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewInAppProvisioningWithDefaults instantiates a new InAppProvisioning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppProvisioningWithDefaults() *InAppProvisioning {
	this := InAppProvisioning{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetAddressVerification returns the AddressVerification field value if set, zero value otherwise.
func (o *InAppProvisioning) GetAddressVerification() DigitalWalletTokenAddressVerification {
	if o == nil || IsNil(o.AddressVerification) {
		var ret DigitalWalletTokenAddressVerification
		return ret
	}
	return *o.AddressVerification
}

// GetAddressVerificationOk returns a tuple with the AddressVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppProvisioning) GetAddressVerificationOk() (*DigitalWalletTokenAddressVerification, bool) {
	if o == nil || IsNil(o.AddressVerification) {
		return nil, false
	}
	return o.AddressVerification, true
}

// HasAddressVerification returns a boolean if a field has been set.
func (o *InAppProvisioning) HasAddressVerification() bool {
	if o != nil && !IsNil(o.AddressVerification) {
		return true
	}

	return false
}

// SetAddressVerification gets a reference to the given DigitalWalletTokenAddressVerification and assigns it to the AddressVerification field.
func (o *InAppProvisioning) SetAddressVerification(v DigitalWalletTokenAddressVerification) {
	o.AddressVerification = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InAppProvisioning) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppProvisioning) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InAppProvisioning) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InAppProvisioning) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InAppProvisioning) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppProvisioning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressVerification) {
		toSerialize["address_verification"] = o.AddressVerification
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableInAppProvisioning struct {
	value *InAppProvisioning
	isSet bool
}

func (v NullableInAppProvisioning) Get() *InAppProvisioning {
	return v.value
}

func (v *NullableInAppProvisioning) Set(val *InAppProvisioning) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppProvisioning) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppProvisioning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppProvisioning(val *InAppProvisioning) *NullableInAppProvisioning {
	return &NullableInAppProvisioning{value: val, isSet: true}
}

func (v NullableInAppProvisioning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppProvisioning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
