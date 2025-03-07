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

// PaymentSubType the model 'PaymentSubType'
type PaymentSubType string

// List of payment_sub_type
const (
	PAYMENTSUBTYPE_ACH_INCOMING_CREDIT                       PaymentSubType = "ACH.INCOMING_CREDIT"
	PAYMENTSUBTYPE_ACH_INCOMING_DEBIT                        PaymentSubType = "ACH.INCOMING_DEBIT"
	PAYMENTSUBTYPE_ACH_OUTGOING_CREDIT                       PaymentSubType = "ACH.OUTGOING_CREDIT"
	PAYMENTSUBTYPE_ACH_OUTGOING_DEBIT                        PaymentSubType = "ACH.OUTGOING_DEBIT"
	PAYMENTSUBTYPE_CARD_ATM_WITHDRAWAL                       PaymentSubType = "CARD.ATM_WITHDRAWAL"
	PAYMENTSUBTYPE_CARD_POS_CASHBACK                         PaymentSubType = "CARD.POS_CASHBACK"
	PAYMENTSUBTYPE_SYNCTERA_PAY_INCOMING_INTERAC_AUTODEPOSIT PaymentSubType = "SYNCTERA_PAY.INCOMING_INTERAC_AUTODEPOSIT"
	PAYMENTSUBTYPE_SYNCTERA_PAY_OUTGOING_BILL_PAYMENT        PaymentSubType = "SYNCTERA_PAY.OUTGOING_BILL_PAYMENT"
	PAYMENTSUBTYPE_SYNCTERA_PAY_OUTGOING_INTERAC_E_TRANSFER  PaymentSubType = "SYNCTERA_PAY.OUTGOING_INTERAC_E_TRANSFER"
)

// All allowed values of PaymentSubType enum
var AllowedPaymentSubTypeEnumValues = []PaymentSubType{
	"ACH.INCOMING_CREDIT",
	"ACH.INCOMING_DEBIT",
	"ACH.OUTGOING_CREDIT",
	"ACH.OUTGOING_DEBIT",
	"CARD.ATM_WITHDRAWAL",
	"CARD.POS_CASHBACK",
	"SYNCTERA_PAY.INCOMING_INTERAC_AUTODEPOSIT",
	"SYNCTERA_PAY.OUTGOING_BILL_PAYMENT",
	"SYNCTERA_PAY.OUTGOING_INTERAC_E_TRANSFER",
}

func (v *PaymentSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentSubType(value)
	for _, existing := range AllowedPaymentSubTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentSubType", value)
}

// NewPaymentSubTypeFromValue returns a pointer to a valid PaymentSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentSubTypeFromValue(v string) (*PaymentSubType, error) {
	ev := PaymentSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentSubType: valid values are %v", v, AllowedPaymentSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentSubType) IsValid() bool {
	for _, existing := range AllowedPaymentSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to payment_sub_type value
func (v PaymentSubType) Ptr() *PaymentSubType {
	return &v
}

type NullablePaymentSubType struct {
	value *PaymentSubType
	isSet bool
}

func (v NullablePaymentSubType) Get() *PaymentSubType {
	return v.value
}

func (v *NullablePaymentSubType) Set(val *PaymentSubType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentSubType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentSubType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentSubType(val *PaymentSubType) *NullablePaymentSubType {
	return &NullablePaymentSubType{value: val, isSet: true}
}

func (v NullablePaymentSubType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentSubType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
