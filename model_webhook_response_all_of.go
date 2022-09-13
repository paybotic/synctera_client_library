/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// WebhookResponseAllOf struct for WebhookResponseAllOf
type WebhookResponseAllOf struct {
	// The timestamp representing when the webhook request was made
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The timestamp representing when the webhook was last modified
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
}

// NewWebhookResponseAllOf instantiates a new WebhookResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponseAllOf() *WebhookResponseAllOf {
	this := WebhookResponseAllOf{}
	return &this
}

// NewWebhookResponseAllOfWithDefaults instantiates a new WebhookResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseAllOfWithDefaults() *WebhookResponseAllOf {
	this := WebhookResponseAllOf{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *WebhookResponseAllOf) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseAllOf) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *WebhookResponseAllOf) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *WebhookResponseAllOf) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *WebhookResponseAllOf) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseAllOf) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *WebhookResponseAllOf) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *WebhookResponseAllOf) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

func (o WebhookResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookResponseAllOf struct {
	value *WebhookResponseAllOf
	isSet bool
}

func (v NullableWebhookResponseAllOf) Get() *WebhookResponseAllOf {
	return v.value
}

func (v *NullableWebhookResponseAllOf) Set(val *WebhookResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponseAllOf(val *WebhookResponseAllOf) *NullableWebhookResponseAllOf {
	return &NullableWebhookResponseAllOf{value: val, isSet: true}
}

func (v NullableWebhookResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


