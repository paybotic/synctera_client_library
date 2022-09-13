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

// CustomerVerificationResultList struct for CustomerVerificationResultList
type CustomerVerificationResultList struct {
	// Array of verification results.
	Verifications []CustomerVerificationResult `json:"verifications"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewCustomerVerificationResultList instantiates a new CustomerVerificationResultList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerVerificationResultList(verifications []CustomerVerificationResult) *CustomerVerificationResultList {
	this := CustomerVerificationResultList{}
	this.Verifications = verifications
	return &this
}

// NewCustomerVerificationResultListWithDefaults instantiates a new CustomerVerificationResultList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerVerificationResultListWithDefaults() *CustomerVerificationResultList {
	this := CustomerVerificationResultList{}
	return &this
}

// GetVerifications returns the Verifications field value
func (o *CustomerVerificationResultList) GetVerifications() []CustomerVerificationResult {
	if o == nil {
		var ret []CustomerVerificationResult
		return ret
	}

	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value
// and a boolean to check if the value has been set.
func (o *CustomerVerificationResultList) GetVerificationsOk() ([]CustomerVerificationResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verifications, true
}

// SetVerifications sets field value
func (o *CustomerVerificationResultList) SetVerifications(v []CustomerVerificationResult) {
	o.Verifications = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CustomerVerificationResultList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerVerificationResultList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CustomerVerificationResultList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CustomerVerificationResultList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o CustomerVerificationResultList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["verifications"] = o.Verifications
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerVerificationResultList struct {
	value *CustomerVerificationResultList
	isSet bool
}

func (v NullableCustomerVerificationResultList) Get() *CustomerVerificationResultList {
	return v.value
}

func (v *NullableCustomerVerificationResultList) Set(val *CustomerVerificationResultList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerVerificationResultList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerVerificationResultList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerVerificationResultList(val *CustomerVerificationResultList) *NullableCustomerVerificationResultList {
	return &NullableCustomerVerificationResultList{value: val, isSet: true}
}

func (v NullableCustomerVerificationResultList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerVerificationResultList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


