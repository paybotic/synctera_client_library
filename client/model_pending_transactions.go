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

// checks if the PendingTransactions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingTransactions{}

// PendingTransactions struct for PendingTransactions
type PendingTransactions struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken NullableString `json:"next_page_token"`
	// List of pending transactions
	Result               []PendingTransaction `json:"result"`
	AdditionalProperties map[string]interface{}
}

type _PendingTransactions PendingTransactions

// NewPendingTransactions instantiates a new PendingTransactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingTransactions(nextPageToken NullableString, result []PendingTransaction) *PendingTransactions {
	this := PendingTransactions{}
	this.NextPageToken = nextPageToken
	this.Result = result
	return &this
}

// NewPendingTransactionsWithDefaults instantiates a new PendingTransactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingTransactionsWithDefaults() *PendingTransactions {
	this := PendingTransactions{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PendingTransactions) GetNextPageToken() string {
	if o == nil || o.NextPageToken.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextPageToken.Get()
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingTransactions) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageToken.Get(), o.NextPageToken.IsSet()
}

// SetNextPageToken sets field value
func (o *PendingTransactions) SetNextPageToken(v string) {
	o.NextPageToken.Set(&v)
}

// GetResult returns the Result field value
func (o *PendingTransactions) GetResult() []PendingTransaction {
	if o == nil {
		var ret []PendingTransaction
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PendingTransactions) GetResultOk() ([]PendingTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *PendingTransactions) SetResult(v []PendingTransaction) {
	o.Result = v
}

func (o PendingTransactions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingTransactions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["next_page_token"] = o.NextPageToken.Get()
	toSerialize["result"] = o.Result

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PendingTransactions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"next_page_token",
		"result",
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

	varPendingTransactions := _PendingTransactions{}

	err = json.Unmarshal(data, &varPendingTransactions)

	if err != nil {
		return err
	}

	*o = PendingTransactions(varPendingTransactions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePendingTransactions struct {
	value *PendingTransactions
	isSet bool
}

func (v NullablePendingTransactions) Get() *PendingTransactions {
	return v.value
}

func (v *NullablePendingTransactions) Set(val *PendingTransactions) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingTransactions) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingTransactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingTransactions(val *PendingTransactions) *NullablePendingTransactions {
	return &NullablePendingTransactions{value: val, isSet: true}
}

func (v NullablePendingTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingTransactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
