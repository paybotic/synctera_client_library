/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// checks if the BaseDisclosure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseDisclosure{}

// BaseDisclosure struct for BaseDisclosure
type BaseDisclosure struct {
	// Unique ID for the person acknowledging the disclosure. Applicable for disclosures where the person acknowledging is different from the subject of the disclosure. Required for OWNER_CERTIFICATION disclosures.
	AcknowledgingPersonId *string `json:"acknowledging_person_id,omitempty"`
	// Unique ID for the business. Exactly one of `business_id` or `person_id` must be set.
	BusinessId *string `json:"business_id,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Date and time the disclosure was made.
	DisclosureDate *time.Time `json:"disclosure_date,omitempty"`
	// Describes how the disclosure was shown and what the user did as a result. One of the following: * `DISPLAYED` —     The document was made visible to the user,     but they did not interact with it.  * `VIEWED` —     The document was made visible to the user,     and they interacted enough to see the whole document (e.g. scrolled to the bottom).  * `ACKNOWLEDGED` —     The document was made visible to the user,     and they took positive action to confirm that they have read and accepted the document.
	EventType *string `json:"event_type,omitempty"`
	// The unique identifier for this resource.
	Id *string `json:"id,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Unique ID for the person. Exactly one of `person_id` or `business_id` must be set.
	PersonId *string `json:"person_id,omitempty"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant *string         `json:"tenant,omitempty"`
	Type   *DisclosureType `json:"type,omitempty"`
	// Version of the disclosure document.
	Version *string `json:"version,omitempty" validate:"regexp=^v?[0-9]+\\\\.[0-9]+$"`
}

// NewBaseDisclosure instantiates a new BaseDisclosure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseDisclosure() *BaseDisclosure {
	this := BaseDisclosure{}
	return &this
}

// NewBaseDisclosureWithDefaults instantiates a new BaseDisclosure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseDisclosureWithDefaults() *BaseDisclosure {
	this := BaseDisclosure{}
	return &this
}

// GetAcknowledgingPersonId returns the AcknowledgingPersonId field value if set, zero value otherwise.
func (o *BaseDisclosure) GetAcknowledgingPersonId() string {
	if o == nil || IsNil(o.AcknowledgingPersonId) {
		var ret string
		return ret
	}
	return *o.AcknowledgingPersonId
}

// GetAcknowledgingPersonIdOk returns a tuple with the AcknowledgingPersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetAcknowledgingPersonIdOk() (*string, bool) {
	if o == nil || IsNil(o.AcknowledgingPersonId) {
		return nil, false
	}
	return o.AcknowledgingPersonId, true
}

// HasAcknowledgingPersonId returns a boolean if a field has been set.
func (o *BaseDisclosure) HasAcknowledgingPersonId() bool {
	if o != nil && !IsNil(o.AcknowledgingPersonId) {
		return true
	}

	return false
}

// SetAcknowledgingPersonId gets a reference to the given string and assigns it to the AcknowledgingPersonId field.
func (o *BaseDisclosure) SetAcknowledgingPersonId(v string) {
	o.AcknowledgingPersonId = &v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *BaseDisclosure) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *BaseDisclosure) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *BaseDisclosure) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BaseDisclosure) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BaseDisclosure) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BaseDisclosure) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDisclosureDate returns the DisclosureDate field value if set, zero value otherwise.
func (o *BaseDisclosure) GetDisclosureDate() time.Time {
	if o == nil || IsNil(o.DisclosureDate) {
		var ret time.Time
		return ret
	}
	return *o.DisclosureDate
}

// GetDisclosureDateOk returns a tuple with the DisclosureDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetDisclosureDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DisclosureDate) {
		return nil, false
	}
	return o.DisclosureDate, true
}

// HasDisclosureDate returns a boolean if a field has been set.
func (o *BaseDisclosure) HasDisclosureDate() bool {
	if o != nil && !IsNil(o.DisclosureDate) {
		return true
	}

	return false
}

// SetDisclosureDate gets a reference to the given time.Time and assigns it to the DisclosureDate field.
func (o *BaseDisclosure) SetDisclosureDate(v time.Time) {
	o.DisclosureDate = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *BaseDisclosure) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *BaseDisclosure) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *BaseDisclosure) SetEventType(v string) {
	o.EventType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseDisclosure) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseDisclosure) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseDisclosure) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *BaseDisclosure) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *BaseDisclosure) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *BaseDisclosure) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BaseDisclosure) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BaseDisclosure) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *BaseDisclosure) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *BaseDisclosure) GetPersonId() string {
	if o == nil || IsNil(o.PersonId) {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetPersonIdOk() (*string, bool) {
	if o == nil || IsNil(o.PersonId) {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *BaseDisclosure) HasPersonId() bool {
	if o != nil && !IsNil(o.PersonId) {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
func (o *BaseDisclosure) SetPersonId(v string) {
	o.PersonId = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *BaseDisclosure) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *BaseDisclosure) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *BaseDisclosure) SetTenant(v string) {
	o.Tenant = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BaseDisclosure) GetType() DisclosureType {
	if o == nil || IsNil(o.Type) {
		var ret DisclosureType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetTypeOk() (*DisclosureType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BaseDisclosure) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DisclosureType and assigns it to the Type field.
func (o *BaseDisclosure) SetType(v DisclosureType) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BaseDisclosure) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDisclosure) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BaseDisclosure) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BaseDisclosure) SetVersion(v string) {
	o.Version = &v
}

func (o BaseDisclosure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseDisclosure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcknowledgingPersonId) {
		toSerialize["acknowledging_person_id"] = o.AcknowledgingPersonId
	}
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.DisclosureDate) {
		toSerialize["disclosure_date"] = o.DisclosureDate
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
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
	if !IsNil(o.PersonId) {
		toSerialize["person_id"] = o.PersonId
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableBaseDisclosure struct {
	value *BaseDisclosure
	isSet bool
}

func (v NullableBaseDisclosure) Get() *BaseDisclosure {
	return v.value
}

func (v *NullableBaseDisclosure) Set(val *BaseDisclosure) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseDisclosure) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseDisclosure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseDisclosure(val *BaseDisclosure) *NullableBaseDisclosure {
	return &NullableBaseDisclosure{value: val, isSet: true}
}

func (v NullableBaseDisclosure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseDisclosure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
