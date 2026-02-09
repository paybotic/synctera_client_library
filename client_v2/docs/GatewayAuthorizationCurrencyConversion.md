# GatewayAuthorizationCurrencyConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionRate** | Pointer to **float32** |  | [optional] 
**DynamicCurrencyConversion** | Pointer to **bool** |  | [optional] 
**OriginalAmount** | Pointer to **int64** |  | [optional] 
**OriginalCurrencyCode** | Pointer to **string** |  | [optional] 
**OriginalCurrencyCodeAlpha** | Pointer to **string** | ISO 4217  Alpha-3 currency code | [optional] 
**RawOriginalAmount** | Pointer to **float32** |  | [optional] 
**SettlementCurrencyConversion** | Pointer to [**GatewayAuthorizationSettlementCurrencyConversion**](GatewayAuthorizationSettlementCurrencyConversion.md) |  | [optional] 

## Methods

### NewGatewayAuthorizationCurrencyConversion

`func NewGatewayAuthorizationCurrencyConversion() *GatewayAuthorizationCurrencyConversion`

NewGatewayAuthorizationCurrencyConversion instantiates a new GatewayAuthorizationCurrencyConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayAuthorizationCurrencyConversionWithDefaults

`func NewGatewayAuthorizationCurrencyConversionWithDefaults() *GatewayAuthorizationCurrencyConversion`

NewGatewayAuthorizationCurrencyConversionWithDefaults instantiates a new GatewayAuthorizationCurrencyConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionRate

`func (o *GatewayAuthorizationCurrencyConversion) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *GatewayAuthorizationCurrencyConversion) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *GatewayAuthorizationCurrencyConversion) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *GatewayAuthorizationCurrencyConversion) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetDynamicCurrencyConversion

`func (o *GatewayAuthorizationCurrencyConversion) GetDynamicCurrencyConversion() bool`

GetDynamicCurrencyConversion returns the DynamicCurrencyConversion field if non-nil, zero value otherwise.

### GetDynamicCurrencyConversionOk

`func (o *GatewayAuthorizationCurrencyConversion) GetDynamicCurrencyConversionOk() (*bool, bool)`

GetDynamicCurrencyConversionOk returns a tuple with the DynamicCurrencyConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicCurrencyConversion

`func (o *GatewayAuthorizationCurrencyConversion) SetDynamicCurrencyConversion(v bool)`

SetDynamicCurrencyConversion sets DynamicCurrencyConversion field to given value.

### HasDynamicCurrencyConversion

`func (o *GatewayAuthorizationCurrencyConversion) HasDynamicCurrencyConversion() bool`

HasDynamicCurrencyConversion returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *GatewayAuthorizationCurrencyConversion) GetOriginalAmount() int64`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *GatewayAuthorizationCurrencyConversion) GetOriginalAmountOk() (*int64, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *GatewayAuthorizationCurrencyConversion) SetOriginalAmount(v int64)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *GatewayAuthorizationCurrencyConversion) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetOriginalCurrencyCode

`func (o *GatewayAuthorizationCurrencyConversion) GetOriginalCurrencyCode() string`

GetOriginalCurrencyCode returns the OriginalCurrencyCode field if non-nil, zero value otherwise.

### GetOriginalCurrencyCodeOk

`func (o *GatewayAuthorizationCurrencyConversion) GetOriginalCurrencyCodeOk() (*string, bool)`

GetOriginalCurrencyCodeOk returns a tuple with the OriginalCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCurrencyCode

`func (o *GatewayAuthorizationCurrencyConversion) SetOriginalCurrencyCode(v string)`

SetOriginalCurrencyCode sets OriginalCurrencyCode field to given value.

### HasOriginalCurrencyCode

`func (o *GatewayAuthorizationCurrencyConversion) HasOriginalCurrencyCode() bool`

HasOriginalCurrencyCode returns a boolean if a field has been set.

### GetOriginalCurrencyCodeAlpha

`func (o *GatewayAuthorizationCurrencyConversion) GetOriginalCurrencyCodeAlpha() string`

GetOriginalCurrencyCodeAlpha returns the OriginalCurrencyCodeAlpha field if non-nil, zero value otherwise.

### GetOriginalCurrencyCodeAlphaOk

`func (o *GatewayAuthorizationCurrencyConversion) GetOriginalCurrencyCodeAlphaOk() (*string, bool)`

GetOriginalCurrencyCodeAlphaOk returns a tuple with the OriginalCurrencyCodeAlpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCurrencyCodeAlpha

`func (o *GatewayAuthorizationCurrencyConversion) SetOriginalCurrencyCodeAlpha(v string)`

SetOriginalCurrencyCodeAlpha sets OriginalCurrencyCodeAlpha field to given value.

### HasOriginalCurrencyCodeAlpha

`func (o *GatewayAuthorizationCurrencyConversion) HasOriginalCurrencyCodeAlpha() bool`

HasOriginalCurrencyCodeAlpha returns a boolean if a field has been set.

### GetRawOriginalAmount

`func (o *GatewayAuthorizationCurrencyConversion) GetRawOriginalAmount() float32`

GetRawOriginalAmount returns the RawOriginalAmount field if non-nil, zero value otherwise.

### GetRawOriginalAmountOk

`func (o *GatewayAuthorizationCurrencyConversion) GetRawOriginalAmountOk() (*float32, bool)`

GetRawOriginalAmountOk returns a tuple with the RawOriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawOriginalAmount

`func (o *GatewayAuthorizationCurrencyConversion) SetRawOriginalAmount(v float32)`

SetRawOriginalAmount sets RawOriginalAmount field to given value.

### HasRawOriginalAmount

`func (o *GatewayAuthorizationCurrencyConversion) HasRawOriginalAmount() bool`

HasRawOriginalAmount returns a boolean if a field has been set.

### GetSettlementCurrencyConversion

`func (o *GatewayAuthorizationCurrencyConversion) GetSettlementCurrencyConversion() GatewayAuthorizationSettlementCurrencyConversion`

GetSettlementCurrencyConversion returns the SettlementCurrencyConversion field if non-nil, zero value otherwise.

### GetSettlementCurrencyConversionOk

`func (o *GatewayAuthorizationCurrencyConversion) GetSettlementCurrencyConversionOk() (*GatewayAuthorizationSettlementCurrencyConversion, bool)`

GetSettlementCurrencyConversionOk returns a tuple with the SettlementCurrencyConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCurrencyConversion

`func (o *GatewayAuthorizationCurrencyConversion) SetSettlementCurrencyConversion(v GatewayAuthorizationSettlementCurrencyConversion)`

SetSettlementCurrencyConversion sets SettlementCurrencyConversion field to given value.

### HasSettlementCurrencyConversion

`func (o *GatewayAuthorizationCurrencyConversion) HasSettlementCurrencyConversion() bool`

HasSettlementCurrencyConversion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


