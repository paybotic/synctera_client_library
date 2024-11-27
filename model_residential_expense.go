/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResidentialExpense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResidentialExpense{}

// ResidentialExpense The residential expense.
type ResidentialExpense struct {
	// The amount paid in residential expenses at the specified frequency. For example, $112.35 USD is represented as 11235 cents).
	Amount *int32 `json:"amount,omitempty"`
	// The currency in ISO 4217 format.
	Currency  *string           `json:"currency,omitempty" validate:"regexp=^[A-Z]{3}$"`
	Frequency NullableFrequency `json:"frequency,omitempty"`
}

// NewResidentialExpense instantiates a new ResidentialExpense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResidentialExpense() *ResidentialExpense {
	this := ResidentialExpense{}
	return &this
}

// NewResidentialExpenseWithDefaults instantiates a new ResidentialExpense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResidentialExpenseWithDefaults() *ResidentialExpense {
	this := ResidentialExpense{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ResidentialExpense) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResidentialExpense) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ResidentialExpense) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ResidentialExpense) SetAmount(v int32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ResidentialExpense) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResidentialExpense) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ResidentialExpense) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ResidentialExpense) SetCurrency(v string) {
	o.Currency = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResidentialExpense) GetFrequency() Frequency {
	if o == nil || IsNil(o.Frequency.Get()) {
		var ret Frequency
		return ret
	}
	return *o.Frequency.Get()
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResidentialExpense) GetFrequencyOk() (*Frequency, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frequency.Get(), o.Frequency.IsSet()
}

// HasFrequency returns a boolean if a field has been set.
func (o *ResidentialExpense) HasFrequency() bool {
	if o != nil && o.Frequency.IsSet() {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given NullableFrequency and assigns it to the Frequency field.
func (o *ResidentialExpense) SetFrequency(v Frequency) {
	o.Frequency.Set(&v)
}

// SetFrequencyNil sets the value for Frequency to be an explicit nil
func (o *ResidentialExpense) SetFrequencyNil() {
	o.Frequency.Set(nil)
}

// UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
func (o *ResidentialExpense) UnsetFrequency() {
	o.Frequency.Unset()
}

func (o ResidentialExpense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResidentialExpense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if o.Frequency.IsSet() {
		toSerialize["frequency"] = o.Frequency.Get()
	}
	return toSerialize, nil
}

type NullableResidentialExpense struct {
	value *ResidentialExpense
	isSet bool
}

func (v NullableResidentialExpense) Get() *ResidentialExpense {
	return v.value
}

func (v *NullableResidentialExpense) Set(val *ResidentialExpense) {
	v.value = val
	v.isSet = true
}

func (v NullableResidentialExpense) IsSet() bool {
	return v.isSet
}

func (v *NullableResidentialExpense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResidentialExpense(val *ResidentialExpense) *NullableResidentialExpense {
	return &NullableResidentialExpense{value: val, isSet: true}
}

func (v NullableResidentialExpense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResidentialExpense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}