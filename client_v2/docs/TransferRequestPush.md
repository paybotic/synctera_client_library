# TransferRequestPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount of the transfer in cents (USD) | 
**Merchant** | Pointer to [**Merchant**](Merchant.md) |  | [optional] 
**OriginatingAccountId** | **string** | The ID of the account to which the transfer will be initiated/received | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 
**ExternalCardId** | **string** | The ID of the External Card associated with the operation | 
**OriginatingCustomerId** | Pointer to **string** | For person-to-person PUSH transactions this is the customer_id of the sender who must have privileges to access funds in the originating account in order to send funds to the recipient cardholder | [optional] 
**Type** | [**CreateTransferType**](CreateTransferType.md) |  | 

## Methods

### NewTransferRequestPush

`func NewTransferRequestPush(amount int32, originatingAccountId string, externalCardId string, type_ CreateTransferType, ) *TransferRequestPush`

NewTransferRequestPush instantiates a new TransferRequestPush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestPushWithDefaults

`func NewTransferRequestPushWithDefaults() *TransferRequestPush`

NewTransferRequestPushWithDefaults instantiates a new TransferRequestPush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferRequestPush) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRequestPush) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRequestPush) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetMerchant

`func (o *TransferRequestPush) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransferRequestPush) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransferRequestPush) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *TransferRequestPush) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *TransferRequestPush) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *TransferRequestPush) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *TransferRequestPush) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetTenant

`func (o *TransferRequestPush) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransferRequestPush) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransferRequestPush) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TransferRequestPush) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetExternalCardId

`func (o *TransferRequestPush) GetExternalCardId() string`

GetExternalCardId returns the ExternalCardId field if non-nil, zero value otherwise.

### GetExternalCardIdOk

`func (o *TransferRequestPush) GetExternalCardIdOk() (*string, bool)`

GetExternalCardIdOk returns a tuple with the ExternalCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCardId

`func (o *TransferRequestPush) SetExternalCardId(v string)`

SetExternalCardId sets ExternalCardId field to given value.


### GetOriginatingCustomerId

`func (o *TransferRequestPush) GetOriginatingCustomerId() string`

GetOriginatingCustomerId returns the OriginatingCustomerId field if non-nil, zero value otherwise.

### GetOriginatingCustomerIdOk

`func (o *TransferRequestPush) GetOriginatingCustomerIdOk() (*string, bool)`

GetOriginatingCustomerIdOk returns a tuple with the OriginatingCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingCustomerId

`func (o *TransferRequestPush) SetOriginatingCustomerId(v string)`

SetOriginatingCustomerId sets OriginatingCustomerId field to given value.

### HasOriginatingCustomerId

`func (o *TransferRequestPush) HasOriginatingCustomerId() bool`

HasOriginatingCustomerId returns a boolean if a field has been set.

### GetType

`func (o *TransferRequestPush) GetType() CreateTransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferRequestPush) GetTypeOk() (*CreateTransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferRequestPush) SetType(v CreateTransferType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


