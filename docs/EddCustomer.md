# EddCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalQuestions** | Pointer to [**[]Question**](Question.md) | Additional questions regarding the related resource | [optional] 
**CaseId** | Pointer to **int32** | The ID of the case related to this EDD record | [optional] 
**Reason** | **string** | The reason for this EDD record to be requested | 
**RelatedResourceId** | **string** | related resource UUID | 
**RelatedResourceType** | [**RelatedResourceType1**](RelatedResourceType1.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
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

## Methods

### NewEddCustomer

`func NewEddCustomer(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, ) *EddCustomer`

NewEddCustomer instantiates a new EddCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEddCustomerWithDefaults

`func NewEddCustomerWithDefaults() *EddCustomer`

NewEddCustomerWithDefaults instantiates a new EddCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalQuestions

`func (o *EddCustomer) GetAdditionalQuestions() []Question`

GetAdditionalQuestions returns the AdditionalQuestions field if non-nil, zero value otherwise.

### GetAdditionalQuestionsOk

`func (o *EddCustomer) GetAdditionalQuestionsOk() (*[]Question, bool)`

GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuestions

`func (o *EddCustomer) SetAdditionalQuestions(v []Question)`

SetAdditionalQuestions sets AdditionalQuestions field to given value.

### HasAdditionalQuestions

`func (o *EddCustomer) HasAdditionalQuestions() bool`

HasAdditionalQuestions returns a boolean if a field has been set.

### GetCaseId

`func (o *EddCustomer) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *EddCustomer) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *EddCustomer) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *EddCustomer) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReason

`func (o *EddCustomer) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EddCustomer) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EddCustomer) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRelatedResourceId

`func (o *EddCustomer) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *EddCustomer) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *EddCustomer) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *EddCustomer) GetRelatedResourceType() RelatedResourceType1`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *EddCustomer) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *EddCustomer) SetRelatedResourceType(v RelatedResourceType1)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetTenant

`func (o *EddCustomer) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *EddCustomer) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *EddCustomer) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *EddCustomer) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetCitizenshipCountries

`func (o *EddCustomer) GetCitizenshipCountries() []string`

GetCitizenshipCountries returns the CitizenshipCountries field if non-nil, zero value otherwise.

### GetCitizenshipCountriesOk

`func (o *EddCustomer) GetCitizenshipCountriesOk() (*[]string, bool)`

GetCitizenshipCountriesOk returns a tuple with the CitizenshipCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenshipCountries

`func (o *EddCustomer) SetCitizenshipCountries(v []string)`

SetCitizenshipCountries sets CitizenshipCountries field to given value.

### HasCitizenshipCountries

`func (o *EddCustomer) HasCitizenshipCountries() bool`

HasCitizenshipCountries returns a boolean if a field has been set.

### GetEmploymentType

`func (o *EddCustomer) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *EddCustomer) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *EddCustomer) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.

### HasEmploymentType

`func (o *EddCustomer) HasEmploymentType() bool`

HasEmploymentType returns a boolean if a field has been set.

### GetIncome

`func (o *EddCustomer) GetIncome() Income`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *EddCustomer) GetIncomeOk() (*Income, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *EddCustomer) SetIncome(v Income)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *EddCustomer) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetNegativeNewsFindings

`func (o *EddCustomer) GetNegativeNewsFindings() int32`

GetNegativeNewsFindings returns the NegativeNewsFindings field if non-nil, zero value otherwise.

### GetNegativeNewsFindingsOk

`func (o *EddCustomer) GetNegativeNewsFindingsOk() (*int32, bool)`

GetNegativeNewsFindingsOk returns a tuple with the NegativeNewsFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeNewsFindings

`func (o *EddCustomer) SetNegativeNewsFindings(v int32)`

SetNegativeNewsFindings sets NegativeNewsFindings field to given value.

### HasNegativeNewsFindings

`func (o *EddCustomer) HasNegativeNewsFindings() bool`

HasNegativeNewsFindings returns a boolean if a field has been set.

### GetOccupation

`func (o *EddCustomer) GetOccupation() string`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *EddCustomer) GetOccupationOk() (*string, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *EddCustomer) SetOccupation(v string)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *EddCustomer) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### GetOccupationIndustry

`func (o *EddCustomer) GetOccupationIndustry() IndustryType`

GetOccupationIndustry returns the OccupationIndustry field if non-nil, zero value otherwise.

### GetOccupationIndustryOk

`func (o *EddCustomer) GetOccupationIndustryOk() (*IndustryType, bool)`

GetOccupationIndustryOk returns a tuple with the OccupationIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupationIndustry

`func (o *EddCustomer) SetOccupationIndustry(v IndustryType)`

SetOccupationIndustry sets OccupationIndustry field to given value.

### HasOccupationIndustry

`func (o *EddCustomer) HasOccupationIndustry() bool`

HasOccupationIndustry returns a boolean if a field has been set.

### GetRecurringDirectDeposit

`func (o *EddCustomer) GetRecurringDirectDeposit() bool`

GetRecurringDirectDeposit returns the RecurringDirectDeposit field if non-nil, zero value otherwise.

### GetRecurringDirectDepositOk

`func (o *EddCustomer) GetRecurringDirectDepositOk() (*bool, bool)`

GetRecurringDirectDepositOk returns a tuple with the RecurringDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDirectDeposit

`func (o *EddCustomer) SetRecurringDirectDeposit(v bool)`

SetRecurringDirectDeposit sets RecurringDirectDeposit field to given value.

### HasRecurringDirectDeposit

`func (o *EddCustomer) HasRecurringDirectDeposit() bool`

HasRecurringDirectDeposit returns a boolean if a field has been set.

### GetResidenceType

`func (o *EddCustomer) GetResidenceType() string`

GetResidenceType returns the ResidenceType field if non-nil, zero value otherwise.

### GetResidenceTypeOk

`func (o *EddCustomer) GetResidenceTypeOk() (*string, bool)`

GetResidenceTypeOk returns a tuple with the ResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceType

`func (o *EddCustomer) SetResidenceType(v string)`

SetResidenceType sets ResidenceType field to given value.

### HasResidenceType

`func (o *EddCustomer) HasResidenceType() bool`

HasResidenceType returns a boolean if a field has been set.

### GetResidentialExpense

`func (o *EddCustomer) GetResidentialExpense() ResidentialExpense`

GetResidentialExpense returns the ResidentialExpense field if non-nil, zero value otherwise.

### GetResidentialExpenseOk

`func (o *EddCustomer) GetResidentialExpenseOk() (*ResidentialExpense, bool)`

GetResidentialExpenseOk returns a tuple with the ResidentialExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentialExpense

`func (o *EddCustomer) SetResidentialExpense(v ResidentialExpense)`

SetResidentialExpense sets ResidentialExpense field to given value.

### HasResidentialExpense

`func (o *EddCustomer) HasResidentialExpense() bool`

HasResidentialExpense returns a boolean if a field has been set.

### GetSourceOfWealth

`func (o *EddCustomer) GetSourceOfWealth() []SourceOfWealth`

GetSourceOfWealth returns the SourceOfWealth field if non-nil, zero value otherwise.

### GetSourceOfWealthOk

`func (o *EddCustomer) GetSourceOfWealthOk() (*[]SourceOfWealth, bool)`

GetSourceOfWealthOk returns a tuple with the SourceOfWealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfWealth

`func (o *EddCustomer) SetSourceOfWealth(v []SourceOfWealth)`

SetSourceOfWealth sets SourceOfWealth field to given value.

### HasSourceOfWealth

`func (o *EddCustomer) HasSourceOfWealth() bool`

HasSourceOfWealth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


