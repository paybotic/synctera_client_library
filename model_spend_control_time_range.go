/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SpendControlTimeRange - The time range to which the spend control applies
type SpendControlTimeRange struct {
	SpendControlRollingWindowDays *SpendControlRollingWindowDays
}

// SpendControlRollingWindowDaysAsSpendControlTimeRange is a convenience function that returns SpendControlRollingWindowDays wrapped in SpendControlTimeRange
func SpendControlRollingWindowDaysAsSpendControlTimeRange(v *SpendControlRollingWindowDays) SpendControlTimeRange {
	return SpendControlTimeRange{
		SpendControlRollingWindowDays: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SpendControlTimeRange) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ROLLING_WINDOW_DAYS'
	if jsonDict["time_range_type"] == "ROLLING_WINDOW_DAYS" {
		// try to unmarshal JSON data into SpendControlRollingWindowDays
		err = json.Unmarshal(data, &dst.SpendControlRollingWindowDays)
		if err == nil {
			return nil // data stored in dst.SpendControlRollingWindowDays, return on the first match
		} else {
			dst.SpendControlRollingWindowDays = nil
			return fmt.Errorf("Failed to unmarshal SpendControlTimeRange as SpendControlRollingWindowDays: %s", err.Error())
		}
	}

	// check if the discriminator value is 'spend_control_rolling_window_days'
	if jsonDict["time_range_type"] == "spend_control_rolling_window_days" {
		// try to unmarshal JSON data into SpendControlRollingWindowDays
		err = json.Unmarshal(data, &dst.SpendControlRollingWindowDays)
		if err == nil {
			return nil // data stored in dst.SpendControlRollingWindowDays, return on the first match
		} else {
			dst.SpendControlRollingWindowDays = nil
			return fmt.Errorf("Failed to unmarshal SpendControlTimeRange as SpendControlRollingWindowDays: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SpendControlTimeRange) MarshalJSON() ([]byte, error) {
	if src.SpendControlRollingWindowDays != nil {
		return json.Marshal(&src.SpendControlRollingWindowDays)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SpendControlTimeRange) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SpendControlRollingWindowDays != nil {
		return obj.SpendControlRollingWindowDays
	}

	// all schemas are nil
	return nil
}

type NullableSpendControlTimeRange struct {
	value *SpendControlTimeRange
	isSet bool
}

func (v NullableSpendControlTimeRange) Get() *SpendControlTimeRange {
	return v.value
}

func (v *NullableSpendControlTimeRange) Set(val *SpendControlTimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendControlTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendControlTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendControlTimeRange(val *SpendControlTimeRange) *NullableSpendControlTimeRange {
	return &NullableSpendControlTimeRange{value: val, isSet: true}
}

func (v NullableSpendControlTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendControlTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


