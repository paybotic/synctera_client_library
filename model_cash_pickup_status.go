/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CashPickupStatus the model 'CashPickupStatus'
type CashPickupStatus string

// List of cash_pickup_status
const (
	CASHPICKUPSTATUS_PENDING  CashPickupStatus = "PENDING"
	CASHPICKUPSTATUS_CANCELED CashPickupStatus = "CANCELED"
	CASHPICKUPSTATUS_SUSPENSE CashPickupStatus = "SUSPENSE"
	CASHPICKUPSTATUS_COMPLETE CashPickupStatus = "COMPLETE"
)

// All allowed values of CashPickupStatus enum
var AllowedCashPickupStatusEnumValues = []CashPickupStatus{
	"PENDING",
	"CANCELED",
	"SUSPENSE",
	"COMPLETE",
}

func (v *CashPickupStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CashPickupStatus(value)
	for _, existing := range AllowedCashPickupStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CashPickupStatus", value)
}

// NewCashPickupStatusFromValue returns a pointer to a valid CashPickupStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCashPickupStatusFromValue(v string) (*CashPickupStatus, error) {
	ev := CashPickupStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CashPickupStatus: valid values are %v", v, AllowedCashPickupStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CashPickupStatus) IsValid() bool {
	for _, existing := range AllowedCashPickupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cash_pickup_status value
func (v CashPickupStatus) Ptr() *CashPickupStatus {
	return &v
}

type NullableCashPickupStatus struct {
	value *CashPickupStatus
	isSet bool
}

func (v NullableCashPickupStatus) Get() *CashPickupStatus {
	return v.value
}

func (v *NullableCashPickupStatus) Set(val *CashPickupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCashPickupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCashPickupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashPickupStatus(val *CashPickupStatus) *NullableCashPickupStatus {
	return &NullableCashPickupStatus{value: val, isSet: true}
}

func (v NullableCashPickupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashPickupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
