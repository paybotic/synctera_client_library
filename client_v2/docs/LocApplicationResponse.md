# LocApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountTemplateId** | **string** | Mainapi account template ID to use for creating the LOC account | 
**Applicant** | [**LocApplicant**](LocApplicant.md) |  | 
**Attributes** | Pointer to **map[string]interface{}** | Raw JSON key-value pairs for flexible data | [optional] 
**Metadata** | Pointer to [**LocApplicationMetadata**](LocApplicationMetadata.md) |  | [optional] 
**Provider** | [**LocProvider**](LocProvider.md) |  | 
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
**CustomerType** | [**LocApplicantType**](LocApplicantType.md) |  | 
**Decision** | Pointer to [**LocDecision**](LocDecision.md) |  | [optional] 
**DecisionTime** | Pointer to **time.Time** | Decision timestamp in RFC3339 format | [optional] 
**ExternalId** | Pointer to **string** | External identifier (e.g., Taktile decision_id) | [optional] 
**Id** | **string** | Application ID | 
**InterestRate** | Pointer to **int32** | Interest rate in basis points (e.g., 1250 &#x3D; 12.50%) | [optional] 
**LastUpdatedTime** | **time.Time** | Timestamp of the last application modification in RFC3339 format | 
**PersonId** | Pointer to **string** | Person ID if customer_type is personal | [optional] 
**RejectionReasons** | Pointer to **[]string** |  | [optional] 
**Status** | [**LocApplicationStatus**](LocApplicationStatus.md) |  | 
**VendorResponse** | Pointer to [**LocVendorResponse**](LocVendorResponse.md) |  | [optional] 

## Methods

### NewLocApplicationResponse

`func NewLocApplicationResponse(accountTemplateId string, applicant LocApplicant, provider LocProvider, tenant string, type_ ApplicationType, applicationWorkflowId string, creationTime time.Time, customerType LocApplicantType, id string, lastUpdatedTime time.Time, status LocApplicationStatus, ) *LocApplicationResponse`

NewLocApplicationResponse instantiates a new LocApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocApplicationResponseWithDefaults

`func NewLocApplicationResponseWithDefaults() *LocApplicationResponse`

NewLocApplicationResponseWithDefaults instantiates a new LocApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountTemplateId

`func (o *LocApplicationResponse) GetAccountTemplateId() string`

GetAccountTemplateId returns the AccountTemplateId field if non-nil, zero value otherwise.

### GetAccountTemplateIdOk

`func (o *LocApplicationResponse) GetAccountTemplateIdOk() (*string, bool)`

GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateId

`func (o *LocApplicationResponse) SetAccountTemplateId(v string)`

SetAccountTemplateId sets AccountTemplateId field to given value.


### GetApplicant

`func (o *LocApplicationResponse) GetApplicant() LocApplicant`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *LocApplicationResponse) GetApplicantOk() (*LocApplicant, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *LocApplicationResponse) SetApplicant(v LocApplicant)`

SetApplicant sets Applicant field to given value.


### GetAttributes

`func (o *LocApplicationResponse) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LocApplicationResponse) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LocApplicationResponse) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LocApplicationResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMetadata

`func (o *LocApplicationResponse) GetMetadata() LocApplicationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LocApplicationResponse) GetMetadataOk() (*LocApplicationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LocApplicationResponse) SetMetadata(v LocApplicationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LocApplicationResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProvider

`func (o *LocApplicationResponse) GetProvider() LocProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LocApplicationResponse) GetProviderOk() (*LocProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LocApplicationResponse) SetProvider(v LocProvider)`

SetProvider sets Provider field to given value.


### GetRequestedCreditLimit

`func (o *LocApplicationResponse) GetRequestedCreditLimit() int32`

GetRequestedCreditLimit returns the RequestedCreditLimit field if non-nil, zero value otherwise.

### GetRequestedCreditLimitOk

`func (o *LocApplicationResponse) GetRequestedCreditLimitOk() (*int32, bool)`

GetRequestedCreditLimitOk returns a tuple with the RequestedCreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCreditLimit

`func (o *LocApplicationResponse) SetRequestedCreditLimit(v int32)`

SetRequestedCreditLimit sets RequestedCreditLimit field to given value.

### HasRequestedCreditLimit

`func (o *LocApplicationResponse) HasRequestedCreditLimit() bool`

HasRequestedCreditLimit returns a boolean if a field has been set.

### GetTenant

`func (o *LocApplicationResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *LocApplicationResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *LocApplicationResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetType

`func (o *LocApplicationResponse) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocApplicationResponse) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocApplicationResponse) SetType(v ApplicationType)`

SetType sets Type field to given value.


### GetAccountId

`func (o *LocApplicationResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LocApplicationResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LocApplicationResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LocApplicationResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetApplicationSubmittedTime

`func (o *LocApplicationResponse) GetApplicationSubmittedTime() time.Time`

GetApplicationSubmittedTime returns the ApplicationSubmittedTime field if non-nil, zero value otherwise.

### GetApplicationSubmittedTimeOk

`func (o *LocApplicationResponse) GetApplicationSubmittedTimeOk() (*time.Time, bool)`

GetApplicationSubmittedTimeOk returns a tuple with the ApplicationSubmittedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSubmittedTime

`func (o *LocApplicationResponse) SetApplicationSubmittedTime(v time.Time)`

SetApplicationSubmittedTime sets ApplicationSubmittedTime field to given value.

### HasApplicationSubmittedTime

`func (o *LocApplicationResponse) HasApplicationSubmittedTime() bool`

HasApplicationSubmittedTime returns a boolean if a field has been set.

### GetApplicationWorkflowId

`func (o *LocApplicationResponse) GetApplicationWorkflowId() string`

GetApplicationWorkflowId returns the ApplicationWorkflowId field if non-nil, zero value otherwise.

### GetApplicationWorkflowIdOk

`func (o *LocApplicationResponse) GetApplicationWorkflowIdOk() (*string, bool)`

GetApplicationWorkflowIdOk returns a tuple with the ApplicationWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationWorkflowId

`func (o *LocApplicationResponse) SetApplicationWorkflowId(v string)`

SetApplicationWorkflowId sets ApplicationWorkflowId field to given value.


### GetBusinessId

`func (o *LocApplicationResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *LocApplicationResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *LocApplicationResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *LocApplicationResponse) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *LocApplicationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *LocApplicationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *LocApplicationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCreditLimit

`func (o *LocApplicationResponse) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *LocApplicationResponse) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *LocApplicationResponse) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *LocApplicationResponse) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCreditReport

`func (o *LocApplicationResponse) GetCreditReport() map[string]interface{}`

GetCreditReport returns the CreditReport field if non-nil, zero value otherwise.

### GetCreditReportOk

`func (o *LocApplicationResponse) GetCreditReportOk() (*map[string]interface{}, bool)`

GetCreditReportOk returns a tuple with the CreditReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditReport

`func (o *LocApplicationResponse) SetCreditReport(v map[string]interface{})`

SetCreditReport sets CreditReport field to given value.

### HasCreditReport

`func (o *LocApplicationResponse) HasCreditReport() bool`

HasCreditReport returns a boolean if a field has been set.

### GetCustomerType

`func (o *LocApplicationResponse) GetCustomerType() LocApplicantType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *LocApplicationResponse) GetCustomerTypeOk() (*LocApplicantType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *LocApplicationResponse) SetCustomerType(v LocApplicantType)`

SetCustomerType sets CustomerType field to given value.


### GetDecision

`func (o *LocApplicationResponse) GetDecision() LocDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *LocApplicationResponse) GetDecisionOk() (*LocDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *LocApplicationResponse) SetDecision(v LocDecision)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *LocApplicationResponse) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetDecisionTime

`func (o *LocApplicationResponse) GetDecisionTime() time.Time`

GetDecisionTime returns the DecisionTime field if non-nil, zero value otherwise.

### GetDecisionTimeOk

`func (o *LocApplicationResponse) GetDecisionTimeOk() (*time.Time, bool)`

GetDecisionTimeOk returns a tuple with the DecisionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionTime

`func (o *LocApplicationResponse) SetDecisionTime(v time.Time)`

SetDecisionTime sets DecisionTime field to given value.

### HasDecisionTime

`func (o *LocApplicationResponse) HasDecisionTime() bool`

HasDecisionTime returns a boolean if a field has been set.

### GetExternalId

`func (o *LocApplicationResponse) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *LocApplicationResponse) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *LocApplicationResponse) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *LocApplicationResponse) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *LocApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInterestRate

`func (o *LocApplicationResponse) GetInterestRate() int32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *LocApplicationResponse) GetInterestRateOk() (*int32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *LocApplicationResponse) SetInterestRate(v int32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *LocApplicationResponse) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *LocApplicationResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *LocApplicationResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *LocApplicationResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetPersonId

`func (o *LocApplicationResponse) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *LocApplicationResponse) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *LocApplicationResponse) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *LocApplicationResponse) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetRejectionReasons

`func (o *LocApplicationResponse) GetRejectionReasons() []string`

GetRejectionReasons returns the RejectionReasons field if non-nil, zero value otherwise.

### GetRejectionReasonsOk

`func (o *LocApplicationResponse) GetRejectionReasonsOk() (*[]string, bool)`

GetRejectionReasonsOk returns a tuple with the RejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReasons

`func (o *LocApplicationResponse) SetRejectionReasons(v []string)`

SetRejectionReasons sets RejectionReasons field to given value.

### HasRejectionReasons

`func (o *LocApplicationResponse) HasRejectionReasons() bool`

HasRejectionReasons returns a boolean if a field has been set.

### GetStatus

`func (o *LocApplicationResponse) GetStatus() LocApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LocApplicationResponse) GetStatusOk() (*LocApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LocApplicationResponse) SetStatus(v LocApplicationStatus)`

SetStatus sets Status field to given value.


### GetVendorResponse

`func (o *LocApplicationResponse) GetVendorResponse() LocVendorResponse`

GetVendorResponse returns the VendorResponse field if non-nil, zero value otherwise.

### GetVendorResponseOk

`func (o *LocApplicationResponse) GetVendorResponseOk() (*LocVendorResponse, bool)`

GetVendorResponseOk returns a tuple with the VendorResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorResponse

`func (o *LocApplicationResponse) SetVendorResponse(v LocVendorResponse)`

SetVendorResponse sets VendorResponse field to given value.

### HasVendorResponse

`func (o *LocApplicationResponse) HasVendorResponse() bool`

HasVendorResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


