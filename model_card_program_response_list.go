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

// CardProgramResponseList struct for CardProgramResponseList
type CardProgramResponseList struct {
	// Array of Card Programs
	Programs []CardProgramResponse `json:"programs"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewCardProgramResponseList instantiates a new CardProgramResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProgramResponseList(programs []CardProgramResponse) *CardProgramResponseList {
	this := CardProgramResponseList{}
	this.Programs = programs
	return &this
}

// NewCardProgramResponseListWithDefaults instantiates a new CardProgramResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProgramResponseListWithDefaults() *CardProgramResponseList {
	this := CardProgramResponseList{}
	return &this
}

// GetPrograms returns the Programs field value
func (o *CardProgramResponseList) GetPrograms() []CardProgramResponse {
	if o == nil {
		var ret []CardProgramResponse
		return ret
	}

	return o.Programs
}

// GetProgramsOk returns a tuple with the Programs field value
// and a boolean to check if the value has been set.
func (o *CardProgramResponseList) GetProgramsOk() ([]CardProgramResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Programs, true
}

// SetPrograms sets field value
func (o *CardProgramResponseList) SetPrograms(v []CardProgramResponse) {
	o.Programs = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CardProgramResponseList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramResponseList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CardProgramResponseList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CardProgramResponseList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o CardProgramResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["programs"] = o.Programs
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableCardProgramResponseList struct {
	value *CardProgramResponseList
	isSet bool
}

func (v NullableCardProgramResponseList) Get() *CardProgramResponseList {
	return v.value
}

func (v *NullableCardProgramResponseList) Set(val *CardProgramResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProgramResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProgramResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProgramResponseList(val *CardProgramResponseList) *NullableCardProgramResponseList {
	return &NullableCardProgramResponseList{value: val, isSet: true}
}

func (v NullableCardProgramResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProgramResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


