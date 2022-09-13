/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// HoldPostingRequest struct for HoldPostingRequest
type HoldPostingRequest struct {
	// The account number associated with the hold
	AccountNo string `json:"account_no"`
	// The amount of the hold.
	Amount int64 `json:"amount"`
	// The effective date of the transaction once it gets posted
	EffectiveDate time.Time `json:"effective_date"`
	// an unstructured json blob representing additional transaction information supplied by the integrator.
	ExternalData map[string]interface{} `json:"external_data"`
	// The amount of the hold.
	HoldAmount int64 `json:"hold_amount"`
	// Information received by the transaction risk/fraud service related to this transaction
	RiskInfo map[string]interface{} `json:"risk_info"`
	// The specific transaction type. For example, for `ach`, this may be \"outgoing_debit\".
	Subtype string `json:"subtype"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data"`
}

// NewHoldPostingRequest instantiates a new HoldPostingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldPostingRequest(accountNo string, amount int64, effectiveDate time.Time, externalData map[string]interface{}, holdAmount int64, riskInfo map[string]interface{}, subtype string, userData map[string]interface{}) *HoldPostingRequest {
	this := HoldPostingRequest{}
	this.AccountNo = accountNo
	this.Amount = amount
	this.EffectiveDate = effectiveDate
	this.ExternalData = externalData
	this.HoldAmount = holdAmount
	this.RiskInfo = riskInfo
	this.Subtype = subtype
	this.UserData = userData
	return &this
}

// NewHoldPostingRequestWithDefaults instantiates a new HoldPostingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldPostingRequestWithDefaults() *HoldPostingRequest {
	this := HoldPostingRequest{}
	return &this
}

// GetAccountNo returns the AccountNo field value
func (o *HoldPostingRequest) GetAccountNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value
// and a boolean to check if the value has been set.
func (o *HoldPostingRequest) GetAccountNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNo, true
}

// SetAccountNo sets field value
func (o *HoldPostingRequest) SetAccountNo(v string) {
	o.AccountNo = v
}

// GetAmount returns the Amount field value
func (o *HoldPostingRequest) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *HoldPostingRequest) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *HoldPostingRequest) SetAmount(v int64) {
	o.Amount = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *HoldPostingRequest) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *HoldPostingRequest) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *HoldPostingRequest) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetExternalData returns the ExternalData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *HoldPostingRequest) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldPostingRequest) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ExternalData == nil {
		return nil, false
	}
	return o.ExternalData, true
}

// SetExternalData sets field value
func (o *HoldPostingRequest) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetHoldAmount returns the HoldAmount field value
func (o *HoldPostingRequest) GetHoldAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HoldAmount
}

// GetHoldAmountOk returns a tuple with the HoldAmount field value
// and a boolean to check if the value has been set.
func (o *HoldPostingRequest) GetHoldAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HoldAmount, true
}

// SetHoldAmount sets field value
func (o *HoldPostingRequest) SetHoldAmount(v int64) {
	o.HoldAmount = v
}

// GetRiskInfo returns the RiskInfo field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *HoldPostingRequest) GetRiskInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.RiskInfo
}

// GetRiskInfoOk returns a tuple with the RiskInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldPostingRequest) GetRiskInfoOk() (map[string]interface{}, bool) {
	if o == nil || o.RiskInfo == nil {
		return nil, false
	}
	return o.RiskInfo, true
}

// SetRiskInfo sets field value
func (o *HoldPostingRequest) SetRiskInfo(v map[string]interface{}) {
	o.RiskInfo = v
}

// GetSubtype returns the Subtype field value
func (o *HoldPostingRequest) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *HoldPostingRequest) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *HoldPostingRequest) SetSubtype(v string) {
	o.Subtype = v
}

// GetUserData returns the UserData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *HoldPostingRequest) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldPostingRequest) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// SetUserData sets field value
func (o *HoldPostingRequest) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

func (o HoldPostingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_no"] = o.AccountNo
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["effective_date"] = o.EffectiveDate
	}
	if o.ExternalData != nil {
		toSerialize["external_data"] = o.ExternalData
	}
	if true {
		toSerialize["hold_amount"] = o.HoldAmount
	}
	if o.RiskInfo != nil {
		toSerialize["risk_info"] = o.RiskInfo
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if o.UserData != nil {
		toSerialize["user_data"] = o.UserData
	}
	return json.Marshal(toSerialize)
}

type NullableHoldPostingRequest struct {
	value *HoldPostingRequest
	isSet bool
}

func (v NullableHoldPostingRequest) Get() *HoldPostingRequest {
	return v.value
}

func (v *NullableHoldPostingRequest) Set(val *HoldPostingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldPostingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldPostingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldPostingRequest(val *HoldPostingRequest) *NullableHoldPostingRequest {
	return &NullableHoldPostingRequest{value: val, isSet: true}
}

func (v NullableHoldPostingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldPostingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


