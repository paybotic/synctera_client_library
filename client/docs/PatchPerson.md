# PatchPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BanStatus** | Pointer to [**BanStatus**](BanStatus.md) |  | [optional] 
**ChosenName** | Pointer to **string** | Person&#39;s chosen name. | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Dob** | Pointer to **string** | Person&#39;s date of birth in RFC 3339 full-date format (YYYY-MM-DD). Must be on or after 1900-01-01 and before current date. | [optional] 
**Email** | Pointer to **string** | Person&#39;s email. | [optional] 
**FirstName** | Pointer to **string** | Person&#39;s first name. | [optional] 
**HasAccounts** | Pointer to **bool** | This flag indicates whether the person or business has accounts. | [optional] [readonly] 
**Id** | Pointer to **string** | Person&#39;s unique identifier. | [optional] [readonly] 
**IsCustomer** | Pointer to **bool** | True for personal and business customers with a direct relationship with the fintech or bank. Set this to true for any customer related to an account.  | [optional] 
**IsUser** | Pointer to **bool** | True for console users. Set this to true for any users direct access to the Synctera console.  | [optional] [readonly] 
**LastName** | Pointer to **string** | Person&#39;s last name. | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**LegalAddress** | Pointer to [**LegalAddress**](LegalAddress.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**MiddleName** | Pointer to **string** | Person&#39;s middle name. | [optional] 
**PhoneNumber** | Pointer to **string** | Person&#39;s mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated | [optional] 
**ShippingAddress** | Pointer to [**ShippingAddress**](ShippingAddress.md) |  | [optional] 
**Ssn** | Pointer to **string** | Person&#39;s full tax ID eg SSN formatted with hyphens. The response contains the last 4 digits only (e.g. 6789). | [optional] 
**SsnSource** | Pointer to [**SsnSource**](SsnSource.md) |  | [optional] 
**Status** | Pointer to [**PersonStatus**](PersonStatus.md) |  | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**VerificationLastRun** | Pointer to **time.Time** | Date and time KYC verification was last run on the person. | [optional] [readonly] 
**VerificationStatus** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] 
**Note** | Pointer to **string** | Text to be added to a note when updating a person. A note is required when changing a person&#39;s ban_status between SUSPENDED and ALLOWED. | [optional] 

## Methods

### NewPatchPerson

`func NewPatchPerson() *PatchPerson`

NewPatchPerson instantiates a new PatchPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPersonWithDefaults

`func NewPatchPersonWithDefaults() *PatchPerson`

NewPatchPersonWithDefaults instantiates a new PatchPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanStatus

`func (o *PatchPerson) GetBanStatus() BanStatus`

GetBanStatus returns the BanStatus field if non-nil, zero value otherwise.

### GetBanStatusOk

`func (o *PatchPerson) GetBanStatusOk() (*BanStatus, bool)`

GetBanStatusOk returns a tuple with the BanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStatus

`func (o *PatchPerson) SetBanStatus(v BanStatus)`

SetBanStatus sets BanStatus field to given value.

### HasBanStatus

`func (o *PatchPerson) HasBanStatus() bool`

HasBanStatus returns a boolean if a field has been set.

### GetChosenName

`func (o *PatchPerson) GetChosenName() string`

GetChosenName returns the ChosenName field if non-nil, zero value otherwise.

### GetChosenNameOk

`func (o *PatchPerson) GetChosenNameOk() (*string, bool)`

GetChosenNameOk returns a tuple with the ChosenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenName

`func (o *PatchPerson) SetChosenName(v string)`

SetChosenName sets ChosenName field to given value.

### HasChosenName

`func (o *PatchPerson) HasChosenName() bool`

HasChosenName returns a boolean if a field has been set.

### GetCreationTime

`func (o *PatchPerson) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PatchPerson) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PatchPerson) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PatchPerson) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDob

`func (o *PatchPerson) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *PatchPerson) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *PatchPerson) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *PatchPerson) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetEmail

`func (o *PatchPerson) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchPerson) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchPerson) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchPerson) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *PatchPerson) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchPerson) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchPerson) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchPerson) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHasAccounts

`func (o *PatchPerson) GetHasAccounts() bool`

GetHasAccounts returns the HasAccounts field if non-nil, zero value otherwise.

### GetHasAccountsOk

`func (o *PatchPerson) GetHasAccountsOk() (*bool, bool)`

GetHasAccountsOk returns a tuple with the HasAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccounts

`func (o *PatchPerson) SetHasAccounts(v bool)`

SetHasAccounts sets HasAccounts field to given value.

### HasHasAccounts

`func (o *PatchPerson) HasHasAccounts() bool`

HasHasAccounts returns a boolean if a field has been set.

### GetId

`func (o *PatchPerson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchPerson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchPerson) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchPerson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCustomer

`func (o *PatchPerson) GetIsCustomer() bool`

GetIsCustomer returns the IsCustomer field if non-nil, zero value otherwise.

### GetIsCustomerOk

`func (o *PatchPerson) GetIsCustomerOk() (*bool, bool)`

GetIsCustomerOk returns a tuple with the IsCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomer

`func (o *PatchPerson) SetIsCustomer(v bool)`

SetIsCustomer sets IsCustomer field to given value.

### HasIsCustomer

`func (o *PatchPerson) HasIsCustomer() bool`

HasIsCustomer returns a boolean if a field has been set.

### GetIsUser

`func (o *PatchPerson) GetIsUser() bool`

GetIsUser returns the IsUser field if non-nil, zero value otherwise.

### GetIsUserOk

`func (o *PatchPerson) GetIsUserOk() (*bool, bool)`

GetIsUserOk returns a tuple with the IsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUser

`func (o *PatchPerson) SetIsUser(v bool)`

SetIsUser sets IsUser field to given value.

### HasIsUser

`func (o *PatchPerson) HasIsUser() bool`

HasIsUser returns a boolean if a field has been set.

### GetLastName

`func (o *PatchPerson) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchPerson) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchPerson) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchPerson) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *PatchPerson) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *PatchPerson) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *PatchPerson) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *PatchPerson) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *PatchPerson) GetLegalAddress() LegalAddress`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *PatchPerson) GetLegalAddressOk() (*LegalAddress, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *PatchPerson) SetLegalAddress(v LegalAddress)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *PatchPerson) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchPerson) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchPerson) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchPerson) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchPerson) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *PatchPerson) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PatchPerson) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PatchPerson) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PatchPerson) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PatchPerson) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PatchPerson) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PatchPerson) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PatchPerson) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetShippingAddress

`func (o *PatchPerson) GetShippingAddress() ShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *PatchPerson) GetShippingAddressOk() (*ShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *PatchPerson) SetShippingAddress(v ShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *PatchPerson) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSsn

`func (o *PatchPerson) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *PatchPerson) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *PatchPerson) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *PatchPerson) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetSsnSource

`func (o *PatchPerson) GetSsnSource() SsnSource`

GetSsnSource returns the SsnSource field if non-nil, zero value otherwise.

### GetSsnSourceOk

`func (o *PatchPerson) GetSsnSourceOk() (*SsnSource, bool)`

GetSsnSourceOk returns a tuple with the SsnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnSource

`func (o *PatchPerson) SetSsnSource(v SsnSource)`

SetSsnSource sets SsnSource field to given value.

### HasSsnSource

`func (o *PatchPerson) HasSsnSource() bool`

HasSsnSource returns a boolean if a field has been set.

### GetStatus

`func (o *PatchPerson) GetStatus() PersonStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchPerson) GetStatusOk() (*PersonStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchPerson) SetStatus(v PersonStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchPerson) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *PatchPerson) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchPerson) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchPerson) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchPerson) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetVerificationLastRun

`func (o *PatchPerson) GetVerificationLastRun() time.Time`

GetVerificationLastRun returns the VerificationLastRun field if non-nil, zero value otherwise.

### GetVerificationLastRunOk

`func (o *PatchPerson) GetVerificationLastRunOk() (*time.Time, bool)`

GetVerificationLastRunOk returns a tuple with the VerificationLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLastRun

`func (o *PatchPerson) SetVerificationLastRun(v time.Time)`

SetVerificationLastRun sets VerificationLastRun field to given value.

### HasVerificationLastRun

`func (o *PatchPerson) HasVerificationLastRun() bool`

HasVerificationLastRun returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *PatchPerson) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *PatchPerson) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *PatchPerson) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *PatchPerson) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetNote

`func (o *PatchPerson) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PatchPerson) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PatchPerson) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PatchPerson) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


