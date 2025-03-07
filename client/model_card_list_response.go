/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the CardListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardListResponse{}

// CardListResponse struct for CardListResponse
type CardListResponse struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of Cards
	Cards                []CardResponse `json:"cards"`
	AdditionalProperties map[string]interface{}
}

type _CardListResponse CardListResponse

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
	toSerialize, err := o.ToMap()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cards",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCardListResponse := _CardListResponse{}

	err = json.Unmarshal(data, &varCardListResponse)

	if err != nil {
		return err
	}

	*o = CardListResponse(varCardListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "cards")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
