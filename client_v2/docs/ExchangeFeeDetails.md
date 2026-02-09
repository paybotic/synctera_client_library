# ExchangeFeeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The amount in the source currency&#39;s minor unit. For example, 10000 would be $100 for USD.  One of the amount or percentage is required.  | [optional] 
**Currency** | **string** | The ISO 4217 currency code of the fee. The currency must match either the source or target currency of the exchange details. If the fee is a percentage of the source amount, the currency must match the source currency. If the fee is a percentage of the target amount, the currency must match the target currency.  | 
**Description** | Pointer to **string** | The description of the fee. | [optional] 
**FeeType** | [**ExchangeFeeDetailsFeeType**](ExchangeFeeDetailsFeeType.md) |  | 
**Percentage** | Pointer to **string** | The percentage of the amount that is the fee. For example, \&quot;0.05403\&quot; would be 5.403%.  One of the amount or percentage is required.  | [optional] 

## Methods

### NewExchangeFeeDetails

`func NewExchangeFeeDetails(currency string, feeType ExchangeFeeDetailsFeeType, ) *ExchangeFeeDetails`

NewExchangeFeeDetails instantiates a new ExchangeFeeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeFeeDetailsWithDefaults

`func NewExchangeFeeDetailsWithDefaults() *ExchangeFeeDetails`

NewExchangeFeeDetailsWithDefaults instantiates a new ExchangeFeeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ExchangeFeeDetails) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExchangeFeeDetails) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExchangeFeeDetails) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExchangeFeeDetails) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *ExchangeFeeDetails) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExchangeFeeDetails) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExchangeFeeDetails) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *ExchangeFeeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExchangeFeeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExchangeFeeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExchangeFeeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeeType

`func (o *ExchangeFeeDetails) GetFeeType() ExchangeFeeDetailsFeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *ExchangeFeeDetails) GetFeeTypeOk() (*ExchangeFeeDetailsFeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *ExchangeFeeDetails) SetFeeType(v ExchangeFeeDetailsFeeType)`

SetFeeType sets FeeType field to given value.


### GetPercentage

`func (o *ExchangeFeeDetails) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ExchangeFeeDetails) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ExchangeFeeDetails) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ExchangeFeeDetails) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


