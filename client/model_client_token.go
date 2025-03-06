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

// checks if the ClientToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientToken{}

// ClientToken A short-lived, one-time token used for accessing client PINs/PANs
type ClientToken struct {
	ClientToken          *string `json:"client_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientToken ClientToken

// NewClientToken instantiates a new ClientToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientToken() *ClientToken {
	this := ClientToken{}
	return &this
}

// NewClientTokenWithDefaults instantiates a new ClientToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientTokenWithDefaults() *ClientToken {
	this := ClientToken{}
	return &this
}

// GetClientToken returns the ClientToken field value if set, zero value otherwise.
func (o *ClientToken) GetClientToken() string {
	if o == nil || IsNil(o.ClientToken) {
		var ret string
		return ret
	}
	return *o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetClientTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ClientToken) {
		return nil, false
	}
	return o.ClientToken, true
}

// HasClientToken returns a boolean if a field has been set.
func (o *ClientToken) HasClientToken() bool {
	if o != nil && !IsNil(o.ClientToken) {
		return true
	}

	return false
}

// SetClientToken gets a reference to the given string and assigns it to the ClientToken field.
func (o *ClientToken) SetClientToken(v string) {
	o.ClientToken = &v
}

func (o ClientToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientToken) {
		toSerialize["client_token"] = o.ClientToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientToken) UnmarshalJSON(data []byte) (err error) {
	varClientToken := _ClientToken{}

	err = json.Unmarshal(data, &varClientToken)

	if err != nil {
		return err
	}

	*o = ClientToken(varClientToken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientToken struct {
	value *ClientToken
	isSet bool
}

func (v NullableClientToken) Get() *ClientToken {
	return v.value
}

func (v *NullableClientToken) Set(val *ClientToken) {
	v.value = val
	v.isSet = true
}

func (v NullableClientToken) IsSet() bool {
	return v.isSet
}

func (v *NullableClientToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientToken(val *ClientToken) *NullableClientToken {
	return &NullableClientToken{value: val, isSet: true}
}

func (v NullableClientToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
