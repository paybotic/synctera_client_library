# SpendControlRollingWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | **string** |  | 
**TimeRangeType** | [**SpendControlTimeRangeType**](SpendControlTimeRangeType.md) |  | 
**Value** | **int32** |  | 

## Methods

### NewSpendControlRollingWindow

`func NewSpendControlRollingWindow(period string, timeRangeType SpendControlTimeRangeType, value int32, ) *SpendControlRollingWindow`

NewSpendControlRollingWindow instantiates a new SpendControlRollingWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlRollingWindowWithDefaults

`func NewSpendControlRollingWindowWithDefaults() *SpendControlRollingWindow`

NewSpendControlRollingWindowWithDefaults instantiates a new SpendControlRollingWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *SpendControlRollingWindow) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SpendControlRollingWindow) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SpendControlRollingWindow) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetTimeRangeType

`func (o *SpendControlRollingWindow) GetTimeRangeType() SpendControlTimeRangeType`

GetTimeRangeType returns the TimeRangeType field if non-nil, zero value otherwise.

### GetTimeRangeTypeOk

`func (o *SpendControlRollingWindow) GetTimeRangeTypeOk() (*SpendControlTimeRangeType, bool)`

GetTimeRangeTypeOk returns a tuple with the TimeRangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeType

`func (o *SpendControlRollingWindow) SetTimeRangeType(v SpendControlTimeRangeType)`

SetTimeRangeType sets TimeRangeType field to given value.


### GetValue

`func (o *SpendControlRollingWindow) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SpendControlRollingWindow) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SpendControlRollingWindow) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


