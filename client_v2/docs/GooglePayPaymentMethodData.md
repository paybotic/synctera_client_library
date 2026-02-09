# GooglePayPaymentMethodData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | User-facing message to describe the payment method that funds this transaction | 
**Info** | [**GooglePayInfo**](GooglePayInfo.md) |  | 
**TokenizationData** | [**GooglePayTokenizationData**](GooglePayTokenizationData.md) |  | 
**Type** | **string** | Payment method type selected in the Google Pay payment sheet. &#x60;CARD&#x60; is the only supported value. | 

## Methods

### NewGooglePayPaymentMethodData

`func NewGooglePayPaymentMethodData(description string, info GooglePayInfo, tokenizationData GooglePayTokenizationData, type_ string, ) *GooglePayPaymentMethodData`

NewGooglePayPaymentMethodData instantiates a new GooglePayPaymentMethodData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayPaymentMethodDataWithDefaults

`func NewGooglePayPaymentMethodDataWithDefaults() *GooglePayPaymentMethodData`

NewGooglePayPaymentMethodDataWithDefaults instantiates a new GooglePayPaymentMethodData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GooglePayPaymentMethodData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GooglePayPaymentMethodData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GooglePayPaymentMethodData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInfo

`func (o *GooglePayPaymentMethodData) GetInfo() GooglePayInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GooglePayPaymentMethodData) GetInfoOk() (*GooglePayInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GooglePayPaymentMethodData) SetInfo(v GooglePayInfo)`

SetInfo sets Info field to given value.


### GetTokenizationData

`func (o *GooglePayPaymentMethodData) GetTokenizationData() GooglePayTokenizationData`

GetTokenizationData returns the TokenizationData field if non-nil, zero value otherwise.

### GetTokenizationDataOk

`func (o *GooglePayPaymentMethodData) GetTokenizationDataOk() (*GooglePayTokenizationData, bool)`

GetTokenizationDataOk returns a tuple with the TokenizationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizationData

`func (o *GooglePayPaymentMethodData) SetTokenizationData(v GooglePayTokenizationData)`

SetTokenizationData sets TokenizationData field to given value.


### GetType

`func (o *GooglePayPaymentMethodData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GooglePayPaymentMethodData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GooglePayPaymentMethodData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


