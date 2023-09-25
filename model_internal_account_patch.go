/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InternalAccountPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalAccountPatch{}

// InternalAccountPatch struct for InternalAccountPatch
type InternalAccountPatch struct {
	AccountType *InternalAccountType `json:"account_type,omitempty"`
	// A user provided description for the current account
	Description *string `json:"description,omitempty"`
	Purpose *InternalAccountPurpose `json:"purpose,omitempty"`
}

// NewInternalAccountPatch instantiates a new InternalAccountPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalAccountPatch() *InternalAccountPatch {
	this := InternalAccountPatch{}
	return &this
}

// NewInternalAccountPatchWithDefaults instantiates a new InternalAccountPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalAccountPatchWithDefaults() *InternalAccountPatch {
	this := InternalAccountPatch{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *InternalAccountPatch) GetAccountType() InternalAccountType {
	if o == nil || IsNil(o.AccountType) {
		var ret InternalAccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccountPatch) GetAccountTypeOk() (*InternalAccountType, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *InternalAccountPatch) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given InternalAccountType and assigns it to the AccountType field.
func (o *InternalAccountPatch) SetAccountType(v InternalAccountType) {
	o.AccountType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InternalAccountPatch) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccountPatch) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InternalAccountPatch) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InternalAccountPatch) SetDescription(v string) {
	o.Description = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *InternalAccountPatch) GetPurpose() InternalAccountPurpose {
	if o == nil || IsNil(o.Purpose) {
		var ret InternalAccountPurpose
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccountPatch) GetPurposeOk() (*InternalAccountPurpose, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *InternalAccountPatch) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given InternalAccountPurpose and assigns it to the Purpose field.
func (o *InternalAccountPatch) SetPurpose(v InternalAccountPurpose) {
	o.Purpose = &v
}

func (o InternalAccountPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalAccountPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	return toSerialize, nil
}

type NullableInternalAccountPatch struct {
	value *InternalAccountPatch
	isSet bool
}

func (v NullableInternalAccountPatch) Get() *InternalAccountPatch {
	return v.value
}

func (v *NullableInternalAccountPatch) Set(val *InternalAccountPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalAccountPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalAccountPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalAccountPatch(val *InternalAccountPatch) *NullableInternalAccountPatch {
	return &NullableInternalAccountPatch{value: val, isSet: true}
}

func (v NullableInternalAccountPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalAccountPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


