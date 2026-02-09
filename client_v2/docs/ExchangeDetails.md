# ExchangeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fees** | Pointer to [**[]ExchangeFeeDetails**](ExchangeFeeDetails.md) | The fees associated with the exchange.  | [optional] 
**Rate** | **string** | The exchange rate from source to target currency. For example: 1.30445  | 
**SourceAmount** | **int64** | The amount in the source currency&#39;s minor unit. For example, 10000 would be $100 for USD. This is the amount inclusive of fees.  | 
**SourceCurrency** | **string** | The ISO 4217 currency code | 
**TargetAmount** | **int64** | The amount in the target currency&#39;s minor unit. For example, 13045 would be £130.45 for GBP. This is the amount inclusive of fees.  | 
**TargetCurrency** | **string** | The ISO 4217 currency code | 

## Methods

### NewExchangeDetails

`func NewExchangeDetails(rate string, sourceAmount int64, sourceCurrency string, targetAmount int64, targetCurrency string, ) *ExchangeDetails`

NewExchangeDetails instantiates a new ExchangeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeDetailsWithDefaults

`func NewExchangeDetailsWithDefaults() *ExchangeDetails`

NewExchangeDetailsWithDefaults instantiates a new ExchangeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFees

`func (o *ExchangeDetails) GetFees() []ExchangeFeeDetails`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *ExchangeDetails) GetFeesOk() (*[]ExchangeFeeDetails, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *ExchangeDetails) SetFees(v []ExchangeFeeDetails)`

SetFees sets Fees field to given value.

### HasFees

`func (o *ExchangeDetails) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetRate

`func (o *ExchangeDetails) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ExchangeDetails) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ExchangeDetails) SetRate(v string)`

SetRate sets Rate field to given value.


### GetSourceAmount

`func (o *ExchangeDetails) GetSourceAmount() int64`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *ExchangeDetails) GetSourceAmountOk() (*int64, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *ExchangeDetails) SetSourceAmount(v int64)`

SetSourceAmount sets SourceAmount field to given value.


### GetSourceCurrency

`func (o *ExchangeDetails) GetSourceCurrency() string`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *ExchangeDetails) GetSourceCurrencyOk() (*string, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *ExchangeDetails) SetSourceCurrency(v string)`

SetSourceCurrency sets SourceCurrency field to given value.


### GetTargetAmount

`func (o *ExchangeDetails) GetTargetAmount() int64`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *ExchangeDetails) GetTargetAmountOk() (*int64, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *ExchangeDetails) SetTargetAmount(v int64)`

SetTargetAmount sets TargetAmount field to given value.


### GetTargetCurrency

`func (o *ExchangeDetails) GetTargetCurrency() string`

GetTargetCurrency returns the TargetCurrency field if non-nil, zero value otherwise.

### GetTargetCurrencyOk

`func (o *ExchangeDetails) GetTargetCurrencyOk() (*string, bool)`

GetTargetCurrencyOk returns a tuple with the TargetCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCurrency

`func (o *ExchangeDetails) SetTargetCurrency(v string)`

SetTargetCurrency sets TargetCurrency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


