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

// checks if the PersonalIdCountryCodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonalIdCountryCodeResponse{}

// PersonalIdCountryCodeResponse struct for PersonalIdCountryCodeResponse
type PersonalIdCountryCodeResponse struct {
	// The ISO 3166 Alpha-2 country code for the country that issued the personal identifier. 
	CountryCode *string `json:"country_code,omitempty"`
}

// NewPersonalIdCountryCodeResponse instantiates a new PersonalIdCountryCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonalIdCountryCodeResponse() *PersonalIdCountryCodeResponse {
	this := PersonalIdCountryCodeResponse{}
	return &this
}

// NewPersonalIdCountryCodeResponseWithDefaults instantiates a new PersonalIdCountryCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonalIdCountryCodeResponseWithDefaults() *PersonalIdCountryCodeResponse {
	this := PersonalIdCountryCodeResponse{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PersonalIdCountryCodeResponse) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalIdCountryCodeResponse) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PersonalIdCountryCodeResponse) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PersonalIdCountryCodeResponse) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o PersonalIdCountryCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonalIdCountryCodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	return toSerialize, nil
}

type NullablePersonalIdCountryCodeResponse struct {
	value *PersonalIdCountryCodeResponse
	isSet bool
}

func (v NullablePersonalIdCountryCodeResponse) Get() *PersonalIdCountryCodeResponse {
	return v.value
}

func (v *NullablePersonalIdCountryCodeResponse) Set(val *PersonalIdCountryCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalIdCountryCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalIdCountryCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalIdCountryCodeResponse(val *PersonalIdCountryCodeResponse) *NullablePersonalIdCountryCodeResponse {
	return &NullablePersonalIdCountryCodeResponse{value: val, isSet: true}
}

func (v NullablePersonalIdCountryCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalIdCountryCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


