/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// checks if the EftCaPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EftCaPatch{}

// EftCaPatch Properties for updating a transfer
type EftCaPatch struct {
	// Additional information to be added to the transfer
	SourceData map[string]interface{} `json:"source_data,omitempty"`
	Status     *string                `json:"status,omitempty"`
}

// NewEftCaPatch instantiates a new EftCaPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEftCaPatch() *EftCaPatch {
	this := EftCaPatch{}
	return &this
}

// NewEftCaPatchWithDefaults instantiates a new EftCaPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEftCaPatchWithDefaults() *EftCaPatch {
	this := EftCaPatch{}
	return &this
}

// GetSourceData returns the SourceData field value if set, zero value otherwise.
func (o *EftCaPatch) GetSourceData() map[string]interface{} {
	if o == nil || IsNil(o.SourceData) {
		var ret map[string]interface{}
		return ret
	}
	return o.SourceData
}

// GetSourceDataOk returns a tuple with the SourceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftCaPatch) GetSourceDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SourceData) {
		return map[string]interface{}{}, false
	}
	return o.SourceData, true
}

// HasSourceData returns a boolean if a field has been set.
func (o *EftCaPatch) HasSourceData() bool {
	if o != nil && !IsNil(o.SourceData) {
		return true
	}

	return false
}

// SetSourceData gets a reference to the given map[string]interface{} and assigns it to the SourceData field.
func (o *EftCaPatch) SetSourceData(v map[string]interface{}) {
	o.SourceData = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EftCaPatch) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftCaPatch) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EftCaPatch) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EftCaPatch) SetStatus(v string) {
	o.Status = &v
}

func (o EftCaPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EftCaPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceData) {
		toSerialize["source_data"] = o.SourceData
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableEftCaPatch struct {
	value *EftCaPatch
	isSet bool
}

func (v NullableEftCaPatch) Get() *EftCaPatch {
	return v.value
}

func (v *NullableEftCaPatch) Set(val *EftCaPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableEftCaPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableEftCaPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEftCaPatch(val *EftCaPatch) *NullableEftCaPatch {
	return &NullableEftCaPatch{value: val, isSet: true}
}

func (v NullableEftCaPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEftCaPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
