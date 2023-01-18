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

// TransactionUpdateMetaRequest struct for TransactionUpdateMetaRequest
type TransactionUpdateMetaRequest struct {
	// an unstructured json blob representing additional transaction information supplied by the integrator.
	ExternalData map[string]interface{} `json:"external_data,omitempty"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data,omitempty"`
}

// NewTransactionUpdateMetaRequest instantiates a new TransactionUpdateMetaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionUpdateMetaRequest() *TransactionUpdateMetaRequest {
	this := TransactionUpdateMetaRequest{}
	return &this
}

// NewTransactionUpdateMetaRequestWithDefaults instantiates a new TransactionUpdateMetaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionUpdateMetaRequestWithDefaults() *TransactionUpdateMetaRequest {
	this := TransactionUpdateMetaRequest{}
	return &this
}

// GetExternalData returns the ExternalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionUpdateMetaRequest) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionUpdateMetaRequest) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ExternalData == nil {
		return nil, false
	}
	return o.ExternalData, true
}

// HasExternalData returns a boolean if a field has been set.
func (o *TransactionUpdateMetaRequest) HasExternalData() bool {
	if o != nil && o.ExternalData != nil {
		return true
	}

	return false
}

// SetExternalData gets a reference to the given map[string]interface{} and assigns it to the ExternalData field.
func (o *TransactionUpdateMetaRequest) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetUserData returns the UserData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionUpdateMetaRequest) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionUpdateMetaRequest) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *TransactionUpdateMetaRequest) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]interface{} and assigns it to the UserData field.
func (o *TransactionUpdateMetaRequest) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

func (o TransactionUpdateMetaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalData != nil {
		toSerialize["external_data"] = o.ExternalData
	}
	if o.UserData != nil {
		toSerialize["user_data"] = o.UserData
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionUpdateMetaRequest struct {
	value *TransactionUpdateMetaRequest
	isSet bool
}

func (v NullableTransactionUpdateMetaRequest) Get() *TransactionUpdateMetaRequest {
	return v.value
}

func (v *NullableTransactionUpdateMetaRequest) Set(val *TransactionUpdateMetaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionUpdateMetaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionUpdateMetaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionUpdateMetaRequest(val *TransactionUpdateMetaRequest) *NullableTransactionUpdateMetaRequest {
	return &NullableTransactionUpdateMetaRequest{value: val, isSet: true}
}

func (v NullableTransactionUpdateMetaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionUpdateMetaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


