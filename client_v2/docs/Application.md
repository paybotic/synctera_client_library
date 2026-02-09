# Application

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
**Status** | [**RestrictedApplicationStatus**](RestrictedApplicationStatus.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 
**ApplicationDetails** | [**RestrictedApplicationDetails**](RestrictedApplicationDetails.md) |  | 
**BusinessId** | Pointer to **string** | Business ID for the application. An application must have either a Business or a customer associated with it. | [optional] 
**CustomerId** | Pointer to **string** | Customer ID for the application. An application must have either a Business or a customer associated with it. | [optional] 
**Description** | Pointer to **string** | A description of the restricted account application | [optional] 
**AccountTemplateId** | **string** | Mainapi account template ID to use for creating the credit card account | 
**Applicant** | [**CreditCardApplicant**](CreditCardApplicant.md) |  | 
**Attributes** | Pointer to **map[string]interface{}** | Raw JSON key-value pairs for flexible data | [optional] 
**Metadata** | Pointer to [**CreditCardApplicationMetadata**](CreditCardApplicationMetadata.md) |  | [optional] 
**Provider** | [**CreditCardProvider**](CreditCardProvider.md) |  | 
**RequestedCreditLimit** | Pointer to **int32** | Requested credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 

## Methods

### NewApplication

`func NewApplication(accountType ApplicationAccountType, applicants []Applicant, purpose CreditApplicationPurpose, status RestrictedApplicationStatus, tenant string, type_ ApplicationType, applicationDetails RestrictedApplicationDetails, accountTemplateId string, applicant CreditCardApplicant, provider CreditCardProvider, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Application) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Application) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Application) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Application) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountType

`func (o *Application) GetAccountType() ApplicationAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Application) GetAccountTypeOk() (*ApplicationAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Application) SetAccountType(v ApplicationAccountType)`

SetAccountType sets AccountType field to given value.


### GetApplicants

`func (o *Application) GetApplicants() []Applicant`

GetApplicants returns the Applicants field if non-nil, zero value otherwise.

### GetApplicantsOk

`func (o *Application) GetApplicantsOk() (*[]Applicant, bool)`

GetApplicantsOk returns a tuple with the Applicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicants

`func (o *Application) SetApplicants(v []Applicant)`

SetApplicants sets Applicants field to given value.


### GetApplicationSubmittedTime

`func (o *Application) GetApplicationSubmittedTime() time.Time`

GetApplicationSubmittedTime returns the ApplicationSubmittedTime field if non-nil, zero value otherwise.

### GetApplicationSubmittedTimeOk

`func (o *Application) GetApplicationSubmittedTimeOk() (*time.Time, bool)`

GetApplicationSubmittedTimeOk returns a tuple with the ApplicationSubmittedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSubmittedTime

`func (o *Application) SetApplicationSubmittedTime(v time.Time)`

SetApplicationSubmittedTime sets ApplicationSubmittedTime field to given value.

### HasApplicationSubmittedTime

`func (o *Application) HasApplicationSubmittedTime() bool`

HasApplicationSubmittedTime returns a boolean if a field has been set.

### GetCreditDecisionTime

`func (o *Application) GetCreditDecisionTime() time.Time`

GetCreditDecisionTime returns the CreditDecisionTime field if non-nil, zero value otherwise.

### GetCreditDecisionTimeOk

`func (o *Application) GetCreditDecisionTimeOk() (*time.Time, bool)`

GetCreditDecisionTimeOk returns a tuple with the CreditDecisionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDecisionTime

`func (o *Application) SetCreditDecisionTime(v time.Time)`

SetCreditDecisionTime sets CreditDecisionTime field to given value.

### HasCreditDecisionTime

`func (o *Application) HasCreditDecisionTime() bool`

HasCreditDecisionTime returns a boolean if a field has been set.

### GetCustomerResponseTime

`func (o *Application) GetCustomerResponseTime() time.Time`

GetCustomerResponseTime returns the CustomerResponseTime field if non-nil, zero value otherwise.

### GetCustomerResponseTimeOk

`func (o *Application) GetCustomerResponseTimeOk() (*time.Time, bool)`

GetCustomerResponseTimeOk returns a tuple with the CustomerResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerResponseTime

`func (o *Application) SetCustomerResponseTime(v time.Time)`

SetCustomerResponseTime sets CustomerResponseTime field to given value.

### HasCustomerResponseTime

`func (o *Application) HasCustomerResponseTime() bool`

HasCustomerResponseTime returns a boolean if a field has been set.

### GetPurpose

`func (o *Application) GetPurpose() CreditApplicationPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *Application) GetPurposeOk() (*CreditApplicationPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *Application) SetPurpose(v CreditApplicationPurpose)`

SetPurpose sets Purpose field to given value.


### GetStatus

`func (o *Application) GetStatus() RestrictedApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Application) GetStatusOk() (*RestrictedApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Application) SetStatus(v RestrictedApplicationStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *Application) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Application) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Application) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetType

`func (o *Application) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Application) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Application) SetType(v ApplicationType)`

SetType sets Type field to given value.


### GetApplicationDetails

`func (o *Application) GetApplicationDetails() RestrictedApplicationDetails`

GetApplicationDetails returns the ApplicationDetails field if non-nil, zero value otherwise.

### GetApplicationDetailsOk

`func (o *Application) GetApplicationDetailsOk() (*RestrictedApplicationDetails, bool)`

GetApplicationDetailsOk returns a tuple with the ApplicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDetails

`func (o *Application) SetApplicationDetails(v RestrictedApplicationDetails)`

SetApplicationDetails sets ApplicationDetails field to given value.


### GetBusinessId

`func (o *Application) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *Application) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *Application) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *Application) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *Application) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Application) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Application) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Application) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountTemplateId

`func (o *Application) GetAccountTemplateId() string`

GetAccountTemplateId returns the AccountTemplateId field if non-nil, zero value otherwise.

### GetAccountTemplateIdOk

`func (o *Application) GetAccountTemplateIdOk() (*string, bool)`

GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateId

`func (o *Application) SetAccountTemplateId(v string)`

SetAccountTemplateId sets AccountTemplateId field to given value.


### GetApplicant

`func (o *Application) GetApplicant() CreditCardApplicant`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *Application) GetApplicantOk() (*CreditCardApplicant, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *Application) SetApplicant(v CreditCardApplicant)`

SetApplicant sets Applicant field to given value.


### GetAttributes

`func (o *Application) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Application) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Application) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Application) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMetadata

`func (o *Application) GetMetadata() CreditCardApplicationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Application) GetMetadataOk() (*CreditCardApplicationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Application) SetMetadata(v CreditCardApplicationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Application) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProvider

`func (o *Application) GetProvider() CreditCardProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Application) GetProviderOk() (*CreditCardProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Application) SetProvider(v CreditCardProvider)`

SetProvider sets Provider field to given value.


### GetRequestedCreditLimit

`func (o *Application) GetRequestedCreditLimit() int32`

GetRequestedCreditLimit returns the RequestedCreditLimit field if non-nil, zero value otherwise.

### GetRequestedCreditLimitOk

`func (o *Application) GetRequestedCreditLimitOk() (*int32, bool)`

GetRequestedCreditLimitOk returns a tuple with the RequestedCreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCreditLimit

`func (o *Application) SetRequestedCreditLimit(v int32)`

SetRequestedCreditLimit sets RequestedCreditLimit field to given value.

### HasRequestedCreditLimit

`func (o *Application) HasRequestedCreditLimit() bool`

HasRequestedCreditLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


