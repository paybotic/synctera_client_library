/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RelationshipsList struct for RelationshipsList
type RelationshipsList struct {
	// Array of business/person relationships.
	Relationships []RelationshipIn `json:"relationships"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewRelationshipsList instantiates a new RelationshipsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipsList(relationships []RelationshipIn) *RelationshipsList {
	this := RelationshipsList{}
	this.Relationships = relationships
	return &this
}

// NewRelationshipsListWithDefaults instantiates a new RelationshipsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipsListWithDefaults() *RelationshipsList {
	this := RelationshipsList{}
	return &this
}

// GetRelationships returns the Relationships field value
func (o *RelationshipsList) GetRelationships() []RelationshipIn {
	if o == nil {
		var ret []RelationshipIn
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *RelationshipsList) GetRelationshipsOk() ([]RelationshipIn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relationships, true
}

// SetRelationships sets field value
func (o *RelationshipsList) SetRelationships(v []RelationshipIn) {
	o.Relationships = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *RelationshipsList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipsList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *RelationshipsList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *RelationshipsList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o RelationshipsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["relationships"] = o.Relationships
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableRelationshipsList struct {
	value *RelationshipsList
	isSet bool
}

func (v NullableRelationshipsList) Get() *RelationshipsList {
	return v.value
}

func (v *NullableRelationshipsList) Set(val *RelationshipsList) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipsList) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipsList(val *RelationshipsList) *NullableRelationshipsList {
	return &NullableRelationshipsList{value: val, isSet: true}
}

func (v NullableRelationshipsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


