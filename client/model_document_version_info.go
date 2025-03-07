/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// checks if the DocumentVersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentVersionInfo{}

// DocumentVersionInfo struct for DocumentVersionInfo
type DocumentVersionInfo struct {
	// The date and time the resource was created
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The file name of the document
	FileName *string `json:"file_name,omitempty"`
	// The date and time the resource was last updated
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Positive integer representing the version of the document
	Version *int32 `json:"version,omitempty"`
}

// NewDocumentVersionInfo instantiates a new DocumentVersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentVersionInfo() *DocumentVersionInfo {
	this := DocumentVersionInfo{}
	return &this
}

// NewDocumentVersionInfoWithDefaults instantiates a new DocumentVersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentVersionInfoWithDefaults() *DocumentVersionInfo {
	this := DocumentVersionInfo{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *DocumentVersionInfo) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentVersionInfo) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *DocumentVersionInfo) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *DocumentVersionInfo) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *DocumentVersionInfo) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentVersionInfo) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *DocumentVersionInfo) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *DocumentVersionInfo) SetFileName(v string) {
	o.FileName = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *DocumentVersionInfo) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentVersionInfo) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *DocumentVersionInfo) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *DocumentVersionInfo) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DocumentVersionInfo) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentVersionInfo) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DocumentVersionInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DocumentVersionInfo) SetVersion(v int32) {
	o.Version = &v
}

func (o DocumentVersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentVersionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableDocumentVersionInfo struct {
	value *DocumentVersionInfo
	isSet bool
}

func (v NullableDocumentVersionInfo) Get() *DocumentVersionInfo {
	return v.value
}

func (v *NullableDocumentVersionInfo) Set(val *DocumentVersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentVersionInfo(val *DocumentVersionInfo) *NullableDocumentVersionInfo {
	return &NullableDocumentVersionInfo{value: val, isSet: true}
}

func (v NullableDocumentVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
