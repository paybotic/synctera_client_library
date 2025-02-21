# EddBusinessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | ISO-3166-1 Alpha-2 country code | [optional] 
**EstimatedRevenue** | Pointer to [**EstimatedRevenue**](EstimatedRevenue.md) |  | [optional] 
**IndustryType** | Pointer to [**IndustryType**](IndustryType.md) |  | [optional] 
**NegativeNewsFindings** | Pointer to **int32** | The number of negative news findings. | [optional] 
**RecurringWireUsage** | Pointer to **bool** | True if the customer is expected to send or receive wire transfers at a regular frequency. | [optional] 
**SpecificInvolvement** | Pointer to [**SpecificInvolvement**](SpecificInvolvement.md) |  | [optional] 
**TransactionVolume** | Pointer to [**[]TransactionVolume**](TransactionVolume.md) | Array of transaction volumes. | [optional] 
**AdditionalQuestions** | Pointer to [**[]Question**](Question.md) | Additional questions regarding the related resource | [optional] 
**CaseId** | Pointer to **int32** | The ID of the case related to this EDD record | [optional] 
**Reason** | **string** | The reason for this EDD record to be requested | 
**RelatedResourceId** | **string** | related resource UUID | 
**RelatedResourceType** | [**RelatedResourceType1**](RelatedResourceType1.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**CreationTime** | **time.Time** |  | [readonly] 
**DeletionTime** | **NullableTime** |  | [readonly] 
**Id** | **string** | EDD record unique identifier | [readonly] 

## Methods

### NewEddBusinessResponse

`func NewEddBusinessResponse(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, creationTime time.Time, deletionTime NullableTime, id string, ) *EddBusinessResponse`

NewEddBusinessResponse instantiates a new EddBusinessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEddBusinessResponseWithDefaults

`func NewEddBusinessResponseWithDefaults() *EddBusinessResponse`

NewEddBusinessResponseWithDefaults instantiates a new EddBusinessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *EddBusinessResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *EddBusinessResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *EddBusinessResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *EddBusinessResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEstimatedRevenue

`func (o *EddBusinessResponse) GetEstimatedRevenue() EstimatedRevenue`

GetEstimatedRevenue returns the EstimatedRevenue field if non-nil, zero value otherwise.

### GetEstimatedRevenueOk

`func (o *EddBusinessResponse) GetEstimatedRevenueOk() (*EstimatedRevenue, bool)`

GetEstimatedRevenueOk returns a tuple with the EstimatedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRevenue

`func (o *EddBusinessResponse) SetEstimatedRevenue(v EstimatedRevenue)`

SetEstimatedRevenue sets EstimatedRevenue field to given value.

### HasEstimatedRevenue

`func (o *EddBusinessResponse) HasEstimatedRevenue() bool`

HasEstimatedRevenue returns a boolean if a field has been set.

### GetIndustryType

`func (o *EddBusinessResponse) GetIndustryType() IndustryType`

GetIndustryType returns the IndustryType field if non-nil, zero value otherwise.

### GetIndustryTypeOk

`func (o *EddBusinessResponse) GetIndustryTypeOk() (*IndustryType, bool)`

GetIndustryTypeOk returns a tuple with the IndustryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryType

`func (o *EddBusinessResponse) SetIndustryType(v IndustryType)`

SetIndustryType sets IndustryType field to given value.

### HasIndustryType

`func (o *EddBusinessResponse) HasIndustryType() bool`

HasIndustryType returns a boolean if a field has been set.

### GetNegativeNewsFindings

`func (o *EddBusinessResponse) GetNegativeNewsFindings() int32`

GetNegativeNewsFindings returns the NegativeNewsFindings field if non-nil, zero value otherwise.

### GetNegativeNewsFindingsOk

`func (o *EddBusinessResponse) GetNegativeNewsFindingsOk() (*int32, bool)`

GetNegativeNewsFindingsOk returns a tuple with the NegativeNewsFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeNewsFindings

`func (o *EddBusinessResponse) SetNegativeNewsFindings(v int32)`

SetNegativeNewsFindings sets NegativeNewsFindings field to given value.

### HasNegativeNewsFindings

`func (o *EddBusinessResponse) HasNegativeNewsFindings() bool`

HasNegativeNewsFindings returns a boolean if a field has been set.

### GetRecurringWireUsage

`func (o *EddBusinessResponse) GetRecurringWireUsage() bool`

GetRecurringWireUsage returns the RecurringWireUsage field if non-nil, zero value otherwise.

### GetRecurringWireUsageOk

`func (o *EddBusinessResponse) GetRecurringWireUsageOk() (*bool, bool)`

GetRecurringWireUsageOk returns a tuple with the RecurringWireUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringWireUsage

`func (o *EddBusinessResponse) SetRecurringWireUsage(v bool)`

SetRecurringWireUsage sets RecurringWireUsage field to given value.

### HasRecurringWireUsage

`func (o *EddBusinessResponse) HasRecurringWireUsage() bool`

HasRecurringWireUsage returns a boolean if a field has been set.

### GetSpecificInvolvement

`func (o *EddBusinessResponse) GetSpecificInvolvement() SpecificInvolvement`

GetSpecificInvolvement returns the SpecificInvolvement field if non-nil, zero value otherwise.

### GetSpecificInvolvementOk

`func (o *EddBusinessResponse) GetSpecificInvolvementOk() (*SpecificInvolvement, bool)`

GetSpecificInvolvementOk returns a tuple with the SpecificInvolvement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificInvolvement

`func (o *EddBusinessResponse) SetSpecificInvolvement(v SpecificInvolvement)`

SetSpecificInvolvement sets SpecificInvolvement field to given value.

### HasSpecificInvolvement

`func (o *EddBusinessResponse) HasSpecificInvolvement() bool`

HasSpecificInvolvement returns a boolean if a field has been set.

### GetTransactionVolume

`func (o *EddBusinessResponse) GetTransactionVolume() []TransactionVolume`

GetTransactionVolume returns the TransactionVolume field if non-nil, zero value otherwise.

### GetTransactionVolumeOk

`func (o *EddBusinessResponse) GetTransactionVolumeOk() (*[]TransactionVolume, bool)`

GetTransactionVolumeOk returns a tuple with the TransactionVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionVolume

`func (o *EddBusinessResponse) SetTransactionVolume(v []TransactionVolume)`

SetTransactionVolume sets TransactionVolume field to given value.

### HasTransactionVolume

`func (o *EddBusinessResponse) HasTransactionVolume() bool`

HasTransactionVolume returns a boolean if a field has been set.

### GetAdditionalQuestions

`func (o *EddBusinessResponse) GetAdditionalQuestions() []Question`

GetAdditionalQuestions returns the AdditionalQuestions field if non-nil, zero value otherwise.

### GetAdditionalQuestionsOk

`func (o *EddBusinessResponse) GetAdditionalQuestionsOk() (*[]Question, bool)`

GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuestions

`func (o *EddBusinessResponse) SetAdditionalQuestions(v []Question)`

SetAdditionalQuestions sets AdditionalQuestions field to given value.

### HasAdditionalQuestions

`func (o *EddBusinessResponse) HasAdditionalQuestions() bool`

HasAdditionalQuestions returns a boolean if a field has been set.

### GetCaseId

`func (o *EddBusinessResponse) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *EddBusinessResponse) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *EddBusinessResponse) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *EddBusinessResponse) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReason

`func (o *EddBusinessResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EddBusinessResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EddBusinessResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRelatedResourceId

`func (o *EddBusinessResponse) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *EddBusinessResponse) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *EddBusinessResponse) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *EddBusinessResponse) GetRelatedResourceType() RelatedResourceType1`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *EddBusinessResponse) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *EddBusinessResponse) SetRelatedResourceType(v RelatedResourceType1)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetTenant

`func (o *EddBusinessResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *EddBusinessResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *EddBusinessResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *EddBusinessResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetCreationTime

`func (o *EddBusinessResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *EddBusinessResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *EddBusinessResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetDeletionTime

`func (o *EddBusinessResponse) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *EddBusinessResponse) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *EddBusinessResponse) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.


### SetDeletionTimeNil

`func (o *EddBusinessResponse) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *EddBusinessResponse) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetId

`func (o *EddBusinessResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EddBusinessResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EddBusinessResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


