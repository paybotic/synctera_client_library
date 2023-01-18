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

// BalanceFloor struct for BalanceFloor
type BalanceFloor struct {
	// Minimum balance in the account's currency. Unit in cents.
	Balance int64 `json:"balance"`
	// ID of linked backing account for just-in-time (JIT) funding of transactions to maintain the balance floor 
	LinkedAccountId *string `json:"linked_account_id,omitempty"`
	// ID of linked backing account for just-in-time (JIT) funding of transactions to maintain the balance floor This attribute is a deprecated alias for linked_account_id. 
	// Deprecated
	OverdraftAccountId *string `json:"overdraft_account_id,omitempty"`
}

// NewBalanceFloor instantiates a new BalanceFloor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceFloor(balance int64) *BalanceFloor {
	this := BalanceFloor{}
	this.Balance = balance
	return &this
}

// NewBalanceFloorWithDefaults instantiates a new BalanceFloor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceFloorWithDefaults() *BalanceFloor {
	this := BalanceFloor{}
	return &this
}

// GetBalance returns the Balance field value
func (o *BalanceFloor) GetBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *BalanceFloor) GetBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *BalanceFloor) SetBalance(v int64) {
	o.Balance = v
}

// GetLinkedAccountId returns the LinkedAccountId field value if set, zero value otherwise.
func (o *BalanceFloor) GetLinkedAccountId() string {
	if o == nil || o.LinkedAccountId == nil {
		var ret string
		return ret
	}
	return *o.LinkedAccountId
}

// GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceFloor) GetLinkedAccountIdOk() (*string, bool) {
	if o == nil || o.LinkedAccountId == nil {
		return nil, false
	}
	return o.LinkedAccountId, true
}

// HasLinkedAccountId returns a boolean if a field has been set.
func (o *BalanceFloor) HasLinkedAccountId() bool {
	if o != nil && o.LinkedAccountId != nil {
		return true
	}

	return false
}

// SetLinkedAccountId gets a reference to the given string and assigns it to the LinkedAccountId field.
func (o *BalanceFloor) SetLinkedAccountId(v string) {
	o.LinkedAccountId = &v
}

// GetOverdraftAccountId returns the OverdraftAccountId field value if set, zero value otherwise.
// Deprecated
func (o *BalanceFloor) GetOverdraftAccountId() string {
	if o == nil || o.OverdraftAccountId == nil {
		var ret string
		return ret
	}
	return *o.OverdraftAccountId
}

// GetOverdraftAccountIdOk returns a tuple with the OverdraftAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BalanceFloor) GetOverdraftAccountIdOk() (*string, bool) {
	if o == nil || o.OverdraftAccountId == nil {
		return nil, false
	}
	return o.OverdraftAccountId, true
}

// HasOverdraftAccountId returns a boolean if a field has been set.
func (o *BalanceFloor) HasOverdraftAccountId() bool {
	if o != nil && o.OverdraftAccountId != nil {
		return true
	}

	return false
}

// SetOverdraftAccountId gets a reference to the given string and assigns it to the OverdraftAccountId field.
// Deprecated
func (o *BalanceFloor) SetOverdraftAccountId(v string) {
	o.OverdraftAccountId = &v
}

func (o BalanceFloor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if o.LinkedAccountId != nil {
		toSerialize["linked_account_id"] = o.LinkedAccountId
	}
	if o.OverdraftAccountId != nil {
		toSerialize["overdraft_account_id"] = o.OverdraftAccountId
	}
	return json.Marshal(toSerialize)
}

type NullableBalanceFloor struct {
	value *BalanceFloor
	isSet bool
}

func (v NullableBalanceFloor) Get() *BalanceFloor {
	return v.value
}

func (v *NullableBalanceFloor) Set(val *BalanceFloor) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceFloor) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceFloor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceFloor(val *BalanceFloor) *NullableBalanceFloor {
	return &NullableBalanceFloor{value: val, isSet: true}
}

func (v NullableBalanceFloor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceFloor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


