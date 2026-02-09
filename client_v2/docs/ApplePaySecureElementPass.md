# ApplePaySecureElementPass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceAccountIdentifier** | **string** | The unique identifier for the device-specific account number. | 
**DeviceAccountNumberSuffix** | **string** | A display-ready version of the device-specific account number. | 
**DevicePassIdentifier** | Pointer to **string** | An opaque value for the pass. | [optional] 
**PairedTerminalIdentifier** | Pointer to **string** | The unique identifier of the paired terminal. | [optional] 
**PassActivationState** | **string** | The activation state of the pass. | 
**PrimaryAccountIdentifier** | **string** | An opaque value that identifies the primary account number that funds the pass’s transactions. | 
**PrimaryAccountNumberSuffix** | **string** | A display-ready version of the primary account number. | 

## Methods

### NewApplePaySecureElementPass

`func NewApplePaySecureElementPass(deviceAccountIdentifier string, deviceAccountNumberSuffix string, passActivationState string, primaryAccountIdentifier string, primaryAccountNumberSuffix string, ) *ApplePaySecureElementPass`

NewApplePaySecureElementPass instantiates a new ApplePaySecureElementPass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePaySecureElementPassWithDefaults

`func NewApplePaySecureElementPassWithDefaults() *ApplePaySecureElementPass`

NewApplePaySecureElementPassWithDefaults instantiates a new ApplePaySecureElementPass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceAccountIdentifier

`func (o *ApplePaySecureElementPass) GetDeviceAccountIdentifier() string`

GetDeviceAccountIdentifier returns the DeviceAccountIdentifier field if non-nil, zero value otherwise.

### GetDeviceAccountIdentifierOk

`func (o *ApplePaySecureElementPass) GetDeviceAccountIdentifierOk() (*string, bool)`

GetDeviceAccountIdentifierOk returns a tuple with the DeviceAccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccountIdentifier

`func (o *ApplePaySecureElementPass) SetDeviceAccountIdentifier(v string)`

SetDeviceAccountIdentifier sets DeviceAccountIdentifier field to given value.


### GetDeviceAccountNumberSuffix

`func (o *ApplePaySecureElementPass) GetDeviceAccountNumberSuffix() string`

GetDeviceAccountNumberSuffix returns the DeviceAccountNumberSuffix field if non-nil, zero value otherwise.

### GetDeviceAccountNumberSuffixOk

`func (o *ApplePaySecureElementPass) GetDeviceAccountNumberSuffixOk() (*string, bool)`

GetDeviceAccountNumberSuffixOk returns a tuple with the DeviceAccountNumberSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccountNumberSuffix

`func (o *ApplePaySecureElementPass) SetDeviceAccountNumberSuffix(v string)`

SetDeviceAccountNumberSuffix sets DeviceAccountNumberSuffix field to given value.


### GetDevicePassIdentifier

`func (o *ApplePaySecureElementPass) GetDevicePassIdentifier() string`

GetDevicePassIdentifier returns the DevicePassIdentifier field if non-nil, zero value otherwise.

### GetDevicePassIdentifierOk

`func (o *ApplePaySecureElementPass) GetDevicePassIdentifierOk() (*string, bool)`

GetDevicePassIdentifierOk returns a tuple with the DevicePassIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePassIdentifier

`func (o *ApplePaySecureElementPass) SetDevicePassIdentifier(v string)`

SetDevicePassIdentifier sets DevicePassIdentifier field to given value.

### HasDevicePassIdentifier

`func (o *ApplePaySecureElementPass) HasDevicePassIdentifier() bool`

HasDevicePassIdentifier returns a boolean if a field has been set.

### GetPairedTerminalIdentifier

`func (o *ApplePaySecureElementPass) GetPairedTerminalIdentifier() string`

GetPairedTerminalIdentifier returns the PairedTerminalIdentifier field if non-nil, zero value otherwise.

### GetPairedTerminalIdentifierOk

`func (o *ApplePaySecureElementPass) GetPairedTerminalIdentifierOk() (*string, bool)`

GetPairedTerminalIdentifierOk returns a tuple with the PairedTerminalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairedTerminalIdentifier

`func (o *ApplePaySecureElementPass) SetPairedTerminalIdentifier(v string)`

SetPairedTerminalIdentifier sets PairedTerminalIdentifier field to given value.

### HasPairedTerminalIdentifier

`func (o *ApplePaySecureElementPass) HasPairedTerminalIdentifier() bool`

HasPairedTerminalIdentifier returns a boolean if a field has been set.

### GetPassActivationState

`func (o *ApplePaySecureElementPass) GetPassActivationState() string`

GetPassActivationState returns the PassActivationState field if non-nil, zero value otherwise.

### GetPassActivationStateOk

`func (o *ApplePaySecureElementPass) GetPassActivationStateOk() (*string, bool)`

GetPassActivationStateOk returns a tuple with the PassActivationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassActivationState

`func (o *ApplePaySecureElementPass) SetPassActivationState(v string)`

SetPassActivationState sets PassActivationState field to given value.


### GetPrimaryAccountIdentifier

`func (o *ApplePaySecureElementPass) GetPrimaryAccountIdentifier() string`

GetPrimaryAccountIdentifier returns the PrimaryAccountIdentifier field if non-nil, zero value otherwise.

### GetPrimaryAccountIdentifierOk

`func (o *ApplePaySecureElementPass) GetPrimaryAccountIdentifierOk() (*string, bool)`

GetPrimaryAccountIdentifierOk returns a tuple with the PrimaryAccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountIdentifier

`func (o *ApplePaySecureElementPass) SetPrimaryAccountIdentifier(v string)`

SetPrimaryAccountIdentifier sets PrimaryAccountIdentifier field to given value.


### GetPrimaryAccountNumberSuffix

`func (o *ApplePaySecureElementPass) GetPrimaryAccountNumberSuffix() string`

GetPrimaryAccountNumberSuffix returns the PrimaryAccountNumberSuffix field if non-nil, zero value otherwise.

### GetPrimaryAccountNumberSuffixOk

`func (o *ApplePaySecureElementPass) GetPrimaryAccountNumberSuffixOk() (*string, bool)`

GetPrimaryAccountNumberSuffixOk returns a tuple with the PrimaryAccountNumberSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountNumberSuffix

`func (o *ApplePaySecureElementPass) SetPrimaryAccountNumberSuffix(v string)`

SetPrimaryAccountNumberSuffix sets PrimaryAccountNumberSuffix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


