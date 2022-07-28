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

// WatchlistSubscriptionList struct for WatchlistSubscriptionList
type WatchlistSubscriptionList struct {
	Subscriptions []WatchlistSubscription `json:"subscriptions"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewWatchlistSubscriptionList instantiates a new WatchlistSubscriptionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistSubscriptionList(subscriptions []WatchlistSubscription) *WatchlistSubscriptionList {
	this := WatchlistSubscriptionList{}
	this.Subscriptions = subscriptions
	return &this
}

// NewWatchlistSubscriptionListWithDefaults instantiates a new WatchlistSubscriptionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistSubscriptionListWithDefaults() *WatchlistSubscriptionList {
	this := WatchlistSubscriptionList{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value
func (o *WatchlistSubscriptionList) GetSubscriptions() []WatchlistSubscription {
	if o == nil {
		var ret []WatchlistSubscription
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *WatchlistSubscriptionList) GetSubscriptionsOk() ([]WatchlistSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *WatchlistSubscriptionList) SetSubscriptions(v []WatchlistSubscription) {
	o.Subscriptions = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *WatchlistSubscriptionList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistSubscriptionList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *WatchlistSubscriptionList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *WatchlistSubscriptionList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o WatchlistSubscriptionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistSubscriptionList struct {
	value *WatchlistSubscriptionList
	isSet bool
}

func (v NullableWatchlistSubscriptionList) Get() *WatchlistSubscriptionList {
	return v.value
}

func (v *NullableWatchlistSubscriptionList) Set(val *WatchlistSubscriptionList) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistSubscriptionList) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistSubscriptionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistSubscriptionList(val *WatchlistSubscriptionList) *NullableWatchlistSubscriptionList {
	return &NullableWatchlistSubscriptionList{value: val, isSet: true}
}

func (v NullableWatchlistSubscriptionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistSubscriptionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


