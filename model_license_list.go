/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LicenseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseList{}

// LicenseList struct for LicenseList
type LicenseList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of license records
	Licenses []ResponseLicense `json:"licenses,omitempty"`
}

// NewLicenseList instantiates a new LicenseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseList() *LicenseList {
	this := LicenseList{}
	return &this
}

// NewLicenseListWithDefaults instantiates a new LicenseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseListWithDefaults() *LicenseList {
	this := LicenseList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *LicenseList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *LicenseList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *LicenseList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *LicenseList) GetLicenses() []ResponseLicense {
	if o == nil || IsNil(o.Licenses) {
		var ret []ResponseLicense
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseList) GetLicensesOk() ([]ResponseLicense, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *LicenseList) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []ResponseLicense and assigns it to the Licenses field.
func (o *LicenseList) SetLicenses(v []ResponseLicense) {
	o.Licenses = v
}

func (o LicenseList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return toSerialize, nil
}

type NullableLicenseList struct {
	value *LicenseList
	isSet bool
}

func (v NullableLicenseList) Get() *LicenseList {
	return v.value
}

func (v *NullableLicenseList) Set(val *LicenseList) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseList) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseList(val *LicenseList) *NullableLicenseList {
	return &NullableLicenseList{value: val, isSet: true}
}

func (v NullableLicenseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


