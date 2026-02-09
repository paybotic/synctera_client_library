# LocApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountTemplateId** | **string** | Mainapi account template ID to use for creating the LOC account | 
**Applicant** | [**LocApplicant**](LocApplicant.md) |  | 
**Attributes** | Pointer to **map[string]interface{}** | Raw JSON key-value pairs for flexible data | [optional] 
**Metadata** | Pointer to [**LocApplicationMetadata**](LocApplicationMetadata.md) |  | [optional] 
**Provider** | [**LocProvider**](LocProvider.md) |  | 
**RequestedCreditLimit** | Pointer to **int32** | Requested credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 

## Methods

### NewLocApplication

`func NewLocApplication(accountTemplateId string, applicant LocApplicant, provider LocProvider, tenant string, type_ ApplicationType, ) *LocApplication`

NewLocApplication instantiates a new LocApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocApplicationWithDefaults

`func NewLocApplicationWithDefaults() *LocApplication`

NewLocApplicationWithDefaults instantiates a new LocApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountTemplateId

`func (o *LocApplication) GetAccountTemplateId() string`

GetAccountTemplateId returns the AccountTemplateId field if non-nil, zero value otherwise.

### GetAccountTemplateIdOk

`func (o *LocApplication) GetAccountTemplateIdOk() (*string, bool)`

GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateId

`func (o *LocApplication) SetAccountTemplateId(v string)`

SetAccountTemplateId sets AccountTemplateId field to given value.


### GetApplicant

`func (o *LocApplication) GetApplicant() LocApplicant`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *LocApplication) GetApplicantOk() (*LocApplicant, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *LocApplication) SetApplicant(v LocApplicant)`

SetApplicant sets Applicant field to given value.


### GetAttributes

`func (o *LocApplication) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LocApplication) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LocApplication) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LocApplication) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMetadata

`func (o *LocApplication) GetMetadata() LocApplicationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LocApplication) GetMetadataOk() (*LocApplicationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LocApplication) SetMetadata(v LocApplicationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LocApplication) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProvider

`func (o *LocApplication) GetProvider() LocProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LocApplication) GetProviderOk() (*LocProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LocApplication) SetProvider(v LocProvider)`

SetProvider sets Provider field to given value.


### GetRequestedCreditLimit

`func (o *LocApplication) GetRequestedCreditLimit() int32`

GetRequestedCreditLimit returns the RequestedCreditLimit field if non-nil, zero value otherwise.

### GetRequestedCreditLimitOk

`func (o *LocApplication) GetRequestedCreditLimitOk() (*int32, bool)`

GetRequestedCreditLimitOk returns a tuple with the RequestedCreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCreditLimit

`func (o *LocApplication) SetRequestedCreditLimit(v int32)`

SetRequestedCreditLimit sets RequestedCreditLimit field to given value.

### HasRequestedCreditLimit

`func (o *LocApplication) HasRequestedCreditLimit() bool`

HasRequestedCreditLimit returns a boolean if a field has been set.

### GetTenant

`func (o *LocApplication) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *LocApplication) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *LocApplication) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetType

`func (o *LocApplication) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocApplication) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocApplication) SetType(v ApplicationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


