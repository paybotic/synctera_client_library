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

// checks if the InternationalWireList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternationalWireList{}

// InternationalWireList struct for InternationalWireList
type InternationalWireList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of international wires
	InternationalWires []InternationalWireResponse `json:"international_wires"`
}

// NewInternationalWireList instantiates a new InternationalWireList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternationalWireList(internationalWires []InternationalWireResponse) *InternationalWireList {
	this := InternationalWireList{}
	this.InternationalWires = internationalWires
	return &this
}

// NewInternationalWireListWithDefaults instantiates a new InternationalWireList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternationalWireListWithDefaults() *InternationalWireList {
	this := InternationalWireList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *InternationalWireList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalWireList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *InternationalWireList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *InternationalWireList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetInternationalWires returns the InternationalWires field value
func (o *InternationalWireList) GetInternationalWires() []InternationalWireResponse {
	if o == nil {
		var ret []InternationalWireResponse
		return ret
	}

	return o.InternationalWires
}

// GetInternationalWiresOk returns a tuple with the InternationalWires field value
// and a boolean to check if the value has been set.
func (o *InternationalWireList) GetInternationalWiresOk() ([]InternationalWireResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternationalWires, true
}

// SetInternationalWires sets field value
func (o *InternationalWireList) SetInternationalWires(v []InternationalWireResponse) {
	o.InternationalWires = v
}

func (o InternationalWireList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternationalWireList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["international_wires"] = o.InternationalWires
	return toSerialize, nil
}

type NullableInternationalWireList struct {
	value *InternationalWireList
	isSet bool
}

func (v NullableInternationalWireList) Get() *InternationalWireList {
	return v.value
}

func (v *NullableInternationalWireList) Set(val *InternationalWireList) {
	v.value = val
	v.isSet = true
}

func (v NullableInternationalWireList) IsSet() bool {
	return v.isSet
}

func (v *NullableInternationalWireList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternationalWireList(val *InternationalWireList) *NullableInternationalWireList {
	return &NullableInternationalWireList{value: val, isSet: true}
}

func (v NullableInternationalWireList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternationalWireList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


