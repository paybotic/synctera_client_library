/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PatchRelationshipIn - struct for PatchRelationshipIn
type PatchRelationshipIn struct {
	PatchBusinessBusinessOwnerRelationship *PatchBusinessBusinessOwnerRelationship
	PatchPayerPayeeRelationship *PatchPayerPayeeRelationship
	PatchPersonBusinessOwnerRelationship *PatchPersonBusinessOwnerRelationship
	PatchPersonBusinessRelationship *PatchPersonBusinessRelationship
}

// PatchBusinessBusinessOwnerRelationshipAsPatchRelationshipIn is a convenience function that returns PatchBusinessBusinessOwnerRelationship wrapped in PatchRelationshipIn
func PatchBusinessBusinessOwnerRelationshipAsPatchRelationshipIn(v *PatchBusinessBusinessOwnerRelationship) PatchRelationshipIn {
	return PatchRelationshipIn{
		PatchBusinessBusinessOwnerRelationship: v,
	}
}

// PatchPayerPayeeRelationshipAsPatchRelationshipIn is a convenience function that returns PatchPayerPayeeRelationship wrapped in PatchRelationshipIn
func PatchPayerPayeeRelationshipAsPatchRelationshipIn(v *PatchPayerPayeeRelationship) PatchRelationshipIn {
	return PatchRelationshipIn{
		PatchPayerPayeeRelationship: v,
	}
}

// PatchPersonBusinessOwnerRelationshipAsPatchRelationshipIn is a convenience function that returns PatchPersonBusinessOwnerRelationship wrapped in PatchRelationshipIn
func PatchPersonBusinessOwnerRelationshipAsPatchRelationshipIn(v *PatchPersonBusinessOwnerRelationship) PatchRelationshipIn {
	return PatchRelationshipIn{
		PatchPersonBusinessOwnerRelationship: v,
	}
}

// PatchPersonBusinessRelationshipAsPatchRelationshipIn is a convenience function that returns PatchPersonBusinessRelationship wrapped in PatchRelationshipIn
func PatchPersonBusinessRelationshipAsPatchRelationshipIn(v *PatchPersonBusinessRelationship) PatchRelationshipIn {
	return PatchRelationshipIn{
		PatchPersonBusinessRelationship: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchRelationshipIn) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BENEFICIAL_OWNER_OF'
	if jsonDict["relationship_type"] == "BENEFICIAL_OWNER_OF" {
		// try to unmarshal JSON data into PatchPersonBusinessOwnerRelationship
		err = json.Unmarshal(data, &dst.PatchPersonBusinessOwnerRelationship)
		if err == nil {
			return nil // data stored in dst.PatchPersonBusinessOwnerRelationship, return on the first match
		} else {
			dst.PatchPersonBusinessOwnerRelationship = nil
			return fmt.Errorf("failed to unmarshal PatchRelationshipIn as PatchPersonBusinessOwnerRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MANAGING_PERSON_OF'
	if jsonDict["relationship_type"] == "MANAGING_PERSON_OF" {
		// try to unmarshal JSON data into PatchPersonBusinessRelationship
		err = json.Unmarshal(data, &dst.PatchPersonBusinessRelationship)
		if err == nil {
			return nil // data stored in dst.PatchPersonBusinessRelationship, return on the first match
		} else {
			dst.PatchPersonBusinessRelationship = nil
			return fmt.Errorf("failed to unmarshal PatchRelationshipIn as PatchPersonBusinessRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OWNER_OF'
	if jsonDict["relationship_type"] == "OWNER_OF" {
		// try to unmarshal JSON data into PatchBusinessBusinessOwnerRelationship
		err = json.Unmarshal(data, &dst.PatchBusinessBusinessOwnerRelationship)
		if err == nil {
			return nil // data stored in dst.PatchBusinessBusinessOwnerRelationship, return on the first match
		} else {
			dst.PatchBusinessBusinessOwnerRelationship = nil
			return fmt.Errorf("failed to unmarshal PatchRelationshipIn as PatchBusinessBusinessOwnerRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PAYER_PAYEE'
	if jsonDict["relationship_type"] == "PAYER_PAYEE" {
		// try to unmarshal JSON data into PatchPayerPayeeRelationship
		err = json.Unmarshal(data, &dst.PatchPayerPayeeRelationship)
		if err == nil {
			return nil // data stored in dst.PatchPayerPayeeRelationship, return on the first match
		} else {
			dst.PatchPayerPayeeRelationship = nil
			return fmt.Errorf("failed to unmarshal PatchRelationshipIn as PatchPayerPayeeRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'patch_business_business_owner_relationship'
	if jsonDict["relationship_type"] == "patch_business_business_owner_relationship" {
		// try to unmarshal JSON data into PatchBusinessBusinessOwnerRelationship
		err = json.Unmarshal(data, &dst.PatchBusinessBusinessOwnerRelationship)
		if err == nil {
			return nil // data stored in dst.PatchBusinessBusinessOwnerRelationship, return on the first match
		} else {
			dst.PatchBusinessBusinessOwnerRelationship = nil
			return fmt.Errorf("failed to unmarshal PatchRelationshipIn as PatchBusinessBusinessOwnerRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'patch_payer_payee_relationship'
	if jsonDict["relationship_type"] == "patch_payer_payee_relationship" {
		// try to unmarshal JSON data into PatchPayerPayeeRelationship
		err = json.Unmarshal(data, &dst.PatchPayerPayeeRelationship)
		if err == nil {
			return nil // data stored in dst.PatchPayerPayeeRelationship, return on the first match
		} else {
			dst.PatchPayerPayeeRelationship = nil
			return fmt.Errorf("failed to unmarshal PatchRelationshipIn as PatchPayerPayeeRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'patch_person_business_owner_relationship'
	if jsonDict["relationship_type"] == "patch_person_business_owner_relationship" {
		// try to unmarshal JSON data into PatchPersonBusinessOwnerRelationship
		err = json.Unmarshal(data, &dst.PatchPersonBusinessOwnerRelationship)
		if err == nil {
			return nil // data stored in dst.PatchPersonBusinessOwnerRelationship, return on the first match
		} else {
			dst.PatchPersonBusinessOwnerRelationship = nil
			return fmt.Errorf("failed to unmarshal PatchRelationshipIn as PatchPersonBusinessOwnerRelationship: %s", err.Error())
		}
	}

	// check if the discriminator value is 'patch_person_business_relationship'
	if jsonDict["relationship_type"] == "patch_person_business_relationship" {
		// try to unmarshal JSON data into PatchPersonBusinessRelationship
		err = json.Unmarshal(data, &dst.PatchPersonBusinessRelationship)
		if err == nil {
			return nil // data stored in dst.PatchPersonBusinessRelationship, return on the first match
		} else {
			dst.PatchPersonBusinessRelationship = nil
			return fmt.Errorf("failed to unmarshal PatchRelationshipIn as PatchPersonBusinessRelationship: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchRelationshipIn) MarshalJSON() ([]byte, error) {
	if src.PatchBusinessBusinessOwnerRelationship != nil {
		return json.Marshal(&src.PatchBusinessBusinessOwnerRelationship)
	}

	if src.PatchPayerPayeeRelationship != nil {
		return json.Marshal(&src.PatchPayerPayeeRelationship)
	}

	if src.PatchPersonBusinessOwnerRelationship != nil {
		return json.Marshal(&src.PatchPersonBusinessOwnerRelationship)
	}

	if src.PatchPersonBusinessRelationship != nil {
		return json.Marshal(&src.PatchPersonBusinessRelationship)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchRelationshipIn) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PatchBusinessBusinessOwnerRelationship != nil {
		return obj.PatchBusinessBusinessOwnerRelationship
	}

	if obj.PatchPayerPayeeRelationship != nil {
		return obj.PatchPayerPayeeRelationship
	}

	if obj.PatchPersonBusinessOwnerRelationship != nil {
		return obj.PatchPersonBusinessOwnerRelationship
	}

	if obj.PatchPersonBusinessRelationship != nil {
		return obj.PatchPersonBusinessRelationship
	}

	// all schemas are nil
	return nil
}

type NullablePatchRelationshipIn struct {
	value *PatchRelationshipIn
	isSet bool
}

func (v NullablePatchRelationshipIn) Get() *PatchRelationshipIn {
	return v.value
}

func (v *NullablePatchRelationshipIn) Set(val *PatchRelationshipIn) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchRelationshipIn) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchRelationshipIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchRelationshipIn(val *PatchRelationshipIn) *NullablePatchRelationshipIn {
	return &NullablePatchRelationshipIn{value: val, isSet: true}
}

func (v NullablePatchRelationshipIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchRelationshipIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


