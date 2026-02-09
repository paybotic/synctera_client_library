# TransferRequestGooglePay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount of the transfer in cents (USD) | 
**Merchant** | Pointer to [**Merchant**](Merchant.md) |  | [optional] 
**OriginatingAccountId** | **string** | The ID of the account to which the transfer will be initiated/received | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 
**CustomerId** | **string** | The customer_id of the cardholder | 
**GooglePayPaymentData** | [**GooglePayPaymentData**](GooglePayPaymentData.md) |  | 

## Methods

### NewTransferRequestGooglePay

`func NewTransferRequestGooglePay(amount int32, originatingAccountId string, customerId string, googlePayPaymentData GooglePayPaymentData, ) *TransferRequestGooglePay`

NewTransferRequestGooglePay instantiates a new TransferRequestGooglePay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestGooglePayWithDefaults

`func NewTransferRequestGooglePayWithDefaults() *TransferRequestGooglePay`

NewTransferRequestGooglePayWithDefaults instantiates a new TransferRequestGooglePay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferRequestGooglePay) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRequestGooglePay) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRequestGooglePay) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetMerchant

`func (o *TransferRequestGooglePay) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransferRequestGooglePay) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransferRequestGooglePay) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *TransferRequestGooglePay) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *TransferRequestGooglePay) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *TransferRequestGooglePay) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *TransferRequestGooglePay) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetTenant

`func (o *TransferRequestGooglePay) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransferRequestGooglePay) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransferRequestGooglePay) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TransferRequestGooglePay) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetCustomerId

`func (o *TransferRequestGooglePay) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TransferRequestGooglePay) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TransferRequestGooglePay) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetGooglePayPaymentData

`func (o *TransferRequestGooglePay) GetGooglePayPaymentData() GooglePayPaymentData`

GetGooglePayPaymentData returns the GooglePayPaymentData field if non-nil, zero value otherwise.

### GetGooglePayPaymentDataOk

`func (o *TransferRequestGooglePay) GetGooglePayPaymentDataOk() (*GooglePayPaymentData, bool)`

GetGooglePayPaymentDataOk returns a tuple with the GooglePayPaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePayPaymentData

`func (o *TransferRequestGooglePay) SetGooglePayPaymentData(v GooglePayPaymentData)`

SetGooglePayPaymentData sets GooglePayPaymentData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


