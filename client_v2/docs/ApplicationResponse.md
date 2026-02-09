# ApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Account ID created for approved application | [optional] 
**AccountType** | [**ApplicationAccountType**](ApplicationAccountType.md) |  | 
**Applicants** | [**[]Applicant**](Applicant.md) |  | 
**ApplicationSubmittedTime** | Pointer to **time.Time** | Application submitted timestamp in RFC3339 format | [optional] 
**CreditDecisionTime** | Pointer to **time.Time** | Credit decision timestamp in RFC3339 format | [optional] 
**CustomerResponseTime** | Pointer to **time.Time** | Credit decision timestamp in RFC3339 format | [optional] 
**Purpose** | [**CreditApplicationPurpose**](CreditApplicationPurpose.md) |  | 
**Status** | [**CreditCardApplicationStatus**](CreditCardApplicationStatus.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 
**CreationTime** | **time.Time** | Application creation timestamp in RFC3339 format | 
**Id** | **string** | Application ID | 
**LastUpdatedTime** | **time.Time** | Timestamp of the last application modification in RFC3339 format | 
**ApplicationDetails** | [**RestrictedApplicationDetails**](RestrictedApplicationDetails.md) |  | 
**BusinessId** | Pointer to **string** | Business ID if customer_type is business | [optional] 
**CustomerId** | Pointer to **string** | Customer ID for the application. An application must have either a Business or a customer associated with it. | [optional] 
**Description** | Pointer to **string** | A description of the restricted account application | [optional] 
**AccountTemplateId** | **string** | Mainapi account template ID to use for creating the credit card account | 
**Applicant** | [**CreditCardApplicant**](CreditCardApplicant.md) |  | 
**Attributes** | Pointer to **map[string]interface{}** | Raw JSON key-value pairs for flexible data | [optional] 
**Metadata** | Pointer to [**CreditCardApplicationMetadata**](CreditCardApplicationMetadata.md) |  | [optional] 
**Provider** | [**CreditCardProvider**](CreditCardProvider.md) |  | 
**RequestedCreditLimit** | Pointer to **int32** | Requested credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 
**ApplicationWorkflowId** | **string** | ID of the workflow used for processing. Obtained from the account template. | 
**CreditLimit** | Pointer to **int32** | Credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 
**CreditReport** | Pointer to **map[string]interface{}** | Credit report data object | [optional] 
**CustomerType** | [**CreditCardApplicantType**](CreditCardApplicantType.md) |  | 
**Decision** | Pointer to [**CreditCardDecision**](CreditCardDecision.md) |  | [optional] 
**DecisionTime** | Pointer to **time.Time** | Decision timestamp in RFC3339 format | [optional] 
**ExternalId** | Pointer to **string** | External identifier (e.g., Taktile decision_id) | [optional] 
**InterestRate** | Pointer to **int32** | Interest rate in basis points (e.g., 1250 &#x3D; 12.50%) | [optional] 
**PersonId** | Pointer to **string** | Person ID if customer_type is personal | [optional] 
**RejectionReasons** | Pointer to **[]string** |  | [optional] 
**VendorResponse** | Pointer to [**CreditCardVendorResponse**](CreditCardVendorResponse.md) |  | [optional] 

## Methods

### NewApplicationResponse

`func NewApplicationResponse(accountType ApplicationAccountType, applicants []Applicant, purpose CreditApplicationPurpose, status CreditCardApplicationStatus, tenant string, type_ ApplicationType, creationTime time.Time, id string, lastUpdatedTime time.Time, applicationDetails RestrictedApplicationDetails, accountTemplateId string, applicant CreditCardApplicant, provider CreditCardProvider, applicationWorkflowId string, customerType CreditCardApplicantType, ) *ApplicationResponse`

NewApplicationResponse instantiates a new ApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseWithDefaults

`func NewApplicationResponseWithDefaults() *ApplicationResponse`

NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ApplicationResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ApplicationResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ApplicationResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ApplicationResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountType

`func (o *ApplicationResponse) GetAccountType() ApplicationAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ApplicationResponse) GetAccountTypeOk() (*ApplicationAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ApplicationResponse) SetAccountType(v ApplicationAccountType)`

SetAccountType sets AccountType field to given value.


### GetApplicants

`func (o *ApplicationResponse) GetApplicants() []Applicant`

GetApplicants returns the Applicants field if non-nil, zero value otherwise.

### GetApplicantsOk

`func (o *ApplicationResponse) GetApplicantsOk() (*[]Applicant, bool)`

GetApplicantsOk returns a tuple with the Applicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicants

`func (o *ApplicationResponse) SetApplicants(v []Applicant)`

SetApplicants sets Applicants field to given value.


### GetApplicationSubmittedTime

`func (o *ApplicationResponse) GetApplicationSubmittedTime() time.Time`

GetApplicationSubmittedTime returns the ApplicationSubmittedTime field if non-nil, zero value otherwise.

### GetApplicationSubmittedTimeOk

`func (o *ApplicationResponse) GetApplicationSubmittedTimeOk() (*time.Time, bool)`

GetApplicationSubmittedTimeOk returns a tuple with the ApplicationSubmittedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSubmittedTime

`func (o *ApplicationResponse) SetApplicationSubmittedTime(v time.Time)`

SetApplicationSubmittedTime sets ApplicationSubmittedTime field to given value.

### HasApplicationSubmittedTime

`func (o *ApplicationResponse) HasApplicationSubmittedTime() bool`

HasApplicationSubmittedTime returns a boolean if a field has been set.

### GetCreditDecisionTime

`func (o *ApplicationResponse) GetCreditDecisionTime() time.Time`

GetCreditDecisionTime returns the CreditDecisionTime field if non-nil, zero value otherwise.

### GetCreditDecisionTimeOk

`func (o *ApplicationResponse) GetCreditDecisionTimeOk() (*time.Time, bool)`

GetCreditDecisionTimeOk returns a tuple with the CreditDecisionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDecisionTime

`func (o *ApplicationResponse) SetCreditDecisionTime(v time.Time)`

SetCreditDecisionTime sets CreditDecisionTime field to given value.

### HasCreditDecisionTime

`func (o *ApplicationResponse) HasCreditDecisionTime() bool`

HasCreditDecisionTime returns a boolean if a field has been set.

### GetCustomerResponseTime

`func (o *ApplicationResponse) GetCustomerResponseTime() time.Time`

GetCustomerResponseTime returns the CustomerResponseTime field if non-nil, zero value otherwise.

### GetCustomerResponseTimeOk

`func (o *ApplicationResponse) GetCustomerResponseTimeOk() (*time.Time, bool)`

GetCustomerResponseTimeOk returns a tuple with the CustomerResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerResponseTime

`func (o *ApplicationResponse) SetCustomerResponseTime(v time.Time)`

SetCustomerResponseTime sets CustomerResponseTime field to given value.

### HasCustomerResponseTime

`func (o *ApplicationResponse) HasCustomerResponseTime() bool`

HasCustomerResponseTime returns a boolean if a field has been set.

### GetPurpose

`func (o *ApplicationResponse) GetPurpose() CreditApplicationPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ApplicationResponse) GetPurposeOk() (*CreditApplicationPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ApplicationResponse) SetPurpose(v CreditApplicationPurpose)`

SetPurpose sets Purpose field to given value.


### GetStatus

`func (o *ApplicationResponse) GetStatus() CreditCardApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationResponse) GetStatusOk() (*CreditCardApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationResponse) SetStatus(v CreditCardApplicationStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *ApplicationResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ApplicationResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ApplicationResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetType

`func (o *ApplicationResponse) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationResponse) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationResponse) SetType(v ApplicationType)`

SetType sets Type field to given value.


### GetCreationTime

`func (o *ApplicationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ApplicationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ApplicationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *ApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *ApplicationResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ApplicationResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ApplicationResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetApplicationDetails

`func (o *ApplicationResponse) GetApplicationDetails() RestrictedApplicationDetails`

GetApplicationDetails returns the ApplicationDetails field if non-nil, zero value otherwise.

### GetApplicationDetailsOk

`func (o *ApplicationResponse) GetApplicationDetailsOk() (*RestrictedApplicationDetails, bool)`

GetApplicationDetailsOk returns a tuple with the ApplicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDetails

`func (o *ApplicationResponse) SetApplicationDetails(v RestrictedApplicationDetails)`

SetApplicationDetails sets ApplicationDetails field to given value.


### GetBusinessId

`func (o *ApplicationResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *ApplicationResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *ApplicationResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *ApplicationResponse) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *ApplicationResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ApplicationResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ApplicationResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ApplicationResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountTemplateId

`func (o *ApplicationResponse) GetAccountTemplateId() string`

GetAccountTemplateId returns the AccountTemplateId field if non-nil, zero value otherwise.

### GetAccountTemplateIdOk

`func (o *ApplicationResponse) GetAccountTemplateIdOk() (*string, bool)`

GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateId

`func (o *ApplicationResponse) SetAccountTemplateId(v string)`

SetAccountTemplateId sets AccountTemplateId field to given value.


### GetApplicant

`func (o *ApplicationResponse) GetApplicant() CreditCardApplicant`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *ApplicationResponse) GetApplicantOk() (*CreditCardApplicant, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *ApplicationResponse) SetApplicant(v CreditCardApplicant)`

SetApplicant sets Applicant field to given value.


### GetAttributes

`func (o *ApplicationResponse) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationResponse) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationResponse) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApplicationResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationResponse) GetMetadata() CreditCardApplicationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationResponse) GetMetadataOk() (*CreditCardApplicationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationResponse) SetMetadata(v CreditCardApplicationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProvider

`func (o *ApplicationResponse) GetProvider() CreditCardProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApplicationResponse) GetProviderOk() (*CreditCardProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApplicationResponse) SetProvider(v CreditCardProvider)`

SetProvider sets Provider field to given value.


### GetRequestedCreditLimit

`func (o *ApplicationResponse) GetRequestedCreditLimit() int32`

GetRequestedCreditLimit returns the RequestedCreditLimit field if non-nil, zero value otherwise.

### GetRequestedCreditLimitOk

`func (o *ApplicationResponse) GetRequestedCreditLimitOk() (*int32, bool)`

GetRequestedCreditLimitOk returns a tuple with the RequestedCreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCreditLimit

`func (o *ApplicationResponse) SetRequestedCreditLimit(v int32)`

SetRequestedCreditLimit sets RequestedCreditLimit field to given value.

### HasRequestedCreditLimit

`func (o *ApplicationResponse) HasRequestedCreditLimit() bool`

HasRequestedCreditLimit returns a boolean if a field has been set.

### GetApplicationWorkflowId

`func (o *ApplicationResponse) GetApplicationWorkflowId() string`

GetApplicationWorkflowId returns the ApplicationWorkflowId field if non-nil, zero value otherwise.

### GetApplicationWorkflowIdOk

`func (o *ApplicationResponse) GetApplicationWorkflowIdOk() (*string, bool)`

GetApplicationWorkflowIdOk returns a tuple with the ApplicationWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationWorkflowId

`func (o *ApplicationResponse) SetApplicationWorkflowId(v string)`

SetApplicationWorkflowId sets ApplicationWorkflowId field to given value.


### GetCreditLimit

`func (o *ApplicationResponse) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *ApplicationResponse) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *ApplicationResponse) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *ApplicationResponse) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCreditReport

`func (o *ApplicationResponse) GetCreditReport() map[string]interface{}`

GetCreditReport returns the CreditReport field if non-nil, zero value otherwise.

### GetCreditReportOk

`func (o *ApplicationResponse) GetCreditReportOk() (*map[string]interface{}, bool)`

GetCreditReportOk returns a tuple with the CreditReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditReport

`func (o *ApplicationResponse) SetCreditReport(v map[string]interface{})`

SetCreditReport sets CreditReport field to given value.

### HasCreditReport

`func (o *ApplicationResponse) HasCreditReport() bool`

HasCreditReport returns a boolean if a field has been set.

### GetCustomerType

`func (o *ApplicationResponse) GetCustomerType() CreditCardApplicantType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *ApplicationResponse) GetCustomerTypeOk() (*CreditCardApplicantType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *ApplicationResponse) SetCustomerType(v CreditCardApplicantType)`

SetCustomerType sets CustomerType field to given value.


### GetDecision

`func (o *ApplicationResponse) GetDecision() CreditCardDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ApplicationResponse) GetDecisionOk() (*CreditCardDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ApplicationResponse) SetDecision(v CreditCardDecision)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *ApplicationResponse) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetDecisionTime

`func (o *ApplicationResponse) GetDecisionTime() time.Time`

GetDecisionTime returns the DecisionTime field if non-nil, zero value otherwise.

### GetDecisionTimeOk

`func (o *ApplicationResponse) GetDecisionTimeOk() (*time.Time, bool)`

GetDecisionTimeOk returns a tuple with the DecisionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionTime

`func (o *ApplicationResponse) SetDecisionTime(v time.Time)`

SetDecisionTime sets DecisionTime field to given value.

### HasDecisionTime

`func (o *ApplicationResponse) HasDecisionTime() bool`

HasDecisionTime returns a boolean if a field has been set.

### GetExternalId

`func (o *ApplicationResponse) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ApplicationResponse) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ApplicationResponse) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ApplicationResponse) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInterestRate

`func (o *ApplicationResponse) GetInterestRate() int32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *ApplicationResponse) GetInterestRateOk() (*int32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *ApplicationResponse) SetInterestRate(v int32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *ApplicationResponse) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetPersonId

`func (o *ApplicationResponse) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *ApplicationResponse) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *ApplicationResponse) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *ApplicationResponse) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetRejectionReasons

`func (o *ApplicationResponse) GetRejectionReasons() []string`

GetRejectionReasons returns the RejectionReasons field if non-nil, zero value otherwise.

### GetRejectionReasonsOk

`func (o *ApplicationResponse) GetRejectionReasonsOk() (*[]string, bool)`

GetRejectionReasonsOk returns a tuple with the RejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReasons

`func (o *ApplicationResponse) SetRejectionReasons(v []string)`

SetRejectionReasons sets RejectionReasons field to given value.

### HasRejectionReasons

`func (o *ApplicationResponse) HasRejectionReasons() bool`

HasRejectionReasons returns a boolean if a field has been set.

### GetVendorResponse

`func (o *ApplicationResponse) GetVendorResponse() CreditCardVendorResponse`

GetVendorResponse returns the VendorResponse field if non-nil, zero value otherwise.

### GetVendorResponseOk

`func (o *ApplicationResponse) GetVendorResponseOk() (*CreditCardVendorResponse, bool)`

GetVendorResponseOk returns a tuple with the VendorResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorResponse

`func (o *ApplicationResponse) SetVendorResponse(v CreditCardVendorResponse)`

SetVendorResponse sets VendorResponse field to given value.

### HasVendorResponse

`func (o *ApplicationResponse) HasVendorResponse() bool`

HasVendorResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


