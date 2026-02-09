# CreditApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Account ID for the application. Only required if application purpose is not ACCOUNT_OPENING | [optional] 
**AccountType** | [**ApplicationAccountType**](ApplicationAccountType.md) |  | 
**Applicants** | [**[]Applicant**](Applicant.md) |  | 
**ApplicationSubmittedTime** | Pointer to **time.Time** | Application submitted timestamp in RFC3339 format | [optional] 
**CreditDecisionTime** | Pointer to **time.Time** | Credit decision timestamp in RFC3339 format | [optional] 
**CustomerResponseTime** | Pointer to **time.Time** | Credit decision timestamp in RFC3339 format | [optional] 
**Purpose** | [**CreditApplicationPurpose**](CreditApplicationPurpose.md) |  | 
**Status** | [**CreditApplicationStatus**](CreditApplicationStatus.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 

## Methods

### NewCreditApplication

`func NewCreditApplication(accountType ApplicationAccountType, applicants []Applicant, purpose CreditApplicationPurpose, status CreditApplicationStatus, type_ ApplicationType, ) *CreditApplication`

NewCreditApplication instantiates a new CreditApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditApplicationWithDefaults

`func NewCreditApplicationWithDefaults() *CreditApplication`

NewCreditApplicationWithDefaults instantiates a new CreditApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreditApplication) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreditApplication) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreditApplication) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreditApplication) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountType

`func (o *CreditApplication) GetAccountType() ApplicationAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *CreditApplication) GetAccountTypeOk() (*ApplicationAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *CreditApplication) SetAccountType(v ApplicationAccountType)`

SetAccountType sets AccountType field to given value.


### GetApplicants

`func (o *CreditApplication) GetApplicants() []Applicant`

GetApplicants returns the Applicants field if non-nil, zero value otherwise.

### GetApplicantsOk

`func (o *CreditApplication) GetApplicantsOk() (*[]Applicant, bool)`

GetApplicantsOk returns a tuple with the Applicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicants

`func (o *CreditApplication) SetApplicants(v []Applicant)`

SetApplicants sets Applicants field to given value.


### GetApplicationSubmittedTime

`func (o *CreditApplication) GetApplicationSubmittedTime() time.Time`

GetApplicationSubmittedTime returns the ApplicationSubmittedTime field if non-nil, zero value otherwise.

### GetApplicationSubmittedTimeOk

`func (o *CreditApplication) GetApplicationSubmittedTimeOk() (*time.Time, bool)`

GetApplicationSubmittedTimeOk returns a tuple with the ApplicationSubmittedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSubmittedTime

`func (o *CreditApplication) SetApplicationSubmittedTime(v time.Time)`

SetApplicationSubmittedTime sets ApplicationSubmittedTime field to given value.

### HasApplicationSubmittedTime

`func (o *CreditApplication) HasApplicationSubmittedTime() bool`

HasApplicationSubmittedTime returns a boolean if a field has been set.

### GetCreditDecisionTime

`func (o *CreditApplication) GetCreditDecisionTime() time.Time`

GetCreditDecisionTime returns the CreditDecisionTime field if non-nil, zero value otherwise.

### GetCreditDecisionTimeOk

`func (o *CreditApplication) GetCreditDecisionTimeOk() (*time.Time, bool)`

GetCreditDecisionTimeOk returns a tuple with the CreditDecisionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDecisionTime

`func (o *CreditApplication) SetCreditDecisionTime(v time.Time)`

SetCreditDecisionTime sets CreditDecisionTime field to given value.

### HasCreditDecisionTime

`func (o *CreditApplication) HasCreditDecisionTime() bool`

HasCreditDecisionTime returns a boolean if a field has been set.

### GetCustomerResponseTime

`func (o *CreditApplication) GetCustomerResponseTime() time.Time`

GetCustomerResponseTime returns the CustomerResponseTime field if non-nil, zero value otherwise.

### GetCustomerResponseTimeOk

`func (o *CreditApplication) GetCustomerResponseTimeOk() (*time.Time, bool)`

GetCustomerResponseTimeOk returns a tuple with the CustomerResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerResponseTime

`func (o *CreditApplication) SetCustomerResponseTime(v time.Time)`

SetCustomerResponseTime sets CustomerResponseTime field to given value.

### HasCustomerResponseTime

`func (o *CreditApplication) HasCustomerResponseTime() bool`

HasCustomerResponseTime returns a boolean if a field has been set.

### GetPurpose

`func (o *CreditApplication) GetPurpose() CreditApplicationPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreditApplication) GetPurposeOk() (*CreditApplicationPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreditApplication) SetPurpose(v CreditApplicationPurpose)`

SetPurpose sets Purpose field to given value.


### GetStatus

`func (o *CreditApplication) GetStatus() CreditApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditApplication) GetStatusOk() (*CreditApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditApplication) SetStatus(v CreditApplicationStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *CreditApplication) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreditApplication) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreditApplication) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CreditApplication) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *CreditApplication) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditApplication) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditApplication) SetType(v ApplicationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


