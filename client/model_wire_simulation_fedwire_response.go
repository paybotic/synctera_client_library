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

// checks if the WireSimulationFedwireResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireSimulationFedwireResponse{}

// WireSimulationFedwireResponse Incoming Wire simulation result with the created file name
type WireSimulationFedwireResponse struct {
	FileName             string `json:"file_name"`
	AdditionalProperties map[string]interface{}
}

type _WireSimulationFedwireResponse WireSimulationFedwireResponse

// NewWireSimulationFedwireResponse instantiates a new WireSimulationFedwireResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireSimulationFedwireResponse(fileName string) *WireSimulationFedwireResponse {
	this := WireSimulationFedwireResponse{}
	this.FileName = fileName
	return &this
}

// NewWireSimulationFedwireResponseWithDefaults instantiates a new WireSimulationFedwireResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireSimulationFedwireResponseWithDefaults() *WireSimulationFedwireResponse {
	this := WireSimulationFedwireResponse{}
	return &this
}

// GetFileName returns the FileName field value
func (o *WireSimulationFedwireResponse) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *WireSimulationFedwireResponse) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *WireSimulationFedwireResponse) SetFileName(v string) {
	o.FileName = v
}

func (o WireSimulationFedwireResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WireSimulationFedwireResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_name"] = o.FileName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WireSimulationFedwireResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file_name",
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

	varWireSimulationFedwireResponse := _WireSimulationFedwireResponse{}

	err = json.Unmarshal(data, &varWireSimulationFedwireResponse)

	if err != nil {
		return err
	}

	*o = WireSimulationFedwireResponse(varWireSimulationFedwireResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWireSimulationFedwireResponse struct {
	value *WireSimulationFedwireResponse
	isSet bool
}

func (v NullableWireSimulationFedwireResponse) Get() *WireSimulationFedwireResponse {
	return v.value
}

func (v *NullableWireSimulationFedwireResponse) Set(val *WireSimulationFedwireResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWireSimulationFedwireResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWireSimulationFedwireResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireSimulationFedwireResponse(val *WireSimulationFedwireResponse) *NullableWireSimulationFedwireResponse {
	return &NullableWireSimulationFedwireResponse{value: val, isSet: true}
}

func (v NullableWireSimulationFedwireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireSimulationFedwireResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
