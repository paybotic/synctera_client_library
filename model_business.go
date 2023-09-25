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

// checks if the Business type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Business{}

// Business Represents a business customer.
type Business struct {
	// The types of compliance that the business needs to adhere to * `LICENSED_CANNABIS` – A type of compliance restriction where the business would need a cannabis license in order to operate. 
	ComplianceRestrictions []string `json:"compliance_restrictions,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// U.S. Employer Identification Number (EIN) for this business, in the format xx-xxxxxxx.
	Ein *string `json:"ein,omitempty"`
	// Business's email.
	Email *string `json:"email,omitempty"`
	// Business's legal name.
	EntityName *string `json:"entity_name,omitempty"`
	// Date the business was legally registered in RFC 3339 full-date format (YYYY-MM-DD).
	FormationDate *string `json:"formation_date,omitempty"`
	// U.S. state where the business is legally registered (2-letter abbreviation).
	FormationState *string `json:"formation_state,omitempty"`
	// This flag indicates whether the person or business has accounts.
	HasAccounts *bool `json:"has_accounts,omitempty"`
	// Business's unique identifier.
	Id *string `json:"id,omitempty"`
	// True for personal and business customers with a direct relationship with the fintech or bank. Set this to true for any customer related to an account. 
	IsCustomer bool `json:"is_customer"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	LegalAddress *LegalAddress `json:"legal_address,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Business's phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Status of the business. One of the following: * `PROSPECT` – a potential customer, used for information-gathering and disclosures. * `ACTIVE` –  is an integrator defined status.  Integrators should set a business to active if they believe the person to be qualified for conducting business.  Synctera will combine this status with other statuses such a verification to determine if the business is eligible for specific actions such as initiating transactions or issuing a card. * `FROZEN` – business's actions are blocked for security, legal, or other reasons. * `SANCTION` – business is on a sanctions list and should be carefully monitored. * `DISSOLVED` – an inactive status indicating a business entity has filed articles of dissolution or a certificate of termination to terminate its existence. * `CANCELLED` – an inactive status indicating that a business entity has filed a cancellation or has failed to file its periodic report after notice of forfeiture of its rights to do business. * `SUSPENDED` – an inactive status indicating that the business entity has lost the right to operate in it's registered jurisdiction. * `MERGED` – an inactive status indicating that the business entity has terminated existence by merging into another entity. * `INACTIVE` – an inactive status indicating that the business entity is no longer active. * `CONVERTED` – An inactive status indicating that the business entity has been converted to another type of business entity in the same jurisdiction. 
	Status string `json:"status"`
	// Business's legal structure.
	Structure *string `json:"structure,omitempty"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces. 
	Tenant *string `json:"tenant,omitempty"`
	// Other names by which this business is known.
	TradeNames []string `json:"trade_names,omitempty"`
	// Date and time KYB verification was last run on the business.
	VerificationLastRun *time.Time `json:"verification_last_run,omitempty"`
	VerificationStatus *VerificationStatus `json:"verification_status,omitempty"`
	// Business's website.
	Website *string `json:"website,omitempty"`
}

// NewBusiness instantiates a new Business object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusiness(isCustomer bool, status string) *Business {
	this := Business{}
	this.IsCustomer = isCustomer
	this.Status = status
	return &this
}

// NewBusinessWithDefaults instantiates a new Business object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessWithDefaults() *Business {
	this := Business{}
	return &this
}

// GetComplianceRestrictions returns the ComplianceRestrictions field value if set, zero value otherwise.
func (o *Business) GetComplianceRestrictions() []string {
	if o == nil || IsNil(o.ComplianceRestrictions) {
		var ret []string
		return ret
	}
	return o.ComplianceRestrictions
}

// GetComplianceRestrictionsOk returns a tuple with the ComplianceRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetComplianceRestrictionsOk() ([]string, bool) {
	if o == nil || IsNil(o.ComplianceRestrictions) {
		return nil, false
	}
	return o.ComplianceRestrictions, true
}

// HasComplianceRestrictions returns a boolean if a field has been set.
func (o *Business) HasComplianceRestrictions() bool {
	if o != nil && !IsNil(o.ComplianceRestrictions) {
		return true
	}

	return false
}

// SetComplianceRestrictions gets a reference to the given []string and assigns it to the ComplianceRestrictions field.
func (o *Business) SetComplianceRestrictions(v []string) {
	o.ComplianceRestrictions = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Business) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Business) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Business) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEin returns the Ein field value if set, zero value otherwise.
func (o *Business) GetEin() string {
	if o == nil || IsNil(o.Ein) {
		var ret string
		return ret
	}
	return *o.Ein
}

// GetEinOk returns a tuple with the Ein field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetEinOk() (*string, bool) {
	if o == nil || IsNil(o.Ein) {
		return nil, false
	}
	return o.Ein, true
}

// HasEin returns a boolean if a field has been set.
func (o *Business) HasEin() bool {
	if o != nil && !IsNil(o.Ein) {
		return true
	}

	return false
}

// SetEin gets a reference to the given string and assigns it to the Ein field.
func (o *Business) SetEin(v string) {
	o.Ein = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Business) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Business) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Business) SetEmail(v string) {
	o.Email = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *Business) GetEntityName() string {
	if o == nil || IsNil(o.EntityName) {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetEntityNameOk() (*string, bool) {
	if o == nil || IsNil(o.EntityName) {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *Business) HasEntityName() bool {
	if o != nil && !IsNil(o.EntityName) {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *Business) SetEntityName(v string) {
	o.EntityName = &v
}

// GetFormationDate returns the FormationDate field value if set, zero value otherwise.
func (o *Business) GetFormationDate() string {
	if o == nil || IsNil(o.FormationDate) {
		var ret string
		return ret
	}
	return *o.FormationDate
}

// GetFormationDateOk returns a tuple with the FormationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetFormationDateOk() (*string, bool) {
	if o == nil || IsNil(o.FormationDate) {
		return nil, false
	}
	return o.FormationDate, true
}

// HasFormationDate returns a boolean if a field has been set.
func (o *Business) HasFormationDate() bool {
	if o != nil && !IsNil(o.FormationDate) {
		return true
	}

	return false
}

// SetFormationDate gets a reference to the given string and assigns it to the FormationDate field.
func (o *Business) SetFormationDate(v string) {
	o.FormationDate = &v
}

// GetFormationState returns the FormationState field value if set, zero value otherwise.
func (o *Business) GetFormationState() string {
	if o == nil || IsNil(o.FormationState) {
		var ret string
		return ret
	}
	return *o.FormationState
}

// GetFormationStateOk returns a tuple with the FormationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetFormationStateOk() (*string, bool) {
	if o == nil || IsNil(o.FormationState) {
		return nil, false
	}
	return o.FormationState, true
}

// HasFormationState returns a boolean if a field has been set.
func (o *Business) HasFormationState() bool {
	if o != nil && !IsNil(o.FormationState) {
		return true
	}

	return false
}

// SetFormationState gets a reference to the given string and assigns it to the FormationState field.
func (o *Business) SetFormationState(v string) {
	o.FormationState = &v
}

// GetHasAccounts returns the HasAccounts field value if set, zero value otherwise.
func (o *Business) GetHasAccounts() bool {
	if o == nil || IsNil(o.HasAccounts) {
		var ret bool
		return ret
	}
	return *o.HasAccounts
}

// GetHasAccountsOk returns a tuple with the HasAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetHasAccountsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAccounts) {
		return nil, false
	}
	return o.HasAccounts, true
}

// HasHasAccounts returns a boolean if a field has been set.
func (o *Business) HasHasAccounts() bool {
	if o != nil && !IsNil(o.HasAccounts) {
		return true
	}

	return false
}

// SetHasAccounts gets a reference to the given bool and assigns it to the HasAccounts field.
func (o *Business) SetHasAccounts(v bool) {
	o.HasAccounts = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Business) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Business) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Business) SetId(v string) {
	o.Id = &v
}

// GetIsCustomer returns the IsCustomer field value
func (o *Business) GetIsCustomer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCustomer
}

// GetIsCustomerOk returns a tuple with the IsCustomer field value
// and a boolean to check if the value has been set.
func (o *Business) GetIsCustomerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCustomer, true
}

// SetIsCustomer sets field value
func (o *Business) SetIsCustomer(v bool) {
	o.IsCustomer = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Business) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Business) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Business) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetLegalAddress returns the LegalAddress field value if set, zero value otherwise.
func (o *Business) GetLegalAddress() LegalAddress {
	if o == nil || IsNil(o.LegalAddress) {
		var ret LegalAddress
		return ret
	}
	return *o.LegalAddress
}

// GetLegalAddressOk returns a tuple with the LegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetLegalAddressOk() (*LegalAddress, bool) {
	if o == nil || IsNil(o.LegalAddress) {
		return nil, false
	}
	return o.LegalAddress, true
}

// HasLegalAddress returns a boolean if a field has been set.
func (o *Business) HasLegalAddress() bool {
	if o != nil && !IsNil(o.LegalAddress) {
		return true
	}

	return false
}

// SetLegalAddress gets a reference to the given LegalAddress and assigns it to the LegalAddress field.
func (o *Business) SetLegalAddress(v LegalAddress) {
	o.LegalAddress = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Business) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Business) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Business) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Business) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Business) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Business) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetStatus returns the Status field value
func (o *Business) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Business) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Business) SetStatus(v string) {
	o.Status = v
}

// GetStructure returns the Structure field value if set, zero value otherwise.
func (o *Business) GetStructure() string {
	if o == nil || IsNil(o.Structure) {
		var ret string
		return ret
	}
	return *o.Structure
}

// GetStructureOk returns a tuple with the Structure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetStructureOk() (*string, bool) {
	if o == nil || IsNil(o.Structure) {
		return nil, false
	}
	return o.Structure, true
}

// HasStructure returns a boolean if a field has been set.
func (o *Business) HasStructure() bool {
	if o != nil && !IsNil(o.Structure) {
		return true
	}

	return false
}

// SetStructure gets a reference to the given string and assigns it to the Structure field.
func (o *Business) SetStructure(v string) {
	o.Structure = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *Business) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *Business) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *Business) SetTenant(v string) {
	o.Tenant = &v
}

// GetTradeNames returns the TradeNames field value if set, zero value otherwise.
func (o *Business) GetTradeNames() []string {
	if o == nil || IsNil(o.TradeNames) {
		var ret []string
		return ret
	}
	return o.TradeNames
}

// GetTradeNamesOk returns a tuple with the TradeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetTradeNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.TradeNames) {
		return nil, false
	}
	return o.TradeNames, true
}

// HasTradeNames returns a boolean if a field has been set.
func (o *Business) HasTradeNames() bool {
	if o != nil && !IsNil(o.TradeNames) {
		return true
	}

	return false
}

// SetTradeNames gets a reference to the given []string and assigns it to the TradeNames field.
func (o *Business) SetTradeNames(v []string) {
	o.TradeNames = v
}

// GetVerificationLastRun returns the VerificationLastRun field value if set, zero value otherwise.
func (o *Business) GetVerificationLastRun() time.Time {
	if o == nil || IsNil(o.VerificationLastRun) {
		var ret time.Time
		return ret
	}
	return *o.VerificationLastRun
}

// GetVerificationLastRunOk returns a tuple with the VerificationLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetVerificationLastRunOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VerificationLastRun) {
		return nil, false
	}
	return o.VerificationLastRun, true
}

// HasVerificationLastRun returns a boolean if a field has been set.
func (o *Business) HasVerificationLastRun() bool {
	if o != nil && !IsNil(o.VerificationLastRun) {
		return true
	}

	return false
}

// SetVerificationLastRun gets a reference to the given time.Time and assigns it to the VerificationLastRun field.
func (o *Business) SetVerificationLastRun(v time.Time) {
	o.VerificationLastRun = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *Business) GetVerificationStatus() VerificationStatus {
	if o == nil || IsNil(o.VerificationStatus) {
		var ret VerificationStatus
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil || IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *Business) HasVerificationStatus() bool {
	if o != nil && !IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given VerificationStatus and assigns it to the VerificationStatus field.
func (o *Business) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *Business) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *Business) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *Business) SetWebsite(v string) {
	o.Website = &v
}

func (o Business) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Business) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComplianceRestrictions) {
		toSerialize["compliance_restrictions"] = o.ComplianceRestrictions
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.Ein) {
		toSerialize["ein"] = o.Ein
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EntityName) {
		toSerialize["entity_name"] = o.EntityName
	}
	if !IsNil(o.FormationDate) {
		toSerialize["formation_date"] = o.FormationDate
	}
	if !IsNil(o.FormationState) {
		toSerialize["formation_state"] = o.FormationState
	}
	if !IsNil(o.HasAccounts) {
		toSerialize["has_accounts"] = o.HasAccounts
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["is_customer"] = o.IsCustomer
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if !IsNil(o.LegalAddress) {
		toSerialize["legal_address"] = o.LegalAddress
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Structure) {
		toSerialize["structure"] = o.Structure
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.TradeNames) {
		toSerialize["trade_names"] = o.TradeNames
	}
	if !IsNil(o.VerificationLastRun) {
		toSerialize["verification_last_run"] = o.VerificationLastRun
	}
	if !IsNil(o.VerificationStatus) {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	return toSerialize, nil
}

type NullableBusiness struct {
	value *Business
	isSet bool
}

func (v NullableBusiness) Get() *Business {
	return v.value
}

func (v *NullableBusiness) Set(val *Business) {
	v.value = val
	v.isSet = true
}

func (v NullableBusiness) IsSet() bool {
	return v.isSet
}

func (v *NullableBusiness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusiness(val *Business) *NullableBusiness {
	return &NullableBusiness{value: val, isSet: true}
}

func (v NullableBusiness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusiness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


