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

// CardProductListResponse struct for CardProductListResponse
type CardProductListResponse struct {
	// Array of Card Products
	CardProducts []CardProductResponse `json:"card_products"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewCardProductListResponse instantiates a new CardProductListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProductListResponse(cardProducts []CardProductResponse) *CardProductListResponse {
	this := CardProductListResponse{}
	this.CardProducts = cardProducts
	return &this
}

// NewCardProductListResponseWithDefaults instantiates a new CardProductListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProductListResponseWithDefaults() *CardProductListResponse {
	this := CardProductListResponse{}
	return &this
}

// GetCardProducts returns the CardProducts field value
func (o *CardProductListResponse) GetCardProducts() []CardProductResponse {
	if o == nil {
		var ret []CardProductResponse
		return ret
	}

	return o.CardProducts
}

// GetCardProductsOk returns a tuple with the CardProducts field value
// and a boolean to check if the value has been set.
func (o *CardProductListResponse) GetCardProductsOk() ([]CardProductResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardProducts, true
}

// SetCardProducts sets field value
func (o *CardProductListResponse) SetCardProducts(v []CardProductResponse) {
	o.CardProducts = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CardProductListResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CardProductListResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CardProductListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o CardProductListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["card_products"] = o.CardProducts
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableCardProductListResponse struct {
	value *CardProductListResponse
	isSet bool
}

func (v NullableCardProductListResponse) Get() *CardProductListResponse {
	return v.value
}

func (v *NullableCardProductListResponse) Set(val *CardProductListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProductListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProductListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProductListResponse(val *CardProductListResponse) *NullableCardProductListResponse {
	return &NullableCardProductListResponse{value: val, isSet: true}
}

func (v NullableCardProductListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProductListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


