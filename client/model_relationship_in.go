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

// RelationshipIn - struct for RelationshipIn
type RelationshipIn struct {
	BusinessBusinessOwnerRelationship *BusinessBusinessOwnerRelationship
	PayerPayeeRelationship            *PayerPayeeRelationship
	PersonBusinessOwnerRelationship   *PersonBusinessOwnerRelationship
	PersonBusinessRelationship        *PersonBusinessRelationship
}

// BusinessBusinessOwnerRelationshipAsRelationshipIn is a convenience function that returns BusinessBusinessOwnerRelationship wrapped in RelationshipIn
func BusinessBusinessOwnerRelationshipAsRelationshipIn(v *BusinessBusinessOwnerRelationship) RelationshipIn {
	return RelationshipIn{
		BusinessBusinessOwnerRelationship: v,
	}
}

// PayerPayeeRelationshipAsRelationshipIn is a convenience function that returns PayerPayeeRelationship wrapped in RelationshipIn
func PayerPayeeRelationshipAsRelationshipIn(v *PayerPayeeRelationship) RelationshipIn {
	return RelationshipIn{
		PayerPayeeRelationship: v,
	}
}

// PersonBusinessOwnerRelationshipAsRelationshipIn is a convenience function that returns PersonBusinessOwnerRelationship wrapped in RelationshipIn
func PersonBusinessOwnerRelationshipAsRelationshipIn(v *PersonBusinessOwnerRelationship) RelationshipIn {
	return RelationshipIn{
		PersonBusinessOwnerRelationship: v,
	}
}

// PersonBusinessRelationshipAsRelationshipIn is a convenience function that returns PersonBusinessRelationship wrapped in RelationshipIn
func PersonBusinessRelationshipAsRelationshipIn(v *PersonBusinessRelationship) RelationshipIn {
	return RelationshipIn{
		PersonBusinessRelationship: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RelationshipIn) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BENEFICIAL_OWNER_OF'
	if jsonDict["relationship_type"] == "BENEFICIAL_OWNER_OF" {
		// try to unmarshal JSON data into PersonBusinessOwnerRelationship
		err = json.Unmarshal(data, &dst.PersonBusinessOwnerRelationship)
		if err == nil {
			return nil // data stored in dst.PersonBusinessOwnerRelationship, return on the first match
		} else {
			dst.PersonBusinessOwnerRelationship = nil
			return fmt.Errorf("failed to unmarshal RelationshipIn as PersonBusinessOwnerRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MANAGING_PERSON_OF'
	if jsonDict["relationship_type"] == "MANAGING_PERSON_OF" {
		// try to unmarshal JSON data into PersonBusinessRelationship
		err = json.Unmarshal(data, &dst.PersonBusinessRelationship)
		if err == nil {
			return nil // data stored in dst.PersonBusinessRelationship, return on the first match
		} else {
			dst.PersonBusinessRelationship = nil
			return fmt.Errorf("failed to unmarshal RelationshipIn as PersonBusinessRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OWNER_OF'
	if jsonDict["relationship_type"] == "OWNER_OF" {
		// try to unmarshal JSON data into BusinessBusinessOwnerRelationship
		err = json.Unmarshal(data, &dst.BusinessBusinessOwnerRelationship)
		if err == nil {
			return nil // data stored in dst.BusinessBusinessOwnerRelationship, return on the first match
		} else {
			dst.BusinessBusinessOwnerRelationship = nil
			return fmt.Errorf("failed to unmarshal RelationshipIn as BusinessBusinessOwnerRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PAYER_PAYEE'
	if jsonDict["relationship_type"] == "PAYER_PAYEE" {
		// try to unmarshal JSON data into PayerPayeeRelationship
		err = json.Unmarshal(data, &dst.PayerPayeeRelationship)
		if err == nil {
			return nil // data stored in dst.PayerPayeeRelationship, return on the first match
		} else {
			dst.PayerPayeeRelationship = nil
			return fmt.Errorf("failed to unmarshal RelationshipIn as PayerPayeeRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'business_business_owner_relationship'
	if jsonDict["relationship_type"] == "business_business_owner_relationship" {
		// try to unmarshal JSON data into BusinessBusinessOwnerRelationship
		err = json.Unmarshal(data, &dst.BusinessBusinessOwnerRelationship)
		if err == nil {
			return nil // data stored in dst.BusinessBusinessOwnerRelationship, return on the first match
		} else {
			dst.BusinessBusinessOwnerRelationship = nil
			return fmt.Errorf("failed to unmarshal RelationshipIn as BusinessBusinessOwnerRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'payer_payee_relationship'
	if jsonDict["relationship_type"] == "payer_payee_relationship" {
		// try to unmarshal JSON data into PayerPayeeRelationship
		err = json.Unmarshal(data, &dst.PayerPayeeRelationship)
		if err == nil {
			return nil // data stored in dst.PayerPayeeRelationship, return on the first match
		} else {
			dst.PayerPayeeRelationship = nil
			return fmt.Errorf("failed to unmarshal RelationshipIn as PayerPayeeRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'person_business_owner_relationship'
	if jsonDict["relationship_type"] == "person_business_owner_relationship" {
		// try to unmarshal JSON data into PersonBusinessOwnerRelationship
		err = json.Unmarshal(data, &dst.PersonBusinessOwnerRelationship)
		if err == nil {
			return nil // data stored in dst.PersonBusinessOwnerRelationship, return on the first match
		} else {
			dst.PersonBusinessOwnerRelationship = nil
			return fmt.Errorf("failed to unmarshal RelationshipIn as PersonBusinessOwnerRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'person_business_relationship'
	if jsonDict["relationship_type"] == "person_business_relationship" {
		// try to unmarshal JSON data into PersonBusinessRelationship
		err = json.Unmarshal(data, &dst.PersonBusinessRelationship)
		if err == nil {
			return nil // data stored in dst.PersonBusinessRelationship, return on the first match
		} else {
			dst.PersonBusinessRelationship = nil
			return fmt.Errorf("failed to unmarshal RelationshipIn as PersonBusinessRelationship: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RelationshipIn) MarshalJSON() ([]byte, error) {
	if src.BusinessBusinessOwnerRelationship != nil {
		return json.Marshal(&src.BusinessBusinessOwnerRelationship)
	}

	if src.PayerPayeeRelationship != nil {
		return json.Marshal(&src.PayerPayeeRelationship)
	}

	if src.PersonBusinessOwnerRelationship != nil {
		return json.Marshal(&src.PersonBusinessOwnerRelationship)
	}

	if src.PersonBusinessRelationship != nil {
		return json.Marshal(&src.PersonBusinessRelationship)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RelationshipIn) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BusinessBusinessOwnerRelationship != nil {
		return obj.BusinessBusinessOwnerRelationship
	}

	if obj.PayerPayeeRelationship != nil {
		return obj.PayerPayeeRelationship
	}

	if obj.PersonBusinessOwnerRelationship != nil {
		return obj.PersonBusinessOwnerRelationship
	}

	if obj.PersonBusinessRelationship != nil {
		return obj.PersonBusinessRelationship
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RelationshipIn) GetActualInstanceValue() interface{} {
	if obj.BusinessBusinessOwnerRelationship != nil {
		return *obj.BusinessBusinessOwnerRelationship
	}

	if obj.PayerPayeeRelationship != nil {
		return *obj.PayerPayeeRelationship
	}

	if obj.PersonBusinessOwnerRelationship != nil {
		return *obj.PersonBusinessOwnerRelationship
	}

	if obj.PersonBusinessRelationship != nil {
		return *obj.PersonBusinessRelationship
	}

	// all schemas are nil
	return nil
}

type NullableRelationshipIn struct {
	value *RelationshipIn
	isSet bool
}

func (v NullableRelationshipIn) Get() *RelationshipIn {
	return v.value
}

func (v *NullableRelationshipIn) Set(val *RelationshipIn) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipIn) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipIn(val *RelationshipIn) *NullableRelationshipIn {
	return &NullableRelationshipIn{value: val, isSet: true}
}

func (v NullableRelationshipIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
