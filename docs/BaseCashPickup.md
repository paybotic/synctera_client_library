# BaseCashPickup

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

## Methods

### NewBaseCashPickup

`func NewBaseCashPickup(accountId string, amount int32, referenceId string, ) *BaseCashPickup`

NewBaseCashPickup instantiates a new BaseCashPickup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCashPickupWithDefaults

`func NewBaseCashPickupWithDefaults() *BaseCashPickup`

NewBaseCashPickupWithDefaults instantiates a new BaseCashPickup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BaseCashPickup) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BaseCashPickup) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BaseCashPickup) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *BaseCashPickup) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BaseCashPickup) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BaseCashPickup) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetBusinessId

`func (o *BaseCashPickup) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *BaseCashPickup) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *BaseCashPickup) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *BaseCashPickup) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *BaseCashPickup) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BaseCashPickup) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BaseCashPickup) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BaseCashPickup) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *BaseCashPickup) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BaseCashPickup) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BaseCashPickup) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BaseCashPickup) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmployeeCustomerId

`func (o *BaseCashPickup) GetEmployeeCustomerId() string`

GetEmployeeCustomerId returns the EmployeeCustomerId field if non-nil, zero value otherwise.

### GetEmployeeCustomerIdOk

`func (o *BaseCashPickup) GetEmployeeCustomerIdOk() (*string, bool)`

GetEmployeeCustomerIdOk returns a tuple with the EmployeeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCustomerId

`func (o *BaseCashPickup) SetEmployeeCustomerId(v string)`

SetEmployeeCustomerId sets EmployeeCustomerId field to given value.

### HasEmployeeCustomerId

`func (o *BaseCashPickup) HasEmployeeCustomerId() bool`

HasEmployeeCustomerId returns a boolean if a field has been set.

### GetId

`func (o *BaseCashPickup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseCashPickup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseCashPickup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseCashPickup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *BaseCashPickup) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BaseCashPickup) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BaseCashPickup) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *BaseCashPickup) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *BaseCashPickup) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BaseCashPickup) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BaseCashPickup) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BaseCashPickup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPickupTime

`func (o *BaseCashPickup) GetPickupTime() time.Time`

GetPickupTime returns the PickupTime field if non-nil, zero value otherwise.

### GetPickupTimeOk

`func (o *BaseCashPickup) GetPickupTimeOk() (*time.Time, bool)`

GetPickupTimeOk returns a tuple with the PickupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTime

`func (o *BaseCashPickup) SetPickupTime(v time.Time)`

SetPickupTime sets PickupTime field to given value.

### HasPickupTime

`func (o *BaseCashPickup) HasPickupTime() bool`

HasPickupTime returns a boolean if a field has been set.

### GetPostedAmount

`func (o *BaseCashPickup) GetPostedAmount() int32`

GetPostedAmount returns the PostedAmount field if non-nil, zero value otherwise.

### GetPostedAmountOk

`func (o *BaseCashPickup) GetPostedAmountOk() (*int32, bool)`

GetPostedAmountOk returns a tuple with the PostedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAmount

`func (o *BaseCashPickup) SetPostedAmount(v int32)`

SetPostedAmount sets PostedAmount field to given value.

### HasPostedAmount

`func (o *BaseCashPickup) HasPostedAmount() bool`

HasPostedAmount returns a boolean if a field has been set.

### GetPostedDate

`func (o *BaseCashPickup) GetPostedDate() string`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *BaseCashPickup) GetPostedDateOk() (*string, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *BaseCashPickup) SetPostedDate(v string)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *BaseCashPickup) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *BaseCashPickup) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *BaseCashPickup) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *BaseCashPickup) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetScheduledPickupDate

`func (o *BaseCashPickup) GetScheduledPickupDate() string`

GetScheduledPickupDate returns the ScheduledPickupDate field if non-nil, zero value otherwise.

### GetScheduledPickupDateOk

`func (o *BaseCashPickup) GetScheduledPickupDateOk() (*string, bool)`

GetScheduledPickupDateOk returns a tuple with the ScheduledPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPickupDate

`func (o *BaseCashPickup) SetScheduledPickupDate(v string)`

SetScheduledPickupDate sets ScheduledPickupDate field to given value.

### HasScheduledPickupDate

`func (o *BaseCashPickup) HasScheduledPickupDate() bool`

HasScheduledPickupDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *BaseCashPickup) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BaseCashPickup) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BaseCashPickup) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BaseCashPickup) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


