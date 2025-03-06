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

// checks if the DocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentResponse{}

// DocumentResponse struct for DocumentResponse
type DocumentResponse struct {
	// All document versions
	AvailableVersions []int32 `json:"available_versions,omitempty"`
	// Metadata of all document versions
	AvailableVersionsInfo []DocumentVersionInfo `json:"available_versions_info,omitempty"`
	// The ID of the batch that the document belongs to
	BatchId *string `json:"batch_id,omitempty"`
	// The date and time the resource was created
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// An explanation why the file was deleted. You must set a document's deletion_reason before deleting it.
	DeletionReason *string `json:"deletion_reason,omitempty"`
	// A description of the document
	Description *string             `json:"description,omitempty"`
	Encryption  *DocumentEncryption `json:"encryption,omitempty"`
	// The file name of the document
	FileName *string `json:"file_name,omitempty"`
	// The unique identifier for this resource
	Id *string `json:"id,omitempty"`
	// The date and time the resource was last updated
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// A user-friendly name for the document
	Name *string `json:"name,omitempty"`
	// The ID of the resource related to the document
	RelatedResourceId   *string              `json:"related_resource_id,omitempty"`
	RelatedResourceType *RelatedResourceType `json:"related_resource_type,omitempty"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant *string       `json:"tenant,omitempty"`
	Type   *DocumentType `json:"type,omitempty"`
	// Positive integer representing the version of the document
	Version *int32 `json:"version,omitempty"`
}

// NewDocumentResponse instantiates a new DocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentResponse() *DocumentResponse {
	this := DocumentResponse{}
	return &this
}

// NewDocumentResponseWithDefaults instantiates a new DocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentResponseWithDefaults() *DocumentResponse {
	this := DocumentResponse{}
	return &this
}

// GetAvailableVersions returns the AvailableVersions field value if set, zero value otherwise.
func (o *DocumentResponse) GetAvailableVersions() []int32 {
	if o == nil || IsNil(o.AvailableVersions) {
		var ret []int32
		return ret
	}
	return o.AvailableVersions
}

// GetAvailableVersionsOk returns a tuple with the AvailableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetAvailableVersionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.AvailableVersions) {
		return nil, false
	}
	return o.AvailableVersions, true
}

// HasAvailableVersions returns a boolean if a field has been set.
func (o *DocumentResponse) HasAvailableVersions() bool {
	if o != nil && !IsNil(o.AvailableVersions) {
		return true
	}

	return false
}

// SetAvailableVersions gets a reference to the given []int32 and assigns it to the AvailableVersions field.
func (o *DocumentResponse) SetAvailableVersions(v []int32) {
	o.AvailableVersions = v
}

// GetAvailableVersionsInfo returns the AvailableVersionsInfo field value if set, zero value otherwise.
func (o *DocumentResponse) GetAvailableVersionsInfo() []DocumentVersionInfo {
	if o == nil || IsNil(o.AvailableVersionsInfo) {
		var ret []DocumentVersionInfo
		return ret
	}
	return o.AvailableVersionsInfo
}

// GetAvailableVersionsInfoOk returns a tuple with the AvailableVersionsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetAvailableVersionsInfoOk() ([]DocumentVersionInfo, bool) {
	if o == nil || IsNil(o.AvailableVersionsInfo) {
		return nil, false
	}
	return o.AvailableVersionsInfo, true
}

// HasAvailableVersionsInfo returns a boolean if a field has been set.
func (o *DocumentResponse) HasAvailableVersionsInfo() bool {
	if o != nil && !IsNil(o.AvailableVersionsInfo) {
		return true
	}

	return false
}

// SetAvailableVersionsInfo gets a reference to the given []DocumentVersionInfo and assigns it to the AvailableVersionsInfo field.
func (o *DocumentResponse) SetAvailableVersionsInfo(v []DocumentVersionInfo) {
	o.AvailableVersionsInfo = v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *DocumentResponse) GetBatchId() string {
	if o == nil || IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *DocumentResponse) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *DocumentResponse) SetBatchId(v string) {
	o.BatchId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *DocumentResponse) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *DocumentResponse) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *DocumentResponse) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDeletionReason returns the DeletionReason field value if set, zero value otherwise.
func (o *DocumentResponse) GetDeletionReason() string {
	if o == nil || IsNil(o.DeletionReason) {
		var ret string
		return ret
	}
	return *o.DeletionReason
}

// GetDeletionReasonOk returns a tuple with the DeletionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetDeletionReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DeletionReason) {
		return nil, false
	}
	return o.DeletionReason, true
}

// HasDeletionReason returns a boolean if a field has been set.
func (o *DocumentResponse) HasDeletionReason() bool {
	if o != nil && !IsNil(o.DeletionReason) {
		return true
	}

	return false
}

// SetDeletionReason gets a reference to the given string and assigns it to the DeletionReason field.
func (o *DocumentResponse) SetDeletionReason(v string) {
	o.DeletionReason = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DocumentResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DocumentResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DocumentResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *DocumentResponse) GetEncryption() DocumentEncryption {
	if o == nil || IsNil(o.Encryption) {
		var ret DocumentEncryption
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetEncryptionOk() (*DocumentEncryption, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *DocumentResponse) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given DocumentEncryption and assigns it to the Encryption field.
func (o *DocumentResponse) SetEncryption(v DocumentEncryption) {
	o.Encryption = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *DocumentResponse) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *DocumentResponse) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *DocumentResponse) SetFileName(v string) {
	o.FileName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *DocumentResponse) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *DocumentResponse) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *DocumentResponse) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DocumentResponse) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DocumentResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *DocumentResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DocumentResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DocumentResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DocumentResponse) SetName(v string) {
	o.Name = &v
}

// GetRelatedResourceId returns the RelatedResourceId field value if set, zero value otherwise.
func (o *DocumentResponse) GetRelatedResourceId() string {
	if o == nil || IsNil(o.RelatedResourceId) {
		var ret string
		return ret
	}
	return *o.RelatedResourceId
}

// GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetRelatedResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.RelatedResourceId) {
		return nil, false
	}
	return o.RelatedResourceId, true
}

// HasRelatedResourceId returns a boolean if a field has been set.
func (o *DocumentResponse) HasRelatedResourceId() bool {
	if o != nil && !IsNil(o.RelatedResourceId) {
		return true
	}

	return false
}

// SetRelatedResourceId gets a reference to the given string and assigns it to the RelatedResourceId field.
func (o *DocumentResponse) SetRelatedResourceId(v string) {
	o.RelatedResourceId = &v
}

// GetRelatedResourceType returns the RelatedResourceType field value if set, zero value otherwise.
func (o *DocumentResponse) GetRelatedResourceType() RelatedResourceType {
	if o == nil || IsNil(o.RelatedResourceType) {
		var ret RelatedResourceType
		return ret
	}
	return *o.RelatedResourceType
}

// GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetRelatedResourceTypeOk() (*RelatedResourceType, bool) {
	if o == nil || IsNil(o.RelatedResourceType) {
		return nil, false
	}
	return o.RelatedResourceType, true
}

// HasRelatedResourceType returns a boolean if a field has been set.
func (o *DocumentResponse) HasRelatedResourceType() bool {
	if o != nil && !IsNil(o.RelatedResourceType) {
		return true
	}

	return false
}

// SetRelatedResourceType gets a reference to the given RelatedResourceType and assigns it to the RelatedResourceType field.
func (o *DocumentResponse) SetRelatedResourceType(v RelatedResourceType) {
	o.RelatedResourceType = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *DocumentResponse) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *DocumentResponse) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *DocumentResponse) SetTenant(v string) {
	o.Tenant = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DocumentResponse) GetType() DocumentType {
	if o == nil || IsNil(o.Type) {
		var ret DocumentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetTypeOk() (*DocumentType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DocumentResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DocumentType and assigns it to the Type field.
func (o *DocumentResponse) SetType(v DocumentType) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DocumentResponse) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DocumentResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DocumentResponse) SetVersion(v int32) {
	o.Version = &v
}

func (o DocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailableVersions) {
		toSerialize["available_versions"] = o.AvailableVersions
	}
	if !IsNil(o.AvailableVersionsInfo) {
		toSerialize["available_versions_info"] = o.AvailableVersionsInfo
	}
	if !IsNil(o.BatchId) {
		toSerialize["batch_id"] = o.BatchId
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.DeletionReason) {
		toSerialize["deletion_reason"] = o.DeletionReason
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Encryption) {
		toSerialize["encryption"] = o.Encryption
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RelatedResourceId) {
		toSerialize["related_resource_id"] = o.RelatedResourceId
	}
	if !IsNil(o.RelatedResourceType) {
		toSerialize["related_resource_type"] = o.RelatedResourceType
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

type NullableDocumentResponse struct {
	value *DocumentResponse
	isSet bool
}

func (v NullableDocumentResponse) Get() *DocumentResponse {
	return v.value
}

func (v *NullableDocumentResponse) Set(val *DocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentResponse(val *DocumentResponse) *NullableDocumentResponse {
	return &NullableDocumentResponse{value: val, isSet: true}
}

func (v NullableDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
