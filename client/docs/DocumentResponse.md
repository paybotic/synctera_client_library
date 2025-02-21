# DocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableVersions** | Pointer to **[]int32** | All document versions | [optional] 
**AvailableVersionsInfo** | Pointer to [**[]DocumentVersionInfo**](DocumentVersionInfo.md) | Metadata of all document versions | [optional] 
**BatchId** | Pointer to **string** | The ID of the batch that the document belongs to | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created | [optional] [readonly] 
**DeletionReason** | Pointer to **string** | An explanation why the file was deleted. You must set a document&#39;s deletion_reason before deleting it. | [optional] 
**Description** | Pointer to **string** | A description of the document | [optional] 
**Encryption** | Pointer to [**DocumentEncryption**](DocumentEncryption.md) |  | [optional] 
**FileName** | Pointer to **string** | The file name of the document | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for this resource | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Name** | Pointer to **string** | A user-friendly name for the document | [optional] 
**RelatedResourceId** | Pointer to **string** | The ID of the resource related to the document | [optional] 
**RelatedResourceType** | Pointer to [**RelatedResourceType**](RelatedResourceType.md) |  | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**Type** | Pointer to [**DocumentType**](DocumentType.md) |  | [optional] 
**Version** | Pointer to **int32** | Positive integer representing the version of the document | [optional] 

## Methods

### NewDocumentResponse

`func NewDocumentResponse() *DocumentResponse`

NewDocumentResponse instantiates a new DocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentResponseWithDefaults

`func NewDocumentResponseWithDefaults() *DocumentResponse`

NewDocumentResponseWithDefaults instantiates a new DocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableVersions

`func (o *DocumentResponse) GetAvailableVersions() []int32`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *DocumentResponse) GetAvailableVersionsOk() (*[]int32, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *DocumentResponse) SetAvailableVersions(v []int32)`

SetAvailableVersions sets AvailableVersions field to given value.

### HasAvailableVersions

`func (o *DocumentResponse) HasAvailableVersions() bool`

HasAvailableVersions returns a boolean if a field has been set.

### GetAvailableVersionsInfo

`func (o *DocumentResponse) GetAvailableVersionsInfo() []DocumentVersionInfo`

GetAvailableVersionsInfo returns the AvailableVersionsInfo field if non-nil, zero value otherwise.

### GetAvailableVersionsInfoOk

`func (o *DocumentResponse) GetAvailableVersionsInfoOk() (*[]DocumentVersionInfo, bool)`

GetAvailableVersionsInfoOk returns a tuple with the AvailableVersionsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersionsInfo

`func (o *DocumentResponse) SetAvailableVersionsInfo(v []DocumentVersionInfo)`

SetAvailableVersionsInfo sets AvailableVersionsInfo field to given value.

### HasAvailableVersionsInfo

`func (o *DocumentResponse) HasAvailableVersionsInfo() bool`

HasAvailableVersionsInfo returns a boolean if a field has been set.

### GetBatchId

`func (o *DocumentResponse) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *DocumentResponse) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *DocumentResponse) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *DocumentResponse) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetCreationTime

`func (o *DocumentResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DocumentResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DocumentResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *DocumentResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDeletionReason

`func (o *DocumentResponse) GetDeletionReason() string`

GetDeletionReason returns the DeletionReason field if non-nil, zero value otherwise.

### GetDeletionReasonOk

`func (o *DocumentResponse) GetDeletionReasonOk() (*string, bool)`

GetDeletionReasonOk returns a tuple with the DeletionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionReason

`func (o *DocumentResponse) SetDeletionReason(v string)`

SetDeletionReason sets DeletionReason field to given value.

### HasDeletionReason

`func (o *DocumentResponse) HasDeletionReason() bool`

HasDeletionReason returns a boolean if a field has been set.

### GetDescription

`func (o *DocumentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DocumentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEncryption

`func (o *DocumentResponse) GetEncryption() DocumentEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *DocumentResponse) GetEncryptionOk() (*DocumentEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *DocumentResponse) SetEncryption(v DocumentEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *DocumentResponse) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFileName

`func (o *DocumentResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DocumentResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DocumentResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DocumentResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetId

`func (o *DocumentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *DocumentResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *DocumentResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *DocumentResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *DocumentResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *DocumentResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DocumentResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DocumentResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DocumentResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DocumentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelatedResourceId

`func (o *DocumentResponse) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *DocumentResponse) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *DocumentResponse) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.

### HasRelatedResourceId

`func (o *DocumentResponse) HasRelatedResourceId() bool`

HasRelatedResourceId returns a boolean if a field has been set.

### GetRelatedResourceType

`func (o *DocumentResponse) GetRelatedResourceType() RelatedResourceType`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *DocumentResponse) GetRelatedResourceTypeOk() (*RelatedResourceType, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *DocumentResponse) SetRelatedResourceType(v RelatedResourceType)`

SetRelatedResourceType sets RelatedResourceType field to given value.

### HasRelatedResourceType

`func (o *DocumentResponse) HasRelatedResourceType() bool`

HasRelatedResourceType returns a boolean if a field has been set.

### GetTenant

`func (o *DocumentResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DocumentResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DocumentResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *DocumentResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *DocumentResponse) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentResponse) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentResponse) SetType(v DocumentType)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *DocumentResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DocumentResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DocumentResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DocumentResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


