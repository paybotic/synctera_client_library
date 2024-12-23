# AccountTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationType** | Pointer to [**ApplicationType**](ApplicationType.md) |  | [optional] 
**Description** | Pointer to **string** | Account template description | [optional] 
**Id** | Pointer to **string** | Generated ID for the template | [optional] 
**IsEnabled** | **bool** | Whether this template can be used for account creation | 
**Name** | **string** | Unique account template name | 
**Template** | [**TemplateFieldsGenericResponse**](TemplateFieldsGenericResponse.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 

## Methods

### NewAccountTemplateResponse

`func NewAccountTemplateResponse(isEnabled bool, name string, template TemplateFieldsGenericResponse, ) *AccountTemplateResponse`

NewAccountTemplateResponse instantiates a new AccountTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTemplateResponseWithDefaults

`func NewAccountTemplateResponseWithDefaults() *AccountTemplateResponse`

NewAccountTemplateResponseWithDefaults instantiates a new AccountTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationType

`func (o *AccountTemplateResponse) GetApplicationType() ApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *AccountTemplateResponse) GetApplicationTypeOk() (*ApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *AccountTemplateResponse) SetApplicationType(v ApplicationType)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *AccountTemplateResponse) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetDescription

`func (o *AccountTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountTemplateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AccountTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountTemplateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountTemplateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AccountTemplateResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AccountTemplateResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AccountTemplateResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetName

`func (o *AccountTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetTemplate

`func (o *AccountTemplateResponse) GetTemplate() TemplateFieldsGenericResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AccountTemplateResponse) GetTemplateOk() (*TemplateFieldsGenericResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AccountTemplateResponse) SetTemplate(v TemplateFieldsGenericResponse)`

SetTemplate sets Template field to given value.


### GetTenant

`func (o *AccountTemplateResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountTemplateResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountTemplateResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AccountTemplateResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


