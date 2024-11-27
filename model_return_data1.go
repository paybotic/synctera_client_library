/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ReturnData1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnData1{}

// ReturnData1 Data associated with a returned wire
type ReturnData1 struct {
	// Wire UUID of the original wire that was returned
	OriginalId *string `json:"original_id,omitempty"`
	// Transaction UUID of the original wire that was returned
	OriginalTransactionId *string `json:"original_transaction_id,omitempty"`
	// Wire reference ID of the original wire that was returned
	PreviousMessageId string `json:"previous_message_id"`
	// The cause of the return
	Reason *string `json:"reason,omitempty"`
}

type _ReturnData1 ReturnData1

// NewReturnData1 instantiates a new ReturnData1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnData1(previousMessageId string) *ReturnData1 {
	this := ReturnData1{}
	this.PreviousMessageId = previousMessageId
	return &this
}

// NewReturnData1WithDefaults instantiates a new ReturnData1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnData1WithDefaults() *ReturnData1 {
	this := ReturnData1{}
	return &this
}

// GetOriginalId returns the OriginalId field value if set, zero value otherwise.
func (o *ReturnData1) GetOriginalId() string {
	if o == nil || IsNil(o.OriginalId) {
		var ret string
		return ret
	}
	return *o.OriginalId
}

// GetOriginalIdOk returns a tuple with the OriginalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnData1) GetOriginalIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalId) {
		return nil, false
	}
	return o.OriginalId, true
}

// HasOriginalId returns a boolean if a field has been set.
func (o *ReturnData1) HasOriginalId() bool {
	if o != nil && !IsNil(o.OriginalId) {
		return true
	}

	return false
}

// SetOriginalId gets a reference to the given string and assigns it to the OriginalId field.
func (o *ReturnData1) SetOriginalId(v string) {
	o.OriginalId = &v
}

// GetOriginalTransactionId returns the OriginalTransactionId field value if set, zero value otherwise.
func (o *ReturnData1) GetOriginalTransactionId() string {
	if o == nil || IsNil(o.OriginalTransactionId) {
		var ret string
		return ret
	}
	return *o.OriginalTransactionId
}

// GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnData1) GetOriginalTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalTransactionId) {
		return nil, false
	}
	return o.OriginalTransactionId, true
}

// HasOriginalTransactionId returns a boolean if a field has been set.
func (o *ReturnData1) HasOriginalTransactionId() bool {
	if o != nil && !IsNil(o.OriginalTransactionId) {
		return true
	}

	return false
}

// SetOriginalTransactionId gets a reference to the given string and assigns it to the OriginalTransactionId field.
func (o *ReturnData1) SetOriginalTransactionId(v string) {
	o.OriginalTransactionId = &v
}

// GetPreviousMessageId returns the PreviousMessageId field value
func (o *ReturnData1) GetPreviousMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousMessageId
}

// GetPreviousMessageIdOk returns a tuple with the PreviousMessageId field value
// and a boolean to check if the value has been set.
func (o *ReturnData1) GetPreviousMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousMessageId, true
}

// SetPreviousMessageId sets field value
func (o *ReturnData1) SetPreviousMessageId(v string) {
	o.PreviousMessageId = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ReturnData1) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnData1) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ReturnData1) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ReturnData1) SetReason(v string) {
	o.Reason = &v
}

func (o ReturnData1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnData1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginalId) {
		toSerialize["original_id"] = o.OriginalId
	}
	if !IsNil(o.OriginalTransactionId) {
		toSerialize["original_transaction_id"] = o.OriginalTransactionId
	}
	toSerialize["previous_message_id"] = o.PreviousMessageId
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

func (o *ReturnData1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"previous_message_id",
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

	varReturnData1 := _ReturnData1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReturnData1)

	if err != nil {
		return err
	}

	*o = ReturnData1(varReturnData1)

	return err
}

type NullableReturnData1 struct {
	value *ReturnData1
	isSet bool
}

func (v NullableReturnData1) Get() *ReturnData1 {
	return v.value
}

func (v *NullableReturnData1) Set(val *ReturnData1) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnData1) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnData1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnData1(val *ReturnData1) *NullableReturnData1 {
	return &NullableReturnData1{value: val, isSet: true}
}

func (v NullableReturnData1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnData1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}