# EddBusiness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalQuestions** | Pointer to [**[]Question**](Question.md) | Additional questions regarding the related resource | [optional] 
**CaseId** | Pointer to **int32** | The ID of the case related to this EDD record | [optional] 
**Reason** | **string** | The reason for this EDD record to be requested | 
**RelatedResourceId** | **string** | related resource UUID | 
**RelatedResourceType** | [**RelatedResourceType1**](RelatedResourceType1.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**Country** | Pointer to **string** | ISO-3166-1 Alpha-2 country code | [optional] 
**EstimatedRevenue** | Pointer to [**EstimatedRevenue**](EstimatedRevenue.md) |  | [optional] 
**IndustryType** | Pointer to [**IndustryType**](IndustryType.md) |  | [optional] 
**NegativeNewsFindings** | Pointer to **int32** | The number of negative news findings. | [optional] 
**RecurringWireUsage** | Pointer to **bool** | True if the customer is expected to send or receive wire transfers at a regular frequency. | [optional] 
**SpecificInvolvement** | Pointer to [**SpecificInvolvement**](SpecificInvolvement.md) |  | [optional] 
**TransactionVolume** | Pointer to [**[]TransactionVolume**](TransactionVolume.md) | Array of transaction volumes. | [optional] 

## Methods

### NewEddBusiness

`func NewEddBusiness(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, ) *EddBusiness`

NewEddBusiness instantiates a new EddBusiness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEddBusinessWithDefaults

`func NewEddBusinessWithDefaults() *EddBusiness`

NewEddBusinessWithDefaults instantiates a new EddBusiness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalQuestions

`func (o *EddBusiness) GetAdditionalQuestions() []Question`

GetAdditionalQuestions returns the AdditionalQuestions field if non-nil, zero value otherwise.

### GetAdditionalQuestionsOk

`func (o *EddBusiness) GetAdditionalQuestionsOk() (*[]Question, bool)`

GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuestions

`func (o *EddBusiness) SetAdditionalQuestions(v []Question)`

SetAdditionalQuestions sets AdditionalQuestions field to given value.

### HasAdditionalQuestions

`func (o *EddBusiness) HasAdditionalQuestions() bool`

HasAdditionalQuestions returns a boolean if a field has been set.

### GetCaseId

`func (o *EddBusiness) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *EddBusiness) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *EddBusiness) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *EddBusiness) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReason

`func (o *EddBusiness) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EddBusiness) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EddBusiness) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRelatedResourceId

`func (o *EddBusiness) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *EddBusiness) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *EddBusiness) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *EddBusiness) GetRelatedResourceType() RelatedResourceType1`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *EddBusiness) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *EddBusiness) SetRelatedResourceType(v RelatedResourceType1)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetTenant

`func (o *EddBusiness) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *EddBusiness) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *EddBusiness) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *EddBusiness) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetCountry

`func (o *EddBusiness) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *EddBusiness) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *EddBusiness) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *EddBusiness) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEstimatedRevenue

`func (o *EddBusiness) GetEstimatedRevenue() EstimatedRevenue`

GetEstimatedRevenue returns the EstimatedRevenue field if non-nil, zero value otherwise.

### GetEstimatedRevenueOk

`func (o *EddBusiness) GetEstimatedRevenueOk() (*EstimatedRevenue, bool)`

GetEstimatedRevenueOk returns a tuple with the EstimatedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRevenue

`func (o *EddBusiness) SetEstimatedRevenue(v EstimatedRevenue)`

SetEstimatedRevenue sets EstimatedRevenue field to given value.

### HasEstimatedRevenue

`func (o *EddBusiness) HasEstimatedRevenue() bool`

HasEstimatedRevenue returns a boolean if a field has been set.

### GetIndustryType

`func (o *EddBusiness) GetIndustryType() IndustryType`

GetIndustryType returns the IndustryType field if non-nil, zero value otherwise.

### GetIndustryTypeOk

`func (o *EddBusiness) GetIndustryTypeOk() (*IndustryType, bool)`

GetIndustryTypeOk returns a tuple with the IndustryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryType

`func (o *EddBusiness) SetIndustryType(v IndustryType)`

SetIndustryType sets IndustryType field to given value.

### HasIndustryType

`func (o *EddBusiness) HasIndustryType() bool`

HasIndustryType returns a boolean if a field has been set.

### GetNegativeNewsFindings

`func (o *EddBusiness) GetNegativeNewsFindings() int32`

GetNegativeNewsFindings returns the NegativeNewsFindings field if non-nil, zero value otherwise.

### GetNegativeNewsFindingsOk

`func (o *EddBusiness) GetNegativeNewsFindingsOk() (*int32, bool)`

GetNegativeNewsFindingsOk returns a tuple with the NegativeNewsFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeNewsFindings

`func (o *EddBusiness) SetNegativeNewsFindings(v int32)`

SetNegativeNewsFindings sets NegativeNewsFindings field to given value.

### HasNegativeNewsFindings

`func (o *EddBusiness) HasNegativeNewsFindings() bool`

HasNegativeNewsFindings returns a boolean if a field has been set.

### GetRecurringWireUsage

`func (o *EddBusiness) GetRecurringWireUsage() bool`

GetRecurringWireUsage returns the RecurringWireUsage field if non-nil, zero value otherwise.

### GetRecurringWireUsageOk

`func (o *EddBusiness) GetRecurringWireUsageOk() (*bool, bool)`

GetRecurringWireUsageOk returns a tuple with the RecurringWireUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringWireUsage

`func (o *EddBusiness) SetRecurringWireUsage(v bool)`

SetRecurringWireUsage sets RecurringWireUsage field to given value.

### HasRecurringWireUsage

`func (o *EddBusiness) HasRecurringWireUsage() bool`

HasRecurringWireUsage returns a boolean if a field has been set.

### GetSpecificInvolvement

`func (o *EddBusiness) GetSpecificInvolvement() SpecificInvolvement`

GetSpecificInvolvement returns the SpecificInvolvement field if non-nil, zero value otherwise.

### GetSpecificInvolvementOk

`func (o *EddBusiness) GetSpecificInvolvementOk() (*SpecificInvolvement, bool)`

GetSpecificInvolvementOk returns a tuple with the SpecificInvolvement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificInvolvement

`func (o *EddBusiness) SetSpecificInvolvement(v SpecificInvolvement)`

SetSpecificInvolvement sets SpecificInvolvement field to given value.

### HasSpecificInvolvement

`func (o *EddBusiness) HasSpecificInvolvement() bool`

HasSpecificInvolvement returns a boolean if a field has been set.

### GetTransactionVolume

`func (o *EddBusiness) GetTransactionVolume() []TransactionVolume`

GetTransactionVolume returns the TransactionVolume field if non-nil, zero value otherwise.

### GetTransactionVolumeOk

`func (o *EddBusiness) GetTransactionVolumeOk() (*[]TransactionVolume, bool)`

GetTransactionVolumeOk returns a tuple with the TransactionVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionVolume

`func (o *EddBusiness) SetTransactionVolume(v []TransactionVolume)`

SetTransactionVolume sets TransactionVolume field to given value.

### HasTransactionVolume

`func (o *EddBusiness) HasTransactionVolume() bool`

HasTransactionVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


