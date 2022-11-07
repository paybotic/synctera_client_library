# CashPickupPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Amount** | **int32** | The amount (in cents) of the transaction | 
**BusinessId** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Currency** | Pointer to **string** |  | [optional] [readonly] 
**EmployeeCustomerId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PickupTime** | Pointer to **time.Time** | the date when the money was actually picked up | [optional] [readonly] 
**PostedAmount** | Pointer to **int32** |  | [optional] [readonly] 
**PostedDate** | Pointer to **string** | The date the transaction was posted. This is the date any money is considered to be added or removed from an account. | [optional] [readonly] 
**ReferenceId** | **string** | An external ID provided by the partner. This is not guaranteed to be globally unique. | 
**ScheduledPickupDate** | Pointer to **string** | the date when the money was expected to be picked up | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCashPickupPatchRequest

`func NewCashPickupPatchRequest(accountId string, amount int32, referenceId string, ) *CashPickupPatchRequest`

NewCashPickupPatchRequest instantiates a new CashPickupPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashPickupPatchRequestWithDefaults

`func NewCashPickupPatchRequestWithDefaults() *CashPickupPatchRequest`

NewCashPickupPatchRequestWithDefaults instantiates a new CashPickupPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CashPickupPatchRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CashPickupPatchRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CashPickupPatchRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *CashPickupPatchRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashPickupPatchRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashPickupPatchRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetBusinessId

`func (o *CashPickupPatchRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *CashPickupPatchRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *CashPickupPatchRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *CashPickupPatchRequest) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *CashPickupPatchRequest) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CashPickupPatchRequest) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CashPickupPatchRequest) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CashPickupPatchRequest) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *CashPickupPatchRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CashPickupPatchRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CashPickupPatchRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CashPickupPatchRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmployeeCustomerId

`func (o *CashPickupPatchRequest) GetEmployeeCustomerId() string`

GetEmployeeCustomerId returns the EmployeeCustomerId field if non-nil, zero value otherwise.

### GetEmployeeCustomerIdOk

`func (o *CashPickupPatchRequest) GetEmployeeCustomerIdOk() (*string, bool)`

GetEmployeeCustomerIdOk returns a tuple with the EmployeeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCustomerId

`func (o *CashPickupPatchRequest) SetEmployeeCustomerId(v string)`

SetEmployeeCustomerId sets EmployeeCustomerId field to given value.

### HasEmployeeCustomerId

`func (o *CashPickupPatchRequest) HasEmployeeCustomerId() bool`

HasEmployeeCustomerId returns a boolean if a field has been set.

### GetId

`func (o *CashPickupPatchRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CashPickupPatchRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CashPickupPatchRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CashPickupPatchRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *CashPickupPatchRequest) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CashPickupPatchRequest) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CashPickupPatchRequest) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *CashPickupPatchRequest) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *CashPickupPatchRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CashPickupPatchRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CashPickupPatchRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CashPickupPatchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPickupTime

`func (o *CashPickupPatchRequest) GetPickupTime() time.Time`

GetPickupTime returns the PickupTime field if non-nil, zero value otherwise.

### GetPickupTimeOk

`func (o *CashPickupPatchRequest) GetPickupTimeOk() (*time.Time, bool)`

GetPickupTimeOk returns a tuple with the PickupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTime

`func (o *CashPickupPatchRequest) SetPickupTime(v time.Time)`

SetPickupTime sets PickupTime field to given value.

### HasPickupTime

`func (o *CashPickupPatchRequest) HasPickupTime() bool`

HasPickupTime returns a boolean if a field has been set.

### GetPostedAmount

`func (o *CashPickupPatchRequest) GetPostedAmount() int32`

GetPostedAmount returns the PostedAmount field if non-nil, zero value otherwise.

### GetPostedAmountOk

`func (o *CashPickupPatchRequest) GetPostedAmountOk() (*int32, bool)`

GetPostedAmountOk returns a tuple with the PostedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAmount

`func (o *CashPickupPatchRequest) SetPostedAmount(v int32)`

SetPostedAmount sets PostedAmount field to given value.

### HasPostedAmount

`func (o *CashPickupPatchRequest) HasPostedAmount() bool`

HasPostedAmount returns a boolean if a field has been set.

### GetPostedDate

`func (o *CashPickupPatchRequest) GetPostedDate() string`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *CashPickupPatchRequest) GetPostedDateOk() (*string, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *CashPickupPatchRequest) SetPostedDate(v string)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *CashPickupPatchRequest) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *CashPickupPatchRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CashPickupPatchRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CashPickupPatchRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetScheduledPickupDate

`func (o *CashPickupPatchRequest) GetScheduledPickupDate() string`

GetScheduledPickupDate returns the ScheduledPickupDate field if non-nil, zero value otherwise.

### GetScheduledPickupDateOk

`func (o *CashPickupPatchRequest) GetScheduledPickupDateOk() (*string, bool)`

GetScheduledPickupDateOk returns a tuple with the ScheduledPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPickupDate

`func (o *CashPickupPatchRequest) SetScheduledPickupDate(v string)`

SetScheduledPickupDate sets ScheduledPickupDate field to given value.

### HasScheduledPickupDate

`func (o *CashPickupPatchRequest) HasScheduledPickupDate() bool`

HasScheduledPickupDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *CashPickupPatchRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CashPickupPatchRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CashPickupPatchRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CashPickupPatchRequest) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetStatus

`func (o *CashPickupPatchRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CashPickupPatchRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CashPickupPatchRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CashPickupPatchRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


