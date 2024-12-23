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

// Alias struct for Alias
type Alias struct {
	// Account ID
	AccountId *string `json:"account_id,omitempty"`
	// Account number
	AccountNumber *string `json:"account_number,omitempty"`
	// Additional information for the alias
	AliasInfo map[string]interface{} `json:"alias_info,omitempty"`
	// Alias name
	AliasName *string `json:"alias_name,omitempty"`
	// Alias source
	AliasSource *string `json:"alias_source,omitempty"`
	// Alias type
	AliasType *string `json:"alias_type,omitempty"`
	// Alias ID
	Id *string `json:"id,omitempty"`
}

// NewAlias instantiates a new Alias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlias() *Alias {
	this := Alias{}
	return &this
}

// NewAliasWithDefaults instantiates a new Alias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasWithDefaults() *Alias {
	this := Alias{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Alias) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Alias) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Alias) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *Alias) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *Alias) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *Alias) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAliasInfo returns the AliasInfo field value if set, zero value otherwise.
func (o *Alias) GetAliasInfo() map[string]interface{} {
	if o == nil || o.AliasInfo == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AliasInfo
}

// GetAliasInfoOk returns a tuple with the AliasInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetAliasInfoOk() (map[string]interface{}, bool) {
	if o == nil || o.AliasInfo == nil {
		return nil, false
	}
	return o.AliasInfo, true
}

// HasAliasInfo returns a boolean if a field has been set.
func (o *Alias) HasAliasInfo() bool {
	if o != nil && o.AliasInfo != nil {
		return true
	}

	return false
}

// SetAliasInfo gets a reference to the given map[string]interface{} and assigns it to the AliasInfo field.
func (o *Alias) SetAliasInfo(v map[string]interface{}) {
	o.AliasInfo = v
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *Alias) GetAliasName() string {
	if o == nil || o.AliasName == nil {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetAliasNameOk() (*string, bool) {
	if o == nil || o.AliasName == nil {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *Alias) HasAliasName() bool {
	if o != nil && o.AliasName != nil {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *Alias) SetAliasName(v string) {
	o.AliasName = &v
}

// GetAliasSource returns the AliasSource field value if set, zero value otherwise.
func (o *Alias) GetAliasSource() string {
	if o == nil || o.AliasSource == nil {
		var ret string
		return ret
	}
	return *o.AliasSource
}

// GetAliasSourceOk returns a tuple with the AliasSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetAliasSourceOk() (*string, bool) {
	if o == nil || o.AliasSource == nil {
		return nil, false
	}
	return o.AliasSource, true
}

// HasAliasSource returns a boolean if a field has been set.
func (o *Alias) HasAliasSource() bool {
	if o != nil && o.AliasSource != nil {
		return true
	}

	return false
}

// SetAliasSource gets a reference to the given string and assigns it to the AliasSource field.
func (o *Alias) SetAliasSource(v string) {
	o.AliasSource = &v
}

// GetAliasType returns the AliasType field value if set, zero value otherwise.
func (o *Alias) GetAliasType() string {
	if o == nil || o.AliasType == nil {
		var ret string
		return ret
	}
	return *o.AliasType
}

// GetAliasTypeOk returns a tuple with the AliasType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetAliasTypeOk() (*string, bool) {
	if o == nil || o.AliasType == nil {
		return nil, false
	}
	return o.AliasType, true
}

// HasAliasType returns a boolean if a field has been set.
func (o *Alias) HasAliasType() bool {
	if o != nil && o.AliasType != nil {
		return true
	}

	return false
}

// SetAliasType gets a reference to the given string and assigns it to the AliasType field.
func (o *Alias) SetAliasType(v string) {
	o.AliasType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Alias) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Alias) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Alias) SetId(v string) {
	o.Id = &v
}

func (o Alias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.AccountNumber != nil {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.AliasInfo != nil {
		toSerialize["alias_info"] = o.AliasInfo
	}
	if o.AliasName != nil {
		toSerialize["alias_name"] = o.AliasName
	}
	if o.AliasSource != nil {
		toSerialize["alias_source"] = o.AliasSource
	}
	if o.AliasType != nil {
		toSerialize["alias_type"] = o.AliasType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAlias struct {
	value *Alias
	isSet bool
}

func (v NullableAlias) Get() *Alias {
	return v.value
}

func (v *NullableAlias) Set(val *Alias) {
	v.value = val
	v.isSet = true
}

func (v NullableAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlias(val *Alias) *NullableAlias {
	return &NullableAlias{value: val, isSet: true}
}

func (v NullableAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
