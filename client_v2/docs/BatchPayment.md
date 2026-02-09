# BatchPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The total amount of the batch transfer.  | [readonly] 
**BatchPaymentTemplateId** | **string** | The ID of the batch Payment template that was used to create the batch.  | 
**BatchedTransferIds** | **[]string** | The IDs of the transfers that are part of the batch. These values can be modified by the client before  the batch is in a terminal status.  | 
**CreationTime** | **time.Time** |  | [readonly] 
**Id** | **string** |  | [readonly] 
**LastUpdatedTime** | **time.Time** |  | [readonly] 
**PaymentRail** | Pointer to **string** | The payment rail that was used to process the batch.  | [optional] 
**PaymentRailTransferId** | Pointer to **string** | The ID of the payment rail transfer that was used to process the batch.  | [optional] 
**Status** | [**BatchStatus**](BatchStatus.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 
**TransactionId** | Pointer to **string** | The transaction ID of the batch which represents a transaction on the ledger.  | [optional] 

## Methods

### NewBatchPayment

`func NewBatchPayment(amount int64, batchPaymentTemplateId string, batchedTransferIds []string, creationTime time.Time, id string, lastUpdatedTime time.Time, status BatchStatus, ) *BatchPayment`

NewBatchPayment instantiates a new BatchPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchPaymentWithDefaults

`func NewBatchPaymentWithDefaults() *BatchPayment`

NewBatchPaymentWithDefaults instantiates a new BatchPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BatchPayment) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BatchPayment) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BatchPayment) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetBatchPaymentTemplateId

`func (o *BatchPayment) GetBatchPaymentTemplateId() string`

GetBatchPaymentTemplateId returns the BatchPaymentTemplateId field if non-nil, zero value otherwise.

### GetBatchPaymentTemplateIdOk

`func (o *BatchPayment) GetBatchPaymentTemplateIdOk() (*string, bool)`

GetBatchPaymentTemplateIdOk returns a tuple with the BatchPaymentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchPaymentTemplateId

`func (o *BatchPayment) SetBatchPaymentTemplateId(v string)`

SetBatchPaymentTemplateId sets BatchPaymentTemplateId field to given value.


### GetBatchedTransferIds

`func (o *BatchPayment) GetBatchedTransferIds() []string`

GetBatchedTransferIds returns the BatchedTransferIds field if non-nil, zero value otherwise.

### GetBatchedTransferIdsOk

`func (o *BatchPayment) GetBatchedTransferIdsOk() (*[]string, bool)`

GetBatchedTransferIdsOk returns a tuple with the BatchedTransferIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchedTransferIds

`func (o *BatchPayment) SetBatchedTransferIds(v []string)`

SetBatchedTransferIds sets BatchedTransferIds field to given value.


### GetCreationTime

`func (o *BatchPayment) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BatchPayment) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BatchPayment) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *BatchPayment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchPayment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchPayment) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *BatchPayment) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BatchPayment) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BatchPayment) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetPaymentRail

`func (o *BatchPayment) GetPaymentRail() string`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *BatchPayment) GetPaymentRailOk() (*string, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *BatchPayment) SetPaymentRail(v string)`

SetPaymentRail sets PaymentRail field to given value.

### HasPaymentRail

`func (o *BatchPayment) HasPaymentRail() bool`

HasPaymentRail returns a boolean if a field has been set.

### GetPaymentRailTransferId

`func (o *BatchPayment) GetPaymentRailTransferId() string`

GetPaymentRailTransferId returns the PaymentRailTransferId field if non-nil, zero value otherwise.

### GetPaymentRailTransferIdOk

`func (o *BatchPayment) GetPaymentRailTransferIdOk() (*string, bool)`

GetPaymentRailTransferIdOk returns a tuple with the PaymentRailTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRailTransferId

`func (o *BatchPayment) SetPaymentRailTransferId(v string)`

SetPaymentRailTransferId sets PaymentRailTransferId field to given value.

### HasPaymentRailTransferId

`func (o *BatchPayment) HasPaymentRailTransferId() bool`

HasPaymentRailTransferId returns a boolean if a field has been set.

### GetStatus

`func (o *BatchPayment) GetStatus() BatchStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchPayment) GetStatusOk() (*BatchStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchPayment) SetStatus(v BatchStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *BatchPayment) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BatchPayment) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BatchPayment) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BatchPayment) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetTransactionId

`func (o *BatchPayment) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BatchPayment) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BatchPayment) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BatchPayment) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


