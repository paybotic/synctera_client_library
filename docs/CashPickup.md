# CashPickup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Amount** | **int32** | The amount (in cents) of the transaction | 
**BusinessId** | **string** |  | 
**CreationTime** | **time.Time** |  | [readonly] 
**Currency** | **string** |  | [readonly] 
**EmployeeCustomerId** | **string** |  | 
**Id** | **string** |  | [readonly] 
**LastUpdatedTime** | **time.Time** |  | [readonly] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PickupTime** | Pointer to **time.Time** | the date when the money was actually picked up | [optional] [readonly] 
**PostedAmount** | Pointer to **int32** |  | [optional] [readonly] 
**PostedDate** | Pointer to **string** | The date the transaction was posted. This is the date any money is considered to be added or removed from an account. | [optional] [readonly] 
**ReferenceId** | **string** | An external ID provided by the partner. This is not guaranteed to be globally unique. | 
**ScheduledPickupDate** | **string** | the date when the money was expected to be picked up | 
**TransactionId** | **string** |  | [readonly] 
**Status** | [**CashPickupStatus**](CashPickupStatus.md) |  | 

## Methods

### NewCashPickup

`func NewCashPickup(accountId string, amount int32, businessId string, creationTime time.Time, currency string, employeeCustomerId string, id string, lastUpdatedTime time.Time, referenceId string, scheduledPickupDate string, transactionId string, status CashPickupStatus, ) *CashPickup`

NewCashPickup instantiates a new CashPickup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashPickupWithDefaults

`func NewCashPickupWithDefaults() *CashPickup`

NewCashPickupWithDefaults instantiates a new CashPickup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CashPickup) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CashPickup) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CashPickup) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *CashPickup) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashPickup) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashPickup) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetBusinessId

`func (o *CashPickup) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *CashPickup) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *CashPickup) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetCreationTime

`func (o *CashPickup) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CashPickup) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CashPickup) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *CashPickup) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CashPickup) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CashPickup) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetEmployeeCustomerId

`func (o *CashPickup) GetEmployeeCustomerId() string`

GetEmployeeCustomerId returns the EmployeeCustomerId field if non-nil, zero value otherwise.

### GetEmployeeCustomerIdOk

`func (o *CashPickup) GetEmployeeCustomerIdOk() (*string, bool)`

GetEmployeeCustomerIdOk returns a tuple with the EmployeeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCustomerId

`func (o *CashPickup) SetEmployeeCustomerId(v string)`

SetEmployeeCustomerId sets EmployeeCustomerId field to given value.


### GetId

`func (o *CashPickup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CashPickup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CashPickup) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *CashPickup) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CashPickup) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CashPickup) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMetadata

`func (o *CashPickup) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CashPickup) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CashPickup) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CashPickup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPickupTime

`func (o *CashPickup) GetPickupTime() time.Time`

GetPickupTime returns the PickupTime field if non-nil, zero value otherwise.

### GetPickupTimeOk

`func (o *CashPickup) GetPickupTimeOk() (*time.Time, bool)`

GetPickupTimeOk returns a tuple with the PickupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTime

`func (o *CashPickup) SetPickupTime(v time.Time)`

SetPickupTime sets PickupTime field to given value.

### HasPickupTime

`func (o *CashPickup) HasPickupTime() bool`

HasPickupTime returns a boolean if a field has been set.

### GetPostedAmount

`func (o *CashPickup) GetPostedAmount() int32`

GetPostedAmount returns the PostedAmount field if non-nil, zero value otherwise.

### GetPostedAmountOk

`func (o *CashPickup) GetPostedAmountOk() (*int32, bool)`

GetPostedAmountOk returns a tuple with the PostedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAmount

`func (o *CashPickup) SetPostedAmount(v int32)`

SetPostedAmount sets PostedAmount field to given value.

### HasPostedAmount

`func (o *CashPickup) HasPostedAmount() bool`

HasPostedAmount returns a boolean if a field has been set.

### GetPostedDate

`func (o *CashPickup) GetPostedDate() string`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *CashPickup) GetPostedDateOk() (*string, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *CashPickup) SetPostedDate(v string)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *CashPickup) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *CashPickup) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CashPickup) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CashPickup) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetScheduledPickupDate

`func (o *CashPickup) GetScheduledPickupDate() string`

GetScheduledPickupDate returns the ScheduledPickupDate field if non-nil, zero value otherwise.

### GetScheduledPickupDateOk

`func (o *CashPickup) GetScheduledPickupDateOk() (*string, bool)`

GetScheduledPickupDateOk returns a tuple with the ScheduledPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPickupDate

`func (o *CashPickup) SetScheduledPickupDate(v string)`

SetScheduledPickupDate sets ScheduledPickupDate field to given value.


### GetTransactionId

`func (o *CashPickup) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CashPickup) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CashPickup) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetStatus

`func (o *CashPickup) GetStatus() CashPickupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CashPickup) GetStatusOk() (*CashPickupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CashPickup) SetStatus(v CashPickupStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


