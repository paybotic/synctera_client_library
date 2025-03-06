/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateCardImageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCardImageRequest{}

// CreateCardImageRequest struct for CreateCardImageRequest
type CreateCardImageRequest struct {
	// The unique identifier of a cards product
	CardProductId string `json:"card_product_id"`
	// The unique identifier of a customer
	CustomerId string `json:"customer_id"`
}

type _CreateCardImageRequest CreateCardImageRequest

// NewCreateCardImageRequest instantiates a new CreateCardImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCardImageRequest(cardProductId string, customerId string) *CreateCardImageRequest {
	this := CreateCardImageRequest{}
	this.CardProductId = cardProductId
	this.CustomerId = customerId
	return &this
}

// NewCreateCardImageRequestWithDefaults instantiates a new CreateCardImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCardImageRequestWithDefaults() *CreateCardImageRequest {
	this := CreateCardImageRequest{}
	return &this
}

// GetCardProductId returns the CardProductId field value
func (o *CreateCardImageRequest) GetCardProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardProductId
}

// GetCardProductIdOk returns a tuple with the CardProductId field value
// and a boolean to check if the value has been set.
func (o *CreateCardImageRequest) GetCardProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProductId, true
}

// SetCardProductId sets field value
func (o *CreateCardImageRequest) SetCardProductId(v string) {
	o.CardProductId = v
}

// GetCustomerId returns the CustomerId field value
func (o *CreateCardImageRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CreateCardImageRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CreateCardImageRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o CreateCardImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCardImageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card_product_id"] = o.CardProductId
	toSerialize["customer_id"] = o.CustomerId
	return toSerialize, nil
}

func (o *CreateCardImageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"card_product_id",
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

	varCreateCardImageRequest := _CreateCardImageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCardImageRequest)

	if err != nil {
		return err
	}

	*o = CreateCardImageRequest(varCreateCardImageRequest)

	return err
}

type NullableCreateCardImageRequest struct {
	value *CreateCardImageRequest
	isSet bool
}

func (v NullableCreateCardImageRequest) Get() *CreateCardImageRequest {
	return v.value
}

func (v *NullableCreateCardImageRequest) Set(val *CreateCardImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCardImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCardImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCardImageRequest(val *CreateCardImageRequest) *NullableCreateCardImageRequest {
	return &NullableCreateCardImageRequest{value: val, isSet: true}
}

func (v NullableCreateCardImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCardImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
