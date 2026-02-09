# GooglePayPaymentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**GooglePayAddress**](GooglePayAddress.md) |  | [optional] 
**ApiVersion** | **int32** | Major API version | 
**ApiVersionMinor** | **int32** | Minor API version | 
**Email** | Pointer to **string** | Email address of cardholder, present if requested in payment request | [optional] 
**PaymentMethodData** | [**GooglePayPaymentMethodData**](GooglePayPaymentMethodData.md) |  | 

## Methods

### NewGooglePayPaymentData

`func NewGooglePayPaymentData(apiVersion int32, apiVersionMinor int32, paymentMethodData GooglePayPaymentMethodData, ) *GooglePayPaymentData`

NewGooglePayPaymentData instantiates a new GooglePayPaymentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayPaymentDataWithDefaults

`func NewGooglePayPaymentDataWithDefaults() *GooglePayPaymentData`

NewGooglePayPaymentDataWithDefaults instantiates a new GooglePayPaymentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GooglePayPaymentData) GetAddress() GooglePayAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GooglePayPaymentData) GetAddressOk() (*GooglePayAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GooglePayPaymentData) SetAddress(v GooglePayAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GooglePayPaymentData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetApiVersion

`func (o *GooglePayPaymentData) GetApiVersion() int32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GooglePayPaymentData) GetApiVersionOk() (*int32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GooglePayPaymentData) SetApiVersion(v int32)`

SetApiVersion sets ApiVersion field to given value.


### GetApiVersionMinor

`func (o *GooglePayPaymentData) GetApiVersionMinor() int32`

GetApiVersionMinor returns the ApiVersionMinor field if non-nil, zero value otherwise.

### GetApiVersionMinorOk

`func (o *GooglePayPaymentData) GetApiVersionMinorOk() (*int32, bool)`

GetApiVersionMinorOk returns a tuple with the ApiVersionMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersionMinor

`func (o *GooglePayPaymentData) SetApiVersionMinor(v int32)`

SetApiVersionMinor sets ApiVersionMinor field to given value.


### GetEmail

`func (o *GooglePayPaymentData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GooglePayPaymentData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GooglePayPaymentData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GooglePayPaymentData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPaymentMethodData

`func (o *GooglePayPaymentData) GetPaymentMethodData() GooglePayPaymentMethodData`

GetPaymentMethodData returns the PaymentMethodData field if non-nil, zero value otherwise.

### GetPaymentMethodDataOk

`func (o *GooglePayPaymentData) GetPaymentMethodDataOk() (*GooglePayPaymentMethodData, bool)`

GetPaymentMethodDataOk returns a tuple with the PaymentMethodData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodData

`func (o *GooglePayPaymentData) SetPaymentMethodData(v GooglePayPaymentMethodData)`

SetPaymentMethodData sets PaymentMethodData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


