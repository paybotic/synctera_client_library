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

// VerificationList struct for VerificationList
type VerificationList struct {
	// Array of verification results.
	Verifications []Verification `json:"verifications"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewVerificationList instantiates a new VerificationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationList(verifications []Verification) *VerificationList {
	this := VerificationList{}
	this.Verifications = verifications
	return &this
}

// NewVerificationListWithDefaults instantiates a new VerificationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationListWithDefaults() *VerificationList {
	this := VerificationList{}
	return &this
}

// GetVerifications returns the Verifications field value
func (o *VerificationList) GetVerifications() []Verification {
	if o == nil {
		var ret []Verification
		return ret
	}

	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value
// and a boolean to check if the value has been set.
func (o *VerificationList) GetVerificationsOk() ([]Verification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verifications, true
}

// SetVerifications sets field value
func (o *VerificationList) SetVerifications(v []Verification) {
	o.Verifications = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *VerificationList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *VerificationList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *VerificationList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o VerificationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["verifications"] = o.Verifications
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableVerificationList struct {
	value *VerificationList
	isSet bool
}

func (v NullableVerificationList) Get() *VerificationList {
	return v.value
}

func (v *NullableVerificationList) Set(val *VerificationList) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationList) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationList(val *VerificationList) *NullableVerificationList {
	return &NullableVerificationList{value: val, isSet: true}
}

func (v NullableVerificationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


