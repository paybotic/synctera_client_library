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

// WatchlistAlertListAllOf struct for WatchlistAlertListAllOf
type WatchlistAlertListAllOf struct {
	Alerts []WatchlistAlert `json:"alerts"`
}

// NewWatchlistAlertListAllOf instantiates a new WatchlistAlertListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistAlertListAllOf(alerts []WatchlistAlert) *WatchlistAlertListAllOf {
	this := WatchlistAlertListAllOf{}
	this.Alerts = alerts
	return &this
}

// NewWatchlistAlertListAllOfWithDefaults instantiates a new WatchlistAlertListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistAlertListAllOfWithDefaults() *WatchlistAlertListAllOf {
	this := WatchlistAlertListAllOf{}
	return &this
}

// GetAlerts returns the Alerts field value
func (o *WatchlistAlertListAllOf) GetAlerts() []WatchlistAlert {
	if o == nil {
		var ret []WatchlistAlert
		return ret
	}

	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value
// and a boolean to check if the value has been set.
func (o *WatchlistAlertListAllOf) GetAlertsOk() ([]WatchlistAlert, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alerts, true
}

// SetAlerts sets field value
func (o *WatchlistAlertListAllOf) SetAlerts(v []WatchlistAlert) {
	o.Alerts = v
}

func (o WatchlistAlertListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alerts"] = o.Alerts
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistAlertListAllOf struct {
	value *WatchlistAlertListAllOf
	isSet bool
}

func (v NullableWatchlistAlertListAllOf) Get() *WatchlistAlertListAllOf {
	return v.value
}

func (v *NullableWatchlistAlertListAllOf) Set(val *WatchlistAlertListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistAlertListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistAlertListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistAlertListAllOf(val *WatchlistAlertListAllOf) *NullableWatchlistAlertListAllOf {
	return &NullableWatchlistAlertListAllOf{value: val, isSet: true}
}

func (v NullableWatchlistAlertListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistAlertListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
