# CreditCardApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountTemplateId** | **string** | Mainapi account template ID to use for creating the credit card account | 
**Applicant** | [**CreditCardApplicant**](CreditCardApplicant.md) |  | 
**Attributes** | Pointer to **map[string]interface{}** | Raw JSON key-value pairs for flexible data | [optional] 
**Metadata** | Pointer to [**CreditCardApplicationMetadata**](CreditCardApplicationMetadata.md) |  | [optional] 
**Provider** | [**CreditCardProvider**](CreditCardProvider.md) |  | 
**RequestedCreditLimit** | Pointer to **int32** | Requested credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 

## Methods

### NewCreditCardApplication

`func NewCreditCardApplication(accountTemplateId string, applicant CreditCardApplicant, provider CreditCardProvider, tenant string, type_ ApplicationType, ) *CreditCardApplication`

NewCreditCardApplication instantiates a new CreditCardApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardApplicationWithDefaults

`func NewCreditCardApplicationWithDefaults() *CreditCardApplication`

NewCreditCardApplicationWithDefaults instantiates a new CreditCardApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountTemplateId

`func (o *CreditCardApplication) GetAccountTemplateId() string`

GetAccountTemplateId returns the AccountTemplateId field if non-nil, zero value otherwise.

### GetAccountTemplateIdOk

`func (o *CreditCardApplication) GetAccountTemplateIdOk() (*string, bool)`

GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateId

`func (o *CreditCardApplication) SetAccountTemplateId(v string)`

SetAccountTemplateId sets AccountTemplateId field to given value.


### GetApplicant

`func (o *CreditCardApplication) GetApplicant() CreditCardApplicant`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *CreditCardApplication) GetApplicantOk() (*CreditCardApplicant, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *CreditCardApplication) SetApplicant(v CreditCardApplicant)`

SetApplicant sets Applicant field to given value.


### GetAttributes

`func (o *CreditCardApplication) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreditCardApplication) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreditCardApplication) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreditCardApplication) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMetadata

`func (o *CreditCardApplication) GetMetadata() CreditCardApplicationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreditCardApplication) GetMetadataOk() (*CreditCardApplicationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreditCardApplication) SetMetadata(v CreditCardApplicationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreditCardApplication) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProvider

`func (o *CreditCardApplication) GetProvider() CreditCardProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreditCardApplication) GetProviderOk() (*CreditCardProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreditCardApplication) SetProvider(v CreditCardProvider)`

SetProvider sets Provider field to given value.


### GetRequestedCreditLimit

`func (o *CreditCardApplication) GetRequestedCreditLimit() int32`

GetRequestedCreditLimit returns the RequestedCreditLimit field if non-nil, zero value otherwise.

### GetRequestedCreditLimitOk

`func (o *CreditCardApplication) GetRequestedCreditLimitOk() (*int32, bool)`

GetRequestedCreditLimitOk returns a tuple with the RequestedCreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCreditLimit

`func (o *CreditCardApplication) SetRequestedCreditLimit(v int32)`

SetRequestedCreditLimit sets RequestedCreditLimit field to given value.

### HasRequestedCreditLimit

`func (o *CreditCardApplication) HasRequestedCreditLimit() bool`

HasRequestedCreditLimit returns a boolean if a field has been set.

### GetTenant

`func (o *CreditCardApplication) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreditCardApplication) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreditCardApplication) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetType

`func (o *CreditCardApplication) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditCardApplication) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditCardApplication) SetType(v ApplicationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


