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

// checks if the GatewayListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayListResponse{}

// GatewayListResponse struct for GatewayListResponse
type GatewayListResponse struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of Authorization gateway configuration
	Gateways []GatewayResponse `json:"gateways,omitempty"`
}

// NewGatewayListResponse instantiates a new GatewayListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayListResponse() *GatewayListResponse {
	this := GatewayListResponse{}
	return &this
}

// NewGatewayListResponseWithDefaults instantiates a new GatewayListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayListResponseWithDefaults() *GatewayListResponse {
	this := GatewayListResponse{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *GatewayListResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *GatewayListResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *GatewayListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *GatewayListResponse) GetGateways() []GatewayResponse {
	if o == nil || IsNil(o.Gateways) {
		var ret []GatewayResponse
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListResponse) GetGatewaysOk() ([]GatewayResponse, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *GatewayListResponse) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []GatewayResponse and assigns it to the Gateways field.
func (o *GatewayListResponse) SetGateways(v []GatewayResponse) {
	o.Gateways = v
}

func (o GatewayListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	return toSerialize, nil
}

type NullableGatewayListResponse struct {
	value *GatewayListResponse
	isSet bool
}

func (v NullableGatewayListResponse) Get() *GatewayListResponse {
	return v.value
}

func (v *NullableGatewayListResponse) Set(val *GatewayListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayListResponse(val *GatewayListResponse) *NullableGatewayListResponse {
	return &NullableGatewayListResponse{value: val, isSet: true}
}

func (v NullableGatewayListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


