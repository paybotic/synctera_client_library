/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// RelationshipRole CUSTODIAN - Related party is the custodian e.g. the parent, BENEFICIARY - Related party is the beneficiary e.g. the dependent, PARTNER - Related party is the partner
type RelationshipRole string

// List of relationship_role
const (
	RELATIONSHIPROLE_CUSTODIAN RelationshipRole = "CUSTODIAN"
	RELATIONSHIPROLE_BENEFICIARY RelationshipRole = "BENEFICIARY"
	RELATIONSHIPROLE_PARTNER RelationshipRole = "PARTNER"
)

// All allowed values of RelationshipRole enum
var AllowedRelationshipRoleEnumValues = []RelationshipRole{
	"CUSTODIAN",
	"BENEFICIARY",
	"PARTNER",
}

func (v *RelationshipRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RelationshipRole(value)
	for _, existing := range AllowedRelationshipRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RelationshipRole", value)
}

// NewRelationshipRoleFromValue returns a pointer to a valid RelationshipRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRelationshipRoleFromValue(v string) (*RelationshipRole, error) {
	ev := RelationshipRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RelationshipRole: valid values are %v", v, AllowedRelationshipRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RelationshipRole) IsValid() bool {
	for _, existing := range AllowedRelationshipRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to relationship_role value
func (v RelationshipRole) Ptr() *RelationshipRole {
	return &v
}

type NullableRelationshipRole struct {
	value *RelationshipRole
	isSet bool
}

func (v NullableRelationshipRole) Get() *RelationshipRole {
	return v.value
}

func (v *NullableRelationshipRole) Set(val *RelationshipRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipRole(val *RelationshipRole) *NullableRelationshipRole {
	return &NullableRelationshipRole{value: val, isSet: true}
}

func (v NullableRelationshipRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

