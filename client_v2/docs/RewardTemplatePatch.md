# RewardTemplatePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The default amount of the reward in ISO 4217 minor currency units, e.g. cents. The internal account will be debited this amount. Can be overridden when creating a reward.  | [optional] 
**Currency** | Pointer to **string** | currency for the reward, as a three character ISO 4217 alphabetic currency code. | [optional] 
**Description** | Pointer to **string** | The description of the reward template. | [optional] 
**InternalAccountId** | Pointer to **string** | The ID of default internal_account to use as the source of the reward transfer. Cannot be a system internal account. | [optional] 
**IsEnabled** | Pointer to **bool** | Whether the reward template is enabled. If false, rewards cannot be created from this template.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Subtype** | Pointer to [**RewardSubtypePost**](RewardSubtypePost.md) |  | [optional] 

## Methods

### NewRewardTemplatePatch

`func NewRewardTemplatePatch() *RewardTemplatePatch`

NewRewardTemplatePatch instantiates a new RewardTemplatePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardTemplatePatchWithDefaults

`func NewRewardTemplatePatchWithDefaults() *RewardTemplatePatch`

NewRewardTemplatePatchWithDefaults instantiates a new RewardTemplatePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RewardTemplatePatch) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardTemplatePatch) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardTemplatePatch) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RewardTemplatePatch) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *RewardTemplatePatch) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RewardTemplatePatch) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RewardTemplatePatch) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RewardTemplatePatch) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *RewardTemplatePatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardTemplatePatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardTemplatePatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RewardTemplatePatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalAccountId

`func (o *RewardTemplatePatch) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *RewardTemplatePatch) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *RewardTemplatePatch) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.

### HasInternalAccountId

`func (o *RewardTemplatePatch) HasInternalAccountId() bool`

HasInternalAccountId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *RewardTemplatePatch) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *RewardTemplatePatch) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *RewardTemplatePatch) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *RewardTemplatePatch) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *RewardTemplatePatch) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RewardTemplatePatch) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RewardTemplatePatch) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RewardTemplatePatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubtype

`func (o *RewardTemplatePatch) GetSubtype() RewardSubtypePost`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *RewardTemplatePatch) GetSubtypeOk() (*RewardSubtypePost, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *RewardTemplatePatch) SetSubtype(v RewardSubtypePost)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *RewardTemplatePatch) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


