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

// MasterDisclosureListAllOf struct for MasterDisclosureListAllOf
type MasterDisclosureListAllOf struct {
	// Array of master disclosures.
	MasterDisclosures []MasterDisclosure `json:"master_disclosures"`
}

// NewMasterDisclosureListAllOf instantiates a new MasterDisclosureListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterDisclosureListAllOf(masterDisclosures []MasterDisclosure) *MasterDisclosureListAllOf {
	this := MasterDisclosureListAllOf{}
	this.MasterDisclosures = masterDisclosures
	return &this
}

// NewMasterDisclosureListAllOfWithDefaults instantiates a new MasterDisclosureListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterDisclosureListAllOfWithDefaults() *MasterDisclosureListAllOf {
	this := MasterDisclosureListAllOf{}
	return &this
}

// GetMasterDisclosures returns the MasterDisclosures field value
func (o *MasterDisclosureListAllOf) GetMasterDisclosures() []MasterDisclosure {
	if o == nil {
		var ret []MasterDisclosure
		return ret
	}

	return o.MasterDisclosures
}

// GetMasterDisclosuresOk returns a tuple with the MasterDisclosures field value
// and a boolean to check if the value has been set.
func (o *MasterDisclosureListAllOf) GetMasterDisclosuresOk() ([]MasterDisclosure, bool) {
	if o == nil {
		return nil, false
	}
	return o.MasterDisclosures, true
}

// SetMasterDisclosures sets field value
func (o *MasterDisclosureListAllOf) SetMasterDisclosures(v []MasterDisclosure) {
	o.MasterDisclosures = v
}

func (o MasterDisclosureListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["master_disclosures"] = o.MasterDisclosures
	}
	return json.Marshal(toSerialize)
}

type NullableMasterDisclosureListAllOf struct {
	value *MasterDisclosureListAllOf
	isSet bool
}

func (v NullableMasterDisclosureListAllOf) Get() *MasterDisclosureListAllOf {
	return v.value
}

func (v *NullableMasterDisclosureListAllOf) Set(val *MasterDisclosureListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterDisclosureListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterDisclosureListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterDisclosureListAllOf(val *MasterDisclosureListAllOf) *NullableMasterDisclosureListAllOf {
	return &NullableMasterDisclosureListAllOf{value: val, isSet: true}
}

func (v NullableMasterDisclosureListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterDisclosureListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
