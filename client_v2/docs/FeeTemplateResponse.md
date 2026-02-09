# FeeTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The default amount of the fee in ISO 4217 minor currency units, e.g. cents. The internal account will be debited this amount. Can be overridden when creating a fee.  | [optional] 
**Currency** | **string** | currency for the fee, as a three character ISO 4217 alphabetic currency code. | 
**Description** | Pointer to **string** | The description of the fee template. | [optional] 
**InternalAccountId** | **string** | The ID of default internal_account to use as the destination of the fee transfer. Cannot be a system internal account. | 
**IsEnabled** | **bool** | Whether the fee template is enabled. If false, fees cannot be created from this template.  | 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**CreationTime** | **time.Time** | The timestamp representing when the fee template was created | [readonly] 
**Id** | **string** | The ID of the fee template. | 
**LastUpdatedTime** | **time.Time** | The timestamp representing when the fee template was last modified | [readonly] 
**Subtype** | [**FeeSubtype**](FeeSubtype.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewFeeTemplateResponse

`func NewFeeTemplateResponse(currency string, internalAccountId string, isEnabled bool, creationTime time.Time, id string, lastUpdatedTime time.Time, subtype FeeSubtype, tenant string, ) *FeeTemplateResponse`

NewFeeTemplateResponse instantiates a new FeeTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeTemplateResponseWithDefaults

`func NewFeeTemplateResponseWithDefaults() *FeeTemplateResponse`

NewFeeTemplateResponseWithDefaults instantiates a new FeeTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *FeeTemplateResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeTemplateResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeTemplateResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeeTemplateResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *FeeTemplateResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeeTemplateResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeeTemplateResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *FeeTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeeTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeeTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeeTemplateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalAccountId

`func (o *FeeTemplateResponse) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *FeeTemplateResponse) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *FeeTemplateResponse) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.


### GetIsEnabled

`func (o *FeeTemplateResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FeeTemplateResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FeeTemplateResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMetadata

`func (o *FeeTemplateResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeeTemplateResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeeTemplateResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeeTemplateResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreationTime

`func (o *FeeTemplateResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *FeeTemplateResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *FeeTemplateResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *FeeTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeeTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeeTemplateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *FeeTemplateResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *FeeTemplateResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *FeeTemplateResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetSubtype

`func (o *FeeTemplateResponse) GetSubtype() FeeSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FeeTemplateResponse) GetSubtypeOk() (*FeeSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FeeTemplateResponse) SetSubtype(v FeeSubtype)`

SetSubtype sets Subtype field to given value.


### GetTenant

`func (o *FeeTemplateResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *FeeTemplateResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *FeeTemplateResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


