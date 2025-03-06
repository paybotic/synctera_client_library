/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the VendorJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorJson{}

// VendorJson struct for VendorJson
type VendorJson struct {
	// Describes the content-type encoding received from the vendor
	ContentType string `json:"content_type"`
	// Data representation in JSON
	Json   map[string]interface{} `json:"json"`
	Vendor string                 `json:"vendor"`
}

type _VendorJson VendorJson

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
		return map[string]interface{}{}, false
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VendorJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content_type"] = o.ContentType
	toSerialize["json"] = o.Json
	toSerialize["vendor"] = o.Vendor
	return toSerialize, nil
}

func (o *VendorJson) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content_type",
		"json",
		"vendor",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVendorJson := _VendorJson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVendorJson)

	if err != nil {
		return err
	}

	*o = VendorJson(varVendorJson)

	return err
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
