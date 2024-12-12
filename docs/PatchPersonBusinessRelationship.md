# PatchPersonBusinessRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to [**AdditionalData**](AdditionalData.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**FromPersonId** | Pointer to **string** | Unique ID for the subject person.  | [optional] 
**Id** | Pointer to **string** | Relationship unique identifier. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**RelationshipType** | [**RelationshipTypes**](RelationshipTypes.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**ToBusinessId** | Pointer to **string** | Unique ID for the related business.  | [optional] 

## Methods

### NewPatchPersonBusinessRelationship

`func NewPatchPersonBusinessRelationship(relationshipType RelationshipTypes, ) *PatchPersonBusinessRelationship`

NewPatchPersonBusinessRelationship instantiates a new PatchPersonBusinessRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPersonBusinessRelationshipWithDefaults

`func NewPatchPersonBusinessRelationshipWithDefaults() *PatchPersonBusinessRelationship`

NewPatchPersonBusinessRelationshipWithDefaults instantiates a new PatchPersonBusinessRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PatchPersonBusinessRelationship) GetAdditionalData() AdditionalData`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PatchPersonBusinessRelationship) GetAdditionalDataOk() (*AdditionalData, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PatchPersonBusinessRelationship) SetAdditionalData(v AdditionalData)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PatchPersonBusinessRelationship) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetCreationTime

`func (o *PatchPersonBusinessRelationship) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PatchPersonBusinessRelationship) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PatchPersonBusinessRelationship) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PatchPersonBusinessRelationship) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetFromPersonId

`func (o *PatchPersonBusinessRelationship) GetFromPersonId() string`

GetFromPersonId returns the FromPersonId field if non-nil, zero value otherwise.

### GetFromPersonIdOk

`func (o *PatchPersonBusinessRelationship) GetFromPersonIdOk() (*string, bool)`

GetFromPersonIdOk returns a tuple with the FromPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPersonId

`func (o *PatchPersonBusinessRelationship) SetFromPersonId(v string)`

SetFromPersonId sets FromPersonId field to given value.

### HasFromPersonId

`func (o *PatchPersonBusinessRelationship) HasFromPersonId() bool`

HasFromPersonId returns a boolean if a field has been set.

### GetId

`func (o *PatchPersonBusinessRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchPersonBusinessRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchPersonBusinessRelationship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchPersonBusinessRelationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *PatchPersonBusinessRelationship) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *PatchPersonBusinessRelationship) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *PatchPersonBusinessRelationship) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *PatchPersonBusinessRelationship) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchPersonBusinessRelationship) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchPersonBusinessRelationship) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchPersonBusinessRelationship) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchPersonBusinessRelationship) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationshipType

`func (o *PatchPersonBusinessRelationship) GetRelationshipType() RelationshipTypes`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *PatchPersonBusinessRelationship) GetRelationshipTypeOk() (*RelationshipTypes, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *PatchPersonBusinessRelationship) SetRelationshipType(v RelationshipTypes)`

SetRelationshipType sets RelationshipType field to given value.


### GetTenant

`func (o *PatchPersonBusinessRelationship) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchPersonBusinessRelationship) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchPersonBusinessRelationship) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchPersonBusinessRelationship) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetToBusinessId

`func (o *PatchPersonBusinessRelationship) GetToBusinessId() string`

GetToBusinessId returns the ToBusinessId field if non-nil, zero value otherwise.

### GetToBusinessIdOk

`func (o *PatchPersonBusinessRelationship) GetToBusinessIdOk() (*string, bool)`

GetToBusinessIdOk returns a tuple with the ToBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBusinessId

`func (o *PatchPersonBusinessRelationship) SetToBusinessId(v string)`

SetToBusinessId sets ToBusinessId field to given value.

### HasToBusinessId

`func (o *PatchPersonBusinessRelationship) HasToBusinessId() bool`

HasToBusinessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


