# CreateTransferResponsePull

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
**Type** | [**CreateTransferType**](CreateTransferType.md) |  | 

## Methods

### NewCreateTransferResponsePull

`func NewCreateTransferResponsePull(accountId string, amount int32, creationTime time.Time, currency string, customerId string, id string, lastUpdatedTime time.Time, merchant Merchant, status TransferStatus, tenant string, externalCardId string, type_ CreateTransferType, ) *CreateTransferResponsePull`

NewCreateTransferResponsePull instantiates a new CreateTransferResponsePull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferResponsePullWithDefaults

`func NewCreateTransferResponsePullWithDefaults() *CreateTransferResponsePull`

NewCreateTransferResponsePullWithDefaults instantiates a new CreateTransferResponsePull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateTransferResponsePull) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateTransferResponsePull) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateTransferResponsePull) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *CreateTransferResponsePull) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateTransferResponsePull) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateTransferResponsePull) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCreationTime

`func (o *CreateTransferResponsePull) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreateTransferResponsePull) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreateTransferResponsePull) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *CreateTransferResponsePull) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateTransferResponsePull) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateTransferResponsePull) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *CreateTransferResponsePull) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateTransferResponsePull) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateTransferResponsePull) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *CreateTransferResponsePull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateTransferResponsePull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateTransferResponsePull) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *CreateTransferResponsePull) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CreateTransferResponsePull) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CreateTransferResponsePull) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMerchant

`func (o *CreateTransferResponsePull) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *CreateTransferResponsePull) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *CreateTransferResponsePull) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.


### GetNetworkDeclineDetails

`func (o *CreateTransferResponsePull) GetNetworkDeclineDetails() string`

GetNetworkDeclineDetails returns the NetworkDeclineDetails field if non-nil, zero value otherwise.

### GetNetworkDeclineDetailsOk

`func (o *CreateTransferResponsePull) GetNetworkDeclineDetailsOk() (*string, bool)`

GetNetworkDeclineDetailsOk returns a tuple with the NetworkDeclineDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeclineDetails

`func (o *CreateTransferResponsePull) SetNetworkDeclineDetails(v string)`

SetNetworkDeclineDetails sets NetworkDeclineDetails field to given value.

### HasNetworkDeclineDetails

`func (o *CreateTransferResponsePull) HasNetworkDeclineDetails() bool`

HasNetworkDeclineDetails returns a boolean if a field has been set.

### GetReason

`func (o *CreateTransferResponsePull) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateTransferResponsePull) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateTransferResponsePull) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateTransferResponsePull) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *CreateTransferResponsePull) GetStatus() TransferStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTransferResponsePull) GetStatusOk() (*TransferStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTransferResponsePull) SetStatus(v TransferStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *CreateTransferResponsePull) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreateTransferResponsePull) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreateTransferResponsePull) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionId

`func (o *CreateTransferResponsePull) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CreateTransferResponsePull) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CreateTransferResponsePull) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CreateTransferResponsePull) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetExternalCardId

`func (o *CreateTransferResponsePull) GetExternalCardId() string`

GetExternalCardId returns the ExternalCardId field if non-nil, zero value otherwise.

### GetExternalCardIdOk

`func (o *CreateTransferResponsePull) GetExternalCardIdOk() (*string, bool)`

GetExternalCardIdOk returns a tuple with the ExternalCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCardId

`func (o *CreateTransferResponsePull) SetExternalCardId(v string)`

SetExternalCardId sets ExternalCardId field to given value.


### GetThreeDsId

`func (o *CreateTransferResponsePull) GetThreeDsId() string`

GetThreeDsId returns the ThreeDsId field if non-nil, zero value otherwise.

### GetThreeDsIdOk

`func (o *CreateTransferResponsePull) GetThreeDsIdOk() (*string, bool)`

GetThreeDsIdOk returns a tuple with the ThreeDsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsId

`func (o *CreateTransferResponsePull) SetThreeDsId(v string)`

SetThreeDsId sets ThreeDsId field to given value.

### HasThreeDsId

`func (o *CreateTransferResponsePull) HasThreeDsId() bool`

HasThreeDsId returns a boolean if a field has been set.

### GetType

`func (o *CreateTransferResponsePull) GetType() CreateTransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTransferResponsePull) GetTypeOk() (*CreateTransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTransferResponsePull) SetType(v CreateTransferType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


