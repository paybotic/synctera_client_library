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

// checks if the PersonBusinessRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonBusinessRelationship{}

// PersonBusinessRelationship Denotes the relationship between specified person and business.
type PersonBusinessRelationship struct {
	AdditionalData AdditionalData `json:"additional_data"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Unique ID for the subject person.
	FromPersonId string `json:"from_person_id"`
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
	ToBusinessId         string `json:"to_business_id"`
	AdditionalProperties map[string]interface{}
}

type _PersonBusinessRelationship PersonBusinessRelationship

// NewPersonBusinessRelationship instantiates a new PersonBusinessRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonBusinessRelationship(additionalData AdditionalData, fromPersonId string, relationshipType RelationshipTypes, toBusinessId string) *PersonBusinessRelationship {
	this := PersonBusinessRelationship{}
	this.AdditionalData = additionalData
	this.FromPersonId = fromPersonId
	this.RelationshipType = relationshipType
	this.ToBusinessId = toBusinessId
	return &this
}

// NewPersonBusinessRelationshipWithDefaults instantiates a new PersonBusinessRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonBusinessRelationshipWithDefaults() *PersonBusinessRelationship {
	this := PersonBusinessRelationship{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value
func (o *PersonBusinessRelationship) GetAdditionalData() AdditionalData {
	if o == nil {
		var ret AdditionalData
		return ret
	}

	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value
// and a boolean to check if the value has been set.
func (o *PersonBusinessRelationship) GetAdditionalDataOk() (*AdditionalData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalData, true
}

// SetAdditionalData sets field value
func (o *PersonBusinessRelationship) SetAdditionalData(v AdditionalData) {
	o.AdditionalData = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PersonBusinessRelationship) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonBusinessRelationship) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PersonBusinessRelationship) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PersonBusinessRelationship) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetFromPersonId returns the FromPersonId field value
func (o *PersonBusinessRelationship) GetFromPersonId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromPersonId
}

// GetFromPersonIdOk returns a tuple with the FromPersonId field value
// and a boolean to check if the value has been set.
func (o *PersonBusinessRelationship) GetFromPersonIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromPersonId, true
}

// SetFromPersonId sets field value
func (o *PersonBusinessRelationship) SetFromPersonId(v string) {
	o.FromPersonId = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PersonBusinessRelationship) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonBusinessRelationship) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PersonBusinessRelationship) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PersonBusinessRelationship) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *PersonBusinessRelationship) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonBusinessRelationship) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *PersonBusinessRelationship) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *PersonBusinessRelationship) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PersonBusinessRelationship) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonBusinessRelationship) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PersonBusinessRelationship) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PersonBusinessRelationship) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetRelationshipType returns the RelationshipType field value
func (o *PersonBusinessRelationship) GetRelationshipType() RelationshipTypes {
	if o == nil {
		var ret RelationshipTypes
		return ret
	}

	return o.RelationshipType
}

// GetRelationshipTypeOk returns a tuple with the RelationshipType field value
// and a boolean to check if the value has been set.
func (o *PersonBusinessRelationship) GetRelationshipTypeOk() (*RelationshipTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationshipType, true
}

// SetRelationshipType sets field value
func (o *PersonBusinessRelationship) SetRelationshipType(v RelationshipTypes) {
	o.RelationshipType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *PersonBusinessRelationship) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonBusinessRelationship) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *PersonBusinessRelationship) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *PersonBusinessRelationship) SetTenant(v string) {
	o.Tenant = &v
}

// GetToBusinessId returns the ToBusinessId field value
func (o *PersonBusinessRelationship) GetToBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToBusinessId
}

// GetToBusinessIdOk returns a tuple with the ToBusinessId field value
// and a boolean to check if the value has been set.
func (o *PersonBusinessRelationship) GetToBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToBusinessId, true
}

// SetToBusinessId sets field value
func (o *PersonBusinessRelationship) SetToBusinessId(v string) {
	o.ToBusinessId = v
}

func (o PersonBusinessRelationship) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonBusinessRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["additional_data"] = o.AdditionalData
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	toSerialize["from_person_id"] = o.FromPersonId
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
	toSerialize["to_business_id"] = o.ToBusinessId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PersonBusinessRelationship) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"additional_data",
		"from_person_id",
		"relationship_type",
		"to_business_id",
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

	varPersonBusinessRelationship := _PersonBusinessRelationship{}

	err = json.Unmarshal(data, &varPersonBusinessRelationship)

	if err != nil {
		return err
	}

	*o = PersonBusinessRelationship(varPersonBusinessRelationship)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_data")
		delete(additionalProperties, "creation_time")
		delete(additionalProperties, "from_person_id")
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

type NullablePersonBusinessRelationship struct {
	value *PersonBusinessRelationship
	isSet bool
}

func (v NullablePersonBusinessRelationship) Get() *PersonBusinessRelationship {
	return v.value
}

func (v *NullablePersonBusinessRelationship) Set(val *PersonBusinessRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonBusinessRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonBusinessRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonBusinessRelationship(val *PersonBusinessRelationship) *NullablePersonBusinessRelationship {
	return &NullablePersonBusinessRelationship{value: val, isSet: true}
}

func (v NullablePersonBusinessRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonBusinessRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
