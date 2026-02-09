# TransferResponseToken

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
**CardDetails** | [**ExternalCardDetails**](ExternalCardDetails.md) |  | 
**Type** | [**TransferType**](TransferType.md) |  | 

## Methods

### NewTransferResponseToken

`func NewTransferResponseToken(accountId string, amount int32, creationTime time.Time, currency string, customerId string, id string, lastUpdatedTime time.Time, merchant Merchant, status TransferStatus, tenant string, cardDetails ExternalCardDetails, type_ TransferType, ) *TransferResponseToken`

NewTransferResponseToken instantiates a new TransferResponseToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResponseTokenWithDefaults

`func NewTransferResponseTokenWithDefaults() *TransferResponseToken`

NewTransferResponseTokenWithDefaults instantiates a new TransferResponseToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransferResponseToken) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferResponseToken) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferResponseToken) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *TransferResponseToken) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferResponseToken) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferResponseToken) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCreationTime

`func (o *TransferResponseToken) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *TransferResponseToken) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *TransferResponseToken) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *TransferResponseToken) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransferResponseToken) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransferResponseToken) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *TransferResponseToken) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TransferResponseToken) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TransferResponseToken) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *TransferResponseToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferResponseToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferResponseToken) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *TransferResponseToken) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *TransferResponseToken) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *TransferResponseToken) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMerchant

`func (o *TransferResponseToken) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransferResponseToken) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransferResponseToken) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.


### GetNetworkDeclineDetails

`func (o *TransferResponseToken) GetNetworkDeclineDetails() string`

GetNetworkDeclineDetails returns the NetworkDeclineDetails field if non-nil, zero value otherwise.

### GetNetworkDeclineDetailsOk

`func (o *TransferResponseToken) GetNetworkDeclineDetailsOk() (*string, bool)`

GetNetworkDeclineDetailsOk returns a tuple with the NetworkDeclineDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeclineDetails

`func (o *TransferResponseToken) SetNetworkDeclineDetails(v string)`

SetNetworkDeclineDetails sets NetworkDeclineDetails field to given value.

### HasNetworkDeclineDetails

`func (o *TransferResponseToken) HasNetworkDeclineDetails() bool`

HasNetworkDeclineDetails returns a boolean if a field has been set.

### GetReason

`func (o *TransferResponseToken) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransferResponseToken) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransferResponseToken) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransferResponseToken) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *TransferResponseToken) GetStatus() TransferStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferResponseToken) GetStatusOk() (*TransferStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferResponseToken) SetStatus(v TransferStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *TransferResponseToken) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransferResponseToken) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransferResponseToken) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionId

`func (o *TransferResponseToken) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransferResponseToken) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransferResponseToken) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TransferResponseToken) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetCardDetails

`func (o *TransferResponseToken) GetCardDetails() ExternalCardDetails`

GetCardDetails returns the CardDetails field if non-nil, zero value otherwise.

### GetCardDetailsOk

`func (o *TransferResponseToken) GetCardDetailsOk() (*ExternalCardDetails, bool)`

GetCardDetailsOk returns a tuple with the CardDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDetails

`func (o *TransferResponseToken) SetCardDetails(v ExternalCardDetails)`

SetCardDetails sets CardDetails field to given value.


### GetType

`func (o *TransferResponseToken) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferResponseToken) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferResponseToken) SetType(v TransferType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


