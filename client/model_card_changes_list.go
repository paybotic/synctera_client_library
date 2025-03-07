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

// checks if the CardChangesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardChangesList{}

// CardChangesList struct for CardChangesList
type CardChangesList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// List of changes in descending chronological order
	Changes              []CardChange `json:"changes"`
	AdditionalProperties map[string]interface{}
}

type _CardChangesList CardChangesList

// NewCardChangesList instantiates a new CardChangesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardChangesList(changes []CardChange) *CardChangesList {
	this := CardChangesList{}
	this.Changes = changes
	return &this
}

// NewCardChangesListWithDefaults instantiates a new CardChangesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardChangesListWithDefaults() *CardChangesList {
	this := CardChangesList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CardChangesList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardChangesList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CardChangesList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CardChangesList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetChanges returns the Changes field value
func (o *CardChangesList) GetChanges() []CardChange {
	if o == nil {
		var ret []CardChange
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *CardChangesList) GetChangesOk() ([]CardChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *CardChangesList) SetChanges(v []CardChange) {
	o.Changes = v
}

func (o CardChangesList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardChangesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["changes"] = o.Changes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardChangesList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"changes",
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

	varCardChangesList := _CardChangesList{}

	err = json.Unmarshal(data, &varCardChangesList)

	if err != nil {
		return err
	}

	*o = CardChangesList(varCardChangesList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "changes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardChangesList struct {
	value *CardChangesList
	isSet bool
}

func (v NullableCardChangesList) Get() *CardChangesList {
	return v.value
}

func (v *NullableCardChangesList) Set(val *CardChangesList) {
	v.value = val
	v.isSet = true
}

func (v NullableCardChangesList) IsSet() bool {
	return v.isSet
}

func (v *NullableCardChangesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardChangesList(val *CardChangesList) *NullableCardChangesList {
	return &NullableCardChangesList{value: val, isSet: true}
}

func (v NullableCardChangesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardChangesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
