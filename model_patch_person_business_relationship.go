/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// PatchPersonBusinessRelationship Denotes the relationship between specified person and business.
type PatchPersonBusinessRelationship struct {
	AdditionalData *AdditionalData `json:"additional_data,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Unique ID for the subject person. 
	FromPersonId *string `json:"from_person_id,omitempty"`
	// Relationship unique identifier.
	Id *string `json:"id,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The relationship type. One of the following: * `BENEFICIAL_OWNER_OF` – a person who directly or indirectly owns a portion of the business. * `MANAGING_PERSON_OF` – a person who is an officer, director, or other notable person of an organization. * `OWNER_OF` – a business with ownership of another business. 
	RelationshipType string `json:"relationship_type"`
	// Unique ID for the related business. 
	ToBusinessId *string `json:"to_business_id,omitempty"`
}

// NewPatchPersonBusinessRelationship instantiates a new PatchPersonBusinessRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPersonBusinessRelationship(relationshipType string) *PatchPersonBusinessRelationship {
	this := PatchPersonBusinessRelationship{}
	this.RelationshipType = relationshipType
	return &this
}

// NewPatchPersonBusinessRelationshipWithDefaults instantiates a new PatchPersonBusinessRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPersonBusinessRelationshipWithDefaults() *PatchPersonBusinessRelationship {
	this := PatchPersonBusinessRelationship{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *PatchPersonBusinessRelationship) GetAdditionalData() AdditionalData {
	if o == nil || o.AdditionalData == nil {
		var ret AdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPersonBusinessRelationship) GetAdditionalDataOk() (*AdditionalData, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *PatchPersonBusinessRelationship) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given AdditionalData and assigns it to the AdditionalData field.
func (o *PatchPersonBusinessRelationship) SetAdditionalData(v AdditionalData) {
	o.AdditionalData = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PatchPersonBusinessRelationship) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPersonBusinessRelationship) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PatchPersonBusinessRelationship) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PatchPersonBusinessRelationship) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetFromPersonId returns the FromPersonId field value if set, zero value otherwise.
func (o *PatchPersonBusinessRelationship) GetFromPersonId() string {
	if o == nil || o.FromPersonId == nil {
		var ret string
		return ret
	}
	return *o.FromPersonId
}

// GetFromPersonIdOk returns a tuple with the FromPersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPersonBusinessRelationship) GetFromPersonIdOk() (*string, bool) {
	if o == nil || o.FromPersonId == nil {
		return nil, false
	}
	return o.FromPersonId, true
}

// HasFromPersonId returns a boolean if a field has been set.
func (o *PatchPersonBusinessRelationship) HasFromPersonId() bool {
	if o != nil && o.FromPersonId != nil {
		return true
	}

	return false
}

// SetFromPersonId gets a reference to the given string and assigns it to the FromPersonId field.
func (o *PatchPersonBusinessRelationship) SetFromPersonId(v string) {
	o.FromPersonId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchPersonBusinessRelationship) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPersonBusinessRelationship) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchPersonBusinessRelationship) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchPersonBusinessRelationship) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *PatchPersonBusinessRelationship) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPersonBusinessRelationship) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *PatchPersonBusinessRelationship) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *PatchPersonBusinessRelationship) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchPersonBusinessRelationship) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPersonBusinessRelationship) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchPersonBusinessRelationship) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchPersonBusinessRelationship) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetRelationshipType returns the RelationshipType field value
func (o *PatchPersonBusinessRelationship) GetRelationshipType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelationshipType
}

// GetRelationshipTypeOk returns a tuple with the RelationshipType field value
// and a boolean to check if the value has been set.
func (o *PatchPersonBusinessRelationship) GetRelationshipTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationshipType, true
}

// SetRelationshipType sets field value
func (o *PatchPersonBusinessRelationship) SetRelationshipType(v string) {
	o.RelationshipType = v
}

// GetToBusinessId returns the ToBusinessId field value if set, zero value otherwise.
func (o *PatchPersonBusinessRelationship) GetToBusinessId() string {
	if o == nil || o.ToBusinessId == nil {
		var ret string
		return ret
	}
	return *o.ToBusinessId
}

// GetToBusinessIdOk returns a tuple with the ToBusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPersonBusinessRelationship) GetToBusinessIdOk() (*string, bool) {
	if o == nil || o.ToBusinessId == nil {
		return nil, false
	}
	return o.ToBusinessId, true
}

// HasToBusinessId returns a boolean if a field has been set.
func (o *PatchPersonBusinessRelationship) HasToBusinessId() bool {
	if o != nil && o.ToBusinessId != nil {
		return true
	}

	return false
}

// SetToBusinessId gets a reference to the given string and assigns it to the ToBusinessId field.
func (o *PatchPersonBusinessRelationship) SetToBusinessId(v string) {
	o.ToBusinessId = &v
}

func (o PatchPersonBusinessRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalData != nil {
		toSerialize["additional_data"] = o.AdditionalData
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.FromPersonId != nil {
		toSerialize["from_person_id"] = o.FromPersonId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["relationship_type"] = o.RelationshipType
	}
	if o.ToBusinessId != nil {
		toSerialize["to_business_id"] = o.ToBusinessId
	}
	return json.Marshal(toSerialize)
}

type NullablePatchPersonBusinessRelationship struct {
	value *PatchPersonBusinessRelationship
	isSet bool
}

func (v NullablePatchPersonBusinessRelationship) Get() *PatchPersonBusinessRelationship {
	return v.value
}

func (v *NullablePatchPersonBusinessRelationship) Set(val *PatchPersonBusinessRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPersonBusinessRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPersonBusinessRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPersonBusinessRelationship(val *PatchPersonBusinessRelationship) *NullablePatchPersonBusinessRelationship {
	return &NullablePatchPersonBusinessRelationship{value: val, isSet: true}
}

func (v NullablePatchPersonBusinessRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPersonBusinessRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


