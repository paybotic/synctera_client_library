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

// checks if the AddressLines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressLines{}

// AddressLines struct for AddressLines
type AddressLines struct {
	// address line 1 from wire file
	AddressLine1 *string `json:"address_line_1,omitempty"`
	// address line 2 from wire file
	AddressLine2 *string `json:"address_line_2,omitempty"`
	// address line 3 from wire file
	AddressLine3 *string `json:"address_line_3,omitempty"`
}

// NewAddressLines instantiates a new AddressLines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressLines() *AddressLines {
	this := AddressLines{}
	return &this
}

// NewAddressLinesWithDefaults instantiates a new AddressLines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressLinesWithDefaults() *AddressLines {
	this := AddressLines{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *AddressLines) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressLines) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *AddressLines) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *AddressLines) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *AddressLines) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressLines) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *AddressLines) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *AddressLines) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *AddressLines) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressLines) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *AddressLines) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *AddressLines) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

func (o AddressLines) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressLines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressLine1) {
		toSerialize["address_line_1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["address_line_2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["address_line_3"] = o.AddressLine3
	}
	return toSerialize, nil
}

type NullableAddressLines struct {
	value *AddressLines
	isSet bool
}

func (v NullableAddressLines) Get() *AddressLines {
	return v.value
}

func (v *NullableAddressLines) Set(val *AddressLines) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressLines) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressLines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressLines(val *AddressLines) *NullableAddressLines {
	return &NullableAddressLines{value: val, isSet: true}
}

func (v NullableAddressLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}