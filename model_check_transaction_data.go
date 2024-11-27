/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CheckTransactionData Transaction metadata specific to transactions with type `WIRE`
type CheckTransactionData struct {
	// The back image id of the captured check
	BackImageId *string `json:"back_image_id,omitempty"`
	// The front image id of the captured check
	FrontImageId *string `json:"front_image_id,omitempty"`
	// The Synctera check deposit ID
	Id *string `json:"id,omitempty"`
}

// NewCheckTransactionData instantiates a new CheckTransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckTransactionData() *CheckTransactionData {
	this := CheckTransactionData{}
	return &this
}

// NewCheckTransactionDataWithDefaults instantiates a new CheckTransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckTransactionDataWithDefaults() *CheckTransactionData {
	this := CheckTransactionData{}
	return &this
}

// GetBackImageId returns the BackImageId field value if set, zero value otherwise.
func (o *CheckTransactionData) GetBackImageId() string {
	if o == nil || o.BackImageId == nil {
		var ret string
		return ret
	}
	return *o.BackImageId
}

// GetBackImageIdOk returns a tuple with the BackImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckTransactionData) GetBackImageIdOk() (*string, bool) {
	if o == nil || o.BackImageId == nil {
		return nil, false
	}
	return o.BackImageId, true
}

// HasBackImageId returns a boolean if a field has been set.
func (o *CheckTransactionData) HasBackImageId() bool {
	if o != nil && o.BackImageId != nil {
		return true
	}

	return false
}

// SetBackImageId gets a reference to the given string and assigns it to the BackImageId field.
func (o *CheckTransactionData) SetBackImageId(v string) {
	o.BackImageId = &v
}

// GetFrontImageId returns the FrontImageId field value if set, zero value otherwise.
func (o *CheckTransactionData) GetFrontImageId() string {
	if o == nil || o.FrontImageId == nil {
		var ret string
		return ret
	}
	return *o.FrontImageId
}

// GetFrontImageIdOk returns a tuple with the FrontImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckTransactionData) GetFrontImageIdOk() (*string, bool) {
	if o == nil || o.FrontImageId == nil {
		return nil, false
	}
	return o.FrontImageId, true
}

// HasFrontImageId returns a boolean if a field has been set.
func (o *CheckTransactionData) HasFrontImageId() bool {
	if o != nil && o.FrontImageId != nil {
		return true
	}

	return false
}

// SetFrontImageId gets a reference to the given string and assigns it to the FrontImageId field.
func (o *CheckTransactionData) SetFrontImageId(v string) {
	o.FrontImageId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CheckTransactionData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckTransactionData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CheckTransactionData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CheckTransactionData) SetId(v string) {
	o.Id = &v
}

func (o CheckTransactionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackImageId != nil {
		toSerialize["back_image_id"] = o.BackImageId
	}
	if o.FrontImageId != nil {
		toSerialize["front_image_id"] = o.FrontImageId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCheckTransactionData struct {
	value *CheckTransactionData
	isSet bool
}

func (v NullableCheckTransactionData) Get() *CheckTransactionData {
	return v.value
}

func (v *NullableCheckTransactionData) Set(val *CheckTransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckTransactionData(val *CheckTransactionData) *NullableCheckTransactionData {
	return &NullableCheckTransactionData{value: val, isSet: true}
}

func (v NullableCheckTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
