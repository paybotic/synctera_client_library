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

// checks if the PersonalIdCountryCodePost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonalIdCountryCodePost{}

// PersonalIdCountryCodePost struct for PersonalIdCountryCodePost
type PersonalIdCountryCodePost struct {
	// The ISO 3166 Alpha-2 country code for the country that issued the personal identifier. This is optional for personal identifier types that have an implicit country, e.g. SSN. This is required for other types, e.g. PASSPORT_NUMBER.
	CountryCode          *string `json:"country_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PersonalIdCountryCodePost PersonalIdCountryCodePost

// NewPersonalIdCountryCodePost instantiates a new PersonalIdCountryCodePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonalIdCountryCodePost() *PersonalIdCountryCodePost {
	this := PersonalIdCountryCodePost{}
	return &this
}

// NewPersonalIdCountryCodePostWithDefaults instantiates a new PersonalIdCountryCodePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonalIdCountryCodePostWithDefaults() *PersonalIdCountryCodePost {
	this := PersonalIdCountryCodePost{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PersonalIdCountryCodePost) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalIdCountryCodePost) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PersonalIdCountryCodePost) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PersonalIdCountryCodePost) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o PersonalIdCountryCodePost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonalIdCountryCodePost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PersonalIdCountryCodePost) UnmarshalJSON(data []byte) (err error) {
	varPersonalIdCountryCodePost := _PersonalIdCountryCodePost{}

	err = json.Unmarshal(data, &varPersonalIdCountryCodePost)

	if err != nil {
		return err
	}

	*o = PersonalIdCountryCodePost(varPersonalIdCountryCodePost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "country_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePersonalIdCountryCodePost struct {
	value *PersonalIdCountryCodePost
	isSet bool
}

func (v NullablePersonalIdCountryCodePost) Get() *PersonalIdCountryCodePost {
	return v.value
}

func (v *NullablePersonalIdCountryCodePost) Set(val *PersonalIdCountryCodePost) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalIdCountryCodePost) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalIdCountryCodePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalIdCountryCodePost(val *PersonalIdCountryCodePost) *NullablePersonalIdCountryCodePost {
	return &NullablePersonalIdCountryCodePost{value: val, isSet: true}
}

func (v NullablePersonalIdCountryCodePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalIdCountryCodePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
