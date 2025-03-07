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

// checks if the UpdateGatewayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGatewayRequest{}

// UpdateGatewayRequest struct for UpdateGatewayRequest
type UpdateGatewayRequest struct {
	// Current status of the Authorization gateway
	Active *bool `json:"active,omitempty"`
	// List of Card Product unique identifiers that will utilize the Gateway
	CardProducts []string `json:"card_products,omitempty"`
	// Custom Headers of the Authorization gateway
	CustomHeaders *map[string]string `json:"custom_headers,omitempty"`
	Standin       *GatewayStandin    `json:"standin,omitempty"`
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
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGatewayRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateGatewayRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateGatewayRequest) SetActive(v bool) {
	o.Active = &v
}

// GetCardProducts returns the CardProducts field value if set, zero value otherwise.
func (o *UpdateGatewayRequest) GetCardProducts() []string {
	if o == nil || IsNil(o.CardProducts) {
		var ret []string
		return ret
	}
	return o.CardProducts
}

// GetCardProductsOk returns a tuple with the CardProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGatewayRequest) GetCardProductsOk() ([]string, bool) {
	if o == nil || IsNil(o.CardProducts) {
		return nil, false
	}
	return o.CardProducts, true
}

// HasCardProducts returns a boolean if a field has been set.
func (o *UpdateGatewayRequest) HasCardProducts() bool {
	if o != nil && !IsNil(o.CardProducts) {
		return true
	}

	return false
}

// SetCardProducts gets a reference to the given []string and assigns it to the CardProducts field.
func (o *UpdateGatewayRequest) SetCardProducts(v []string) {
	o.CardProducts = v
}

// GetCustomHeaders returns the CustomHeaders field value if set, zero value otherwise.
func (o *UpdateGatewayRequest) GetCustomHeaders() map[string]string {
	if o == nil || IsNil(o.CustomHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.CustomHeaders
}

// GetCustomHeadersOk returns a tuple with the CustomHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGatewayRequest) GetCustomHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomHeaders) {
		return nil, false
	}
	return o.CustomHeaders, true
}

// HasCustomHeaders returns a boolean if a field has been set.
func (o *UpdateGatewayRequest) HasCustomHeaders() bool {
	if o != nil && !IsNil(o.CustomHeaders) {
		return true
	}

	return false
}

// SetCustomHeaders gets a reference to the given map[string]string and assigns it to the CustomHeaders field.
func (o *UpdateGatewayRequest) SetCustomHeaders(v map[string]string) {
	o.CustomHeaders = &v
}

// GetStandin returns the Standin field value if set, zero value otherwise.
func (o *UpdateGatewayRequest) GetStandin() GatewayStandin {
	if o == nil || IsNil(o.Standin) {
		var ret GatewayStandin
		return ret
	}
	return *o.Standin
}

// GetStandinOk returns a tuple with the Standin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGatewayRequest) GetStandinOk() (*GatewayStandin, bool) {
	if o == nil || IsNil(o.Standin) {
		return nil, false
	}
	return o.Standin, true
}

// HasStandin returns a boolean if a field has been set.
func (o *UpdateGatewayRequest) HasStandin() bool {
	if o != nil && !IsNil(o.Standin) {
		return true
	}

	return false
}

// SetStandin gets a reference to the given GatewayStandin and assigns it to the Standin field.
func (o *UpdateGatewayRequest) SetStandin(v GatewayStandin) {
	o.Standin = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateGatewayRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGatewayRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateGatewayRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateGatewayRequest) SetUrl(v string) {
	o.Url = &v
}

func (o UpdateGatewayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGatewayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CardProducts) {
		toSerialize["card_products"] = o.CardProducts
	}
	if !IsNil(o.CustomHeaders) {
		toSerialize["custom_headers"] = o.CustomHeaders
	}
	if !IsNil(o.Standin) {
		toSerialize["standin"] = o.Standin
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
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
