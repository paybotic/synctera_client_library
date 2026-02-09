# DisputeDocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | **time.Time** | The timestamp representing when the object was created | [readonly] 
**DisputeId** | **string** | The unique identifier of the dispute | [readonly] 
**FileName** | **string** |  | 
**Id** | **string** | The unique identifier of a dispute document. | [readonly] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewDisputeDocumentResponse

`func NewDisputeDocumentResponse(creationTime time.Time, disputeId string, fileName string, id string, tenant string, ) *DisputeDocumentResponse`

NewDisputeDocumentResponse instantiates a new DisputeDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeDocumentResponseWithDefaults

`func NewDisputeDocumentResponseWithDefaults() *DisputeDocumentResponse`

NewDisputeDocumentResponseWithDefaults instantiates a new DisputeDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *DisputeDocumentResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DisputeDocumentResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DisputeDocumentResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetDisputeId

`func (o *DisputeDocumentResponse) GetDisputeId() string`

GetDisputeId returns the DisputeId field if non-nil, zero value otherwise.

### GetDisputeIdOk

`func (o *DisputeDocumentResponse) GetDisputeIdOk() (*string, bool)`

GetDisputeIdOk returns a tuple with the DisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeId

`func (o *DisputeDocumentResponse) SetDisputeId(v string)`

SetDisputeId sets DisputeId field to given value.


### GetFileName

`func (o *DisputeDocumentResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DisputeDocumentResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DisputeDocumentResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetId

`func (o *DisputeDocumentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisputeDocumentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisputeDocumentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenant

`func (o *DisputeDocumentResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DisputeDocumentResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DisputeDocumentResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


