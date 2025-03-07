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

// VerificationStatus The result of a KYC/KYB verification. One of the following: * `UNVERIFIED` – verification has not been completed for this customer. * `PENDING` – verification is in progress for this customer. * `PROVISIONAL` – partially verified or verified with restrictions. * `ACCEPTED` – the customer has been verified. * `REVIEW` – verification has run and issues have been identified and require review. * `REJECTED` – the customer was rejected and should not be allowed to take certain actions e.g., open an account.
type VerificationStatus string

// List of verification_status
const (
	VERIFICATIONSTATUS_ACCEPTED    VerificationStatus = "ACCEPTED"
	VERIFICATIONSTATUS_PENDING     VerificationStatus = "PENDING"
	VERIFICATIONSTATUS_PROVISIONAL VerificationStatus = "PROVISIONAL"
	VERIFICATIONSTATUS_REJECTED    VerificationStatus = "REJECTED"
	VERIFICATIONSTATUS_REVIEW      VerificationStatus = "REVIEW"
	VERIFICATIONSTATUS_UNVERIFIED  VerificationStatus = "UNVERIFIED"
)

// All allowed values of VerificationStatus enum
var AllowedVerificationStatusEnumValues = []VerificationStatus{
	"ACCEPTED",
	"PENDING",
	"PROVISIONAL",
	"REJECTED",
	"REVIEW",
	"UNVERIFIED",
}

func (v *VerificationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerificationStatus(value)
	for _, existing := range AllowedVerificationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerificationStatus", value)
}

// NewVerificationStatusFromValue returns a pointer to a valid VerificationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerificationStatusFromValue(v string) (*VerificationStatus, error) {
	ev := VerificationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VerificationStatus: valid values are %v", v, AllowedVerificationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerificationStatus) IsValid() bool {
	for _, existing := range AllowedVerificationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to verification_status value
func (v VerificationStatus) Ptr() *VerificationStatus {
	return &v
}

type NullableVerificationStatus struct {
	value *VerificationStatus
	isSet bool
}

func (v NullableVerificationStatus) Get() *VerificationStatus {
	return v.value
}

func (v *NullableVerificationStatus) Set(val *VerificationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationStatus(val *VerificationStatus) *NullableVerificationStatus {
	return &NullableVerificationStatus{value: val, isSet: true}
}

func (v NullableVerificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
