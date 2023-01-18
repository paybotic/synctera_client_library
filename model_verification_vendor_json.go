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

// VerificationVendorJson struct for VerificationVendorJson
type VerificationVendorJson struct {
	// Describes the content-type encoding received from the vendor.
	ContentType string `json:"content_type"`
	// Array of vendor specific information.
	Details []VerificationVendorInfoDetail `json:"details,omitempty"`
	// Data representation in JSON.
	Json map[string]interface{} `json:"json"`
	// Name of the vendor used.
	Vendor string `json:"vendor"`
}

// NewVerificationVendorJson instantiates a new VerificationVendorJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationVendorJson(contentType string, json map[string]interface{}, vendor string) *VerificationVendorJson {
	this := VerificationVendorJson{}
	this.ContentType = contentType
	this.Json = json
	this.Vendor = vendor
	return &this
}

// NewVerificationVendorJsonWithDefaults instantiates a new VerificationVendorJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationVendorJsonWithDefaults() *VerificationVendorJson {
	this := VerificationVendorJson{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *VerificationVendorJson) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *VerificationVendorJson) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *VerificationVendorJson) SetContentType(v string) {
	o.ContentType = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VerificationVendorJson) GetDetails() []VerificationVendorInfoDetail {
	if o == nil || o.Details == nil {
		var ret []VerificationVendorInfoDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationVendorJson) GetDetailsOk() ([]VerificationVendorInfoDetail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VerificationVendorJson) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []VerificationVendorInfoDetail and assigns it to the Details field.
func (o *VerificationVendorJson) SetDetails(v []VerificationVendorInfoDetail) {
	o.Details = v
}

// GetJson returns the Json field value
func (o *VerificationVendorJson) GetJson() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Json
}

// GetJsonOk returns a tuple with the Json field value
// and a boolean to check if the value has been set.
func (o *VerificationVendorJson) GetJsonOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Json, true
}

// SetJson sets field value
func (o *VerificationVendorJson) SetJson(v map[string]interface{}) {
	o.Json = v
}

// GetVendor returns the Vendor field value
func (o *VerificationVendorJson) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *VerificationVendorJson) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *VerificationVendorJson) SetVendor(v string) {
	o.Vendor = v
}

func (o VerificationVendorJson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if true {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableVerificationVendorJson struct {
	value *VerificationVendorJson
	isSet bool
}

func (v NullableVerificationVendorJson) Get() *VerificationVendorJson {
	return v.value
}

func (v *NullableVerificationVendorJson) Set(val *VerificationVendorJson) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationVendorJson) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationVendorJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationVendorJson(val *VerificationVendorJson) *NullableVerificationVendorJson {
	return &NullableVerificationVendorJson{value: val, isSet: true}
}

func (v NullableVerificationVendorJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationVendorJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


