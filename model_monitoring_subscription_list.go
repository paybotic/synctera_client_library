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

// MonitoringSubscriptionList struct for MonitoringSubscriptionList
type MonitoringSubscriptionList struct {
	Subscriptions []MonitoringSubscription `json:"subscriptions"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewMonitoringSubscriptionList instantiates a new MonitoringSubscriptionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringSubscriptionList(subscriptions []MonitoringSubscription) *MonitoringSubscriptionList {
	this := MonitoringSubscriptionList{}
	this.Subscriptions = subscriptions
	return &this
}

// NewMonitoringSubscriptionListWithDefaults instantiates a new MonitoringSubscriptionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringSubscriptionListWithDefaults() *MonitoringSubscriptionList {
	this := MonitoringSubscriptionList{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value
func (o *MonitoringSubscriptionList) GetSubscriptions() []MonitoringSubscription {
	if o == nil {
		var ret []MonitoringSubscription
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *MonitoringSubscriptionList) GetSubscriptionsOk() ([]MonitoringSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *MonitoringSubscriptionList) SetSubscriptions(v []MonitoringSubscription) {
	o.Subscriptions = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *MonitoringSubscriptionList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSubscriptionList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *MonitoringSubscriptionList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *MonitoringSubscriptionList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o MonitoringSubscriptionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringSubscriptionList struct {
	value *MonitoringSubscriptionList
	isSet bool
}

func (v NullableMonitoringSubscriptionList) Get() *MonitoringSubscriptionList {
	return v.value
}

func (v *NullableMonitoringSubscriptionList) Set(val *MonitoringSubscriptionList) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringSubscriptionList) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringSubscriptionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringSubscriptionList(val *MonitoringSubscriptionList) *NullableMonitoringSubscriptionList {
	return &NullableMonitoringSubscriptionList{value: val, isSet: true}
}

func (v NullableMonitoringSubscriptionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringSubscriptionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


