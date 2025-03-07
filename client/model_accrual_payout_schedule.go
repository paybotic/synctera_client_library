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

// AccrualPayoutSchedule the model 'AccrualPayoutSchedule'
type AccrualPayoutSchedule string

// List of accrual_payout_schedule
const (
	ACCRUALPAYOUTSCHEDULE_MONTHLY AccrualPayoutSchedule = "MONTHLY"
	ACCRUALPAYOUTSCHEDULE_NONE    AccrualPayoutSchedule = "NONE"
)

// All allowed values of AccrualPayoutSchedule enum
var AllowedAccrualPayoutScheduleEnumValues = []AccrualPayoutSchedule{
	"MONTHLY",
	"NONE",
}

func (v *AccrualPayoutSchedule) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccrualPayoutSchedule(value)
	for _, existing := range AllowedAccrualPayoutScheduleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccrualPayoutSchedule", value)
}

// NewAccrualPayoutScheduleFromValue returns a pointer to a valid AccrualPayoutSchedule
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccrualPayoutScheduleFromValue(v string) (*AccrualPayoutSchedule, error) {
	ev := AccrualPayoutSchedule(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccrualPayoutSchedule: valid values are %v", v, AllowedAccrualPayoutScheduleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccrualPayoutSchedule) IsValid() bool {
	for _, existing := range AllowedAccrualPayoutScheduleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to accrual_payout_schedule value
func (v AccrualPayoutSchedule) Ptr() *AccrualPayoutSchedule {
	return &v
}

type NullableAccrualPayoutSchedule struct {
	value *AccrualPayoutSchedule
	isSet bool
}

func (v NullableAccrualPayoutSchedule) Get() *AccrualPayoutSchedule {
	return v.value
}

func (v *NullableAccrualPayoutSchedule) Set(val *AccrualPayoutSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableAccrualPayoutSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableAccrualPayoutSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccrualPayoutSchedule(val *AccrualPayoutSchedule) *NullableAccrualPayoutSchedule {
	return &NullableAccrualPayoutSchedule{value: val, isSet: true}
}

func (v NullableAccrualPayoutSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccrualPayoutSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
