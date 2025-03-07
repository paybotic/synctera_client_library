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

// checks if the ExternalAccountsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalAccountsList{}

// ExternalAccountsList struct for ExternalAccountsList
type ExternalAccountsList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of external accounts
	ExternalAccounts     []ExternalAccount `json:"external_accounts"`
	AdditionalProperties map[string]interface{}
}

type _ExternalAccountsList ExternalAccountsList

// NewExternalAccountsList instantiates a new ExternalAccountsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccountsList(externalAccounts []ExternalAccount) *ExternalAccountsList {
	this := ExternalAccountsList{}
	this.ExternalAccounts = externalAccounts
	return &this
}

// NewExternalAccountsListWithDefaults instantiates a new ExternalAccountsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountsListWithDefaults() *ExternalAccountsList {
	this := ExternalAccountsList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ExternalAccountsList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountsList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ExternalAccountsList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ExternalAccountsList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetExternalAccounts returns the ExternalAccounts field value
func (o *ExternalAccountsList) GetExternalAccounts() []ExternalAccount {
	if o == nil {
		var ret []ExternalAccount
		return ret
	}

	return o.ExternalAccounts
}

// GetExternalAccountsOk returns a tuple with the ExternalAccounts field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountsList) GetExternalAccountsOk() ([]ExternalAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalAccounts, true
}

// SetExternalAccounts sets field value
func (o *ExternalAccountsList) SetExternalAccounts(v []ExternalAccount) {
	o.ExternalAccounts = v
}

func (o ExternalAccountsList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalAccountsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["external_accounts"] = o.ExternalAccounts

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExternalAccountsList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"external_accounts",
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

	varExternalAccountsList := _ExternalAccountsList{}

	err = json.Unmarshal(data, &varExternalAccountsList)

	if err != nil {
		return err
	}

	*o = ExternalAccountsList(varExternalAccountsList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "external_accounts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalAccountsList struct {
	value *ExternalAccountsList
	isSet bool
}

func (v NullableExternalAccountsList) Get() *ExternalAccountsList {
	return v.value
}

func (v *NullableExternalAccountsList) Set(val *ExternalAccountsList) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccountsList) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccountsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccountsList(val *ExternalAccountsList) *NullableExternalAccountsList {
	return &NullableExternalAccountsList{value: val, isSet: true}
}

func (v NullableExternalAccountsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccountsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
