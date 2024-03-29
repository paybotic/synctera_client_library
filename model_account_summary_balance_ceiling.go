/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AccountSummaryBalanceCeiling struct for AccountSummaryBalanceCeiling
type AccountSummaryBalanceCeiling struct {
	// Maximum balance in the account's currency. Unit in cents.
	Balance *int64 `json:"balance,omitempty"`
}

// NewAccountSummaryBalanceCeiling instantiates a new AccountSummaryBalanceCeiling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSummaryBalanceCeiling() *AccountSummaryBalanceCeiling {
	this := AccountSummaryBalanceCeiling{}
	return &this
}

// NewAccountSummaryBalanceCeilingWithDefaults instantiates a new AccountSummaryBalanceCeiling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSummaryBalanceCeilingWithDefaults() *AccountSummaryBalanceCeiling {
	this := AccountSummaryBalanceCeiling{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *AccountSummaryBalanceCeiling) GetBalance() int64 {
	if o == nil || o.Balance == nil {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummaryBalanceCeiling) GetBalanceOk() (*int64, bool) {
	if o == nil || o.Balance == nil {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *AccountSummaryBalanceCeiling) HasBalance() bool {
	if o != nil && o.Balance != nil {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *AccountSummaryBalanceCeiling) SetBalance(v int64) {
	o.Balance = &v
}

func (o AccountSummaryBalanceCeiling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Balance != nil {
		toSerialize["balance"] = o.Balance
	}
	return json.Marshal(toSerialize)
}

type NullableAccountSummaryBalanceCeiling struct {
	value *AccountSummaryBalanceCeiling
	isSet bool
}

func (v NullableAccountSummaryBalanceCeiling) Get() *AccountSummaryBalanceCeiling {
	return v.value
}

func (v *NullableAccountSummaryBalanceCeiling) Set(val *AccountSummaryBalanceCeiling) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSummaryBalanceCeiling) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSummaryBalanceCeiling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSummaryBalanceCeiling(val *AccountSummaryBalanceCeiling) *NullableAccountSummaryBalanceCeiling {
	return &NullableAccountSummaryBalanceCeiling{value: val, isSet: true}
}

func (v NullableAccountSummaryBalanceCeiling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSummaryBalanceCeiling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


