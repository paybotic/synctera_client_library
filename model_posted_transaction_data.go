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

// PostedTransactionData struct for PostedTransactionData
type PostedTransactionData struct {
	// an unstructured json blob representing additional transaction information supplied by the integrator.
	ExternalData map[string]interface{} `json:"external_data"`
	// Whether or not the hold was forced (spending controls ignored)
	ForcePost bool `json:"force_post"`
	// The uuid of the hold (pending transaction) that this transaction originated from, if any.
	HoldId *string `json:"hold_id,omitempty"`
	// The set of accounting entries associated with this transaction. For example, a debit to a customer account will have a corresponding credit in a general ledger account.
	Lines []TransactionLine1 `json:"lines"`
	// A short note to the recipient
	Memo string `json:"memo"`
	Metadata map[string]interface{} `json:"metadata"`
	// The \"original\" transaction that this transaction is related to. This is only populated in the case of reversed transactions.
	OriginalTrx *string `json:"original_trx,omitempty"`
	// Information received by the transaction risk/fraud service related to this transaction
	RiskInfo map[string]interface{} `json:"risk_info"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data"`
}

// NewPostedTransactionData instantiates a new PostedTransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostedTransactionData(externalData map[string]interface{}, forcePost bool, lines []TransactionLine1, memo string, metadata map[string]interface{}, riskInfo map[string]interface{}, userData map[string]interface{}) *PostedTransactionData {
	this := PostedTransactionData{}
	this.ExternalData = externalData
	this.ForcePost = forcePost
	this.Lines = lines
	this.Memo = memo
	this.Metadata = metadata
	this.RiskInfo = riskInfo
	this.UserData = userData
	return &this
}

// NewPostedTransactionDataWithDefaults instantiates a new PostedTransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostedTransactionDataWithDefaults() *PostedTransactionData {
	this := PostedTransactionData{}
	return &this
}

// GetExternalData returns the ExternalData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PostedTransactionData) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostedTransactionData) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ExternalData == nil {
		return nil, false
	}
	return o.ExternalData, true
}

// SetExternalData sets field value
func (o *PostedTransactionData) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetForcePost returns the ForcePost field value
func (o *PostedTransactionData) GetForcePost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value
// and a boolean to check if the value has been set.
func (o *PostedTransactionData) GetForcePostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForcePost, true
}

// SetForcePost sets field value
func (o *PostedTransactionData) SetForcePost(v bool) {
	o.ForcePost = v
}

// GetHoldId returns the HoldId field value if set, zero value otherwise.
func (o *PostedTransactionData) GetHoldId() string {
	if o == nil || o.HoldId == nil {
		var ret string
		return ret
	}
	return *o.HoldId
}

// GetHoldIdOk returns a tuple with the HoldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostedTransactionData) GetHoldIdOk() (*string, bool) {
	if o == nil || o.HoldId == nil {
		return nil, false
	}
	return o.HoldId, true
}

// HasHoldId returns a boolean if a field has been set.
func (o *PostedTransactionData) HasHoldId() bool {
	if o != nil && o.HoldId != nil {
		return true
	}

	return false
}

// SetHoldId gets a reference to the given string and assigns it to the HoldId field.
func (o *PostedTransactionData) SetHoldId(v string) {
	o.HoldId = &v
}

// GetLines returns the Lines field value
func (o *PostedTransactionData) GetLines() []TransactionLine1 {
	if o == nil {
		var ret []TransactionLine1
		return ret
	}

	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value
// and a boolean to check if the value has been set.
func (o *PostedTransactionData) GetLinesOk() ([]TransactionLine1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lines, true
}

// SetLines sets field value
func (o *PostedTransactionData) SetLines(v []TransactionLine1) {
	o.Lines = v
}

// GetMemo returns the Memo field value
func (o *PostedTransactionData) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *PostedTransactionData) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *PostedTransactionData) SetMemo(v string) {
	o.Memo = v
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PostedTransactionData) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostedTransactionData) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *PostedTransactionData) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetOriginalTrx returns the OriginalTrx field value if set, zero value otherwise.
func (o *PostedTransactionData) GetOriginalTrx() string {
	if o == nil || o.OriginalTrx == nil {
		var ret string
		return ret
	}
	return *o.OriginalTrx
}

// GetOriginalTrxOk returns a tuple with the OriginalTrx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostedTransactionData) GetOriginalTrxOk() (*string, bool) {
	if o == nil || o.OriginalTrx == nil {
		return nil, false
	}
	return o.OriginalTrx, true
}

// HasOriginalTrx returns a boolean if a field has been set.
func (o *PostedTransactionData) HasOriginalTrx() bool {
	if o != nil && o.OriginalTrx != nil {
		return true
	}

	return false
}

// SetOriginalTrx gets a reference to the given string and assigns it to the OriginalTrx field.
func (o *PostedTransactionData) SetOriginalTrx(v string) {
	o.OriginalTrx = &v
}

// GetRiskInfo returns the RiskInfo field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PostedTransactionData) GetRiskInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.RiskInfo
}

// GetRiskInfoOk returns a tuple with the RiskInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostedTransactionData) GetRiskInfoOk() (map[string]interface{}, bool) {
	if o == nil || o.RiskInfo == nil {
		return nil, false
	}
	return o.RiskInfo, true
}

// SetRiskInfo sets field value
func (o *PostedTransactionData) SetRiskInfo(v map[string]interface{}) {
	o.RiskInfo = v
}

// GetUserData returns the UserData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PostedTransactionData) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostedTransactionData) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// SetUserData sets field value
func (o *PostedTransactionData) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

func (o PostedTransactionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalData != nil {
		toSerialize["external_data"] = o.ExternalData
	}
	if true {
		toSerialize["force_post"] = o.ForcePost
	}
	if o.HoldId != nil {
		toSerialize["hold_id"] = o.HoldId
	}
	if true {
		toSerialize["lines"] = o.Lines
	}
	if true {
		toSerialize["memo"] = o.Memo
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OriginalTrx != nil {
		toSerialize["original_trx"] = o.OriginalTrx
	}
	if o.RiskInfo != nil {
		toSerialize["risk_info"] = o.RiskInfo
	}
	if o.UserData != nil {
		toSerialize["user_data"] = o.UserData
	}
	return json.Marshal(toSerialize)
}

type NullablePostedTransactionData struct {
	value *PostedTransactionData
	isSet bool
}

func (v NullablePostedTransactionData) Get() *PostedTransactionData {
	return v.value
}

func (v *NullablePostedTransactionData) Set(val *PostedTransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePostedTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePostedTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostedTransactionData(val *PostedTransactionData) *NullablePostedTransactionData {
	return &NullablePostedTransactionData{value: val, isSet: true}
}

func (v NullablePostedTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostedTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


