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

// EventType - struct for EventType
type EventType struct {
	EventTypeExplicit *EventTypeExplicit
	EventTypeWildcard *EventTypeWildcard
}

// EventTypeExplicitAsEventType is a convenience function that returns EventTypeExplicit wrapped in EventType
func EventTypeExplicitAsEventType(v *EventTypeExplicit) EventType {
	return EventType{
		EventTypeExplicit: v,
	}
}

// EventTypeWildcardAsEventType is a convenience function that returns EventTypeWildcard wrapped in EventType
func EventTypeWildcardAsEventType(v *EventTypeWildcard) EventType {
	return EventType{
		EventTypeWildcard: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EventType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EventTypeExplicit
	err = json.Unmarshal(data, &dst.EventTypeExplicit)
	if err == nil {
		jsonEventTypeExplicit, _ := json.Marshal(dst.EventTypeExplicit)
		if string(jsonEventTypeExplicit) == "{}" { // empty struct
			dst.EventTypeExplicit = nil
		} else {
			match++
		}
	} else {
		dst.EventTypeExplicit = nil
	}

	// try to unmarshal data into EventTypeWildcard
	err = json.Unmarshal(data, &dst.EventTypeWildcard)
	if err == nil {
		jsonEventTypeWildcard, _ := json.Marshal(dst.EventTypeWildcard)
		if string(jsonEventTypeWildcard) == "{}" { // empty struct
			dst.EventTypeWildcard = nil
		} else {
			match++
		}
	} else {
		dst.EventTypeWildcard = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EventTypeExplicit = nil
		dst.EventTypeWildcard = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EventType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EventType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EventType) MarshalJSON() ([]byte, error) {
	if src.EventTypeExplicit != nil {
		return json.Marshal(&src.EventTypeExplicit)
	}

	if src.EventTypeWildcard != nil {
		return json.Marshal(&src.EventTypeWildcard)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EventType) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EventTypeExplicit != nil {
		return obj.EventTypeExplicit
	}

	if obj.EventTypeWildcard != nil {
		return obj.EventTypeWildcard
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj EventType) GetActualInstanceValue() interface{} {
	if obj.EventTypeExplicit != nil {
		return *obj.EventTypeExplicit
	}

	if obj.EventTypeWildcard != nil {
		return *obj.EventTypeWildcard
	}

	// all schemas are nil
	return nil
}

type NullableEventType struct {
	value *EventType
	isSet bool
}

func (v NullableEventType) Get() *EventType {
	return v.value
}

func (v *NullableEventType) Set(val *EventType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventType(val *EventType) *NullableEventType {
	return &NullableEventType{value: val, isSet: true}
}

func (v NullableEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
