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

// ResendResponse struct for ResendResponse
type ResendResponse struct {
	// response of the event notification
	Body *string `json:"body,omitempty"`
}

// NewResendResponse instantiates a new ResendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResendResponse() *ResendResponse {
	this := ResendResponse{}
	return &this
}

// NewResendResponseWithDefaults instantiates a new ResendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResendResponseWithDefaults() *ResendResponse {
	this := ResendResponse{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ResendResponse) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResendResponse) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ResendResponse) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ResendResponse) SetBody(v string) {
	o.Body = &v
}

func (o ResendResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableResendResponse struct {
	value *ResendResponse
	isSet bool
}

func (v NullableResendResponse) Get() *ResendResponse {
	return v.value
}

func (v *NullableResendResponse) Set(val *ResendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResendResponse(val *ResendResponse) *NullableResendResponse {
	return &NullableResendResponse{value: val, isSet: true}
}

func (v NullableResendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
