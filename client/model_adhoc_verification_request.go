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

// checks if the AdhocVerificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdhocVerificationRequest{}

// AdhocVerificationRequest Basic identifying information about the person being verified.
type AdhocVerificationRequest struct {
	// Synctera party (non-customer) who is receiving money from a customer (the payer)
	PayeeId string `json:"payee_id"`
	// Synctera customer who is sending money to a non-customer (the payee)
	PayerId              string `json:"payer_id"`
	AdditionalProperties map[string]interface{}
}

type _AdhocVerificationRequest AdhocVerificationRequest

// NewAdhocVerificationRequest instantiates a new AdhocVerificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdhocVerificationRequest(payeeId string, payerId string) *AdhocVerificationRequest {
	this := AdhocVerificationRequest{}
	this.PayeeId = payeeId
	this.PayerId = payerId
	return &this
}

// NewAdhocVerificationRequestWithDefaults instantiates a new AdhocVerificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdhocVerificationRequestWithDefaults() *AdhocVerificationRequest {
	this := AdhocVerificationRequest{}
	return &this
}

// GetPayeeId returns the PayeeId field value
func (o *AdhocVerificationRequest) GetPayeeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value
// and a boolean to check if the value has been set.
func (o *AdhocVerificationRequest) GetPayeeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayeeId, true
}

// SetPayeeId sets field value
func (o *AdhocVerificationRequest) SetPayeeId(v string) {
	o.PayeeId = v
}

// GetPayerId returns the PayerId field value
func (o *AdhocVerificationRequest) GetPayerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayerId
}

// GetPayerIdOk returns a tuple with the PayerId field value
// and a boolean to check if the value has been set.
func (o *AdhocVerificationRequest) GetPayerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayerId, true
}

// SetPayerId sets field value
func (o *AdhocVerificationRequest) SetPayerId(v string) {
	o.PayerId = v
}

func (o AdhocVerificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdhocVerificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payee_id"] = o.PayeeId
	toSerialize["payer_id"] = o.PayerId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdhocVerificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payee_id",
		"payer_id",
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

	varAdhocVerificationRequest := _AdhocVerificationRequest{}

	err = json.Unmarshal(data, &varAdhocVerificationRequest)

	if err != nil {
		return err
	}

	*o = AdhocVerificationRequest(varAdhocVerificationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payee_id")
		delete(additionalProperties, "payer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdhocVerificationRequest struct {
	value *AdhocVerificationRequest
	isSet bool
}

func (v NullableAdhocVerificationRequest) Get() *AdhocVerificationRequest {
	return v.value
}

func (v *NullableAdhocVerificationRequest) Set(val *AdhocVerificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdhocVerificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdhocVerificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdhocVerificationRequest(val *AdhocVerificationRequest) *NullableAdhocVerificationRequest {
	return &NullableAdhocVerificationRequest{value: val, isSet: true}
}

func (v NullableAdhocVerificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdhocVerificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
