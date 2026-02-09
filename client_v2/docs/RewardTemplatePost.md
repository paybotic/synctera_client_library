# RewardTemplatePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The default amount of the reward in ISO 4217 minor currency units, e.g. cents. The internal account will be debited this amount. Can be overridden when creating a reward.  | [optional] 
**Currency** | **string** | currency for the reward, as a three character ISO 4217 alphabetic currency code. | 
**Description** | **string** | The description of the reward template. | 
**InternalAccountId** | **string** | The ID of default internal_account to use as the source of the reward transfer. Cannot be a system internal account. | 
**IsEnabled** | Pointer to **bool** | Whether the reward template is enabled. If false, rewards cannot be created from this template.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Subtype** | [**RewardSubtypePost**](RewardSubtypePost.md) |  | 

## Methods

### NewRewardTemplatePost

`func NewRewardTemplatePost(currency string, description string, internalAccountId string, subtype RewardSubtypePost, ) *RewardTemplatePost`

NewRewardTemplatePost instantiates a new RewardTemplatePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardTemplatePostWithDefaults

`func NewRewardTemplatePostWithDefaults() *RewardTemplatePost`

NewRewardTemplatePostWithDefaults instantiates a new RewardTemplatePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RewardTemplatePost) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardTemplatePost) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardTemplatePost) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RewardTemplatePost) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *RewardTemplatePost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RewardTemplatePost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RewardTemplatePost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *RewardTemplatePost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardTemplatePost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardTemplatePost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInternalAccountId

`func (o *RewardTemplatePost) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *RewardTemplatePost) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *RewardTemplatePost) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.


### GetIsEnabled

`func (o *RewardTemplatePost) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *RewardTemplatePost) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *RewardTemplatePost) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *RewardTemplatePost) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *RewardTemplatePost) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RewardTemplatePost) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RewardTemplatePost) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RewardTemplatePost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubtype

`func (o *RewardTemplatePost) GetSubtype() RewardSubtypePost`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *RewardTemplatePost) GetSubtypeOk() (*RewardSubtypePost, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *RewardTemplatePost) SetSubtype(v RewardSubtypePost)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


