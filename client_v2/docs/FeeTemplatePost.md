# FeeTemplatePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The default amount of the fee in ISO 4217 minor currency units, e.g. cents. The internal account will be debited this amount. Can be overridden when creating a fee.  | [optional] 
**Currency** | **string** | currency for the fee, as a three character ISO 4217 alphabetic currency code. | 
**Description** | **string** | The description of the fee template. | 
**InternalAccountId** | **string** | The ID of default internal_account to use as the destination of the fee transfer. Cannot be a system internal account. | 
**IsEnabled** | Pointer to **bool** | Whether the fee template is enabled. If false, fees cannot be created from this template.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Subtype** | [**FeeSubtypePost**](FeeSubtypePost.md) |  | 

## Methods

### NewFeeTemplatePost

`func NewFeeTemplatePost(currency string, description string, internalAccountId string, subtype FeeSubtypePost, ) *FeeTemplatePost`

NewFeeTemplatePost instantiates a new FeeTemplatePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeTemplatePostWithDefaults

`func NewFeeTemplatePostWithDefaults() *FeeTemplatePost`

NewFeeTemplatePostWithDefaults instantiates a new FeeTemplatePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *FeeTemplatePost) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeTemplatePost) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeTemplatePost) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeeTemplatePost) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *FeeTemplatePost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeeTemplatePost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeeTemplatePost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *FeeTemplatePost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeeTemplatePost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeeTemplatePost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInternalAccountId

`func (o *FeeTemplatePost) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *FeeTemplatePost) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *FeeTemplatePost) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.


### GetIsEnabled

`func (o *FeeTemplatePost) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FeeTemplatePost) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FeeTemplatePost) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FeeTemplatePost) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *FeeTemplatePost) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeeTemplatePost) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeeTemplatePost) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeeTemplatePost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubtype

`func (o *FeeTemplatePost) GetSubtype() FeeSubtypePost`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FeeTemplatePost) GetSubtypeOk() (*FeeSubtypePost, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FeeTemplatePost) SetSubtype(v FeeSubtypePost)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


