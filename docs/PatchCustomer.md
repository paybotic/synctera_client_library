# PatchCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BanStatus** | Pointer to [**BanStatus**](BanStatus.md) |  | [optional] 
**Dob** | Pointer to **string** | Customer&#39;s date of birth in RFC 3339 full-date format (YYYY-MM-DD) | [optional] 
**Email** | Pointer to **string** | Customer&#39;s email | [optional] 
**FirstName** | Pointer to **string** | Customer&#39;s first name | [optional] 
**LastName** | Pointer to **string** | Customer&#39;s last name | [optional] 
**LegalAddress** | Pointer to [**LegalAddress**](LegalAddress.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | User-supplied JSON format metadata. Do not use to store PII. | [optional] 
**MiddleName** | Pointer to **string** | Customer&#39;s middle name | [optional] 
**Note** | Pointer to **string** | Add an optional note when creating or updating a customer. A note is required when updating a customers&#39;s ban_status between SUSPENDED and ALLOWED. | [optional] 
**PhoneNumber** | Pointer to **string** | Customer&#39;s mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated. | [optional] 
**ShippingAddress** | Pointer to [**ShippingAddress**](ShippingAddress.md) |  | [optional] 
**Ssn** | Pointer to **string** | Customer&#39;s full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC on a customer. Must be compiled with ^\\d{3}-\\d{2}-\\d{4}$. Response contains the last 4 digits only (e.g. 6789). | [optional] 
**Status** | Pointer to **string** | Customer&#39;s status | [optional] 

## Methods

### NewPatchCustomer

`func NewPatchCustomer() *PatchCustomer`

NewPatchCustomer instantiates a new PatchCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCustomerWithDefaults

`func NewPatchCustomerWithDefaults() *PatchCustomer`

NewPatchCustomerWithDefaults instantiates a new PatchCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanStatus

`func (o *PatchCustomer) GetBanStatus() BanStatus`

GetBanStatus returns the BanStatus field if non-nil, zero value otherwise.

### GetBanStatusOk

`func (o *PatchCustomer) GetBanStatusOk() (*BanStatus, bool)`

GetBanStatusOk returns a tuple with the BanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStatus

`func (o *PatchCustomer) SetBanStatus(v BanStatus)`

SetBanStatus sets BanStatus field to given value.

### HasBanStatus

`func (o *PatchCustomer) HasBanStatus() bool`

HasBanStatus returns a boolean if a field has been set.

### GetDob

`func (o *PatchCustomer) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *PatchCustomer) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *PatchCustomer) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *PatchCustomer) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetEmail

`func (o *PatchCustomer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchCustomer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchCustomer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchCustomer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *PatchCustomer) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchCustomer) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchCustomer) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchCustomer) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchCustomer) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchCustomer) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchCustomer) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchCustomer) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLegalAddress

`func (o *PatchCustomer) GetLegalAddress() LegalAddress`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *PatchCustomer) GetLegalAddressOk() (*LegalAddress, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *PatchCustomer) SetLegalAddress(v LegalAddress)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *PatchCustomer) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchCustomer) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchCustomer) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchCustomer) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchCustomer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *PatchCustomer) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PatchCustomer) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PatchCustomer) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PatchCustomer) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNote

`func (o *PatchCustomer) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PatchCustomer) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PatchCustomer) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PatchCustomer) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PatchCustomer) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PatchCustomer) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PatchCustomer) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PatchCustomer) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetShippingAddress

`func (o *PatchCustomer) GetShippingAddress() ShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *PatchCustomer) GetShippingAddressOk() (*ShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *PatchCustomer) SetShippingAddress(v ShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *PatchCustomer) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetSsn

`func (o *PatchCustomer) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *PatchCustomer) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *PatchCustomer) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *PatchCustomer) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetStatus

`func (o *PatchCustomer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchCustomer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchCustomer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchCustomer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


