/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PersonalIdCustomerId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonalIdCustomerId{}

// PersonalIdCustomerId struct for PersonalIdCustomerId
type PersonalIdCustomerId struct {
	// Id of the customer having this personal identifier
	CustomerId *string `json:"customer_id,omitempty"`
}

// NewPersonalIdCustomerId instantiates a new PersonalIdCustomerId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonalIdCustomerId() *PersonalIdCustomerId {
	this := PersonalIdCustomerId{}
	return &this
}

// NewPersonalIdCustomerIdWithDefaults instantiates a new PersonalIdCustomerId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonalIdCustomerIdWithDefaults() *PersonalIdCustomerId {
	this := PersonalIdCustomerId{}
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *PersonalIdCustomerId) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalIdCustomerId) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PersonalIdCustomerId) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *PersonalIdCustomerId) SetCustomerId(v string) {
	o.CustomerId = &v
}

func (o PersonalIdCustomerId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonalIdCustomerId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	return toSerialize, nil
}

type NullablePersonalIdCustomerId struct {
	value *PersonalIdCustomerId
	isSet bool
}

func (v NullablePersonalIdCustomerId) Get() *PersonalIdCustomerId {
	return v.value
}

func (v *NullablePersonalIdCustomerId) Set(val *PersonalIdCustomerId) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalIdCustomerId) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalIdCustomerId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalIdCustomerId(val *PersonalIdCustomerId) *NullablePersonalIdCustomerId {
	return &NullablePersonalIdCustomerId{value: val, isSet: true}
}

func (v NullablePersonalIdCustomerId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalIdCustomerId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}