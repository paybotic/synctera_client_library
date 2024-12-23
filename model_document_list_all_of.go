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

// DocumentListAllOf struct for DocumentListAllOf
type DocumentListAllOf struct {
	// Array of documents
	Documents []Document `json:"documents"`
}

// NewDocumentListAllOf instantiates a new DocumentListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentListAllOf(documents []Document) *DocumentListAllOf {
	this := DocumentListAllOf{}
	this.Documents = documents
	return &this
}

// NewDocumentListAllOfWithDefaults instantiates a new DocumentListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentListAllOfWithDefaults() *DocumentListAllOf {
	this := DocumentListAllOf{}
	return &this
}

// GetDocuments returns the Documents field value
func (o *DocumentListAllOf) GetDocuments() []Document {
	if o == nil {
		var ret []Document
		return ret
	}

	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value
// and a boolean to check if the value has been set.
func (o *DocumentListAllOf) GetDocumentsOk() ([]Document, bool) {
	if o == nil {
		return nil, false
	}
	return o.Documents, true
}

// SetDocuments sets field value
func (o *DocumentListAllOf) SetDocuments(v []Document) {
	o.Documents = v
}

func (o DocumentListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["documents"] = o.Documents
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentListAllOf struct {
	value *DocumentListAllOf
	isSet bool
}

func (v NullableDocumentListAllOf) Get() *DocumentListAllOf {
	return v.value
}

func (v *NullableDocumentListAllOf) Set(val *DocumentListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentListAllOf(val *DocumentListAllOf) *NullableDocumentListAllOf {
	return &NullableDocumentListAllOf{value: val, isSet: true}
}

func (v NullableDocumentListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
