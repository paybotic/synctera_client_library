/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// checks if the BaseCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseCustomer{}

// BaseCustomer struct for BaseCustomer
type BaseCustomer struct {
	// All of the customer's addresses
	Addresses []Address  `json:"addresses,omitempty"`
	BanStatus *BanStatus `json:"ban_status,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Customer's email
	Email *string `json:"email,omitempty"`
	// This flag indicates whether the person or business has accounts.
	HasAccounts *bool `json:"has_accounts,omitempty"`
	// Customer unique identifier
	Id *string `json:"id,omitempty"`
	// Customer's KYC exemption
	KycExempt *bool `json:"kyc_exempt,omitempty"`
	// Date and time KYC was last run on the customer
	KycLastRun *time.Time         `json:"kyc_last_run,omitempty"`
	KycStatus  *CustomerKycStatus `json:"kyc_status,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time    `json:"last_updated_time,omitempty"`
	LegalAddress    *LegalAddress `json:"legal_address,omitempty"`
	// User-supplied metadata. Do not use to store PII.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Customer's middle name
	MiddleName *string `json:"middle_name,omitempty"`
	// Add an optional note when creating or updating a customer. A note is required when updating a customers's ban_status between SUSPENDED and ALLOWED.
	Note *string `json:"note,omitempty"`
	// Customer's mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated.
	PhoneNumber *string `json:"phone_number,omitempty" validate:"regexp=^\\\\+[1-9]\\\\d{1,14}$"`
	// Customer's relationships with other accounts eg. guardian. This property is no longer supported. Setting it will return an error.
	// Deprecated
	RelatedCustomers []Relationship   `json:"related_customers,omitempty"`
	ShippingAddress  *ShippingAddress `json:"shipping_address,omitempty"`
	// Customer's full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC on a customer. Input must match the pattern ^\\d{3}-\\d{2}-\\d{4}$. The response contains the last 4 digits only (e.g. 6789).
	Ssn       *string    `json:"ssn,omitempty"`
	SsnSource *SsnSource `json:"ssn_source,omitempty"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant *string `json:"tenant,omitempty"`
}

// NewBaseCustomer instantiates a new BaseCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseCustomer() *BaseCustomer {
	this := BaseCustomer{}
	return &this
}

// NewBaseCustomerWithDefaults instantiates a new BaseCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseCustomerWithDefaults() *BaseCustomer {
	this := BaseCustomer{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *BaseCustomer) GetAddresses() []Address {
	if o == nil || IsNil(o.Addresses) {
		var ret []Address
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetAddressesOk() ([]Address, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *BaseCustomer) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []Address and assigns it to the Addresses field.
func (o *BaseCustomer) SetAddresses(v []Address) {
	o.Addresses = v
}

// GetBanStatus returns the BanStatus field value if set, zero value otherwise.
func (o *BaseCustomer) GetBanStatus() BanStatus {
	if o == nil || IsNil(o.BanStatus) {
		var ret BanStatus
		return ret
	}
	return *o.BanStatus
}

// GetBanStatusOk returns a tuple with the BanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetBanStatusOk() (*BanStatus, bool) {
	if o == nil || IsNil(o.BanStatus) {
		return nil, false
	}
	return o.BanStatus, true
}

// HasBanStatus returns a boolean if a field has been set.
func (o *BaseCustomer) HasBanStatus() bool {
	if o != nil && !IsNil(o.BanStatus) {
		return true
	}

	return false
}

// SetBanStatus gets a reference to the given BanStatus and assigns it to the BanStatus field.
func (o *BaseCustomer) SetBanStatus(v BanStatus) {
	o.BanStatus = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BaseCustomer) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BaseCustomer) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BaseCustomer) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BaseCustomer) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BaseCustomer) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BaseCustomer) SetEmail(v string) {
	o.Email = &v
}

// GetHasAccounts returns the HasAccounts field value if set, zero value otherwise.
func (o *BaseCustomer) GetHasAccounts() bool {
	if o == nil || IsNil(o.HasAccounts) {
		var ret bool
		return ret
	}
	return *o.HasAccounts
}

// GetHasAccountsOk returns a tuple with the HasAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetHasAccountsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAccounts) {
		return nil, false
	}
	return o.HasAccounts, true
}

// HasHasAccounts returns a boolean if a field has been set.
func (o *BaseCustomer) HasHasAccounts() bool {
	if o != nil && !IsNil(o.HasAccounts) {
		return true
	}

	return false
}

// SetHasAccounts gets a reference to the given bool and assigns it to the HasAccounts field.
func (o *BaseCustomer) SetHasAccounts(v bool) {
	o.HasAccounts = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseCustomer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseCustomer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseCustomer) SetId(v string) {
	o.Id = &v
}

// GetKycExempt returns the KycExempt field value if set, zero value otherwise.
func (o *BaseCustomer) GetKycExempt() bool {
	if o == nil || IsNil(o.KycExempt) {
		var ret bool
		return ret
	}
	return *o.KycExempt
}

// GetKycExemptOk returns a tuple with the KycExempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetKycExemptOk() (*bool, bool) {
	if o == nil || IsNil(o.KycExempt) {
		return nil, false
	}
	return o.KycExempt, true
}

// HasKycExempt returns a boolean if a field has been set.
func (o *BaseCustomer) HasKycExempt() bool {
	if o != nil && !IsNil(o.KycExempt) {
		return true
	}

	return false
}

// SetKycExempt gets a reference to the given bool and assigns it to the KycExempt field.
func (o *BaseCustomer) SetKycExempt(v bool) {
	o.KycExempt = &v
}

// GetKycLastRun returns the KycLastRun field value if set, zero value otherwise.
func (o *BaseCustomer) GetKycLastRun() time.Time {
	if o == nil || IsNil(o.KycLastRun) {
		var ret time.Time
		return ret
	}
	return *o.KycLastRun
}

// GetKycLastRunOk returns a tuple with the KycLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetKycLastRunOk() (*time.Time, bool) {
	if o == nil || IsNil(o.KycLastRun) {
		return nil, false
	}
	return o.KycLastRun, true
}

// HasKycLastRun returns a boolean if a field has been set.
func (o *BaseCustomer) HasKycLastRun() bool {
	if o != nil && !IsNil(o.KycLastRun) {
		return true
	}

	return false
}

// SetKycLastRun gets a reference to the given time.Time and assigns it to the KycLastRun field.
func (o *BaseCustomer) SetKycLastRun(v time.Time) {
	o.KycLastRun = &v
}

// GetKycStatus returns the KycStatus field value if set, zero value otherwise.
func (o *BaseCustomer) GetKycStatus() CustomerKycStatus {
	if o == nil || IsNil(o.KycStatus) {
		var ret CustomerKycStatus
		return ret
	}
	return *o.KycStatus
}

// GetKycStatusOk returns a tuple with the KycStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetKycStatusOk() (*CustomerKycStatus, bool) {
	if o == nil || IsNil(o.KycStatus) {
		return nil, false
	}
	return o.KycStatus, true
}

// HasKycStatus returns a boolean if a field has been set.
func (o *BaseCustomer) HasKycStatus() bool {
	if o != nil && !IsNil(o.KycStatus) {
		return true
	}

	return false
}

// SetKycStatus gets a reference to the given CustomerKycStatus and assigns it to the KycStatus field.
func (o *BaseCustomer) SetKycStatus(v CustomerKycStatus) {
	o.KycStatus = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *BaseCustomer) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *BaseCustomer) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *BaseCustomer) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetLegalAddress returns the LegalAddress field value if set, zero value otherwise.
func (o *BaseCustomer) GetLegalAddress() LegalAddress {
	if o == nil || IsNil(o.LegalAddress) {
		var ret LegalAddress
		return ret
	}
	return *o.LegalAddress
}

// GetLegalAddressOk returns a tuple with the LegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetLegalAddressOk() (*LegalAddress, bool) {
	if o == nil || IsNil(o.LegalAddress) {
		return nil, false
	}
	return o.LegalAddress, true
}

// HasLegalAddress returns a boolean if a field has been set.
func (o *BaseCustomer) HasLegalAddress() bool {
	if o != nil && !IsNil(o.LegalAddress) {
		return true
	}

	return false
}

// SetLegalAddress gets a reference to the given LegalAddress and assigns it to the LegalAddress field.
func (o *BaseCustomer) SetLegalAddress(v LegalAddress) {
	o.LegalAddress = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BaseCustomer) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BaseCustomer) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *BaseCustomer) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *BaseCustomer) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *BaseCustomer) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *BaseCustomer) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *BaseCustomer) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *BaseCustomer) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *BaseCustomer) SetNote(v string) {
	o.Note = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *BaseCustomer) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *BaseCustomer) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *BaseCustomer) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetRelatedCustomers returns the RelatedCustomers field value if set, zero value otherwise.
// Deprecated
func (o *BaseCustomer) GetRelatedCustomers() []Relationship {
	if o == nil || IsNil(o.RelatedCustomers) {
		var ret []Relationship
		return ret
	}
	return o.RelatedCustomers
}

// GetRelatedCustomersOk returns a tuple with the RelatedCustomers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BaseCustomer) GetRelatedCustomersOk() ([]Relationship, bool) {
	if o == nil || IsNil(o.RelatedCustomers) {
		return nil, false
	}
	return o.RelatedCustomers, true
}

// HasRelatedCustomers returns a boolean if a field has been set.
func (o *BaseCustomer) HasRelatedCustomers() bool {
	if o != nil && !IsNil(o.RelatedCustomers) {
		return true
	}

	return false
}

// SetRelatedCustomers gets a reference to the given []Relationship and assigns it to the RelatedCustomers field.
// Deprecated
func (o *BaseCustomer) SetRelatedCustomers(v []Relationship) {
	o.RelatedCustomers = v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *BaseCustomer) GetShippingAddress() ShippingAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret ShippingAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetShippingAddressOk() (*ShippingAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *BaseCustomer) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given ShippingAddress and assigns it to the ShippingAddress field.
func (o *BaseCustomer) SetShippingAddress(v ShippingAddress) {
	o.ShippingAddress = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *BaseCustomer) GetSsn() string {
	if o == nil || IsNil(o.Ssn) {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetSsnOk() (*string, bool) {
	if o == nil || IsNil(o.Ssn) {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *BaseCustomer) HasSsn() bool {
	if o != nil && !IsNil(o.Ssn) {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *BaseCustomer) SetSsn(v string) {
	o.Ssn = &v
}

// GetSsnSource returns the SsnSource field value if set, zero value otherwise.
func (o *BaseCustomer) GetSsnSource() SsnSource {
	if o == nil || IsNil(o.SsnSource) {
		var ret SsnSource
		return ret
	}
	return *o.SsnSource
}

// GetSsnSourceOk returns a tuple with the SsnSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetSsnSourceOk() (*SsnSource, bool) {
	if o == nil || IsNil(o.SsnSource) {
		return nil, false
	}
	return o.SsnSource, true
}

// HasSsnSource returns a boolean if a field has been set.
func (o *BaseCustomer) HasSsnSource() bool {
	if o != nil && !IsNil(o.SsnSource) {
		return true
	}

	return false
}

// SetSsnSource gets a reference to the given SsnSource and assigns it to the SsnSource field.
func (o *BaseCustomer) SetSsnSource(v SsnSource) {
	o.SsnSource = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *BaseCustomer) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCustomer) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *BaseCustomer) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *BaseCustomer) SetTenant(v string) {
	o.Tenant = &v
}

func (o BaseCustomer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.BanStatus) {
		toSerialize["ban_status"] = o.BanStatus
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.HasAccounts) {
		toSerialize["has_accounts"] = o.HasAccounts
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.KycExempt) {
		toSerialize["kyc_exempt"] = o.KycExempt
	}
	if !IsNil(o.KycLastRun) {
		toSerialize["kyc_last_run"] = o.KycLastRun
	}
	if !IsNil(o.KycStatus) {
		toSerialize["kyc_status"] = o.KycStatus
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
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.RelatedCustomers) {
		toSerialize["related_customers"] = o.RelatedCustomers
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
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	return toSerialize, nil
}

type NullableBaseCustomer struct {
	value *BaseCustomer
	isSet bool
}

func (v NullableBaseCustomer) Get() *BaseCustomer {
	return v.value
}

func (v *NullableBaseCustomer) Set(val *BaseCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseCustomer(val *BaseCustomer) *NullableBaseCustomer {
	return &NullableBaseCustomer{value: val, isSet: true}
}

func (v NullableBaseCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
