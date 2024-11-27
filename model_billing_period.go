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
	"time"
)

// checks if the BillingPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPeriod{}

// BillingPeriod The scheme for determining an account's billing period.
type BillingPeriod struct {
	Frequency BillingFrequency `json:"frequency"`
	// The first day of the first billing cycle for this account. For a monthly billing cycle, this would determine the day of the month each billing cycle will start on. Note that, although this is returned as a UTC timestamp, the date always corresponds to the bank's calendar, and therefore the time and timezone should be ignored.
	StartDate time.Time `json:"start_date"`
}

type _BillingPeriod BillingPeriod

// NewBillingPeriod instantiates a new BillingPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPeriod(frequency BillingFrequency, startDate time.Time) *BillingPeriod {
	this := BillingPeriod{}
	this.Frequency = frequency
	this.StartDate = startDate
	return &this
}

// NewBillingPeriodWithDefaults instantiates a new BillingPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPeriodWithDefaults() *BillingPeriod {
	this := BillingPeriod{}
	return &this
}

// GetFrequency returns the Frequency field value
func (o *BillingPeriod) GetFrequency() BillingFrequency {
	if o == nil {
		var ret BillingFrequency
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *BillingPeriod) GetFrequencyOk() (*BillingFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *BillingPeriod) SetFrequency(v BillingFrequency) {
	o.Frequency = v
}

// GetStartDate returns the StartDate field value
func (o *BillingPeriod) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *BillingPeriod) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *BillingPeriod) SetStartDate(v time.Time) {
	o.StartDate = v
}

func (o BillingPeriod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["frequency"] = o.Frequency
	toSerialize["start_date"] = o.StartDate
	return toSerialize, nil
}

func (o *BillingPeriod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"frequency",
		"start_date",
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

	varBillingPeriod := _BillingPeriod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingPeriod)

	if err != nil {
		return err
	}

	*o = BillingPeriod(varBillingPeriod)

	return err
}

type NullableBillingPeriod struct {
	value *BillingPeriod
	isSet bool
}

func (v NullableBillingPeriod) Get() *BillingPeriod {
	return v.value
}

func (v *NullableBillingPeriod) Set(val *BillingPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPeriod(val *BillingPeriod) *NullableBillingPeriod {
	return &NullableBillingPeriod{value: val, isSet: true}
}

func (v NullableBillingPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
