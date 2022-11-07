/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// HoldPatchRequest struct for HoldPatchRequest
type HoldPatchRequest struct {
	// The time the transaction will be automatically posted.
	AutoPostAt time.Time `json:"auto_post_at"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	EnhancedTransaction map[string]interface{} `json:"enhanced_transaction,omitempty"`
	// The date that at which this hold is no longer valid.
	ExpiresAt time.Time `json:"expires_at"`
	// an unstructured json blob representing additional transaction information supplied by the integrator.
	ExternalData map[string]interface{} `json:"external_data,omitempty"`
	// An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers.
	ReferenceId string `json:"reference_id"`
	// Information received by the transaction risk/fraud service related to this transaction
	RiskInfo map[string]interface{} `json:"risk_info,omitempty"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data,omitempty"`
}

// NewHoldPatchRequest instantiates a new HoldPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldPatchRequest(autoPostAt time.Time, expiresAt time.Time, referenceId string) *HoldPatchRequest {
	this := HoldPatchRequest{}
	this.AutoPostAt = autoPostAt
	this.ExpiresAt = expiresAt
	this.ReferenceId = referenceId
	return &this
}

// NewHoldPatchRequestWithDefaults instantiates a new HoldPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldPatchRequestWithDefaults() *HoldPatchRequest {
	this := HoldPatchRequest{}
	return &this
}

// GetAutoPostAt returns the AutoPostAt field value
func (o *HoldPatchRequest) GetAutoPostAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AutoPostAt
}

// GetAutoPostAtOk returns a tuple with the AutoPostAt field value
// and a boolean to check if the value has been set.
func (o *HoldPatchRequest) GetAutoPostAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoPostAt, true
}

// SetAutoPostAt sets field value
func (o *HoldPatchRequest) SetAutoPostAt(v time.Time) {
	o.AutoPostAt = v
}

// GetEnhancedTransaction returns the EnhancedTransaction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HoldPatchRequest) GetEnhancedTransaction() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EnhancedTransaction
}

// GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldPatchRequest) GetEnhancedTransactionOk() (map[string]interface{}, bool) {
	if o == nil || o.EnhancedTransaction == nil {
		return nil, false
	}
	return o.EnhancedTransaction, true
}

// HasEnhancedTransaction returns a boolean if a field has been set.
func (o *HoldPatchRequest) HasEnhancedTransaction() bool {
	if o != nil && o.EnhancedTransaction != nil {
		return true
	}

	return false
}

// SetEnhancedTransaction gets a reference to the given map[string]interface{} and assigns it to the EnhancedTransaction field.
func (o *HoldPatchRequest) SetEnhancedTransaction(v map[string]interface{}) {
	o.EnhancedTransaction = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *HoldPatchRequest) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *HoldPatchRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *HoldPatchRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetExternalData returns the ExternalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HoldPatchRequest) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldPatchRequest) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ExternalData == nil {
		return nil, false
	}
	return o.ExternalData, true
}

// HasExternalData returns a boolean if a field has been set.
func (o *HoldPatchRequest) HasExternalData() bool {
	if o != nil && o.ExternalData != nil {
		return true
	}

	return false
}

// SetExternalData gets a reference to the given map[string]interface{} and assigns it to the ExternalData field.
func (o *HoldPatchRequest) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetReferenceId returns the ReferenceId field value
func (o *HoldPatchRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *HoldPatchRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *HoldPatchRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetRiskInfo returns the RiskInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HoldPatchRequest) GetRiskInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RiskInfo
}

// GetRiskInfoOk returns a tuple with the RiskInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldPatchRequest) GetRiskInfoOk() (map[string]interface{}, bool) {
	if o == nil || o.RiskInfo == nil {
		return nil, false
	}
	return o.RiskInfo, true
}

// HasRiskInfo returns a boolean if a field has been set.
func (o *HoldPatchRequest) HasRiskInfo() bool {
	if o != nil && o.RiskInfo != nil {
		return true
	}

	return false
}

// SetRiskInfo gets a reference to the given map[string]interface{} and assigns it to the RiskInfo field.
func (o *HoldPatchRequest) SetRiskInfo(v map[string]interface{}) {
	o.RiskInfo = v
}

// GetUserData returns the UserData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HoldPatchRequest) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldPatchRequest) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *HoldPatchRequest) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]interface{} and assigns it to the UserData field.
func (o *HoldPatchRequest) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

func (o HoldPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auto_post_at"] = o.AutoPostAt
	}
	if o.EnhancedTransaction != nil {
		toSerialize["enhanced_transaction"] = o.EnhancedTransaction
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.ExternalData != nil {
		toSerialize["external_data"] = o.ExternalData
	}
	if true {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if o.RiskInfo != nil {
		toSerialize["risk_info"] = o.RiskInfo
	}
	if o.UserData != nil {
		toSerialize["user_data"] = o.UserData
	}
	return json.Marshal(toSerialize)
}

type NullableHoldPatchRequest struct {
	value *HoldPatchRequest
	isSet bool
}

func (v NullableHoldPatchRequest) Get() *HoldPatchRequest {
	return v.value
}

func (v *NullableHoldPatchRequest) Set(val *HoldPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldPatchRequest(val *HoldPatchRequest) *NullableHoldPatchRequest {
	return &NullableHoldPatchRequest{value: val, isSet: true}
}

func (v NullableHoldPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


