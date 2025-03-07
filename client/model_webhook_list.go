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

// checks if the WebhookList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookList{}

// WebhookList struct for WebhookList
type WebhookList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of webhooks
	Webhooks             []Webhook `json:"webhooks"`
	AdditionalProperties map[string]interface{}
}

type _WebhookList WebhookList

// NewWebhookList instantiates a new WebhookList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookList(webhooks []Webhook) *WebhookList {
	this := WebhookList{}
	this.Webhooks = webhooks
	return &this
}

// NewWebhookListWithDefaults instantiates a new WebhookList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookListWithDefaults() *WebhookList {
	this := WebhookList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *WebhookList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *WebhookList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *WebhookList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetWebhooks returns the Webhooks field value
func (o *WebhookList) GetWebhooks() []Webhook {
	if o == nil {
		var ret []Webhook
		return ret
	}

	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value
// and a boolean to check if the value has been set.
func (o *WebhookList) GetWebhooksOk() ([]Webhook, bool) {
	if o == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// SetWebhooks sets field value
func (o *WebhookList) SetWebhooks(v []Webhook) {
	o.Webhooks = v
}

func (o WebhookList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["webhooks"] = o.Webhooks

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"webhooks",
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

	varWebhookList := _WebhookList{}

	err = json.Unmarshal(data, &varWebhookList)

	if err != nil {
		return err
	}

	*o = WebhookList(varWebhookList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "webhooks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookList struct {
	value *WebhookList
	isSet bool
}

func (v NullableWebhookList) Get() *WebhookList {
	return v.value
}

func (v *NullableWebhookList) Set(val *WebhookList) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookList) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookList(val *WebhookList) *NullableWebhookList {
	return &NullableWebhookList{value: val, isSet: true}
}

func (v NullableWebhookList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
