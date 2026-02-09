# TransferResponse

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
**OriginatingCustomerId** | Pointer to **string** | For person-to-person PUSH transactions this is the customer_id of the sender who must have privileges to access funds in the originating account in order to send funds to the recipient cardholder | [optional] 
**CardDetails** | [**ExternalCardDetails**](ExternalCardDetails.md) |  | 

## Methods

### NewTransferResponse

`func NewTransferResponse(accountId string, amount int32, creationTime time.Time, currency string, customerId string, id string, lastUpdatedTime time.Time, merchant Merchant, status TransferStatus, tenant string, externalCardId string, type_ TransferType, cardDetails ExternalCardDetails, ) *TransferResponse`

NewTransferResponse instantiates a new TransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResponseWithDefaults

`func NewTransferResponseWithDefaults() *TransferResponse`

NewTransferResponseWithDefaults instantiates a new TransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransferResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *TransferResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCreationTime

`func (o *TransferResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *TransferResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *TransferResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *TransferResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransferResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransferResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *TransferResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TransferResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TransferResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *TransferResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *TransferResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *TransferResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *TransferResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMerchant

`func (o *TransferResponse) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransferResponse) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransferResponse) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.


### GetNetworkDeclineDetails

`func (o *TransferResponse) GetNetworkDeclineDetails() string`

GetNetworkDeclineDetails returns the NetworkDeclineDetails field if non-nil, zero value otherwise.

### GetNetworkDeclineDetailsOk

`func (o *TransferResponse) GetNetworkDeclineDetailsOk() (*string, bool)`

GetNetworkDeclineDetailsOk returns a tuple with the NetworkDeclineDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeclineDetails

`func (o *TransferResponse) SetNetworkDeclineDetails(v string)`

SetNetworkDeclineDetails sets NetworkDeclineDetails field to given value.

### HasNetworkDeclineDetails

`func (o *TransferResponse) HasNetworkDeclineDetails() bool`

HasNetworkDeclineDetails returns a boolean if a field has been set.

### GetReason

`func (o *TransferResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransferResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransferResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransferResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *TransferResponse) GetStatus() TransferStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferResponse) GetStatusOk() (*TransferStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferResponse) SetStatus(v TransferStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *TransferResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransferResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransferResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionId

`func (o *TransferResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransferResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransferResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TransferResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetExternalCardId

`func (o *TransferResponse) GetExternalCardId() string`

GetExternalCardId returns the ExternalCardId field if non-nil, zero value otherwise.

### GetExternalCardIdOk

`func (o *TransferResponse) GetExternalCardIdOk() (*string, bool)`

GetExternalCardIdOk returns a tuple with the ExternalCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCardId

`func (o *TransferResponse) SetExternalCardId(v string)`

SetExternalCardId sets ExternalCardId field to given value.


### GetThreeDsId

`func (o *TransferResponse) GetThreeDsId() string`

GetThreeDsId returns the ThreeDsId field if non-nil, zero value otherwise.

### GetThreeDsIdOk

`func (o *TransferResponse) GetThreeDsIdOk() (*string, bool)`

GetThreeDsIdOk returns a tuple with the ThreeDsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsId

`func (o *TransferResponse) SetThreeDsId(v string)`

SetThreeDsId sets ThreeDsId field to given value.

### HasThreeDsId

`func (o *TransferResponse) HasThreeDsId() bool`

HasThreeDsId returns a boolean if a field has been set.

### GetType

`func (o *TransferResponse) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferResponse) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferResponse) SetType(v TransferType)`

SetType sets Type field to given value.


### GetOriginatingCustomerId

`func (o *TransferResponse) GetOriginatingCustomerId() string`

GetOriginatingCustomerId returns the OriginatingCustomerId field if non-nil, zero value otherwise.

### GetOriginatingCustomerIdOk

`func (o *TransferResponse) GetOriginatingCustomerIdOk() (*string, bool)`

GetOriginatingCustomerIdOk returns a tuple with the OriginatingCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingCustomerId

`func (o *TransferResponse) SetOriginatingCustomerId(v string)`

SetOriginatingCustomerId sets OriginatingCustomerId field to given value.

### HasOriginatingCustomerId

`func (o *TransferResponse) HasOriginatingCustomerId() bool`

HasOriginatingCustomerId returns a boolean if a field has been set.

### GetCardDetails

`func (o *TransferResponse) GetCardDetails() ExternalCardDetails`

GetCardDetails returns the CardDetails field if non-nil, zero value otherwise.

### GetCardDetailsOk

`func (o *TransferResponse) GetCardDetailsOk() (*ExternalCardDetails, bool)`

GetCardDetailsOk returns a tuple with the CardDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDetails

`func (o *TransferResponse) SetCardDetails(v ExternalCardDetails)`

SetCardDetails sets CardDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


