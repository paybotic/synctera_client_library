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

// checks if the LicensePatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicensePatch{}

// LicensePatch struct for LicensePatch
type LicensePatch struct {
	// The entity's license number
	LicenseNumber string `json:"license_number"`
}

// NewLicensePatch instantiates a new LicensePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensePatch(licenseNumber string) *LicensePatch {
	this := LicensePatch{}
	this.LicenseNumber = licenseNumber
	return &this
}

// NewLicensePatchWithDefaults instantiates a new LicensePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensePatchWithDefaults() *LicensePatch {
	this := LicensePatch{}
	return &this
}

// GetLicenseNumber returns the LicenseNumber field value
func (o *LicensePatch) GetLicenseNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseNumber
}

// GetLicenseNumberOk returns a tuple with the LicenseNumber field value
// and a boolean to check if the value has been set.
func (o *LicensePatch) GetLicenseNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseNumber, true
}

// SetLicenseNumber sets field value
func (o *LicensePatch) SetLicenseNumber(v string) {
	o.LicenseNumber = v
}

func (o LicensePatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicensePatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["license_number"] = o.LicenseNumber
	return toSerialize, nil
}

type NullableLicensePatch struct {
	value *LicensePatch
	isSet bool
}

func (v NullableLicensePatch) Get() *LicensePatch {
	return v.value
}

func (v *NullableLicensePatch) Set(val *LicensePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensePatch(val *LicensePatch) *NullableLicensePatch {
	return &NullableLicensePatch{value: val, isSet: true}
}

func (v NullableLicensePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

