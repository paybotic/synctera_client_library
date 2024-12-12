# PostPerson

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
**IsCustomer** | **bool** | True for personal and business customers with a direct relationship with the fintech or bank. Set this to true for any customer related to an account.  | 
**LastName** | Pointer to **string** | Person&#39;s last name. | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**LegalAddress** | Pointer to [**LegalAddress**](LegalAddress.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**MiddleName** | Pointer to **string** | Person&#39;s middle name. | [optional] 
**PhoneNumber** | Pointer to **string** | Person&#39;s mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated | [optional] 
**ShippingAddress** | Pointer to [**ShippingAddress**](ShippingAddress.md) |  | [optional] 
**Ssn** | Pointer to **string** | Person&#39;s full tax ID eg SSN formatted with hyphens. The response contains the last 4 digits only (e.g. 6789). | [optional] 
**SsnSource** | Pointer to [**SsnSource**](SsnSource.md) |  | [optional] 
**Status** | [**PersonStatus**](PersonStatus.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**VerificationLastRun** | Pointer to **time.Time** | Date and time KYC verification was last run on the person. | [optional] [readonly] 
**VerificationStatus** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] 
**PersonalIds** | Pointer to [**[]PostPersonalId**](PostPersonalId.md) | Array of personal identifiers  | [optional] 
**Note** | Pointer to **string** | Text to be added to a note when creating a person. A note is required when creating a person with a ban_status of SUSPENDED. | [optional] 

## Methods

### NewPostPerson

`func NewPostPerson(isCustomer bool, status PersonStatus, ) *PostPerson`

NewPostPerson instantiates a new PostPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPersonWithDefaults

`func NewPostPersonWithDefaults() *PostPerson`

NewPostPersonWithDefaults instantiates a new PostPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanStatus

`func (o *PostPerson) GetBanStatus() BanStatus`

GetBanStatus returns the BanStatus field if non-nil, zero value otherwise.

### GetBanStatusOk

`func (o *PostPerson) GetBanStatusOk() (*BanStatus, bool)`

GetBanStatusOk returns a tuple with the BanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStatus

`func (o *PostPerson) SetBanStatus(v BanStatus)`

SetBanStatus sets BanStatus field to given value.

### HasBanStatus

`func (o *PostPerson) HasBanStatus() bool`

HasBanStatus returns a boolean if a field has been set.

### GetChosenName

`func (o *PostPerson) GetChosenName() string`

GetChosenName returns the ChosenName field if non-nil, zero value otherwise.

### GetChosenNameOk

`func (o *PostPerson) GetChosenNameOk() (*string, bool)`

GetChosenNameOk returns a tuple with the ChosenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenName

`func (o *PostPerson) SetChosenName(v string)`

SetChosenName sets ChosenName field to given value.

### HasChosenName

`func (o *PostPerson) HasChosenName() bool`

HasChosenName returns a boolean if a field has been set.

### GetCreationTime

`func (o *PostPerson) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PostPerson) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PostPerson) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PostPerson) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDob

`func (o *PostPerson) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *PostPerson) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *PostPerson) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *PostPerson) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetEmail

`func (o *PostPerson) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostPerson) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostPerson) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PostPerson) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *PostPerson) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PostPerson) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PostPerson) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PostPerson) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHasAccounts

`func (o *PostPerson) GetHasAccounts() bool`

GetHasAccounts returns the HasAccounts field if non-nil, zero value otherwise.

### GetHasAccountsOk

`func (o *PostPerson) GetHasAccountsOk() (*bool, bool)`

GetHasAccountsOk returns a tuple with the HasAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccounts

`func (o *PostPerson) SetHasAccounts(v bool)`

SetHasAccounts sets HasAccounts field to given value.

### HasHasAccounts

`func (o *PostPerson) HasHasAccounts() bool`

HasHasAccounts returns a boolean if a field has been set.

### GetId

`func (o *PostPerson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostPerson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostPerson) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostPerson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCustomer

`func (o *PostPerson) GetIsCustomer() bool`

GetIsCustomer returns the IsCustomer field if non-nil, zero value otherwise.

### GetIsCustomerOk

`func (o *PostPerson) GetIsCustomerOk() (*bool, bool)`

GetIsCustomerOk returns a tuple with the IsCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomer

`func (o *PostPerson) SetIsCustomer(v bool)`

SetIsCustomer sets IsCustomer field to given value.


### GetLastName

`func (o *PostPerson) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PostPerson) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PostPerson) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PostPerson) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *PostPerson) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *PostPerson) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *PostPerson) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *PostPerson) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLegalAddress

`func (o *PostPerson) GetLegalAddress() LegalAddress`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *PostPerson) GetLegalAddressOk() (*LegalAddress, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *PostPerson) SetLegalAddress(v LegalAddress)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *PostPerson) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *PostPerson) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostPerson) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostPerson) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PostPerson) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *PostPerson) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PostPerson) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PostPerson) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PostPerson) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PostPerson) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PostPerson) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PostPerson) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PostPerson) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetShippingAddress

`func (o *PostPerson) GetShippingAddress() ShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *PostPerson) GetShippingAddressOk() (*ShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *PostPerson) SetShippingAddress(v ShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *PostPerson) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSsn

`func (o *PostPerson) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *PostPerson) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *PostPerson) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *PostPerson) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetSsnSource

`func (o *PostPerson) GetSsnSource() SsnSource`

GetSsnSource returns the SsnSource field if non-nil, zero value otherwise.

### GetSsnSourceOk

`func (o *PostPerson) GetSsnSourceOk() (*SsnSource, bool)`

GetSsnSourceOk returns a tuple with the SsnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnSource

`func (o *PostPerson) SetSsnSource(v SsnSource)`

SetSsnSource sets SsnSource field to given value.

### HasSsnSource

`func (o *PostPerson) HasSsnSource() bool`

HasSsnSource returns a boolean if a field has been set.

### GetStatus

`func (o *PostPerson) GetStatus() PersonStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostPerson) GetStatusOk() (*PersonStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostPerson) SetStatus(v PersonStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *PostPerson) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PostPerson) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PostPerson) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PostPerson) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetVerificationLastRun

`func (o *PostPerson) GetVerificationLastRun() time.Time`

GetVerificationLastRun returns the VerificationLastRun field if non-nil, zero value otherwise.

### GetVerificationLastRunOk

`func (o *PostPerson) GetVerificationLastRunOk() (*time.Time, bool)`

GetVerificationLastRunOk returns a tuple with the VerificationLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLastRun

`func (o *PostPerson) SetVerificationLastRun(v time.Time)`

SetVerificationLastRun sets VerificationLastRun field to given value.

### HasVerificationLastRun

`func (o *PostPerson) HasVerificationLastRun() bool`

HasVerificationLastRun returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *PostPerson) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *PostPerson) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *PostPerson) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *PostPerson) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetPersonalIds

`func (o *PostPerson) GetPersonalIds() []PostPersonalId`

GetPersonalIds returns the PersonalIds field if non-nil, zero value otherwise.

### GetPersonalIdsOk

`func (o *PostPerson) GetPersonalIdsOk() (*[]PostPersonalId, bool)`

GetPersonalIdsOk returns a tuple with the PersonalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIds

`func (o *PostPerson) SetPersonalIds(v []PostPersonalId)`

SetPersonalIds sets PersonalIds field to given value.

### HasPersonalIds

`func (o *PostPerson) HasPersonalIds() bool`

HasPersonalIds returns a boolean if a field has been set.

### GetNote

`func (o *PostPerson) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PostPerson) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PostPerson) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PostPerson) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


