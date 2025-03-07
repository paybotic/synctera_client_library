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

// AddVendorAccountsErrorReason A machine-readable code describing the reason for the failure.
type AddVendorAccountsErrorReason string

// List of add_vendor_accounts_error_reason
const (
	ADDVENDORACCOUNTSERRORREASON_ACCOUNT_NOT_FOUND        AddVendorAccountsErrorReason = "ACCOUNT_NOT_FOUND"
	ADDVENDORACCOUNTSERRORREASON_DUPLICATE_ACCOUNT        AddVendorAccountsErrorReason = "DUPLICATE_ACCOUNT"
	ADDVENDORACCOUNTSERRORREASON_FAILED_VERIFICATION      AddVendorAccountsErrorReason = "FAILED_VERIFICATION"
	ADDVENDORACCOUNTSERRORREASON_PROVIDER_ERROR           AddVendorAccountsErrorReason = "PROVIDER_ERROR"
	ADDVENDORACCOUNTSERRORREASON_UNSUPPORTED_ACCOUNT_TYPE AddVendorAccountsErrorReason = "UNSUPPORTED_ACCOUNT_TYPE"
)

// All allowed values of AddVendorAccountsErrorReason enum
var AllowedAddVendorAccountsErrorReasonEnumValues = []AddVendorAccountsErrorReason{
	"ACCOUNT_NOT_FOUND",
	"DUPLICATE_ACCOUNT",
	"FAILED_VERIFICATION",
	"PROVIDER_ERROR",
	"UNSUPPORTED_ACCOUNT_TYPE",
}

func (v *AddVendorAccountsErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AddVendorAccountsErrorReason(value)
	for _, existing := range AllowedAddVendorAccountsErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AddVendorAccountsErrorReason", value)
}

// NewAddVendorAccountsErrorReasonFromValue returns a pointer to a valid AddVendorAccountsErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAddVendorAccountsErrorReasonFromValue(v string) (*AddVendorAccountsErrorReason, error) {
	ev := AddVendorAccountsErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AddVendorAccountsErrorReason: valid values are %v", v, AllowedAddVendorAccountsErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AddVendorAccountsErrorReason) IsValid() bool {
	for _, existing := range AllowedAddVendorAccountsErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to add_vendor_accounts_error_reason value
func (v AddVendorAccountsErrorReason) Ptr() *AddVendorAccountsErrorReason {
	return &v
}

type NullableAddVendorAccountsErrorReason struct {
	value *AddVendorAccountsErrorReason
	isSet bool
}

func (v NullableAddVendorAccountsErrorReason) Get() *AddVendorAccountsErrorReason {
	return v.value
}

func (v *NullableAddVendorAccountsErrorReason) Set(val *AddVendorAccountsErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVendorAccountsErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVendorAccountsErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVendorAccountsErrorReason(val *AddVendorAccountsErrorReason) *NullableAddVendorAccountsErrorReason {
	return &NullableAddVendorAccountsErrorReason{value: val, isSet: true}
}

func (v NullableAddVendorAccountsErrorReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVendorAccountsErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
