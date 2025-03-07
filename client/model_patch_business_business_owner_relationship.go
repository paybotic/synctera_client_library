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
	"time"
)

// checks if the PatchBusinessBusinessOwnerRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchBusinessBusinessOwnerRelationship{}

// PatchBusinessBusinessOwnerRelationship Denotes the relationship between specified businesses.
type PatchBusinessBusinessOwnerRelationship struct {
	AdditionalData *AdditionalOwnerData `json:"additional_data,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Unique ID for the subject business.
	FromBusinessId *string `json:"from_business_id,omitempty"`
	// Relationship unique identifier.
	Id *string `json:"id,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	RelationshipType RelationshipTypes      `json:"relationship_type"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant *string `json:"tenant,omitempty"`
	// Unique ID for the related business.
	ToBusinessId         *string `json:"to_business_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchBusinessBusinessOwnerRelationship PatchBusinessBusinessOwnerRelationship

// NewPatchBusinessBusinessOwnerRelationship instantiates a new PatchBusinessBusinessOwnerRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchBusinessBusinessOwnerRelationship(relationshipType RelationshipTypes) *PatchBusinessBusinessOwnerRelationship {
	this := PatchBusinessBusinessOwnerRelationship{}
	this.RelationshipType = relationshipType
	return &this
}

// NewPatchBusinessBusinessOwnerRelationshipWithDefaults instantiates a new PatchBusinessBusinessOwnerRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchBusinessBusinessOwnerRelationshipWithDefaults() *PatchBusinessBusinessOwnerRelationship {
	this := PatchBusinessBusinessOwnerRelationship{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetAdditionalData() AdditionalOwnerData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret AdditionalOwnerData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetAdditionalDataOk() (*AdditionalOwnerData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given AdditionalOwnerData and assigns it to the AdditionalData field.
func (o *PatchBusinessBusinessOwnerRelationship) SetAdditionalData(v AdditionalOwnerData) {
	o.AdditionalData = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PatchBusinessBusinessOwnerRelationship) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetFromBusinessId returns the FromBusinessId field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetFromBusinessId() string {
	if o == nil || IsNil(o.FromBusinessId) {
		var ret string
		return ret
	}
	return *o.FromBusinessId
}

// GetFromBusinessIdOk returns a tuple with the FromBusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetFromBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.FromBusinessId) {
		return nil, false
	}
	return o.FromBusinessId, true
}

// HasFromBusinessId returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasFromBusinessId() bool {
	if o != nil && !IsNil(o.FromBusinessId) {
		return true
	}

	return false
}

// SetFromBusinessId gets a reference to the given string and assigns it to the FromBusinessId field.
func (o *PatchBusinessBusinessOwnerRelationship) SetFromBusinessId(v string) {
	o.FromBusinessId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchBusinessBusinessOwnerRelationship) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *PatchBusinessBusinessOwnerRelationship) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchBusinessBusinessOwnerRelationship) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetRelationshipType returns the RelationshipType field value
func (o *PatchBusinessBusinessOwnerRelationship) GetRelationshipType() RelationshipTypes {
	if o == nil {
		var ret RelationshipTypes
		return ret
	}

	return o.RelationshipType
}

// GetRelationshipTypeOk returns a tuple with the RelationshipType field value
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetRelationshipTypeOk() (*RelationshipTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationshipType, true
}

// SetRelationshipType sets field value
func (o *PatchBusinessBusinessOwnerRelationship) SetRelationshipType(v RelationshipTypes) {
	o.RelationshipType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *PatchBusinessBusinessOwnerRelationship) SetTenant(v string) {
	o.Tenant = &v
}

// GetToBusinessId returns the ToBusinessId field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetToBusinessId() string {
	if o == nil || IsNil(o.ToBusinessId) {
		var ret string
		return ret
	}
	return *o.ToBusinessId
}

// GetToBusinessIdOk returns a tuple with the ToBusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetToBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.ToBusinessId) {
		return nil, false
	}
	return o.ToBusinessId, true
}

// HasToBusinessId returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasToBusinessId() bool {
	if o != nil && !IsNil(o.ToBusinessId) {
		return true
	}

	return false
}

// SetToBusinessId gets a reference to the given string and assigns it to the ToBusinessId field.
func (o *PatchBusinessBusinessOwnerRelationship) SetToBusinessId(v string) {
	o.ToBusinessId = &v
}

func (o PatchBusinessBusinessOwnerRelationship) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchBusinessBusinessOwnerRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalData) {
		toSerialize["additional_data"] = o.AdditionalData
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.FromBusinessId) {
		toSerialize["from_business_id"] = o.FromBusinessId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["relationship_type"] = o.RelationshipType
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.ToBusinessId) {
		toSerialize["to_business_id"] = o.ToBusinessId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchBusinessBusinessOwnerRelationship) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"relationship_type",
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

	varPatchBusinessBusinessOwnerRelationship := _PatchBusinessBusinessOwnerRelationship{}

	err = json.Unmarshal(data, &varPatchBusinessBusinessOwnerRelationship)

	if err != nil {
		return err
	}

	*o = PatchBusinessBusinessOwnerRelationship(varPatchBusinessBusinessOwnerRelationship)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_data")
		delete(additionalProperties, "creation_time")
		delete(additionalProperties, "from_business_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "last_updated_time")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "relationship_type")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "to_business_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchBusinessBusinessOwnerRelationship struct {
	value *PatchBusinessBusinessOwnerRelationship
	isSet bool
}

func (v NullablePatchBusinessBusinessOwnerRelationship) Get() *PatchBusinessBusinessOwnerRelationship {
	return v.value
}

func (v *NullablePatchBusinessBusinessOwnerRelationship) Set(val *PatchBusinessBusinessOwnerRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBusinessBusinessOwnerRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBusinessBusinessOwnerRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBusinessBusinessOwnerRelationship(val *PatchBusinessBusinessOwnerRelationship) *NullablePatchBusinessBusinessOwnerRelationship {
	return &NullablePatchBusinessBusinessOwnerRelationship{value: val, isSet: true}
}

func (v NullablePatchBusinessBusinessOwnerRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBusinessBusinessOwnerRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
