/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// InternalTransferResponseStatus The status of the internal transfer auth. A value of `PENDING` indicates that the funds have been reserved and the transaction is ready to be either completed or canceled. A value of `COMPLETE` indicates the funds have been successfully moved and no more action can be performed. A value of `CANCELED` or `EXPIRED` means that the transaction has rolled back and the funds have been returned to the originating account, either by explicitly canceling via the API, or due to the expiry time having passed.
type InternalTransferResponseStatus string

// List of internal_transfer_response_status
const (
	INTERNALTRANSFERRESPONSESTATUS_CANCELED InternalTransferResponseStatus = "CANCELED"
	INTERNALTRANSFERRESPONSESTATUS_COMPLETE InternalTransferResponseStatus = "COMPLETE"
	INTERNALTRANSFERRESPONSESTATUS_EXPIRED  InternalTransferResponseStatus = "EXPIRED"
	INTERNALTRANSFERRESPONSESTATUS_PENDING  InternalTransferResponseStatus = "PENDING"
)

// All allowed values of InternalTransferResponseStatus enum
var AllowedInternalTransferResponseStatusEnumValues = []InternalTransferResponseStatus{
	"CANCELED",
	"COMPLETE",
	"EXPIRED",
	"PENDING",
}

func (v *InternalTransferResponseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InternalTransferResponseStatus(value)
	for _, existing := range AllowedInternalTransferResponseStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InternalTransferResponseStatus", value)
}

// NewInternalTransferResponseStatusFromValue returns a pointer to a valid InternalTransferResponseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInternalTransferResponseStatusFromValue(v string) (*InternalTransferResponseStatus, error) {
	ev := InternalTransferResponseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InternalTransferResponseStatus: valid values are %v", v, AllowedInternalTransferResponseStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InternalTransferResponseStatus) IsValid() bool {
	for _, existing := range AllowedInternalTransferResponseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to internal_transfer_response_status value
func (v InternalTransferResponseStatus) Ptr() *InternalTransferResponseStatus {
	return &v
}

type NullableInternalTransferResponseStatus struct {
	value *InternalTransferResponseStatus
	isSet bool
}

func (v NullableInternalTransferResponseStatus) Get() *InternalTransferResponseStatus {
	return v.value
}

func (v *NullableInternalTransferResponseStatus) Set(val *InternalTransferResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalTransferResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalTransferResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalTransferResponseStatus(val *InternalTransferResponseStatus) *NullableInternalTransferResponseStatus {
	return &NullableInternalTransferResponseStatus{value: val, isSet: true}
}

func (v NullableInternalTransferResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalTransferResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}