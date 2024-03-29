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

// PaymentSchedule Payment schedule
type PaymentSchedule struct {
	// User provided description for the payment schedule
	Description string `json:"description"`
	// Payment schedule ID
	Id *string `json:"id,omitempty"`
	// User provided JSON format data
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	NextPaymentDate *PaymentDate `json:"next_payment_date,omitempty"`
	PaymentInstruction PaymentInstruction `json:"payment_instruction"`
	Schedule ScheduleConfig `json:"schedule"`
	Status *PaymentScheduleStatus `json:"status,omitempty"`
}

// NewPaymentSchedule instantiates a new PaymentSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentSchedule(description string, paymentInstruction PaymentInstruction, schedule ScheduleConfig) *PaymentSchedule {
	this := PaymentSchedule{}
	this.Description = description
	this.PaymentInstruction = paymentInstruction
	this.Schedule = schedule
	return &this
}

// NewPaymentScheduleWithDefaults instantiates a new PaymentSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentScheduleWithDefaults() *PaymentSchedule {
	this := PaymentSchedule{}
	return &this
}

// GetDescription returns the Description field value
func (o *PaymentSchedule) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PaymentSchedule) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PaymentSchedule) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentSchedule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSchedule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentSchedule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentSchedule) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PaymentSchedule) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSchedule) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentSchedule) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentSchedule) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNextPaymentDate returns the NextPaymentDate field value if set, zero value otherwise.
func (o *PaymentSchedule) GetNextPaymentDate() PaymentDate {
	if o == nil || o.NextPaymentDate == nil {
		var ret PaymentDate
		return ret
	}
	return *o.NextPaymentDate
}

// GetNextPaymentDateOk returns a tuple with the NextPaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSchedule) GetNextPaymentDateOk() (*PaymentDate, bool) {
	if o == nil || o.NextPaymentDate == nil {
		return nil, false
	}
	return o.NextPaymentDate, true
}

// HasNextPaymentDate returns a boolean if a field has been set.
func (o *PaymentSchedule) HasNextPaymentDate() bool {
	if o != nil && o.NextPaymentDate != nil {
		return true
	}

	return false
}

// SetNextPaymentDate gets a reference to the given PaymentDate and assigns it to the NextPaymentDate field.
func (o *PaymentSchedule) SetNextPaymentDate(v PaymentDate) {
	o.NextPaymentDate = &v
}

// GetPaymentInstruction returns the PaymentInstruction field value
func (o *PaymentSchedule) GetPaymentInstruction() PaymentInstruction {
	if o == nil {
		var ret PaymentInstruction
		return ret
	}

	return o.PaymentInstruction
}

// GetPaymentInstructionOk returns a tuple with the PaymentInstruction field value
// and a boolean to check if the value has been set.
func (o *PaymentSchedule) GetPaymentInstructionOk() (*PaymentInstruction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentInstruction, true
}

// SetPaymentInstruction sets field value
func (o *PaymentSchedule) SetPaymentInstruction(v PaymentInstruction) {
	o.PaymentInstruction = v
}

// GetSchedule returns the Schedule field value
func (o *PaymentSchedule) GetSchedule() ScheduleConfig {
	if o == nil {
		var ret ScheduleConfig
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *PaymentSchedule) GetScheduleOk() (*ScheduleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *PaymentSchedule) SetSchedule(v ScheduleConfig) {
	o.Schedule = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentSchedule) GetStatus() PaymentScheduleStatus {
	if o == nil || o.Status == nil {
		var ret PaymentScheduleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSchedule) GetStatusOk() (*PaymentScheduleStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentSchedule) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PaymentScheduleStatus and assigns it to the Status field.
func (o *PaymentSchedule) SetStatus(v PaymentScheduleStatus) {
	o.Status = &v
}

func (o PaymentSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.NextPaymentDate != nil {
		toSerialize["next_payment_date"] = o.NextPaymentDate
	}
	if true {
		toSerialize["payment_instruction"] = o.PaymentInstruction
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentSchedule struct {
	value *PaymentSchedule
	isSet bool
}

func (v NullablePaymentSchedule) Get() *PaymentSchedule {
	return v.value
}

func (v *NullablePaymentSchedule) Set(val *PaymentSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentSchedule(val *PaymentSchedule) *NullablePaymentSchedule {
	return &NullablePaymentSchedule{value: val, isSet: true}
}

func (v NullablePaymentSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


