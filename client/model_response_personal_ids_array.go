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

// checks if the ResponsePersonalIdsArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsePersonalIdsArray{}

// ResponsePersonalIdsArray struct for ResponsePersonalIdsArray
type ResponsePersonalIdsArray struct {
	// Array of personal identifiers
	PersonalIds []ResponsePersonalId `json:"personal_ids,omitempty"`
}

// NewResponsePersonalIdsArray instantiates a new ResponsePersonalIdsArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePersonalIdsArray() *ResponsePersonalIdsArray {
	this := ResponsePersonalIdsArray{}
	return &this
}

// NewResponsePersonalIdsArrayWithDefaults instantiates a new ResponsePersonalIdsArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePersonalIdsArrayWithDefaults() *ResponsePersonalIdsArray {
	this := ResponsePersonalIdsArray{}
	return &this
}

// GetPersonalIds returns the PersonalIds field value if set, zero value otherwise.
func (o *ResponsePersonalIdsArray) GetPersonalIds() []ResponsePersonalId {
	if o == nil || IsNil(o.PersonalIds) {
		var ret []ResponsePersonalId
		return ret
	}
	return o.PersonalIds
}

// GetPersonalIdsOk returns a tuple with the PersonalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePersonalIdsArray) GetPersonalIdsOk() ([]ResponsePersonalId, bool) {
	if o == nil || IsNil(o.PersonalIds) {
		return nil, false
	}
	return o.PersonalIds, true
}

// HasPersonalIds returns a boolean if a field has been set.
func (o *ResponsePersonalIdsArray) HasPersonalIds() bool {
	if o != nil && !IsNil(o.PersonalIds) {
		return true
	}

	return false
}

// SetPersonalIds gets a reference to the given []ResponsePersonalId and assigns it to the PersonalIds field.
func (o *ResponsePersonalIdsArray) SetPersonalIds(v []ResponsePersonalId) {
	o.PersonalIds = v
}

func (o ResponsePersonalIdsArray) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsePersonalIdsArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PersonalIds) {
		toSerialize["personal_ids"] = o.PersonalIds
	}
	return toSerialize, nil
}

type NullableResponsePersonalIdsArray struct {
	value *ResponsePersonalIdsArray
	isSet bool
}

func (v NullableResponsePersonalIdsArray) Get() *ResponsePersonalIdsArray {
	return v.value
}

func (v *NullableResponsePersonalIdsArray) Set(val *ResponsePersonalIdsArray) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePersonalIdsArray) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePersonalIdsArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePersonalIdsArray(val *ResponsePersonalIdsArray) *NullableResponsePersonalIdsArray {
	return &NullableResponsePersonalIdsArray{value: val, isSet: true}
}

func (v NullableResponsePersonalIdsArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePersonalIdsArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
