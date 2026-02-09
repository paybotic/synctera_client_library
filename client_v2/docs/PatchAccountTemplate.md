# PatchAccountTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | User provided account template description | [optional] 
**IsEnabled** | Pointer to **bool** | Whether this template can be used for account creation | [optional] 
**Name** | Pointer to **string** | Unique account template name | [optional] 
**Template** | Pointer to [**PatchAccountTemplateFields**](PatchAccountTemplateFields.md) |  | [optional] 

## Methods

### NewPatchAccountTemplate

`func NewPatchAccountTemplate() *PatchAccountTemplate`

NewPatchAccountTemplate instantiates a new PatchAccountTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchAccountTemplateWithDefaults

`func NewPatchAccountTemplateWithDefaults() *PatchAccountTemplate`

NewPatchAccountTemplateWithDefaults instantiates a new PatchAccountTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PatchAccountTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchAccountTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchAccountTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchAccountTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PatchAccountTemplate) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PatchAccountTemplate) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PatchAccountTemplate) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PatchAccountTemplate) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *PatchAccountTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchAccountTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchAccountTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchAccountTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplate

`func (o *PatchAccountTemplate) GetTemplate() PatchAccountTemplateFields`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PatchAccountTemplate) GetTemplateOk() (*PatchAccountTemplateFields, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PatchAccountTemplate) SetTemplate(v PatchAccountTemplateFields)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PatchAccountTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


