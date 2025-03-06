/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// checks if the Income type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Income{}

// Income The customer's income.
type Income struct {
	// The amount earned at the specified frequency. For example, $112.35 USD is represented as 11235 cents.
	Amount *int32 `json:"amount,omitempty"`
	// The currency in ISO 4217 format.
	Currency  *string           `json:"currency,omitempty" validate:"regexp=^[A-Z]{3}$"`
	Frequency NullableFrequency `json:"frequency,omitempty"`
	// The source of the income
	Source               NullableString `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Income Income

// NewIncome instantiates a new Income object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncome() *Income {
	this := Income{}
	return &this
}

// NewIncomeWithDefaults instantiates a new Income object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeWithDefaults() *Income {
	this := Income{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Income) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Income) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Income) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *Income) SetAmount(v int32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Income) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Income) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Income) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Income) SetCurrency(v string) {
	o.Currency = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Income) GetFrequency() Frequency {
	if o == nil || IsNil(o.Frequency.Get()) {
		var ret Frequency
		return ret
	}
	return *o.Frequency.Get()
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Income) GetFrequencyOk() (*Frequency, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frequency.Get(), o.Frequency.IsSet()
}

// HasFrequency returns a boolean if a field has been set.
func (o *Income) HasFrequency() bool {
	if o != nil && o.Frequency.IsSet() {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given NullableFrequency and assigns it to the Frequency field.
func (o *Income) SetFrequency(v Frequency) {
	o.Frequency.Set(&v)
}

// SetFrequencyNil sets the value for Frequency to be an explicit nil
func (o *Income) SetFrequencyNil() {
	o.Frequency.Set(nil)
}

// UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
func (o *Income) UnsetFrequency() {
	o.Frequency.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Income) GetSource() string {
	if o == nil || IsNil(o.Source.Get()) {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Income) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *Income) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *Income) SetSource(v string) {
	o.Source.Set(&v)
}

// SetSourceNil sets the value for Source to be an explicit nil
func (o *Income) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *Income) UnsetSource() {
	o.Source.Unset()
}

func (o Income) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Income) ToMap() (map[string]interface{}, error) {
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
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Income) UnmarshalJSON(data []byte) (err error) {
	varIncome := _Income{}

	err = json.Unmarshal(data, &varIncome)

	if err != nil {
		return err
	}

	*o = Income(varIncome)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "frequency")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncome struct {
	value *Income
	isSet bool
}

func (v NullableIncome) Get() *Income {
	return v.value
}

func (v *NullableIncome) Set(val *Income) {
	v.value = val
	v.isSet = true
}

func (v NullableIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncome(val *Income) *NullableIncome {
	return &NullableIncome{value: val, isSet: true}
}

func (v NullableIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
