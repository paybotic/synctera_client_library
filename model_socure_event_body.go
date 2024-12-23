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

// SocureEventBody struct for SocureEventBody
type SocureEventBody struct {
	// Environment the event belongs to
	EnvironmentName string                `json:"environmentName"`
	Event           SocureWatchlistResult `json:"event"`
	// Unique identifier for the monitoring event
	Id string `json:"id"`
}

// NewSocureEventBody instantiates a new SocureEventBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocureEventBody(environmentName string, event SocureWatchlistResult, id string) *SocureEventBody {
	this := SocureEventBody{}
	this.EnvironmentName = environmentName
	this.Event = event
	this.Id = id
	return &this
}

// NewSocureEventBodyWithDefaults instantiates a new SocureEventBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocureEventBodyWithDefaults() *SocureEventBody {
	this := SocureEventBody{}
	return &this
}

// GetEnvironmentName returns the EnvironmentName field value
func (o *SocureEventBody) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *SocureEventBody) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value
func (o *SocureEventBody) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetEvent returns the Event field value
func (o *SocureEventBody) GetEvent() SocureWatchlistResult {
	if o == nil {
		var ret SocureWatchlistResult
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *SocureEventBody) GetEventOk() (*SocureWatchlistResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *SocureEventBody) SetEvent(v SocureWatchlistResult) {
	o.Event = v
}

// GetId returns the Id field value
func (o *SocureEventBody) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SocureEventBody) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SocureEventBody) SetId(v string) {
	o.Id = v
}

func (o SocureEventBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environmentName"] = o.EnvironmentName
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSocureEventBody struct {
	value *SocureEventBody
	isSet bool
}

func (v NullableSocureEventBody) Get() *SocureEventBody {
	return v.value
}

func (v *NullableSocureEventBody) Set(val *SocureEventBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSocureEventBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSocureEventBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocureEventBody(val *SocureEventBody) *NullableSocureEventBody {
	return &NullableSocureEventBody{value: val, isSet: true}
}

func (v NullableSocureEventBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocureEventBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
