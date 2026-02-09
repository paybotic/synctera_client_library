# RewardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the business or customer account being charged the reward. | 
**Amount** | **int32** | The amount of the reward in ISO 4217 minor currency units, e.g. cents. The internal account referenced by the reward template will be debited this amount. Defaults to the value in the reward template.  | 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Note** | Pointer to **string** | An optional note for this instance of the reward. | [optional] 
**TemplateId** | **string** | The ID of the reward template to use to create the reward. Values from the reward template will be used as defaults for the reward. Note that the reward template may have been updated since the reward was created and that such subsequent updates to the reward template do not affect existing rewards.  | 
**CreationTime** | **time.Time** | The timestamp representing when the reward was created | [readonly] 
**Currency** | **string** | currency of the reward, as a three character ISO 4217 alphabetic currency code. | 
**Description** | **string** | The description of the reward template. | 
**Id** | **string** | The ID of the reward. | 
**InternalAccountId** | **string** | The ID of internal_account that is the source of the reward transfer. | 
**LastUpdatedTime** | **time.Time** | The timestamp representing when the reward was last updated | [readonly] 
**Subtype** | [**RewardSubtype**](RewardSubtype.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**TransactionId** | **string** | The ID the resulting transaction resource. | 

## Methods

### NewRewardResponse

`func NewRewardResponse(accountId string, amount int32, templateId string, creationTime time.Time, currency string, description string, id string, internalAccountId string, lastUpdatedTime time.Time, subtype RewardSubtype, tenant string, transactionId string, ) *RewardResponse`

NewRewardResponse instantiates a new RewardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardResponseWithDefaults

`func NewRewardResponseWithDefaults() *RewardResponse`

NewRewardResponseWithDefaults instantiates a new RewardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RewardResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RewardResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RewardResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *RewardResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetMetadata

`func (o *RewardResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RewardResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RewardResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RewardResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNote

`func (o *RewardResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RewardResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RewardResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *RewardResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTemplateId

`func (o *RewardResponse) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *RewardResponse) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *RewardResponse) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetCreationTime

`func (o *RewardResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RewardResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RewardResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *RewardResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RewardResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RewardResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *RewardResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *RewardResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInternalAccountId

`func (o *RewardResponse) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *RewardResponse) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *RewardResponse) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.


### GetLastUpdatedTime

`func (o *RewardResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *RewardResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *RewardResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetSubtype

`func (o *RewardResponse) GetSubtype() RewardSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *RewardResponse) GetSubtypeOk() (*RewardSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *RewardResponse) SetSubtype(v RewardSubtype)`

SetSubtype sets Subtype field to given value.


### GetTenant

`func (o *RewardResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *RewardResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *RewardResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionId

`func (o *RewardResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *RewardResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *RewardResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


