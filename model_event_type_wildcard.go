/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// EventTypeWildcard the model 'EventTypeWildcard'
type EventTypeWildcard string

// List of event_type_wildcard
const (
	EVENTTYPEWILDCARD_ACCOUNT                EventTypeWildcard = "ACCOUNT.*"
	EVENTTYPEWILDCARD_ACH                    EventTypeWildcard = "ACH.*"
	EVENTTYPEWILDCARD_APPLICATION            EventTypeWildcard = "APPLICATION.*"
	EVENTTYPEWILDCARD_BUSINESS               EventTypeWildcard = "BUSINESS.*"
	EVENTTYPEWILDCARD_CARD                   EventTypeWildcard = "CARD.*"
	EVENTTYPEWILDCARD_CASE                   EventTypeWildcard = "CASE.*"
	EVENTTYPEWILDCARD_CASH_PICKUP            EventTypeWildcard = "CASH_PICKUP.*"
	EVENTTYPEWILDCARD_CREDIT_SCORE           EventTypeWildcard = "CREDIT_SCORE.*"
	EVENTTYPEWILDCARD_CUSTOMER               EventTypeWildcard = "CUSTOMER.*"
	EVENTTYPEWILDCARD_EXTERNAL_CARD          EventTypeWildcard = "EXTERNAL_CARD.*"
	EVENTTYPEWILDCARD_EXTERNAL_CARD_TRANSFER EventTypeWildcard = "EXTERNAL_CARD_TRANSFER.*"
	EVENTTYPEWILDCARD_INTEREST               EventTypeWildcard = "INTEREST.*"
	EVENTTYPEWILDCARD_INTERNAL_TRANSFER      EventTypeWildcard = "INTERNAL_TRANSFER.*"
	EVENTTYPEWILDCARD_NOTE                   EventTypeWildcard = "NOTE.*"
	EVENTTYPEWILDCARD_PAYMENT_SCHEDULE       EventTypeWildcard = "PAYMENT_SCHEDULE.*"
	EVENTTYPEWILDCARD_PERSON                 EventTypeWildcard = "PERSON.*"
	EVENTTYPEWILDCARD_STATEMENT              EventTypeWildcard = "STATEMENT.*"
	EVENTTYPEWILDCARD_TRANSACTION            EventTypeWildcard = "TRANSACTION.*"
)

// All allowed values of EventTypeWildcard enum
var AllowedEventTypeWildcardEnumValues = []EventTypeWildcard{
	"ACCOUNT.*",
	"ACH.*",
	"APPLICATION.*",
	"BUSINESS.*",
	"CARD.*",
	"CASE.*",
	"CASH_PICKUP.*",
	"CREDIT_SCORE.*",
	"CUSTOMER.*",
	"EXTERNAL_CARD.*",
	"EXTERNAL_CARD_TRANSFER.*",
	"INTEREST.*",
	"INTERNAL_TRANSFER.*",
	"NOTE.*",
	"PAYMENT_SCHEDULE.*",
	"PERSON.*",
	"STATEMENT.*",
	"TRANSACTION.*",
}

func (v *EventTypeWildcard) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventTypeWildcard(value)
	for _, existing := range AllowedEventTypeWildcardEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventTypeWildcard", value)
}

// NewEventTypeWildcardFromValue returns a pointer to a valid EventTypeWildcard
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypeWildcardFromValue(v string) (*EventTypeWildcard, error) {
	ev := EventTypeWildcard(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventTypeWildcard: valid values are %v", v, AllowedEventTypeWildcardEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventTypeWildcard) IsValid() bool {
	for _, existing := range AllowedEventTypeWildcardEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_type_wildcard value
func (v EventTypeWildcard) Ptr() *EventTypeWildcard {
	return &v
}

type NullableEventTypeWildcard struct {
	value *EventTypeWildcard
	isSet bool
}

func (v NullableEventTypeWildcard) Get() *EventTypeWildcard {
	return v.value
}

func (v *NullableEventTypeWildcard) Set(val *EventTypeWildcard) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeWildcard) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeWildcard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeWildcard(val *EventTypeWildcard) *NullableEventTypeWildcard {
	return &NullableEventTypeWildcard{value: val, isSet: true}
}

func (v NullableEventTypeWildcard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeWildcard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
