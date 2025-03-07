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

// checks if the Balance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Balance{}

// Balance struct for Balance
type Balance struct {
	// balance in ISO 4217 minor currency units. Unit in cents.
	Balance              int64       `json:"balance"`
	Type                 BalanceType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _Balance Balance

// NewBalance instantiates a new Balance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalance(balance int64, type_ BalanceType) *Balance {
	this := Balance{}
	this.Balance = balance
	this.Type = type_
	return &this
}

// NewBalanceWithDefaults instantiates a new Balance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceWithDefaults() *Balance {
	this := Balance{}
	return &this
}

// GetBalance returns the Balance field value
func (o *Balance) GetBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *Balance) GetBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *Balance) SetBalance(v int64) {
	o.Balance = v
}

// GetType returns the Type field value
func (o *Balance) GetType() BalanceType {
	if o == nil {
		var ret BalanceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Balance) GetTypeOk() (*BalanceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Balance) SetType(v BalanceType) {
	o.Type = v
}

func (o Balance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Balance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance"] = o.Balance
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Balance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"balance",
		"type",
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

	varBalance := _Balance{}

	err = json.Unmarshal(data, &varBalance)

	if err != nil {
		return err
	}

	*o = Balance(varBalance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "balance")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBalance struct {
	value *Balance
	isSet bool
}

func (v NullableBalance) Get() *Balance {
	return v.value
}

func (v *NullableBalance) Set(val *Balance) {
	v.value = val
	v.isSet = true
}

func (v NullableBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalance(val *Balance) *NullableBalance {
	return &NullableBalance{value: val, isSet: true}
}

func (v NullableBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
