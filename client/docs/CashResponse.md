# CashResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transfer amount in cents | 
**DcSign** | **string** | Debit or credit sign | 
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 
**ClientBankAccount** | Pointer to **string** | The bank account of the client. | [optional] 
**ClientName** | Pointer to **string** | The name of the client. | [optional] 
**CustomerId** | Pointer to **string** | The UUID of the Synctera customer resource that is the originator of the transfer.  | [optional] 
**DestinationAccountId** | Pointer to **string** | The UUID of the Synctera account that is the destination of the transfer. For a transfer originated by the Synctera platform, this will be an external account resource, while for a transfer originated by the external account, this account will be an account resource.  | [optional] 
**DestinationAccountNumber** | Pointer to **string** | The account number of the destination account. | [optional] 
**DestinationAccountOwnerName** | Pointer to **string** | The account owner name of the destination account. | [optional] 
**EffectiveDate** | **string** | The effective date of the transaction once it gets posted | 
**Failed** | Pointer to **bool** | Whether the transfer failed or not. | [optional] 
**History** | Pointer to [**[]Action**](Action.md) |  | [optional] 
**Id** | **string** | ID of the transfer | 
**IsSameDay** | **bool** | Send the same day (use only is_same_day without specific effective_date). | 
**NetworkStatus** | Pointer to [**CashNetworkStatus**](CashNetworkStatus.md) |  | [optional] 
**OriginalReferenceId** | Pointer to **string** | The original reference id of the transfer if it&#39;s a return. | [optional] 
**OriginatingAccountId** | Pointer to **string** | The UUID of the Synctera account that is the origination of the transfer. For a transfer originated by the Synctera platform, this will be an account resource, while for a transfer originated by the external account, this will be an external account resource.  | [optional] 
**OriginatingAccountNumber** | Pointer to **string** | The account number of the originating account. | [optional] 
**OriginatingAccountOwnerName** | Pointer to **string** | The account owner name of the origination account. | [optional] 
**PostingDate** | Pointer to **string** | The posting date of the transaction once it gets posted | [optional] 
**ReferenceId** | **string** | The reference id of the transfer. | 
**Status** | [**CashStatus**](CashStatus.md) |  | 
**Subtype** | **string** | The subtype of the transfer | 
**Suspended** | Pointer to **bool** | Whether the transfer is suspended or not. | [optional] 
**TenantId** | **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | 
**TransactionId** | Pointer to **string** | The related transaction id of the transfer. | [optional] 

## Methods

### NewCashResponse

`func NewCashResponse(amount int64, dcSign string, effectiveDate string, id string, isSameDay bool, referenceId string, status CashStatus, subtype string, tenantId string, ) *CashResponse`

NewCashResponse instantiates a new CashResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashResponseWithDefaults

`func NewCashResponseWithDefaults() *CashResponse`

NewCashResponseWithDefaults instantiates a new CashResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CashResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDcSign

`func (o *CashResponse) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *CashResponse) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *CashResponse) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetSourceData

`func (o *CashResponse) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *CashResponse) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *CashResponse) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *CashResponse) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetClientBankAccount

`func (o *CashResponse) GetClientBankAccount() string`

GetClientBankAccount returns the ClientBankAccount field if non-nil, zero value otherwise.

### GetClientBankAccountOk

`func (o *CashResponse) GetClientBankAccountOk() (*string, bool)`

GetClientBankAccountOk returns a tuple with the ClientBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBankAccount

`func (o *CashResponse) SetClientBankAccount(v string)`

SetClientBankAccount sets ClientBankAccount field to given value.

### HasClientBankAccount

`func (o *CashResponse) HasClientBankAccount() bool`

HasClientBankAccount returns a boolean if a field has been set.

### GetClientName

`func (o *CashResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *CashResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *CashResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *CashResponse) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetCustomerId

`func (o *CashResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CashResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CashResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CashResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDestinationAccountId

`func (o *CashResponse) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *CashResponse) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *CashResponse) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.

### HasDestinationAccountId

`func (o *CashResponse) HasDestinationAccountId() bool`

HasDestinationAccountId returns a boolean if a field has been set.

### GetDestinationAccountNumber

`func (o *CashResponse) GetDestinationAccountNumber() string`

GetDestinationAccountNumber returns the DestinationAccountNumber field if non-nil, zero value otherwise.

### GetDestinationAccountNumberOk

`func (o *CashResponse) GetDestinationAccountNumberOk() (*string, bool)`

GetDestinationAccountNumberOk returns a tuple with the DestinationAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountNumber

`func (o *CashResponse) SetDestinationAccountNumber(v string)`

SetDestinationAccountNumber sets DestinationAccountNumber field to given value.

### HasDestinationAccountNumber

`func (o *CashResponse) HasDestinationAccountNumber() bool`

HasDestinationAccountNumber returns a boolean if a field has been set.

### GetDestinationAccountOwnerName

`func (o *CashResponse) GetDestinationAccountOwnerName() string`

GetDestinationAccountOwnerName returns the DestinationAccountOwnerName field if non-nil, zero value otherwise.

### GetDestinationAccountOwnerNameOk

`func (o *CashResponse) GetDestinationAccountOwnerNameOk() (*string, bool)`

GetDestinationAccountOwnerNameOk returns a tuple with the DestinationAccountOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountOwnerName

`func (o *CashResponse) SetDestinationAccountOwnerName(v string)`

SetDestinationAccountOwnerName sets DestinationAccountOwnerName field to given value.

### HasDestinationAccountOwnerName

`func (o *CashResponse) HasDestinationAccountOwnerName() bool`

HasDestinationAccountOwnerName returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *CashResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *CashResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *CashResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetFailed

`func (o *CashResponse) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *CashResponse) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *CashResponse) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *CashResponse) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetHistory

`func (o *CashResponse) GetHistory() []Action`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *CashResponse) GetHistoryOk() (*[]Action, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *CashResponse) SetHistory(v []Action)`

SetHistory sets History field to given value.

### HasHistory

`func (o *CashResponse) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetId

`func (o *CashResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CashResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CashResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsSameDay

`func (o *CashResponse) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *CashResponse) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *CashResponse) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.


### GetNetworkStatus

`func (o *CashResponse) GetNetworkStatus() CashNetworkStatus`

GetNetworkStatus returns the NetworkStatus field if non-nil, zero value otherwise.

### GetNetworkStatusOk

`func (o *CashResponse) GetNetworkStatusOk() (*CashNetworkStatus, bool)`

GetNetworkStatusOk returns a tuple with the NetworkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatus

`func (o *CashResponse) SetNetworkStatus(v CashNetworkStatus)`

SetNetworkStatus sets NetworkStatus field to given value.

### HasNetworkStatus

`func (o *CashResponse) HasNetworkStatus() bool`

HasNetworkStatus returns a boolean if a field has been set.

### GetOriginalReferenceId

`func (o *CashResponse) GetOriginalReferenceId() string`

GetOriginalReferenceId returns the OriginalReferenceId field if non-nil, zero value otherwise.

### GetOriginalReferenceIdOk

`func (o *CashResponse) GetOriginalReferenceIdOk() (*string, bool)`

GetOriginalReferenceIdOk returns a tuple with the OriginalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceId

`func (o *CashResponse) SetOriginalReferenceId(v string)`

SetOriginalReferenceId sets OriginalReferenceId field to given value.

### HasOriginalReferenceId

`func (o *CashResponse) HasOriginalReferenceId() bool`

HasOriginalReferenceId returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *CashResponse) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *CashResponse) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *CashResponse) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.

### HasOriginatingAccountId

`func (o *CashResponse) HasOriginatingAccountId() bool`

HasOriginatingAccountId returns a boolean if a field has been set.

### GetOriginatingAccountNumber

`func (o *CashResponse) GetOriginatingAccountNumber() string`

GetOriginatingAccountNumber returns the OriginatingAccountNumber field if non-nil, zero value otherwise.

### GetOriginatingAccountNumberOk

`func (o *CashResponse) GetOriginatingAccountNumberOk() (*string, bool)`

GetOriginatingAccountNumberOk returns a tuple with the OriginatingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountNumber

`func (o *CashResponse) SetOriginatingAccountNumber(v string)`

SetOriginatingAccountNumber sets OriginatingAccountNumber field to given value.

### HasOriginatingAccountNumber

`func (o *CashResponse) HasOriginatingAccountNumber() bool`

HasOriginatingAccountNumber returns a boolean if a field has been set.

### GetOriginatingAccountOwnerName

`func (o *CashResponse) GetOriginatingAccountOwnerName() string`

GetOriginatingAccountOwnerName returns the OriginatingAccountOwnerName field if non-nil, zero value otherwise.

### GetOriginatingAccountOwnerNameOk

`func (o *CashResponse) GetOriginatingAccountOwnerNameOk() (*string, bool)`

GetOriginatingAccountOwnerNameOk returns a tuple with the OriginatingAccountOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountOwnerName

`func (o *CashResponse) SetOriginatingAccountOwnerName(v string)`

SetOriginatingAccountOwnerName sets OriginatingAccountOwnerName field to given value.

### HasOriginatingAccountOwnerName

`func (o *CashResponse) HasOriginatingAccountOwnerName() bool`

HasOriginatingAccountOwnerName returns a boolean if a field has been set.

### GetPostingDate

`func (o *CashResponse) GetPostingDate() string`

GetPostingDate returns the PostingDate field if non-nil, zero value otherwise.

### GetPostingDateOk

`func (o *CashResponse) GetPostingDateOk() (*string, bool)`

GetPostingDateOk returns a tuple with the PostingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingDate

`func (o *CashResponse) SetPostingDate(v string)`

SetPostingDate sets PostingDate field to given value.

### HasPostingDate

`func (o *CashResponse) HasPostingDate() bool`

HasPostingDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *CashResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CashResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CashResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetStatus

`func (o *CashResponse) GetStatus() CashStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CashResponse) GetStatusOk() (*CashStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CashResponse) SetStatus(v CashStatus)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *CashResponse) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CashResponse) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CashResponse) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetSuspended

`func (o *CashResponse) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *CashResponse) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *CashResponse) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *CashResponse) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetTenantId

`func (o *CashResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CashResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CashResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTransactionId

`func (o *CashResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CashResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CashResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CashResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


