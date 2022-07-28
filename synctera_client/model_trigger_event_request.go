/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TriggerEventRequest struct for TriggerEventRequest
type TriggerEventRequest struct {
	Event *EventType1 `json:"event,omitempty"`
}

// NewTriggerEventRequest instantiates a new TriggerEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerEventRequest() *TriggerEventRequest {
	this := TriggerEventRequest{}
	return &this
}

// NewTriggerEventRequestWithDefaults instantiates a new TriggerEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerEventRequestWithDefaults() *TriggerEventRequest {
	this := TriggerEventRequest{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *TriggerEventRequest) GetEvent() EventType1 {
	if o == nil || o.Event == nil {
		var ret EventType1
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerEventRequest) GetEventOk() (*EventType1, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *TriggerEventRequest) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given EventType1 and assigns it to the Event field.
func (o *TriggerEventRequest) SetEvent(v EventType1) {
	o.Event = &v
}

func (o TriggerEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	return json.Marshal(toSerialize)
}

type NullableTriggerEventRequest struct {
	value *TriggerEventRequest
	isSet bool
}

func (v NullableTriggerEventRequest) Get() *TriggerEventRequest {
	return v.value
}

func (v *NullableTriggerEventRequest) Set(val *TriggerEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerEventRequest(val *TriggerEventRequest) *NullableTriggerEventRequest {
	return &NullableTriggerEventRequest{value: val, isSet: true}
}

func (v NullableTriggerEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


