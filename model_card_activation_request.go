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

// CardActivationRequest struct for CardActivationRequest
type CardActivationRequest struct {
	// An activation code provided with the card required to prove possession of the card
	ActivationCode string `json:"activation_code"`
	// The ID of the customer for which card is being activated
	CustomerId string `json:"customer_id"`
}

// NewCardActivationRequest instantiates a new CardActivationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardActivationRequest(activationCode string, customerId string) *CardActivationRequest {
	this := CardActivationRequest{}
	this.ActivationCode = activationCode
	this.CustomerId = customerId
	return &this
}

// NewCardActivationRequestWithDefaults instantiates a new CardActivationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardActivationRequestWithDefaults() *CardActivationRequest {
	this := CardActivationRequest{}
	return &this
}

// GetActivationCode returns the ActivationCode field value
func (o *CardActivationRequest) GetActivationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivationCode
}

// GetActivationCodeOk returns a tuple with the ActivationCode field value
// and a boolean to check if the value has been set.
func (o *CardActivationRequest) GetActivationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivationCode, true
}

// SetActivationCode sets field value
func (o *CardActivationRequest) SetActivationCode(v string) {
	o.ActivationCode = v
}

// GetCustomerId returns the CustomerId field value
func (o *CardActivationRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CardActivationRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CardActivationRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o CardActivationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["activation_code"] = o.ActivationCode
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	return json.Marshal(toSerialize)
}

type NullableCardActivationRequest struct {
	value *CardActivationRequest
	isSet bool
}

func (v NullableCardActivationRequest) Get() *CardActivationRequest {
	return v.value
}

func (v *NullableCardActivationRequest) Set(val *CardActivationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardActivationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardActivationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardActivationRequest(val *CardActivationRequest) *NullableCardActivationRequest {
	return &NullableCardActivationRequest{value: val, isSet: true}
}

func (v NullableCardActivationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardActivationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


