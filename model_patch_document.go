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

// PatchDocument struct for PatchDocument
type PatchDocument struct {
	// A description of the document
	Description *string `json:"description,omitempty"`
	// A user-friendly name for the document
	Name *string       `json:"name,omitempty"`
	Type *DocumentType `json:"type,omitempty"`
}

// NewPatchDocument instantiates a new PatchDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchDocument() *PatchDocument {
	this := PatchDocument{}
	return &this
}

// NewPatchDocumentWithDefaults instantiates a new PatchDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchDocumentWithDefaults() *PatchDocument {
	this := PatchDocument{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchDocument) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchDocument) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchDocument) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchDocument) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchDocument) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchDocument) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchDocument) GetType() DocumentType {
	if o == nil || o.Type == nil {
		var ret DocumentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetTypeOk() (*DocumentType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchDocument) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given DocumentType and assigns it to the Type field.
func (o *PatchDocument) SetType(v DocumentType) {
	o.Type = &v
}

func (o PatchDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePatchDocument struct {
	value *PatchDocument
	isSet bool
}

func (v NullablePatchDocument) Get() *PatchDocument {
	return v.value
}

func (v *NullablePatchDocument) Set(val *PatchDocument) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchDocument) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchDocument(val *PatchDocument) *NullablePatchDocument {
	return &NullablePatchDocument{value: val, isSet: true}
}

func (v NullablePatchDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
