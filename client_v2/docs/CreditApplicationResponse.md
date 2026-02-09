# CreditApplicationResponse

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
**CreationTime** | **time.Time** | Application creation timestamp in RFC3339 format | [readonly] 
**Id** | **string** | Generated ID for the application | [readonly] 
**LastUpdatedTime** | **time.Time** | Timestamp of the last application modification in RFC3339 format | [readonly] 

## Methods

### NewCreditApplicationResponse

`func NewCreditApplicationResponse(accountType ApplicationAccountType, applicants []Applicant, purpose CreditApplicationPurpose, status CreditApplicationStatus, type_ ApplicationType, creationTime time.Time, id string, lastUpdatedTime time.Time, ) *CreditApplicationResponse`

NewCreditApplicationResponse instantiates a new CreditApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditApplicationResponseWithDefaults

`func NewCreditApplicationResponseWithDefaults() *CreditApplicationResponse`

NewCreditApplicationResponseWithDefaults instantiates a new CreditApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreditApplicationResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreditApplicationResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreditApplicationResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreditApplicationResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountType

`func (o *CreditApplicationResponse) GetAccountType() ApplicationAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *CreditApplicationResponse) GetAccountTypeOk() (*ApplicationAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *CreditApplicationResponse) SetAccountType(v ApplicationAccountType)`

SetAccountType sets AccountType field to given value.


### GetApplicants

`func (o *CreditApplicationResponse) GetApplicants() []Applicant`

GetApplicants returns the Applicants field if non-nil, zero value otherwise.

### GetApplicantsOk

`func (o *CreditApplicationResponse) GetApplicantsOk() (*[]Applicant, bool)`

GetApplicantsOk returns a tuple with the Applicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicants

`func (o *CreditApplicationResponse) SetApplicants(v []Applicant)`

SetApplicants sets Applicants field to given value.


### GetApplicationSubmittedTime

`func (o *CreditApplicationResponse) GetApplicationSubmittedTime() time.Time`

GetApplicationSubmittedTime returns the ApplicationSubmittedTime field if non-nil, zero value otherwise.

### GetApplicationSubmittedTimeOk

`func (o *CreditApplicationResponse) GetApplicationSubmittedTimeOk() (*time.Time, bool)`

GetApplicationSubmittedTimeOk returns a tuple with the ApplicationSubmittedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSubmittedTime

`func (o *CreditApplicationResponse) SetApplicationSubmittedTime(v time.Time)`

SetApplicationSubmittedTime sets ApplicationSubmittedTime field to given value.

### HasApplicationSubmittedTime

`func (o *CreditApplicationResponse) HasApplicationSubmittedTime() bool`

HasApplicationSubmittedTime returns a boolean if a field has been set.

### GetCreditDecisionTime

`func (o *CreditApplicationResponse) GetCreditDecisionTime() time.Time`

GetCreditDecisionTime returns the CreditDecisionTime field if non-nil, zero value otherwise.

### GetCreditDecisionTimeOk

`func (o *CreditApplicationResponse) GetCreditDecisionTimeOk() (*time.Time, bool)`

GetCreditDecisionTimeOk returns a tuple with the CreditDecisionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDecisionTime

`func (o *CreditApplicationResponse) SetCreditDecisionTime(v time.Time)`

SetCreditDecisionTime sets CreditDecisionTime field to given value.

### HasCreditDecisionTime

`func (o *CreditApplicationResponse) HasCreditDecisionTime() bool`

HasCreditDecisionTime returns a boolean if a field has been set.

### GetCustomerResponseTime

`func (o *CreditApplicationResponse) GetCustomerResponseTime() time.Time`

GetCustomerResponseTime returns the CustomerResponseTime field if non-nil, zero value otherwise.

### GetCustomerResponseTimeOk

`func (o *CreditApplicationResponse) GetCustomerResponseTimeOk() (*time.Time, bool)`

GetCustomerResponseTimeOk returns a tuple with the CustomerResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerResponseTime

`func (o *CreditApplicationResponse) SetCustomerResponseTime(v time.Time)`

SetCustomerResponseTime sets CustomerResponseTime field to given value.

### HasCustomerResponseTime

`func (o *CreditApplicationResponse) HasCustomerResponseTime() bool`

HasCustomerResponseTime returns a boolean if a field has been set.

### GetPurpose

`func (o *CreditApplicationResponse) GetPurpose() CreditApplicationPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreditApplicationResponse) GetPurposeOk() (*CreditApplicationPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreditApplicationResponse) SetPurpose(v CreditApplicationPurpose)`

SetPurpose sets Purpose field to given value.


### GetStatus

`func (o *CreditApplicationResponse) GetStatus() CreditApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditApplicationResponse) GetStatusOk() (*CreditApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditApplicationResponse) SetStatus(v CreditApplicationStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *CreditApplicationResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreditApplicationResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreditApplicationResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CreditApplicationResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *CreditApplicationResponse) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditApplicationResponse) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditApplicationResponse) SetType(v ApplicationType)`

SetType sets Type field to given value.


### GetCreationTime

`func (o *CreditApplicationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreditApplicationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreditApplicationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *CreditApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *CreditApplicationResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CreditApplicationResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CreditApplicationResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


