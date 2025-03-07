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

// checks if the AchRequestHoldData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AchRequestHoldData{}

// AchRequestHoldData struct for AchRequestHoldData
type AchRequestHoldData struct {
	Amount               int32 `json:"amount"`
	Duration             int32 `json:"duration"`
	AdditionalProperties map[string]interface{}
}

type _AchRequestHoldData AchRequestHoldData

// NewAchRequestHoldData instantiates a new AchRequestHoldData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchRequestHoldData(amount int32, duration int32) *AchRequestHoldData {
	this := AchRequestHoldData{}
	this.Amount = amount
	this.Duration = duration
	return &this
}

// NewAchRequestHoldDataWithDefaults instantiates a new AchRequestHoldData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchRequestHoldDataWithDefaults() *AchRequestHoldData {
	this := AchRequestHoldData{}
	return &this
}

// GetAmount returns the Amount field value
func (o *AchRequestHoldData) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AchRequestHoldData) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AchRequestHoldData) SetAmount(v int32) {
	o.Amount = v
}

// GetDuration returns the Duration field value
func (o *AchRequestHoldData) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *AchRequestHoldData) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *AchRequestHoldData) SetDuration(v int32) {
	o.Duration = v
}

func (o AchRequestHoldData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AchRequestHoldData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["duration"] = o.Duration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AchRequestHoldData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"duration",
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

	varAchRequestHoldData := _AchRequestHoldData{}

	err = json.Unmarshal(data, &varAchRequestHoldData)

	if err != nil {
		return err
	}

	*o = AchRequestHoldData(varAchRequestHoldData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "duration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAchRequestHoldData struct {
	value *AchRequestHoldData
	isSet bool
}

func (v NullableAchRequestHoldData) Get() *AchRequestHoldData {
	return v.value
}

func (v *NullableAchRequestHoldData) Set(val *AchRequestHoldData) {
	v.value = val
	v.isSet = true
}

func (v NullableAchRequestHoldData) IsSet() bool {
	return v.isSet
}

func (v *NullableAchRequestHoldData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchRequestHoldData(val *AchRequestHoldData) *NullableAchRequestHoldData {
	return &NullableAchRequestHoldData{value: val, isSet: true}
}

func (v NullableAchRequestHoldData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchRequestHoldData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
