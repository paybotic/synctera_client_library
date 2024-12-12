/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the ResponsePerson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsePerson{}

// ResponsePerson Details of a person
type ResponsePerson struct {
	BanStatus BanStatus `json:"ban_status"`
	// Person's chosen name.
	ChosenName *string `json:"chosen_name,omitempty"`
	// The date and time the resource was created.
	CreationTime time.Time `json:"creation_time"`
	// Person's date of birth in RFC 3339 full-date format (YYYY-MM-DD). Must be on or after 1900-01-01 and before current date.
	Dob *string `json:"dob,omitempty"`
	// Person's email.
	Email *string `json:"email,omitempty"`
	// Person's first name.
	FirstName *string `json:"first_name,omitempty"`
	// This flag indicates whether the person or business has accounts.
	HasAccounts *bool `json:"has_accounts,omitempty"`
	// Person's unique identifier.
	Id string `json:"id"`
	// True for personal and business customers with a direct relationship with the fintech or bank. Set this to true for any customer related to an account.
	IsCustomer bool `json:"is_customer"`
	// Person's last name.
	LastName *string `json:"last_name,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime time.Time     `json:"last_updated_time"`
	LegalAddress    *LegalAddress `json:"legal_address,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Person's middle name.
	MiddleName *string `json:"middle_name,omitempty"`
	// Person's mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated
	PhoneNumber     *string          `json:"phone_number,omitempty" validate:"regexp=^\\\\+[1-9]\\\\d{1,14}$"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	// Person's full tax ID eg SSN formatted with hyphens. The response contains the last 4 digits only (e.g. 6789).
	Ssn       *string      `json:"ssn,omitempty"`
	SsnSource *SsnSource   `json:"ssn_source,omitempty"`
	Status    PersonStatus `json:"status"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant string `json:"tenant"`
	// Date and time KYC verification was last run on the person.
	VerificationLastRun *time.Time         `json:"verification_last_run,omitempty"`
	VerificationStatus  VerificationStatus `json:"verification_status"`
	// Array of personal identifiers
	PersonalIds []ResponsePersonalId `json:"personal_ids,omitempty"`
}

type _ResponsePerson ResponsePerson

// NewResponsePerson instantiates a new ResponsePerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePerson(banStatus BanStatus, creationTime time.Time, id string, isCustomer bool, lastUpdatedTime time.Time, status PersonStatus, tenant string, verificationStatus VerificationStatus) *ResponsePerson {
	this := ResponsePerson{}
	this.BanStatus = banStatus
	this.CreationTime = creationTime
	this.Id = id
	this.IsCustomer = isCustomer
	this.LastUpdatedTime = lastUpdatedTime
	this.Status = status
	this.Tenant = tenant
	this.VerificationStatus = verificationStatus
	return &this
}

// NewResponsePersonWithDefaults instantiates a new ResponsePerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePersonWithDefaults() *ResponsePerson {
	this := ResponsePerson{}
	return &this
}

// GetBanStatus returns the BanStatus field value
func (o *ResponsePerson) GetBanStatus() BanStatus {
	if o == nil {
		var ret BanStatus
		return ret
	}

	return o.BanStatus
}

// GetBanStatusOk returns a tuple with the BanStatus field value
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetBanStatusOk() (*BanStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BanStatus, true
}

// SetBanStatus sets field value
func (o *ResponsePerson) SetBanStatus(v BanStatus) {
	o.BanStatus = v
}

// GetChosenName returns the ChosenName field value if set, zero value otherwise.
func (o *ResponsePerson) GetChosenName() string {
	if o == nil || IsNil(o.ChosenName) {
		var ret string
		return ret
	}
	return *o.ChosenName
}

// GetChosenNameOk returns a tuple with the ChosenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetChosenNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChosenName) {
		return nil, false
	}
	return o.ChosenName, true
}

// HasChosenName returns a boolean if a field has been set.
func (o *ResponsePerson) HasChosenName() bool {
	if o != nil && !IsNil(o.ChosenName) {
		return true
	}

	return false
}

// SetChosenName gets a reference to the given string and assigns it to the ChosenName field.
func (o *ResponsePerson) SetChosenName(v string) {
	o.ChosenName = &v
}

// GetCreationTime returns the CreationTime field value
func (o *ResponsePerson) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *ResponsePerson) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *ResponsePerson) GetDob() string {
	if o == nil || IsNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetDobOk() (*string, bool) {
	if o == nil || IsNil(o.Dob) {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *ResponsePerson) HasDob() bool {
	if o != nil && !IsNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *ResponsePerson) SetDob(v string) {
	o.Dob = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ResponsePerson) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ResponsePerson) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ResponsePerson) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ResponsePerson) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ResponsePerson) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ResponsePerson) SetFirstName(v string) {
	o.FirstName = &v
}

// GetHasAccounts returns the HasAccounts field value if set, zero value otherwise.
func (o *ResponsePerson) GetHasAccounts() bool {
	if o == nil || IsNil(o.HasAccounts) {
		var ret bool
		return ret
	}
	return *o.HasAccounts
}

// GetHasAccountsOk returns a tuple with the HasAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetHasAccountsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAccounts) {
		return nil, false
	}
	return o.HasAccounts, true
}

// HasHasAccounts returns a boolean if a field has been set.
func (o *ResponsePerson) HasHasAccounts() bool {
	if o != nil && !IsNil(o.HasAccounts) {
		return true
	}

	return false
}

// SetHasAccounts gets a reference to the given bool and assigns it to the HasAccounts field.
func (o *ResponsePerson) SetHasAccounts(v bool) {
	o.HasAccounts = &v
}

// GetId returns the Id field value
func (o *ResponsePerson) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponsePerson) SetId(v string) {
	o.Id = v
}

// GetIsCustomer returns the IsCustomer field value
func (o *ResponsePerson) GetIsCustomer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCustomer
}

// GetIsCustomerOk returns a tuple with the IsCustomer field value
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetIsCustomerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCustomer, true
}

// SetIsCustomer sets field value
func (o *ResponsePerson) SetIsCustomer(v bool) {
	o.IsCustomer = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ResponsePerson) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ResponsePerson) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ResponsePerson) SetLastName(v string) {
	o.LastName = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *ResponsePerson) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *ResponsePerson) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetLegalAddress returns the LegalAddress field value if set, zero value otherwise.
func (o *ResponsePerson) GetLegalAddress() LegalAddress {
	if o == nil || IsNil(o.LegalAddress) {
		var ret LegalAddress
		return ret
	}
	return *o.LegalAddress
}

// GetLegalAddressOk returns a tuple with the LegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetLegalAddressOk() (*LegalAddress, bool) {
	if o == nil || IsNil(o.LegalAddress) {
		return nil, false
	}
	return o.LegalAddress, true
}

// HasLegalAddress returns a boolean if a field has been set.
func (o *ResponsePerson) HasLegalAddress() bool {
	if o != nil && !IsNil(o.LegalAddress) {
		return true
	}

	return false
}

// SetLegalAddress gets a reference to the given LegalAddress and assigns it to the LegalAddress field.
func (o *ResponsePerson) SetLegalAddress(v LegalAddress) {
	o.LegalAddress = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ResponsePerson) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ResponsePerson) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ResponsePerson) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *ResponsePerson) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *ResponsePerson) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *ResponsePerson) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ResponsePerson) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ResponsePerson) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ResponsePerson) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *ResponsePerson) GetShippingAddress() ShippingAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret ShippingAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetShippingAddressOk() (*ShippingAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *ResponsePerson) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given ShippingAddress and assigns it to the ShippingAddress field.
func (o *ResponsePerson) SetShippingAddress(v ShippingAddress) {
	o.ShippingAddress = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *ResponsePerson) GetSsn() string {
	if o == nil || IsNil(o.Ssn) {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetSsnOk() (*string, bool) {
	if o == nil || IsNil(o.Ssn) {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *ResponsePerson) HasSsn() bool {
	if o != nil && !IsNil(o.Ssn) {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *ResponsePerson) SetSsn(v string) {
	o.Ssn = &v
}

// GetSsnSource returns the SsnSource field value if set, zero value otherwise.
func (o *ResponsePerson) GetSsnSource() SsnSource {
	if o == nil || IsNil(o.SsnSource) {
		var ret SsnSource
		return ret
	}
	return *o.SsnSource
}

// GetSsnSourceOk returns a tuple with the SsnSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetSsnSourceOk() (*SsnSource, bool) {
	if o == nil || IsNil(o.SsnSource) {
		return nil, false
	}
	return o.SsnSource, true
}

// HasSsnSource returns a boolean if a field has been set.
func (o *ResponsePerson) HasSsnSource() bool {
	if o != nil && !IsNil(o.SsnSource) {
		return true
	}

	return false
}

// SetSsnSource gets a reference to the given SsnSource and assigns it to the SsnSource field.
func (o *ResponsePerson) SetSsnSource(v SsnSource) {
	o.SsnSource = &v
}

// GetStatus returns the Status field value
func (o *ResponsePerson) GetStatus() PersonStatus {
	if o == nil {
		var ret PersonStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetStatusOk() (*PersonStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResponsePerson) SetStatus(v PersonStatus) {
	o.Status = v
}

// GetTenant returns the Tenant field value
func (o *ResponsePerson) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *ResponsePerson) SetTenant(v string) {
	o.Tenant = v
}

// GetVerificationLastRun returns the VerificationLastRun field value if set, zero value otherwise.
func (o *ResponsePerson) GetVerificationLastRun() time.Time {
	if o == nil || IsNil(o.VerificationLastRun) {
		var ret time.Time
		return ret
	}
	return *o.VerificationLastRun
}

// GetVerificationLastRunOk returns a tuple with the VerificationLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetVerificationLastRunOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VerificationLastRun) {
		return nil, false
	}
	return o.VerificationLastRun, true
}

// HasVerificationLastRun returns a boolean if a field has been set.
func (o *ResponsePerson) HasVerificationLastRun() bool {
	if o != nil && !IsNil(o.VerificationLastRun) {
		return true
	}

	return false
}

// SetVerificationLastRun gets a reference to the given time.Time and assigns it to the VerificationLastRun field.
func (o *ResponsePerson) SetVerificationLastRun(v time.Time) {
	o.VerificationLastRun = &v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *ResponsePerson) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *ResponsePerson) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

// GetPersonalIds returns the PersonalIds field value if set, zero value otherwise.
func (o *ResponsePerson) GetPersonalIds() []ResponsePersonalId {
	if o == nil || IsNil(o.PersonalIds) {
		var ret []ResponsePersonalId
		return ret
	}
	return o.PersonalIds
}

// GetPersonalIdsOk returns a tuple with the PersonalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePerson) GetPersonalIdsOk() ([]ResponsePersonalId, bool) {
	if o == nil || IsNil(o.PersonalIds) {
		return nil, false
	}
	return o.PersonalIds, true
}

// HasPersonalIds returns a boolean if a field has been set.
func (o *ResponsePerson) HasPersonalIds() bool {
	if o != nil && !IsNil(o.PersonalIds) {
		return true
	}

	return false
}

// SetPersonalIds gets a reference to the given []ResponsePersonalId and assigns it to the PersonalIds field.
func (o *ResponsePerson) SetPersonalIds(v []ResponsePersonalId) {
	o.PersonalIds = v
}

func (o ResponsePerson) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsePerson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ban_status"] = o.BanStatus
	if !IsNil(o.ChosenName) {
		toSerialize["chosen_name"] = o.ChosenName
	}
	toSerialize["creation_time"] = o.CreationTime
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
	toSerialize["id"] = o.Id
	toSerialize["is_customer"] = o.IsCustomer
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	toSerialize["last_updated_time"] = o.LastUpdatedTime
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
	toSerialize["tenant"] = o.Tenant
	if !IsNil(o.VerificationLastRun) {
		toSerialize["verification_last_run"] = o.VerificationLastRun
	}
	toSerialize["verification_status"] = o.VerificationStatus
	if !IsNil(o.PersonalIds) {
		toSerialize["personal_ids"] = o.PersonalIds
	}
	return toSerialize, nil
}

func (o *ResponsePerson) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ban_status",
		"creation_time",
		"id",
		"is_customer",
		"last_updated_time",
		"status",
		"tenant",
		"verification_status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResponsePerson := _ResponsePerson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponsePerson)

	if err != nil {
		return err
	}

	*o = ResponsePerson(varResponsePerson)

	return err
}

type NullableResponsePerson struct {
	value *ResponsePerson
	isSet bool
}

func (v NullableResponsePerson) Get() *ResponsePerson {
	return v.value
}

func (v *NullableResponsePerson) Set(val *ResponsePerson) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePerson) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePerson(val *ResponsePerson) *NullableResponsePerson {
	return &NullableResponsePerson{value: val, isSet: true}
}

func (v NullableResponsePerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
