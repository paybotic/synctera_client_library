/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CardListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardListResponse{}

// CardListResponse struct for CardListResponse
type CardListResponse struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of Cards
	Cards []CardResponse `json:"cards"`
}

// NewCardListResponse instantiates a new CardListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardListResponse(cards []CardResponse) *CardListResponse {
	this := CardListResponse{}
	this.Cards = cards
	return &this
}

// NewCardListResponseWithDefaults instantiates a new CardListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardListResponseWithDefaults() *CardListResponse {
	this := CardListResponse{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CardListResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CardListResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CardListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetCards returns the Cards field value
func (o *CardListResponse) GetCards() []CardResponse {
	if o == nil {
		var ret []CardResponse
		return ret
	}

	return o.Cards
}

// GetCardsOk returns a tuple with the Cards field value
// and a boolean to check if the value has been set.
func (o *CardListResponse) GetCardsOk() ([]CardResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cards, true
}

// SetCards sets field value
func (o *CardListResponse) SetCards(v []CardResponse) {
	o.Cards = v
}

func (o CardListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["cards"] = o.Cards
	return toSerialize, nil
}

type NullableCardListResponse struct {
	value *CardListResponse
	isSet bool
}

func (v NullableCardListResponse) Get() *CardListResponse {
	return v.value
}

func (v *NullableCardListResponse) Set(val *CardListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCardListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCardListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardListResponse(val *CardListResponse) *NullableCardListResponse {
	return &NullableCardListResponse{value: val, isSet: true}
}

func (v NullableCardListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


