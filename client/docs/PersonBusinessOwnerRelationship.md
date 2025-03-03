# PersonBusinessOwnerRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | [**AdditionalOwnerData**](AdditionalOwnerData.md) |  | 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**FromPersonId** | **string** | Unique ID for the subject person.  | 
**Id** | Pointer to **string** | Relationship unique identifier. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**RelationshipType** | [**RelationshipTypes**](RelationshipTypes.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**ToBusinessId** | **string** | Unique ID for the related business.  | 

## Methods

### NewPersonBusinessOwnerRelationship

`func NewPersonBusinessOwnerRelationship(additionalData AdditionalOwnerData, fromPersonId string, relationshipType RelationshipTypes, toBusinessId string, ) *PersonBusinessOwnerRelationship`

NewPersonBusinessOwnerRelationship instantiates a new PersonBusinessOwnerRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonBusinessOwnerRelationshipWithDefaults

`func NewPersonBusinessOwnerRelationshipWithDefaults() *PersonBusinessOwnerRelationship`

NewPersonBusinessOwnerRelationshipWithDefaults instantiates a new PersonBusinessOwnerRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PersonBusinessOwnerRelationship) GetAdditionalData() AdditionalOwnerData`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PersonBusinessOwnerRelationship) GetAdditionalDataOk() (*AdditionalOwnerData, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PersonBusinessOwnerRelationship) SetAdditionalData(v AdditionalOwnerData)`

SetAdditionalData sets AdditionalData field to given value.


### GetCreationTime

`func (o *PersonBusinessOwnerRelationship) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PersonBusinessOwnerRelationship) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PersonBusinessOwnerRelationship) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PersonBusinessOwnerRelationship) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetFromPersonId

`func (o *PersonBusinessOwnerRelationship) GetFromPersonId() string`

GetFromPersonId returns the FromPersonId field if non-nil, zero value otherwise.

### GetFromPersonIdOk

`func (o *PersonBusinessOwnerRelationship) GetFromPersonIdOk() (*string, bool)`

GetFromPersonIdOk returns a tuple with the FromPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPersonId

`func (o *PersonBusinessOwnerRelationship) SetFromPersonId(v string)`

SetFromPersonId sets FromPersonId field to given value.


### GetId

`func (o *PersonBusinessOwnerRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonBusinessOwnerRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonBusinessOwnerRelationship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PersonBusinessOwnerRelationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *PersonBusinessOwnerRelationship) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *PersonBusinessOwnerRelationship) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *PersonBusinessOwnerRelationship) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *PersonBusinessOwnerRelationship) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *PersonBusinessOwnerRelationship) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PersonBusinessOwnerRelationship) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PersonBusinessOwnerRelationship) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PersonBusinessOwnerRelationship) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationshipType

`func (o *PersonBusinessOwnerRelationship) GetRelationshipType() RelationshipTypes`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *PersonBusinessOwnerRelationship) GetRelationshipTypeOk() (*RelationshipTypes, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *PersonBusinessOwnerRelationship) SetRelationshipType(v RelationshipTypes)`

SetRelationshipType sets RelationshipType field to given value.


### GetTenant

`func (o *PersonBusinessOwnerRelationship) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PersonBusinessOwnerRelationship) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PersonBusinessOwnerRelationship) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PersonBusinessOwnerRelationship) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetToBusinessId

`func (o *PersonBusinessOwnerRelationship) GetToBusinessId() string`

GetToBusinessId returns the ToBusinessId field if non-nil, zero value otherwise.

### GetToBusinessIdOk

`func (o *PersonBusinessOwnerRelationship) GetToBusinessIdOk() (*string, bool)`

GetToBusinessIdOk returns a tuple with the ToBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBusinessId

`func (o *PersonBusinessOwnerRelationship) SetToBusinessId(v string)`

SetToBusinessId sets ToBusinessId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


