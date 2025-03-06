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

// checks if the WireSimulationDatasoftResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireSimulationDatasoftResponse{}

// WireSimulationDatasoftResponse Incoming Wire simulation result with the webhook ID
type WireSimulationDatasoftResponse struct {
	WebhookId            string `json:"webhook_id"`
	AdditionalProperties map[string]interface{}
}

type _WireSimulationDatasoftResponse WireSimulationDatasoftResponse

// NewWireSimulationDatasoftResponse instantiates a new WireSimulationDatasoftResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireSimulationDatasoftResponse(webhookId string) *WireSimulationDatasoftResponse {
	this := WireSimulationDatasoftResponse{}
	this.WebhookId = webhookId
	return &this
}

// NewWireSimulationDatasoftResponseWithDefaults instantiates a new WireSimulationDatasoftResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireSimulationDatasoftResponseWithDefaults() *WireSimulationDatasoftResponse {
	this := WireSimulationDatasoftResponse{}
	return &this
}

// GetWebhookId returns the WebhookId field value
func (o *WireSimulationDatasoftResponse) GetWebhookId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value
// and a boolean to check if the value has been set.
func (o *WireSimulationDatasoftResponse) GetWebhookIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookId, true
}

// SetWebhookId sets field value
func (o *WireSimulationDatasoftResponse) SetWebhookId(v string) {
	o.WebhookId = v
}

func (o WireSimulationDatasoftResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WireSimulationDatasoftResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["webhook_id"] = o.WebhookId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WireSimulationDatasoftResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"webhook_id",
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

	varWireSimulationDatasoftResponse := _WireSimulationDatasoftResponse{}

	err = json.Unmarshal(data, &varWireSimulationDatasoftResponse)

	if err != nil {
		return err
	}

	*o = WireSimulationDatasoftResponse(varWireSimulationDatasoftResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWireSimulationDatasoftResponse struct {
	value *WireSimulationDatasoftResponse
	isSet bool
}

func (v NullableWireSimulationDatasoftResponse) Get() *WireSimulationDatasoftResponse {
	return v.value
}

func (v *NullableWireSimulationDatasoftResponse) Set(val *WireSimulationDatasoftResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWireSimulationDatasoftResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWireSimulationDatasoftResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireSimulationDatasoftResponse(val *WireSimulationDatasoftResponse) *NullableWireSimulationDatasoftResponse {
	return &NullableWireSimulationDatasoftResponse{value: val, isSet: true}
}

func (v NullableWireSimulationDatasoftResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireSimulationDatasoftResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
