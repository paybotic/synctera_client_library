# FeeTemplatePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The default amount of the fee in ISO 4217 minor currency units, e.g. cents. The internal account will be debited this amount. Can be overridden when creating a fee.  | [optional] 
**Currency** | Pointer to **string** | currency for the fee, as a three character ISO 4217 alphabetic currency code. | [optional] 
**Description** | Pointer to **string** | The description of the fee template. | [optional] 
**InternalAccountId** | Pointer to **string** | The ID of default internal_account to use as the destination of the fee transfer. Cannot be a system internal account. | [optional] 
**IsEnabled** | Pointer to **bool** | Whether the fee template is enabled. If false, fees cannot be created from this template.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Subtype** | Pointer to [**FeeSubtypePost**](FeeSubtypePost.md) |  | [optional] 

## Methods

### NewFeeTemplatePatch

`func NewFeeTemplatePatch() *FeeTemplatePatch`

NewFeeTemplatePatch instantiates a new FeeTemplatePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeTemplatePatchWithDefaults

`func NewFeeTemplatePatchWithDefaults() *FeeTemplatePatch`

NewFeeTemplatePatchWithDefaults instantiates a new FeeTemplatePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *FeeTemplatePatch) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeTemplatePatch) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeTemplatePatch) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeeTemplatePatch) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *FeeTemplatePatch) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeeTemplatePatch) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeeTemplatePatch) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FeeTemplatePatch) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *FeeTemplatePatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeeTemplatePatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeeTemplatePatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeeTemplatePatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalAccountId

`func (o *FeeTemplatePatch) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *FeeTemplatePatch) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *FeeTemplatePatch) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.

### HasInternalAccountId

`func (o *FeeTemplatePatch) HasInternalAccountId() bool`

HasInternalAccountId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *FeeTemplatePatch) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FeeTemplatePatch) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FeeTemplatePatch) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FeeTemplatePatch) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *FeeTemplatePatch) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeeTemplatePatch) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeeTemplatePatch) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeeTemplatePatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubtype

`func (o *FeeTemplatePatch) GetSubtype() FeeSubtypePost`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FeeTemplatePatch) GetSubtypeOk() (*FeeSubtypePost, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FeeTemplatePatch) SetSubtype(v FeeSubtypePost)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *FeeTemplatePatch) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


