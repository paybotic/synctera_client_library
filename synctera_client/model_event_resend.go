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

// EventResend struct for EventResend
type EventResend struct {
	// id of the event notification you want to resend
	EventId string `json:"event_id"`
	EventType EventType `json:"event_type"`
}

// NewEventResend instantiates a new EventResend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResend(eventId string, eventType EventType) *EventResend {
	this := EventResend{}
	this.EventId = eventId
	this.EventType = eventType
	return &this
}

// NewEventResendWithDefaults instantiates a new EventResend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResendWithDefaults() *EventResend {
	this := EventResend{}
	return &this
}

// GetEventId returns the EventId field value
func (o *EventResend) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *EventResend) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *EventResend) SetEventId(v string) {
	o.EventId = v
}

// GetEventType returns the EventType field value
func (o *EventResend) GetEventType() EventType {
	if o == nil {
		var ret EventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EventResend) GetEventTypeOk() (*EventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EventResend) SetEventType(v EventType) {
	o.EventType = v
}

func (o EventResend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_id"] = o.EventId
	}
	if true {
		toSerialize["event_type"] = o.EventType
	}
	return json.Marshal(toSerialize)
}

type NullableEventResend struct {
	value *EventResend
	isSet bool
}

func (v NullableEventResend) Get() *EventResend {
	return v.value
}

func (v *NullableEventResend) Set(val *EventResend) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResend) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResend(val *EventResend) *NullableEventResend {
	return &NullableEventResend{value: val, isSet: true}
}

func (v NullableEventResend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


