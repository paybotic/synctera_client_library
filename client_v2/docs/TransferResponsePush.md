# TransferResponsePush

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
**OriginatingCustomerId** | Pointer to **string** | For person-to-person PUSH transactions this is the customer_id of the sender who must have privileges to access funds in the originating account in order to send funds to the recipient cardholder | [optional] 
**Type** | [**TransferType**](TransferType.md) |  | 

## Methods

### NewTransferResponsePush

`func NewTransferResponsePush(accountId string, amount int32, creationTime time.Time, currency string, customerId string, id string, lastUpdatedTime time.Time, merchant Merchant, status TransferStatus, tenant string, externalCardId string, type_ TransferType, ) *TransferResponsePush`

NewTransferResponsePush instantiates a new TransferResponsePush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResponsePushWithDefaults

`func NewTransferResponsePushWithDefaults() *TransferResponsePush`

NewTransferResponsePushWithDefaults instantiates a new TransferResponsePush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransferResponsePush) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferResponsePush) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferResponsePush) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *TransferResponsePush) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferResponsePush) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferResponsePush) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCreationTime

`func (o *TransferResponsePush) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *TransferResponsePush) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *TransferResponsePush) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *TransferResponsePush) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransferResponsePush) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransferResponsePush) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *TransferResponsePush) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TransferResponsePush) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TransferResponsePush) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *TransferResponsePush) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferResponsePush) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferResponsePush) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *TransferResponsePush) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *TransferResponsePush) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *TransferResponsePush) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMerchant

`func (o *TransferResponsePush) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransferResponsePush) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransferResponsePush) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.


### GetNetworkDeclineDetails

`func (o *TransferResponsePush) GetNetworkDeclineDetails() string`

GetNetworkDeclineDetails returns the NetworkDeclineDetails field if non-nil, zero value otherwise.

### GetNetworkDeclineDetailsOk

`func (o *TransferResponsePush) GetNetworkDeclineDetailsOk() (*string, bool)`

GetNetworkDeclineDetailsOk returns a tuple with the NetworkDeclineDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeclineDetails

`func (o *TransferResponsePush) SetNetworkDeclineDetails(v string)`

SetNetworkDeclineDetails sets NetworkDeclineDetails field to given value.

### HasNetworkDeclineDetails

`func (o *TransferResponsePush) HasNetworkDeclineDetails() bool`

HasNetworkDeclineDetails returns a boolean if a field has been set.

### GetReason

`func (o *TransferResponsePush) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransferResponsePush) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransferResponsePush) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransferResponsePush) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *TransferResponsePush) GetStatus() TransferStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferResponsePush) GetStatusOk() (*TransferStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferResponsePush) SetStatus(v TransferStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *TransferResponsePush) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransferResponsePush) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransferResponsePush) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionId

`func (o *TransferResponsePush) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransferResponsePush) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransferResponsePush) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TransferResponsePush) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetExternalCardId

`func (o *TransferResponsePush) GetExternalCardId() string`

GetExternalCardId returns the ExternalCardId field if non-nil, zero value otherwise.

### GetExternalCardIdOk

`func (o *TransferResponsePush) GetExternalCardIdOk() (*string, bool)`

GetExternalCardIdOk returns a tuple with the ExternalCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCardId

`func (o *TransferResponsePush) SetExternalCardId(v string)`

SetExternalCardId sets ExternalCardId field to given value.


### GetOriginatingCustomerId

`func (o *TransferResponsePush) GetOriginatingCustomerId() string`

GetOriginatingCustomerId returns the OriginatingCustomerId field if non-nil, zero value otherwise.

### GetOriginatingCustomerIdOk

`func (o *TransferResponsePush) GetOriginatingCustomerIdOk() (*string, bool)`

GetOriginatingCustomerIdOk returns a tuple with the OriginatingCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingCustomerId

`func (o *TransferResponsePush) SetOriginatingCustomerId(v string)`

SetOriginatingCustomerId sets OriginatingCustomerId field to given value.

### HasOriginatingCustomerId

`func (o *TransferResponsePush) HasOriginatingCustomerId() bool`

HasOriginatingCustomerId returns a boolean if a field has been set.

### GetType

`func (o *TransferResponsePush) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferResponsePush) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferResponsePush) SetType(v TransferType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


