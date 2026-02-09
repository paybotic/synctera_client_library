# TransferRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount of the transfer in cents (USD) | 
**Merchant** | Pointer to [**Merchant**](Merchant.md) |  | [optional] 
**OriginatingAccountId** | **string** | The ID of the account to which the transfer will be initiated/received | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 

## Methods

### NewTransferRequestBase

`func NewTransferRequestBase(amount int32, originatingAccountId string, ) *TransferRequestBase`

NewTransferRequestBase instantiates a new TransferRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestBaseWithDefaults

`func NewTransferRequestBaseWithDefaults() *TransferRequestBase`

NewTransferRequestBaseWithDefaults instantiates a new TransferRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferRequestBase) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRequestBase) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRequestBase) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetMerchant

`func (o *TransferRequestBase) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransferRequestBase) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransferRequestBase) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *TransferRequestBase) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *TransferRequestBase) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *TransferRequestBase) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *TransferRequestBase) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetTenant

`func (o *TransferRequestBase) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransferRequestBase) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransferRequestBase) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TransferRequestBase) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


