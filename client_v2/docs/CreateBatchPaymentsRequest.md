# CreateBatchPaymentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | Whether or not the batch should be created as a dry run. If the batch is created as a dry run, the batch will not be processed.  | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 
**TransactionIds** | **[]string** | The IDs of the transfers that are part of the batch. These values can be modified by the client before  the batch is in a terminal status.  | 

## Methods

### NewCreateBatchPaymentsRequest

`func NewCreateBatchPaymentsRequest(transactionIds []string, ) *CreateBatchPaymentsRequest`

NewCreateBatchPaymentsRequest instantiates a new CreateBatchPaymentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchPaymentsRequestWithDefaults

`func NewCreateBatchPaymentsRequestWithDefaults() *CreateBatchPaymentsRequest`

NewCreateBatchPaymentsRequestWithDefaults instantiates a new CreateBatchPaymentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateBatchPaymentsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateBatchPaymentsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateBatchPaymentsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateBatchPaymentsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetTenant

`func (o *CreateBatchPaymentsRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreateBatchPaymentsRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreateBatchPaymentsRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CreateBatchPaymentsRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetTransactionIds

`func (o *CreateBatchPaymentsRequest) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *CreateBatchPaymentsRequest) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *CreateBatchPaymentsRequest) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


