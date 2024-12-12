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

// WebhookConfig struct for WebhookConfig
type WebhookConfig struct {
	CustomHeader *CustomHeaders `json:"custom_header,omitempty"`
	// password for access webhook endpoint
	Password *string `json:"password,omitempty"`
	// url of webhook endpoint
	Url *string `json:"url,omitempty"`
	// username for access webhook endpoint
	Username *string `json:"username,omitempty"`
}

// NewWebhookConfig instantiates a new WebhookConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookConfig() *WebhookConfig {
	this := WebhookConfig{}
	return &this
}

// NewWebhookConfigWithDefaults instantiates a new WebhookConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookConfigWithDefaults() *WebhookConfig {
	this := WebhookConfig{}
	return &this
}

// GetCustomHeader returns the CustomHeader field value if set, zero value otherwise.
func (o *WebhookConfig) GetCustomHeader() CustomHeaders {
	if o == nil || o.CustomHeader == nil {
		var ret CustomHeaders
		return ret
	}
	return *o.CustomHeader
}

// GetCustomHeaderOk returns a tuple with the CustomHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookConfig) GetCustomHeaderOk() (*CustomHeaders, bool) {
	if o == nil || o.CustomHeader == nil {
		return nil, false
	}
	return o.CustomHeader, true
}

// HasCustomHeader returns a boolean if a field has been set.
func (o *WebhookConfig) HasCustomHeader() bool {
	if o != nil && o.CustomHeader != nil {
		return true
	}

	return false
}

// SetCustomHeader gets a reference to the given CustomHeaders and assigns it to the CustomHeader field.
func (o *WebhookConfig) SetCustomHeader(v CustomHeaders) {
	o.CustomHeader = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *WebhookConfig) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *WebhookConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *WebhookConfig) SetPassword(v string) {
	o.Password = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WebhookConfig) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookConfig) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WebhookConfig) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WebhookConfig) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *WebhookConfig) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookConfig) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *WebhookConfig) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *WebhookConfig) SetUsername(v string) {
	o.Username = &v
}

func (o WebhookConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomHeader != nil {
		toSerialize["custom_header"] = o.CustomHeader
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookConfig struct {
	value *WebhookConfig
	isSet bool
}

func (v NullableWebhookConfig) Get() *WebhookConfig {
	return v.value
}

func (v *NullableWebhookConfig) Set(val *WebhookConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookConfig(val *WebhookConfig) *NullableWebhookConfig {
	return &NullableWebhookConfig{value: val, isSet: true}
}

func (v NullableWebhookConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
