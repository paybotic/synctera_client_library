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

// MinimumPayment - The scheme for calculating the minimum payment due for outstanding balances in a billing period.
type MinimumPayment struct {
	MinimumPaymentTypeFull         *MinimumPaymentTypeFull
	MinimumPaymentTypeRateOrAmount *MinimumPaymentTypeRateOrAmount
}

// MinimumPaymentTypeFullAsMinimumPayment is a convenience function that returns MinimumPaymentTypeFull wrapped in MinimumPayment
func MinimumPaymentTypeFullAsMinimumPayment(v *MinimumPaymentTypeFull) MinimumPayment {
	return MinimumPayment{
		MinimumPaymentTypeFull: v,
	}
}

// MinimumPaymentTypeRateOrAmountAsMinimumPayment is a convenience function that returns MinimumPaymentTypeRateOrAmount wrapped in MinimumPayment
func MinimumPaymentTypeRateOrAmountAsMinimumPayment(v *MinimumPaymentTypeRateOrAmount) MinimumPayment {
	return MinimumPayment{
		MinimumPaymentTypeRateOrAmount: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MinimumPayment) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal MinimumPayment as MinimumPaymentTypeFull: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RATE_OR_AMOUNT'
	if jsonDict["type"] == "RATE_OR_AMOUNT" {
		// try to unmarshal JSON data into MinimumPaymentTypeRateOrAmount
		err = json.Unmarshal(data, &dst.MinimumPaymentTypeRateOrAmount)
		if err == nil {
			return nil // data stored in dst.MinimumPaymentTypeRateOrAmount, return on the first match
		} else {
			dst.MinimumPaymentTypeRateOrAmount = nil
			return fmt.Errorf("failed to unmarshal MinimumPayment as MinimumPaymentTypeRateOrAmount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MinimumPayment as MinimumPaymentTypeFull: %s", err.Error())
		}
	}

	// check if the discriminator value is 'minimum_payment_type_rate_or_amount'
	if jsonDict["type"] == "minimum_payment_type_rate_or_amount" {
		// try to unmarshal JSON data into MinimumPaymentTypeRateOrAmount
		err = json.Unmarshal(data, &dst.MinimumPaymentTypeRateOrAmount)
		if err == nil {
			return nil // data stored in dst.MinimumPaymentTypeRateOrAmount, return on the first match
		} else {
			dst.MinimumPaymentTypeRateOrAmount = nil
			return fmt.Errorf("failed to unmarshal MinimumPayment as MinimumPaymentTypeRateOrAmount: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MinimumPayment) MarshalJSON() ([]byte, error) {
	if src.MinimumPaymentTypeFull != nil {
		return json.Marshal(&src.MinimumPaymentTypeFull)
	}

	if src.MinimumPaymentTypeRateOrAmount != nil {
		return json.Marshal(&src.MinimumPaymentTypeRateOrAmount)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MinimumPayment) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MinimumPaymentTypeFull != nil {
		return obj.MinimumPaymentTypeFull
	}

	if obj.MinimumPaymentTypeRateOrAmount != nil {
		return obj.MinimumPaymentTypeRateOrAmount
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj MinimumPayment) GetActualInstanceValue() interface{} {
	if obj.MinimumPaymentTypeFull != nil {
		return *obj.MinimumPaymentTypeFull
	}

	if obj.MinimumPaymentTypeRateOrAmount != nil {
		return *obj.MinimumPaymentTypeRateOrAmount
	}

	// all schemas are nil
	return nil
}

type NullableMinimumPayment struct {
	value *MinimumPayment
	isSet bool
}

func (v NullableMinimumPayment) Get() *MinimumPayment {
	return v.value
}

func (v *NullableMinimumPayment) Set(val *MinimumPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimumPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimumPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimumPayment(val *MinimumPayment) *NullableMinimumPayment {
	return &NullableMinimumPayment{value: val, isSet: true}
}

func (v NullableMinimumPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimumPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
