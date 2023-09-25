/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ReplaceSecret200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceSecret200Response{}

// ReplaceSecret200Response struct for ReplaceSecret200Response
type ReplaceSecret200Response struct {
	// Timestamp that the old secret is delete
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Generated secret. Do not share. This secret will be used to verify that webhook requests were sent from Synctera.
	Secret *string `json:"secret,omitempty"`
}

// NewReplaceSecret200Response instantiates a new ReplaceSecret200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceSecret200Response() *ReplaceSecret200Response {
	this := ReplaceSecret200Response{}
	return &this
}

// NewReplaceSecret200ResponseWithDefaults instantiates a new ReplaceSecret200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceSecret200ResponseWithDefaults() *ReplaceSecret200Response {
	this := ReplaceSecret200Response{}
	return &this
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ReplaceSecret200Response) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceSecret200Response) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ReplaceSecret200Response) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *ReplaceSecret200Response) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ReplaceSecret200Response) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceSecret200Response) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ReplaceSecret200Response) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ReplaceSecret200Response) SetSecret(v string) {
	o.Secret = &v
}

func (o ReplaceSecret200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceSecret200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullableReplaceSecret200Response struct {
	value *ReplaceSecret200Response
	isSet bool
}

func (v NullableReplaceSecret200Response) Get() *ReplaceSecret200Response {
	return v.value
}

func (v *NullableReplaceSecret200Response) Set(val *ReplaceSecret200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceSecret200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceSecret200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceSecret200Response(val *ReplaceSecret200Response) *NullableReplaceSecret200Response {
	return &NullableReplaceSecret200Response{value: val, isSet: true}
}

func (v NullableReplaceSecret200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceSecret200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


