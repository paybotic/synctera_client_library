/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PingResponse struct for PingResponse
type PingResponse struct {
	// status of webhook endpoint
	WebhookStatus string `json:"webhook_status"`
}

// NewPingResponse instantiates a new PingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingResponse(webhookStatus string) *PingResponse {
	this := PingResponse{}
	this.WebhookStatus = webhookStatus
	return &this
}

// NewPingResponseWithDefaults instantiates a new PingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingResponseWithDefaults() *PingResponse {
	this := PingResponse{}
	return &this
}

// GetWebhookStatus returns the WebhookStatus field value
func (o *PingResponse) GetWebhookStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookStatus
}

// GetWebhookStatusOk returns a tuple with the WebhookStatus field value
// and a boolean to check if the value has been set.
func (o *PingResponse) GetWebhookStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookStatus, true
}

// SetWebhookStatus sets field value
func (o *PingResponse) SetWebhookStatus(v string) {
	o.WebhookStatus = v
}

func (o PingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_status"] = o.WebhookStatus
	}
	return json.Marshal(toSerialize)
}

type NullablePingResponse struct {
	value *PingResponse
	isSet bool
}

func (v NullablePingResponse) Get() *PingResponse {
	return v.value
}

func (v *NullablePingResponse) Set(val *PingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingResponse(val *PingResponse) *NullablePingResponse {
	return &NullablePingResponse{value: val, isSet: true}
}

func (v NullablePingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


