# AccountTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationType** | Pointer to [**ApplicationType**](ApplicationType.md) |  | [optional] 
**Description** | Pointer to **string** | User provided account template description | [optional] 
**Id** | Pointer to **string** | Generated ID for the template | [optional] [readonly] 
**IsEnabled** | **bool** | Whether this template can be used for account creation | 
**Name** | **string** | Unique account template name | 
**Template** | [**TemplateFields**](TemplateFields.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 

## Methods

### NewAccountTemplate

`func NewAccountTemplate(isEnabled bool, name string, template TemplateFields, ) *AccountTemplate`

NewAccountTemplate instantiates a new AccountTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTemplateWithDefaults

`func NewAccountTemplateWithDefaults() *AccountTemplate`

NewAccountTemplateWithDefaults instantiates a new AccountTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationType

`func (o *AccountTemplate) GetApplicationType() ApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *AccountTemplate) GetApplicationTypeOk() (*ApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *AccountTemplate) SetApplicationType(v ApplicationType)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *AccountTemplate) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetDescription

`func (o *AccountTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AccountTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AccountTemplate) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AccountTemplate) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AccountTemplate) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetName

`func (o *AccountTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetTemplate

`func (o *AccountTemplate) GetTemplate() TemplateFields`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AccountTemplate) GetTemplateOk() (*TemplateFields, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AccountTemplate) SetTemplate(v TemplateFields)`

SetTemplate sets Template field to given value.


### GetTenant

`func (o *AccountTemplate) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountTemplate) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountTemplate) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AccountTemplate) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


