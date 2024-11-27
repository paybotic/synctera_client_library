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

// checks if the PatchNote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchNote{}

// PatchNote struct for PatchNote
type PatchNote struct {
	Status *NoteStatus `json:"status,omitempty"`
}

// NewPatchNote instantiates a new PatchNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchNote() *PatchNote {
	this := PatchNote{}
	return &this
}

// NewPatchNoteWithDefaults instantiates a new PatchNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchNoteWithDefaults() *PatchNote {
	this := PatchNote{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchNote) GetStatus() NoteStatus {
	if o == nil || IsNil(o.Status) {
		var ret NoteStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchNote) GetStatusOk() (*NoteStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchNote) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NoteStatus and assigns it to the Status field.
func (o *PatchNote) SetStatus(v NoteStatus) {
	o.Status = &v
}

func (o PatchNote) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchNote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePatchNote struct {
	value *PatchNote
	isSet bool
}

func (v NullablePatchNote) Get() *PatchNote {
	return v.value
}

func (v *NullablePatchNote) Set(val *PatchNote) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchNote) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchNote(val *PatchNote) *NullablePatchNote {
	return &NullablePatchNote{value: val, isSet: true}
}

func (v NullablePatchNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
