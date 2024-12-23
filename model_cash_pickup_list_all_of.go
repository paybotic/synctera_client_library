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

// CashPickupListAllOf struct for CashPickupListAllOf
type CashPickupListAllOf struct {
	// Array of cash pickups
	CashPickups []CashPickup `json:"cash_pickups"`
}

// NewCashPickupListAllOf instantiates a new CashPickupListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashPickupListAllOf(cashPickups []CashPickup) *CashPickupListAllOf {
	this := CashPickupListAllOf{}
	this.CashPickups = cashPickups
	return &this
}

// NewCashPickupListAllOfWithDefaults instantiates a new CashPickupListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashPickupListAllOfWithDefaults() *CashPickupListAllOf {
	this := CashPickupListAllOf{}
	return &this
}

// GetCashPickups returns the CashPickups field value
func (o *CashPickupListAllOf) GetCashPickups() []CashPickup {
	if o == nil {
		var ret []CashPickup
		return ret
	}

	return o.CashPickups
}

// GetCashPickupsOk returns a tuple with the CashPickups field value
// and a boolean to check if the value has been set.
func (o *CashPickupListAllOf) GetCashPickupsOk() ([]CashPickup, bool) {
	if o == nil {
		return nil, false
	}
	return o.CashPickups, true
}

// SetCashPickups sets field value
func (o *CashPickupListAllOf) SetCashPickups(v []CashPickup) {
	o.CashPickups = v
}

func (o CashPickupListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cash_pickups"] = o.CashPickups
	}
	return json.Marshal(toSerialize)
}

type NullableCashPickupListAllOf struct {
	value *CashPickupListAllOf
	isSet bool
}

func (v NullableCashPickupListAllOf) Get() *CashPickupListAllOf {
	return v.value
}

func (v *NullableCashPickupListAllOf) Set(val *CashPickupListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCashPickupListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCashPickupListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashPickupListAllOf(val *CashPickupListAllOf) *NullableCashPickupListAllOf {
	return &NullableCashPickupListAllOf{value: val, isSet: true}
}

func (v NullableCashPickupListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashPickupListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
