# TransferRequestApplePay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount of the transfer in cents (USD) | 
**Merchant** | Pointer to [**Merchant**](Merchant.md) |  | [optional] 
**OriginatingAccountId** | **string** | The ID of the account to which the transfer will be initiated/received | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 
**ApplePayPayment** | [**ApplePayPayment**](ApplePayPayment.md) |  | 
**CustomerId** | **string** | The customer_id of the cardholder | 

## Methods

### NewTransferRequestApplePay

`func NewTransferRequestApplePay(amount int32, originatingAccountId string, applePayPayment ApplePayPayment, customerId string, ) *TransferRequestApplePay`

NewTransferRequestApplePay instantiates a new TransferRequestApplePay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestApplePayWithDefaults

`func NewTransferRequestApplePayWithDefaults() *TransferRequestApplePay`

NewTransferRequestApplePayWithDefaults instantiates a new TransferRequestApplePay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferRequestApplePay) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRequestApplePay) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRequestApplePay) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetMerchant

`func (o *TransferRequestApplePay) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransferRequestApplePay) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransferRequestApplePay) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *TransferRequestApplePay) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *TransferRequestApplePay) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *TransferRequestApplePay) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *TransferRequestApplePay) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetTenant

`func (o *TransferRequestApplePay) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransferRequestApplePay) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransferRequestApplePay) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TransferRequestApplePay) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetApplePayPayment

`func (o *TransferRequestApplePay) GetApplePayPayment() ApplePayPayment`

GetApplePayPayment returns the ApplePayPayment field if non-nil, zero value otherwise.

### GetApplePayPaymentOk

`func (o *TransferRequestApplePay) GetApplePayPaymentOk() (*ApplePayPayment, bool)`

GetApplePayPaymentOk returns a tuple with the ApplePayPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplePayPayment

`func (o *TransferRequestApplePay) SetApplePayPayment(v ApplePayPayment)`

SetApplePayPayment sets ApplePayPayment field to given value.


### GetCustomerId

`func (o *TransferRequestApplePay) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TransferRequestApplePay) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TransferRequestApplePay) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


