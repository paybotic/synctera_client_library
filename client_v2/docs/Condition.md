# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lower** | Pointer to **int32** | If the condition type is RANGE. condition_lower is the field that stores the lower value of the condition | [optional] 
**Type** | [**ConditionType**](ConditionType.md) |  | 
**Upper** | Pointer to **int32** | If the condition type is RANGE. condition_upper is the field that stores the upper value of the condition | [optional] 
**Value** | Pointer to **string** | If the condition type is CATEGORICAL. condition_value is the field that stores the discrete value | [optional] 
**Weight** | **int32** | The weight of the condition as specific in the risk config. | 

## Methods

### NewCondition

`func NewCondition(type_ ConditionType, weight int32, ) *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLower

`func (o *Condition) GetLower() int32`

GetLower returns the Lower field if non-nil, zero value otherwise.

### GetLowerOk

`func (o *Condition) GetLowerOk() (*int32, bool)`

GetLowerOk returns a tuple with the Lower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLower

`func (o *Condition) SetLower(v int32)`

SetLower sets Lower field to given value.

### HasLower

`func (o *Condition) HasLower() bool`

HasLower returns a boolean if a field has been set.

### GetType

`func (o *Condition) GetType() ConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Condition) GetTypeOk() (*ConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Condition) SetType(v ConditionType)`

SetType sets Type field to given value.


### GetUpper

`func (o *Condition) GetUpper() int32`

GetUpper returns the Upper field if non-nil, zero value otherwise.

### GetUpperOk

`func (o *Condition) GetUpperOk() (*int32, bool)`

GetUpperOk returns a tuple with the Upper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpper

`func (o *Condition) SetUpper(v int32)`

SetUpper sets Upper field to given value.

### HasUpper

`func (o *Condition) HasUpper() bool`

HasUpper returns a boolean if a field has been set.

### GetValue

`func (o *Condition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Condition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Condition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Condition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWeight

`func (o *Condition) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Condition) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Condition) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


