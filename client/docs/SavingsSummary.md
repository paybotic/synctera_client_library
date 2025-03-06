# SavingsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apy** | Pointer to **float32** | The annual percentage yield (APY) for this account for this statement period, rounded to two decimal points. For example, an APY of 5.5% will display as 5.50.  | [optional] 
**InterestEarned** | Pointer to **int64** | The total interest earned by the depository account for this statement period in ISO 4217 minor currency units. For example, $1.50 USD of interest will be displayed as 150.  | [optional] 
**InterestEarnedPreviousMonth** | Pointer to **int64** | The total interest earned by the depository account in the previous statement period in ISO 4217 minor currency units. For example, $1.50 USD of interest will be displayed as 150.  | [optional] 
**InterestEarnedPreviousYear** | Pointer to **int64** | The total interest earned by the depository account in the previous year in ISO 4217 minor currency units. For example, $100 USD of interest will be displayed as 10000.  | [optional] 
**InterestEarnedYtd** | Pointer to **int64** | The total interest earned by the depository account for this year to date in ISO 4217 minor currency units. For example, $100 USD of interest will be displayed as 10000.  | [optional] 

## Methods

### NewSavingsSummary

`func NewSavingsSummary() *SavingsSummary`

NewSavingsSummary instantiates a new SavingsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavingsSummaryWithDefaults

`func NewSavingsSummaryWithDefaults() *SavingsSummary`

NewSavingsSummaryWithDefaults instantiates a new SavingsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApy

`func (o *SavingsSummary) GetApy() float32`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *SavingsSummary) GetApyOk() (*float32, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *SavingsSummary) SetApy(v float32)`

SetApy sets Apy field to given value.

### HasApy

`func (o *SavingsSummary) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetInterestEarned

`func (o *SavingsSummary) GetInterestEarned() int64`

GetInterestEarned returns the InterestEarned field if non-nil, zero value otherwise.

### GetInterestEarnedOk

`func (o *SavingsSummary) GetInterestEarnedOk() (*int64, bool)`

GetInterestEarnedOk returns a tuple with the InterestEarned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestEarned

`func (o *SavingsSummary) SetInterestEarned(v int64)`

SetInterestEarned sets InterestEarned field to given value.

### HasInterestEarned

`func (o *SavingsSummary) HasInterestEarned() bool`

HasInterestEarned returns a boolean if a field has been set.

### GetInterestEarnedPreviousMonth

`func (o *SavingsSummary) GetInterestEarnedPreviousMonth() int64`

GetInterestEarnedPreviousMonth returns the InterestEarnedPreviousMonth field if non-nil, zero value otherwise.

### GetInterestEarnedPreviousMonthOk

`func (o *SavingsSummary) GetInterestEarnedPreviousMonthOk() (*int64, bool)`

GetInterestEarnedPreviousMonthOk returns a tuple with the InterestEarnedPreviousMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestEarnedPreviousMonth

`func (o *SavingsSummary) SetInterestEarnedPreviousMonth(v int64)`

SetInterestEarnedPreviousMonth sets InterestEarnedPreviousMonth field to given value.

### HasInterestEarnedPreviousMonth

`func (o *SavingsSummary) HasInterestEarnedPreviousMonth() bool`

HasInterestEarnedPreviousMonth returns a boolean if a field has been set.

### GetInterestEarnedPreviousYear

`func (o *SavingsSummary) GetInterestEarnedPreviousYear() int64`

GetInterestEarnedPreviousYear returns the InterestEarnedPreviousYear field if non-nil, zero value otherwise.

### GetInterestEarnedPreviousYearOk

`func (o *SavingsSummary) GetInterestEarnedPreviousYearOk() (*int64, bool)`

GetInterestEarnedPreviousYearOk returns a tuple with the InterestEarnedPreviousYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestEarnedPreviousYear

`func (o *SavingsSummary) SetInterestEarnedPreviousYear(v int64)`

SetInterestEarnedPreviousYear sets InterestEarnedPreviousYear field to given value.

### HasInterestEarnedPreviousYear

`func (o *SavingsSummary) HasInterestEarnedPreviousYear() bool`

HasInterestEarnedPreviousYear returns a boolean if a field has been set.

### GetInterestEarnedYtd

`func (o *SavingsSummary) GetInterestEarnedYtd() int64`

GetInterestEarnedYtd returns the InterestEarnedYtd field if non-nil, zero value otherwise.

### GetInterestEarnedYtdOk

`func (o *SavingsSummary) GetInterestEarnedYtdOk() (*int64, bool)`

GetInterestEarnedYtdOk returns a tuple with the InterestEarnedYtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestEarnedYtd

`func (o *SavingsSummary) SetInterestEarnedYtd(v int64)`

SetInterestEarnedYtd sets InterestEarnedYtd field to given value.

### HasInterestEarnedYtd

`func (o *SavingsSummary) HasInterestEarnedYtd() bool`

HasInterestEarnedYtd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


