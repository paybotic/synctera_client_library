# TransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount of the transfer in cents (USD) | 
**Merchant** | Pointer to [**Merchant**](Merchant.md) |  | [optional] 
**OriginatingAccountId** | **string** | The ID of the account to which the transfer will be initiated/received | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 
**ExternalCardId** | **string** | The ID of the External Card associated with the operation | 
**ThreeDsId** | Pointer to **string** | Unique identifier of an External Card Transfer 3-D Secure Authorization - conditionally required according to your program&#39;s 3DS policy | [optional] 
**Type** | [**CreateTransferType**](CreateTransferType.md) |  | 
**OriginatingCustomerId** | Pointer to **string** | For person-to-person PUSH transactions this is the customer_id of the sender who must have privileges to access funds in the originating account in order to send funds to the recipient cardholder | [optional] 

## Methods

### NewTransferRequest

`func NewTransferRequest(amount int32, originatingAccountId string, externalCardId string, type_ CreateTransferType, ) *TransferRequest`

NewTransferRequest instantiates a new TransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestWithDefaults

`func NewTransferRequestWithDefaults() *TransferRequest`

NewTransferRequestWithDefaults instantiates a new TransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetMerchant

`func (o *TransferRequest) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransferRequest) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransferRequest) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *TransferRequest) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *TransferRequest) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *TransferRequest) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *TransferRequest) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetTenant

`func (o *TransferRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransferRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransferRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TransferRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetExternalCardId

`func (o *TransferRequest) GetExternalCardId() string`

GetExternalCardId returns the ExternalCardId field if non-nil, zero value otherwise.

### GetExternalCardIdOk

`func (o *TransferRequest) GetExternalCardIdOk() (*string, bool)`

GetExternalCardIdOk returns a tuple with the ExternalCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCardId

`func (o *TransferRequest) SetExternalCardId(v string)`

SetExternalCardId sets ExternalCardId field to given value.


### GetThreeDsId

`func (o *TransferRequest) GetThreeDsId() string`

GetThreeDsId returns the ThreeDsId field if non-nil, zero value otherwise.

### GetThreeDsIdOk

`func (o *TransferRequest) GetThreeDsIdOk() (*string, bool)`

GetThreeDsIdOk returns a tuple with the ThreeDsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsId

`func (o *TransferRequest) SetThreeDsId(v string)`

SetThreeDsId sets ThreeDsId field to given value.

### HasThreeDsId

`func (o *TransferRequest) HasThreeDsId() bool`

HasThreeDsId returns a boolean if a field has been set.

### GetType

`func (o *TransferRequest) GetType() CreateTransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferRequest) GetTypeOk() (*CreateTransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferRequest) SetType(v CreateTransferType)`

SetType sets Type field to given value.


### GetOriginatingCustomerId

`func (o *TransferRequest) GetOriginatingCustomerId() string`

GetOriginatingCustomerId returns the OriginatingCustomerId field if non-nil, zero value otherwise.

### GetOriginatingCustomerIdOk

`func (o *TransferRequest) GetOriginatingCustomerIdOk() (*string, bool)`

GetOriginatingCustomerIdOk returns a tuple with the OriginatingCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingCustomerId

`func (o *TransferRequest) SetOriginatingCustomerId(v string)`

SetOriginatingCustomerId sets OriginatingCustomerId field to given value.

### HasOriginatingCustomerId

`func (o *TransferRequest) HasOriginatingCustomerId() bool`

HasOriginatingCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


