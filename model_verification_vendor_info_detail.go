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

// checks if the VerificationVendorInfoDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationVendorInfoDetail{}

// VerificationVendorInfoDetail struct for VerificationVendorInfoDetail
type VerificationVendorInfoDetail struct {
	// Vendor specific code.
	Code *string `json:"code,omitempty"`
	// Description of vendor specific code.
	Description *string `json:"description,omitempty"`
}

// NewVerificationVendorInfoDetail instantiates a new VerificationVendorInfoDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationVendorInfoDetail() *VerificationVendorInfoDetail {
	this := VerificationVendorInfoDetail{}
	return &this
}

// NewVerificationVendorInfoDetailWithDefaults instantiates a new VerificationVendorInfoDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationVendorInfoDetailWithDefaults() *VerificationVendorInfoDetail {
	this := VerificationVendorInfoDetail{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VerificationVendorInfoDetail) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationVendorInfoDetail) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VerificationVendorInfoDetail) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *VerificationVendorInfoDetail) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VerificationVendorInfoDetail) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationVendorInfoDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VerificationVendorInfoDetail) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VerificationVendorInfoDetail) SetDescription(v string) {
	o.Description = &v
}

func (o VerificationVendorInfoDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationVendorInfoDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableVerificationVendorInfoDetail struct {
	value *VerificationVendorInfoDetail
	isSet bool
}

func (v NullableVerificationVendorInfoDetail) Get() *VerificationVendorInfoDetail {
	return v.value
}

func (v *NullableVerificationVendorInfoDetail) Set(val *VerificationVendorInfoDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationVendorInfoDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationVendorInfoDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationVendorInfoDetail(val *VerificationVendorInfoDetail) *NullableVerificationVendorInfoDetail {
	return &NullableVerificationVendorInfoDetail{value: val, isSet: true}
}

func (v NullableVerificationVendorInfoDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationVendorInfoDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
