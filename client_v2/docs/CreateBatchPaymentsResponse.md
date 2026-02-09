# CreateBatchPaymentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batches** | [**[]BatchPayment**](BatchPayment.md) | A collection of batch payments  | 
**DryRun** | Pointer to **bool** | Whether or not the batch should be created as a dry run. If the batch is created as a dry run, the batch will not be processed.  | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 

## Methods

### NewCreateBatchPaymentsResponse

`func NewCreateBatchPaymentsResponse(batches []BatchPayment, ) *CreateBatchPaymentsResponse`

NewCreateBatchPaymentsResponse instantiates a new CreateBatchPaymentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchPaymentsResponseWithDefaults

`func NewCreateBatchPaymentsResponseWithDefaults() *CreateBatchPaymentsResponse`

NewCreateBatchPaymentsResponseWithDefaults instantiates a new CreateBatchPaymentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatches

`func (o *CreateBatchPaymentsResponse) GetBatches() []BatchPayment`

GetBatches returns the Batches field if non-nil, zero value otherwise.

### GetBatchesOk

`func (o *CreateBatchPaymentsResponse) GetBatchesOk() (*[]BatchPayment, bool)`

GetBatchesOk returns a tuple with the Batches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatches

`func (o *CreateBatchPaymentsResponse) SetBatches(v []BatchPayment)`

SetBatches sets Batches field to given value.


### GetDryRun

`func (o *CreateBatchPaymentsResponse) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateBatchPaymentsResponse) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateBatchPaymentsResponse) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateBatchPaymentsResponse) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetTenant

`func (o *CreateBatchPaymentsResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreateBatchPaymentsResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreateBatchPaymentsResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CreateBatchPaymentsResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


