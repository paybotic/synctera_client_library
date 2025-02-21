# EddTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceOfFunds** | Pointer to **string** | The source of funds for the transaction. | [optional] 
**TransactionPurpose** | Pointer to **string** | The purpose of the transaction. | [optional] 
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

### NewEddTransactionResponse

`func NewEddTransactionResponse(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, creationTime time.Time, deletionTime NullableTime, id string, ) *EddTransactionResponse`

NewEddTransactionResponse instantiates a new EddTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEddTransactionResponseWithDefaults

`func NewEddTransactionResponseWithDefaults() *EddTransactionResponse`

NewEddTransactionResponseWithDefaults instantiates a new EddTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceOfFunds

`func (o *EddTransactionResponse) GetSourceOfFunds() string`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *EddTransactionResponse) GetSourceOfFundsOk() (*string, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *EddTransactionResponse) SetSourceOfFunds(v string)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *EddTransactionResponse) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetTransactionPurpose

`func (o *EddTransactionResponse) GetTransactionPurpose() string`

GetTransactionPurpose returns the TransactionPurpose field if non-nil, zero value otherwise.

### GetTransactionPurposeOk

`func (o *EddTransactionResponse) GetTransactionPurposeOk() (*string, bool)`

GetTransactionPurposeOk returns a tuple with the TransactionPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPurpose

`func (o *EddTransactionResponse) SetTransactionPurpose(v string)`

SetTransactionPurpose sets TransactionPurpose field to given value.

### HasTransactionPurpose

`func (o *EddTransactionResponse) HasTransactionPurpose() bool`

HasTransactionPurpose returns a boolean if a field has been set.

### GetAdditionalQuestions

`func (o *EddTransactionResponse) GetAdditionalQuestions() []Question`

GetAdditionalQuestions returns the AdditionalQuestions field if non-nil, zero value otherwise.

### GetAdditionalQuestionsOk

`func (o *EddTransactionResponse) GetAdditionalQuestionsOk() (*[]Question, bool)`

GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuestions

`func (o *EddTransactionResponse) SetAdditionalQuestions(v []Question)`

SetAdditionalQuestions sets AdditionalQuestions field to given value.

### HasAdditionalQuestions

`func (o *EddTransactionResponse) HasAdditionalQuestions() bool`

HasAdditionalQuestions returns a boolean if a field has been set.

### GetCaseId

`func (o *EddTransactionResponse) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *EddTransactionResponse) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *EddTransactionResponse) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *EddTransactionResponse) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReason

`func (o *EddTransactionResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EddTransactionResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EddTransactionResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRelatedResourceId

`func (o *EddTransactionResponse) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *EddTransactionResponse) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *EddTransactionResponse) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *EddTransactionResponse) GetRelatedResourceType() RelatedResourceType1`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *EddTransactionResponse) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *EddTransactionResponse) SetRelatedResourceType(v RelatedResourceType1)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetTenant

`func (o *EddTransactionResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *EddTransactionResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *EddTransactionResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *EddTransactionResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetCreationTime

`func (o *EddTransactionResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *EddTransactionResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *EddTransactionResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetDeletionTime

`func (o *EddTransactionResponse) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *EddTransactionResponse) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *EddTransactionResponse) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.


### SetDeletionTimeNil

`func (o *EddTransactionResponse) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *EddTransactionResponse) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetId

`func (o *EddTransactionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EddTransactionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EddTransactionResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


