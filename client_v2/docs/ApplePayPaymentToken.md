# ApplePayPaymentToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentData** | [**ApplePayPaymentData**](ApplePayPaymentData.md) |  | 
**PaymentMethod** | [**ApplePayPaymentMethod**](ApplePayPaymentMethod.md) |  | 
**TransactionIdentifier** | **string** | Unique identifier for this payment | 

## Methods

### NewApplePayPaymentToken

`func NewApplePayPaymentToken(paymentData ApplePayPaymentData, paymentMethod ApplePayPaymentMethod, transactionIdentifier string, ) *ApplePayPaymentToken`

NewApplePayPaymentToken instantiates a new ApplePayPaymentToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePayPaymentTokenWithDefaults

`func NewApplePayPaymentTokenWithDefaults() *ApplePayPaymentToken`

NewApplePayPaymentTokenWithDefaults instantiates a new ApplePayPaymentToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentData

`func (o *ApplePayPaymentToken) GetPaymentData() ApplePayPaymentData`

GetPaymentData returns the PaymentData field if non-nil, zero value otherwise.

### GetPaymentDataOk

`func (o *ApplePayPaymentToken) GetPaymentDataOk() (*ApplePayPaymentData, bool)`

GetPaymentDataOk returns a tuple with the PaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentData

`func (o *ApplePayPaymentToken) SetPaymentData(v ApplePayPaymentData)`

SetPaymentData sets PaymentData field to given value.


### GetPaymentMethod

`func (o *ApplePayPaymentToken) GetPaymentMethod() ApplePayPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ApplePayPaymentToken) GetPaymentMethodOk() (*ApplePayPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ApplePayPaymentToken) SetPaymentMethod(v ApplePayPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetTransactionIdentifier

`func (o *ApplePayPaymentToken) GetTransactionIdentifier() string`

GetTransactionIdentifier returns the TransactionIdentifier field if non-nil, zero value otherwise.

### GetTransactionIdentifierOk

`func (o *ApplePayPaymentToken) GetTransactionIdentifierOk() (*string, bool)`

GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentifier

`func (o *ApplePayPaymentToken) SetTransactionIdentifier(v string)`

SetTransactionIdentifier sets TransactionIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


