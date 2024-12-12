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

// checks if the RiskData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskData{}

// RiskData struct for RiskData
type RiskData struct {
	// Client IP
	ClientIp *string `json:"client_ip,omitempty"`
}

// NewRiskData instantiates a new RiskData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskData() *RiskData {
	this := RiskData{}
	return &this
}

// NewRiskDataWithDefaults instantiates a new RiskData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskDataWithDefaults() *RiskData {
	this := RiskData{}
	return &this
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *RiskData) GetClientIp() string {
	if o == nil || IsNil(o.ClientIp) {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskData) GetClientIpOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIp) {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *RiskData) HasClientIp() bool {
	if o != nil && !IsNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *RiskData) SetClientIp(v string) {
	o.ClientIp = &v
}

func (o RiskData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientIp) {
		toSerialize["client_ip"] = o.ClientIp
	}
	return toSerialize, nil
}

type NullableRiskData struct {
	value *RiskData
	isSet bool
}

func (v NullableRiskData) Get() *RiskData {
	return v.value
}

func (v *NullableRiskData) Set(val *RiskData) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskData) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskData(val *RiskData) *NullableRiskData {
	return &NullableRiskData{value: val, isSet: true}
}

func (v NullableRiskData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
