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

// checks if the AdditionalOwnerData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalOwnerData{}

// AdditionalOwnerData Contains additional information about the relationship.
type AdditionalOwnerData struct {
	// Percentage ownership of the related business.
	PercentOwnership float64 `json:"percent_ownership"`
}

// NewAdditionalOwnerData instantiates a new AdditionalOwnerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalOwnerData(percentOwnership float64) *AdditionalOwnerData {
	this := AdditionalOwnerData{}
	this.PercentOwnership = percentOwnership
	return &this
}

// NewAdditionalOwnerDataWithDefaults instantiates a new AdditionalOwnerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalOwnerDataWithDefaults() *AdditionalOwnerData {
	this := AdditionalOwnerData{}
	return &this
}

// GetPercentOwnership returns the PercentOwnership field value
func (o *AdditionalOwnerData) GetPercentOwnership() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.PercentOwnership
}

// GetPercentOwnershipOk returns a tuple with the PercentOwnership field value
// and a boolean to check if the value has been set.
func (o *AdditionalOwnerData) GetPercentOwnershipOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentOwnership, true
}

// SetPercentOwnership sets field value
func (o *AdditionalOwnerData) SetPercentOwnership(v float64) {
	o.PercentOwnership = v
}

func (o AdditionalOwnerData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalOwnerData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["percent_ownership"] = o.PercentOwnership
	return toSerialize, nil
}

type NullableAdditionalOwnerData struct {
	value *AdditionalOwnerData
	isSet bool
}

func (v NullableAdditionalOwnerData) Get() *AdditionalOwnerData {
	return v.value
}

func (v *NullableAdditionalOwnerData) Set(val *AdditionalOwnerData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalOwnerData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalOwnerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalOwnerData(val *AdditionalOwnerData) *NullableAdditionalOwnerData {
	return &NullableAdditionalOwnerData{value: val, isSet: true}
}

func (v NullableAdditionalOwnerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalOwnerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


