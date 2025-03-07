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
	"time"
)

// checks if the HoldData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HoldData{}

// HoldData struct for HoldData
type HoldData struct {
	Amount               int32     `json:"amount"`
	AvailabilityTime     time.Time `json:"availability_time"`
	AdditionalProperties map[string]interface{}
}

type _HoldData HoldData

// NewHoldData instantiates a new HoldData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldData(amount int32, availabilityTime time.Time) *HoldData {
	this := HoldData{}
	this.Amount = amount
	this.AvailabilityTime = availabilityTime
	return &this
}

// NewHoldDataWithDefaults instantiates a new HoldData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldDataWithDefaults() *HoldData {
	this := HoldData{}
	return &this
}

// GetAmount returns the Amount field value
func (o *HoldData) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *HoldData) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *HoldData) SetAmount(v int32) {
	o.Amount = v
}

// GetAvailabilityTime returns the AvailabilityTime field value
func (o *HoldData) GetAvailabilityTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AvailabilityTime
}

// GetAvailabilityTimeOk returns a tuple with the AvailabilityTime field value
// and a boolean to check if the value has been set.
func (o *HoldData) GetAvailabilityTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityTime, true
}

// SetAvailabilityTime sets field value
func (o *HoldData) SetAvailabilityTime(v time.Time) {
	o.AvailabilityTime = v
}

func (o HoldData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HoldData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["availability_time"] = o.AvailabilityTime

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HoldData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"availability_time",
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

	varHoldData := _HoldData{}

	err = json.Unmarshal(data, &varHoldData)

	if err != nil {
		return err
	}

	*o = HoldData(varHoldData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "availability_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHoldData struct {
	value *HoldData
	isSet bool
}

func (v NullableHoldData) Get() *HoldData {
	return v.value
}

func (v *NullableHoldData) Set(val *HoldData) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldData) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldData(val *HoldData) *NullableHoldData {
	return &NullableHoldData{value: val, isSet: true}
}

func (v NullableHoldData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
