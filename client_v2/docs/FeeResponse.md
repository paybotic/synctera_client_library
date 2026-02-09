# FeeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the business or customer account being charged the fee. | 
**Amount** | **int32** | The amount of the fee in ISO 4217 minor currency units, e.g. cents. The internal account referenced by the fee template will be debited this amount. Defaults to the value in the fee template.  | 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Note** | Pointer to **string** | An optional note for this instance of the fee. | [optional] 
**ReferenceId** | Pointer to **string** | An optional reference ID that uniquely identifies this fee instance. This ID will be included in the posted fee transaction. | [optional] 
**TemplateId** | **string** | The ID of the fee template to use to create the fee. Values from the fee template will be used as defaults for the fee. Note that the fee template may have been updated since the fee was created and that such subsequent updates to the fee template do not affect existing fees.  | 
**CreationTime** | **time.Time** | The timestamp representing when the fee was created | [readonly] 
**Currency** | **string** | currency of the fee, as a three character ISO 4217 alphabetic currency code. | 
**Description** | **string** | The description of the fee template. | 
**Id** | **string** | The ID of the fee. | 
**InternalAccountId** | **string** | The ID of internal_account that is the destination of the fee transfer. | 
**LastUpdatedTime** | **time.Time** | The timestamp representing when the fee was last updated | [readonly] 
**Subtype** | [**FeeSubtype**](FeeSubtype.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**TransactionId** | **string** | The ID the resulting transaction resource. | 

## Methods

### NewFeeResponse

`func NewFeeResponse(accountId string, amount int32, templateId string, creationTime time.Time, currency string, description string, id string, internalAccountId string, lastUpdatedTime time.Time, subtype FeeSubtype, tenant string, transactionId string, ) *FeeResponse`

NewFeeResponse instantiates a new FeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeResponseWithDefaults

`func NewFeeResponseWithDefaults() *FeeResponse`

NewFeeResponseWithDefaults instantiates a new FeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *FeeResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *FeeResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *FeeResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *FeeResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetMetadata

`func (o *FeeResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeeResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeeResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeeResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNote

`func (o *FeeResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *FeeResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *FeeResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *FeeResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReferenceId

`func (o *FeeResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *FeeResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *FeeResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *FeeResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetTemplateId

`func (o *FeeResponse) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *FeeResponse) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *FeeResponse) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetCreationTime

`func (o *FeeResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *FeeResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *FeeResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *FeeResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeeResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeeResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *FeeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *FeeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInternalAccountId

`func (o *FeeResponse) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *FeeResponse) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *FeeResponse) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.


### GetLastUpdatedTime

`func (o *FeeResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *FeeResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *FeeResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetSubtype

`func (o *FeeResponse) GetSubtype() FeeSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FeeResponse) GetSubtypeOk() (*FeeSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FeeResponse) SetSubtype(v FeeSubtype)`

SetSubtype sets Subtype field to given value.


### GetTenant

`func (o *FeeResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *FeeResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *FeeResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionId

`func (o *FeeResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *FeeResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *FeeResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


