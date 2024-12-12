# CreateEddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalQuestions** | Pointer to [**[]Question**](Question.md) | Additional questions regarding the related resource | [optional] 
**CaseId** | Pointer to **int32** | The ID of the case related to this EDD record | [optional] 
**Reason** | **string** | The reason for this EDD record to be requested | 
**RelatedResourceId** | **string** | related resource UUID | 
**RelatedResourceType** | [**RelatedResourceType1**](RelatedResourceType1.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**SourceOfFunds** | Pointer to **string** | The source of funds for the transaction. | [optional] 
**TransactionPurpose** | Pointer to **string** | The purpose of the transaction. | [optional] 
**Country** | Pointer to **string** | ISO-3166-1 Alpha-2 country code | [optional] 
**EstimatedRevenue** | Pointer to [**EstimatedRevenue**](EstimatedRevenue.md) |  | [optional] 
**IndustryType** | Pointer to [**IndustryType**](IndustryType.md) |  | [optional] 
**NegativeNewsFindings** | Pointer to **int32** | The number of negative news findings. | [optional] 
**RecurringWireUsage** | Pointer to **bool** | True if the customer is expected to send or receive wire transfers at a regular frequency. | [optional] 
**SpecificInvolvement** | Pointer to [**SpecificInvolvement**](SpecificInvolvement.md) |  | [optional] 
**TransactionVolume** | Pointer to [**[]TransactionVolume**](TransactionVolume.md) | Array of transaction volumes. | [optional] 
**CitizenshipCountries** | Pointer to **[]string** | List of countries where the related customer holds citizenship. | [optional] 
**EmploymentType** | Pointer to **string** | The type of employment. | [optional] 
**Income** | Pointer to [**Income**](Income.md) |  | [optional] 
**Occupation** | Pointer to **string** | The occupation of the related resource. | [optional] 
**OccupationIndustry** | Pointer to [**IndustryType**](IndustryType.md) |  | [optional] 
**RecurringDirectDeposit** | Pointer to **bool** | True if the customer is expected to use direct deposit at a regular frequency. | [optional] 
**ResidenceType** | Pointer to **string** | The type of residence. | [optional] 
**ResidentialExpense** | Pointer to [**ResidentialExpense**](ResidentialExpense.md) |  | [optional] 
**SourceOfWealth** | Pointer to [**[]SourceOfWealth**](SourceOfWealth.md) | The sources of wealth for the customer. | [optional] 

## Methods

### NewCreateEddRequest

`func NewCreateEddRequest(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, ) *CreateEddRequest`

NewCreateEddRequest instantiates a new CreateEddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEddRequestWithDefaults

`func NewCreateEddRequestWithDefaults() *CreateEddRequest`

NewCreateEddRequestWithDefaults instantiates a new CreateEddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalQuestions

`func (o *CreateEddRequest) GetAdditionalQuestions() []Question`

GetAdditionalQuestions returns the AdditionalQuestions field if non-nil, zero value otherwise.

### GetAdditionalQuestionsOk

`func (o *CreateEddRequest) GetAdditionalQuestionsOk() (*[]Question, bool)`

GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuestions

`func (o *CreateEddRequest) SetAdditionalQuestions(v []Question)`

SetAdditionalQuestions sets AdditionalQuestions field to given value.

### HasAdditionalQuestions

`func (o *CreateEddRequest) HasAdditionalQuestions() bool`

HasAdditionalQuestions returns a boolean if a field has been set.

### GetCaseId

`func (o *CreateEddRequest) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *CreateEddRequest) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *CreateEddRequest) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *CreateEddRequest) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReason

`func (o *CreateEddRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateEddRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateEddRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRelatedResourceId

`func (o *CreateEddRequest) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *CreateEddRequest) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *CreateEddRequest) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *CreateEddRequest) GetRelatedResourceType() RelatedResourceType1`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *CreateEddRequest) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *CreateEddRequest) SetRelatedResourceType(v RelatedResourceType1)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetTenant

`func (o *CreateEddRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreateEddRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreateEddRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CreateEddRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetSourceOfFunds

`func (o *CreateEddRequest) GetSourceOfFunds() string`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *CreateEddRequest) GetSourceOfFundsOk() (*string, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *CreateEddRequest) SetSourceOfFunds(v string)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *CreateEddRequest) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetTransactionPurpose

`func (o *CreateEddRequest) GetTransactionPurpose() string`

GetTransactionPurpose returns the TransactionPurpose field if non-nil, zero value otherwise.

### GetTransactionPurposeOk

`func (o *CreateEddRequest) GetTransactionPurposeOk() (*string, bool)`

GetTransactionPurposeOk returns a tuple with the TransactionPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPurpose

`func (o *CreateEddRequest) SetTransactionPurpose(v string)`

SetTransactionPurpose sets TransactionPurpose field to given value.

### HasTransactionPurpose

`func (o *CreateEddRequest) HasTransactionPurpose() bool`

HasTransactionPurpose returns a boolean if a field has been set.

### GetCountry

`func (o *CreateEddRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateEddRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateEddRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateEddRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEstimatedRevenue

`func (o *CreateEddRequest) GetEstimatedRevenue() EstimatedRevenue`

GetEstimatedRevenue returns the EstimatedRevenue field if non-nil, zero value otherwise.

### GetEstimatedRevenueOk

`func (o *CreateEddRequest) GetEstimatedRevenueOk() (*EstimatedRevenue, bool)`

GetEstimatedRevenueOk returns a tuple with the EstimatedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRevenue

`func (o *CreateEddRequest) SetEstimatedRevenue(v EstimatedRevenue)`

SetEstimatedRevenue sets EstimatedRevenue field to given value.

### HasEstimatedRevenue

`func (o *CreateEddRequest) HasEstimatedRevenue() bool`

HasEstimatedRevenue returns a boolean if a field has been set.

### GetIndustryType

`func (o *CreateEddRequest) GetIndustryType() IndustryType`

GetIndustryType returns the IndustryType field if non-nil, zero value otherwise.

### GetIndustryTypeOk

`func (o *CreateEddRequest) GetIndustryTypeOk() (*IndustryType, bool)`

GetIndustryTypeOk returns a tuple with the IndustryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryType

`func (o *CreateEddRequest) SetIndustryType(v IndustryType)`

SetIndustryType sets IndustryType field to given value.

### HasIndustryType

`func (o *CreateEddRequest) HasIndustryType() bool`

HasIndustryType returns a boolean if a field has been set.

### GetNegativeNewsFindings

`func (o *CreateEddRequest) GetNegativeNewsFindings() int32`

GetNegativeNewsFindings returns the NegativeNewsFindings field if non-nil, zero value otherwise.

### GetNegativeNewsFindingsOk

`func (o *CreateEddRequest) GetNegativeNewsFindingsOk() (*int32, bool)`

GetNegativeNewsFindingsOk returns a tuple with the NegativeNewsFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeNewsFindings

`func (o *CreateEddRequest) SetNegativeNewsFindings(v int32)`

SetNegativeNewsFindings sets NegativeNewsFindings field to given value.

### HasNegativeNewsFindings

`func (o *CreateEddRequest) HasNegativeNewsFindings() bool`

HasNegativeNewsFindings returns a boolean if a field has been set.

### GetRecurringWireUsage

`func (o *CreateEddRequest) GetRecurringWireUsage() bool`

GetRecurringWireUsage returns the RecurringWireUsage field if non-nil, zero value otherwise.

### GetRecurringWireUsageOk

`func (o *CreateEddRequest) GetRecurringWireUsageOk() (*bool, bool)`

GetRecurringWireUsageOk returns a tuple with the RecurringWireUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringWireUsage

`func (o *CreateEddRequest) SetRecurringWireUsage(v bool)`

SetRecurringWireUsage sets RecurringWireUsage field to given value.

### HasRecurringWireUsage

`func (o *CreateEddRequest) HasRecurringWireUsage() bool`

HasRecurringWireUsage returns a boolean if a field has been set.

### GetSpecificInvolvement

`func (o *CreateEddRequest) GetSpecificInvolvement() SpecificInvolvement`

GetSpecificInvolvement returns the SpecificInvolvement field if non-nil, zero value otherwise.

### GetSpecificInvolvementOk

`func (o *CreateEddRequest) GetSpecificInvolvementOk() (*SpecificInvolvement, bool)`

GetSpecificInvolvementOk returns a tuple with the SpecificInvolvement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificInvolvement

`func (o *CreateEddRequest) SetSpecificInvolvement(v SpecificInvolvement)`

SetSpecificInvolvement sets SpecificInvolvement field to given value.

### HasSpecificInvolvement

`func (o *CreateEddRequest) HasSpecificInvolvement() bool`

HasSpecificInvolvement returns a boolean if a field has been set.

### GetTransactionVolume

`func (o *CreateEddRequest) GetTransactionVolume() []TransactionVolume`

GetTransactionVolume returns the TransactionVolume field if non-nil, zero value otherwise.

### GetTransactionVolumeOk

`func (o *CreateEddRequest) GetTransactionVolumeOk() (*[]TransactionVolume, bool)`

GetTransactionVolumeOk returns a tuple with the TransactionVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionVolume

`func (o *CreateEddRequest) SetTransactionVolume(v []TransactionVolume)`

SetTransactionVolume sets TransactionVolume field to given value.

### HasTransactionVolume

`func (o *CreateEddRequest) HasTransactionVolume() bool`

HasTransactionVolume returns a boolean if a field has been set.

### GetCitizenshipCountries

`func (o *CreateEddRequest) GetCitizenshipCountries() []string`

GetCitizenshipCountries returns the CitizenshipCountries field if non-nil, zero value otherwise.

### GetCitizenshipCountriesOk

`func (o *CreateEddRequest) GetCitizenshipCountriesOk() (*[]string, bool)`

GetCitizenshipCountriesOk returns a tuple with the CitizenshipCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenshipCountries

`func (o *CreateEddRequest) SetCitizenshipCountries(v []string)`

SetCitizenshipCountries sets CitizenshipCountries field to given value.

### HasCitizenshipCountries

`func (o *CreateEddRequest) HasCitizenshipCountries() bool`

HasCitizenshipCountries returns a boolean if a field has been set.

### GetEmploymentType

`func (o *CreateEddRequest) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *CreateEddRequest) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *CreateEddRequest) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.

### HasEmploymentType

`func (o *CreateEddRequest) HasEmploymentType() bool`

HasEmploymentType returns a boolean if a field has been set.

### GetIncome

`func (o *CreateEddRequest) GetIncome() Income`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *CreateEddRequest) GetIncomeOk() (*Income, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *CreateEddRequest) SetIncome(v Income)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *CreateEddRequest) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetOccupation

`func (o *CreateEddRequest) GetOccupation() string`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *CreateEddRequest) GetOccupationOk() (*string, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *CreateEddRequest) SetOccupation(v string)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *CreateEddRequest) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### GetOccupationIndustry

`func (o *CreateEddRequest) GetOccupationIndustry() IndustryType`

GetOccupationIndustry returns the OccupationIndustry field if non-nil, zero value otherwise.

### GetOccupationIndustryOk

`func (o *CreateEddRequest) GetOccupationIndustryOk() (*IndustryType, bool)`

GetOccupationIndustryOk returns a tuple with the OccupationIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupationIndustry

`func (o *CreateEddRequest) SetOccupationIndustry(v IndustryType)`

SetOccupationIndustry sets OccupationIndustry field to given value.

### HasOccupationIndustry

`func (o *CreateEddRequest) HasOccupationIndustry() bool`

HasOccupationIndustry returns a boolean if a field has been set.

### GetRecurringDirectDeposit

`func (o *CreateEddRequest) GetRecurringDirectDeposit() bool`

GetRecurringDirectDeposit returns the RecurringDirectDeposit field if non-nil, zero value otherwise.

### GetRecurringDirectDepositOk

`func (o *CreateEddRequest) GetRecurringDirectDepositOk() (*bool, bool)`

GetRecurringDirectDepositOk returns a tuple with the RecurringDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDirectDeposit

`func (o *CreateEddRequest) SetRecurringDirectDeposit(v bool)`

SetRecurringDirectDeposit sets RecurringDirectDeposit field to given value.

### HasRecurringDirectDeposit

`func (o *CreateEddRequest) HasRecurringDirectDeposit() bool`

HasRecurringDirectDeposit returns a boolean if a field has been set.

### GetResidenceType

`func (o *CreateEddRequest) GetResidenceType() string`

GetResidenceType returns the ResidenceType field if non-nil, zero value otherwise.

### GetResidenceTypeOk

`func (o *CreateEddRequest) GetResidenceTypeOk() (*string, bool)`

GetResidenceTypeOk returns a tuple with the ResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceType

`func (o *CreateEddRequest) SetResidenceType(v string)`

SetResidenceType sets ResidenceType field to given value.

### HasResidenceType

`func (o *CreateEddRequest) HasResidenceType() bool`

HasResidenceType returns a boolean if a field has been set.

### GetResidentialExpense

`func (o *CreateEddRequest) GetResidentialExpense() ResidentialExpense`

GetResidentialExpense returns the ResidentialExpense field if non-nil, zero value otherwise.

### GetResidentialExpenseOk

`func (o *CreateEddRequest) GetResidentialExpenseOk() (*ResidentialExpense, bool)`

GetResidentialExpenseOk returns a tuple with the ResidentialExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentialExpense

`func (o *CreateEddRequest) SetResidentialExpense(v ResidentialExpense)`

SetResidentialExpense sets ResidentialExpense field to given value.

### HasResidentialExpense

`func (o *CreateEddRequest) HasResidentialExpense() bool`

HasResidentialExpense returns a boolean if a field has been set.

### GetSourceOfWealth

`func (o *CreateEddRequest) GetSourceOfWealth() []SourceOfWealth`

GetSourceOfWealth returns the SourceOfWealth field if non-nil, zero value otherwise.

### GetSourceOfWealthOk

`func (o *CreateEddRequest) GetSourceOfWealthOk() (*[]SourceOfWealth, bool)`

GetSourceOfWealthOk returns a tuple with the SourceOfWealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfWealth

`func (o *CreateEddRequest) SetSourceOfWealth(v []SourceOfWealth)`

SetSourceOfWealth sets SourceOfWealth field to given value.

### HasSourceOfWealth

`func (o *CreateEddRequest) HasSourceOfWealth() bool`

HasSourceOfWealth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


