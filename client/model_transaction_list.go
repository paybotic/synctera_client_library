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

// checks if the TransactionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionList{}

// TransactionList struct for TransactionList
type TransactionList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of statement transactions
	Transactions         []Transaction `json:"transactions"`
	AdditionalProperties map[string]interface{}
}

type _TransactionList TransactionList

// NewTransactionList instantiates a new TransactionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionList(transactions []Transaction) *TransactionList {
	this := TransactionList{}
	this.Transactions = transactions
	return &this
}

// NewTransactionListWithDefaults instantiates a new TransactionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionListWithDefaults() *TransactionList {
	this := TransactionList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *TransactionList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *TransactionList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *TransactionList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetTransactions returns the Transactions field value
func (o *TransactionList) GetTransactions() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *TransactionList) GetTransactionsOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *TransactionList) SetTransactions(v []Transaction) {
	o.Transactions = v
}

func (o TransactionList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["transactions"] = o.Transactions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransactionList) UnmarshalJSON(data []byte) (err error) {
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

	varTransactionList := _TransactionList{}

	err = json.Unmarshal(data, &varTransactionList)

	if err != nil {
		return err
	}

	*o = TransactionList(varTransactionList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionList struct {
	value *TransactionList
	isSet bool
}

func (v NullableTransactionList) Get() *TransactionList {
	return v.value
}

func (v *NullableTransactionList) Set(val *TransactionList) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionList) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionList(val *TransactionList) *NullableTransactionList {
	return &NullableTransactionList{value: val, isSet: true}
}

func (v NullableTransactionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
