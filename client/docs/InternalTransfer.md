# InternalTransfer

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
**ReferenceId** | Pointer to **string** | Network reference id | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**Type** | **string** | The desired transaction type to use for this transfer | 

## Methods

### NewInternalTransfer

`func NewInternalTransfer(amount int64, currency string, type_ string, ) *InternalTransfer`

NewInternalTransfer instantiates a new InternalTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalTransferWithDefaults

`func NewInternalTransferWithDefaults() *InternalTransfer`

NewInternalTransferWithDefaults instantiates a new InternalTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalTransfer) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalTransfer) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalTransfer) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCaptureMode

`func (o *InternalTransfer) GetCaptureMode() string`

GetCaptureMode returns the CaptureMode field if non-nil, zero value otherwise.

### GetCaptureModeOk

`func (o *InternalTransfer) GetCaptureModeOk() (*string, bool)`

GetCaptureModeOk returns a tuple with the CaptureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMode

`func (o *InternalTransfer) SetCaptureMode(v string)`

SetCaptureMode sets CaptureMode field to given value.

### HasCaptureMode

`func (o *InternalTransfer) HasCaptureMode() bool`

HasCaptureMode returns a boolean if a field has been set.

### GetCurrency

`func (o *InternalTransfer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalTransfer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalTransfer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExpiresAt

`func (o *InternalTransfer) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InternalTransfer) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InternalTransfer) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InternalTransfer) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFinalCustomerId

`func (o *InternalTransfer) GetFinalCustomerId() string`

GetFinalCustomerId returns the FinalCustomerId field if non-nil, zero value otherwise.

### GetFinalCustomerIdOk

`func (o *InternalTransfer) GetFinalCustomerIdOk() (*string, bool)`

GetFinalCustomerIdOk returns a tuple with the FinalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCustomerId

`func (o *InternalTransfer) SetFinalCustomerId(v string)`

SetFinalCustomerId sets FinalCustomerId field to given value.

### HasFinalCustomerId

`func (o *InternalTransfer) HasFinalCustomerId() bool`

HasFinalCustomerId returns a boolean if a field has been set.

### GetMemo

`func (o *InternalTransfer) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *InternalTransfer) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *InternalTransfer) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *InternalTransfer) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMetadata

`func (o *InternalTransfer) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternalTransfer) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternalTransfer) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternalTransfer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginatingAccountAlias

`func (o *InternalTransfer) GetOriginatingAccountAlias() string`

GetOriginatingAccountAlias returns the OriginatingAccountAlias field if non-nil, zero value otherwise.

### GetOriginatingAccountAliasOk

`func (o *InternalTransfer) GetOriginatingAccountAliasOk() (*string, bool)`

GetOriginatingAccountAliasOk returns a tuple with the OriginatingAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountAlias

`func (o *InternalTransfer) SetOriginatingAccountAlias(v string)`

SetOriginatingAccountAlias sets OriginatingAccountAlias field to given value.

### HasOriginatingAccountAlias

`func (o *InternalTransfer) HasOriginatingAccountAlias() bool`

HasOriginatingAccountAlias returns a boolean if a field has been set.

### GetOriginatingAccountCustomerId

`func (o *InternalTransfer) GetOriginatingAccountCustomerId() string`

GetOriginatingAccountCustomerId returns the OriginatingAccountCustomerId field if non-nil, zero value otherwise.

### GetOriginatingAccountCustomerIdOk

`func (o *InternalTransfer) GetOriginatingAccountCustomerIdOk() (*string, bool)`

GetOriginatingAccountCustomerIdOk returns a tuple with the OriginatingAccountCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountCustomerId

`func (o *InternalTransfer) SetOriginatingAccountCustomerId(v string)`

SetOriginatingAccountCustomerId sets OriginatingAccountCustomerId field to given value.

### HasOriginatingAccountCustomerId

`func (o *InternalTransfer) HasOriginatingAccountCustomerId() bool`

HasOriginatingAccountCustomerId returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *InternalTransfer) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *InternalTransfer) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *InternalTransfer) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.

### HasOriginatingAccountId

`func (o *InternalTransfer) HasOriginatingAccountId() bool`

HasOriginatingAccountId returns a boolean if a field has been set.

### GetReceivingAccountAlias

`func (o *InternalTransfer) GetReceivingAccountAlias() string`

GetReceivingAccountAlias returns the ReceivingAccountAlias field if non-nil, zero value otherwise.

### GetReceivingAccountAliasOk

`func (o *InternalTransfer) GetReceivingAccountAliasOk() (*string, bool)`

GetReceivingAccountAliasOk returns a tuple with the ReceivingAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountAlias

`func (o *InternalTransfer) SetReceivingAccountAlias(v string)`

SetReceivingAccountAlias sets ReceivingAccountAlias field to given value.

### HasReceivingAccountAlias

`func (o *InternalTransfer) HasReceivingAccountAlias() bool`

HasReceivingAccountAlias returns a boolean if a field has been set.

### GetReceivingAccountCustomerId

`func (o *InternalTransfer) GetReceivingAccountCustomerId() string`

GetReceivingAccountCustomerId returns the ReceivingAccountCustomerId field if non-nil, zero value otherwise.

### GetReceivingAccountCustomerIdOk

`func (o *InternalTransfer) GetReceivingAccountCustomerIdOk() (*string, bool)`

GetReceivingAccountCustomerIdOk returns a tuple with the ReceivingAccountCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountCustomerId

`func (o *InternalTransfer) SetReceivingAccountCustomerId(v string)`

SetReceivingAccountCustomerId sets ReceivingAccountCustomerId field to given value.

### HasReceivingAccountCustomerId

`func (o *InternalTransfer) HasReceivingAccountCustomerId() bool`

HasReceivingAccountCustomerId returns a boolean if a field has been set.

### GetReceivingAccountId

`func (o *InternalTransfer) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *InternalTransfer) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *InternalTransfer) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.

### HasReceivingAccountId

`func (o *InternalTransfer) HasReceivingAccountId() bool`

HasReceivingAccountId returns a boolean if a field has been set.

### GetReferenceId

`func (o *InternalTransfer) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InternalTransfer) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InternalTransfer) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InternalTransfer) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetTenant

`func (o *InternalTransfer) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *InternalTransfer) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *InternalTransfer) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *InternalTransfer) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *InternalTransfer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalTransfer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalTransfer) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


