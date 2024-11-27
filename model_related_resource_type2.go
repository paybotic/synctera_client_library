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

// RelatedResourceType2 Type of the resource associated with the note.
type RelatedResourceType2 string

// List of related_resource_type2
const (
	RELATEDRESOURCETYPE2_ACCOUNT            RelatedResourceType2 = "ACCOUNT"
	RELATEDRESOURCETYPE2_APPLICATION        RelatedResourceType2 = "APPLICATION"
	RELATEDRESOURCETYPE2_BUSINESS           RelatedResourceType2 = "BUSINESS"
	RELATEDRESOURCETYPE2_CUSTOMER           RelatedResourceType2 = "CUSTOMER"
	RELATEDRESOURCETYPE2_FILE               RelatedResourceType2 = "FILE"
	RELATEDRESOURCETYPE2_SHADOW_TRANSACTION RelatedResourceType2 = "SHADOW_TRANSACTION"
	RELATEDRESOURCETYPE2_SNAPSHOT           RelatedResourceType2 = "SNAPSHOT"
	RELATEDRESOURCETYPE2_TENANT             RelatedResourceType2 = "TENANT"
	RELATEDRESOURCETYPE2_TRANSACTION        RelatedResourceType2 = "TRANSACTION"
	RELATEDRESOURCETYPE2_USER               RelatedResourceType2 = "USER"
)

// All allowed values of RelatedResourceType2 enum
var AllowedRelatedResourceType2EnumValues = []RelatedResourceType2{
	"ACCOUNT",
	"APPLICATION",
	"BUSINESS",
	"CUSTOMER",
	"FILE",
	"SHADOW_TRANSACTION",
	"SNAPSHOT",
	"TENANT",
	"TRANSACTION",
	"USER",
}

func (v *RelatedResourceType2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RelatedResourceType2(value)
	for _, existing := range AllowedRelatedResourceType2EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RelatedResourceType2", value)
}

// NewRelatedResourceType2FromValue returns a pointer to a valid RelatedResourceType2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRelatedResourceType2FromValue(v string) (*RelatedResourceType2, error) {
	ev := RelatedResourceType2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RelatedResourceType2: valid values are %v", v, AllowedRelatedResourceType2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RelatedResourceType2) IsValid() bool {
	for _, existing := range AllowedRelatedResourceType2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to related_resource_type2 value
func (v RelatedResourceType2) Ptr() *RelatedResourceType2 {
	return &v
}

type NullableRelatedResourceType2 struct {
	value *RelatedResourceType2
	isSet bool
}

func (v NullableRelatedResourceType2) Get() *RelatedResourceType2 {
	return v.value
}

func (v *NullableRelatedResourceType2) Set(val *RelatedResourceType2) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedResourceType2) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedResourceType2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedResourceType2(val *RelatedResourceType2) *NullableRelatedResourceType2 {
	return &NullableRelatedResourceType2{value: val, isSet: true}
}

func (v NullableRelatedResourceType2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedResourceType2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}