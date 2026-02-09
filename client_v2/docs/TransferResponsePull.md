# TransferResponsePull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the Synctera account into which or from which funds were moved | 
**Amount** | **int32** | Amount of the transfer in cents | 
**CreationTime** | **time.Time** |  | 
**Currency** | **string** | ISO 4217  Alpha-3 currency code | 
**CustomerId** | **string** | The customer_id of the cardholder | 
**Id** | **string** | The ID of the transfer | 
**LastUpdatedTime** | **time.Time** |  | 
**Merchant** | [**Merchant**](Merchant.md) |  | 
**NetworkDeclineDetails** | Pointer to **string** | If available, a human readable string indicating why a transfer was declined downstream of our system | [optional] 
**Reason** | Pointer to **string** | The reason for the status, e.g. INSUFFICIENT_FUNDS, SUSPECTED_FRAUD, NETWORK_DECLINED | [optional] 
**Status** | [**TransferStatus**](TransferStatus.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**TransactionId** | Pointer to **string** | The transaction ID | [optional] 
**ExternalCardId** | **string** | The ID of the External Card associated with the operation | 
**ThreeDsId** | Pointer to **string** | Unique identifier of an External Card Transfer 3-D Secure Authorization - conditionally required according to your program&#39;s 3DS policy | [optional] 
**Type** | [**TransferType**](TransferType.md) |  | 

## Methods

### NewTransferResponsePull

`func NewTransferResponsePull(accountId string, amount int32, creationTime time.Time, currency string, customerId string, id string, lastUpdatedTime time.Time, merchant Merchant, status TransferStatus, tenant string, externalCardId string, type_ TransferType, ) *TransferResponsePull`

NewTransferResponsePull instantiates a new TransferResponsePull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResponsePullWithDefaults

`func NewTransferResponsePullWithDefaults() *TransferResponsePull`

NewTransferResponsePullWithDefaults instantiates a new TransferResponsePull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransferResponsePull) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferResponsePull) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferResponsePull) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *TransferResponsePull) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferResponsePull) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferResponsePull) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCreationTime

`func (o *TransferResponsePull) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *TransferResponsePull) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *TransferResponsePull) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *TransferResponsePull) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransferResponsePull) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransferResponsePull) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *TransferResponsePull) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TransferResponsePull) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TransferResponsePull) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *TransferResponsePull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferResponsePull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferResponsePull) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *TransferResponsePull) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *TransferResponsePull) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *TransferResponsePull) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMerchant

`func (o *TransferResponsePull) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransferResponsePull) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransferResponsePull) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.


### GetNetworkDeclineDetails

`func (o *TransferResponsePull) GetNetworkDeclineDetails() string`

GetNetworkDeclineDetails returns the NetworkDeclineDetails field if non-nil, zero value otherwise.

### GetNetworkDeclineDetailsOk

`func (o *TransferResponsePull) GetNetworkDeclineDetailsOk() (*string, bool)`

GetNetworkDeclineDetailsOk returns a tuple with the NetworkDeclineDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeclineDetails

`func (o *TransferResponsePull) SetNetworkDeclineDetails(v string)`

SetNetworkDeclineDetails sets NetworkDeclineDetails field to given value.

### HasNetworkDeclineDetails

`func (o *TransferResponsePull) HasNetworkDeclineDetails() bool`

HasNetworkDeclineDetails returns a boolean if a field has been set.

### GetReason

`func (o *TransferResponsePull) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransferResponsePull) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransferResponsePull) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransferResponsePull) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *TransferResponsePull) GetStatus() TransferStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferResponsePull) GetStatusOk() (*TransferStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferResponsePull) SetStatus(v TransferStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *TransferResponsePull) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransferResponsePull) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransferResponsePull) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionId

`func (o *TransferResponsePull) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransferResponsePull) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransferResponsePull) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TransferResponsePull) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetExternalCardId

`func (o *TransferResponsePull) GetExternalCardId() string`

GetExternalCardId returns the ExternalCardId field if non-nil, zero value otherwise.

### GetExternalCardIdOk

`func (o *TransferResponsePull) GetExternalCardIdOk() (*string, bool)`

GetExternalCardIdOk returns a tuple with the ExternalCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCardId

`func (o *TransferResponsePull) SetExternalCardId(v string)`

SetExternalCardId sets ExternalCardId field to given value.


### GetThreeDsId

`func (o *TransferResponsePull) GetThreeDsId() string`

GetThreeDsId returns the ThreeDsId field if non-nil, zero value otherwise.

### GetThreeDsIdOk

`func (o *TransferResponsePull) GetThreeDsIdOk() (*string, bool)`

GetThreeDsIdOk returns a tuple with the ThreeDsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsId

`func (o *TransferResponsePull) SetThreeDsId(v string)`

SetThreeDsId sets ThreeDsId field to given value.

### HasThreeDsId

`func (o *TransferResponsePull) HasThreeDsId() bool`

HasThreeDsId returns a boolean if a field has been set.

### GetType

`func (o *TransferResponsePull) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferResponsePull) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferResponsePull) SetType(v TransferType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


