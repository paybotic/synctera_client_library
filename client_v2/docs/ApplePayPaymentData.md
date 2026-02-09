# ApplePayPaymentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Encrypted payment data, Base64 encoded | 
**Header** | [**ApplePayPaymentHeader**](ApplePayPaymentHeader.md) |  | 
**Signature** | **string** | Signature of the payment and header data | 
**Version** | **string** | * RSA_v1 encryption is not currently supported. Version information about the payment token  | 

## Methods

### NewApplePayPaymentData

`func NewApplePayPaymentData(data string, header ApplePayPaymentHeader, signature string, version string, ) *ApplePayPaymentData`

NewApplePayPaymentData instantiates a new ApplePayPaymentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePayPaymentDataWithDefaults

`func NewApplePayPaymentDataWithDefaults() *ApplePayPaymentData`

NewApplePayPaymentDataWithDefaults instantiates a new ApplePayPaymentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApplePayPaymentData) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplePayPaymentData) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplePayPaymentData) SetData(v string)`

SetData sets Data field to given value.


### GetHeader

`func (o *ApplePayPaymentData) GetHeader() ApplePayPaymentHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ApplePayPaymentData) GetHeaderOk() (*ApplePayPaymentHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ApplePayPaymentData) SetHeader(v ApplePayPaymentHeader)`

SetHeader sets Header field to given value.


### GetSignature

`func (o *ApplePayPaymentData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ApplePayPaymentData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ApplePayPaymentData) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetVersion

`func (o *ApplePayPaymentData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplePayPaymentData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplePayPaymentData) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


