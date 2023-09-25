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

// checks if the PostPersonalIdsArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPersonalIdsArray{}

// PostPersonalIdsArray struct for PostPersonalIdsArray
type PostPersonalIdsArray struct {
	// > 🚧 Beta > This is a Beta property. Feedback from the community is welcome. We may make breaking changes to this property. Array of personal identifiers 
	PersonalIds []PostPersonalId `json:"personal_ids,omitempty"`
}

// NewPostPersonalIdsArray instantiates a new PostPersonalIdsArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPersonalIdsArray() *PostPersonalIdsArray {
	this := PostPersonalIdsArray{}
	return &this
}

// NewPostPersonalIdsArrayWithDefaults instantiates a new PostPersonalIdsArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPersonalIdsArrayWithDefaults() *PostPersonalIdsArray {
	this := PostPersonalIdsArray{}
	return &this
}

// GetPersonalIds returns the PersonalIds field value if set, zero value otherwise.
func (o *PostPersonalIdsArray) GetPersonalIds() []PostPersonalId {
	if o == nil || IsNil(o.PersonalIds) {
		var ret []PostPersonalId
		return ret
	}
	return o.PersonalIds
}

// GetPersonalIdsOk returns a tuple with the PersonalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPersonalIdsArray) GetPersonalIdsOk() ([]PostPersonalId, bool) {
	if o == nil || IsNil(o.PersonalIds) {
		return nil, false
	}
	return o.PersonalIds, true
}

// HasPersonalIds returns a boolean if a field has been set.
func (o *PostPersonalIdsArray) HasPersonalIds() bool {
	if o != nil && !IsNil(o.PersonalIds) {
		return true
	}

	return false
}

// SetPersonalIds gets a reference to the given []PostPersonalId and assigns it to the PersonalIds field.
func (o *PostPersonalIdsArray) SetPersonalIds(v []PostPersonalId) {
	o.PersonalIds = v
}

func (o PostPersonalIdsArray) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPersonalIdsArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PersonalIds) {
		toSerialize["personal_ids"] = o.PersonalIds
	}
	return toSerialize, nil
}

type NullablePostPersonalIdsArray struct {
	value *PostPersonalIdsArray
	isSet bool
}

func (v NullablePostPersonalIdsArray) Get() *PostPersonalIdsArray {
	return v.value
}

func (v *NullablePostPersonalIdsArray) Set(val *PostPersonalIdsArray) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPersonalIdsArray) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPersonalIdsArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPersonalIdsArray(val *PostPersonalIdsArray) *NullablePostPersonalIdsArray {
	return &NullablePostPersonalIdsArray{value: val, isSet: true}
}

func (v NullablePostPersonalIdsArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPersonalIdsArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


