/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// checks if the PatchExternalAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchExternalAccount{}

// PatchExternalAccount struct for PatchExternalAccount
type PatchExternalAccount struct {
	AccountIdentifiers *PatchAccountsRequestAccountIdentifiers `json:"account_identifiers,omitempty"`
	// The names of the account owners.
	AccountOwnerNames []string `json:"account_owner_names,omitempty"`
	// The currency of the account in ISO 4217 format
	Currency *string `json:"currency,omitempty" validate:"regexp=^[A-Z]{3}$"`
	// A user-meaningful name for the account
	Nickname           *string                                 `json:"nickname,omitempty"`
	RoutingIdentifiers *PatchAccountsRequestRoutingIdentifiers `json:"routing_identifiers,omitempty"`
	// The type of the account
	Type                 *string                     `json:"type,omitempty"`
	Verification         NullableAccountVerification `json:"verification,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchExternalAccount PatchExternalAccount

// NewPatchExternalAccount instantiates a new PatchExternalAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchExternalAccount() *PatchExternalAccount {
	this := PatchExternalAccount{}
	return &this
}

// NewPatchExternalAccountWithDefaults instantiates a new PatchExternalAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchExternalAccountWithDefaults() *PatchExternalAccount {
	this := PatchExternalAccount{}
	return &this
}

// GetAccountIdentifiers returns the AccountIdentifiers field value if set, zero value otherwise.
func (o *PatchExternalAccount) GetAccountIdentifiers() PatchAccountsRequestAccountIdentifiers {
	if o == nil || IsNil(o.AccountIdentifiers) {
		var ret PatchAccountsRequestAccountIdentifiers
		return ret
	}
	return *o.AccountIdentifiers
}

// GetAccountIdentifiersOk returns a tuple with the AccountIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchExternalAccount) GetAccountIdentifiersOk() (*PatchAccountsRequestAccountIdentifiers, bool) {
	if o == nil || IsNil(o.AccountIdentifiers) {
		return nil, false
	}
	return o.AccountIdentifiers, true
}

// HasAccountIdentifiers returns a boolean if a field has been set.
func (o *PatchExternalAccount) HasAccountIdentifiers() bool {
	if o != nil && !IsNil(o.AccountIdentifiers) {
		return true
	}

	return false
}

// SetAccountIdentifiers gets a reference to the given PatchAccountsRequestAccountIdentifiers and assigns it to the AccountIdentifiers field.
func (o *PatchExternalAccount) SetAccountIdentifiers(v PatchAccountsRequestAccountIdentifiers) {
	o.AccountIdentifiers = &v
}

// GetAccountOwnerNames returns the AccountOwnerNames field value if set, zero value otherwise.
func (o *PatchExternalAccount) GetAccountOwnerNames() []string {
	if o == nil || IsNil(o.AccountOwnerNames) {
		var ret []string
		return ret
	}
	return o.AccountOwnerNames
}

// GetAccountOwnerNamesOk returns a tuple with the AccountOwnerNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchExternalAccount) GetAccountOwnerNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AccountOwnerNames) {
		return nil, false
	}
	return o.AccountOwnerNames, true
}

// HasAccountOwnerNames returns a boolean if a field has been set.
func (o *PatchExternalAccount) HasAccountOwnerNames() bool {
	if o != nil && !IsNil(o.AccountOwnerNames) {
		return true
	}

	return false
}

// SetAccountOwnerNames gets a reference to the given []string and assigns it to the AccountOwnerNames field.
func (o *PatchExternalAccount) SetAccountOwnerNames(v []string) {
	o.AccountOwnerNames = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PatchExternalAccount) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchExternalAccount) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PatchExternalAccount) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PatchExternalAccount) SetCurrency(v string) {
	o.Currency = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *PatchExternalAccount) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchExternalAccount) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *PatchExternalAccount) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *PatchExternalAccount) SetNickname(v string) {
	o.Nickname = &v
}

// GetRoutingIdentifiers returns the RoutingIdentifiers field value if set, zero value otherwise.
func (o *PatchExternalAccount) GetRoutingIdentifiers() PatchAccountsRequestRoutingIdentifiers {
	if o == nil || IsNil(o.RoutingIdentifiers) {
		var ret PatchAccountsRequestRoutingIdentifiers
		return ret
	}
	return *o.RoutingIdentifiers
}

// GetRoutingIdentifiersOk returns a tuple with the RoutingIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchExternalAccount) GetRoutingIdentifiersOk() (*PatchAccountsRequestRoutingIdentifiers, bool) {
	if o == nil || IsNil(o.RoutingIdentifiers) {
		return nil, false
	}
	return o.RoutingIdentifiers, true
}

// HasRoutingIdentifiers returns a boolean if a field has been set.
func (o *PatchExternalAccount) HasRoutingIdentifiers() bool {
	if o != nil && !IsNil(o.RoutingIdentifiers) {
		return true
	}

	return false
}

// SetRoutingIdentifiers gets a reference to the given PatchAccountsRequestRoutingIdentifiers and assigns it to the RoutingIdentifiers field.
func (o *PatchExternalAccount) SetRoutingIdentifiers(v PatchAccountsRequestRoutingIdentifiers) {
	o.RoutingIdentifiers = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchExternalAccount) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchExternalAccount) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchExternalAccount) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchExternalAccount) SetType(v string) {
	o.Type = &v
}

// GetVerification returns the Verification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchExternalAccount) GetVerification() AccountVerification {
	if o == nil || IsNil(o.Verification.Get()) {
		var ret AccountVerification
		return ret
	}
	return *o.Verification.Get()
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchExternalAccount) GetVerificationOk() (*AccountVerification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verification.Get(), o.Verification.IsSet()
}

// HasVerification returns a boolean if a field has been set.
func (o *PatchExternalAccount) HasVerification() bool {
	if o != nil && o.Verification.IsSet() {
		return true
	}

	return false
}

// SetVerification gets a reference to the given NullableAccountVerification and assigns it to the Verification field.
func (o *PatchExternalAccount) SetVerification(v AccountVerification) {
	o.Verification.Set(&v)
}

// SetVerificationNil sets the value for Verification to be an explicit nil
func (o *PatchExternalAccount) SetVerificationNil() {
	o.Verification.Set(nil)
}

// UnsetVerification ensures that no value is present for Verification, not even an explicit nil
func (o *PatchExternalAccount) UnsetVerification() {
	o.Verification.Unset()
}

func (o PatchExternalAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchExternalAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountIdentifiers) {
		toSerialize["account_identifiers"] = o.AccountIdentifiers
	}
	if !IsNil(o.AccountOwnerNames) {
		toSerialize["account_owner_names"] = o.AccountOwnerNames
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.RoutingIdentifiers) {
		toSerialize["routing_identifiers"] = o.RoutingIdentifiers
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Verification.IsSet() {
		toSerialize["verification"] = o.Verification.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchExternalAccount) UnmarshalJSON(data []byte) (err error) {
	varPatchExternalAccount := _PatchExternalAccount{}

	err = json.Unmarshal(data, &varPatchExternalAccount)

	if err != nil {
		return err
	}

	*o = PatchExternalAccount(varPatchExternalAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_identifiers")
		delete(additionalProperties, "account_owner_names")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "routing_identifiers")
		delete(additionalProperties, "type")
		delete(additionalProperties, "verification")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchExternalAccount struct {
	value *PatchExternalAccount
	isSet bool
}

func (v NullablePatchExternalAccount) Get() *PatchExternalAccount {
	return v.value
}

func (v *NullablePatchExternalAccount) Set(val *PatchExternalAccount) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchExternalAccount) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchExternalAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchExternalAccount(val *PatchExternalAccount) *NullablePatchExternalAccount {
	return &NullablePatchExternalAccount{value: val, isSet: true}
}

func (v NullablePatchExternalAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchExternalAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
