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

// DisclosureType Describes the regulatory requirement that triggered the disclosure. One of the following: * `ACH_AUTHORIZATION` –     The ACH Authorization & Agreement Disclosure serves as a method of disclosing     and obtaining consent from a consumer to conduct ACH and pre-authorized     electronic transactions to and from their account. Record of obtaining this     consent should be retained, in accordance with Reg E, for no less than 2 years     from the date the disclosure was made to the consumer. * `CARDHOLDER_AGREEMENT` –     The Cardholder Agreement is a legal document that details the terms of a card     agreement between either a consumer or a business, and the financial     institution that issues the card. The Agreement governs the use of the     account. * `E_SIGN` –     US law with rules around electronic agreements/documents/disclosures:     used to obtain consent from consumers to receive electronic communications     (agreements, disclosures, statements, etc) about their accounts. * `KYC_DATA_COLLECTION` –     Document advising the customer that you will collect their personal information     (name, date of birth, tax ID, etc.)     and will be validating their information against external data/documentation. * `PRIVACY_NOTICE` –     Document that tells customers what is done with their non-public information,     who it is shared with, how is is secured,     and how they can opt out of it being shared beyond Synctera. * `REG_CC` –     US regulation that implements the Expedited Funds Availability Act:     describes standards for when a financial institution     makes funds available in a deposit account. * `REG_DD` –     US regulation that implements the Truth in Savings Act,     to inform customers about the terms and rules for a deposit account. * `REG_E` –     US regulation that implements the Electronic Funds Transfer Act:     covers liability for electronic transactions,     disputes for fraudulent or unrecognized electronic transactions,     and consent for electronic debits from a consumer's account. * `TERMS_AND_CONDITIONS` –     A detailed agreement between you and the consumer for the     structure, terms, fees, charges, rates of the product or service,     and the rules for the relationship between you and the consumer. 
type DisclosureType string

// List of disclosure_type
const (
	DISCLOSURETYPE_ACH_AUTHORIZATION DisclosureType = "ACH_AUTHORIZATION"
	DISCLOSURETYPE_CARDHOLDER_AGREEMENT DisclosureType = "CARDHOLDER_AGREEMENT"
	DISCLOSURETYPE_E_SIGN DisclosureType = "E_SIGN"
	DISCLOSURETYPE_KYC_DATA_COLLECTION DisclosureType = "KYC_DATA_COLLECTION"
	DISCLOSURETYPE_PRIVACY_NOTICE DisclosureType = "PRIVACY_NOTICE"
	DISCLOSURETYPE_REG_CC DisclosureType = "REG_CC"
	DISCLOSURETYPE_REG_DD DisclosureType = "REG_DD"
	DISCLOSURETYPE_REG_E DisclosureType = "REG_E"
	DISCLOSURETYPE_TERMS_AND_CONDITIONS DisclosureType = "TERMS_AND_CONDITIONS"
)

// All allowed values of DisclosureType enum
var AllowedDisclosureTypeEnumValues = []DisclosureType{
	"ACH_AUTHORIZATION",
	"CARDHOLDER_AGREEMENT",
	"E_SIGN",
	"KYC_DATA_COLLECTION",
	"PRIVACY_NOTICE",
	"REG_CC",
	"REG_DD",
	"REG_E",
	"TERMS_AND_CONDITIONS",
}

func (v *DisclosureType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DisclosureType(value)
	for _, existing := range AllowedDisclosureTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DisclosureType", value)
}

// NewDisclosureTypeFromValue returns a pointer to a valid DisclosureType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDisclosureTypeFromValue(v string) (*DisclosureType, error) {
	ev := DisclosureType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DisclosureType: valid values are %v", v, AllowedDisclosureTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DisclosureType) IsValid() bool {
	for _, existing := range AllowedDisclosureTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to disclosure_type value
func (v DisclosureType) Ptr() *DisclosureType {
	return &v
}

type NullableDisclosureType struct {
	value *DisclosureType
	isSet bool
}

func (v NullableDisclosureType) Get() *DisclosureType {
	return v.value
}

func (v *NullableDisclosureType) Set(val *DisclosureType) {
	v.value = val
	v.isSet = true
}

func (v NullableDisclosureType) IsSet() bool {
	return v.isSet
}

func (v *NullableDisclosureType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisclosureType(val *DisclosureType) *NullableDisclosureType {
	return &NullableDisclosureType{value: val, isSet: true}
}

func (v NullableDisclosureType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisclosureType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

