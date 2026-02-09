# ApplePayPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingContact** | Pointer to [**ApplePayContact**](ApplePayContact.md) |  | [optional] 
**ShippingContact** | Pointer to [**ApplePayContact**](ApplePayContact.md) |  | [optional] 
**Token** | [**ApplePayPaymentToken**](ApplePayPaymentToken.md) |  | 

## Methods

### NewApplePayPayment

`func NewApplePayPayment(token ApplePayPaymentToken, ) *ApplePayPayment`

NewApplePayPayment instantiates a new ApplePayPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePayPaymentWithDefaults

`func NewApplePayPaymentWithDefaults() *ApplePayPayment`

NewApplePayPaymentWithDefaults instantiates a new ApplePayPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingContact

`func (o *ApplePayPayment) GetBillingContact() ApplePayContact`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *ApplePayPayment) GetBillingContactOk() (*ApplePayContact, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *ApplePayPayment) SetBillingContact(v ApplePayContact)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *ApplePayPayment) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### GetShippingContact

`func (o *ApplePayPayment) GetShippingContact() ApplePayContact`

GetShippingContact returns the ShippingContact field if non-nil, zero value otherwise.

### GetShippingContactOk

`func (o *ApplePayPayment) GetShippingContactOk() (*ApplePayContact, bool)`

GetShippingContactOk returns a tuple with the ShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingContact

`func (o *ApplePayPayment) SetShippingContact(v ApplePayContact)`

SetShippingContact sets ShippingContact field to given value.

### HasShippingContact

`func (o *ApplePayPayment) HasShippingContact() bool`

HasShippingContact returns a boolean if a field has been set.

### GetToken

`func (o *ApplePayPayment) GetToken() ApplePayPaymentToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApplePayPayment) GetTokenOk() (*ApplePayPaymentToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApplePayPayment) SetToken(v ApplePayPaymentToken)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


