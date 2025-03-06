# StopPaymentResponseWAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisputeId** | Pointer to **string** | ID of the dispute that created the stop payment | [optional] 
**ExpiresOn** | Pointer to **time.Time** | The date when this stop payment is no longer valid. This is only for business accounts. | [optional] 
**OriginatorName** | **string** | Name of the originator | 
**StopPaymentId** | **string** |  | 
**TransactionId** | Pointer to **string** | If this stop payment was created from a disputed transaction, transaction_id references the posted transaction. | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** | Timestamp when the stop payment was created. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | Timestamp when stop payment was updated. | [optional] [readonly] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 

## Methods

### NewStopPaymentResponseWAccount

`func NewStopPaymentResponseWAccount(originatorName string, stopPaymentId string, ) *StopPaymentResponseWAccount`

NewStopPaymentResponseWAccount instantiates a new StopPaymentResponseWAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopPaymentResponseWAccountWithDefaults

`func NewStopPaymentResponseWAccountWithDefaults() *StopPaymentResponseWAccount`

NewStopPaymentResponseWAccountWithDefaults instantiates a new StopPaymentResponseWAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisputeId

`func (o *StopPaymentResponseWAccount) GetDisputeId() string`

GetDisputeId returns the DisputeId field if non-nil, zero value otherwise.

### GetDisputeIdOk

`func (o *StopPaymentResponseWAccount) GetDisputeIdOk() (*string, bool)`

GetDisputeIdOk returns a tuple with the DisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeId

`func (o *StopPaymentResponseWAccount) SetDisputeId(v string)`

SetDisputeId sets DisputeId field to given value.

### HasDisputeId

`func (o *StopPaymentResponseWAccount) HasDisputeId() bool`

HasDisputeId returns a boolean if a field has been set.

### GetExpiresOn

`func (o *StopPaymentResponseWAccount) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *StopPaymentResponseWAccount) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *StopPaymentResponseWAccount) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *StopPaymentResponseWAccount) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetOriginatorName

`func (o *StopPaymentResponseWAccount) GetOriginatorName() string`

GetOriginatorName returns the OriginatorName field if non-nil, zero value otherwise.

### GetOriginatorNameOk

`func (o *StopPaymentResponseWAccount) GetOriginatorNameOk() (*string, bool)`

GetOriginatorNameOk returns a tuple with the OriginatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorName

`func (o *StopPaymentResponseWAccount) SetOriginatorName(v string)`

SetOriginatorName sets OriginatorName field to given value.


### GetStopPaymentId

`func (o *StopPaymentResponseWAccount) GetStopPaymentId() string`

GetStopPaymentId returns the StopPaymentId field if non-nil, zero value otherwise.

### GetStopPaymentIdOk

`func (o *StopPaymentResponseWAccount) GetStopPaymentIdOk() (*string, bool)`

GetStopPaymentIdOk returns a tuple with the StopPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPaymentId

`func (o *StopPaymentResponseWAccount) SetStopPaymentId(v string)`

SetStopPaymentId sets StopPaymentId field to given value.


### GetTransactionId

`func (o *StopPaymentResponseWAccount) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *StopPaymentResponseWAccount) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *StopPaymentResponseWAccount) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *StopPaymentResponseWAccount) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetAccountId

`func (o *StopPaymentResponseWAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StopPaymentResponseWAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StopPaymentResponseWAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StopPaymentResponseWAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreationTime

`func (o *StopPaymentResponseWAccount) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *StopPaymentResponseWAccount) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *StopPaymentResponseWAccount) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *StopPaymentResponseWAccount) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *StopPaymentResponseWAccount) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *StopPaymentResponseWAccount) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *StopPaymentResponseWAccount) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *StopPaymentResponseWAccount) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetTenant

`func (o *StopPaymentResponseWAccount) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StopPaymentResponseWAccount) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StopPaymentResponseWAccount) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StopPaymentResponseWAccount) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


