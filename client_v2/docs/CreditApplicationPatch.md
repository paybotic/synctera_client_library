# CreditApplicationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applicants** | Pointer to [**[]Applicant**](Applicant.md) |  | [optional] 
**Status** | Pointer to [**CreditApplicationStatus**](CreditApplicationStatus.md) |  | [optional] 

## Methods

### NewCreditApplicationPatch

`func NewCreditApplicationPatch() *CreditApplicationPatch`

NewCreditApplicationPatch instantiates a new CreditApplicationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditApplicationPatchWithDefaults

`func NewCreditApplicationPatchWithDefaults() *CreditApplicationPatch`

NewCreditApplicationPatchWithDefaults instantiates a new CreditApplicationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicants

`func (o *CreditApplicationPatch) GetApplicants() []Applicant`

GetApplicants returns the Applicants field if non-nil, zero value otherwise.

### GetApplicantsOk

`func (o *CreditApplicationPatch) GetApplicantsOk() (*[]Applicant, bool)`

GetApplicantsOk returns a tuple with the Applicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicants

`func (o *CreditApplicationPatch) SetApplicants(v []Applicant)`

SetApplicants sets Applicants field to given value.

### HasApplicants

`func (o *CreditApplicationPatch) HasApplicants() bool`

HasApplicants returns a boolean if a field has been set.

### GetStatus

`func (o *CreditApplicationPatch) GetStatus() CreditApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditApplicationPatch) GetStatusOk() (*CreditApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditApplicationPatch) SetStatus(v CreditApplicationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreditApplicationPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


