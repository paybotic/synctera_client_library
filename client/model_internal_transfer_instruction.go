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

// checks if the InternalTransferInstruction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalTransferInstruction{}

// InternalTransferInstruction struct for InternalTransferInstruction
type InternalTransferInstruction struct {
	Request              InternalTransfer `json:"request"`
	Type                 string           `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _InternalTransferInstruction InternalTransferInstruction

// NewInternalTransferInstruction instantiates a new InternalTransferInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalTransferInstruction(request InternalTransfer, type_ string) *InternalTransferInstruction {
	this := InternalTransferInstruction{}
	this.Request = request
	this.Type = type_
	return &this
}

// NewInternalTransferInstructionWithDefaults instantiates a new InternalTransferInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalTransferInstructionWithDefaults() *InternalTransferInstruction {
	this := InternalTransferInstruction{}
	return &this
}

// GetRequest returns the Request field value
func (o *InternalTransferInstruction) GetRequest() InternalTransfer {
	if o == nil {
		var ret InternalTransfer
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *InternalTransferInstruction) GetRequestOk() (*InternalTransfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *InternalTransferInstruction) SetRequest(v InternalTransfer) {
	o.Request = v
}

// GetType returns the Type field value
func (o *InternalTransferInstruction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InternalTransferInstruction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InternalTransferInstruction) SetType(v string) {
	o.Type = v
}

func (o InternalTransferInstruction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalTransferInstruction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request"] = o.Request
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InternalTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request",
		"type",
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

	varInternalTransferInstruction := _InternalTransferInstruction{}

	err = json.Unmarshal(data, &varInternalTransferInstruction)

	if err != nil {
		return err
	}

	*o = InternalTransferInstruction(varInternalTransferInstruction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInternalTransferInstruction struct {
	value *InternalTransferInstruction
	isSet bool
}

func (v NullableInternalTransferInstruction) Get() *InternalTransferInstruction {
	return v.value
}

func (v *NullableInternalTransferInstruction) Set(val *InternalTransferInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalTransferInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalTransferInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalTransferInstruction(val *InternalTransferInstruction) *NullableInternalTransferInstruction {
	return &NullableInternalTransferInstruction{value: val, isSet: true}
}

func (v NullableInternalTransferInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalTransferInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
