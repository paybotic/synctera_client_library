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

// checks if the DigitalWalletTokenEditRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalWalletTokenEditRequest{}

// DigitalWalletTokenEditRequest struct for DigitalWalletTokenEditRequest
type DigitalWalletTokenEditRequest struct {
	// The status indicating the digital wallet token lifecycle state
	TokenStatus          string `json:"token_status"`
	AdditionalProperties map[string]interface{}
}

type _DigitalWalletTokenEditRequest DigitalWalletTokenEditRequest

// NewDigitalWalletTokenEditRequest instantiates a new DigitalWalletTokenEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletTokenEditRequest(tokenStatus string) *DigitalWalletTokenEditRequest {
	this := DigitalWalletTokenEditRequest{}
	this.TokenStatus = tokenStatus
	return &this
}

// NewDigitalWalletTokenEditRequestWithDefaults instantiates a new DigitalWalletTokenEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletTokenEditRequestWithDefaults() *DigitalWalletTokenEditRequest {
	this := DigitalWalletTokenEditRequest{}
	return &this
}

// GetTokenStatus returns the TokenStatus field value
func (o *DigitalWalletTokenEditRequest) GetTokenStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenStatus
}

// GetTokenStatusOk returns a tuple with the TokenStatus field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenEditRequest) GetTokenStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenStatus, true
}

// SetTokenStatus sets field value
func (o *DigitalWalletTokenEditRequest) SetTokenStatus(v string) {
	o.TokenStatus = v
}

func (o DigitalWalletTokenEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalWalletTokenEditRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_status"] = o.TokenStatus

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DigitalWalletTokenEditRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token_status",
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

	varDigitalWalletTokenEditRequest := _DigitalWalletTokenEditRequest{}

	err = json.Unmarshal(data, &varDigitalWalletTokenEditRequest)

	if err != nil {
		return err
	}

	*o = DigitalWalletTokenEditRequest(varDigitalWalletTokenEditRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDigitalWalletTokenEditRequest struct {
	value *DigitalWalletTokenEditRequest
	isSet bool
}

func (v NullableDigitalWalletTokenEditRequest) Get() *DigitalWalletTokenEditRequest {
	return v.value
}

func (v *NullableDigitalWalletTokenEditRequest) Set(val *DigitalWalletTokenEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletTokenEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletTokenEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletTokenEditRequest(val *DigitalWalletTokenEditRequest) *NullableDigitalWalletTokenEditRequest {
	return &NullableDigitalWalletTokenEditRequest{value: val, isSet: true}
}

func (v NullableDigitalWalletTokenEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletTokenEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
