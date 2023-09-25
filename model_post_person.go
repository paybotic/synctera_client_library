/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the PostPerson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPerson{}

// PostPerson Details of a person
type PostPerson struct {
	BanStatus *BanStatus `json:"ban_status,omitempty"`
	// Person's chosen name.
	ChosenName *string `json:"chosen_name,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Person's date of birth in RFC 3339 full-date format (YYYY-MM-DD).
	Dob *string `json:"dob,omitempty"`
	// Person's email.
	Email *string `json:"email,omitempty"`
	// Person's first name.
	FirstName *string `json:"first_name,omitempty"`
	// This flag indicates whether the person or business has accounts.
	HasAccounts *bool `json:"has_accounts,omitempty"`
	// Person's unique identifier.
	Id *string `json:"id,omitempty"`
	// True for personal and business customers with a direct relationship with the fintech or bank. Set this to true for any customer related to an account. 
	IsCustomer bool `json:"is_customer"`
	// Person's last name.
	LastName *string `json:"last_name,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	LegalAddress *LegalAddress `json:"legal_address,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Person's middle name.
	MiddleName *string `json:"middle_name,omitempty"`
	// Person's mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated
	PhoneNumber *string `json:"phone_number,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	// Person's full tax ID eg SSN formatted with hyphens. The response contains the last 4 digits only (e.g. 6789).
	Ssn *string `json:"ssn,omitempty"`
	SsnSource *SsnSource `json:"ssn_source,omitempty"`
	Status Status2 `json:"status"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces. 
	Tenant *string `json:"tenant,omitempty"`
	// Date and time KYC verification was last run on the person.
	VerificationLastRun *time.Time `json:"verification_last_run,omitempty"`
	VerificationStatus *VerificationStatus `json:"verification_status,omitempty"`
	// > 🚧 Beta > This is a Beta property. Feedback from the community is welcome. We may make breaking changes to this property. Array of personal identifiers 
	PersonalIds []PostPersonalId `json:"personal_ids,omitempty"`
	// Text to be added to a note when creating a person. A note is required when creating a person with a ban_status of SUSPENDED.
	Note *string `json:"note,omitempty"`
}

// NewPostPerson instantiates a new PostPerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPerson(isCustomer bool, status Status2) *PostPerson {
	this := PostPerson{}
	this.IsCustomer = isCustomer
	this.Status = status
	return &this
}

// NewPostPersonWithDefaults instantiates a new PostPerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPersonWithDefaults() *PostPerson {
	this := PostPerson{}
	return &this
}

// GetBanStatus returns the BanStatus field value if set, zero value otherwise.
func (o *PostPerson) GetBanStatus() BanStatus {
	if o == nil || IsNil(o.BanStatus) {
		var ret BanStatus
		return ret
	}
	return *o.BanStatus
}

// GetBanStatusOk returns a tuple with the BanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetBanStatusOk() (*BanStatus, bool) {
	if o == nil || IsNil(o.BanStatus) {
		return nil, false
	}
	return o.BanStatus, true
}

// HasBanStatus returns a boolean if a field has been set.
func (o *PostPerson) HasBanStatus() bool {
	if o != nil && !IsNil(o.BanStatus) {
		return true
	}

	return false
}

// SetBanStatus gets a reference to the given BanStatus and assigns it to the BanStatus field.
func (o *PostPerson) SetBanStatus(v BanStatus) {
	o.BanStatus = &v
}

// GetChosenName returns the ChosenName field value if set, zero value otherwise.
func (o *PostPerson) GetChosenName() string {
	if o == nil || IsNil(o.ChosenName) {
		var ret string
		return ret
	}
	return *o.ChosenName
}

// GetChosenNameOk returns a tuple with the ChosenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetChosenNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChosenName) {
		return nil, false
	}
	return o.ChosenName, true
}

// HasChosenName returns a boolean if a field has been set.
func (o *PostPerson) HasChosenName() bool {
	if o != nil && !IsNil(o.ChosenName) {
		return true
	}

	return false
}

// SetChosenName gets a reference to the given string and assigns it to the ChosenName field.
func (o *PostPerson) SetChosenName(v string) {
	o.ChosenName = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PostPerson) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PostPerson) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PostPerson) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *PostPerson) GetDob() string {
	if o == nil || IsNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetDobOk() (*string, bool) {
	if o == nil || IsNil(o.Dob) {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *PostPerson) HasDob() bool {
	if o != nil && !IsNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *PostPerson) SetDob(v string) {
	o.Dob = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PostPerson) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PostPerson) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PostPerson) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PostPerson) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PostPerson) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PostPerson) SetFirstName(v string) {
	o.FirstName = &v
}

// GetHasAccounts returns the HasAccounts field value if set, zero value otherwise.
func (o *PostPerson) GetHasAccounts() bool {
	if o == nil || IsNil(o.HasAccounts) {
		var ret bool
		return ret
	}
	return *o.HasAccounts
}

// GetHasAccountsOk returns a tuple with the HasAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetHasAccountsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAccounts) {
		return nil, false
	}
	return o.HasAccounts, true
}

// HasHasAccounts returns a boolean if a field has been set.
func (o *PostPerson) HasHasAccounts() bool {
	if o != nil && !IsNil(o.HasAccounts) {
		return true
	}

	return false
}

// SetHasAccounts gets a reference to the given bool and assigns it to the HasAccounts field.
func (o *PostPerson) SetHasAccounts(v bool) {
	o.HasAccounts = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostPerson) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostPerson) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PostPerson) SetId(v string) {
	o.Id = &v
}

// GetIsCustomer returns the IsCustomer field value
func (o *PostPerson) GetIsCustomer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCustomer
}

// GetIsCustomerOk returns a tuple with the IsCustomer field value
// and a boolean to check if the value has been set.
func (o *PostPerson) GetIsCustomerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCustomer, true
}

// SetIsCustomer sets field value
func (o *PostPerson) SetIsCustomer(v bool) {
	o.IsCustomer = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PostPerson) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PostPerson) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PostPerson) SetLastName(v string) {
	o.LastName = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *PostPerson) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *PostPerson) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *PostPerson) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetLegalAddress returns the LegalAddress field value if set, zero value otherwise.
func (o *PostPerson) GetLegalAddress() LegalAddress {
	if o == nil || IsNil(o.LegalAddress) {
		var ret LegalAddress
		return ret
	}
	return *o.LegalAddress
}

// GetLegalAddressOk returns a tuple with the LegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetLegalAddressOk() (*LegalAddress, bool) {
	if o == nil || IsNil(o.LegalAddress) {
		return nil, false
	}
	return o.LegalAddress, true
}

// HasLegalAddress returns a boolean if a field has been set.
func (o *PostPerson) HasLegalAddress() bool {
	if o != nil && !IsNil(o.LegalAddress) {
		return true
	}

	return false
}

// SetLegalAddress gets a reference to the given LegalAddress and assigns it to the LegalAddress field.
func (o *PostPerson) SetLegalAddress(v LegalAddress) {
	o.LegalAddress = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PostPerson) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PostPerson) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PostPerson) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *PostPerson) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *PostPerson) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *PostPerson) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PostPerson) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PostPerson) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PostPerson) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *PostPerson) GetShippingAddress() ShippingAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret ShippingAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetShippingAddressOk() (*ShippingAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *PostPerson) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given ShippingAddress and assigns it to the ShippingAddress field.
func (o *PostPerson) SetShippingAddress(v ShippingAddress) {
	o.ShippingAddress = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *PostPerson) GetSsn() string {
	if o == nil || IsNil(o.Ssn) {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetSsnOk() (*string, bool) {
	if o == nil || IsNil(o.Ssn) {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *PostPerson) HasSsn() bool {
	if o != nil && !IsNil(o.Ssn) {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *PostPerson) SetSsn(v string) {
	o.Ssn = &v
}

// GetSsnSource returns the SsnSource field value if set, zero value otherwise.
func (o *PostPerson) GetSsnSource() SsnSource {
	if o == nil || IsNil(o.SsnSource) {
		var ret SsnSource
		return ret
	}
	return *o.SsnSource
}

// GetSsnSourceOk returns a tuple with the SsnSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetSsnSourceOk() (*SsnSource, bool) {
	if o == nil || IsNil(o.SsnSource) {
		return nil, false
	}
	return o.SsnSource, true
}

// HasSsnSource returns a boolean if a field has been set.
func (o *PostPerson) HasSsnSource() bool {
	if o != nil && !IsNil(o.SsnSource) {
		return true
	}

	return false
}

// SetSsnSource gets a reference to the given SsnSource and assigns it to the SsnSource field.
func (o *PostPerson) SetSsnSource(v SsnSource) {
	o.SsnSource = &v
}

// GetStatus returns the Status field value
func (o *PostPerson) GetStatus() Status2 {
	if o == nil {
		var ret Status2
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostPerson) GetStatusOk() (*Status2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PostPerson) SetStatus(v Status2) {
	o.Status = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *PostPerson) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *PostPerson) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *PostPerson) SetTenant(v string) {
	o.Tenant = &v
}

// GetVerificationLastRun returns the VerificationLastRun field value if set, zero value otherwise.
func (o *PostPerson) GetVerificationLastRun() time.Time {
	if o == nil || IsNil(o.VerificationLastRun) {
		var ret time.Time
		return ret
	}
	return *o.VerificationLastRun
}

// GetVerificationLastRunOk returns a tuple with the VerificationLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetVerificationLastRunOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VerificationLastRun) {
		return nil, false
	}
	return o.VerificationLastRun, true
}

// HasVerificationLastRun returns a boolean if a field has been set.
func (o *PostPerson) HasVerificationLastRun() bool {
	if o != nil && !IsNil(o.VerificationLastRun) {
		return true
	}

	return false
}

// SetVerificationLastRun gets a reference to the given time.Time and assigns it to the VerificationLastRun field.
func (o *PostPerson) SetVerificationLastRun(v time.Time) {
	o.VerificationLastRun = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *PostPerson) GetVerificationStatus() VerificationStatus {
	if o == nil || IsNil(o.VerificationStatus) {
		var ret VerificationStatus
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil || IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *PostPerson) HasVerificationStatus() bool {
	if o != nil && !IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given VerificationStatus and assigns it to the VerificationStatus field.
func (o *PostPerson) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = &v
}

// GetPersonalIds returns the PersonalIds field value if set, zero value otherwise.
func (o *PostPerson) GetPersonalIds() []PostPersonalId {
	if o == nil || IsNil(o.PersonalIds) {
		var ret []PostPersonalId
		return ret
	}
	return o.PersonalIds
}

// GetPersonalIdsOk returns a tuple with the PersonalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetPersonalIdsOk() ([]PostPersonalId, bool) {
	if o == nil || IsNil(o.PersonalIds) {
		return nil, false
	}
	return o.PersonalIds, true
}

// HasPersonalIds returns a boolean if a field has been set.
func (o *PostPerson) HasPersonalIds() bool {
	if o != nil && !IsNil(o.PersonalIds) {
		return true
	}

	return false
}

// SetPersonalIds gets a reference to the given []PostPersonalId and assigns it to the PersonalIds field.
func (o *PostPerson) SetPersonalIds(v []PostPersonalId) {
	o.PersonalIds = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *PostPerson) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPerson) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *PostPerson) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *PostPerson) SetNote(v string) {
	o.Note = &v
}

func (o PostPerson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPerson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BanStatus) {
		toSerialize["ban_status"] = o.BanStatus
	}
	if !IsNil(o.ChosenName) {
		toSerialize["chosen_name"] = o.ChosenName
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.Dob) {
		toSerialize["dob"] = o.Dob
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.HasAccounts) {
		toSerialize["has_accounts"] = o.HasAccounts
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["is_customer"] = o.IsCustomer
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if !IsNil(o.LegalAddress) {
		toSerialize["legal_address"] = o.LegalAddress
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if !IsNil(o.Ssn) {
		toSerialize["ssn"] = o.Ssn
	}
	if !IsNil(o.SsnSource) {
		toSerialize["ssn_source"] = o.SsnSource
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.VerificationLastRun) {
		toSerialize["verification_last_run"] = o.VerificationLastRun
	}
	if !IsNil(o.VerificationStatus) {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if !IsNil(o.PersonalIds) {
		toSerialize["personal_ids"] = o.PersonalIds
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	return toSerialize, nil
}

type NullablePostPerson struct {
	value *PostPerson
	isSet bool
}

func (v NullablePostPerson) Get() *PostPerson {
	return v.value
}

func (v *NullablePostPerson) Set(val *PostPerson) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPerson) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPerson(val *PostPerson) *NullablePostPerson {
	return &NullablePostPerson{value: val, isSet: true}
}

func (v NullablePostPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


