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

// ExternalAccountVendorData struct for ExternalAccountVendorData
type ExternalAccountVendorData struct {
	// The last alphanumeric characters of an account's official account number. Note that the mask may be non-unique between accounts, and it may also not match the mask that the bank displays to the user. 
	AccountNumberMask *string `json:"account_number_mask,omitempty"`
	// The ID of the institution external account belongs
	InstitutionId *string `json:"institution_id,omitempty"`
}

// NewExternalAccountVendorData instantiates a new ExternalAccountVendorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccountVendorData() *ExternalAccountVendorData {
	this := ExternalAccountVendorData{}
	return &this
}

// NewExternalAccountVendorDataWithDefaults instantiates a new ExternalAccountVendorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountVendorDataWithDefaults() *ExternalAccountVendorData {
	this := ExternalAccountVendorData{}
	return &this
}

// GetAccountNumberMask returns the AccountNumberMask field value if set, zero value otherwise.
func (o *ExternalAccountVendorData) GetAccountNumberMask() string {
	if o == nil || o.AccountNumberMask == nil {
		var ret string
		return ret
	}
	return *o.AccountNumberMask
}

// GetAccountNumberMaskOk returns a tuple with the AccountNumberMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountVendorData) GetAccountNumberMaskOk() (*string, bool) {
	if o == nil || o.AccountNumberMask == nil {
		return nil, false
	}
	return o.AccountNumberMask, true
}

// HasAccountNumberMask returns a boolean if a field has been set.
func (o *ExternalAccountVendorData) HasAccountNumberMask() bool {
	if o != nil && o.AccountNumberMask != nil {
		return true
	}

	return false
}

// SetAccountNumberMask gets a reference to the given string and assigns it to the AccountNumberMask field.
func (o *ExternalAccountVendorData) SetAccountNumberMask(v string) {
	o.AccountNumberMask = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *ExternalAccountVendorData) GetInstitutionId() string {
	if o == nil || o.InstitutionId == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountVendorData) GetInstitutionIdOk() (*string, bool) {
	if o == nil || o.InstitutionId == nil {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *ExternalAccountVendorData) HasInstitutionId() bool {
	if o != nil && o.InstitutionId != nil {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *ExternalAccountVendorData) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

func (o ExternalAccountVendorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountNumberMask != nil {
		toSerialize["account_number_mask"] = o.AccountNumberMask
	}
	if o.InstitutionId != nil {
		toSerialize["institution_id"] = o.InstitutionId
	}
	return json.Marshal(toSerialize)
}

type NullableExternalAccountVendorData struct {
	value *ExternalAccountVendorData
	isSet bool
}

func (v NullableExternalAccountVendorData) Get() *ExternalAccountVendorData {
	return v.value
}

func (v *NullableExternalAccountVendorData) Set(val *ExternalAccountVendorData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccountVendorData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccountVendorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccountVendorData(val *ExternalAccountVendorData) *NullableExternalAccountVendorData {
	return &NullableExternalAccountVendorData{value: val, isSet: true}
}

func (v NullableExternalAccountVendorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccountVendorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


