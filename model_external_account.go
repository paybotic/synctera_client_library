/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ExternalAccount struct for ExternalAccount
type ExternalAccount struct {
	AccountIdentifiers AccountIdentifiers `json:"account_identifiers"`
	// The names of the account owners. Values may be masked, in which case the array will be empty. 
	AccountOwnerNames []string `json:"account_owner_names"`
	// The identifier for the business customer associated with this external account. Exactly one of `business_id` or `customer_id` will be returned. 
	BusinessId *string `json:"business_id,omitempty"`
	CreationTime time.Time `json:"creation_time"`
	// The identifier for the personal customer associated with this external account. Exactly one of `customer_id` or `business_id` will be returned. 
	CustomerId *string `json:"customer_id,omitempty"`
	// External account unique identifier
	Id string `json:"id"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
	// User-supplied JSON format metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The official name of the account
	Name *string `json:"name,omitempty"`
	// A user-meaningful name for the account
	Nickname NullableString `json:"nickname,omitempty"`
	RoutingIdentifiers AccountRouting `json:"routing_identifiers"`
	// The current state of the account
	Status string `json:"status"`
	// The type of the account
	Type string `json:"type"`
	VendorData *ExternalAccountVendorData `json:"vendor_data,omitempty"`
	VendorInfo *VendorInfo `json:"vendor_info,omitempty"`
	Verification NullableAccountVerification `json:"verification"`
}

// NewExternalAccount instantiates a new ExternalAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccount(accountIdentifiers AccountIdentifiers, accountOwnerNames []string, creationTime time.Time, id string, lastUpdatedTime time.Time, routingIdentifiers AccountRouting, status string, type_ string, verification NullableAccountVerification) *ExternalAccount {
	this := ExternalAccount{}
	this.AccountIdentifiers = accountIdentifiers
	this.AccountOwnerNames = accountOwnerNames
	this.CreationTime = creationTime
	this.Id = id
	this.LastUpdatedTime = lastUpdatedTime
	this.RoutingIdentifiers = routingIdentifiers
	this.Status = status
	this.Type = type_
	this.Verification = verification
	return &this
}

// NewExternalAccountWithDefaults instantiates a new ExternalAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountWithDefaults() *ExternalAccount {
	this := ExternalAccount{}
	return &this
}

// GetAccountIdentifiers returns the AccountIdentifiers field value
func (o *ExternalAccount) GetAccountIdentifiers() AccountIdentifiers {
	if o == nil {
		var ret AccountIdentifiers
		return ret
	}

	return o.AccountIdentifiers
}

// GetAccountIdentifiersOk returns a tuple with the AccountIdentifiers field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetAccountIdentifiersOk() (*AccountIdentifiers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifiers, true
}

// SetAccountIdentifiers sets field value
func (o *ExternalAccount) SetAccountIdentifiers(v AccountIdentifiers) {
	o.AccountIdentifiers = v
}

// GetAccountOwnerNames returns the AccountOwnerNames field value
func (o *ExternalAccount) GetAccountOwnerNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccountOwnerNames
}

// GetAccountOwnerNamesOk returns a tuple with the AccountOwnerNames field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetAccountOwnerNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountOwnerNames, true
}

// SetAccountOwnerNames sets field value
func (o *ExternalAccount) SetAccountOwnerNames(v []string) {
	o.AccountOwnerNames = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *ExternalAccount) GetBusinessId() string {
	if o == nil || o.BusinessId == nil {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetBusinessIdOk() (*string, bool) {
	if o == nil || o.BusinessId == nil {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *ExternalAccount) HasBusinessId() bool {
	if o != nil && o.BusinessId != nil {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *ExternalAccount) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCreationTime returns the CreationTime field value
func (o *ExternalAccount) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *ExternalAccount) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *ExternalAccount) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *ExternalAccount) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *ExternalAccount) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetId returns the Id field value
func (o *ExternalAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalAccount) SetId(v string) {
	o.Id = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *ExternalAccount) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *ExternalAccount) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ExternalAccount) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ExternalAccount) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ExternalAccount) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalAccount) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalAccount) SetName(v string) {
	o.Name = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalAccount) GetNickname() string {
	if o == nil || o.Nickname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalAccount) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *ExternalAccount) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *ExternalAccount) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *ExternalAccount) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *ExternalAccount) UnsetNickname() {
	o.Nickname.Unset()
}

// GetRoutingIdentifiers returns the RoutingIdentifiers field value
func (o *ExternalAccount) GetRoutingIdentifiers() AccountRouting {
	if o == nil {
		var ret AccountRouting
		return ret
	}

	return o.RoutingIdentifiers
}

// GetRoutingIdentifiersOk returns a tuple with the RoutingIdentifiers field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetRoutingIdentifiersOk() (*AccountRouting, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingIdentifiers, true
}

// SetRoutingIdentifiers sets field value
func (o *ExternalAccount) SetRoutingIdentifiers(v AccountRouting) {
	o.RoutingIdentifiers = v
}

// GetStatus returns the Status field value
func (o *ExternalAccount) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExternalAccount) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *ExternalAccount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalAccount) SetType(v string) {
	o.Type = v
}

// GetVendorData returns the VendorData field value if set, zero value otherwise.
func (o *ExternalAccount) GetVendorData() ExternalAccountVendorData {
	if o == nil || o.VendorData == nil {
		var ret ExternalAccountVendorData
		return ret
	}
	return *o.VendorData
}

// GetVendorDataOk returns a tuple with the VendorData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetVendorDataOk() (*ExternalAccountVendorData, bool) {
	if o == nil || o.VendorData == nil {
		return nil, false
	}
	return o.VendorData, true
}

// HasVendorData returns a boolean if a field has been set.
func (o *ExternalAccount) HasVendorData() bool {
	if o != nil && o.VendorData != nil {
		return true
	}

	return false
}

// SetVendorData gets a reference to the given ExternalAccountVendorData and assigns it to the VendorData field.
func (o *ExternalAccount) SetVendorData(v ExternalAccountVendorData) {
	o.VendorData = &v
}

// GetVendorInfo returns the VendorInfo field value if set, zero value otherwise.
func (o *ExternalAccount) GetVendorInfo() VendorInfo {
	if o == nil || o.VendorInfo == nil {
		var ret VendorInfo
		return ret
	}
	return *o.VendorInfo
}

// GetVendorInfoOk returns a tuple with the VendorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetVendorInfoOk() (*VendorInfo, bool) {
	if o == nil || o.VendorInfo == nil {
		return nil, false
	}
	return o.VendorInfo, true
}

// HasVendorInfo returns a boolean if a field has been set.
func (o *ExternalAccount) HasVendorInfo() bool {
	if o != nil && o.VendorInfo != nil {
		return true
	}

	return false
}

// SetVendorInfo gets a reference to the given VendorInfo and assigns it to the VendorInfo field.
func (o *ExternalAccount) SetVendorInfo(v VendorInfo) {
	o.VendorInfo = &v
}

// GetVerification returns the Verification field value
// If the value is explicit nil, the zero value for AccountVerification will be returned
func (o *ExternalAccount) GetVerification() AccountVerification {
	if o == nil || o.Verification.Get() == nil {
		var ret AccountVerification
		return ret
	}

	return *o.Verification.Get()
}

// GetVerificationOk returns a tuple with the Verification field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalAccount) GetVerificationOk() (*AccountVerification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verification.Get(), o.Verification.IsSet()
}

// SetVerification sets field value
func (o *ExternalAccount) SetVerification(v AccountVerification) {
	o.Verification.Set(&v)
}

func (o ExternalAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_identifiers"] = o.AccountIdentifiers
	}
	if true {
		toSerialize["account_owner_names"] = o.AccountOwnerNames
	}
	if o.BusinessId != nil {
		toSerialize["business_id"] = o.BusinessId
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Nickname.IsSet() {
		toSerialize["nickname"] = o.Nickname.Get()
	}
	if true {
		toSerialize["routing_identifiers"] = o.RoutingIdentifiers
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.VendorData != nil {
		toSerialize["vendor_data"] = o.VendorData
	}
	if o.VendorInfo != nil {
		toSerialize["vendor_info"] = o.VendorInfo
	}
	if true {
		toSerialize["verification"] = o.Verification.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExternalAccount struct {
	value *ExternalAccount
	isSet bool
}

func (v NullableExternalAccount) Get() *ExternalAccount {
	return v.value
}

func (v *NullableExternalAccount) Set(val *ExternalAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccount(val *ExternalAccount) *NullableExternalAccount {
	return &NullableExternalAccount{value: val, isSet: true}
}

func (v NullableExternalAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


