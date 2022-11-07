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

// CardListResponseAllOf struct for CardListResponseAllOf
type CardListResponseAllOf struct {
	// Array of Cards
	Cards []CardResponse `json:"cards"`
}

// NewCardListResponseAllOf instantiates a new CardListResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardListResponseAllOf(cards []CardResponse) *CardListResponseAllOf {
	this := CardListResponseAllOf{}
	this.Cards = cards
	return &this
}

// NewCardListResponseAllOfWithDefaults instantiates a new CardListResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardListResponseAllOfWithDefaults() *CardListResponseAllOf {
	this := CardListResponseAllOf{}
	return &this
}

// GetCards returns the Cards field value
func (o *CardListResponseAllOf) GetCards() []CardResponse {
	if o == nil {
		var ret []CardResponse
		return ret
	}

	return o.Cards
}

// GetCardsOk returns a tuple with the Cards field value
// and a boolean to check if the value has been set.
func (o *CardListResponseAllOf) GetCardsOk() ([]CardResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cards, true
}

// SetCards sets field value
func (o *CardListResponseAllOf) SetCards(v []CardResponse) {
	o.Cards = v
}

func (o CardListResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cards"] = o.Cards
	}
	return json.Marshal(toSerialize)
}

type NullableCardListResponseAllOf struct {
	value *CardListResponseAllOf
	isSet bool
}

func (v NullableCardListResponseAllOf) Get() *CardListResponseAllOf {
	return v.value
}

func (v *NullableCardListResponseAllOf) Set(val *CardListResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCardListResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCardListResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardListResponseAllOf(val *CardListResponseAllOf) *NullableCardListResponseAllOf {
	return &NullableCardListResponseAllOf{value: val, isSet: true}
}

func (v NullableCardListResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardListResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


