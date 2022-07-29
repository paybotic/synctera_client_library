/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ApplicationResponse1 struct for ApplicationResponse1
type ApplicationResponse1 struct {
	CreatedTime time.Time `json:"created_time"`
	// The id of the application from processor
	ExternalId *string `json:"external_id,omitempty"`
	// The id of the application
	Id string `json:"id"`
	LastModifiedTime time.Time `json:"last_modified_time"`
}

// NewApplicationResponse1 instantiates a new ApplicationResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResponse1(createdTime time.Time, id string, lastModifiedTime time.Time) *ApplicationResponse1 {
	this := ApplicationResponse1{}
	this.CreatedTime = createdTime
	this.Id = id
	this.LastModifiedTime = lastModifiedTime
	return &this
}

// NewApplicationResponse1WithDefaults instantiates a new ApplicationResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResponse1WithDefaults() *ApplicationResponse1 {
	this := ApplicationResponse1{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value
func (o *ApplicationResponse1) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponse1) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *ApplicationResponse1) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ApplicationResponse1) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse1) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ApplicationResponse1) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ApplicationResponse1) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value
func (o *ApplicationResponse1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponse1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationResponse1) SetId(v string) {
	o.Id = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *ApplicationResponse1) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponse1) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *ApplicationResponse1) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

func (o ApplicationResponse1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_time"] = o.CreatedTime
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationResponse1 struct {
	value *ApplicationResponse1
	isSet bool
}

func (v NullableApplicationResponse1) Get() *ApplicationResponse1 {
	return v.value
}

func (v *NullableApplicationResponse1) Set(val *ApplicationResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResponse1(val *ApplicationResponse1) *NullableApplicationResponse1 {
	return &NullableApplicationResponse1{value: val, isSet: true}
}

func (v NullableApplicationResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


