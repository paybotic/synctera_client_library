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

// PersonalIdType The type of the personal identifier. This cannot be changed once the personal identifier is created. One of the following: * `ITIN` - US Individual Tax Identification Number. Format is 987-65-4321. Country code will default to US. * `PASSPORT_NUMBER` - Passport Number. Format varies by country. Country code is required. * `SIN` - Canadian Social Insurance Number. Format is 123-456-789. Country code will default to CA. * `SSN` - US Social Security Number. Format is 123-45-6789. Country code will default to US. * `DRIVER_LICENSE` - Driver License Number. Format varies by country and state. * `STATE_ID` - US State ID Number. Format varies by state. * `NATIONAL_ID` - National ID Number. Format varies by country. * `TAX_ID` - Tax ID Number. Format varies by country. * `CITIZENSHIP_CARD` - Canadian Citizenship Card Number. Format is A1234567. * `PR_CARD` - Permanent Resident Card Number. Format varies by country. * `PROVINCIAL_ID` - Canadian Provincial ID Number. Format varies by province. * `SECURE_STATUS_CARD` - Canadian Secure Status Card Number. Format is 1998-12-123456.
type PersonalIdType string

// List of personal_id_type
const (
	PERSONALIDTYPE_CITIZENSHIP_CARD   PersonalIdType = "CITIZENSHIP_CARD"
	PERSONALIDTYPE_DRIVER_LICENSE     PersonalIdType = "DRIVER_LICENSE"
	PERSONALIDTYPE_ITIN               PersonalIdType = "ITIN"
	PERSONALIDTYPE_NATIONAL_ID        PersonalIdType = "NATIONAL_ID"
	PERSONALIDTYPE_PASSPORT_NUMBER    PersonalIdType = "PASSPORT_NUMBER"
	PERSONALIDTYPE_PROVINCIAL_ID      PersonalIdType = "PROVINCIAL_ID"
	PERSONALIDTYPE_PR_CARD            PersonalIdType = "PR_CARD"
	PERSONALIDTYPE_SECURE_STATUS_CARD PersonalIdType = "SECURE_STATUS_CARD"
	PERSONALIDTYPE_SIN                PersonalIdType = "SIN"
	PERSONALIDTYPE_SSN                PersonalIdType = "SSN"
	PERSONALIDTYPE_STATE_ID           PersonalIdType = "STATE_ID"
	PERSONALIDTYPE_TAX_ID             PersonalIdType = "TAX_ID"
)

// All allowed values of PersonalIdType enum
var AllowedPersonalIdTypeEnumValues = []PersonalIdType{
	"CITIZENSHIP_CARD",
	"DRIVER_LICENSE",
	"ITIN",
	"NATIONAL_ID",
	"PASSPORT_NUMBER",
	"PROVINCIAL_ID",
	"PR_CARD",
	"SECURE_STATUS_CARD",
	"SIN",
	"SSN",
	"STATE_ID",
	"TAX_ID",
}

func (v *PersonalIdType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PersonalIdType(value)
	for _, existing := range AllowedPersonalIdTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PersonalIdType", value)
}

// NewPersonalIdTypeFromValue returns a pointer to a valid PersonalIdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPersonalIdTypeFromValue(v string) (*PersonalIdType, error) {
	ev := PersonalIdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PersonalIdType: valid values are %v", v, AllowedPersonalIdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PersonalIdType) IsValid() bool {
	for _, existing := range AllowedPersonalIdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to personal_id_type value
func (v PersonalIdType) Ptr() *PersonalIdType {
	return &v
}

type NullablePersonalIdType struct {
	value *PersonalIdType
	isSet bool
}

func (v NullablePersonalIdType) Get() *PersonalIdType {
	return v.value
}

func (v *NullablePersonalIdType) Set(val *PersonalIdType) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalIdType) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalIdType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalIdType(val *PersonalIdType) *NullablePersonalIdType {
	return &NullablePersonalIdType{value: val, isSet: true}
}

func (v NullablePersonalIdType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalIdType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}