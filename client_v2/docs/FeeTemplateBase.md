# FeeTemplateBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The default amount of the fee in ISO 4217 minor currency units, e.g. cents. The internal account will be debited this amount. Can be overridden when creating a fee.  | [optional] 
**Currency** | Pointer to **string** | currency for the fee, as a three character ISO 4217 alphabetic currency code. | [optional] 
**Description** | Pointer to **string** | The description of the fee template. | [optional] 
**InternalAccountId** | Pointer to **string** | The ID of default internal_account to use as the destination of the fee transfer. Cannot be a system internal account. | [optional] 
**IsEnabled** | Pointer to **bool** | Whether the fee template is enabled. If false, fees cannot be created from this template.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 

## Methods

### NewFeeTemplateBase

`func NewFeeTemplateBase() *FeeTemplateBase`

NewFeeTemplateBase instantiates a new FeeTemplateBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeTemplateBaseWithDefaults

`func NewFeeTemplateBaseWithDefaults() *FeeTemplateBase`

NewFeeTemplateBaseWithDefaults instantiates a new FeeTemplateBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *FeeTemplateBase) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeTemplateBase) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeTemplateBase) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeeTemplateBase) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *FeeTemplateBase) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeeTemplateBase) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeeTemplateBase) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FeeTemplateBase) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *FeeTemplateBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeeTemplateBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeeTemplateBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeeTemplateBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalAccountId

`func (o *FeeTemplateBase) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *FeeTemplateBase) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *FeeTemplateBase) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.

### HasInternalAccountId

`func (o *FeeTemplateBase) HasInternalAccountId() bool`

HasInternalAccountId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *FeeTemplateBase) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FeeTemplateBase) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FeeTemplateBase) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FeeTemplateBase) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *FeeTemplateBase) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeeTemplateBase) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeeTemplateBase) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeeTemplateBase) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


