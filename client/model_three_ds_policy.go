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

// ThreeDsPolicy EMV 3-D Secure policy. Whenever a 3DS decision gateway refers to a card product, the policy for that card product is always DECISION_GATEWAY.  Policy           | Description ---------------- | ----------- SMS_OTP          | Use the card holder's phone number on file to perform advanced authentication via SMS EXEMPT           | Transactions will be exempted from advanced authentication DECISION_GATEWAY | Fintech 3DS decision gateway will decide the 3DS action for each transaction
type ThreeDsPolicy string

// List of three_ds_policy
const (
	THREEDSPOLICY_DECISION_GATEWAY ThreeDsPolicy = "DECISION_GATEWAY"
	THREEDSPOLICY_EXEMPT           ThreeDsPolicy = "EXEMPT"
	THREEDSPOLICY_SMS_OTP          ThreeDsPolicy = "SMS_OTP"
)

// All allowed values of ThreeDsPolicy enum
var AllowedThreeDsPolicyEnumValues = []ThreeDsPolicy{
	"DECISION_GATEWAY",
	"EXEMPT",
	"SMS_OTP",
}

func (v *ThreeDsPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThreeDsPolicy(value)
	for _, existing := range AllowedThreeDsPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThreeDsPolicy", value)
}

// NewThreeDsPolicyFromValue returns a pointer to a valid ThreeDsPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewThreeDsPolicyFromValue(v string) (*ThreeDsPolicy, error) {
	ev := ThreeDsPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ThreeDsPolicy: valid values are %v", v, AllowedThreeDsPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ThreeDsPolicy) IsValid() bool {
	for _, existing := range AllowedThreeDsPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to three_ds_policy value
func (v ThreeDsPolicy) Ptr() *ThreeDsPolicy {
	return &v
}

type NullableThreeDsPolicy struct {
	value *ThreeDsPolicy
	isSet bool
}

func (v NullableThreeDsPolicy) Get() *ThreeDsPolicy {
	return v.value
}

func (v *NullableThreeDsPolicy) Set(val *ThreeDsPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDsPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDsPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDsPolicy(val *ThreeDsPolicy) *NullableThreeDsPolicy {
	return &NullableThreeDsPolicy{value: val, isSet: true}
}

func (v NullableThreeDsPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDsPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
