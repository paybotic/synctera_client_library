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

// MinimumPaymentFull - The scheme for calculating the minimum payment due for outstanding balances in a billing period.
type MinimumPaymentFull struct {
	MinimumPaymentTypeFull *MinimumPaymentTypeFull
}

// MinimumPaymentTypeFullAsMinimumPaymentFull is a convenience function that returns MinimumPaymentTypeFull wrapped in MinimumPaymentFull
func MinimumPaymentTypeFullAsMinimumPaymentFull(v *MinimumPaymentTypeFull) MinimumPaymentFull {
	return MinimumPaymentFull{
		MinimumPaymentTypeFull: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MinimumPaymentFull) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'FULL'
	if jsonDict["type"] == "FULL" {
		// try to unmarshal JSON data into MinimumPaymentTypeFull
		err = json.Unmarshal(data, &dst.MinimumPaymentTypeFull)
		if err == nil {
			return nil // data stored in dst.MinimumPaymentTypeFull, return on the first match
		} else {
			dst.MinimumPaymentTypeFull = nil
			return fmt.Errorf("failed to unmarshal MinimumPaymentFull as MinimumPaymentTypeFull: %s", err.Error())
		}
	}

	// check if the discriminator value is 'minimum_payment_type_full'
	if jsonDict["type"] == "minimum_payment_type_full" {
		// try to unmarshal JSON data into MinimumPaymentTypeFull
		err = json.Unmarshal(data, &dst.MinimumPaymentTypeFull)
		if err == nil {
			return nil // data stored in dst.MinimumPaymentTypeFull, return on the first match
		} else {
			dst.MinimumPaymentTypeFull = nil
			return fmt.Errorf("failed to unmarshal MinimumPaymentFull as MinimumPaymentTypeFull: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MinimumPaymentFull) MarshalJSON() ([]byte, error) {
	if src.MinimumPaymentTypeFull != nil {
		return json.Marshal(&src.MinimumPaymentTypeFull)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MinimumPaymentFull) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MinimumPaymentTypeFull != nil {
		return obj.MinimumPaymentTypeFull
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj MinimumPaymentFull) GetActualInstanceValue() interface{} {
	if obj.MinimumPaymentTypeFull != nil {
		return *obj.MinimumPaymentTypeFull
	}

	// all schemas are nil
	return nil
}

type NullableMinimumPaymentFull struct {
	value *MinimumPaymentFull
	isSet bool
}

func (v NullableMinimumPaymentFull) Get() *MinimumPaymentFull {
	return v.value
}

func (v *NullableMinimumPaymentFull) Set(val *MinimumPaymentFull) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimumPaymentFull) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimumPaymentFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimumPaymentFull(val *MinimumPaymentFull) *NullableMinimumPaymentFull {
	return &NullableMinimumPaymentFull{value: val, isSet: true}
}

func (v NullableMinimumPaymentFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimumPaymentFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
