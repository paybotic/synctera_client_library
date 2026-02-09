# BatchPaymentPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchedTransferIds** | Pointer to **[]string** | The IDs of the transfers that are part of the batch. These values can be modified by the client before  the batch is in a terminal status.  | [optional] 
**CreationTime** | **time.Time** |  | [readonly] 
**LastUpdatedTime** | **time.Time** |  | [readonly] 
**Status** | Pointer to [**BatchStatus**](BatchStatus.md) |  | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 

## Methods

### NewBatchPaymentPatchRequest

`func NewBatchPaymentPatchRequest(creationTime time.Time, lastUpdatedTime time.Time, ) *BatchPaymentPatchRequest`

NewBatchPaymentPatchRequest instantiates a new BatchPaymentPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchPaymentPatchRequestWithDefaults

`func NewBatchPaymentPatchRequestWithDefaults() *BatchPaymentPatchRequest`

NewBatchPaymentPatchRequestWithDefaults instantiates a new BatchPaymentPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchedTransferIds

`func (o *BatchPaymentPatchRequest) GetBatchedTransferIds() []string`

GetBatchedTransferIds returns the BatchedTransferIds field if non-nil, zero value otherwise.

### GetBatchedTransferIdsOk

`func (o *BatchPaymentPatchRequest) GetBatchedTransferIdsOk() (*[]string, bool)`

GetBatchedTransferIdsOk returns a tuple with the BatchedTransferIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchedTransferIds

`func (o *BatchPaymentPatchRequest) SetBatchedTransferIds(v []string)`

SetBatchedTransferIds sets BatchedTransferIds field to given value.

### HasBatchedTransferIds

`func (o *BatchPaymentPatchRequest) HasBatchedTransferIds() bool`

HasBatchedTransferIds returns a boolean if a field has been set.

### GetCreationTime

`func (o *BatchPaymentPatchRequest) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BatchPaymentPatchRequest) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BatchPaymentPatchRequest) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetLastUpdatedTime

`func (o *BatchPaymentPatchRequest) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BatchPaymentPatchRequest) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BatchPaymentPatchRequest) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetStatus

`func (o *BatchPaymentPatchRequest) GetStatus() BatchStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchPaymentPatchRequest) GetStatusOk() (*BatchStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchPaymentPatchRequest) SetStatus(v BatchStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchPaymentPatchRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *BatchPaymentPatchRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BatchPaymentPatchRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BatchPaymentPatchRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BatchPaymentPatchRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


