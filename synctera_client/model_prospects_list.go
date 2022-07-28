/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ProspectsList struct for ProspectsList
type ProspectsList struct {
	// Array of Prospects
	Prospects []Prospect1 `json:"prospects"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewProspectsList instantiates a new ProspectsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProspectsList(prospects []Prospect1) *ProspectsList {
	this := ProspectsList{}
	this.Prospects = prospects
	return &this
}

// NewProspectsListWithDefaults instantiates a new ProspectsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProspectsListWithDefaults() *ProspectsList {
	this := ProspectsList{}
	return &this
}

// GetProspects returns the Prospects field value
func (o *ProspectsList) GetProspects() []Prospect1 {
	if o == nil {
		var ret []Prospect1
		return ret
	}

	return o.Prospects
}

// GetProspectsOk returns a tuple with the Prospects field value
// and a boolean to check if the value has been set.
func (o *ProspectsList) GetProspectsOk() ([]Prospect1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prospects, true
}

// SetProspects sets field value
func (o *ProspectsList) SetProspects(v []Prospect1) {
	o.Prospects = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ProspectsList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProspectsList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ProspectsList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ProspectsList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ProspectsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["prospects"] = o.Prospects
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableProspectsList struct {
	value *ProspectsList
	isSet bool
}

func (v NullableProspectsList) Get() *ProspectsList {
	return v.value
}

func (v *NullableProspectsList) Set(val *ProspectsList) {
	v.value = val
	v.isSet = true
}

func (v NullableProspectsList) IsSet() bool {
	return v.isSet
}

func (v *NullableProspectsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProspectsList(val *ProspectsList) *NullableProspectsList {
	return &NullableProspectsList{value: val, isSet: true}
}

func (v NullableProspectsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProspectsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


