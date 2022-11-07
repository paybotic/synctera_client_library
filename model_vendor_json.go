/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VendorJson struct for VendorJson
type VendorJson struct {
	// Describes the content-type encoding received from the vendor
	ContentType string `json:"content_type"`
	// Data representation in JSON
	Json map[string]interface{} `json:"json"`
	Vendor string `json:"vendor"`
}

// NewVendorJson instantiates a new VendorJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorJson(contentType string, json map[string]interface{}, vendor string) *VendorJson {
	this := VendorJson{}
	this.ContentType = contentType
	this.Json = json
	this.Vendor = vendor
	return &this
}

// NewVendorJsonWithDefaults instantiates a new VendorJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorJsonWithDefaults() *VendorJson {
	this := VendorJson{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *VendorJson) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *VendorJson) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *VendorJson) SetContentType(v string) {
	o.ContentType = v
}

// GetJson returns the Json field value
func (o *VendorJson) GetJson() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Json
}

// GetJsonOk returns a tuple with the Json field value
// and a boolean to check if the value has been set.
func (o *VendorJson) GetJsonOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Json, true
}

// SetJson sets field value
func (o *VendorJson) SetJson(v map[string]interface{}) {
	o.Json = v
}

// GetVendor returns the Vendor field value
func (o *VendorJson) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *VendorJson) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *VendorJson) SetVendor(v string) {
	o.Vendor = v
}

func (o VendorJson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if true {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableVendorJson struct {
	value *VendorJson
	isSet bool
}

func (v NullableVendorJson) Get() *VendorJson {
	return v.value
}

func (v *NullableVendorJson) Set(val *VendorJson) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorJson) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorJson(val *VendorJson) *NullableVendorJson {
	return &NullableVendorJson{value: val, isSet: true}
}

func (v NullableVendorJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


