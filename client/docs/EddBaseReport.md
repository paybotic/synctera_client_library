# EddBaseReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalQuestions** | Pointer to [**[]Question**](Question.md) | Additional questions regarding the related resource | [optional] 
**CaseId** | Pointer to **int32** | The ID of the case related to this EDD record | [optional] 
**Reason** | **string** | The reason for this EDD record to be requested | 
**RelatedResourceId** | **string** | related resource UUID | 
**RelatedResourceType** | [**RelatedResourceType1**](RelatedResourceType1.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 

## Methods

### NewEddBaseReport

`func NewEddBaseReport(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, ) *EddBaseReport`

NewEddBaseReport instantiates a new EddBaseReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEddBaseReportWithDefaults

`func NewEddBaseReportWithDefaults() *EddBaseReport`

NewEddBaseReportWithDefaults instantiates a new EddBaseReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalQuestions

`func (o *EddBaseReport) GetAdditionalQuestions() []Question`

GetAdditionalQuestions returns the AdditionalQuestions field if non-nil, zero value otherwise.

### GetAdditionalQuestionsOk

`func (o *EddBaseReport) GetAdditionalQuestionsOk() (*[]Question, bool)`

GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuestions

`func (o *EddBaseReport) SetAdditionalQuestions(v []Question)`

SetAdditionalQuestions sets AdditionalQuestions field to given value.

### HasAdditionalQuestions

`func (o *EddBaseReport) HasAdditionalQuestions() bool`

HasAdditionalQuestions returns a boolean if a field has been set.

### GetCaseId

`func (o *EddBaseReport) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *EddBaseReport) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *EddBaseReport) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *EddBaseReport) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReason

`func (o *EddBaseReport) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EddBaseReport) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EddBaseReport) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRelatedResourceId

`func (o *EddBaseReport) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *EddBaseReport) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *EddBaseReport) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *EddBaseReport) GetRelatedResourceType() RelatedResourceType1`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *EddBaseReport) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *EddBaseReport) SetRelatedResourceType(v RelatedResourceType1)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetTenant

`func (o *EddBaseReport) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *EddBaseReport) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *EddBaseReport) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *EddBaseReport) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


