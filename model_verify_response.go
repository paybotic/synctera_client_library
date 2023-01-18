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

// VerifyResponse struct for VerifyResponse
type VerifyResponse struct {
	VerificationStatus VerificationStatus `json:"verification_status"`
	// Array of verification results.
	Verifications []Verification `json:"verifications"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewVerifyResponse instantiates a new VerifyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyResponse(verificationStatus VerificationStatus, verifications []Verification) *VerifyResponse {
	this := VerifyResponse{}
	this.VerificationStatus = verificationStatus
	this.Verifications = verifications
	return &this
}

// NewVerifyResponseWithDefaults instantiates a new VerifyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyResponseWithDefaults() *VerifyResponse {
	this := VerifyResponse{}
	return &this
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *VerifyResponse) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *VerifyResponse) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

// GetVerifications returns the Verifications field value
func (o *VerifyResponse) GetVerifications() []Verification {
	if o == nil {
		var ret []Verification
		return ret
	}

	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetVerificationsOk() ([]Verification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verifications, true
}

// SetVerifications sets field value
func (o *VerifyResponse) SetVerifications(v []Verification) {
	o.Verifications = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *VerifyResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *VerifyResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *VerifyResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o VerifyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if true {
		toSerialize["verifications"] = o.Verifications
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyResponse struct {
	value *VerifyResponse
	isSet bool
}

func (v NullableVerifyResponse) Get() *VerifyResponse {
	return v.value
}

func (v *NullableVerifyResponse) Set(val *VerifyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyResponse(val *VerifyResponse) *NullableVerifyResponse {
	return &NullableVerifyResponse{value: val, isSet: true}
}

func (v NullableVerifyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


