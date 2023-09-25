# InternalTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount (in cents) to transfer from originating account to receiving account. | 
**CaptureMode** | Pointer to **string** | Controls when the transfer will take effect. A value of &#x60;IMMEDIATE&#x60; (the default) means that the transfer will be completed immediately. A value of &#x60;MANUAL&#x60; means that the transaction will remain in a \&quot;pending\&quot; state until explicitly completed or cancelled (or the auth expires). | [optional] [default to "IMMEDIATE"]
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**ExpiresAt** | Pointer to **time.Time** | When &#x60;capture_mode&#x60; is &#x60;MANUAL&#x60;, this field describes when the pending transaction should expire. | [optional] 
**FinalCustomerId** | Pointer to **string** | The customer id of the international customer that receives the final remittance transfer (required for remittance payments). | [optional] 
**Memo** | Pointer to **string** | A short note to the recipient | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value metadata to associate with the transaction | [optional] 
**OriginatingAccountAlias** | Pointer to **string** | An alias representing a GL account to debit. This is alternative to specifying by account id | [optional] 
**OriginatingAccountCustomerId** | Pointer to **string** | The customer id of the owner of the originating account. | [optional] 
**OriginatingAccountId** | Pointer to **string** | The UUID of the account being debited | [optional] 
**ReceivingAccountAlias** | Pointer to **string** | An alias representing a GL account to credit. This is an alternative to specifying by account id | [optional] 
**ReceivingAccountCustomerId** | Pointer to **string** | The customer id of the owner of the receiving account. | [optional] 
**ReceivingAccountId** | Pointer to **string** | The UUID of the account being credited | [optional] 
**Tenant** | Pointer to **string** | The tenant associated with this resource, in the form \&quot;&lt;bankid&gt;_&lt;partnerid&gt;\&quot; | [optional] 
**Type** | **string** | The desired transaction type to use for this transfer | 
**Id** | **string** | The transaction id associated with the transfer | 
**Status** | **string** | The status of the internal transfer auth. A value of &#x60;PENDING&#x60; indicates that the funds have been reserved and the transaction is ready to be either completed or canceled. A value of &#x60;COMPLETE&#x60; indicates the funds have been successfully moved and no more action can be performed. A value of &#x60;CANCELED&#x60; or &#x60;EXPIRED&#x60; means that the transaction has rolled back and the funds have been returned to the originating account, either by explicitly canceling via the API, or due to the expiry time having passed. | 

## Methods

### NewInternalTransferResponse

`func NewInternalTransferResponse(amount int64, currency string, type_ string, id string, status string, ) *InternalTransferResponse`

NewInternalTransferResponse instantiates a new InternalTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalTransferResponseWithDefaults

`func NewInternalTransferResponseWithDefaults() *InternalTransferResponse`

NewInternalTransferResponseWithDefaults instantiates a new InternalTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalTransferResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalTransferResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalTransferResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCaptureMode

`func (o *InternalTransferResponse) GetCaptureMode() string`

GetCaptureMode returns the CaptureMode field if non-nil, zero value otherwise.

### GetCaptureModeOk

`func (o *InternalTransferResponse) GetCaptureModeOk() (*string, bool)`

GetCaptureModeOk returns a tuple with the CaptureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMode

`func (o *InternalTransferResponse) SetCaptureMode(v string)`

SetCaptureMode sets CaptureMode field to given value.

### HasCaptureMode

`func (o *InternalTransferResponse) HasCaptureMode() bool`

HasCaptureMode returns a boolean if a field has been set.

### GetCurrency

`func (o *InternalTransferResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalTransferResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalTransferResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExpiresAt

`func (o *InternalTransferResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InternalTransferResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InternalTransferResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InternalTransferResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFinalCustomerId

`func (o *InternalTransferResponse) GetFinalCustomerId() string`

GetFinalCustomerId returns the FinalCustomerId field if non-nil, zero value otherwise.

### GetFinalCustomerIdOk

`func (o *InternalTransferResponse) GetFinalCustomerIdOk() (*string, bool)`

GetFinalCustomerIdOk returns a tuple with the FinalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCustomerId

`func (o *InternalTransferResponse) SetFinalCustomerId(v string)`

SetFinalCustomerId sets FinalCustomerId field to given value.

### HasFinalCustomerId

`func (o *InternalTransferResponse) HasFinalCustomerId() bool`

HasFinalCustomerId returns a boolean if a field has been set.

### GetMemo

`func (o *InternalTransferResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *InternalTransferResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *InternalTransferResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *InternalTransferResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMetadata

`func (o *InternalTransferResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternalTransferResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternalTransferResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternalTransferResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginatingAccountAlias

`func (o *InternalTransferResponse) GetOriginatingAccountAlias() string`

GetOriginatingAccountAlias returns the OriginatingAccountAlias field if non-nil, zero value otherwise.

### GetOriginatingAccountAliasOk

`func (o *InternalTransferResponse) GetOriginatingAccountAliasOk() (*string, bool)`

GetOriginatingAccountAliasOk returns a tuple with the OriginatingAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountAlias

`func (o *InternalTransferResponse) SetOriginatingAccountAlias(v string)`

SetOriginatingAccountAlias sets OriginatingAccountAlias field to given value.

### HasOriginatingAccountAlias

`func (o *InternalTransferResponse) HasOriginatingAccountAlias() bool`

HasOriginatingAccountAlias returns a boolean if a field has been set.

### GetOriginatingAccountCustomerId

`func (o *InternalTransferResponse) GetOriginatingAccountCustomerId() string`

GetOriginatingAccountCustomerId returns the OriginatingAccountCustomerId field if non-nil, zero value otherwise.

### GetOriginatingAccountCustomerIdOk

`func (o *InternalTransferResponse) GetOriginatingAccountCustomerIdOk() (*string, bool)`

GetOriginatingAccountCustomerIdOk returns a tuple with the OriginatingAccountCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountCustomerId

`func (o *InternalTransferResponse) SetOriginatingAccountCustomerId(v string)`

SetOriginatingAccountCustomerId sets OriginatingAccountCustomerId field to given value.

### HasOriginatingAccountCustomerId

`func (o *InternalTransferResponse) HasOriginatingAccountCustomerId() bool`

HasOriginatingAccountCustomerId returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *InternalTransferResponse) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *InternalTransferResponse) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *InternalTransferResponse) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.

### HasOriginatingAccountId

`func (o *InternalTransferResponse) HasOriginatingAccountId() bool`

HasOriginatingAccountId returns a boolean if a field has been set.

### GetReceivingAccountAlias

`func (o *InternalTransferResponse) GetReceivingAccountAlias() string`

GetReceivingAccountAlias returns the ReceivingAccountAlias field if non-nil, zero value otherwise.

### GetReceivingAccountAliasOk

`func (o *InternalTransferResponse) GetReceivingAccountAliasOk() (*string, bool)`

GetReceivingAccountAliasOk returns a tuple with the ReceivingAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountAlias

`func (o *InternalTransferResponse) SetReceivingAccountAlias(v string)`

SetReceivingAccountAlias sets ReceivingAccountAlias field to given value.

### HasReceivingAccountAlias

`func (o *InternalTransferResponse) HasReceivingAccountAlias() bool`

HasReceivingAccountAlias returns a boolean if a field has been set.

### GetReceivingAccountCustomerId

`func (o *InternalTransferResponse) GetReceivingAccountCustomerId() string`

GetReceivingAccountCustomerId returns the ReceivingAccountCustomerId field if non-nil, zero value otherwise.

### GetReceivingAccountCustomerIdOk

`func (o *InternalTransferResponse) GetReceivingAccountCustomerIdOk() (*string, bool)`

GetReceivingAccountCustomerIdOk returns a tuple with the ReceivingAccountCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountCustomerId

`func (o *InternalTransferResponse) SetReceivingAccountCustomerId(v string)`

SetReceivingAccountCustomerId sets ReceivingAccountCustomerId field to given value.

### HasReceivingAccountCustomerId

`func (o *InternalTransferResponse) HasReceivingAccountCustomerId() bool`

HasReceivingAccountCustomerId returns a boolean if a field has been set.

### GetReceivingAccountId

`func (o *InternalTransferResponse) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *InternalTransferResponse) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *InternalTransferResponse) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.

### HasReceivingAccountId

`func (o *InternalTransferResponse) HasReceivingAccountId() bool`

HasReceivingAccountId returns a boolean if a field has been set.

### GetTenant

`func (o *InternalTransferResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *InternalTransferResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *InternalTransferResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *InternalTransferResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *InternalTransferResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalTransferResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalTransferResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InternalTransferResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalTransferResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalTransferResponse) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *InternalTransferResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalTransferResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalTransferResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


