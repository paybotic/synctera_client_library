/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the MonitoringAlertList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringAlertList{}

// MonitoringAlertList struct for MonitoringAlertList
type MonitoringAlertList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken        *string           `json:"next_page_token,omitempty"`
	Alerts               []MonitoringAlert `json:"alerts"`
	AdditionalProperties map[string]interface{}
}

type _MonitoringAlertList MonitoringAlertList

// NewMonitoringAlertList instantiates a new MonitoringAlertList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAlertList(alerts []MonitoringAlert) *MonitoringAlertList {
	this := MonitoringAlertList{}
	this.Alerts = alerts
	return &this
}

// NewMonitoringAlertListWithDefaults instantiates a new MonitoringAlertList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAlertListWithDefaults() *MonitoringAlertList {
	this := MonitoringAlertList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *MonitoringAlertList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAlertList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *MonitoringAlertList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *MonitoringAlertList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetAlerts returns the Alerts field value
func (o *MonitoringAlertList) GetAlerts() []MonitoringAlert {
	if o == nil {
		var ret []MonitoringAlert
		return ret
	}

	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value
// and a boolean to check if the value has been set.
func (o *MonitoringAlertList) GetAlertsOk() ([]MonitoringAlert, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alerts, true
}

// SetAlerts sets field value
func (o *MonitoringAlertList) SetAlerts(v []MonitoringAlert) {
	o.Alerts = v
}

func (o MonitoringAlertList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringAlertList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["alerts"] = o.Alerts

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MonitoringAlertList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alerts",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMonitoringAlertList := _MonitoringAlertList{}

	err = json.Unmarshal(data, &varMonitoringAlertList)

	if err != nil {
		return err
	}

	*o = MonitoringAlertList(varMonitoringAlertList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "alerts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMonitoringAlertList struct {
	value *MonitoringAlertList
	isSet bool
}

func (v NullableMonitoringAlertList) Get() *MonitoringAlertList {
	return v.value
}

func (v *NullableMonitoringAlertList) Set(val *MonitoringAlertList) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAlertList) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAlertList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAlertList(val *MonitoringAlertList) *NullableMonitoringAlertList {
	return &NullableMonitoringAlertList{value: val, isSet: true}
}

func (v NullableMonitoringAlertList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAlertList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
