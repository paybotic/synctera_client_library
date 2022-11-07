/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SpendingLimitWithTime Limit over a specific time period.
type SpendingLimitWithTime struct {
	// Maximum amount allowed within the time range. Unit in cents.
	Amount *int64 `json:"amount,omitempty"`
	// Maximum number of transactions allowed within the time range
	Transactions *int64 `json:"transactions,omitempty"`
}

// NewSpendingLimitWithTime instantiates a new SpendingLimitWithTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingLimitWithTime() *SpendingLimitWithTime {
	this := SpendingLimitWithTime{}
	return &this
}

// NewSpendingLimitWithTimeWithDefaults instantiates a new SpendingLimitWithTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingLimitWithTimeWithDefaults() *SpendingLimitWithTime {
	this := SpendingLimitWithTime{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SpendingLimitWithTime) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimitWithTime) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SpendingLimitWithTime) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *SpendingLimitWithTime) SetAmount(v int64) {
	o.Amount = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *SpendingLimitWithTime) GetTransactions() int64 {
	if o == nil || o.Transactions == nil {
		var ret int64
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimitWithTime) GetTransactionsOk() (*int64, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *SpendingLimitWithTime) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given int64 and assigns it to the Transactions field.
func (o *SpendingLimitWithTime) SetTransactions(v int64) {
	o.Transactions = &v
}

func (o SpendingLimitWithTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableSpendingLimitWithTime struct {
	value *SpendingLimitWithTime
	isSet bool
}

func (v NullableSpendingLimitWithTime) Get() *SpendingLimitWithTime {
	return v.value
}

func (v *NullableSpendingLimitWithTime) Set(val *SpendingLimitWithTime) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingLimitWithTime) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingLimitWithTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingLimitWithTime(val *SpendingLimitWithTime) *NullableSpendingLimitWithTime {
	return &NullableSpendingLimitWithTime{value: val, isSet: true}
}

func (v NullableSpendingLimitWithTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingLimitWithTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


