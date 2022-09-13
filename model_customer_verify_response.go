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

// CustomerVerifyResponse struct for CustomerVerifyResponse
type CustomerVerifyResponse struct {
	KycStatus CustomerKycStatus `json:"kyc_status"`
	// Array of verification results.
	Verifications []CustomerVerificationResult `json:"verifications"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewCustomerVerifyResponse instantiates a new CustomerVerifyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerVerifyResponse(kycStatus CustomerKycStatus, verifications []CustomerVerificationResult) *CustomerVerifyResponse {
	this := CustomerVerifyResponse{}
	this.KycStatus = kycStatus
	this.Verifications = verifications
	return &this
}

// NewCustomerVerifyResponseWithDefaults instantiates a new CustomerVerifyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerVerifyResponseWithDefaults() *CustomerVerifyResponse {
	this := CustomerVerifyResponse{}
	return &this
}

// GetKycStatus returns the KycStatus field value
func (o *CustomerVerifyResponse) GetKycStatus() CustomerKycStatus {
	if o == nil {
		var ret CustomerKycStatus
		return ret
	}

	return o.KycStatus
}

// GetKycStatusOk returns a tuple with the KycStatus field value
// and a boolean to check if the value has been set.
func (o *CustomerVerifyResponse) GetKycStatusOk() (*CustomerKycStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KycStatus, true
}

// SetKycStatus sets field value
func (o *CustomerVerifyResponse) SetKycStatus(v CustomerKycStatus) {
	o.KycStatus = v
}

// GetVerifications returns the Verifications field value
func (o *CustomerVerifyResponse) GetVerifications() []CustomerVerificationResult {
	if o == nil {
		var ret []CustomerVerificationResult
		return ret
	}

	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value
// and a boolean to check if the value has been set.
func (o *CustomerVerifyResponse) GetVerificationsOk() ([]CustomerVerificationResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verifications, true
}

// SetVerifications sets field value
func (o *CustomerVerifyResponse) SetVerifications(v []CustomerVerificationResult) {
	o.Verifications = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CustomerVerifyResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerVerifyResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CustomerVerifyResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CustomerVerifyResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o CustomerVerifyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kyc_status"] = o.KycStatus
	}
	if true {
		toSerialize["verifications"] = o.Verifications
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerVerifyResponse struct {
	value *CustomerVerifyResponse
	isSet bool
}

func (v NullableCustomerVerifyResponse) Get() *CustomerVerifyResponse {
	return v.value
}

func (v *NullableCustomerVerifyResponse) Set(val *CustomerVerifyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerVerifyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerVerifyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerVerifyResponse(val *CustomerVerifyResponse) *NullableCustomerVerifyResponse {
	return &NullableCustomerVerifyResponse{value: val, isSet: true}
}

func (v NullableCustomerVerifyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerVerifyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


