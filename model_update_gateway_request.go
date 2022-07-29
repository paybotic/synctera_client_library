/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpdateGatewayRequest struct for UpdateGatewayRequest
type UpdateGatewayRequest struct {
	// Current status of the Authorization gateway
	Active *bool `json:"active,omitempty"`
	// Custom Headers of the Authorization gateway
	CustomHeaders *map[string]string `json:"custom_headers,omitempty"`
	// URL of the Authorization gateway
	Url *string `json:"url,omitempty"`
}

// NewUpdateGatewayRequest instantiates a new UpdateGatewayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGatewayRequest() *UpdateGatewayRequest {
	this := UpdateGatewayRequest{}
	return &this
}

// NewUpdateGatewayRequestWithDefaults instantiates a new UpdateGatewayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGatewayRequestWithDefaults() *UpdateGatewayRequest {
	this := UpdateGatewayRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateGatewayRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGatewayRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateGatewayRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateGatewayRequest) SetActive(v bool) {
	o.Active = &v
}

// GetCustomHeaders returns the CustomHeaders field value if set, zero value otherwise.
func (o *UpdateGatewayRequest) GetCustomHeaders() map[string]string {
	if o == nil || o.CustomHeaders == nil {
		var ret map[string]string
		return ret
	}
	return *o.CustomHeaders
}

// GetCustomHeadersOk returns a tuple with the CustomHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGatewayRequest) GetCustomHeadersOk() (*map[string]string, bool) {
	if o == nil || o.CustomHeaders == nil {
		return nil, false
	}
	return o.CustomHeaders, true
}

// HasCustomHeaders returns a boolean if a field has been set.
func (o *UpdateGatewayRequest) HasCustomHeaders() bool {
	if o != nil && o.CustomHeaders != nil {
		return true
	}

	return false
}

// SetCustomHeaders gets a reference to the given map[string]string and assigns it to the CustomHeaders field.
func (o *UpdateGatewayRequest) SetCustomHeaders(v map[string]string) {
	o.CustomHeaders = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateGatewayRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGatewayRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateGatewayRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateGatewayRequest) SetUrl(v string) {
	o.Url = &v
}

func (o UpdateGatewayRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.CustomHeaders != nil {
		toSerialize["custom_headers"] = o.CustomHeaders
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateGatewayRequest struct {
	value *UpdateGatewayRequest
	isSet bool
}

func (v NullableUpdateGatewayRequest) Get() *UpdateGatewayRequest {
	return v.value
}

func (v *NullableUpdateGatewayRequest) Set(val *UpdateGatewayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGatewayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGatewayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGatewayRequest(val *UpdateGatewayRequest) *NullableUpdateGatewayRequest {
	return &NullableUpdateGatewayRequest{value: val, isSet: true}
}

func (v NullableUpdateGatewayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGatewayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


