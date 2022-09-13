/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TemplateList struct for TemplateList
type TemplateList struct {
	// Array of account templates
	AccountTemplates []AccountTemplateResponse `json:"account_templates"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewTemplateList instantiates a new TemplateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateList(accountTemplates []AccountTemplateResponse) *TemplateList {
	this := TemplateList{}
	this.AccountTemplates = accountTemplates
	return &this
}

// NewTemplateListWithDefaults instantiates a new TemplateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateListWithDefaults() *TemplateList {
	this := TemplateList{}
	return &this
}

// GetAccountTemplates returns the AccountTemplates field value
func (o *TemplateList) GetAccountTemplates() []AccountTemplateResponse {
	if o == nil {
		var ret []AccountTemplateResponse
		return ret
	}

	return o.AccountTemplates
}

// GetAccountTemplatesOk returns a tuple with the AccountTemplates field value
// and a boolean to check if the value has been set.
func (o *TemplateList) GetAccountTemplatesOk() ([]AccountTemplateResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountTemplates, true
}

// SetAccountTemplates sets field value
func (o *TemplateList) SetAccountTemplates(v []AccountTemplateResponse) {
	o.AccountTemplates = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *TemplateList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *TemplateList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *TemplateList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o TemplateList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_templates"] = o.AccountTemplates
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateList struct {
	value *TemplateList
	isSet bool
}

func (v NullableTemplateList) Get() *TemplateList {
	return v.value
}

func (v *NullableTemplateList) Set(val *TemplateList) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateList) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateList(val *TemplateList) *NullableTemplateList {
	return &NullableTemplateList{value: val, isSet: true}
}

func (v NullableTemplateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


