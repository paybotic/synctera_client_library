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

// CardFulfillmentStatus The status indicating the state of the card issuance
type CardFulfillmentStatus string

// List of card_fulfillment_status
const (
	CARDFULFILLMENTSTATUS_DIGITALLY_PRESENTED CardFulfillmentStatus = "DIGITALLY_PRESENTED"
	CARDFULFILLMENTSTATUS_ISSUED              CardFulfillmentStatus = "ISSUED"
	CARDFULFILLMENTSTATUS_ORDERED             CardFulfillmentStatus = "ORDERED"
	CARDFULFILLMENTSTATUS_REISSUED            CardFulfillmentStatus = "REISSUED"
	CARDFULFILLMENTSTATUS_REJECTED            CardFulfillmentStatus = "REJECTED"
	CARDFULFILLMENTSTATUS_REORDERED           CardFulfillmentStatus = "REORDERED"
	CARDFULFILLMENTSTATUS_SHIPPED             CardFulfillmentStatus = "SHIPPED"
)

// All allowed values of CardFulfillmentStatus enum
var AllowedCardFulfillmentStatusEnumValues = []CardFulfillmentStatus{
	"DIGITALLY_PRESENTED",
	"ISSUED",
	"ORDERED",
	"REISSUED",
	"REJECTED",
	"REORDERED",
	"SHIPPED",
}

func (v *CardFulfillmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardFulfillmentStatus(value)
	for _, existing := range AllowedCardFulfillmentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardFulfillmentStatus", value)
}

// NewCardFulfillmentStatusFromValue returns a pointer to a valid CardFulfillmentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardFulfillmentStatusFromValue(v string) (*CardFulfillmentStatus, error) {
	ev := CardFulfillmentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardFulfillmentStatus: valid values are %v", v, AllowedCardFulfillmentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardFulfillmentStatus) IsValid() bool {
	for _, existing := range AllowedCardFulfillmentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to card_fulfillment_status value
func (v CardFulfillmentStatus) Ptr() *CardFulfillmentStatus {
	return &v
}

type NullableCardFulfillmentStatus struct {
	value *CardFulfillmentStatus
	isSet bool
}

func (v NullableCardFulfillmentStatus) Get() *CardFulfillmentStatus {
	return v.value
}

func (v *NullableCardFulfillmentStatus) Set(val *CardFulfillmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCardFulfillmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCardFulfillmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardFulfillmentStatus(val *CardFulfillmentStatus) *NullableCardFulfillmentStatus {
	return &NullableCardFulfillmentStatus{value: val, isSet: true}
}

func (v NullableCardFulfillmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardFulfillmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
