# ResidentialExpense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The amount paid in residential expenses at the specified frequency. For example, $112.35 USD is represented as 11235 cents). | [optional] 
**Currency** | Pointer to **string** | The currency in ISO 4217 format. | [optional] 
**Frequency** | Pointer to [**NullableFrequency**](Frequency.md) |  | [optional] 

## Methods

### NewResidentialExpense

`func NewResidentialExpense() *ResidentialExpense`

NewResidentialExpense instantiates a new ResidentialExpense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResidentialExpenseWithDefaults

`func NewResidentialExpenseWithDefaults() *ResidentialExpense`

NewResidentialExpenseWithDefaults instantiates a new ResidentialExpense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ResidentialExpense) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResidentialExpense) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResidentialExpense) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ResidentialExpense) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *ResidentialExpense) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ResidentialExpense) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ResidentialExpense) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ResidentialExpense) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFrequency

`func (o *ResidentialExpense) GetFrequency() Frequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ResidentialExpense) GetFrequencyOk() (*Frequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ResidentialExpense) SetFrequency(v Frequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ResidentialExpense) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *ResidentialExpense) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *ResidentialExpense) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


