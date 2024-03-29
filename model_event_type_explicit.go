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

// EventTypeExplicit All the webhook event types
type EventTypeExplicit string

// List of event_type_explicit
const (
	EVENTTYPEEXPLICIT_ACCOUNT_CREATED EventTypeExplicit = "ACCOUNT.CREATED"
	EVENTTYPEEXPLICIT_ACCOUNT_UPDATED EventTypeExplicit = "ACCOUNT.UPDATED"
	EVENTTYPEEXPLICIT_CARD_UPDATED EventTypeExplicit = "CARD.UPDATED"
	EVENTTYPEEXPLICIT_CARD_IMAGE_UPDATED EventTypeExplicit = "CARD.IMAGE.UPDATED"
	EVENTTYPEEXPLICIT_CASE_UPDATED EventTypeExplicit = "CASE.UPDATED"
	EVENTTYPEEXPLICIT_CUSTOMER_UPDATED EventTypeExplicit = "CUSTOMER.UPDATED"
	EVENTTYPEEXPLICIT_CUSTOMER_KYC_OUTCOME_UPDATED EventTypeExplicit = "CUSTOMER.KYC_OUTCOME.UPDATED"
	EVENTTYPEEXPLICIT_BUSINESS_VERIFICATION_OUTCOME_UPDATED EventTypeExplicit = "BUSINESS.VERIFICATION_OUTCOME.UPDATED"
	EVENTTYPEEXPLICIT_PERSON_VERIFICATION_OUTCOME_UPDATED EventTypeExplicit = "PERSON.VERIFICATION_OUTCOME.UPDATED"
	EVENTTYPEEXPLICIT_INTEREST_MONTHLY_PAYOUT EventTypeExplicit = "INTEREST.MONTHLY_PAYOUT"
	EVENTTYPEEXPLICIT_INTERNAL_TRANSFER_SUCCEEDED EventTypeExplicit = "INTERNAL_TRANSFER.SUCCEEDED"
	EVENTTYPEEXPLICIT_TRANSACTION_POSTED_CREATED EventTypeExplicit = "TRANSACTION.POSTED.CREATED"
	EVENTTYPEEXPLICIT_TRANSACTION_POSTED_UPDATED EventTypeExplicit = "TRANSACTION.POSTED.UPDATED"
	EVENTTYPEEXPLICIT_TRANSACTION_PENDING_CREATED EventTypeExplicit = "TRANSACTION.PENDING.CREATED"
	EVENTTYPEEXPLICIT_TRANSACTION_PENDING_UPDATED EventTypeExplicit = "TRANSACTION.PENDING.UPDATED"
	EVENTTYPEEXPLICIT_CARD_DIGITALWALLETTOKEN_CREATED EventTypeExplicit = "CARD.DIGITALWALLETTOKEN.CREATED"
	EVENTTYPEEXPLICIT_CARD_DIGITALWALLETTOKEN_UPDATED EventTypeExplicit = "CARD.DIGITALWALLETTOKEN.UPDATED"
	EVENTTYPEEXPLICIT_PAYMENT_SCHEDULE_CREATED EventTypeExplicit = "PAYMENT_SCHEDULE.CREATED"
	EVENTTYPEEXPLICIT_PAYMENT_SCHEDULE_UPDATED EventTypeExplicit = "PAYMENT_SCHEDULE.UPDATED"
	EVENTTYPEEXPLICIT_PAYMENT_SCHEDULE_PAYMENT_CREATED EventTypeExplicit = "PAYMENT_SCHEDULE.PAYMENT.CREATED"
	EVENTTYPEEXPLICIT_STATEMENT_CREATED EventTypeExplicit = "STATEMENT.CREATED"
	EVENTTYPEEXPLICIT_EXTERNAL_CARD_CREATED EventTypeExplicit = "EXTERNAL_CARD.CREATED"
	EVENTTYPEEXPLICIT_EXTERNAL_CARD_TRANSFER_CREATED EventTypeExplicit = "EXTERNAL_CARD_TRANSFER.CREATED"
	EVENTTYPEEXPLICIT_APPLICATION_CREATED EventTypeExplicit = "APPLICATION.CREATED"
	EVENTTYPEEXPLICIT_APPLICATION_UPDATED EventTypeExplicit = "APPLICATION.UPDATED"
	EVENTTYPEEXPLICIT_CASH_PICKUP_CREATED EventTypeExplicit = "CASH_PICKUP.CREATED"
	EVENTTYPEEXPLICIT_CASH_PICKUP_UPDATED EventTypeExplicit = "CASH_PICKUP.UPDATED"
)

// All allowed values of EventTypeExplicit enum
var AllowedEventTypeExplicitEnumValues = []EventTypeExplicit{
	"ACCOUNT.CREATED",
	"ACCOUNT.UPDATED",
	"CARD.UPDATED",
	"CARD.IMAGE.UPDATED",
	"CASE.UPDATED",
	"CUSTOMER.UPDATED",
	"CUSTOMER.KYC_OUTCOME.UPDATED",
	"BUSINESS.VERIFICATION_OUTCOME.UPDATED",
	"PERSON.VERIFICATION_OUTCOME.UPDATED",
	"INTEREST.MONTHLY_PAYOUT",
	"INTERNAL_TRANSFER.SUCCEEDED",
	"TRANSACTION.POSTED.CREATED",
	"TRANSACTION.POSTED.UPDATED",
	"TRANSACTION.PENDING.CREATED",
	"TRANSACTION.PENDING.UPDATED",
	"CARD.DIGITALWALLETTOKEN.CREATED",
	"CARD.DIGITALWALLETTOKEN.UPDATED",
	"PAYMENT_SCHEDULE.CREATED",
	"PAYMENT_SCHEDULE.UPDATED",
	"PAYMENT_SCHEDULE.PAYMENT.CREATED",
	"STATEMENT.CREATED",
	"EXTERNAL_CARD.CREATED",
	"EXTERNAL_CARD_TRANSFER.CREATED",
	"APPLICATION.CREATED",
	"APPLICATION.UPDATED",
	"CASH_PICKUP.CREATED",
	"CASH_PICKUP.UPDATED",
}

func (v *EventTypeExplicit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventTypeExplicit(value)
	for _, existing := range AllowedEventTypeExplicitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventTypeExplicit", value)
}

// NewEventTypeExplicitFromValue returns a pointer to a valid EventTypeExplicit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypeExplicitFromValue(v string) (*EventTypeExplicit, error) {
	ev := EventTypeExplicit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventTypeExplicit: valid values are %v", v, AllowedEventTypeExplicitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventTypeExplicit) IsValid() bool {
	for _, existing := range AllowedEventTypeExplicitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_type_explicit value
func (v EventTypeExplicit) Ptr() *EventTypeExplicit {
	return &v
}

type NullableEventTypeExplicit struct {
	value *EventTypeExplicit
	isSet bool
}

func (v NullableEventTypeExplicit) Get() *EventTypeExplicit {
	return v.value
}

func (v *NullableEventTypeExplicit) Set(val *EventTypeExplicit) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeExplicit) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeExplicit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeExplicit(val *EventTypeExplicit) *NullableEventTypeExplicit {
	return &NullableEventTypeExplicit{value: val, isSet: true}
}

func (v NullableEventTypeExplicit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeExplicit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

