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

// ProspectEditable struct for ProspectEditable
type ProspectEditable struct {
	// First Name
	FirstName *string `json:"first_name,omitempty"`
	// Last Name
	LastName *string `json:"last_name,omitempty"`
	// Client supplied json metadata to be stored with the prospect
	Metadata *string `json:"metadata,omitempty"`
	// Source of prospect
	Source *string         `json:"source,omitempty"`
	Status *ProspectStatus `json:"status,omitempty"`
}

// NewProspectEditable instantiates a new ProspectEditable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProspectEditable() *ProspectEditable {
	this := ProspectEditable{}
	return &this
}

// NewProspectEditableWithDefaults instantiates a new ProspectEditable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProspectEditableWithDefaults() *ProspectEditable {
	this := ProspectEditable{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ProspectEditable) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProspectEditable) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ProspectEditable) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ProspectEditable) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ProspectEditable) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProspectEditable) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ProspectEditable) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ProspectEditable) SetLastName(v string) {
	o.LastName = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ProspectEditable) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProspectEditable) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ProspectEditable) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *ProspectEditable) SetMetadata(v string) {
	o.Metadata = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProspectEditable) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProspectEditable) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProspectEditable) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ProspectEditable) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProspectEditable) GetStatus() ProspectStatus {
	if o == nil || o.Status == nil {
		var ret ProspectStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProspectEditable) GetStatusOk() (*ProspectStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProspectEditable) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ProspectStatus and assigns it to the Status field.
func (o *ProspectEditable) SetStatus(v ProspectStatus) {
	o.Status = &v
}

func (o ProspectEditable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableProspectEditable struct {
	value *ProspectEditable
	isSet bool
}

func (v NullableProspectEditable) Get() *ProspectEditable {
	return v.value
}

func (v *NullableProspectEditable) Set(val *ProspectEditable) {
	v.value = val
	v.isSet = true
}

func (v NullableProspectEditable) IsSet() bool {
	return v.isSet
}

func (v *NullableProspectEditable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProspectEditable(val *ProspectEditable) *NullableProspectEditable {
	return &NullableProspectEditable{value: val, isSet: true}
}

func (v NullableProspectEditable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProspectEditable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
