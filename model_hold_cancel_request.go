/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// HoldCancelRequest struct for HoldCancelRequest
type HoldCancelRequest struct {
	// The account number associated with the hold
	AccountNo string `json:"account_no"`
	// an unstructured json blob representing additional transaction information supplied by the integrator.
	ExternalData map[string]interface{} `json:"external_data"`
	// The reason for the cancellation
	Reason string `json:"reason"`
	// An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers.
	ReferenceId NullableString `json:"reference_id"`
	// Information received by the transaction risk/fraud service related to this transaction
	RiskInfo map[string]interface{} `json:"risk_info"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data"`
}

// NewHoldCancelRequest instantiates a new HoldCancelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldCancelRequest(accountNo string, externalData map[string]interface{}, reason string, referenceId NullableString, riskInfo map[string]interface{}, userData map[string]interface{}) *HoldCancelRequest {
	this := HoldCancelRequest{}
	this.AccountNo = accountNo
	this.ExternalData = externalData
	this.Reason = reason
	this.ReferenceId = referenceId
	this.RiskInfo = riskInfo
	this.UserData = userData
	return &this
}

// NewHoldCancelRequestWithDefaults instantiates a new HoldCancelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldCancelRequestWithDefaults() *HoldCancelRequest {
	this := HoldCancelRequest{}
	return &this
}

// GetAccountNo returns the AccountNo field value
func (o *HoldCancelRequest) GetAccountNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value
// and a boolean to check if the value has been set.
func (o *HoldCancelRequest) GetAccountNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNo, true
}

// SetAccountNo sets field value
func (o *HoldCancelRequest) SetAccountNo(v string) {
	o.AccountNo = v
}

// GetExternalData returns the ExternalData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *HoldCancelRequest) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldCancelRequest) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ExternalData == nil {
		return nil, false
	}
	return o.ExternalData, true
}

// SetExternalData sets field value
func (o *HoldCancelRequest) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetReason returns the Reason field value
func (o *HoldCancelRequest) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *HoldCancelRequest) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *HoldCancelRequest) SetReason(v string) {
	o.Reason = v
}

// GetReferenceId returns the ReferenceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HoldCancelRequest) GetReferenceId() string {
	if o == nil || o.ReferenceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldCancelRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// SetReferenceId sets field value
func (o *HoldCancelRequest) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}

// GetRiskInfo returns the RiskInfo field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *HoldCancelRequest) GetRiskInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.RiskInfo
}

// GetRiskInfoOk returns a tuple with the RiskInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldCancelRequest) GetRiskInfoOk() (map[string]interface{}, bool) {
	if o == nil || o.RiskInfo == nil {
		return nil, false
	}
	return o.RiskInfo, true
}

// SetRiskInfo sets field value
func (o *HoldCancelRequest) SetRiskInfo(v map[string]interface{}) {
	o.RiskInfo = v
}

// GetUserData returns the UserData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *HoldCancelRequest) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldCancelRequest) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// SetUserData sets field value
func (o *HoldCancelRequest) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

func (o HoldCancelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_no"] = o.AccountNo
	}
	if o.ExternalData != nil {
		toSerialize["external_data"] = o.ExternalData
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["reference_id"] = o.ReferenceId.Get()
	}
	if o.RiskInfo != nil {
		toSerialize["risk_info"] = o.RiskInfo
	}
	if o.UserData != nil {
		toSerialize["user_data"] = o.UserData
	}
	return json.Marshal(toSerialize)
}

type NullableHoldCancelRequest struct {
	value *HoldCancelRequest
	isSet bool
}

func (v NullableHoldCancelRequest) Get() *HoldCancelRequest {
	return v.value
}

func (v *NullableHoldCancelRequest) Set(val *HoldCancelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldCancelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldCancelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldCancelRequest(val *HoldCancelRequest) *NullableHoldCancelRequest {
	return &NullableHoldCancelRequest{value: val, isSet: true}
}

func (v NullableHoldCancelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldCancelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

