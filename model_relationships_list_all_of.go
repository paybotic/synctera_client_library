/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RelationshipsListAllOf struct for RelationshipsListAllOf
type RelationshipsListAllOf struct {
	// Array of business/person relationships.
	Relationships []RelationshipIn `json:"relationships"`
}

// NewRelationshipsListAllOf instantiates a new RelationshipsListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipsListAllOf(relationships []RelationshipIn) *RelationshipsListAllOf {
	this := RelationshipsListAllOf{}
	this.Relationships = relationships
	return &this
}

// NewRelationshipsListAllOfWithDefaults instantiates a new RelationshipsListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipsListAllOfWithDefaults() *RelationshipsListAllOf {
	this := RelationshipsListAllOf{}
	return &this
}

// GetRelationships returns the Relationships field value
func (o *RelationshipsListAllOf) GetRelationships() []RelationshipIn {
	if o == nil {
		var ret []RelationshipIn
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *RelationshipsListAllOf) GetRelationshipsOk() ([]RelationshipIn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relationships, true
}

// SetRelationships sets field value
func (o *RelationshipsListAllOf) SetRelationships(v []RelationshipIn) {
	o.Relationships = v
}

func (o RelationshipsListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableRelationshipsListAllOf struct {
	value *RelationshipsListAllOf
	isSet bool
}

func (v NullableRelationshipsListAllOf) Get() *RelationshipsListAllOf {
	return v.value
}

func (v *NullableRelationshipsListAllOf) Set(val *RelationshipsListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipsListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipsListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipsListAllOf(val *RelationshipsListAllOf) *NullableRelationshipsListAllOf {
	return &NullableRelationshipsListAllOf{value: val, isSet: true}
}

func (v NullableRelationshipsListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipsListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
