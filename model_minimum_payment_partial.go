/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MinimumPaymentPartial - The scheme for calculating the minimum payment due for outstanding balances in a billing period. 
type MinimumPaymentPartial struct {
	MinimumPaymentTypeRateOrAmount *MinimumPaymentTypeRateOrAmount
}

// MinimumPaymentTypeRateOrAmountAsMinimumPaymentPartial is a convenience function that returns MinimumPaymentTypeRateOrAmount wrapped in MinimumPaymentPartial
func MinimumPaymentTypeRateOrAmountAsMinimumPaymentPartial(v *MinimumPaymentTypeRateOrAmount) MinimumPaymentPartial {
	return MinimumPaymentPartial{
		MinimumPaymentTypeRateOrAmount: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MinimumPaymentPartial) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'RATE_OR_AMOUNT'
	if jsonDict["type"] == "RATE_OR_AMOUNT" {
		// try to unmarshal JSON data into MinimumPaymentTypeRateOrAmount
		err = json.Unmarshal(data, &dst.MinimumPaymentTypeRateOrAmount)
		if err == nil {
			return nil // data stored in dst.MinimumPaymentTypeRateOrAmount, return on the first match
		} else {
			dst.MinimumPaymentTypeRateOrAmount = nil
			return fmt.Errorf("failed to unmarshal MinimumPaymentPartial as MinimumPaymentTypeRateOrAmount: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MinimumPaymentPartial as MinimumPaymentTypeRateOrAmount: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MinimumPaymentPartial) MarshalJSON() ([]byte, error) {
	if src.MinimumPaymentTypeRateOrAmount != nil {
		return json.Marshal(&src.MinimumPaymentTypeRateOrAmount)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MinimumPaymentPartial) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MinimumPaymentTypeRateOrAmount != nil {
		return obj.MinimumPaymentTypeRateOrAmount
	}

	// all schemas are nil
	return nil
}

type NullableMinimumPaymentPartial struct {
	value *MinimumPaymentPartial
	isSet bool
}

func (v NullableMinimumPaymentPartial) Get() *MinimumPaymentPartial {
	return v.value
}

func (v *NullableMinimumPaymentPartial) Set(val *MinimumPaymentPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimumPaymentPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimumPaymentPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimumPaymentPartial(val *MinimumPaymentPartial) *NullableMinimumPaymentPartial {
	return &NullableMinimumPaymentPartial{value: val, isSet: true}
}

func (v NullableMinimumPaymentPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimumPaymentPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


