# BatchPaymentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**BatchPaymentTemplateConfig**](BatchPaymentTemplateConfig.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the template is enabled. If the template is not enabled, it will not be used when creating a batch transfer.  | [optional] 
**ExtraRailParams** | Pointer to **map[string]interface{}** | Additional parameters to be passed through to the payment rail.  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Rules** | [**BatchPaymentTemplateRules**](BatchPaymentTemplateRules.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewBatchPaymentTemplate

`func NewBatchPaymentTemplate(config BatchPaymentTemplateConfig, name string, rules BatchPaymentTemplateRules, tenant string, ) *BatchPaymentTemplate`

NewBatchPaymentTemplate instantiates a new BatchPaymentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchPaymentTemplateWithDefaults

`func NewBatchPaymentTemplateWithDefaults() *BatchPaymentTemplate`

NewBatchPaymentTemplateWithDefaults instantiates a new BatchPaymentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *BatchPaymentTemplate) GetConfig() BatchPaymentTemplateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BatchPaymentTemplate) GetConfigOk() (*BatchPaymentTemplateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BatchPaymentTemplate) SetConfig(v BatchPaymentTemplateConfig)`

SetConfig sets Config field to given value.


### GetDescription

`func (o *BatchPaymentTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BatchPaymentTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BatchPaymentTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BatchPaymentTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BatchPaymentTemplate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BatchPaymentTemplate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BatchPaymentTemplate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BatchPaymentTemplate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraRailParams

`func (o *BatchPaymentTemplate) GetExtraRailParams() map[string]interface{}`

GetExtraRailParams returns the ExtraRailParams field if non-nil, zero value otherwise.

### GetExtraRailParamsOk

`func (o *BatchPaymentTemplate) GetExtraRailParamsOk() (*map[string]interface{}, bool)`

GetExtraRailParamsOk returns a tuple with the ExtraRailParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRailParams

`func (o *BatchPaymentTemplate) SetExtraRailParams(v map[string]interface{})`

SetExtraRailParams sets ExtraRailParams field to given value.

### HasExtraRailParams

`func (o *BatchPaymentTemplate) HasExtraRailParams() bool`

HasExtraRailParams returns a boolean if a field has been set.

### GetId

`func (o *BatchPaymentTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchPaymentTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchPaymentTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BatchPaymentTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BatchPaymentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchPaymentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchPaymentTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetRules

`func (o *BatchPaymentTemplate) GetRules() BatchPaymentTemplateRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BatchPaymentTemplate) GetRulesOk() (*BatchPaymentTemplateRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BatchPaymentTemplate) SetRules(v BatchPaymentTemplateRules)`

SetRules sets Rules field to given value.


### GetTenant

`func (o *BatchPaymentTemplate) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BatchPaymentTemplate) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BatchPaymentTemplate) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


