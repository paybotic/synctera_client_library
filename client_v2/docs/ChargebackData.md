# ChargebackData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount eligible for chargeback in cents. | 
**Currency** | **string** | ISO 4217 Alpha-3 currency code | 

## Methods

### NewChargebackData

`func NewChargebackData(amount int64, currency string, ) *ChargebackData`

NewChargebackData instantiates a new ChargebackData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargebackDataWithDefaults

`func NewChargebackDataWithDefaults() *ChargebackData`

NewChargebackDataWithDefaults instantiates a new ChargebackData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ChargebackData) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ChargebackData) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ChargebackData) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *ChargebackData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChargebackData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChargebackData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


