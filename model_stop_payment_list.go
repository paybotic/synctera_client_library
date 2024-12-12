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

// checks if the StopPaymentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopPaymentList{}

// StopPaymentList struct for StopPaymentList
type StopPaymentList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string                       `json:"next_page_token,omitempty"`
	StopPayments  []StopPaymentResponseWAccount `json:"stop_payments"`
}

type _StopPaymentList StopPaymentList

// NewStopPaymentList instantiates a new StopPaymentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopPaymentList(stopPayments []StopPaymentResponseWAccount) *StopPaymentList {
	this := StopPaymentList{}
	this.StopPayments = stopPayments
	return &this
}

// NewStopPaymentListWithDefaults instantiates a new StopPaymentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopPaymentListWithDefaults() *StopPaymentList {
	this := StopPaymentList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *StopPaymentList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPaymentList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *StopPaymentList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *StopPaymentList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetStopPayments returns the StopPayments field value
func (o *StopPaymentList) GetStopPayments() []StopPaymentResponseWAccount {
	if o == nil {
		var ret []StopPaymentResponseWAccount
		return ret
	}

	return o.StopPayments
}

// GetStopPaymentsOk returns a tuple with the StopPayments field value
// and a boolean to check if the value has been set.
func (o *StopPaymentList) GetStopPaymentsOk() ([]StopPaymentResponseWAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.StopPayments, true
}

// SetStopPayments sets field value
func (o *StopPaymentList) SetStopPayments(v []StopPaymentResponseWAccount) {
	o.StopPayments = v
}

func (o StopPaymentList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopPaymentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["stop_payments"] = o.StopPayments
	return toSerialize, nil
}

func (o *StopPaymentList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stop_payments",
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

	varStopPaymentList := _StopPaymentList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStopPaymentList)

	if err != nil {
		return err
	}

	*o = StopPaymentList(varStopPaymentList)

	return err
}

type NullableStopPaymentList struct {
	value *StopPaymentList
	isSet bool
}

func (v NullableStopPaymentList) Get() *StopPaymentList {
	return v.value
}

func (v *NullableStopPaymentList) Set(val *StopPaymentList) {
	v.value = val
	v.isSet = true
}

func (v NullableStopPaymentList) IsSet() bool {
	return v.isSet
}

func (v *NullableStopPaymentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopPaymentList(val *StopPaymentList) *NullableStopPaymentList {
	return &NullableStopPaymentList{value: val, isSet: true}
}

func (v NullableStopPaymentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopPaymentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
