# CardTransactionDataCurrencyConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionRate** | Pointer to **float32** | The currency conversion rate used | [optional] 
**OriginalAmount** | Pointer to **int32** | The original transaction amount before conversion | [optional] 
**OriginalCurrencyCode** | Pointer to **string** | The original currency code | [optional] 

## Methods

### NewCardTransactionDataCurrencyConversion

`func NewCardTransactionDataCurrencyConversion() *CardTransactionDataCurrencyConversion`

NewCardTransactionDataCurrencyConversion instantiates a new CardTransactionDataCurrencyConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTransactionDataCurrencyConversionWithDefaults

`func NewCardTransactionDataCurrencyConversionWithDefaults() *CardTransactionDataCurrencyConversion`

NewCardTransactionDataCurrencyConversionWithDefaults instantiates a new CardTransactionDataCurrencyConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionRate

`func (o *CardTransactionDataCurrencyConversion) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *CardTransactionDataCurrencyConversion) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *CardTransactionDataCurrencyConversion) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *CardTransactionDataCurrencyConversion) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *CardTransactionDataCurrencyConversion) GetOriginalAmount() int32`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *CardTransactionDataCurrencyConversion) GetOriginalAmountOk() (*int32, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *CardTransactionDataCurrencyConversion) SetOriginalAmount(v int32)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *CardTransactionDataCurrencyConversion) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetOriginalCurrencyCode

`func (o *CardTransactionDataCurrencyConversion) GetOriginalCurrencyCode() string`

GetOriginalCurrencyCode returns the OriginalCurrencyCode field if non-nil, zero value otherwise.

### GetOriginalCurrencyCodeOk

`func (o *CardTransactionDataCurrencyConversion) GetOriginalCurrencyCodeOk() (*string, bool)`

GetOriginalCurrencyCodeOk returns a tuple with the OriginalCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCurrencyCode

`func (o *CardTransactionDataCurrencyConversion) SetOriginalCurrencyCode(v string)`

SetOriginalCurrencyCode sets OriginalCurrencyCode field to given value.

### HasOriginalCurrencyCode

`func (o *CardTransactionDataCurrencyConversion) HasOriginalCurrencyCode() bool`

HasOriginalCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


