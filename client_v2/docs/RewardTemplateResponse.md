# RewardTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The default amount of the reward in ISO 4217 minor currency units, e.g. cents. The internal account will be debited this amount. Can be overridden when creating a reward.  | [optional] 
**Currency** | **string** | currency for the reward, as a three character ISO 4217 alphabetic currency code. | 
**Description** | Pointer to **string** | The description of the reward template. | [optional] 
**InternalAccountId** | **string** | The ID of default internal_account to use as the source of the reward transfer. Cannot be a system internal account. | 
**IsEnabled** | **bool** | Whether the reward template is enabled. If false, rewards cannot be created from this template.  | 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**CreationTime** | **time.Time** | The timestamp representing when the reward template was created | [readonly] 
**Id** | **string** | The ID of the reward template. | 
**LastUpdatedTime** | **time.Time** | The timestamp representing when the reward template was last modified | [readonly] 
**Subtype** | [**RewardSubtype**](RewardSubtype.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewRewardTemplateResponse

`func NewRewardTemplateResponse(currency string, internalAccountId string, isEnabled bool, creationTime time.Time, id string, lastUpdatedTime time.Time, subtype RewardSubtype, tenant string, ) *RewardTemplateResponse`

NewRewardTemplateResponse instantiates a new RewardTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardTemplateResponseWithDefaults

`func NewRewardTemplateResponseWithDefaults() *RewardTemplateResponse`

NewRewardTemplateResponseWithDefaults instantiates a new RewardTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RewardTemplateResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardTemplateResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardTemplateResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RewardTemplateResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *RewardTemplateResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RewardTemplateResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RewardTemplateResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *RewardTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RewardTemplateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalAccountId

`func (o *RewardTemplateResponse) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *RewardTemplateResponse) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *RewardTemplateResponse) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.


### GetIsEnabled

`func (o *RewardTemplateResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *RewardTemplateResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *RewardTemplateResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMetadata

`func (o *RewardTemplateResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RewardTemplateResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RewardTemplateResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RewardTemplateResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreationTime

`func (o *RewardTemplateResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RewardTemplateResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RewardTemplateResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *RewardTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardTemplateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *RewardTemplateResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *RewardTemplateResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *RewardTemplateResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetSubtype

`func (o *RewardTemplateResponse) GetSubtype() RewardSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *RewardTemplateResponse) GetSubtypeOk() (*RewardSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *RewardTemplateResponse) SetSubtype(v RewardSubtype)`

SetSubtype sets Subtype field to given value.


### GetTenant

`func (o *RewardTemplateResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *RewardTemplateResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *RewardTemplateResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


