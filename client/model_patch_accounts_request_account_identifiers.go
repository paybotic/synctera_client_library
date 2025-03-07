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

// checks if the PatchAccountsRequestAccountIdentifiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchAccountsRequestAccountIdentifiers{}

// PatchAccountsRequestAccountIdentifiers struct for PatchAccountsRequestAccountIdentifiers
type PatchAccountsRequestAccountIdentifiers struct {
	// The IBAN of the account. On write, Synctera will store the entire IBAN number; on read, we only return the last 4 characters.
	Iban *string `json:"iban,omitempty"`
	// The account number.
	Number *string `json:"number,omitempty"`
}

// NewPatchAccountsRequestAccountIdentifiers instantiates a new PatchAccountsRequestAccountIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchAccountsRequestAccountIdentifiers() *PatchAccountsRequestAccountIdentifiers {
	this := PatchAccountsRequestAccountIdentifiers{}
	return &this
}

// NewPatchAccountsRequestAccountIdentifiersWithDefaults instantiates a new PatchAccountsRequestAccountIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchAccountsRequestAccountIdentifiersWithDefaults() *PatchAccountsRequestAccountIdentifiers {
	this := PatchAccountsRequestAccountIdentifiers{}
	return &this
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *PatchAccountsRequestAccountIdentifiers) GetIban() string {
	if o == nil || IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountsRequestAccountIdentifiers) GetIbanOk() (*string, bool) {
	if o == nil || IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *PatchAccountsRequestAccountIdentifiers) HasIban() bool {
	if o != nil && !IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *PatchAccountsRequestAccountIdentifiers) SetIban(v string) {
	o.Iban = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *PatchAccountsRequestAccountIdentifiers) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountsRequestAccountIdentifiers) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PatchAccountsRequestAccountIdentifiers) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *PatchAccountsRequestAccountIdentifiers) SetNumber(v string) {
	o.Number = &v
}

func (o PatchAccountsRequestAccountIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchAccountsRequestAccountIdentifiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return toSerialize, nil
}

type NullablePatchAccountsRequestAccountIdentifiers struct {
	value *PatchAccountsRequestAccountIdentifiers
	isSet bool
}

func (v NullablePatchAccountsRequestAccountIdentifiers) Get() *PatchAccountsRequestAccountIdentifiers {
	return v.value
}

func (v *NullablePatchAccountsRequestAccountIdentifiers) Set(val *PatchAccountsRequestAccountIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchAccountsRequestAccountIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchAccountsRequestAccountIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchAccountsRequestAccountIdentifiers(val *PatchAccountsRequestAccountIdentifiers) *NullablePatchAccountsRequestAccountIdentifiers {
	return &NullablePatchAccountsRequestAccountIdentifiers{value: val, isSet: true}
}

func (v NullablePatchAccountsRequestAccountIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchAccountsRequestAccountIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
