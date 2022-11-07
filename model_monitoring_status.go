/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MonitoringStatus The status of the alert. Any of the following: * `ACTIVE` – alert has been issued and should be investigated. * `SUPPRESSED` – alert is a false positive, alert should be dismissed, or has been otherwise investigated. 
type MonitoringStatus string

// List of monitoring_status
const (
	MONITORINGSTATUS_ACTIVE MonitoringStatus = "ACTIVE"
	MONITORINGSTATUS_SUPPRESSED MonitoringStatus = "SUPPRESSED"
)

// All allowed values of MonitoringStatus enum
var AllowedMonitoringStatusEnumValues = []MonitoringStatus{
	"ACTIVE",
	"SUPPRESSED",
}

func (v *MonitoringStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MonitoringStatus(value)
	for _, existing := range AllowedMonitoringStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MonitoringStatus", value)
}

// NewMonitoringStatusFromValue returns a pointer to a valid MonitoringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitoringStatusFromValue(v string) (*MonitoringStatus, error) {
	ev := MonitoringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitoringStatus: valid values are %v", v, AllowedMonitoringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitoringStatus) IsValid() bool {
	for _, existing := range AllowedMonitoringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to monitoring_status value
func (v MonitoringStatus) Ptr() *MonitoringStatus {
	return &v
}

type NullableMonitoringStatus struct {
	value *MonitoringStatus
	isSet bool
}

func (v NullableMonitoringStatus) Get() *MonitoringStatus {
	return v.value
}

func (v *NullableMonitoringStatus) Set(val *MonitoringStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringStatus(val *MonitoringStatus) *NullableMonitoringStatus {
	return &NullableMonitoringStatus{value: val, isSet: true}
}

func (v NullableMonitoringStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

