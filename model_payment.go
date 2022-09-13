/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Payment Executed payment
type Payment struct {
	// User provided description for the payment schedule
	Description *string `json:"description,omitempty"`
	ErrorDetails *PaymentErrorDetails `json:"error_details,omitempty"`
	// Payment ID
	Id *string `json:"id,omitempty"`
	// User provided JSON format data for the payment schedule
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	PaymentDate *PaymentDate `json:"payment_date,omitempty"`
	PaymentInstruction *PaymentInstruction `json:"payment_instruction,omitempty"`
	// ID of the payment schedule that executed this payment
	PaymentScheduleId *string `json:"payment_schedule_id,omitempty"`
	Status *PaymentStatus `json:"status,omitempty"`
	// Transaction ID. It will be included only when status is COMPLETED
	TransactionId *string `json:"transaction_id,omitempty"`
}

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment() *Payment {
	this := Payment{}
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Payment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Payment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Payment) SetDescription(v string) {
	o.Description = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *Payment) GetErrorDetails() PaymentErrorDetails {
	if o == nil || o.ErrorDetails == nil {
		var ret PaymentErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetErrorDetailsOk() (*PaymentErrorDetails, bool) {
	if o == nil || o.ErrorDetails == nil {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *Payment) HasErrorDetails() bool {
	if o != nil && o.ErrorDetails != nil {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given PaymentErrorDetails and assigns it to the ErrorDetails field.
func (o *Payment) SetErrorDetails(v PaymentErrorDetails) {
	o.ErrorDetails = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Payment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Payment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Payment) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Payment) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Payment) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Payment) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPaymentDate returns the PaymentDate field value if set, zero value otherwise.
func (o *Payment) GetPaymentDate() PaymentDate {
	if o == nil || o.PaymentDate == nil {
		var ret PaymentDate
		return ret
	}
	return *o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetPaymentDateOk() (*PaymentDate, bool) {
	if o == nil || o.PaymentDate == nil {
		return nil, false
	}
	return o.PaymentDate, true
}

// HasPaymentDate returns a boolean if a field has been set.
func (o *Payment) HasPaymentDate() bool {
	if o != nil && o.PaymentDate != nil {
		return true
	}

	return false
}

// SetPaymentDate gets a reference to the given PaymentDate and assigns it to the PaymentDate field.
func (o *Payment) SetPaymentDate(v PaymentDate) {
	o.PaymentDate = &v
}

// GetPaymentInstruction returns the PaymentInstruction field value if set, zero value otherwise.
func (o *Payment) GetPaymentInstruction() PaymentInstruction {
	if o == nil || o.PaymentInstruction == nil {
		var ret PaymentInstruction
		return ret
	}
	return *o.PaymentInstruction
}

// GetPaymentInstructionOk returns a tuple with the PaymentInstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetPaymentInstructionOk() (*PaymentInstruction, bool) {
	if o == nil || o.PaymentInstruction == nil {
		return nil, false
	}
	return o.PaymentInstruction, true
}

// HasPaymentInstruction returns a boolean if a field has been set.
func (o *Payment) HasPaymentInstruction() bool {
	if o != nil && o.PaymentInstruction != nil {
		return true
	}

	return false
}

// SetPaymentInstruction gets a reference to the given PaymentInstruction and assigns it to the PaymentInstruction field.
func (o *Payment) SetPaymentInstruction(v PaymentInstruction) {
	o.PaymentInstruction = &v
}

// GetPaymentScheduleId returns the PaymentScheduleId field value if set, zero value otherwise.
func (o *Payment) GetPaymentScheduleId() string {
	if o == nil || o.PaymentScheduleId == nil {
		var ret string
		return ret
	}
	return *o.PaymentScheduleId
}

// GetPaymentScheduleIdOk returns a tuple with the PaymentScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetPaymentScheduleIdOk() (*string, bool) {
	if o == nil || o.PaymentScheduleId == nil {
		return nil, false
	}
	return o.PaymentScheduleId, true
}

// HasPaymentScheduleId returns a boolean if a field has been set.
func (o *Payment) HasPaymentScheduleId() bool {
	if o != nil && o.PaymentScheduleId != nil {
		return true
	}

	return false
}

// SetPaymentScheduleId gets a reference to the given string and assigns it to the PaymentScheduleId field.
func (o *Payment) SetPaymentScheduleId(v string) {
	o.PaymentScheduleId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Payment) GetStatus() PaymentStatus {
	if o == nil || o.Status == nil {
		var ret PaymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetStatusOk() (*PaymentStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Payment) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PaymentStatus and assigns it to the Status field.
func (o *Payment) SetStatus(v PaymentStatus) {
	o.Status = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *Payment) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *Payment) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *Payment) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o Payment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ErrorDetails != nil {
		toSerialize["error_details"] = o.ErrorDetails
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PaymentDate != nil {
		toSerialize["payment_date"] = o.PaymentDate
	}
	if o.PaymentInstruction != nil {
		toSerialize["payment_instruction"] = o.PaymentInstruction
	}
	if o.PaymentScheduleId != nil {
		toSerialize["payment_schedule_id"] = o.PaymentScheduleId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


