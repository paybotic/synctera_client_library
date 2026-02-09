# SpendControlTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** | The number of days to define a rolling window for a spend control | 
**TimeRangeType** | [**SpendControlTimeRangeType**](SpendControlTimeRangeType.md) |  | 
**Period** | **string** |  | 
**Value** | **int32** |  | 

## Methods

### NewSpendControlTimeRange

`func NewSpendControlTimeRange(days int32, timeRangeType SpendControlTimeRangeType, period string, value int32, ) *SpendControlTimeRange`

NewSpendControlTimeRange instantiates a new SpendControlTimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlTimeRangeWithDefaults

`func NewSpendControlTimeRangeWithDefaults() *SpendControlTimeRange`

NewSpendControlTimeRangeWithDefaults instantiates a new SpendControlTimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *SpendControlTimeRange) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *SpendControlTimeRange) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *SpendControlTimeRange) SetDays(v int32)`

SetDays sets Days field to given value.


### GetTimeRangeType

`func (o *SpendControlTimeRange) GetTimeRangeType() SpendControlTimeRangeType`

GetTimeRangeType returns the TimeRangeType field if non-nil, zero value otherwise.

### GetTimeRangeTypeOk

`func (o *SpendControlTimeRange) GetTimeRangeTypeOk() (*SpendControlTimeRangeType, bool)`

GetTimeRangeTypeOk returns a tuple with the TimeRangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeType

`func (o *SpendControlTimeRange) SetTimeRangeType(v SpendControlTimeRangeType)`

SetTimeRangeType sets TimeRangeType field to given value.


### GetPeriod

`func (o *SpendControlTimeRange) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SpendControlTimeRange) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SpendControlTimeRange) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetValue

`func (o *SpendControlTimeRange) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SpendControlTimeRange) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SpendControlTimeRange) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


