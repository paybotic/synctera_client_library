/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the NoteList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoteList{}

// NoteList struct for NoteList
type NoteList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of notes
	Notes                []NoteResponse `json:"notes"`
	AdditionalProperties map[string]interface{}
}

type _NoteList NoteList

// NewNoteList instantiates a new NoteList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteList(notes []NoteResponse) *NoteList {
	this := NoteList{}
	this.Notes = notes
	return &this
}

// NewNoteListWithDefaults instantiates a new NoteList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteListWithDefaults() *NoteList {
	this := NoteList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *NoteList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *NoteList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *NoteList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetNotes returns the Notes field value
func (o *NoteList) GetNotes() []NoteResponse {
	if o == nil {
		var ret []NoteResponse
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *NoteList) GetNotesOk() ([]NoteResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *NoteList) SetNotes(v []NoteResponse) {
	o.Notes = v
}

func (o NoteList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoteList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["notes"] = o.Notes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NoteList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notes",
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

	varNoteList := _NoteList{}

	err = json.Unmarshal(data, &varNoteList)

	if err != nil {
		return err
	}

	*o = NoteList(varNoteList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "notes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNoteList struct {
	value *NoteList
	isSet bool
}

func (v NullableNoteList) Get() *NoteList {
	return v.value
}

func (v *NullableNoteList) Set(val *NoteList) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteList) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteList(val *NoteList) *NullableNoteList {
	return &NullableNoteList{value: val, isSet: true}
}

func (v NullableNoteList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
