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

// RelationshipTypes The relationship type. One of the following: * `BENEFICIAL_OWNER_OF` – a person who directly or indirectly owns a portion of the business. * `MANAGING_PERSON_OF` – a person who is an officer, director, or other notable person of an organization. * `OWNER_OF` – a business with ownership of another business. * `PAYER_PAYEE` - a person or business with a payer payee relationship with another person or business
type RelationshipTypes string

// List of relationship_types
const (
	RELATIONSHIPTYPES_BENEFICIAL_OWNER_OF RelationshipTypes = "BENEFICIAL_OWNER_OF"
	RELATIONSHIPTYPES_MANAGING_PERSON_OF  RelationshipTypes = "MANAGING_PERSON_OF"
	RELATIONSHIPTYPES_OWNER_OF            RelationshipTypes = "OWNER_OF"
	RELATIONSHIPTYPES_PAYER_PAYEE         RelationshipTypes = "PAYER_PAYEE"
)

// All allowed values of RelationshipTypes enum
var AllowedRelationshipTypesEnumValues = []RelationshipTypes{
	"BENEFICIAL_OWNER_OF",
	"MANAGING_PERSON_OF",
	"OWNER_OF",
	"PAYER_PAYEE",
}

func (v *RelationshipTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RelationshipTypes(value)
	for _, existing := range AllowedRelationshipTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RelationshipTypes", value)
}

// NewRelationshipTypesFromValue returns a pointer to a valid RelationshipTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRelationshipTypesFromValue(v string) (*RelationshipTypes, error) {
	ev := RelationshipTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RelationshipTypes: valid values are %v", v, AllowedRelationshipTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RelationshipTypes) IsValid() bool {
	for _, existing := range AllowedRelationshipTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to relationship_types value
func (v RelationshipTypes) Ptr() *RelationshipTypes {
	return &v
}

type NullableRelationshipTypes struct {
	value *RelationshipTypes
	isSet bool
}

func (v NullableRelationshipTypes) Get() *RelationshipTypes {
	return v.value
}

func (v *NullableRelationshipTypes) Set(val *RelationshipTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipTypes(val *RelationshipTypes) *NullableRelationshipTypes {
	return &NullableRelationshipTypes{value: val, isSet: true}
}

func (v NullableRelationshipTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
