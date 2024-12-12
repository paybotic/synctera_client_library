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

// checks if the ExternalAccountsTransactionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalAccountsTransactionList{}

// ExternalAccountsTransactionList struct for ExternalAccountsTransactionList
type ExternalAccountsTransactionList struct {
	// Array of transactions of a given external account
	Transactions []ExternalAccountTransaction `json:"transactions"`
}

type _ExternalAccountsTransactionList ExternalAccountsTransactionList

// NewExternalAccountsTransactionList instantiates a new ExternalAccountsTransactionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccountsTransactionList(transactions []ExternalAccountTransaction) *ExternalAccountsTransactionList {
	this := ExternalAccountsTransactionList{}
	this.Transactions = transactions
	return &this
}

// NewExternalAccountsTransactionListWithDefaults instantiates a new ExternalAccountsTransactionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountsTransactionListWithDefaults() *ExternalAccountsTransactionList {
	this := ExternalAccountsTransactionList{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *ExternalAccountsTransactionList) GetTransactions() []ExternalAccountTransaction {
	if o == nil {
		var ret []ExternalAccountTransaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountsTransactionList) GetTransactionsOk() ([]ExternalAccountTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *ExternalAccountsTransactionList) SetTransactions(v []ExternalAccountTransaction) {
	o.Transactions = v
}

func (o ExternalAccountsTransactionList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalAccountsTransactionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

func (o *ExternalAccountsTransactionList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transactions",
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

	varExternalAccountsTransactionList := _ExternalAccountsTransactionList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalAccountsTransactionList)

	if err != nil {
		return err
	}

	*o = ExternalAccountsTransactionList(varExternalAccountsTransactionList)

	return err
}

type NullableExternalAccountsTransactionList struct {
	value *ExternalAccountsTransactionList
	isSet bool
}

func (v NullableExternalAccountsTransactionList) Get() *ExternalAccountsTransactionList {
	return v.value
}

func (v *NullableExternalAccountsTransactionList) Set(val *ExternalAccountsTransactionList) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccountsTransactionList) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccountsTransactionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccountsTransactionList(val *ExternalAccountsTransactionList) *NullableExternalAccountsTransactionList {
	return &NullableExternalAccountsTransactionList{value: val, isSet: true}
}

func (v NullableExternalAccountsTransactionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccountsTransactionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
