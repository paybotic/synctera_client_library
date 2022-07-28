/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PaymentDate struct for PaymentDate
type PaymentDate struct {
	// Execution date for the next payment
	ExecutionDate string `json:"execution_date"`
	// Scheduled date for the next payment
	ScheduledDate string `json:"scheduled_date"`
}

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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["execution_date"] = o.ExecutionDate
	}
	if true {
		toSerialize["scheduled_date"] = o.ScheduledDate
	}
	return json.Marshal(toSerialize)
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


