/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the WireTransactionSimulationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireTransactionSimulationRequest{}

// WireTransactionSimulationRequest Simulate receiving a Wire transaction
type WireTransactionSimulationRequest struct {
	// Number of the receiving account
	AccountNumber string `json:"account_number"`
	// Amount to transfer in cents (e.g. 100 = $1).
	Amount int32 `json:"amount"`
	// Network to use for the Wire transfer
	Network *string `json:"network,omitempty"`
}

type _WireTransactionSimulationRequest WireTransactionSimulationRequest

// NewWireTransactionSimulationRequest instantiates a new WireTransactionSimulationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireTransactionSimulationRequest(accountNumber string, amount int32) *WireTransactionSimulationRequest {
	this := WireTransactionSimulationRequest{}
	this.AccountNumber = accountNumber
	this.Amount = amount
	var network string = "FEDWIRE"
	this.Network = &network
	return &this
}

// NewWireTransactionSimulationRequestWithDefaults instantiates a new WireTransactionSimulationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireTransactionSimulationRequestWithDefaults() *WireTransactionSimulationRequest {
	this := WireTransactionSimulationRequest{}
	var network string = "FEDWIRE"
	this.Network = &network
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *WireTransactionSimulationRequest) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *WireTransactionSimulationRequest) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *WireTransactionSimulationRequest) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAmount returns the Amount field value
func (o *WireTransactionSimulationRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *WireTransactionSimulationRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *WireTransactionSimulationRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *WireTransactionSimulationRequest) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireTransactionSimulationRequest) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *WireTransactionSimulationRequest) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *WireTransactionSimulationRequest) SetNetwork(v string) {
	o.Network = &v
}

func (o WireTransactionSimulationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WireTransactionSimulationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_number"] = o.AccountNumber
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return toSerialize, nil
}

func (o *WireTransactionSimulationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_number",
		"amount",
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

	varWireTransactionSimulationRequest := _WireTransactionSimulationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWireTransactionSimulationRequest)

	if err != nil {
		return err
	}

	*o = WireTransactionSimulationRequest(varWireTransactionSimulationRequest)

	return err
}

type NullableWireTransactionSimulationRequest struct {
	value *WireTransactionSimulationRequest
	isSet bool
}

func (v NullableWireTransactionSimulationRequest) Get() *WireTransactionSimulationRequest {
	return v.value
}

func (v *NullableWireTransactionSimulationRequest) Set(val *WireTransactionSimulationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWireTransactionSimulationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWireTransactionSimulationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireTransactionSimulationRequest(val *WireTransactionSimulationRequest) *NullableWireTransactionSimulationRequest {
	return &NullableWireTransactionSimulationRequest{value: val, isSet: true}
}

func (v NullableWireTransactionSimulationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireTransactionSimulationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}