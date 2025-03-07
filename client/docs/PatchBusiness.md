# PatchBusiness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]Address**](Address.md) | All of the customer&#39;s addresses | [optional] [readonly] 
**ComplianceRestrictions** | Pointer to **[]string** | The types of compliance that the business needs to adhere to * &#x60;LICENSED_CANNABIS&#x60; – A type of compliance restriction where the business would need a cannabis license in order to operate.  | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Ein** | Pointer to **string** | U.S. Employer Identification Number (EIN) for this business, in the format xx-xxxxxxx. | [optional] 
**Email** | Pointer to **string** | Business&#39;s email. | [optional] 
**EntityName** | Pointer to **string** | Business&#39;s legal name. | [optional] 
**FormationDate** | Pointer to **string** | Date the business was legally registered in RFC 3339 full-date format (YYYY-MM-DD). | [optional] 
**FormationState** | Pointer to **string** | U.S. state where the business is legally registered (2-letter abbreviation). | [optional] 
**HasAccounts** | Pointer to **bool** | This flag indicates whether the person or business has accounts. | [optional] [readonly] 
**Id** | Pointer to **string** | Business&#39;s unique identifier. | [optional] [readonly] 
**IsCustomer** | Pointer to **bool** | True for personal and business customers with a direct relationship with the fintech or bank. Set this to true for any customer related to an account.  | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**LegalAddress** | Pointer to [**LegalAddress**](LegalAddress.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**PhoneNumber** | Pointer to **string** | Business&#39;s phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated. | [optional] 
**Status** | Pointer to **string** | Status of the business. One of the following: * &#x60;PROSPECT&#x60; – a potential customer, used for information-gathering and disclosures. * &#x60;ACTIVE&#x60; –  is an integrator defined status.  Integrators should set a business to active if they believe the person to be qualified for conducting business.  Synctera will combine this status with other statuses such a verification to determine if the business is eligible for specific actions such as initiating transactions or issuing a card. * &#x60;FROZEN&#x60; – business&#39;s actions are blocked for security, legal, or other reasons. * &#x60;SANCTION&#x60; – business is on a sanctions list and should be carefully monitored. * &#x60;DISSOLVED&#x60; – an inactive status indicating a business entity has filed articles of dissolution or a certificate of termination to terminate its existence. * &#x60;CANCELLED&#x60; – an inactive status indicating that a business entity has filed a cancellation or has failed to file its periodic report after notice of forfeiture of its rights to do business. * &#x60;SUSPENDED&#x60; – an inactive status indicating that the business entity has lost the right to operate in it&#39;s registered jurisdiction. * &#x60;MERGED&#x60; – an inactive status indicating that the business entity has terminated existence by merging into another entity. * &#x60;INACTIVE&#x60; – an inactive status indicating that the business entity is no longer active. * &#x60;CONVERTED&#x60; – An inactive status indicating that the business entity has been converted to another type of business entity in the same jurisdiction.  | [optional] 
**Structure** | Pointer to **string** | Business&#39;s legal structure. | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**TradeNames** | Pointer to **[]string** | All registered &#39;doing business as&#39; (DBA) or trade names for this business. | [optional] 
**VerificationLastRun** | Pointer to **time.Time** | Date and time KYB verification was last run on the business. | [optional] [readonly] 
**VerificationStatus** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] 
**Website** | Pointer to **string** | Business&#39;s website. | [optional] 

## Methods

### NewPatchBusiness

`func NewPatchBusiness() *PatchBusiness`

NewPatchBusiness instantiates a new PatchBusiness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchBusinessWithDefaults

`func NewPatchBusinessWithDefaults() *PatchBusiness`

NewPatchBusinessWithDefaults instantiates a new PatchBusiness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *PatchBusiness) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PatchBusiness) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PatchBusiness) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PatchBusiness) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetComplianceRestrictions

`func (o *PatchBusiness) GetComplianceRestrictions() []string`

GetComplianceRestrictions returns the ComplianceRestrictions field if non-nil, zero value otherwise.

### GetComplianceRestrictionsOk

`func (o *PatchBusiness) GetComplianceRestrictionsOk() (*[]string, bool)`

GetComplianceRestrictionsOk returns a tuple with the ComplianceRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceRestrictions

`func (o *PatchBusiness) SetComplianceRestrictions(v []string)`

SetComplianceRestrictions sets ComplianceRestrictions field to given value.

### HasComplianceRestrictions

`func (o *PatchBusiness) HasComplianceRestrictions() bool`

HasComplianceRestrictions returns a boolean if a field has been set.

### GetCreationTime

`func (o *PatchBusiness) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PatchBusiness) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PatchBusiness) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PatchBusiness) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEin

`func (o *PatchBusiness) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *PatchBusiness) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *PatchBusiness) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *PatchBusiness) HasEin() bool`

HasEin returns a boolean if a field has been set.

### GetEmail

`func (o *PatchBusiness) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchBusiness) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchBusiness) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchBusiness) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEntityName

`func (o *PatchBusiness) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *PatchBusiness) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *PatchBusiness) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *PatchBusiness) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetFormationDate

`func (o *PatchBusiness) GetFormationDate() string`

GetFormationDate returns the FormationDate field if non-nil, zero value otherwise.

### GetFormationDateOk

`func (o *PatchBusiness) GetFormationDateOk() (*string, bool)`

GetFormationDateOk returns a tuple with the FormationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormationDate

`func (o *PatchBusiness) SetFormationDate(v string)`

SetFormationDate sets FormationDate field to given value.

### HasFormationDate

`func (o *PatchBusiness) HasFormationDate() bool`

HasFormationDate returns a boolean if a field has been set.

### GetFormationState

`func (o *PatchBusiness) GetFormationState() string`

GetFormationState returns the FormationState field if non-nil, zero value otherwise.

### GetFormationStateOk

`func (o *PatchBusiness) GetFormationStateOk() (*string, bool)`

GetFormationStateOk returns a tuple with the FormationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormationState

`func (o *PatchBusiness) SetFormationState(v string)`

SetFormationState sets FormationState field to given value.

### HasFormationState

`func (o *PatchBusiness) HasFormationState() bool`

HasFormationState returns a boolean if a field has been set.

### GetHasAccounts

`func (o *PatchBusiness) GetHasAccounts() bool`

GetHasAccounts returns the HasAccounts field if non-nil, zero value otherwise.

### GetHasAccountsOk

`func (o *PatchBusiness) GetHasAccountsOk() (*bool, bool)`

GetHasAccountsOk returns a tuple with the HasAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccounts

`func (o *PatchBusiness) SetHasAccounts(v bool)`

SetHasAccounts sets HasAccounts field to given value.

### HasHasAccounts

`func (o *PatchBusiness) HasHasAccounts() bool`

HasHasAccounts returns a boolean if a field has been set.

### GetId

`func (o *PatchBusiness) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchBusiness) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchBusiness) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchBusiness) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCustomer

`func (o *PatchBusiness) GetIsCustomer() bool`

GetIsCustomer returns the IsCustomer field if non-nil, zero value otherwise.

### GetIsCustomerOk

`func (o *PatchBusiness) GetIsCustomerOk() (*bool, bool)`

GetIsCustomerOk returns a tuple with the IsCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomer

`func (o *PatchBusiness) SetIsCustomer(v bool)`

SetIsCustomer sets IsCustomer field to given value.

### HasIsCustomer

`func (o *PatchBusiness) HasIsCustomer() bool`

HasIsCustomer returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *PatchBusiness) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *PatchBusiness) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *PatchBusiness) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *PatchBusiness) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *PatchBusiness) GetLegalAddress() LegalAddress`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *PatchBusiness) GetLegalAddressOk() (*LegalAddress, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *PatchBusiness) SetLegalAddress(v LegalAddress)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *PatchBusiness) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchBusiness) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchBusiness) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchBusiness) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchBusiness) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PatchBusiness) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PatchBusiness) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PatchBusiness) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PatchBusiness) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *PatchBusiness) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchBusiness) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchBusiness) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchBusiness) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStructure

`func (o *PatchBusiness) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *PatchBusiness) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *PatchBusiness) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *PatchBusiness) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetTenant

`func (o *PatchBusiness) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchBusiness) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchBusiness) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchBusiness) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetTradeNames

`func (o *PatchBusiness) GetTradeNames() []string`

GetTradeNames returns the TradeNames field if non-nil, zero value otherwise.

### GetTradeNamesOk

`func (o *PatchBusiness) GetTradeNamesOk() (*[]string, bool)`

GetTradeNamesOk returns a tuple with the TradeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNames

`func (o *PatchBusiness) SetTradeNames(v []string)`

SetTradeNames sets TradeNames field to given value.

### HasTradeNames

`func (o *PatchBusiness) HasTradeNames() bool`

HasTradeNames returns a boolean if a field has been set.

### GetVerificationLastRun

`func (o *PatchBusiness) GetVerificationLastRun() time.Time`

GetVerificationLastRun returns the VerificationLastRun field if non-nil, zero value otherwise.

### GetVerificationLastRunOk

`func (o *PatchBusiness) GetVerificationLastRunOk() (*time.Time, bool)`

GetVerificationLastRunOk returns a tuple with the VerificationLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLastRun

`func (o *PatchBusiness) SetVerificationLastRun(v time.Time)`

SetVerificationLastRun sets VerificationLastRun field to given value.

### HasVerificationLastRun

`func (o *PatchBusiness) HasVerificationLastRun() bool`

HasVerificationLastRun returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *PatchBusiness) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *PatchBusiness) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *PatchBusiness) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *PatchBusiness) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetWebsite

`func (o *PatchBusiness) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *PatchBusiness) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *PatchBusiness) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *PatchBusiness) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


