/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// BasePerson1 struct for BasePerson1
type BasePerson1 struct {
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Person's date of birth in RFC 3339 full-date format (YYYY-MM-DD).
	Dob *string `json:"dob,omitempty"`
	// Person's email.
	Email *string `json:"email,omitempty"`
	// Person's first name.
	FirstName *string `json:"first_name,omitempty"`
	// Person's unique identifier.
	Id *string `json:"id,omitempty"`
	// True for personal and business customers with a direct relationship with the fintech or bank.
	IsCustomer *bool `json:"is_customer,omitempty"`
	// Person's last name.
	LastName *string `json:"last_name,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	LegalAddress *Address `json:"legal_address,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Person's middle name.
	MiddleName *string `json:"middle_name,omitempty"`
	// Person's mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated
	PhoneNumber *string `json:"phone_number,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
	// Person's full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC. The response contains the last 4 digits only (e.g. 6789).
	Ssn *string `json:"ssn,omitempty"`
	SsnSource *SsnSource `json:"ssn_source,omitempty"`
	Status *Status1 `json:"status,omitempty"`
	// Date and time KYC verification was last run on the person.
	VerificationLastRun *time.Time `json:"verification_last_run,omitempty"`
	VerificationStatus *VerificationStatus `json:"verification_status,omitempty"`
}

// NewBasePerson1 instantiates a new BasePerson1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasePerson1() *BasePerson1 {
	this := BasePerson1{}
	return &this
}

// NewBasePerson1WithDefaults instantiates a new BasePerson1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasePerson1WithDefaults() *BasePerson1 {
	this := BasePerson1{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BasePerson1) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BasePerson1) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BasePerson1) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *BasePerson1) GetDob() string {
	if o == nil || o.Dob == nil {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetDobOk() (*string, bool) {
	if o == nil || o.Dob == nil {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *BasePerson1) HasDob() bool {
	if o != nil && o.Dob != nil {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *BasePerson1) SetDob(v string) {
	o.Dob = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BasePerson1) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BasePerson1) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BasePerson1) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BasePerson1) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BasePerson1) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BasePerson1) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BasePerson1) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BasePerson1) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BasePerson1) SetId(v string) {
	o.Id = &v
}

// GetIsCustomer returns the IsCustomer field value if set, zero value otherwise.
func (o *BasePerson1) GetIsCustomer() bool {
	if o == nil || o.IsCustomer == nil {
		var ret bool
		return ret
	}
	return *o.IsCustomer
}

// GetIsCustomerOk returns a tuple with the IsCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetIsCustomerOk() (*bool, bool) {
	if o == nil || o.IsCustomer == nil {
		return nil, false
	}
	return o.IsCustomer, true
}

// HasIsCustomer returns a boolean if a field has been set.
func (o *BasePerson1) HasIsCustomer() bool {
	if o != nil && o.IsCustomer != nil {
		return true
	}

	return false
}

// SetIsCustomer gets a reference to the given bool and assigns it to the IsCustomer field.
func (o *BasePerson1) SetIsCustomer(v bool) {
	o.IsCustomer = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BasePerson1) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BasePerson1) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BasePerson1) SetLastName(v string) {
	o.LastName = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *BasePerson1) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *BasePerson1) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *BasePerson1) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetLegalAddress returns the LegalAddress field value if set, zero value otherwise.
func (o *BasePerson1) GetLegalAddress() Address {
	if o == nil || o.LegalAddress == nil {
		var ret Address
		return ret
	}
	return *o.LegalAddress
}

// GetLegalAddressOk returns a tuple with the LegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetLegalAddressOk() (*Address, bool) {
	if o == nil || o.LegalAddress == nil {
		return nil, false
	}
	return o.LegalAddress, true
}

// HasLegalAddress returns a boolean if a field has been set.
func (o *BasePerson1) HasLegalAddress() bool {
	if o != nil && o.LegalAddress != nil {
		return true
	}

	return false
}

// SetLegalAddress gets a reference to the given Address and assigns it to the LegalAddress field.
func (o *BasePerson1) SetLegalAddress(v Address) {
	o.LegalAddress = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BasePerson1) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BasePerson1) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *BasePerson1) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *BasePerson1) GetMiddleName() string {
	if o == nil || o.MiddleName == nil {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetMiddleNameOk() (*string, bool) {
	if o == nil || o.MiddleName == nil {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *BasePerson1) HasMiddleName() bool {
	if o != nil && o.MiddleName != nil {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *BasePerson1) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *BasePerson1) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *BasePerson1) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *BasePerson1) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *BasePerson1) GetShippingAddress() Address {
	if o == nil || o.ShippingAddress == nil {
		var ret Address
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetShippingAddressOk() (*Address, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *BasePerson1) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given Address and assigns it to the ShippingAddress field.
func (o *BasePerson1) SetShippingAddress(v Address) {
	o.ShippingAddress = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *BasePerson1) GetSsn() string {
	if o == nil || o.Ssn == nil {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetSsnOk() (*string, bool) {
	if o == nil || o.Ssn == nil {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *BasePerson1) HasSsn() bool {
	if o != nil && o.Ssn != nil {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *BasePerson1) SetSsn(v string) {
	o.Ssn = &v
}

// GetSsnSource returns the SsnSource field value if set, zero value otherwise.
func (o *BasePerson1) GetSsnSource() SsnSource {
	if o == nil || o.SsnSource == nil {
		var ret SsnSource
		return ret
	}
	return *o.SsnSource
}

// GetSsnSourceOk returns a tuple with the SsnSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetSsnSourceOk() (*SsnSource, bool) {
	if o == nil || o.SsnSource == nil {
		return nil, false
	}
	return o.SsnSource, true
}

// HasSsnSource returns a boolean if a field has been set.
func (o *BasePerson1) HasSsnSource() bool {
	if o != nil && o.SsnSource != nil {
		return true
	}

	return false
}

// SetSsnSource gets a reference to the given SsnSource and assigns it to the SsnSource field.
func (o *BasePerson1) SetSsnSource(v SsnSource) {
	o.SsnSource = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BasePerson1) GetStatus() Status1 {
	if o == nil || o.Status == nil {
		var ret Status1
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetStatusOk() (*Status1, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BasePerson1) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status1 and assigns it to the Status field.
func (o *BasePerson1) SetStatus(v Status1) {
	o.Status = &v
}

// GetVerificationLastRun returns the VerificationLastRun field value if set, zero value otherwise.
func (o *BasePerson1) GetVerificationLastRun() time.Time {
	if o == nil || o.VerificationLastRun == nil {
		var ret time.Time
		return ret
	}
	return *o.VerificationLastRun
}

// GetVerificationLastRunOk returns a tuple with the VerificationLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetVerificationLastRunOk() (*time.Time, bool) {
	if o == nil || o.VerificationLastRun == nil {
		return nil, false
	}
	return o.VerificationLastRun, true
}

// HasVerificationLastRun returns a boolean if a field has been set.
func (o *BasePerson1) HasVerificationLastRun() bool {
	if o != nil && o.VerificationLastRun != nil {
		return true
	}

	return false
}

// SetVerificationLastRun gets a reference to the given time.Time and assigns it to the VerificationLastRun field.
func (o *BasePerson1) SetVerificationLastRun(v time.Time) {
	o.VerificationLastRun = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *BasePerson1) GetVerificationStatus() VerificationStatus {
	if o == nil || o.VerificationStatus == nil {
		var ret VerificationStatus
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasePerson1) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *BasePerson1) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given VerificationStatus and assigns it to the VerificationStatus field.
func (o *BasePerson1) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = &v
}

func (o BasePerson1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Dob != nil {
		toSerialize["dob"] = o.Dob
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsCustomer != nil {
		toSerialize["is_customer"] = o.IsCustomer
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.LegalAddress != nil {
		toSerialize["legal_address"] = o.LegalAddress
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MiddleName != nil {
		toSerialize["middle_name"] = o.MiddleName
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.ShippingAddress != nil {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.Ssn != nil {
		toSerialize["ssn"] = o.Ssn
	}
	if o.SsnSource != nil {
		toSerialize["ssn_source"] = o.SsnSource
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.VerificationLastRun != nil {
		toSerialize["verification_last_run"] = o.VerificationLastRun
	}
	if o.VerificationStatus != nil {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableBasePerson1 struct {
	value *BasePerson1
	isSet bool
}

func (v NullableBasePerson1) Get() *BasePerson1 {
	return v.value
}

func (v *NullableBasePerson1) Set(val *BasePerson1) {
	v.value = val
	v.isSet = true
}

func (v NullableBasePerson1) IsSet() bool {
	return v.isSet
}

func (v *NullableBasePerson1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasePerson1(val *BasePerson1) *NullableBasePerson1 {
	return &NullableBasePerson1{value: val, isSet: true}
}

func (v NullableBasePerson1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasePerson1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


