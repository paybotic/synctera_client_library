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

// checks if the InternalTransferPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalTransferPatch{}

// InternalTransferPatch struct for InternalTransferPatch
type InternalTransferPatch struct {
	// The desired status of the internal transfer auth
	Status *string `json:"status,omitempty"`
}

// NewInternalTransferPatch instantiates a new InternalTransferPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalTransferPatch() *InternalTransferPatch {
	this := InternalTransferPatch{}
	return &this
}

// NewInternalTransferPatchWithDefaults instantiates a new InternalTransferPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalTransferPatchWithDefaults() *InternalTransferPatch {
	this := InternalTransferPatch{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalTransferPatch) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferPatch) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalTransferPatch) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalTransferPatch) SetStatus(v string) {
	o.Status = &v
}

func (o InternalTransferPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalTransferPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableInternalTransferPatch struct {
	value *InternalTransferPatch
	isSet bool
}

func (v NullableInternalTransferPatch) Get() *InternalTransferPatch {
	return v.value
}

func (v *NullableInternalTransferPatch) Set(val *InternalTransferPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalTransferPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalTransferPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalTransferPatch(val *InternalTransferPatch) *NullableInternalTransferPatch {
	return &NullableInternalTransferPatch{value: val, isSet: true}
}

func (v NullableInternalTransferPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalTransferPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
