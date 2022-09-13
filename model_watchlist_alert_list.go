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

// WatchlistAlertList struct for WatchlistAlertList
type WatchlistAlertList struct {
	Alerts []WatchlistAlert `json:"alerts"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewWatchlistAlertList instantiates a new WatchlistAlertList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistAlertList(alerts []WatchlistAlert) *WatchlistAlertList {
	this := WatchlistAlertList{}
	this.Alerts = alerts
	return &this
}

// NewWatchlistAlertListWithDefaults instantiates a new WatchlistAlertList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistAlertListWithDefaults() *WatchlistAlertList {
	this := WatchlistAlertList{}
	return &this
}

// GetAlerts returns the Alerts field value
func (o *WatchlistAlertList) GetAlerts() []WatchlistAlert {
	if o == nil {
		var ret []WatchlistAlert
		return ret
	}

	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value
// and a boolean to check if the value has been set.
func (o *WatchlistAlertList) GetAlertsOk() ([]WatchlistAlert, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alerts, true
}

// SetAlerts sets field value
func (o *WatchlistAlertList) SetAlerts(v []WatchlistAlert) {
	o.Alerts = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *WatchlistAlertList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistAlertList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *WatchlistAlertList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *WatchlistAlertList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o WatchlistAlertList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alerts"] = o.Alerts
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistAlertList struct {
	value *WatchlistAlertList
	isSet bool
}

func (v NullableWatchlistAlertList) Get() *WatchlistAlertList {
	return v.value
}

func (v *NullableWatchlistAlertList) Set(val *WatchlistAlertList) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistAlertList) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistAlertList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistAlertList(val *WatchlistAlertList) *NullableWatchlistAlertList {
	return &NullableWatchlistAlertList{value: val, isSet: true}
}

func (v NullableWatchlistAlertList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistAlertList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


