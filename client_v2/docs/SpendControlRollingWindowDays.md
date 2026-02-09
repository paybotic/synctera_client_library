# SpendControlRollingWindowDays

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** | The number of days to define a rolling window for a spend control | 
**TimeRangeType** | [**SpendControlTimeRangeType**](SpendControlTimeRangeType.md) |  | 

## Methods

### NewSpendControlRollingWindowDays

`func NewSpendControlRollingWindowDays(days int32, timeRangeType SpendControlTimeRangeType, ) *SpendControlRollingWindowDays`

NewSpendControlRollingWindowDays instantiates a new SpendControlRollingWindowDays object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlRollingWindowDaysWithDefaults

`func NewSpendControlRollingWindowDaysWithDefaults() *SpendControlRollingWindowDays`

NewSpendControlRollingWindowDaysWithDefaults instantiates a new SpendControlRollingWindowDays object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *SpendControlRollingWindowDays) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *SpendControlRollingWindowDays) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *SpendControlRollingWindowDays) SetDays(v int32)`

SetDays sets Days field to given value.


### GetTimeRangeType

`func (o *SpendControlRollingWindowDays) GetTimeRangeType() SpendControlTimeRangeType`

GetTimeRangeType returns the TimeRangeType field if non-nil, zero value otherwise.

### GetTimeRangeTypeOk

`func (o *SpendControlRollingWindowDays) GetTimeRangeTypeOk() (*SpendControlTimeRangeType, bool)`

GetTimeRangeTypeOk returns a tuple with the TimeRangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeType

`func (o *SpendControlRollingWindowDays) SetTimeRangeType(v SpendControlTimeRangeType)`

SetTimeRangeType sets TimeRangeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


