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

// EventList struct for EventList
type EventList struct {
	// Array of events
	EventList []Event `json:"event_list"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewEventList instantiates a new EventList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventList(eventList []Event) *EventList {
	this := EventList{}
	this.EventList = eventList
	return &this
}

// NewEventListWithDefaults instantiates a new EventList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventListWithDefaults() *EventList {
	this := EventList{}
	return &this
}

// GetEventList returns the EventList field value
func (o *EventList) GetEventList() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.EventList
}

// GetEventListOk returns a tuple with the EventList field value
// and a boolean to check if the value has been set.
func (o *EventList) GetEventListOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventList, true
}

// SetEventList sets field value
func (o *EventList) SetEventList(v []Event) {
	o.EventList = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *EventList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *EventList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *EventList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o EventList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_list"] = o.EventList
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableEventList struct {
	value *EventList
	isSet bool
}

func (v NullableEventList) Get() *EventList {
	return v.value
}

func (v *NullableEventList) Set(val *EventList) {
	v.value = val
	v.isSet = true
}

func (v NullableEventList) IsSet() bool {
	return v.isSet
}

func (v *NullableEventList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventList(val *EventList) *NullableEventList {
	return &NullableEventList{value: val, isSet: true}
}

func (v NullableEventList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


