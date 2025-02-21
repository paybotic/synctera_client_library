# PatchPayerPayeeRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to [**PayerPayeeAdditionalData**](PayerPayeeAdditionalData.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**FromBusinessId** | Pointer to **string** | Unique ID for the subject business.  | [optional] 
**FromPersonId** | Pointer to **string** | Unique ID for the subject person.  | [optional] 
**Id** | Pointer to **string** | Relationship unique identifier. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**RelationshipType** | [**RelationshipTypes**](RelationshipTypes.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**ToBusinessId** | Pointer to **string** | Unique ID for the related business.  | [optional] 
**ToPersonId** | Pointer to **string** | Unique ID for the related person.  | [optional] 

## Methods

### NewPatchPayerPayeeRelationship

`func NewPatchPayerPayeeRelationship(relationshipType RelationshipTypes, ) *PatchPayerPayeeRelationship`

NewPatchPayerPayeeRelationship instantiates a new PatchPayerPayeeRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPayerPayeeRelationshipWithDefaults

`func NewPatchPayerPayeeRelationshipWithDefaults() *PatchPayerPayeeRelationship`

NewPatchPayerPayeeRelationshipWithDefaults instantiates a new PatchPayerPayeeRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PatchPayerPayeeRelationship) GetAdditionalData() PayerPayeeAdditionalData`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PatchPayerPayeeRelationship) GetAdditionalDataOk() (*PayerPayeeAdditionalData, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PatchPayerPayeeRelationship) SetAdditionalData(v PayerPayeeAdditionalData)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PatchPayerPayeeRelationship) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetCreationTime

`func (o *PatchPayerPayeeRelationship) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PatchPayerPayeeRelationship) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PatchPayerPayeeRelationship) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PatchPayerPayeeRelationship) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetFromBusinessId

`func (o *PatchPayerPayeeRelationship) GetFromBusinessId() string`

GetFromBusinessId returns the FromBusinessId field if non-nil, zero value otherwise.

### GetFromBusinessIdOk

`func (o *PatchPayerPayeeRelationship) GetFromBusinessIdOk() (*string, bool)`

GetFromBusinessIdOk returns a tuple with the FromBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBusinessId

`func (o *PatchPayerPayeeRelationship) SetFromBusinessId(v string)`

SetFromBusinessId sets FromBusinessId field to given value.

### HasFromBusinessId

`func (o *PatchPayerPayeeRelationship) HasFromBusinessId() bool`

HasFromBusinessId returns a boolean if a field has been set.

### GetFromPersonId

`func (o *PatchPayerPayeeRelationship) GetFromPersonId() string`

GetFromPersonId returns the FromPersonId field if non-nil, zero value otherwise.

### GetFromPersonIdOk

`func (o *PatchPayerPayeeRelationship) GetFromPersonIdOk() (*string, bool)`

GetFromPersonIdOk returns a tuple with the FromPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPersonId

`func (o *PatchPayerPayeeRelationship) SetFromPersonId(v string)`

SetFromPersonId sets FromPersonId field to given value.

### HasFromPersonId

`func (o *PatchPayerPayeeRelationship) HasFromPersonId() bool`

HasFromPersonId returns a boolean if a field has been set.

### GetId

`func (o *PatchPayerPayeeRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchPayerPayeeRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchPayerPayeeRelationship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchPayerPayeeRelationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *PatchPayerPayeeRelationship) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *PatchPayerPayeeRelationship) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *PatchPayerPayeeRelationship) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *PatchPayerPayeeRelationship) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchPayerPayeeRelationship) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchPayerPayeeRelationship) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchPayerPayeeRelationship) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchPayerPayeeRelationship) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationshipType

`func (o *PatchPayerPayeeRelationship) GetRelationshipType() RelationshipTypes`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *PatchPayerPayeeRelationship) GetRelationshipTypeOk() (*RelationshipTypes, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *PatchPayerPayeeRelationship) SetRelationshipType(v RelationshipTypes)`

SetRelationshipType sets RelationshipType field to given value.


### GetTenant

`func (o *PatchPayerPayeeRelationship) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchPayerPayeeRelationship) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchPayerPayeeRelationship) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchPayerPayeeRelationship) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetToBusinessId

`func (o *PatchPayerPayeeRelationship) GetToBusinessId() string`

GetToBusinessId returns the ToBusinessId field if non-nil, zero value otherwise.

### GetToBusinessIdOk

`func (o *PatchPayerPayeeRelationship) GetToBusinessIdOk() (*string, bool)`

GetToBusinessIdOk returns a tuple with the ToBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBusinessId

`func (o *PatchPayerPayeeRelationship) SetToBusinessId(v string)`

SetToBusinessId sets ToBusinessId field to given value.

### HasToBusinessId

`func (o *PatchPayerPayeeRelationship) HasToBusinessId() bool`

HasToBusinessId returns a boolean if a field has been set.

### GetToPersonId

`func (o *PatchPayerPayeeRelationship) GetToPersonId() string`

GetToPersonId returns the ToPersonId field if non-nil, zero value otherwise.

### GetToPersonIdOk

`func (o *PatchPayerPayeeRelationship) GetToPersonIdOk() (*string, bool)`

GetToPersonIdOk returns a tuple with the ToPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPersonId

`func (o *PatchPayerPayeeRelationship) SetToPersonId(v string)`

SetToPersonId sets ToPersonId field to given value.

### HasToPersonId

`func (o *PatchPayerPayeeRelationship) HasToPersonId() bool`

HasToPersonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


