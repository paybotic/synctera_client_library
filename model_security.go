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

// Security struct for Security
type Security struct {
	// ID of linked backing account for use as a security, e.g. for use in a Smart Charge Card offering. Must be of type CHECKING or SAVING. 
	LinkedAccountId string `json:"linked_account_id"`
}

// NewSecurity instantiates a new Security object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurity(linkedAccountId string) *Security {
	this := Security{}
	this.LinkedAccountId = linkedAccountId
	return &this
}

// NewSecurityWithDefaults instantiates a new Security object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityWithDefaults() *Security {
	this := Security{}
	return &this
}

// GetLinkedAccountId returns the LinkedAccountId field value
func (o *Security) GetLinkedAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkedAccountId
}

// GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field value
// and a boolean to check if the value has been set.
func (o *Security) GetLinkedAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkedAccountId, true
}

// SetLinkedAccountId sets field value
func (o *Security) SetLinkedAccountId(v string) {
	o.LinkedAccountId = v
}

func (o Security) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["linked_account_id"] = o.LinkedAccountId
	}
	return json.Marshal(toSerialize)
}

type NullableSecurity struct {
	value *Security
	isSet bool
}

func (v NullableSecurity) Get() *Security {
	return v.value
}

func (v *NullableSecurity) Set(val *Security) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurity(val *Security) *NullableSecurity {
	return &NullableSecurity{value: val, isSet: true}
}

func (v NullableSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


