# SyncteraPayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transfer amount in cents | 
**ConfigurationId** | Pointer to **string** | The ID of the Outgoing Synctera Pay configuration that will be used to match the Synctera Pay transfer to the appropriate configuration. Only required if the transfer is not automatically matched to a configuration.  | [optional] 
**Currency** | **string** | The currency of the transfer in ISO 4217 format | 
**DcSign** | **string** | The debit/credit sign of the transfer. This is a legacy field for backward compatibility. For outgoing synctera pay, customer accounts are always debited. For incoming synctera pay, customer accounts are always credited. The direction of the transfer is determined by the API route that is used to create the transfer. For outgoing synctera pay, the dc_sign must be specified and it must always be &#39;CREDIT&#39;.  | 
**Direction** | [**SyncteraPayDirection**](SyncteraPayDirection.md) |  | 
**ExchangeDetails** | Pointer to [**ExchangeDetails**](ExchangeDetails.md) |  | [optional] 
**FinalExternalAccountId** | Pointer to **string** | The ID of the final external account that will be the receiver of the Outgoing Synctera Pay transfer. Required unless previously agreed with Synctera.  | [optional] 
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 
**Subtype** | [**SyncteraPaySubtype**](SyncteraPaySubtype.md) |  | 
**SyncteraPayNetwork** | **string** | The network of the transfer.  | 
**SyncteraPayVendorId** | Pointer to **string** | The ID of the vendor that will be used to process the Synctera Pay transfer. | [optional] 
**CustomerId** | Pointer to **string** | The UUID of the Synctera customer resource that is the originator of the transfer.  | [optional] 
**DestinationAccountId** | Pointer to **string** | The UUID of the Synctera account that is the destination of the transfer. For a transfer originated by the Synctera platform, this will be an external account resource, while for a transfer originated by the external account, this account will be an account resource.  | [optional] 
**DestinationAccountOwnerName** | Pointer to **string** | The account owner name of the destination account. | [optional] 
**EffectiveDate** | **string** | The effective date of the transaction once it gets posted | 
**Failed** | Pointer to **bool** | Whether the transfer failed or not. | [optional] 
**History** | Pointer to [**[]Action**](Action.md) |  | [optional] 
**Id** | **string** | ID of the transfer | 
**IsSameDay** | **bool** | Send the same day (use only is_same_day without specific effective_date). | 
**NetworkStatus** | Pointer to [**SyncteraPayNetworkStatus**](SyncteraPayNetworkStatus.md) |  | [optional] 
**OriginalReferenceId** | Pointer to **string** | The original reference id of the transfer if it&#39;s a return. | [optional] 
**OriginatingAccountId** | Pointer to **string** | The UUID of the Synctera account that is the origination of the transfer. For a transfer originated by the Synctera platform, this will be an account resource, while for a transfer originated by the external account, this will be an external account resource.  | [optional] 
**OriginatingAccountOwnerName** | Pointer to **string** | The account owner name of the origination account. | [optional] 
**PostingDate** | Pointer to **string** | The posting date of the transaction once it gets posted | [optional] 
**ReferenceId** | Pointer to **string** | The reference id of the transfer. | [optional] 
**Status** | [**SyncteraPayStatus**](SyncteraPayStatus.md) |  | 
**Suspended** | Pointer to **bool** | Whether the transfer is suspended or not. | [optional] 
**TenantId** | **string** | The id of the tenant containing the resource.  | 
**TransactionId** | Pointer to **string** | The related transaction id of the transfer. | [optional] 

## Methods

### NewSyncteraPayResponse

`func NewSyncteraPayResponse(amount int64, currency string, dcSign string, direction SyncteraPayDirection, subtype SyncteraPaySubtype, syncteraPayNetwork string, effectiveDate string, id string, isSameDay bool, status SyncteraPayStatus, tenantId string, ) *SyncteraPayResponse`

NewSyncteraPayResponse instantiates a new SyncteraPayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayResponseWithDefaults

`func NewSyncteraPayResponseWithDefaults() *SyncteraPayResponse`

NewSyncteraPayResponseWithDefaults instantiates a new SyncteraPayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SyncteraPayResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SyncteraPayResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SyncteraPayResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetConfigurationId

`func (o *SyncteraPayResponse) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *SyncteraPayResponse) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *SyncteraPayResponse) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *SyncteraPayResponse) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetCurrency

`func (o *SyncteraPayResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SyncteraPayResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SyncteraPayResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDcSign

`func (o *SyncteraPayResponse) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *SyncteraPayResponse) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *SyncteraPayResponse) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetDirection

`func (o *SyncteraPayResponse) GetDirection() SyncteraPayDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SyncteraPayResponse) GetDirectionOk() (*SyncteraPayDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SyncteraPayResponse) SetDirection(v SyncteraPayDirection)`

SetDirection sets Direction field to given value.


### GetExchangeDetails

`func (o *SyncteraPayResponse) GetExchangeDetails() ExchangeDetails`

GetExchangeDetails returns the ExchangeDetails field if non-nil, zero value otherwise.

### GetExchangeDetailsOk

`func (o *SyncteraPayResponse) GetExchangeDetailsOk() (*ExchangeDetails, bool)`

GetExchangeDetailsOk returns a tuple with the ExchangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDetails

`func (o *SyncteraPayResponse) SetExchangeDetails(v ExchangeDetails)`

SetExchangeDetails sets ExchangeDetails field to given value.

### HasExchangeDetails

`func (o *SyncteraPayResponse) HasExchangeDetails() bool`

HasExchangeDetails returns a boolean if a field has been set.

### GetFinalExternalAccountId

`func (o *SyncteraPayResponse) GetFinalExternalAccountId() string`

GetFinalExternalAccountId returns the FinalExternalAccountId field if non-nil, zero value otherwise.

### GetFinalExternalAccountIdOk

`func (o *SyncteraPayResponse) GetFinalExternalAccountIdOk() (*string, bool)`

GetFinalExternalAccountIdOk returns a tuple with the FinalExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalExternalAccountId

`func (o *SyncteraPayResponse) SetFinalExternalAccountId(v string)`

SetFinalExternalAccountId sets FinalExternalAccountId field to given value.

### HasFinalExternalAccountId

`func (o *SyncteraPayResponse) HasFinalExternalAccountId() bool`

HasFinalExternalAccountId returns a boolean if a field has been set.

### GetSourceData

`func (o *SyncteraPayResponse) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *SyncteraPayResponse) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *SyncteraPayResponse) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *SyncteraPayResponse) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSubtype

`func (o *SyncteraPayResponse) GetSubtype() SyncteraPaySubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SyncteraPayResponse) GetSubtypeOk() (*SyncteraPaySubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SyncteraPayResponse) SetSubtype(v SyncteraPaySubtype)`

SetSubtype sets Subtype field to given value.


### GetSyncteraPayNetwork

`func (o *SyncteraPayResponse) GetSyncteraPayNetwork() string`

GetSyncteraPayNetwork returns the SyncteraPayNetwork field if non-nil, zero value otherwise.

### GetSyncteraPayNetworkOk

`func (o *SyncteraPayResponse) GetSyncteraPayNetworkOk() (*string, bool)`

GetSyncteraPayNetworkOk returns a tuple with the SyncteraPayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncteraPayNetwork

`func (o *SyncteraPayResponse) SetSyncteraPayNetwork(v string)`

SetSyncteraPayNetwork sets SyncteraPayNetwork field to given value.


### GetSyncteraPayVendorId

`func (o *SyncteraPayResponse) GetSyncteraPayVendorId() string`

GetSyncteraPayVendorId returns the SyncteraPayVendorId field if non-nil, zero value otherwise.

### GetSyncteraPayVendorIdOk

`func (o *SyncteraPayResponse) GetSyncteraPayVendorIdOk() (*string, bool)`

GetSyncteraPayVendorIdOk returns a tuple with the SyncteraPayVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncteraPayVendorId

`func (o *SyncteraPayResponse) SetSyncteraPayVendorId(v string)`

SetSyncteraPayVendorId sets SyncteraPayVendorId field to given value.

### HasSyncteraPayVendorId

`func (o *SyncteraPayResponse) HasSyncteraPayVendorId() bool`

HasSyncteraPayVendorId returns a boolean if a field has been set.

### GetCustomerId

`func (o *SyncteraPayResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SyncteraPayResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SyncteraPayResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SyncteraPayResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDestinationAccountId

`func (o *SyncteraPayResponse) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *SyncteraPayResponse) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *SyncteraPayResponse) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.

### HasDestinationAccountId

`func (o *SyncteraPayResponse) HasDestinationAccountId() bool`

HasDestinationAccountId returns a boolean if a field has been set.

### GetDestinationAccountOwnerName

`func (o *SyncteraPayResponse) GetDestinationAccountOwnerName() string`

GetDestinationAccountOwnerName returns the DestinationAccountOwnerName field if non-nil, zero value otherwise.

### GetDestinationAccountOwnerNameOk

`func (o *SyncteraPayResponse) GetDestinationAccountOwnerNameOk() (*string, bool)`

GetDestinationAccountOwnerNameOk returns a tuple with the DestinationAccountOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountOwnerName

`func (o *SyncteraPayResponse) SetDestinationAccountOwnerName(v string)`

SetDestinationAccountOwnerName sets DestinationAccountOwnerName field to given value.

### HasDestinationAccountOwnerName

`func (o *SyncteraPayResponse) HasDestinationAccountOwnerName() bool`

HasDestinationAccountOwnerName returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *SyncteraPayResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *SyncteraPayResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *SyncteraPayResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetFailed

`func (o *SyncteraPayResponse) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *SyncteraPayResponse) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *SyncteraPayResponse) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *SyncteraPayResponse) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetHistory

`func (o *SyncteraPayResponse) GetHistory() []Action`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *SyncteraPayResponse) GetHistoryOk() (*[]Action, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *SyncteraPayResponse) SetHistory(v []Action)`

SetHistory sets History field to given value.

### HasHistory

`func (o *SyncteraPayResponse) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetId

`func (o *SyncteraPayResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncteraPayResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncteraPayResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsSameDay

`func (o *SyncteraPayResponse) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *SyncteraPayResponse) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *SyncteraPayResponse) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.


### GetNetworkStatus

`func (o *SyncteraPayResponse) GetNetworkStatus() SyncteraPayNetworkStatus`

GetNetworkStatus returns the NetworkStatus field if non-nil, zero value otherwise.

### GetNetworkStatusOk

`func (o *SyncteraPayResponse) GetNetworkStatusOk() (*SyncteraPayNetworkStatus, bool)`

GetNetworkStatusOk returns a tuple with the NetworkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatus

`func (o *SyncteraPayResponse) SetNetworkStatus(v SyncteraPayNetworkStatus)`

SetNetworkStatus sets NetworkStatus field to given value.

### HasNetworkStatus

`func (o *SyncteraPayResponse) HasNetworkStatus() bool`

HasNetworkStatus returns a boolean if a field has been set.

### GetOriginalReferenceId

`func (o *SyncteraPayResponse) GetOriginalReferenceId() string`

GetOriginalReferenceId returns the OriginalReferenceId field if non-nil, zero value otherwise.

### GetOriginalReferenceIdOk

`func (o *SyncteraPayResponse) GetOriginalReferenceIdOk() (*string, bool)`

GetOriginalReferenceIdOk returns a tuple with the OriginalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceId

`func (o *SyncteraPayResponse) SetOriginalReferenceId(v string)`

SetOriginalReferenceId sets OriginalReferenceId field to given value.

### HasOriginalReferenceId

`func (o *SyncteraPayResponse) HasOriginalReferenceId() bool`

HasOriginalReferenceId returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *SyncteraPayResponse) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *SyncteraPayResponse) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *SyncteraPayResponse) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.

### HasOriginatingAccountId

`func (o *SyncteraPayResponse) HasOriginatingAccountId() bool`

HasOriginatingAccountId returns a boolean if a field has been set.

### GetOriginatingAccountOwnerName

`func (o *SyncteraPayResponse) GetOriginatingAccountOwnerName() string`

GetOriginatingAccountOwnerName returns the OriginatingAccountOwnerName field if non-nil, zero value otherwise.

### GetOriginatingAccountOwnerNameOk

`func (o *SyncteraPayResponse) GetOriginatingAccountOwnerNameOk() (*string, bool)`

GetOriginatingAccountOwnerNameOk returns a tuple with the OriginatingAccountOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountOwnerName

`func (o *SyncteraPayResponse) SetOriginatingAccountOwnerName(v string)`

SetOriginatingAccountOwnerName sets OriginatingAccountOwnerName field to given value.

### HasOriginatingAccountOwnerName

`func (o *SyncteraPayResponse) HasOriginatingAccountOwnerName() bool`

HasOriginatingAccountOwnerName returns a boolean if a field has been set.

### GetPostingDate

`func (o *SyncteraPayResponse) GetPostingDate() string`

GetPostingDate returns the PostingDate field if non-nil, zero value otherwise.

### GetPostingDateOk

`func (o *SyncteraPayResponse) GetPostingDateOk() (*string, bool)`

GetPostingDateOk returns a tuple with the PostingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingDate

`func (o *SyncteraPayResponse) SetPostingDate(v string)`

SetPostingDate sets PostingDate field to given value.

### HasPostingDate

`func (o *SyncteraPayResponse) HasPostingDate() bool`

HasPostingDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *SyncteraPayResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SyncteraPayResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SyncteraPayResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *SyncteraPayResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetStatus

`func (o *SyncteraPayResponse) GetStatus() SyncteraPayStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncteraPayResponse) GetStatusOk() (*SyncteraPayStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncteraPayResponse) SetStatus(v SyncteraPayStatus)`

SetStatus sets Status field to given value.


### GetSuspended

`func (o *SyncteraPayResponse) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *SyncteraPayResponse) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *SyncteraPayResponse) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *SyncteraPayResponse) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetTenantId

`func (o *SyncteraPayResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SyncteraPayResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SyncteraPayResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTransactionId

`func (o *SyncteraPayResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *SyncteraPayResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *SyncteraPayResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *SyncteraPayResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


