# CalculationBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**Condition**](Condition.md) |  | [optional] 
**ParameterLabel** | Pointer to **string** | The user-friendly label for the parameter | [optional] 
**ParameterName** | Pointer to **string** | The name of the parameter that was used in the score calculation | [optional] 
**ParameterScore** | **float64** | The calculated weight of the risk param. calculated_weight &#x3D; parameter weight/sum(parameters weight) * 100 | 
**ParameterWeight** | **int32** | The weight of the risk param as defined in the risk config | 

## Methods

### NewCalculationBreakdown

`func NewCalculationBreakdown(parameterScore float64, parameterWeight int32, ) *CalculationBreakdown`

NewCalculationBreakdown instantiates a new CalculationBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculationBreakdownWithDefaults

`func NewCalculationBreakdownWithDefaults() *CalculationBreakdown`

NewCalculationBreakdownWithDefaults instantiates a new CalculationBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *CalculationBreakdown) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CalculationBreakdown) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CalculationBreakdown) SetCondition(v Condition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CalculationBreakdown) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetParameterLabel

`func (o *CalculationBreakdown) GetParameterLabel() string`

GetParameterLabel returns the ParameterLabel field if non-nil, zero value otherwise.

### GetParameterLabelOk

`func (o *CalculationBreakdown) GetParameterLabelOk() (*string, bool)`

GetParameterLabelOk returns a tuple with the ParameterLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterLabel

`func (o *CalculationBreakdown) SetParameterLabel(v string)`

SetParameterLabel sets ParameterLabel field to given value.

### HasParameterLabel

`func (o *CalculationBreakdown) HasParameterLabel() bool`

HasParameterLabel returns a boolean if a field has been set.

### GetParameterName

`func (o *CalculationBreakdown) GetParameterName() string`

GetParameterName returns the ParameterName field if non-nil, zero value otherwise.

### GetParameterNameOk

`func (o *CalculationBreakdown) GetParameterNameOk() (*string, bool)`

GetParameterNameOk returns a tuple with the ParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterName

`func (o *CalculationBreakdown) SetParameterName(v string)`

SetParameterName sets ParameterName field to given value.

### HasParameterName

`func (o *CalculationBreakdown) HasParameterName() bool`

HasParameterName returns a boolean if a field has been set.

### GetParameterScore

`func (o *CalculationBreakdown) GetParameterScore() float64`

GetParameterScore returns the ParameterScore field if non-nil, zero value otherwise.

### GetParameterScoreOk

`func (o *CalculationBreakdown) GetParameterScoreOk() (*float64, bool)`

GetParameterScoreOk returns a tuple with the ParameterScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterScore

`func (o *CalculationBreakdown) SetParameterScore(v float64)`

SetParameterScore sets ParameterScore field to given value.


### GetParameterWeight

`func (o *CalculationBreakdown) GetParameterWeight() int32`

GetParameterWeight returns the ParameterWeight field if non-nil, zero value otherwise.

### GetParameterWeightOk

`func (o *CalculationBreakdown) GetParameterWeightOk() (*int32, bool)`

GetParameterWeightOk returns a tuple with the ParameterWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterWeight

`func (o *CalculationBreakdown) SetParameterWeight(v int32)`

SetParameterWeight sets ParameterWeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


