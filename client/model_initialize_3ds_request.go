/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Initialize3dsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Initialize3dsRequest{}

// Initialize3dsRequest Initialization for an External Card Transfer 3-D Secure Authentication request
type Initialize3dsRequest struct {
	// Amount in cents of the External Card Transfer to be authenticated
	Amount int32 `json:"amount"`
	// ISO 4217  Alpha-3 currency code
	Currency string `json:"currency"`
	// The ID of the External Card for which the 3DS Authentication will be performed
	ExternalCardId string `json:"external_card_id"`
}

type _Initialize3dsRequest Initialize3dsRequest

// NewInitialize3dsRequest instantiates a new Initialize3dsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialize3dsRequest(amount int32, currency string, externalCardId string) *Initialize3dsRequest {
	this := Initialize3dsRequest{}
	this.Amount = amount
	this.Currency = currency
	this.ExternalCardId = externalCardId
	return &this
}

// NewInitialize3dsRequestWithDefaults instantiates a new Initialize3dsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialize3dsRequestWithDefaults() *Initialize3dsRequest {
	this := Initialize3dsRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Initialize3dsRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Initialize3dsRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Initialize3dsRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *Initialize3dsRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Initialize3dsRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Initialize3dsRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetExternalCardId returns the ExternalCardId field value
func (o *Initialize3dsRequest) GetExternalCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalCardId
}

// GetExternalCardIdOk returns a tuple with the ExternalCardId field value
// and a boolean to check if the value has been set.
func (o *Initialize3dsRequest) GetExternalCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCardId, true
}

// SetExternalCardId sets field value
func (o *Initialize3dsRequest) SetExternalCardId(v string) {
	o.ExternalCardId = v
}

func (o Initialize3dsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Initialize3dsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	toSerialize["external_card_id"] = o.ExternalCardId
	return toSerialize, nil
}

func (o *Initialize3dsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"external_card_id",
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

	varInitialize3dsRequest := _Initialize3dsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInitialize3dsRequest)

	if err != nil {
		return err
	}

	*o = Initialize3dsRequest(varInitialize3dsRequest)

	return err
}

type NullableInitialize3dsRequest struct {
	value *Initialize3dsRequest
	isSet bool
}

func (v NullableInitialize3dsRequest) Get() *Initialize3dsRequest {
	return v.value
}

func (v *NullableInitialize3dsRequest) Set(val *Initialize3dsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInitialize3dsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInitialize3dsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitialize3dsRequest(val *Initialize3dsRequest) *NullableInitialize3dsRequest {
	return &NullableInitialize3dsRequest{value: val, isSet: true}
}

func (v NullableInitialize3dsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitialize3dsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
