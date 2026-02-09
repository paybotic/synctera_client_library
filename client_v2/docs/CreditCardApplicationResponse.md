# CreditCardApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountTemplateId** | **string** | Mainapi account template ID to use for creating the credit card account | 
**Applicant** | [**CreditCardApplicant**](CreditCardApplicant.md) |  | 
**Attributes** | Pointer to **map[string]interface{}** | Raw JSON key-value pairs for flexible data | [optional] 
**Metadata** | Pointer to [**CreditCardApplicationMetadata**](CreditCardApplicationMetadata.md) |  | [optional] 
**Provider** | [**CreditCardProvider**](CreditCardProvider.md) |  | 
**RequestedCreditLimit** | Pointer to **int32** | Requested credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 
**AccountId** | Pointer to **string** | Account ID created for approved application | [optional] 
**ApplicationSubmittedTime** | Pointer to **time.Time** | Application submitted timestamp in RFC3339 format | [optional] 
**ApplicationWorkflowId** | **string** | ID of the workflow used for processing. Obtained from the account template. | 
**BusinessId** | Pointer to **string** | Business ID if customer_type is business | [optional] 
**CreationTime** | **time.Time** | Application creation timestamp in RFC3339 format | 
**CreditLimit** | Pointer to **int32** | Credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 
**CreditReport** | Pointer to **map[string]interface{}** | Credit report data object | [optional] 
**CustomerType** | [**CreditCardApplicantType**](CreditCardApplicantType.md) |  | 
**Decision** | Pointer to [**CreditCardDecision**](CreditCardDecision.md) |  | [optional] 
**DecisionTime** | Pointer to **time.Time** | Decision timestamp in RFC3339 format | [optional] 
**ExternalId** | Pointer to **string** | External identifier (e.g., Taktile decision_id) | [optional] 
**Id** | **string** | Application ID | 
**InterestRate** | Pointer to **int32** | Interest rate in basis points (e.g., 1250 &#x3D; 12.50%) | [optional] 
**LastUpdatedTime** | **time.Time** | Timestamp of the last application modification in RFC3339 format | 
**PersonId** | Pointer to **string** | Person ID if customer_type is personal | [optional] 
**RejectionReasons** | Pointer to **[]string** |  | [optional] 
**Status** | [**CreditCardApplicationStatus**](CreditCardApplicationStatus.md) |  | 
**VendorResponse** | Pointer to [**CreditCardVendorResponse**](CreditCardVendorResponse.md) |  | [optional] 

## Methods

### NewCreditCardApplicationResponse

`func NewCreditCardApplicationResponse(accountTemplateId string, applicant CreditCardApplicant, provider CreditCardProvider, tenant string, type_ ApplicationType, applicationWorkflowId string, creationTime time.Time, customerType CreditCardApplicantType, id string, lastUpdatedTime time.Time, status CreditCardApplicationStatus, ) *CreditCardApplicationResponse`

NewCreditCardApplicationResponse instantiates a new CreditCardApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardApplicationResponseWithDefaults

`func NewCreditCardApplicationResponseWithDefaults() *CreditCardApplicationResponse`

NewCreditCardApplicationResponseWithDefaults instantiates a new CreditCardApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountTemplateId

`func (o *CreditCardApplicationResponse) GetAccountTemplateId() string`

GetAccountTemplateId returns the AccountTemplateId field if non-nil, zero value otherwise.

### GetAccountTemplateIdOk

`func (o *CreditCardApplicationResponse) GetAccountTemplateIdOk() (*string, bool)`

GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateId

`func (o *CreditCardApplicationResponse) SetAccountTemplateId(v string)`

SetAccountTemplateId sets AccountTemplateId field to given value.


### GetApplicant

`func (o *CreditCardApplicationResponse) GetApplicant() CreditCardApplicant`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *CreditCardApplicationResponse) GetApplicantOk() (*CreditCardApplicant, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *CreditCardApplicationResponse) SetApplicant(v CreditCardApplicant)`

SetApplicant sets Applicant field to given value.


### GetAttributes

`func (o *CreditCardApplicationResponse) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreditCardApplicationResponse) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreditCardApplicationResponse) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreditCardApplicationResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMetadata

`func (o *CreditCardApplicationResponse) GetMetadata() CreditCardApplicationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreditCardApplicationResponse) GetMetadataOk() (*CreditCardApplicationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreditCardApplicationResponse) SetMetadata(v CreditCardApplicationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreditCardApplicationResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProvider

`func (o *CreditCardApplicationResponse) GetProvider() CreditCardProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreditCardApplicationResponse) GetProviderOk() (*CreditCardProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreditCardApplicationResponse) SetProvider(v CreditCardProvider)`

SetProvider sets Provider field to given value.


### GetRequestedCreditLimit

`func (o *CreditCardApplicationResponse) GetRequestedCreditLimit() int32`

GetRequestedCreditLimit returns the RequestedCreditLimit field if non-nil, zero value otherwise.

### GetRequestedCreditLimitOk

`func (o *CreditCardApplicationResponse) GetRequestedCreditLimitOk() (*int32, bool)`

GetRequestedCreditLimitOk returns a tuple with the RequestedCreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCreditLimit

`func (o *CreditCardApplicationResponse) SetRequestedCreditLimit(v int32)`

SetRequestedCreditLimit sets RequestedCreditLimit field to given value.

### HasRequestedCreditLimit

`func (o *CreditCardApplicationResponse) HasRequestedCreditLimit() bool`

HasRequestedCreditLimit returns a boolean if a field has been set.

### GetTenant

`func (o *CreditCardApplicationResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreditCardApplicationResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreditCardApplicationResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetType

`func (o *CreditCardApplicationResponse) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditCardApplicationResponse) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditCardApplicationResponse) SetType(v ApplicationType)`

SetType sets Type field to given value.


### GetAccountId

`func (o *CreditCardApplicationResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreditCardApplicationResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreditCardApplicationResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreditCardApplicationResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetApplicationSubmittedTime

`func (o *CreditCardApplicationResponse) GetApplicationSubmittedTime() time.Time`

GetApplicationSubmittedTime returns the ApplicationSubmittedTime field if non-nil, zero value otherwise.

### GetApplicationSubmittedTimeOk

`func (o *CreditCardApplicationResponse) GetApplicationSubmittedTimeOk() (*time.Time, bool)`

GetApplicationSubmittedTimeOk returns a tuple with the ApplicationSubmittedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSubmittedTime

`func (o *CreditCardApplicationResponse) SetApplicationSubmittedTime(v time.Time)`

SetApplicationSubmittedTime sets ApplicationSubmittedTime field to given value.

### HasApplicationSubmittedTime

`func (o *CreditCardApplicationResponse) HasApplicationSubmittedTime() bool`

HasApplicationSubmittedTime returns a boolean if a field has been set.

### GetApplicationWorkflowId

`func (o *CreditCardApplicationResponse) GetApplicationWorkflowId() string`

GetApplicationWorkflowId returns the ApplicationWorkflowId field if non-nil, zero value otherwise.

### GetApplicationWorkflowIdOk

`func (o *CreditCardApplicationResponse) GetApplicationWorkflowIdOk() (*string, bool)`

GetApplicationWorkflowIdOk returns a tuple with the ApplicationWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationWorkflowId

`func (o *CreditCardApplicationResponse) SetApplicationWorkflowId(v string)`

SetApplicationWorkflowId sets ApplicationWorkflowId field to given value.


### GetBusinessId

`func (o *CreditCardApplicationResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *CreditCardApplicationResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *CreditCardApplicationResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *CreditCardApplicationResponse) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *CreditCardApplicationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreditCardApplicationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreditCardApplicationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCreditLimit

`func (o *CreditCardApplicationResponse) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *CreditCardApplicationResponse) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *CreditCardApplicationResponse) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *CreditCardApplicationResponse) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCreditReport

`func (o *CreditCardApplicationResponse) GetCreditReport() map[string]interface{}`

GetCreditReport returns the CreditReport field if non-nil, zero value otherwise.

### GetCreditReportOk

`func (o *CreditCardApplicationResponse) GetCreditReportOk() (*map[string]interface{}, bool)`

GetCreditReportOk returns a tuple with the CreditReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditReport

`func (o *CreditCardApplicationResponse) SetCreditReport(v map[string]interface{})`

SetCreditReport sets CreditReport field to given value.

### HasCreditReport

`func (o *CreditCardApplicationResponse) HasCreditReport() bool`

HasCreditReport returns a boolean if a field has been set.

### GetCustomerType

`func (o *CreditCardApplicationResponse) GetCustomerType() CreditCardApplicantType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *CreditCardApplicationResponse) GetCustomerTypeOk() (*CreditCardApplicantType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *CreditCardApplicationResponse) SetCustomerType(v CreditCardApplicantType)`

SetCustomerType sets CustomerType field to given value.


### GetDecision

`func (o *CreditCardApplicationResponse) GetDecision() CreditCardDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *CreditCardApplicationResponse) GetDecisionOk() (*CreditCardDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *CreditCardApplicationResponse) SetDecision(v CreditCardDecision)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *CreditCardApplicationResponse) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetDecisionTime

`func (o *CreditCardApplicationResponse) GetDecisionTime() time.Time`

GetDecisionTime returns the DecisionTime field if non-nil, zero value otherwise.

### GetDecisionTimeOk

`func (o *CreditCardApplicationResponse) GetDecisionTimeOk() (*time.Time, bool)`

GetDecisionTimeOk returns a tuple with the DecisionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionTime

`func (o *CreditCardApplicationResponse) SetDecisionTime(v time.Time)`

SetDecisionTime sets DecisionTime field to given value.

### HasDecisionTime

`func (o *CreditCardApplicationResponse) HasDecisionTime() bool`

HasDecisionTime returns a boolean if a field has been set.

### GetExternalId

`func (o *CreditCardApplicationResponse) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreditCardApplicationResponse) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreditCardApplicationResponse) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreditCardApplicationResponse) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *CreditCardApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCardApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCardApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInterestRate

`func (o *CreditCardApplicationResponse) GetInterestRate() int32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *CreditCardApplicationResponse) GetInterestRateOk() (*int32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *CreditCardApplicationResponse) SetInterestRate(v int32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *CreditCardApplicationResponse) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *CreditCardApplicationResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CreditCardApplicationResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CreditCardApplicationResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetPersonId

`func (o *CreditCardApplicationResponse) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *CreditCardApplicationResponse) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *CreditCardApplicationResponse) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *CreditCardApplicationResponse) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetRejectionReasons

`func (o *CreditCardApplicationResponse) GetRejectionReasons() []string`

GetRejectionReasons returns the RejectionReasons field if non-nil, zero value otherwise.

### GetRejectionReasonsOk

`func (o *CreditCardApplicationResponse) GetRejectionReasonsOk() (*[]string, bool)`

GetRejectionReasonsOk returns a tuple with the RejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReasons

`func (o *CreditCardApplicationResponse) SetRejectionReasons(v []string)`

SetRejectionReasons sets RejectionReasons field to given value.

### HasRejectionReasons

`func (o *CreditCardApplicationResponse) HasRejectionReasons() bool`

HasRejectionReasons returns a boolean if a field has been set.

### GetStatus

`func (o *CreditCardApplicationResponse) GetStatus() CreditCardApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditCardApplicationResponse) GetStatusOk() (*CreditCardApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditCardApplicationResponse) SetStatus(v CreditCardApplicationStatus)`

SetStatus sets Status field to given value.


### GetVendorResponse

`func (o *CreditCardApplicationResponse) GetVendorResponse() CreditCardVendorResponse`

GetVendorResponse returns the VendorResponse field if non-nil, zero value otherwise.

### GetVendorResponseOk

`func (o *CreditCardApplicationResponse) GetVendorResponseOk() (*CreditCardVendorResponse, bool)`

GetVendorResponseOk returns a tuple with the VendorResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorResponse

`func (o *CreditCardApplicationResponse) SetVendorResponse(v CreditCardVendorResponse)`

SetVendorResponse sets VendorResponse field to given value.

### HasVendorResponse

`func (o *CreditCardApplicationResponse) HasVendorResponse() bool`

HasVendorResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


