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

// checks if the TransferReversalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferReversalRequest{}

// TransferReversalRequest Reversal for a transfer with type PULL
type TransferReversalRequest struct {
	// Amount of the refund in cents (Amount can be up to the original amount)
	Amount int32 `json:"amount"`
	// ISO 4217  Alpha-3 currency code
	Currency             string `json:"currency"`
	AdditionalProperties map[string]interface{}
}

type _TransferReversalRequest TransferReversalRequest

// NewTransferReversalRequest instantiates a new TransferReversalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferReversalRequest(amount int32, currency string) *TransferReversalRequest {
	this := TransferReversalRequest{}
	this.Amount = amount
	this.Currency = currency
	return &this
}

// NewTransferReversalRequestWithDefaults instantiates a new TransferReversalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferReversalRequestWithDefaults() *TransferReversalRequest {
	this := TransferReversalRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TransferReversalRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferReversalRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferReversalRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *TransferReversalRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransferReversalRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransferReversalRequest) SetCurrency(v string) {
	o.Currency = v
}

func (o TransferReversalRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferReversalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransferReversalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
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

	varTransferReversalRequest := _TransferReversalRequest{}

	err = json.Unmarshal(data, &varTransferReversalRequest)

	if err != nil {
		return err
	}

	*o = TransferReversalRequest(varTransferReversalRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferReversalRequest struct {
	value *TransferReversalRequest
	isSet bool
}

func (v NullableTransferReversalRequest) Get() *TransferReversalRequest {
	return v.value
}

func (v *NullableTransferReversalRequest) Set(val *TransferReversalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferReversalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferReversalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferReversalRequest(val *TransferReversalRequest) *NullableTransferReversalRequest {
	return &NullableTransferReversalRequest{value: val, isSet: true}
}

func (v NullableTransferReversalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferReversalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
