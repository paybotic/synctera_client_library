# BatchPaymentTemplatePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**BatchPaymentTemplateConfig**](BatchPaymentTemplateConfig.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the template is enabled. If the template is not enabled, it will not be used when creating a batch transfer.  | [optional] 
**ExtraRailParams** | Pointer to **map[string]interface{}** | Additional parameters to be passed through to the payment rail.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**BatchPaymentTemplateRules**](BatchPaymentTemplateRules.md) |  | [optional] 

## Methods

### NewBatchPaymentTemplatePatch

`func NewBatchPaymentTemplatePatch() *BatchPaymentTemplatePatch`

NewBatchPaymentTemplatePatch instantiates a new BatchPaymentTemplatePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchPaymentTemplatePatchWithDefaults

`func NewBatchPaymentTemplatePatchWithDefaults() *BatchPaymentTemplatePatch`

NewBatchPaymentTemplatePatchWithDefaults instantiates a new BatchPaymentTemplatePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *BatchPaymentTemplatePatch) GetConfig() BatchPaymentTemplateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BatchPaymentTemplatePatch) GetConfigOk() (*BatchPaymentTemplateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BatchPaymentTemplatePatch) SetConfig(v BatchPaymentTemplateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BatchPaymentTemplatePatch) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *BatchPaymentTemplatePatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BatchPaymentTemplatePatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BatchPaymentTemplatePatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BatchPaymentTemplatePatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BatchPaymentTemplatePatch) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BatchPaymentTemplatePatch) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BatchPaymentTemplatePatch) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BatchPaymentTemplatePatch) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraRailParams

`func (o *BatchPaymentTemplatePatch) GetExtraRailParams() map[string]interface{}`

GetExtraRailParams returns the ExtraRailParams field if non-nil, zero value otherwise.

### GetExtraRailParamsOk

`func (o *BatchPaymentTemplatePatch) GetExtraRailParamsOk() (*map[string]interface{}, bool)`

GetExtraRailParamsOk returns a tuple with the ExtraRailParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRailParams

`func (o *BatchPaymentTemplatePatch) SetExtraRailParams(v map[string]interface{})`

SetExtraRailParams sets ExtraRailParams field to given value.

### HasExtraRailParams

`func (o *BatchPaymentTemplatePatch) HasExtraRailParams() bool`

HasExtraRailParams returns a boolean if a field has been set.

### GetName

`func (o *BatchPaymentTemplatePatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchPaymentTemplatePatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchPaymentTemplatePatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BatchPaymentTemplatePatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *BatchPaymentTemplatePatch) GetRules() BatchPaymentTemplateRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BatchPaymentTemplatePatch) GetRulesOk() (*BatchPaymentTemplateRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BatchPaymentTemplatePatch) SetRules(v BatchPaymentTemplateRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *BatchPaymentTemplatePatch) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


