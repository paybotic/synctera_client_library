/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetCardBarcode200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCardBarcode200Response{}

// GetCardBarcode200Response struct for GetCardBarcode200Response
type GetCardBarcode200Response struct {
	// Barcode of the card
	Barcode *string `json:"barcode,omitempty"`
}

// NewGetCardBarcode200Response instantiates a new GetCardBarcode200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCardBarcode200Response() *GetCardBarcode200Response {
	this := GetCardBarcode200Response{}
	return &this
}

// NewGetCardBarcode200ResponseWithDefaults instantiates a new GetCardBarcode200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCardBarcode200ResponseWithDefaults() *GetCardBarcode200Response {
	this := GetCardBarcode200Response{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *GetCardBarcode200Response) GetBarcode() string {
	if o == nil || IsNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCardBarcode200Response) GetBarcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *GetCardBarcode200Response) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *GetCardBarcode200Response) SetBarcode(v string) {
	o.Barcode = &v
}

func (o GetCardBarcode200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCardBarcode200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	return toSerialize, nil
}

type NullableGetCardBarcode200Response struct {
	value *GetCardBarcode200Response
	isSet bool
}

func (v NullableGetCardBarcode200Response) Get() *GetCardBarcode200Response {
	return v.value
}

func (v *NullableGetCardBarcode200Response) Set(val *GetCardBarcode200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCardBarcode200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCardBarcode200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCardBarcode200Response(val *GetCardBarcode200Response) *NullableGetCardBarcode200Response {
	return &NullableGetCardBarcode200Response{value: val, isSet: true}
}

func (v NullableGetCardBarcode200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCardBarcode200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
