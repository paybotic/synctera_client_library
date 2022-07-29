/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// IngestionStatus Reconciliation ingestion status values
type IngestionStatus string

// List of ingestion_status
const (
	INGESTIONSTATUS_IN_PROCESS IngestionStatus = "IN_PROCESS"
	INGESTIONSTATUS_COMPLETED IngestionStatus = "COMPLETED"
	INGESTIONSTATUS_FAILED IngestionStatus = "FAILED"
)

// All allowed values of IngestionStatus enum
var AllowedIngestionStatusEnumValues = []IngestionStatus{
	"IN_PROCESS",
	"COMPLETED",
	"FAILED",
}

func (v *IngestionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IngestionStatus(value)
	for _, existing := range AllowedIngestionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IngestionStatus", value)
}

// NewIngestionStatusFromValue returns a pointer to a valid IngestionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIngestionStatusFromValue(v string) (*IngestionStatus, error) {
	ev := IngestionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IngestionStatus: valid values are %v", v, AllowedIngestionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IngestionStatus) IsValid() bool {
	for _, existing := range AllowedIngestionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ingestion_status value
func (v IngestionStatus) Ptr() *IngestionStatus {
	return &v
}

type NullableIngestionStatus struct {
	value *IngestionStatus
	isSet bool
}

func (v NullableIngestionStatus) Get() *IngestionStatus {
	return v.value
}

func (v *NullableIngestionStatus) Set(val *IngestionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionStatus(val *IngestionStatus) *NullableIngestionStatus {
	return &NullableIngestionStatus{value: val, isSet: true}
}

func (v NullableIngestionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

