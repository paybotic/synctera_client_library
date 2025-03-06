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

// checks if the CardActivationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardActivationRequest{}

// CardActivationRequest struct for CardActivationRequest
type CardActivationRequest struct {
	// An activation code provided with the card required to prove possession of the card
	ActivationCode string `json:"activation_code"`
	// The ID of the customer for which card is being activated
	CustomerId           string `json:"customer_id"`
	AdditionalProperties map[string]interface{}
}

type _CardActivationRequest CardActivationRequest

// NewCardActivationRequest instantiates a new CardActivationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardActivationRequest(activationCode string, customerId string) *CardActivationRequest {
	this := CardActivationRequest{}
	this.ActivationCode = activationCode
	this.CustomerId = customerId
	return &this
}

// NewCardActivationRequestWithDefaults instantiates a new CardActivationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardActivationRequestWithDefaults() *CardActivationRequest {
	this := CardActivationRequest{}
	return &this
}

// GetActivationCode returns the ActivationCode field value
func (o *CardActivationRequest) GetActivationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivationCode
}

// GetActivationCodeOk returns a tuple with the ActivationCode field value
// and a boolean to check if the value has been set.
func (o *CardActivationRequest) GetActivationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivationCode, true
}

// SetActivationCode sets field value
func (o *CardActivationRequest) SetActivationCode(v string) {
	o.ActivationCode = v
}

// GetCustomerId returns the CustomerId field value
func (o *CardActivationRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CardActivationRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CardActivationRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o CardActivationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardActivationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activation_code"] = o.ActivationCode
	toSerialize["customer_id"] = o.CustomerId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardActivationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activation_code",
		"customer_id",
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

	varCardActivationRequest := _CardActivationRequest{}

	err = json.Unmarshal(data, &varCardActivationRequest)

	if err != nil {
		return err
	}

	*o = CardActivationRequest(varCardActivationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activation_code")
		delete(additionalProperties, "customer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardActivationRequest struct {
	value *CardActivationRequest
	isSet bool
}

func (v NullableCardActivationRequest) Get() *CardActivationRequest {
	return v.value
}

func (v *NullableCardActivationRequest) Set(val *CardActivationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardActivationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardActivationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardActivationRequest(val *CardActivationRequest) *NullableCardActivationRequest {
	return &NullableCardActivationRequest{value: val, isSet: true}
}

func (v NullableCardActivationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardActivationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
