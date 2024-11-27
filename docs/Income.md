# Income

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The amount earned at the specified frequency. For example, $112.35 USD is represented as 11235 cents). | [optional] 
**Currency** | Pointer to **string** | The currency in ISO 4217 format. | [optional] 
**Frequency** | Pointer to [**NullableFrequency**](Frequency.md) |  | [optional] 
**Source** | Pointer to **NullableString** | The source of the income | [optional] 

## Methods

### NewIncome

`func NewIncome() *Income`

NewIncome instantiates a new Income object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeWithDefaults

`func NewIncomeWithDefaults() *Income`

NewIncomeWithDefaults instantiates a new Income object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Income) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Income) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Income) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Income) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *Income) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Income) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Income) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Income) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFrequency

`func (o *Income) GetFrequency() Frequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Income) GetFrequencyOk() (*Frequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Income) SetFrequency(v Frequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *Income) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *Income) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *Income) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetSource

`func (o *Income) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Income) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Income) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Income) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *Income) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *Income) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


