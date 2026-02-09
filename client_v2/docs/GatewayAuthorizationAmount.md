# GatewayAuthorizationAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** |  | 
**ApprovedAmount** | **int64** |  | 
**Currency** | **string** |  | 
**CurrencyConversion** | [**GatewayAuthorizationCurrencyConversion**](GatewayAuthorizationCurrencyConversion.md) |  | 
**TotalPendingAmount** | **int64** |  | 

## Methods

### NewGatewayAuthorizationAmount

`func NewGatewayAuthorizationAmount(amount int64, approvedAmount int64, currency string, currencyConversion GatewayAuthorizationCurrencyConversion, totalPendingAmount int64, ) *GatewayAuthorizationAmount`

NewGatewayAuthorizationAmount instantiates a new GatewayAuthorizationAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayAuthorizationAmountWithDefaults

`func NewGatewayAuthorizationAmountWithDefaults() *GatewayAuthorizationAmount`

NewGatewayAuthorizationAmountWithDefaults instantiates a new GatewayAuthorizationAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GatewayAuthorizationAmount) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GatewayAuthorizationAmount) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GatewayAuthorizationAmount) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetApprovedAmount

`func (o *GatewayAuthorizationAmount) GetApprovedAmount() int64`

GetApprovedAmount returns the ApprovedAmount field if non-nil, zero value otherwise.

### GetApprovedAmountOk

`func (o *GatewayAuthorizationAmount) GetApprovedAmountOk() (*int64, bool)`

GetApprovedAmountOk returns a tuple with the ApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAmount

`func (o *GatewayAuthorizationAmount) SetApprovedAmount(v int64)`

SetApprovedAmount sets ApprovedAmount field to given value.


### GetCurrency

`func (o *GatewayAuthorizationAmount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GatewayAuthorizationAmount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GatewayAuthorizationAmount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrencyConversion

`func (o *GatewayAuthorizationAmount) GetCurrencyConversion() GatewayAuthorizationCurrencyConversion`

GetCurrencyConversion returns the CurrencyConversion field if non-nil, zero value otherwise.

### GetCurrencyConversionOk

`func (o *GatewayAuthorizationAmount) GetCurrencyConversionOk() (*GatewayAuthorizationCurrencyConversion, bool)`

GetCurrencyConversionOk returns a tuple with the CurrencyConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyConversion

`func (o *GatewayAuthorizationAmount) SetCurrencyConversion(v GatewayAuthorizationCurrencyConversion)`

SetCurrencyConversion sets CurrencyConversion field to given value.


### GetTotalPendingAmount

`func (o *GatewayAuthorizationAmount) GetTotalPendingAmount() int64`

GetTotalPendingAmount returns the TotalPendingAmount field if non-nil, zero value otherwise.

### GetTotalPendingAmountOk

`func (o *GatewayAuthorizationAmount) GetTotalPendingAmountOk() (*int64, bool)`

GetTotalPendingAmountOk returns a tuple with the TotalPendingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPendingAmount

`func (o *GatewayAuthorizationAmount) SetTotalPendingAmount(v int64)`

SetTotalPendingAmount sets TotalPendingAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


