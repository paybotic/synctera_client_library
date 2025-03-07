/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the AddVendorAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVendorAccountsResponse{}

// AddVendorAccountsResponse struct for AddVendorAccountsResponse
type AddVendorAccountsResponse struct {
	AddedAccounts        []ExternalAccount         `json:"added_accounts"`
	DeletedAccounts      []ExternalAccount         `json:"deleted_accounts,omitempty"`
	FailedAccounts       []AddVendorAccountFailure `json:"failed_accounts"`
	AdditionalProperties map[string]interface{}
}

type _AddVendorAccountsResponse AddVendorAccountsResponse

// NewAddVendorAccountsResponse instantiates a new AddVendorAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVendorAccountsResponse(addedAccounts []ExternalAccount, failedAccounts []AddVendorAccountFailure) *AddVendorAccountsResponse {
	this := AddVendorAccountsResponse{}
	this.AddedAccounts = addedAccounts
	this.FailedAccounts = failedAccounts
	return &this
}

// NewAddVendorAccountsResponseWithDefaults instantiates a new AddVendorAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVendorAccountsResponseWithDefaults() *AddVendorAccountsResponse {
	this := AddVendorAccountsResponse{}
	return &this
}

// GetAddedAccounts returns the AddedAccounts field value
func (o *AddVendorAccountsResponse) GetAddedAccounts() []ExternalAccount {
	if o == nil {
		var ret []ExternalAccount
		return ret
	}

	return o.AddedAccounts
}

// GetAddedAccountsOk returns a tuple with the AddedAccounts field value
// and a boolean to check if the value has been set.
func (o *AddVendorAccountsResponse) GetAddedAccountsOk() ([]ExternalAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddedAccounts, true
}

// SetAddedAccounts sets field value
func (o *AddVendorAccountsResponse) SetAddedAccounts(v []ExternalAccount) {
	o.AddedAccounts = v
}

// GetDeletedAccounts returns the DeletedAccounts field value if set, zero value otherwise.
func (o *AddVendorAccountsResponse) GetDeletedAccounts() []ExternalAccount {
	if o == nil || IsNil(o.DeletedAccounts) {
		var ret []ExternalAccount
		return ret
	}
	return o.DeletedAccounts
}

// GetDeletedAccountsOk returns a tuple with the DeletedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVendorAccountsResponse) GetDeletedAccountsOk() ([]ExternalAccount, bool) {
	if o == nil || IsNil(o.DeletedAccounts) {
		return nil, false
	}
	return o.DeletedAccounts, true
}

// HasDeletedAccounts returns a boolean if a field has been set.
func (o *AddVendorAccountsResponse) HasDeletedAccounts() bool {
	if o != nil && !IsNil(o.DeletedAccounts) {
		return true
	}

	return false
}

// SetDeletedAccounts gets a reference to the given []ExternalAccount and assigns it to the DeletedAccounts field.
func (o *AddVendorAccountsResponse) SetDeletedAccounts(v []ExternalAccount) {
	o.DeletedAccounts = v
}

// GetFailedAccounts returns the FailedAccounts field value
func (o *AddVendorAccountsResponse) GetFailedAccounts() []AddVendorAccountFailure {
	if o == nil {
		var ret []AddVendorAccountFailure
		return ret
	}

	return o.FailedAccounts
}

// GetFailedAccountsOk returns a tuple with the FailedAccounts field value
// and a boolean to check if the value has been set.
func (o *AddVendorAccountsResponse) GetFailedAccountsOk() ([]AddVendorAccountFailure, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedAccounts, true
}

// SetFailedAccounts sets field value
func (o *AddVendorAccountsResponse) SetFailedAccounts(v []AddVendorAccountFailure) {
	o.FailedAccounts = v
}

func (o AddVendorAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVendorAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["added_accounts"] = o.AddedAccounts
	if !IsNil(o.DeletedAccounts) {
		toSerialize["deleted_accounts"] = o.DeletedAccounts
	}
	toSerialize["failed_accounts"] = o.FailedAccounts

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddVendorAccountsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"added_accounts",
		"failed_accounts",
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

	varAddVendorAccountsResponse := _AddVendorAccountsResponse{}

	err = json.Unmarshal(data, &varAddVendorAccountsResponse)

	if err != nil {
		return err
	}

	*o = AddVendorAccountsResponse(varAddVendorAccountsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "added_accounts")
		delete(additionalProperties, "deleted_accounts")
		delete(additionalProperties, "failed_accounts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddVendorAccountsResponse struct {
	value *AddVendorAccountsResponse
	isSet bool
}

func (v NullableAddVendorAccountsResponse) Get() *AddVendorAccountsResponse {
	return v.value
}

func (v *NullableAddVendorAccountsResponse) Set(val *AddVendorAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVendorAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVendorAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVendorAccountsResponse(val *AddVendorAccountsResponse) *NullableAddVendorAccountsResponse {
	return &NullableAddVendorAccountsResponse{value: val, isSet: true}
}

func (v NullableAddVendorAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVendorAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
