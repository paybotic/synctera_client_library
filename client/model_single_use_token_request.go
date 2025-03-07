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

// checks if the SingleUseTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleUseTokenRequest{}

// SingleUseTokenRequest struct for SingleUseTokenRequest
type SingleUseTokenRequest struct {
	// The ID of the account to which the token will be linked
	AccountId string `json:"account_id"`
	// The ID of the customer to whom the token will be issued
	CustomerId           string `json:"customer_id"`
	AdditionalProperties map[string]interface{}
}

type _SingleUseTokenRequest SingleUseTokenRequest

// NewSingleUseTokenRequest instantiates a new SingleUseTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleUseTokenRequest(accountId string, customerId string) *SingleUseTokenRequest {
	this := SingleUseTokenRequest{}
	this.AccountId = accountId
	this.CustomerId = customerId
	return &this
}

// NewSingleUseTokenRequestWithDefaults instantiates a new SingleUseTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleUseTokenRequestWithDefaults() *SingleUseTokenRequest {
	this := SingleUseTokenRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *SingleUseTokenRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SingleUseTokenRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *SingleUseTokenRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetCustomerId returns the CustomerId field value
func (o *SingleUseTokenRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *SingleUseTokenRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *SingleUseTokenRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o SingleUseTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleUseTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["customer_id"] = o.CustomerId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SingleUseTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
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

	varSingleUseTokenRequest := _SingleUseTokenRequest{}

	err = json.Unmarshal(data, &varSingleUseTokenRequest)

	if err != nil {
		return err
	}

	*o = SingleUseTokenRequest(varSingleUseTokenRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "customer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSingleUseTokenRequest struct {
	value *SingleUseTokenRequest
	isSet bool
}

func (v NullableSingleUseTokenRequest) Get() *SingleUseTokenRequest {
	return v.value
}

func (v *NullableSingleUseTokenRequest) Set(val *SingleUseTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleUseTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleUseTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleUseTokenRequest(val *SingleUseTokenRequest) *NullableSingleUseTokenRequest {
	return &NullableSingleUseTokenRequest{value: val, isSet: true}
}

func (v NullableSingleUseTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleUseTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
