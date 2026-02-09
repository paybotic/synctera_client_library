# ApplicationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applicants** | Pointer to [**[]Applicant**](Applicant.md) |  | [optional] 
**Status** | Pointer to [**CreditCardApplicationStatus**](CreditCardApplicationStatus.md) |  | [optional] 
**ApplicationDetails** | Pointer to [**RestrictedApplicationDetails**](RestrictedApplicationDetails.md) |  | [optional] 
**Description** | Pointer to **string** | A description of the restricted account application | [optional] 
**CreditLimit** | Pointer to **int32** | Credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 
**Decision** | Pointer to [**CreditCardDecision**](CreditCardDecision.md) |  | [optional] 
**ExternalId** | Pointer to **string** | External identifier (e.g., Taktile decision_id) | [optional] 
**InterestRate** | Pointer to **int32** | Interest rate in basis points (e.g., 1250 &#x3D; 12.50%) | [optional] 
**RejectionReasons** | Pointer to **[]string** | Reasons for rejection (when decision is declined) | [optional] 

## Methods

### NewApplicationPatch

`func NewApplicationPatch() *ApplicationPatch`

NewApplicationPatch instantiates a new ApplicationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPatchWithDefaults

`func NewApplicationPatchWithDefaults() *ApplicationPatch`

NewApplicationPatchWithDefaults instantiates a new ApplicationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicants

`func (o *ApplicationPatch) GetApplicants() []Applicant`

GetApplicants returns the Applicants field if non-nil, zero value otherwise.

### GetApplicantsOk

`func (o *ApplicationPatch) GetApplicantsOk() (*[]Applicant, bool)`

GetApplicantsOk returns a tuple with the Applicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicants

`func (o *ApplicationPatch) SetApplicants(v []Applicant)`

SetApplicants sets Applicants field to given value.

### HasApplicants

`func (o *ApplicationPatch) HasApplicants() bool`

HasApplicants returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationPatch) GetStatus() CreditCardApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationPatch) GetStatusOk() (*CreditCardApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationPatch) SetStatus(v CreditCardApplicationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApplicationDetails

`func (o *ApplicationPatch) GetApplicationDetails() RestrictedApplicationDetails`

GetApplicationDetails returns the ApplicationDetails field if non-nil, zero value otherwise.

### GetApplicationDetailsOk

`func (o *ApplicationPatch) GetApplicationDetailsOk() (*RestrictedApplicationDetails, bool)`

GetApplicationDetailsOk returns a tuple with the ApplicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDetails

`func (o *ApplicationPatch) SetApplicationDetails(v RestrictedApplicationDetails)`

SetApplicationDetails sets ApplicationDetails field to given value.

### HasApplicationDetails

`func (o *ApplicationPatch) HasApplicationDetails() bool`

HasApplicationDetails returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreditLimit

`func (o *ApplicationPatch) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *ApplicationPatch) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *ApplicationPatch) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *ApplicationPatch) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetDecision

`func (o *ApplicationPatch) GetDecision() CreditCardDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ApplicationPatch) GetDecisionOk() (*CreditCardDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ApplicationPatch) SetDecision(v CreditCardDecision)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *ApplicationPatch) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetExternalId

`func (o *ApplicationPatch) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ApplicationPatch) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ApplicationPatch) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ApplicationPatch) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInterestRate

`func (o *ApplicationPatch) GetInterestRate() int32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *ApplicationPatch) GetInterestRateOk() (*int32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *ApplicationPatch) SetInterestRate(v int32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *ApplicationPatch) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetRejectionReasons

`func (o *ApplicationPatch) GetRejectionReasons() []string`

GetRejectionReasons returns the RejectionReasons field if non-nil, zero value otherwise.

### GetRejectionReasonsOk

`func (o *ApplicationPatch) GetRejectionReasonsOk() (*[]string, bool)`

GetRejectionReasonsOk returns a tuple with the RejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReasons

`func (o *ApplicationPatch) SetRejectionReasons(v []string)`

SetRejectionReasons sets RejectionReasons field to given value.

### HasRejectionReasons

`func (o *ApplicationPatch) HasRejectionReasons() bool`

HasRejectionReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


