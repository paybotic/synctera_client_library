# ApplePayPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to [**ApplePayContact**](ApplePayContact.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Name that describes the card, suitable for display | [optional] 
**Network** | Pointer to **string** | Name of the payment network backing the card, suitable for display | [optional] 
**SecureElementPass** | Pointer to [**ApplePaySecureElementPass**](ApplePaySecureElementPass.md) |  | [optional] 
**Type** | Pointer to **string** | Card payment type | [optional] 

## Methods

### NewApplePayPaymentMethod

`func NewApplePayPaymentMethod() *ApplePayPaymentMethod`

NewApplePayPaymentMethod instantiates a new ApplePayPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePayPaymentMethodWithDefaults

`func NewApplePayPaymentMethodWithDefaults() *ApplePayPaymentMethod`

NewApplePayPaymentMethodWithDefaults instantiates a new ApplePayPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *ApplePayPaymentMethod) GetBillingAddress() ApplePayContact`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ApplePayPaymentMethod) GetBillingAddressOk() (*ApplePayContact, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ApplePayPaymentMethod) SetBillingAddress(v ApplePayContact)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *ApplePayPaymentMethod) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApplePayPaymentMethod) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApplePayPaymentMethod) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApplePayPaymentMethod) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApplePayPaymentMethod) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetNetwork

`func (o *ApplePayPaymentMethod) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ApplePayPaymentMethod) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ApplePayPaymentMethod) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ApplePayPaymentMethod) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSecureElementPass

`func (o *ApplePayPaymentMethod) GetSecureElementPass() ApplePaySecureElementPass`

GetSecureElementPass returns the SecureElementPass field if non-nil, zero value otherwise.

### GetSecureElementPassOk

`func (o *ApplePayPaymentMethod) GetSecureElementPassOk() (*ApplePaySecureElementPass, bool)`

GetSecureElementPassOk returns a tuple with the SecureElementPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureElementPass

`func (o *ApplePayPaymentMethod) SetSecureElementPass(v ApplePaySecureElementPass)`

SetSecureElementPass sets SecureElementPass field to given value.

### HasSecureElementPass

`func (o *ApplePayPaymentMethod) HasSecureElementPass() bool`

HasSecureElementPass returns a boolean if a field has been set.

### GetType

`func (o *ApplePayPaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplePayPaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplePayPaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApplePayPaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


