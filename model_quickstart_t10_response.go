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

// QuickstartT10Response T-10 cards config details
type QuickstartT10Response struct {
	// Array of card products created
	CardProducts []CardProductResponse `json:"card_products"`
	CardProgram  CardProgramResponse   `json:"card_program"`
}

// NewQuickstartT10Response instantiates a new QuickstartT10Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickstartT10Response(cardProducts []CardProductResponse, cardProgram CardProgramResponse) *QuickstartT10Response {
	this := QuickstartT10Response{}
	this.CardProducts = cardProducts
	this.CardProgram = cardProgram
	return &this
}

// NewQuickstartT10ResponseWithDefaults instantiates a new QuickstartT10Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickstartT10ResponseWithDefaults() *QuickstartT10Response {
	this := QuickstartT10Response{}
	return &this
}

// GetCardProducts returns the CardProducts field value
func (o *QuickstartT10Response) GetCardProducts() []CardProductResponse {
	if o == nil {
		var ret []CardProductResponse
		return ret
	}

	return o.CardProducts
}

// GetCardProductsOk returns a tuple with the CardProducts field value
// and a boolean to check if the value has been set.
func (o *QuickstartT10Response) GetCardProductsOk() ([]CardProductResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardProducts, true
}

// SetCardProducts sets field value
func (o *QuickstartT10Response) SetCardProducts(v []CardProductResponse) {
	o.CardProducts = v
}

// GetCardProgram returns the CardProgram field value
func (o *QuickstartT10Response) GetCardProgram() CardProgramResponse {
	if o == nil {
		var ret CardProgramResponse
		return ret
	}

	return o.CardProgram
}

// GetCardProgramOk returns a tuple with the CardProgram field value
// and a boolean to check if the value has been set.
func (o *QuickstartT10Response) GetCardProgramOk() (*CardProgramResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProgram, true
}

// SetCardProgram sets field value
func (o *QuickstartT10Response) SetCardProgram(v CardProgramResponse) {
	o.CardProgram = v
}

func (o QuickstartT10Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["card_products"] = o.CardProducts
	}
	if true {
		toSerialize["card_program"] = o.CardProgram
	}
	return json.Marshal(toSerialize)
}

type NullableQuickstartT10Response struct {
	value *QuickstartT10Response
	isSet bool
}

func (v NullableQuickstartT10Response) Get() *QuickstartT10Response {
	return v.value
}

func (v *NullableQuickstartT10Response) Set(val *QuickstartT10Response) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickstartT10Response) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickstartT10Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickstartT10Response(val *QuickstartT10Response) *NullableQuickstartT10Response {
	return &NullableQuickstartT10Response{value: val, isSet: true}
}

func (v NullableQuickstartT10Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickstartT10Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
