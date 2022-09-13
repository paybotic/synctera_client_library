/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PersonListAllOf struct for PersonListAllOf
type PersonListAllOf struct {
	// Array of persons.
	Persons []Person `json:"persons"`
}

// NewPersonListAllOf instantiates a new PersonListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonListAllOf(persons []Person) *PersonListAllOf {
	this := PersonListAllOf{}
	this.Persons = persons
	return &this
}

// NewPersonListAllOfWithDefaults instantiates a new PersonListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonListAllOfWithDefaults() *PersonListAllOf {
	this := PersonListAllOf{}
	return &this
}

// GetPersons returns the Persons field value
func (o *PersonListAllOf) GetPersons() []Person {
	if o == nil {
		var ret []Person
		return ret
	}

	return o.Persons
}

// GetPersonsOk returns a tuple with the Persons field value
// and a boolean to check if the value has been set.
func (o *PersonListAllOf) GetPersonsOk() ([]Person, bool) {
	if o == nil {
		return nil, false
	}
	return o.Persons, true
}

// SetPersons sets field value
func (o *PersonListAllOf) SetPersons(v []Person) {
	o.Persons = v
}

func (o PersonListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["persons"] = o.Persons
	}
	return json.Marshal(toSerialize)
}

type NullablePersonListAllOf struct {
	value *PersonListAllOf
	isSet bool
}

func (v NullablePersonListAllOf) Get() *PersonListAllOf {
	return v.value
}

func (v *NullablePersonListAllOf) Set(val *PersonListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonListAllOf(val *PersonListAllOf) *NullablePersonListAllOf {
	return &NullablePersonListAllOf{value: val, isSet: true}
}

func (v NullablePersonListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


