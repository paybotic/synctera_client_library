/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Workspace struct for Workspace
type Workspace struct {
	// Bank ID
	BankId int32 `json:"bank_id"`
	// First Name
	BankName    string      `json:"bank_name"`
	Environment Environment `json:"environment"`
	// Partner ID
	PartnerId int32 `json:"partner_id"`
	// Last Name
	PartnerName string `json:"partner_name"`
	// Each workspace has a rank. The highest-ranked (lowest numerical value) workspace is intended to be presented first within its environment.
	Rank int32 `json:"rank"`
	// The id of the tenant containing the resource.
	Tenant             *string `json:"tenant,omitempty"`
	VerificationStatus string  `json:"verification_status"`
}

// NewWorkspace instantiates a new Workspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspace(bankId int32, bankName string, environment Environment, partnerId int32, partnerName string, rank int32, verificationStatus string) *Workspace {
	this := Workspace{}
	this.BankId = bankId
	this.BankName = bankName
	this.Environment = environment
	this.PartnerId = partnerId
	this.PartnerName = partnerName
	this.Rank = rank
	this.VerificationStatus = verificationStatus
	return &this
}

// NewWorkspaceWithDefaults instantiates a new Workspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceWithDefaults() *Workspace {
	this := Workspace{}
	return &this
}

// GetBankId returns the BankId field value
func (o *Workspace) GetBankId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetBankIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankId, true
}

// SetBankId sets field value
func (o *Workspace) SetBankId(v int32) {
	o.BankId = v
}

// GetBankName returns the BankName field value
func (o *Workspace) GetBankName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankName, true
}

// SetBankName sets field value
func (o *Workspace) SetBankName(v string) {
	o.BankName = v
}

// GetEnvironment returns the Environment field value
func (o *Workspace) GetEnvironment() Environment {
	if o == nil {
		var ret Environment
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetEnvironmentOk() (*Environment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *Workspace) SetEnvironment(v Environment) {
	o.Environment = v
}

// GetPartnerId returns the PartnerId field value
func (o *Workspace) GetPartnerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetPartnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerId, true
}

// SetPartnerId sets field value
func (o *Workspace) SetPartnerId(v int32) {
	o.PartnerId = v
}

// GetPartnerName returns the PartnerName field value
func (o *Workspace) GetPartnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerName
}

// GetPartnerNameOk returns a tuple with the PartnerName field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetPartnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerName, true
}

// SetPartnerName sets field value
func (o *Workspace) SetPartnerName(v string) {
	o.PartnerName = v
}

// GetRank returns the Rank field value
func (o *Workspace) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *Workspace) SetRank(v int32) {
	o.Rank = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *Workspace) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *Workspace) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *Workspace) SetTenant(v string) {
	o.Tenant = &v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *Workspace) GetVerificationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetVerificationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *Workspace) SetVerificationStatus(v string) {
	o.VerificationStatus = v
}

func (o Workspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_id"] = o.BankId
	}
	if true {
		toSerialize["bank_name"] = o.BankName
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId
	}
	if true {
		toSerialize["partner_name"] = o.PartnerName
	}
	if true {
		toSerialize["rank"] = o.Rank
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspace struct {
	value *Workspace
	isSet bool
}

func (v NullableWorkspace) Get() *Workspace {
	return v.value
}

func (v *NullableWorkspace) Set(val *Workspace) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspace(val *Workspace) *NullableWorkspace {
	return &NullableWorkspace{value: val, isSet: true}
}

func (v NullableWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
