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

// checks if the SpendControlSingleTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpendControlSingleTransaction{}

// SpendControlSingleTransaction struct for SpendControlSingleTransaction
type SpendControlSingleTransaction struct {
	TimeRangeType        string `json:"time_range_type"`
	AdditionalProperties map[string]interface{}
}

type _SpendControlSingleTransaction SpendControlSingleTransaction

// NewSpendControlSingleTransaction instantiates a new SpendControlSingleTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendControlSingleTransaction(timeRangeType string) *SpendControlSingleTransaction {
	this := SpendControlSingleTransaction{}
	this.TimeRangeType = timeRangeType
	return &this
}

// NewSpendControlSingleTransactionWithDefaults instantiates a new SpendControlSingleTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendControlSingleTransactionWithDefaults() *SpendControlSingleTransaction {
	this := SpendControlSingleTransaction{}
	return &this
}

// GetTimeRangeType returns the TimeRangeType field value
func (o *SpendControlSingleTransaction) GetTimeRangeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeRangeType
}

// GetTimeRangeTypeOk returns a tuple with the TimeRangeType field value
// and a boolean to check if the value has been set.
func (o *SpendControlSingleTransaction) GetTimeRangeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRangeType, true
}

// SetTimeRangeType sets field value
func (o *SpendControlSingleTransaction) SetTimeRangeType(v string) {
	o.TimeRangeType = v
}

func (o SpendControlSingleTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpendControlSingleTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["time_range_type"] = o.TimeRangeType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpendControlSingleTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"time_range_type",
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

	varSpendControlSingleTransaction := _SpendControlSingleTransaction{}

	err = json.Unmarshal(data, &varSpendControlSingleTransaction)

	if err != nil {
		return err
	}

	*o = SpendControlSingleTransaction(varSpendControlSingleTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time_range_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpendControlSingleTransaction struct {
	value *SpendControlSingleTransaction
	isSet bool
}

func (v NullableSpendControlSingleTransaction) Get() *SpendControlSingleTransaction {
	return v.value
}

func (v *NullableSpendControlSingleTransaction) Set(val *SpendControlSingleTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendControlSingleTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendControlSingleTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendControlSingleTransaction(val *SpendControlSingleTransaction) *NullableSpendControlSingleTransaction {
	return &NullableSpendControlSingleTransaction{value: val, isSet: true}
}

func (v NullableSpendControlSingleTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendControlSingleTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
