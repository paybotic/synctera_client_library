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
	"time"
)

// checks if the StopPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopPayment{}

// StopPayment struct for StopPayment
type StopPayment struct {
	// ID of the dispute that created the stop payment
	DisputeId *string `json:"dispute_id,omitempty"`
	// The date when this stop payment is no longer valid. This is only for business accounts.
	ExpiresOn *time.Time `json:"expires_on,omitempty"`
	// Name of the originator
	OriginatorName string `json:"originator_name"`
	StopPaymentId  string `json:"stop_payment_id"`
	// If this stop payment was created from a disputed transaction, transaction_id references the posted transaction.
	TransactionId        *string `json:"transaction_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StopPayment StopPayment

// NewStopPayment instantiates a new StopPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopPayment(originatorName string, stopPaymentId string) *StopPayment {
	this := StopPayment{}
	this.OriginatorName = originatorName
	this.StopPaymentId = stopPaymentId
	return &this
}

// NewStopPaymentWithDefaults instantiates a new StopPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopPaymentWithDefaults() *StopPayment {
	this := StopPayment{}
	return &this
}

// GetDisputeId returns the DisputeId field value if set, zero value otherwise.
func (o *StopPayment) GetDisputeId() string {
	if o == nil || IsNil(o.DisputeId) {
		var ret string
		return ret
	}
	return *o.DisputeId
}

// GetDisputeIdOk returns a tuple with the DisputeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPayment) GetDisputeIdOk() (*string, bool) {
	if o == nil || IsNil(o.DisputeId) {
		return nil, false
	}
	return o.DisputeId, true
}

// HasDisputeId returns a boolean if a field has been set.
func (o *StopPayment) HasDisputeId() bool {
	if o != nil && !IsNil(o.DisputeId) {
		return true
	}

	return false
}

// SetDisputeId gets a reference to the given string and assigns it to the DisputeId field.
func (o *StopPayment) SetDisputeId(v string) {
	o.DisputeId = &v
}

// GetExpiresOn returns the ExpiresOn field value if set, zero value otherwise.
func (o *StopPayment) GetExpiresOn() time.Time {
	if o == nil || IsNil(o.ExpiresOn) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPayment) GetExpiresOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresOn) {
		return nil, false
	}
	return o.ExpiresOn, true
}

// HasExpiresOn returns a boolean if a field has been set.
func (o *StopPayment) HasExpiresOn() bool {
	if o != nil && !IsNil(o.ExpiresOn) {
		return true
	}

	return false
}

// SetExpiresOn gets a reference to the given time.Time and assigns it to the ExpiresOn field.
func (o *StopPayment) SetExpiresOn(v time.Time) {
	o.ExpiresOn = &v
}

// GetOriginatorName returns the OriginatorName field value
func (o *StopPayment) GetOriginatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatorName
}

// GetOriginatorNameOk returns a tuple with the OriginatorName field value
// and a boolean to check if the value has been set.
func (o *StopPayment) GetOriginatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatorName, true
}

// SetOriginatorName sets field value
func (o *StopPayment) SetOriginatorName(v string) {
	o.OriginatorName = v
}

// GetStopPaymentId returns the StopPaymentId field value
func (o *StopPayment) GetStopPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StopPaymentId
}

// GetStopPaymentIdOk returns a tuple with the StopPaymentId field value
// and a boolean to check if the value has been set.
func (o *StopPayment) GetStopPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopPaymentId, true
}

// SetStopPaymentId sets field value
func (o *StopPayment) SetStopPaymentId(v string) {
	o.StopPaymentId = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *StopPayment) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPayment) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *StopPayment) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *StopPayment) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o StopPayment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisputeId) {
		toSerialize["dispute_id"] = o.DisputeId
	}
	if !IsNil(o.ExpiresOn) {
		toSerialize["expires_on"] = o.ExpiresOn
	}
	toSerialize["originator_name"] = o.OriginatorName
	toSerialize["stop_payment_id"] = o.StopPaymentId
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StopPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"originator_name",
		"stop_payment_id",
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

	varStopPayment := _StopPayment{}

	err = json.Unmarshal(data, &varStopPayment)

	if err != nil {
		return err
	}

	*o = StopPayment(varStopPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dispute_id")
		delete(additionalProperties, "expires_on")
		delete(additionalProperties, "originator_name")
		delete(additionalProperties, "stop_payment_id")
		delete(additionalProperties, "transaction_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStopPayment struct {
	value *StopPayment
	isSet bool
}

func (v NullableStopPayment) Get() *StopPayment {
	return v.value
}

func (v *NullableStopPayment) Set(val *StopPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableStopPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableStopPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopPayment(val *StopPayment) *NullableStopPayment {
	return &NullableStopPayment{value: val, isSet: true}
}

func (v NullableStopPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
