# ApplePayPaymentHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationData** | Pointer to **string** | SHA–256 hash, hex encoded as a string | [optional] 
**EphemeralPublicKey** | **string** | X.509 encoded public key bytes | 
**PublicKeyHash** | **string** | Hash of the X.509 encoded public key bytes | 
**TransactionId** | **string** | Device generated transaction identifier | 

## Methods

### NewApplePayPaymentHeader

`func NewApplePayPaymentHeader(ephemeralPublicKey string, publicKeyHash string, transactionId string, ) *ApplePayPaymentHeader`

NewApplePayPaymentHeader instantiates a new ApplePayPaymentHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePayPaymentHeaderWithDefaults

`func NewApplePayPaymentHeaderWithDefaults() *ApplePayPaymentHeader`

NewApplePayPaymentHeaderWithDefaults instantiates a new ApplePayPaymentHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationData

`func (o *ApplePayPaymentHeader) GetApplicationData() string`

GetApplicationData returns the ApplicationData field if non-nil, zero value otherwise.

### GetApplicationDataOk

`func (o *ApplePayPaymentHeader) GetApplicationDataOk() (*string, bool)`

GetApplicationDataOk returns a tuple with the ApplicationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationData

`func (o *ApplePayPaymentHeader) SetApplicationData(v string)`

SetApplicationData sets ApplicationData field to given value.

### HasApplicationData

`func (o *ApplePayPaymentHeader) HasApplicationData() bool`

HasApplicationData returns a boolean if a field has been set.

### GetEphemeralPublicKey

`func (o *ApplePayPaymentHeader) GetEphemeralPublicKey() string`

GetEphemeralPublicKey returns the EphemeralPublicKey field if non-nil, zero value otherwise.

### GetEphemeralPublicKeyOk

`func (o *ApplePayPaymentHeader) GetEphemeralPublicKeyOk() (*string, bool)`

GetEphemeralPublicKeyOk returns a tuple with the EphemeralPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralPublicKey

`func (o *ApplePayPaymentHeader) SetEphemeralPublicKey(v string)`

SetEphemeralPublicKey sets EphemeralPublicKey field to given value.


### GetPublicKeyHash

`func (o *ApplePayPaymentHeader) GetPublicKeyHash() string`

GetPublicKeyHash returns the PublicKeyHash field if non-nil, zero value otherwise.

### GetPublicKeyHashOk

`func (o *ApplePayPaymentHeader) GetPublicKeyHashOk() (*string, bool)`

GetPublicKeyHashOk returns a tuple with the PublicKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyHash

`func (o *ApplePayPaymentHeader) SetPublicKeyHash(v string)`

SetPublicKeyHash sets PublicKeyHash field to given value.


### GetTransactionId

`func (o *ApplePayPaymentHeader) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ApplePayPaymentHeader) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ApplePayPaymentHeader) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


