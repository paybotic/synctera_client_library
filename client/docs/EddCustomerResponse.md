# EddCustomerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitizenshipCountries** | Pointer to **[]string** | List of countries where the related customer holds citizenship. | [optional] 
**EmploymentType** | Pointer to **string** | The type of employment. | [optional] 
**Income** | Pointer to [**Income**](Income.md) |  | [optional] 
**NegativeNewsFindings** | Pointer to **int32** | The number of negative news findings. | [optional] 
**Occupation** | Pointer to **string** | The occupation of the related resource. | [optional] 
**OccupationIndustry** | Pointer to [**IndustryType**](IndustryType.md) |  | [optional] 
**RecurringDirectDeposit** | Pointer to **bool** | True if the customer is expected to use direct deposit at a regular frequency. | [optional] 
**ResidenceType** | Pointer to **string** | The type of residence. | [optional] 
**ResidentialExpense** | Pointer to [**ResidentialExpense**](ResidentialExpense.md) |  | [optional] 
**SourceOfWealth** | Pointer to [**[]SourceOfWealth**](SourceOfWealth.md) | The sources of wealth for the customer. | [optional] 
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

### NewEddCustomerResponse

`func NewEddCustomerResponse(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, creationTime time.Time, deletionTime NullableTime, id string, ) *EddCustomerResponse`

NewEddCustomerResponse instantiates a new EddCustomerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEddCustomerResponseWithDefaults

`func NewEddCustomerResponseWithDefaults() *EddCustomerResponse`

NewEddCustomerResponseWithDefaults instantiates a new EddCustomerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitizenshipCountries

`func (o *EddCustomerResponse) GetCitizenshipCountries() []string`

GetCitizenshipCountries returns the CitizenshipCountries field if non-nil, zero value otherwise.

### GetCitizenshipCountriesOk

`func (o *EddCustomerResponse) GetCitizenshipCountriesOk() (*[]string, bool)`

GetCitizenshipCountriesOk returns a tuple with the CitizenshipCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenshipCountries

`func (o *EddCustomerResponse) SetCitizenshipCountries(v []string)`

SetCitizenshipCountries sets CitizenshipCountries field to given value.

### HasCitizenshipCountries

`func (o *EddCustomerResponse) HasCitizenshipCountries() bool`

HasCitizenshipCountries returns a boolean if a field has been set.

### GetEmploymentType

`func (o *EddCustomerResponse) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *EddCustomerResponse) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *EddCustomerResponse) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.

### HasEmploymentType

`func (o *EddCustomerResponse) HasEmploymentType() bool`

HasEmploymentType returns a boolean if a field has been set.

### GetIncome

`func (o *EddCustomerResponse) GetIncome() Income`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *EddCustomerResponse) GetIncomeOk() (*Income, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *EddCustomerResponse) SetIncome(v Income)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *EddCustomerResponse) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetNegativeNewsFindings

`func (o *EddCustomerResponse) GetNegativeNewsFindings() int32`

GetNegativeNewsFindings returns the NegativeNewsFindings field if non-nil, zero value otherwise.

### GetNegativeNewsFindingsOk

`func (o *EddCustomerResponse) GetNegativeNewsFindingsOk() (*int32, bool)`

GetNegativeNewsFindingsOk returns a tuple with the NegativeNewsFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeNewsFindings

`func (o *EddCustomerResponse) SetNegativeNewsFindings(v int32)`

SetNegativeNewsFindings sets NegativeNewsFindings field to given value.

### HasNegativeNewsFindings

`func (o *EddCustomerResponse) HasNegativeNewsFindings() bool`

HasNegativeNewsFindings returns a boolean if a field has been set.

### GetOccupation

`func (o *EddCustomerResponse) GetOccupation() string`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *EddCustomerResponse) GetOccupationOk() (*string, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *EddCustomerResponse) SetOccupation(v string)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *EddCustomerResponse) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### GetOccupationIndustry

`func (o *EddCustomerResponse) GetOccupationIndustry() IndustryType`

GetOccupationIndustry returns the OccupationIndustry field if non-nil, zero value otherwise.

### GetOccupationIndustryOk

`func (o *EddCustomerResponse) GetOccupationIndustryOk() (*IndustryType, bool)`

GetOccupationIndustryOk returns a tuple with the OccupationIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupationIndustry

`func (o *EddCustomerResponse) SetOccupationIndustry(v IndustryType)`

SetOccupationIndustry sets OccupationIndustry field to given value.

### HasOccupationIndustry

`func (o *EddCustomerResponse) HasOccupationIndustry() bool`

HasOccupationIndustry returns a boolean if a field has been set.

### GetRecurringDirectDeposit

`func (o *EddCustomerResponse) GetRecurringDirectDeposit() bool`

GetRecurringDirectDeposit returns the RecurringDirectDeposit field if non-nil, zero value otherwise.

### GetRecurringDirectDepositOk

`func (o *EddCustomerResponse) GetRecurringDirectDepositOk() (*bool, bool)`

GetRecurringDirectDepositOk returns a tuple with the RecurringDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDirectDeposit

`func (o *EddCustomerResponse) SetRecurringDirectDeposit(v bool)`

SetRecurringDirectDeposit sets RecurringDirectDeposit field to given value.

### HasRecurringDirectDeposit

`func (o *EddCustomerResponse) HasRecurringDirectDeposit() bool`

HasRecurringDirectDeposit returns a boolean if a field has been set.

### GetResidenceType

`func (o *EddCustomerResponse) GetResidenceType() string`

GetResidenceType returns the ResidenceType field if non-nil, zero value otherwise.

### GetResidenceTypeOk

`func (o *EddCustomerResponse) GetResidenceTypeOk() (*string, bool)`

GetResidenceTypeOk returns a tuple with the ResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceType

`func (o *EddCustomerResponse) SetResidenceType(v string)`

SetResidenceType sets ResidenceType field to given value.

### HasResidenceType

`func (o *EddCustomerResponse) HasResidenceType() bool`

HasResidenceType returns a boolean if a field has been set.

### GetResidentialExpense

`func (o *EddCustomerResponse) GetResidentialExpense() ResidentialExpense`

GetResidentialExpense returns the ResidentialExpense field if non-nil, zero value otherwise.

### GetResidentialExpenseOk

`func (o *EddCustomerResponse) GetResidentialExpenseOk() (*ResidentialExpense, bool)`

GetResidentialExpenseOk returns a tuple with the ResidentialExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentialExpense

`func (o *EddCustomerResponse) SetResidentialExpense(v ResidentialExpense)`

SetResidentialExpense sets ResidentialExpense field to given value.

### HasResidentialExpense

`func (o *EddCustomerResponse) HasResidentialExpense() bool`

HasResidentialExpense returns a boolean if a field has been set.

### GetSourceOfWealth

`func (o *EddCustomerResponse) GetSourceOfWealth() []SourceOfWealth`

GetSourceOfWealth returns the SourceOfWealth field if non-nil, zero value otherwise.

### GetSourceOfWealthOk

`func (o *EddCustomerResponse) GetSourceOfWealthOk() (*[]SourceOfWealth, bool)`

GetSourceOfWealthOk returns a tuple with the SourceOfWealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfWealth

`func (o *EddCustomerResponse) SetSourceOfWealth(v []SourceOfWealth)`

SetSourceOfWealth sets SourceOfWealth field to given value.

### HasSourceOfWealth

`func (o *EddCustomerResponse) HasSourceOfWealth() bool`

HasSourceOfWealth returns a boolean if a field has been set.

### GetAdditionalQuestions

`func (o *EddCustomerResponse) GetAdditionalQuestions() []Question`

GetAdditionalQuestions returns the AdditionalQuestions field if non-nil, zero value otherwise.

### GetAdditionalQuestionsOk

`func (o *EddCustomerResponse) GetAdditionalQuestionsOk() (*[]Question, bool)`

GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuestions

`func (o *EddCustomerResponse) SetAdditionalQuestions(v []Question)`

SetAdditionalQuestions sets AdditionalQuestions field to given value.

### HasAdditionalQuestions

`func (o *EddCustomerResponse) HasAdditionalQuestions() bool`

HasAdditionalQuestions returns a boolean if a field has been set.

### GetCaseId

`func (o *EddCustomerResponse) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *EddCustomerResponse) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *EddCustomerResponse) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *EddCustomerResponse) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReason

`func (o *EddCustomerResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EddCustomerResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EddCustomerResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRelatedResourceId

`func (o *EddCustomerResponse) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *EddCustomerResponse) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *EddCustomerResponse) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *EddCustomerResponse) GetRelatedResourceType() RelatedResourceType1`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *EddCustomerResponse) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *EddCustomerResponse) SetRelatedResourceType(v RelatedResourceType1)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetTenant

`func (o *EddCustomerResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *EddCustomerResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *EddCustomerResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *EddCustomerResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetCreationTime

`func (o *EddCustomerResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *EddCustomerResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *EddCustomerResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetDeletionTime

`func (o *EddCustomerResponse) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *EddCustomerResponse) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *EddCustomerResponse) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.


### SetDeletionTimeNil

`func (o *EddCustomerResponse) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *EddCustomerResponse) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetId

`func (o *EddCustomerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EddCustomerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EddCustomerResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


