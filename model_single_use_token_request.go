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

// SingleUseTokenRequest struct for SingleUseTokenRequest
type SingleUseTokenRequest struct {
	// The ID of the account to which the token will be linked
	AccountId string `json:"account_id"`
	// The ID of the customer to whom the token will be issued
	CustomerId string `json:"customer_id"`
}

// NewSingleUseTokenRequest instantiates a new SingleUseTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleUseTokenRequest(accountId string, customerId string) *SingleUseTokenRequest {
	this := SingleUseTokenRequest{}
	this.AccountId = accountId
	this.CustomerId = customerId
	return &this
}

// NewSingleUseTokenRequestWithDefaults instantiates a new SingleUseTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleUseTokenRequestWithDefaults() *SingleUseTokenRequest {
	this := SingleUseTokenRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *SingleUseTokenRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SingleUseTokenRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *SingleUseTokenRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetCustomerId returns the CustomerId field value
func (o *SingleUseTokenRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *SingleUseTokenRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *SingleUseTokenRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o SingleUseTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	return json.Marshal(toSerialize)
}

type NullableSingleUseTokenRequest struct {
	value *SingleUseTokenRequest
	isSet bool
}

func (v NullableSingleUseTokenRequest) Get() *SingleUseTokenRequest {
	return v.value
}

func (v *NullableSingleUseTokenRequest) Set(val *SingleUseTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleUseTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleUseTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleUseTokenRequest(val *SingleUseTokenRequest) *NullableSingleUseTokenRequest {
	return &NullableSingleUseTokenRequest{value: val, isSet: true}
}

func (v NullableSingleUseTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleUseTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


