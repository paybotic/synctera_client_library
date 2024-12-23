/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// WebhookResponse struct for WebhookResponse
type WebhookResponse struct {
	// indicates whether webhook is active
	Active *bool          `json:"active,omitempty"`
	Config *WebhookConfig `json:"config,omitempty"`
	// list of webhook events, use * to receive all notifications
	Events []string `json:"events,omitempty"`
	// id of the webhook
	Id *string `json:"id,omitempty"`
	// name of the webhook
	Name *string `json:"name,omitempty"`
	// The timestamp representing when the webhook request was made
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The timestamp representing when the webhook was last modified
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
}

// NewWebhookResponse instantiates a new WebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponse() *WebhookResponse {
	this := WebhookResponse{}
	return &this
}

// NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseWithDefaults() *WebhookResponse {
	this := WebhookResponse{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WebhookResponse) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WebhookResponse) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WebhookResponse) SetActive(v bool) {
	o.Active = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *WebhookResponse) GetConfig() WebhookConfig {
	if o == nil || o.Config == nil {
		var ret WebhookConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetConfigOk() (*WebhookConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *WebhookResponse) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given WebhookConfig and assigns it to the Config field.
func (o *WebhookResponse) SetConfig(v WebhookConfig) {
	o.Config = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *WebhookResponse) GetEvents() []string {
	if o == nil || o.Events == nil {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetEventsOk() ([]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *WebhookResponse) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *WebhookResponse) SetEvents(v []string) {
	o.Events = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhookResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookResponse) SetName(v string) {
	o.Name = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *WebhookResponse) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *WebhookResponse) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *WebhookResponse) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *WebhookResponse) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *WebhookResponse) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *WebhookResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

func (o WebhookResponse) MarshalJSON() ([]byte, error) {
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
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookResponse struct {
	value *WebhookResponse
	isSet bool
}

func (v NullableWebhookResponse) Get() *WebhookResponse {
	return v.value
}

func (v *NullableWebhookResponse) Set(val *WebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponse(val *WebhookResponse) *NullableWebhookResponse {
	return &NullableWebhookResponse{value: val, isSet: true}
}

func (v NullableWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
