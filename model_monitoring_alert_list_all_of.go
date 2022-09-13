/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MonitoringAlertListAllOf struct for MonitoringAlertListAllOf
type MonitoringAlertListAllOf struct {
	Alerts []MonitoringAlert `json:"alerts"`
}

// NewMonitoringAlertListAllOf instantiates a new MonitoringAlertListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAlertListAllOf(alerts []MonitoringAlert) *MonitoringAlertListAllOf {
	this := MonitoringAlertListAllOf{}
	this.Alerts = alerts
	return &this
}

// NewMonitoringAlertListAllOfWithDefaults instantiates a new MonitoringAlertListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAlertListAllOfWithDefaults() *MonitoringAlertListAllOf {
	this := MonitoringAlertListAllOf{}
	return &this
}

// GetAlerts returns the Alerts field value
func (o *MonitoringAlertListAllOf) GetAlerts() []MonitoringAlert {
	if o == nil {
		var ret []MonitoringAlert
		return ret
	}

	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value
// and a boolean to check if the value has been set.
func (o *MonitoringAlertListAllOf) GetAlertsOk() ([]MonitoringAlert, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alerts, true
}

// SetAlerts sets field value
func (o *MonitoringAlertListAllOf) SetAlerts(v []MonitoringAlert) {
	o.Alerts = v
}

func (o MonitoringAlertListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alerts"] = o.Alerts
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringAlertListAllOf struct {
	value *MonitoringAlertListAllOf
	isSet bool
}

func (v NullableMonitoringAlertListAllOf) Get() *MonitoringAlertListAllOf {
	return v.value
}

func (v *NullableMonitoringAlertListAllOf) Set(val *MonitoringAlertListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAlertListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAlertListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAlertListAllOf(val *MonitoringAlertListAllOf) *NullableMonitoringAlertListAllOf {
	return &NullableMonitoringAlertListAllOf{value: val, isSet: true}
}

func (v NullableMonitoringAlertListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAlertListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


