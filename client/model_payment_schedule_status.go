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

// PaymentScheduleStatus Status of the payment schedule.
type PaymentScheduleStatus string

// List of payment_schedule_status
const (
	PAYMENTSCHEDULESTATUS_ACTIVE    PaymentScheduleStatus = "ACTIVE"
	PAYMENTSCHEDULESTATUS_CANCELLED PaymentScheduleStatus = "CANCELLED"
	PAYMENTSCHEDULESTATUS_ERRORED   PaymentScheduleStatus = "ERRORED"
	PAYMENTSCHEDULESTATUS_EXPIRED   PaymentScheduleStatus = "EXPIRED"
)

// All allowed values of PaymentScheduleStatus enum
var AllowedPaymentScheduleStatusEnumValues = []PaymentScheduleStatus{
	"ACTIVE",
	"CANCELLED",
	"ERRORED",
	"EXPIRED",
}

func (v *PaymentScheduleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentScheduleStatus(value)
	for _, existing := range AllowedPaymentScheduleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentScheduleStatus", value)
}

// NewPaymentScheduleStatusFromValue returns a pointer to a valid PaymentScheduleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentScheduleStatusFromValue(v string) (*PaymentScheduleStatus, error) {
	ev := PaymentScheduleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentScheduleStatus: valid values are %v", v, AllowedPaymentScheduleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentScheduleStatus) IsValid() bool {
	for _, existing := range AllowedPaymentScheduleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to payment_schedule_status value
func (v PaymentScheduleStatus) Ptr() *PaymentScheduleStatus {
	return &v
}

type NullablePaymentScheduleStatus struct {
	value *PaymentScheduleStatus
	isSet bool
}

func (v NullablePaymentScheduleStatus) Get() *PaymentScheduleStatus {
	return v.value
}

func (v *NullablePaymentScheduleStatus) Set(val *PaymentScheduleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentScheduleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentScheduleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentScheduleStatus(val *PaymentScheduleStatus) *NullablePaymentScheduleStatus {
	return &NullablePaymentScheduleStatus{value: val, isSet: true}
}

func (v NullablePaymentScheduleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentScheduleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
