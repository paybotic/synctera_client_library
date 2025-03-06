/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PaymentDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentDate{}

// PaymentDate struct for PaymentDate
type PaymentDate struct {
	// Execution date for the next payment
	ExecutionDate string `json:"execution_date"`
	// Scheduled date for the next payment
	ScheduledDate string `json:"scheduled_date"`
}

type _PaymentDate PaymentDate

// NewPaymentDate instantiates a new PaymentDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentDate(executionDate string, scheduledDate string) *PaymentDate {
	this := PaymentDate{}
	this.ExecutionDate = executionDate
	this.ScheduledDate = scheduledDate
	return &this
}

// NewPaymentDateWithDefaults instantiates a new PaymentDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentDateWithDefaults() *PaymentDate {
	this := PaymentDate{}
	return &this
}

// GetExecutionDate returns the ExecutionDate field value
func (o *PaymentDate) GetExecutionDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value
// and a boolean to check if the value has been set.
func (o *PaymentDate) GetExecutionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionDate, true
}

// SetExecutionDate sets field value
func (o *PaymentDate) SetExecutionDate(v string) {
	o.ExecutionDate = v
}

// GetScheduledDate returns the ScheduledDate field value
func (o *PaymentDate) GetScheduledDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledDate
}

// GetScheduledDateOk returns a tuple with the ScheduledDate field value
// and a boolean to check if the value has been set.
func (o *PaymentDate) GetScheduledDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledDate, true
}

// SetScheduledDate sets field value
func (o *PaymentDate) SetScheduledDate(v string) {
	o.ScheduledDate = v
}

func (o PaymentDate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["execution_date"] = o.ExecutionDate
	toSerialize["scheduled_date"] = o.ScheduledDate
	return toSerialize, nil
}

func (o *PaymentDate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"execution_date",
		"scheduled_date",
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

	varPaymentDate := _PaymentDate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentDate)

	if err != nil {
		return err
	}

	*o = PaymentDate(varPaymentDate)

	return err
}

type NullablePaymentDate struct {
	value *PaymentDate
	isSet bool
}

func (v NullablePaymentDate) Get() *PaymentDate {
	return v.value
}

func (v *NullablePaymentDate) Set(val *PaymentDate) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentDate) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentDate(val *PaymentDate) *NullablePaymentDate {
	return &NullablePaymentDate{value: val, isSet: true}
}

func (v NullablePaymentDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
