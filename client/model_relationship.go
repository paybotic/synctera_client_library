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

// checks if the Relationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Relationship{}

// Relationship struct for Relationship
type Relationship struct {
	// ID of related entity
	Id                   string           `json:"id"`
	RelationshipRole     RelationshipRole `json:"relationship_role"`
	AdditionalProperties map[string]interface{}
}

type _Relationship Relationship

// NewRelationship instantiates a new Relationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationship(id string, relationshipRole RelationshipRole) *Relationship {
	this := Relationship{}
	this.Id = id
	this.RelationshipRole = relationshipRole
	return &this
}

// NewRelationshipWithDefaults instantiates a new Relationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWithDefaults() *Relationship {
	this := Relationship{}
	return &this
}

// GetId returns the Id field value
func (o *Relationship) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Relationship) SetId(v string) {
	o.Id = v
}

// GetRelationshipRole returns the RelationshipRole field value
func (o *Relationship) GetRelationshipRole() RelationshipRole {
	if o == nil {
		var ret RelationshipRole
		return ret
	}

	return o.RelationshipRole
}

// GetRelationshipRoleOk returns a tuple with the RelationshipRole field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetRelationshipRoleOk() (*RelationshipRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationshipRole, true
}

// SetRelationshipRole sets field value
func (o *Relationship) SetRelationshipRole(v RelationshipRole) {
	o.RelationshipRole = v
}

func (o Relationship) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Relationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["relationship_role"] = o.RelationshipRole

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Relationship) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"relationship_role",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRelationship := _Relationship{}

	err = json.Unmarshal(data, &varRelationship)

	if err != nil {
		return err
	}

	*o = Relationship(varRelationship)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "relationship_role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRelationship struct {
	value *Relationship
	isSet bool
}

func (v NullableRelationship) Get() *Relationship {
	return v.value
}

func (v *NullableRelationship) Set(val *Relationship) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationship(val *Relationship) *NullableRelationship {
	return &NullableRelationship{value: val, isSet: true}
}

func (v NullableRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
