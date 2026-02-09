# AutopayPaymentConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ach** | Pointer to [**AchPaymentConfig**](AchPaymentConfig.md) |  | [optional] 
**InternalTransfer** | Pointer to [**InternalTransferPaymentConfig**](InternalTransferPaymentConfig.md) |  | [optional] 

## Methods

### NewAutopayPaymentConfigs

`func NewAutopayPaymentConfigs() *AutopayPaymentConfigs`

NewAutopayPaymentConfigs instantiates a new AutopayPaymentConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayPaymentConfigsWithDefaults

`func NewAutopayPaymentConfigsWithDefaults() *AutopayPaymentConfigs`

NewAutopayPaymentConfigsWithDefaults instantiates a new AutopayPaymentConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAch

`func (o *AutopayPaymentConfigs) GetAch() AchPaymentConfig`

GetAch returns the Ach field if non-nil, zero value otherwise.

### GetAchOk

`func (o *AutopayPaymentConfigs) GetAchOk() (*AchPaymentConfig, bool)`

GetAchOk returns a tuple with the Ach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAch

`func (o *AutopayPaymentConfigs) SetAch(v AchPaymentConfig)`

SetAch sets Ach field to given value.

### HasAch

`func (o *AutopayPaymentConfigs) HasAch() bool`

HasAch returns a boolean if a field has been set.

### GetInternalTransfer

`func (o *AutopayPaymentConfigs) GetInternalTransfer() InternalTransferPaymentConfig`

GetInternalTransfer returns the InternalTransfer field if non-nil, zero value otherwise.

### GetInternalTransferOk

`func (o *AutopayPaymentConfigs) GetInternalTransferOk() (*InternalTransferPaymentConfig, bool)`

GetInternalTransferOk returns a tuple with the InternalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTransfer

`func (o *AutopayPaymentConfigs) SetInternalTransfer(v InternalTransferPaymentConfig)`

SetInternalTransfer sets InternalTransfer field to given value.

### HasInternalTransfer

`func (o *AutopayPaymentConfigs) HasInternalTransfer() bool`

HasInternalTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


