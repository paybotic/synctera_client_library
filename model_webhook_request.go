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

// WebhookRequest struct for WebhookRequest
type WebhookRequest struct {
	// indicates whether webhook is active
	Active *bool          `json:"active,omitempty"`
	Config *WebhookConfig `json:"config,omitempty"`
	// list of webhook events, use * to receive all notifications
	Events []string `json:"events,omitempty"`
	// id of the webhook
	Id *string `json:"id,omitempty"`
	// name of the webhook
	Name *string `json:"name,omitempty"`
}

// NewWebhookRequest instantiates a new WebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookRequest() *WebhookRequest {
	this := WebhookRequest{}
	return &this
}

// NewWebhookRequestWithDefaults instantiates a new WebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookRequestWithDefaults() *WebhookRequest {
	this := WebhookRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WebhookRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WebhookRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WebhookRequest) SetActive(v bool) {
	o.Active = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *WebhookRequest) GetConfig() WebhookConfig {
	if o == nil || o.Config == nil {
		var ret WebhookConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetConfigOk() (*WebhookConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *WebhookRequest) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given WebhookConfig and assigns it to the Config field.
func (o *WebhookRequest) SetConfig(v WebhookConfig) {
	o.Config = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *WebhookRequest) GetEvents() []string {
	if o == nil || o.Events == nil {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetEventsOk() ([]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *WebhookRequest) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *WebhookRequest) SetEvents(v []string) {
	o.Events = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhookRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookRequest) SetName(v string) {
	o.Name = &v
}

func (o WebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookRequest struct {
	value *WebhookRequest
	isSet bool
}

func (v NullableWebhookRequest) Get() *WebhookRequest {
	return v.value
}

func (v *NullableWebhookRequest) Set(val *WebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookRequest(val *WebhookRequest) *NullableWebhookRequest {
	return &NullableWebhookRequest{value: val, isSet: true}
}

func (v NullableWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
