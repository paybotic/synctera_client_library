/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SimulateCardFulfillment struct for SimulateCardFulfillment
type SimulateCardFulfillment struct {
	CardFulfillmentStatus CardFulfillmentStatus `json:"card_fulfillment_status"`
}

// NewSimulateCardFulfillment instantiates a new SimulateCardFulfillment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulateCardFulfillment(cardFulfillmentStatus CardFulfillmentStatus) *SimulateCardFulfillment {
	this := SimulateCardFulfillment{}
	this.CardFulfillmentStatus = cardFulfillmentStatus
	return &this
}

// NewSimulateCardFulfillmentWithDefaults instantiates a new SimulateCardFulfillment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulateCardFulfillmentWithDefaults() *SimulateCardFulfillment {
	this := SimulateCardFulfillment{}
	return &this
}

// GetCardFulfillmentStatus returns the CardFulfillmentStatus field value
func (o *SimulateCardFulfillment) GetCardFulfillmentStatus() CardFulfillmentStatus {
	if o == nil {
		var ret CardFulfillmentStatus
		return ret
	}

	return o.CardFulfillmentStatus
}

// GetCardFulfillmentStatusOk returns a tuple with the CardFulfillmentStatus field value
// and a boolean to check if the value has been set.
func (o *SimulateCardFulfillment) GetCardFulfillmentStatusOk() (*CardFulfillmentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardFulfillmentStatus, true
}

// SetCardFulfillmentStatus sets field value
func (o *SimulateCardFulfillment) SetCardFulfillmentStatus(v CardFulfillmentStatus) {
	o.CardFulfillmentStatus = v
}

func (o SimulateCardFulfillment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["card_fulfillment_status"] = o.CardFulfillmentStatus
	}
	return json.Marshal(toSerialize)
}

type NullableSimulateCardFulfillment struct {
	value *SimulateCardFulfillment
	isSet bool
}

func (v NullableSimulateCardFulfillment) Get() *SimulateCardFulfillment {
	return v.value
}

func (v *NullableSimulateCardFulfillment) Set(val *SimulateCardFulfillment) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulateCardFulfillment) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulateCardFulfillment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulateCardFulfillment(val *SimulateCardFulfillment) *NullableSimulateCardFulfillment {
	return &NullableSimulateCardFulfillment{value: val, isSet: true}
}

func (v NullableSimulateCardFulfillment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulateCardFulfillment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


