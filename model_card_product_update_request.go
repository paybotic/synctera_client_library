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

// CardProductUpdateRequest struct for CardProductUpdateRequest
type CardProductUpdateRequest struct {
	DigitalWalletTokenization *DigitalWalletTokenization `json:"digital_wallet_tokenization,omitempty"`
	// Allow issuing cards on this product without requiring KYC (customer and account statuses must still be valid, and the customer must still be associated to the account). If set to true on a virtual card product, activation of newly issued cards on that product will no longer be created in an activated state automatically. With this setting enabled, cards will be issued in an unactivated state, and activation of the cards will be blocked until the customer has passed KYC.
	IssueWithoutKyc *bool `json:"issue_without_kyc,omitempty"`
}

// NewCardProductUpdateRequest instantiates a new CardProductUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProductUpdateRequest() *CardProductUpdateRequest {
	this := CardProductUpdateRequest{}
	return &this
}

// NewCardProductUpdateRequestWithDefaults instantiates a new CardProductUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProductUpdateRequestWithDefaults() *CardProductUpdateRequest {
	this := CardProductUpdateRequest{}
	return &this
}

// GetDigitalWalletTokenization returns the DigitalWalletTokenization field value if set, zero value otherwise.
func (o *CardProductUpdateRequest) GetDigitalWalletTokenization() DigitalWalletTokenization {
	if o == nil || o.DigitalWalletTokenization == nil {
		var ret DigitalWalletTokenization
		return ret
	}
	return *o.DigitalWalletTokenization
}

// GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductUpdateRequest) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool) {
	if o == nil || o.DigitalWalletTokenization == nil {
		return nil, false
	}
	return o.DigitalWalletTokenization, true
}

// HasDigitalWalletTokenization returns a boolean if a field has been set.
func (o *CardProductUpdateRequest) HasDigitalWalletTokenization() bool {
	if o != nil && o.DigitalWalletTokenization != nil {
		return true
	}

	return false
}

// SetDigitalWalletTokenization gets a reference to the given DigitalWalletTokenization and assigns it to the DigitalWalletTokenization field.
func (o *CardProductUpdateRequest) SetDigitalWalletTokenization(v DigitalWalletTokenization) {
	o.DigitalWalletTokenization = &v
}

// GetIssueWithoutKyc returns the IssueWithoutKyc field value if set, zero value otherwise.
func (o *CardProductUpdateRequest) GetIssueWithoutKyc() bool {
	if o == nil || o.IssueWithoutKyc == nil {
		var ret bool
		return ret
	}
	return *o.IssueWithoutKyc
}

// GetIssueWithoutKycOk returns a tuple with the IssueWithoutKyc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductUpdateRequest) GetIssueWithoutKycOk() (*bool, bool) {
	if o == nil || o.IssueWithoutKyc == nil {
		return nil, false
	}
	return o.IssueWithoutKyc, true
}

// HasIssueWithoutKyc returns a boolean if a field has been set.
func (o *CardProductUpdateRequest) HasIssueWithoutKyc() bool {
	if o != nil && o.IssueWithoutKyc != nil {
		return true
	}

	return false
}

// SetIssueWithoutKyc gets a reference to the given bool and assigns it to the IssueWithoutKyc field.
func (o *CardProductUpdateRequest) SetIssueWithoutKyc(v bool) {
	o.IssueWithoutKyc = &v
}

func (o CardProductUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DigitalWalletTokenization != nil {
		toSerialize["digital_wallet_tokenization"] = o.DigitalWalletTokenization
	}
	if o.IssueWithoutKyc != nil {
		toSerialize["issue_without_kyc"] = o.IssueWithoutKyc
	}
	return json.Marshal(toSerialize)
}

type NullableCardProductUpdateRequest struct {
	value *CardProductUpdateRequest
	isSet bool
}

func (v NullableCardProductUpdateRequest) Get() *CardProductUpdateRequest {
	return v.value
}

func (v *NullableCardProductUpdateRequest) Set(val *CardProductUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProductUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProductUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProductUpdateRequest(val *CardProductUpdateRequest) *NullableCardProductUpdateRequest {
	return &NullableCardProductUpdateRequest{value: val, isSet: true}
}

func (v NullableCardProductUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProductUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
