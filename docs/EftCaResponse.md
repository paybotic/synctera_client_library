# EftCaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transfer amount in cents | 
**CustomerId** | **string** | The UUID of the Synctera customer resource that is the originator of the transfer.  | 
**DcSign** | **string** | Debit or credit sign | 
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 
**TransactionCode** | **string** | The three digit transaction code that identifies the type of transaction. More information can be found here: https://www.payments.ca/sites/default/files/standard007eng.pdf.  | 
**DestinationAccountId** | **string** | The UUID of the Synctera account that is the destination of the transfer. For a transfer originated by the Synctera platform, this will be an external account resource, while for a transfer originated by the external account, this account will be an account resource.  | 
**DestinationAccountNumber** | **string** | The account number of the destination account. | 
**DestinationAccountOwnerName** | **string** | The account owner name of the destination account. | 
**EffectiveDate** | **string** | The effective date of the transaction once it gets posted | 
**History** | Pointer to [**[]Action**](Action.md) |  | [optional] 
**Id** | **string** | ID of the transfer | 
**IsSameDay** | **bool** | Send the same day (use only is_same_day without specific effective_date). | 
**NetworkStatus** | Pointer to **string** | The network status of the transfer. | [optional] 
**OriginatingAccountId** | **string** | The UUID of the Synctera account that is the origination of the transfer. For a transfer originated by the Synctera platform, this will be an account resource, while for a transfer originated by the external account, this will be an external account resource.  | 
**OriginatingAccountNumber** | **string** | The account number of the originating account. | 
**OriginatingAccountOwnerName** | **string** | The account owner name of the origination account. | 
**PostingDate** | Pointer to **string** | The posting date of the transaction once it gets posted | [optional] 
**ReferenceId** | **string** | The reference id of the transfer. | 
**Status** | [**EftCaStatus**](EftCaStatus.md) |  | 
**Subtype** | **string** | The subtype of the transfer | 
**TenantId** | **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | 
**TransactionId** | Pointer to **string** | The related transaction id of the transfer. | [optional] 

## Methods

### NewEftCaResponse

`func NewEftCaResponse(amount int64, customerId string, dcSign string, transactionCode string, destinationAccountId string, destinationAccountNumber string, destinationAccountOwnerName string, effectiveDate string, id string, isSameDay bool, originatingAccountId string, originatingAccountNumber string, originatingAccountOwnerName string, referenceId string, status EftCaStatus, subtype string, tenantId string, ) *EftCaResponse`

NewEftCaResponse instantiates a new EftCaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEftCaResponseWithDefaults

`func NewEftCaResponseWithDefaults() *EftCaResponse`

NewEftCaResponseWithDefaults instantiates a new EftCaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EftCaResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EftCaResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EftCaResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCustomerId

`func (o *EftCaResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *EftCaResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *EftCaResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDcSign

`func (o *EftCaResponse) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *EftCaResponse) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *EftCaResponse) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetSourceData

`func (o *EftCaResponse) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *EftCaResponse) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *EftCaResponse) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *EftCaResponse) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetTransactionCode

`func (o *EftCaResponse) GetTransactionCode() string`

GetTransactionCode returns the TransactionCode field if non-nil, zero value otherwise.

### GetTransactionCodeOk

`func (o *EftCaResponse) GetTransactionCodeOk() (*string, bool)`

GetTransactionCodeOk returns a tuple with the TransactionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCode

`func (o *EftCaResponse) SetTransactionCode(v string)`

SetTransactionCode sets TransactionCode field to given value.


### GetDestinationAccountId

`func (o *EftCaResponse) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *EftCaResponse) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *EftCaResponse) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.


### GetDestinationAccountNumber

`func (o *EftCaResponse) GetDestinationAccountNumber() string`

GetDestinationAccountNumber returns the DestinationAccountNumber field if non-nil, zero value otherwise.

### GetDestinationAccountNumberOk

`func (o *EftCaResponse) GetDestinationAccountNumberOk() (*string, bool)`

GetDestinationAccountNumberOk returns a tuple with the DestinationAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountNumber

`func (o *EftCaResponse) SetDestinationAccountNumber(v string)`

SetDestinationAccountNumber sets DestinationAccountNumber field to given value.


### GetDestinationAccountOwnerName

`func (o *EftCaResponse) GetDestinationAccountOwnerName() string`

GetDestinationAccountOwnerName returns the DestinationAccountOwnerName field if non-nil, zero value otherwise.

### GetDestinationAccountOwnerNameOk

`func (o *EftCaResponse) GetDestinationAccountOwnerNameOk() (*string, bool)`

GetDestinationAccountOwnerNameOk returns a tuple with the DestinationAccountOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountOwnerName

`func (o *EftCaResponse) SetDestinationAccountOwnerName(v string)`

SetDestinationAccountOwnerName sets DestinationAccountOwnerName field to given value.


### GetEffectiveDate

`func (o *EftCaResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *EftCaResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *EftCaResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetHistory

`func (o *EftCaResponse) GetHistory() []Action`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *EftCaResponse) GetHistoryOk() (*[]Action, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *EftCaResponse) SetHistory(v []Action)`

SetHistory sets History field to given value.

### HasHistory

`func (o *EftCaResponse) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetId

`func (o *EftCaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EftCaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EftCaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsSameDay

`func (o *EftCaResponse) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *EftCaResponse) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *EftCaResponse) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.


### GetNetworkStatus

`func (o *EftCaResponse) GetNetworkStatus() string`

GetNetworkStatus returns the NetworkStatus field if non-nil, zero value otherwise.

### GetNetworkStatusOk

`func (o *EftCaResponse) GetNetworkStatusOk() (*string, bool)`

GetNetworkStatusOk returns a tuple with the NetworkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatus

`func (o *EftCaResponse) SetNetworkStatus(v string)`

SetNetworkStatus sets NetworkStatus field to given value.

### HasNetworkStatus

`func (o *EftCaResponse) HasNetworkStatus() bool`

HasNetworkStatus returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *EftCaResponse) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *EftCaResponse) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *EftCaResponse) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetOriginatingAccountNumber

`func (o *EftCaResponse) GetOriginatingAccountNumber() string`

GetOriginatingAccountNumber returns the OriginatingAccountNumber field if non-nil, zero value otherwise.

### GetOriginatingAccountNumberOk

`func (o *EftCaResponse) GetOriginatingAccountNumberOk() (*string, bool)`

GetOriginatingAccountNumberOk returns a tuple with the OriginatingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountNumber

`func (o *EftCaResponse) SetOriginatingAccountNumber(v string)`

SetOriginatingAccountNumber sets OriginatingAccountNumber field to given value.


### GetOriginatingAccountOwnerName

`func (o *EftCaResponse) GetOriginatingAccountOwnerName() string`

GetOriginatingAccountOwnerName returns the OriginatingAccountOwnerName field if non-nil, zero value otherwise.

### GetOriginatingAccountOwnerNameOk

`func (o *EftCaResponse) GetOriginatingAccountOwnerNameOk() (*string, bool)`

GetOriginatingAccountOwnerNameOk returns a tuple with the OriginatingAccountOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountOwnerName

`func (o *EftCaResponse) SetOriginatingAccountOwnerName(v string)`

SetOriginatingAccountOwnerName sets OriginatingAccountOwnerName field to given value.


### GetPostingDate

`func (o *EftCaResponse) GetPostingDate() string`

GetPostingDate returns the PostingDate field if non-nil, zero value otherwise.

### GetPostingDateOk

`func (o *EftCaResponse) GetPostingDateOk() (*string, bool)`

GetPostingDateOk returns a tuple with the PostingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingDate

`func (o *EftCaResponse) SetPostingDate(v string)`

SetPostingDate sets PostingDate field to given value.

### HasPostingDate

`func (o *EftCaResponse) HasPostingDate() bool`

HasPostingDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *EftCaResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *EftCaResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *EftCaResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetStatus

`func (o *EftCaResponse) GetStatus() EftCaStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EftCaResponse) GetStatusOk() (*EftCaStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EftCaResponse) SetStatus(v EftCaStatus)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *EftCaResponse) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *EftCaResponse) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *EftCaResponse) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetTenantId

`func (o *EftCaResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EftCaResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EftCaResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTransactionId

`func (o *EftCaResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *EftCaResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *EftCaResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *EftCaResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


