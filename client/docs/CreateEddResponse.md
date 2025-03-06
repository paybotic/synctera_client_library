# CreateEddResponse

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
**Country** | Pointer to **string** | ISO-3166-1 Alpha-2 country code | [optional] 
**EstimatedRevenue** | Pointer to [**EstimatedRevenue**](EstimatedRevenue.md) |  | [optional] 
**IndustryType** | Pointer to [**IndustryType**](IndustryType.md) |  | [optional] 
**RecurringWireUsage** | Pointer to **bool** | True if the customer is expected to send or receive wire transfers at a regular frequency. | [optional] 
**SpecificInvolvement** | Pointer to [**SpecificInvolvement**](SpecificInvolvement.md) |  | [optional] 
**TransactionVolume** | Pointer to [**[]TransactionVolume**](TransactionVolume.md) | Array of transaction volumes. | [optional] 

## Methods

### NewCreateEddResponse

`func NewCreateEddResponse(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, creationTime time.Time, deletionTime NullableTime, id string, ) *CreateEddResponse`

NewCreateEddResponse instantiates a new CreateEddResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEddResponseWithDefaults

`func NewCreateEddResponseWithDefaults() *CreateEddResponse`

NewCreateEddResponseWithDefaults instantiates a new CreateEddResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceOfFunds

`func (o *CreateEddResponse) GetSourceOfFunds() string`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *CreateEddResponse) GetSourceOfFundsOk() (*string, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *CreateEddResponse) SetSourceOfFunds(v string)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *CreateEddResponse) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetTransactionPurpose

`func (o *CreateEddResponse) GetTransactionPurpose() string`

GetTransactionPurpose returns the TransactionPurpose field if non-nil, zero value otherwise.

### GetTransactionPurposeOk

`func (o *CreateEddResponse) GetTransactionPurposeOk() (*string, bool)`

GetTransactionPurposeOk returns a tuple with the TransactionPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPurpose

`func (o *CreateEddResponse) SetTransactionPurpose(v string)`

SetTransactionPurpose sets TransactionPurpose field to given value.

### HasTransactionPurpose

`func (o *CreateEddResponse) HasTransactionPurpose() bool`

HasTransactionPurpose returns a boolean if a field has been set.

### GetAdditionalQuestions

`func (o *CreateEddResponse) GetAdditionalQuestions() []Question`

GetAdditionalQuestions returns the AdditionalQuestions field if non-nil, zero value otherwise.

### GetAdditionalQuestionsOk

`func (o *CreateEddResponse) GetAdditionalQuestionsOk() (*[]Question, bool)`

GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuestions

`func (o *CreateEddResponse) SetAdditionalQuestions(v []Question)`

SetAdditionalQuestions sets AdditionalQuestions field to given value.

### HasAdditionalQuestions

`func (o *CreateEddResponse) HasAdditionalQuestions() bool`

HasAdditionalQuestions returns a boolean if a field has been set.

### GetCaseId

`func (o *CreateEddResponse) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *CreateEddResponse) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *CreateEddResponse) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *CreateEddResponse) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReason

`func (o *CreateEddResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateEddResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateEddResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRelatedResourceId

`func (o *CreateEddResponse) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *CreateEddResponse) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *CreateEddResponse) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *CreateEddResponse) GetRelatedResourceType() RelatedResourceType1`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *CreateEddResponse) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *CreateEddResponse) SetRelatedResourceType(v RelatedResourceType1)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetTenant

`func (o *CreateEddResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreateEddResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreateEddResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CreateEddResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetCreationTime

`func (o *CreateEddResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreateEddResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreateEddResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetDeletionTime

`func (o *CreateEddResponse) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *CreateEddResponse) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *CreateEddResponse) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.


### SetDeletionTimeNil

`func (o *CreateEddResponse) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *CreateEddResponse) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetId

`func (o *CreateEddResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateEddResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateEddResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCitizenshipCountries

`func (o *CreateEddResponse) GetCitizenshipCountries() []string`

GetCitizenshipCountries returns the CitizenshipCountries field if non-nil, zero value otherwise.

### GetCitizenshipCountriesOk

`func (o *CreateEddResponse) GetCitizenshipCountriesOk() (*[]string, bool)`

GetCitizenshipCountriesOk returns a tuple with the CitizenshipCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenshipCountries

`func (o *CreateEddResponse) SetCitizenshipCountries(v []string)`

SetCitizenshipCountries sets CitizenshipCountries field to given value.

### HasCitizenshipCountries

`func (o *CreateEddResponse) HasCitizenshipCountries() bool`

HasCitizenshipCountries returns a boolean if a field has been set.

### GetEmploymentType

`func (o *CreateEddResponse) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *CreateEddResponse) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *CreateEddResponse) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.

### HasEmploymentType

`func (o *CreateEddResponse) HasEmploymentType() bool`

HasEmploymentType returns a boolean if a field has been set.

### GetIncome

`func (o *CreateEddResponse) GetIncome() Income`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *CreateEddResponse) GetIncomeOk() (*Income, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *CreateEddResponse) SetIncome(v Income)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *CreateEddResponse) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetNegativeNewsFindings

`func (o *CreateEddResponse) GetNegativeNewsFindings() int32`

GetNegativeNewsFindings returns the NegativeNewsFindings field if non-nil, zero value otherwise.

### GetNegativeNewsFindingsOk

`func (o *CreateEddResponse) GetNegativeNewsFindingsOk() (*int32, bool)`

GetNegativeNewsFindingsOk returns a tuple with the NegativeNewsFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeNewsFindings

`func (o *CreateEddResponse) SetNegativeNewsFindings(v int32)`

SetNegativeNewsFindings sets NegativeNewsFindings field to given value.

### HasNegativeNewsFindings

`func (o *CreateEddResponse) HasNegativeNewsFindings() bool`

HasNegativeNewsFindings returns a boolean if a field has been set.

### GetOccupation

`func (o *CreateEddResponse) GetOccupation() string`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *CreateEddResponse) GetOccupationOk() (*string, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *CreateEddResponse) SetOccupation(v string)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *CreateEddResponse) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### GetOccupationIndustry

`func (o *CreateEddResponse) GetOccupationIndustry() IndustryType`

GetOccupationIndustry returns the OccupationIndustry field if non-nil, zero value otherwise.

### GetOccupationIndustryOk

`func (o *CreateEddResponse) GetOccupationIndustryOk() (*IndustryType, bool)`

GetOccupationIndustryOk returns a tuple with the OccupationIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupationIndustry

`func (o *CreateEddResponse) SetOccupationIndustry(v IndustryType)`

SetOccupationIndustry sets OccupationIndustry field to given value.

### HasOccupationIndustry

`func (o *CreateEddResponse) HasOccupationIndustry() bool`

HasOccupationIndustry returns a boolean if a field has been set.

### GetRecurringDirectDeposit

`func (o *CreateEddResponse) GetRecurringDirectDeposit() bool`

GetRecurringDirectDeposit returns the RecurringDirectDeposit field if non-nil, zero value otherwise.

### GetRecurringDirectDepositOk

`func (o *CreateEddResponse) GetRecurringDirectDepositOk() (*bool, bool)`

GetRecurringDirectDepositOk returns a tuple with the RecurringDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDirectDeposit

`func (o *CreateEddResponse) SetRecurringDirectDeposit(v bool)`

SetRecurringDirectDeposit sets RecurringDirectDeposit field to given value.

### HasRecurringDirectDeposit

`func (o *CreateEddResponse) HasRecurringDirectDeposit() bool`

HasRecurringDirectDeposit returns a boolean if a field has been set.

### GetResidenceType

`func (o *CreateEddResponse) GetResidenceType() string`

GetResidenceType returns the ResidenceType field if non-nil, zero value otherwise.

### GetResidenceTypeOk

`func (o *CreateEddResponse) GetResidenceTypeOk() (*string, bool)`

GetResidenceTypeOk returns a tuple with the ResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceType

`func (o *CreateEddResponse) SetResidenceType(v string)`

SetResidenceType sets ResidenceType field to given value.

### HasResidenceType

`func (o *CreateEddResponse) HasResidenceType() bool`

HasResidenceType returns a boolean if a field has been set.

### GetResidentialExpense

`func (o *CreateEddResponse) GetResidentialExpense() ResidentialExpense`

GetResidentialExpense returns the ResidentialExpense field if non-nil, zero value otherwise.

### GetResidentialExpenseOk

`func (o *CreateEddResponse) GetResidentialExpenseOk() (*ResidentialExpense, bool)`

GetResidentialExpenseOk returns a tuple with the ResidentialExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentialExpense

`func (o *CreateEddResponse) SetResidentialExpense(v ResidentialExpense)`

SetResidentialExpense sets ResidentialExpense field to given value.

### HasResidentialExpense

`func (o *CreateEddResponse) HasResidentialExpense() bool`

HasResidentialExpense returns a boolean if a field has been set.

### GetSourceOfWealth

`func (o *CreateEddResponse) GetSourceOfWealth() []SourceOfWealth`

GetSourceOfWealth returns the SourceOfWealth field if non-nil, zero value otherwise.

### GetSourceOfWealthOk

`func (o *CreateEddResponse) GetSourceOfWealthOk() (*[]SourceOfWealth, bool)`

GetSourceOfWealthOk returns a tuple with the SourceOfWealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfWealth

`func (o *CreateEddResponse) SetSourceOfWealth(v []SourceOfWealth)`

SetSourceOfWealth sets SourceOfWealth field to given value.

### HasSourceOfWealth

`func (o *CreateEddResponse) HasSourceOfWealth() bool`

HasSourceOfWealth returns a boolean if a field has been set.

### GetCountry

`func (o *CreateEddResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateEddResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateEddResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateEddResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEstimatedRevenue

`func (o *CreateEddResponse) GetEstimatedRevenue() EstimatedRevenue`

GetEstimatedRevenue returns the EstimatedRevenue field if non-nil, zero value otherwise.

### GetEstimatedRevenueOk

`func (o *CreateEddResponse) GetEstimatedRevenueOk() (*EstimatedRevenue, bool)`

GetEstimatedRevenueOk returns a tuple with the EstimatedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRevenue

`func (o *CreateEddResponse) SetEstimatedRevenue(v EstimatedRevenue)`

SetEstimatedRevenue sets EstimatedRevenue field to given value.

### HasEstimatedRevenue

`func (o *CreateEddResponse) HasEstimatedRevenue() bool`

HasEstimatedRevenue returns a boolean if a field has been set.

### GetIndustryType

`func (o *CreateEddResponse) GetIndustryType() IndustryType`

GetIndustryType returns the IndustryType field if non-nil, zero value otherwise.

### GetIndustryTypeOk

`func (o *CreateEddResponse) GetIndustryTypeOk() (*IndustryType, bool)`

GetIndustryTypeOk returns a tuple with the IndustryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryType

`func (o *CreateEddResponse) SetIndustryType(v IndustryType)`

SetIndustryType sets IndustryType field to given value.

### HasIndustryType

`func (o *CreateEddResponse) HasIndustryType() bool`

HasIndustryType returns a boolean if a field has been set.

### GetRecurringWireUsage

`func (o *CreateEddResponse) GetRecurringWireUsage() bool`

GetRecurringWireUsage returns the RecurringWireUsage field if non-nil, zero value otherwise.

### GetRecurringWireUsageOk

`func (o *CreateEddResponse) GetRecurringWireUsageOk() (*bool, bool)`

GetRecurringWireUsageOk returns a tuple with the RecurringWireUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringWireUsage

`func (o *CreateEddResponse) SetRecurringWireUsage(v bool)`

SetRecurringWireUsage sets RecurringWireUsage field to given value.

### HasRecurringWireUsage

`func (o *CreateEddResponse) HasRecurringWireUsage() bool`

HasRecurringWireUsage returns a boolean if a field has been set.

### GetSpecificInvolvement

`func (o *CreateEddResponse) GetSpecificInvolvement() SpecificInvolvement`

GetSpecificInvolvement returns the SpecificInvolvement field if non-nil, zero value otherwise.

### GetSpecificInvolvementOk

`func (o *CreateEddResponse) GetSpecificInvolvementOk() (*SpecificInvolvement, bool)`

GetSpecificInvolvementOk returns a tuple with the SpecificInvolvement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificInvolvement

`func (o *CreateEddResponse) SetSpecificInvolvement(v SpecificInvolvement)`

SetSpecificInvolvement sets SpecificInvolvement field to given value.

### HasSpecificInvolvement

`func (o *CreateEddResponse) HasSpecificInvolvement() bool`

HasSpecificInvolvement returns a boolean if a field has been set.

### GetTransactionVolume

`func (o *CreateEddResponse) GetTransactionVolume() []TransactionVolume`

GetTransactionVolume returns the TransactionVolume field if non-nil, zero value otherwise.

### GetTransactionVolumeOk

`func (o *CreateEddResponse) GetTransactionVolumeOk() (*[]TransactionVolume, bool)`

GetTransactionVolumeOk returns a tuple with the TransactionVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionVolume

`func (o *CreateEddResponse) SetTransactionVolume(v []TransactionVolume)`

SetTransactionVolume sets TransactionVolume field to given value.

### HasTransactionVolume

`func (o *CreateEddResponse) HasTransactionVolume() bool`

HasTransactionVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


