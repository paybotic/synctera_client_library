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

// TemplateListAllOf struct for TemplateListAllOf
type TemplateListAllOf struct {
	// Array of account templates
	AccountTemplates []AccountTemplateResponse `json:"account_templates"`
}

// NewTemplateListAllOf instantiates a new TemplateListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateListAllOf(accountTemplates []AccountTemplateResponse) *TemplateListAllOf {
	this := TemplateListAllOf{}
	this.AccountTemplates = accountTemplates
	return &this
}

// NewTemplateListAllOfWithDefaults instantiates a new TemplateListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateListAllOfWithDefaults() *TemplateListAllOf {
	this := TemplateListAllOf{}
	return &this
}

// GetAccountTemplates returns the AccountTemplates field value
func (o *TemplateListAllOf) GetAccountTemplates() []AccountTemplateResponse {
	if o == nil {
		var ret []AccountTemplateResponse
		return ret
	}

	return o.AccountTemplates
}

// GetAccountTemplatesOk returns a tuple with the AccountTemplates field value
// and a boolean to check if the value has been set.
func (o *TemplateListAllOf) GetAccountTemplatesOk() ([]AccountTemplateResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountTemplates, true
}

// SetAccountTemplates sets field value
func (o *TemplateListAllOf) SetAccountTemplates(v []AccountTemplateResponse) {
	o.AccountTemplates = v
}

func (o TemplateListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_templates"] = o.AccountTemplates
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateListAllOf struct {
	value *TemplateListAllOf
	isSet bool
}

func (v NullableTemplateListAllOf) Get() *TemplateListAllOf {
	return v.value
}

func (v *NullableTemplateListAllOf) Set(val *TemplateListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateListAllOf(val *TemplateListAllOf) *NullableTemplateListAllOf {
	return &NullableTemplateListAllOf{value: val, isSet: true}
}

func (v NullableTemplateListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


