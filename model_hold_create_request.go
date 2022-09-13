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

// HoldCreateRequest struct for HoldCreateRequest
type HoldCreateRequest struct {
	// The account number associated with the hold
	AccountNo string `json:"account_no"`
	AllowPartial bool `json:"allow_partial"`
	// The amount of the hold.
	Amount int64 `json:"amount"`
	// The time the transaction will be automatically posted.
	AutoPostAt time.Time `json:"auto_post_at"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	DcSign DcSign `json:"dc_sign"`
	// The reason for the decline, if any
	DeclineReason string `json:"decline_reason"`
	// The description of the transaction
	Description *string `json:"description,omitempty"`
	// The effective date of the transaction once it gets posted
	EffectiveDate time.Time `json:"effective_date"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	EnhancedTransaction map[string]interface{} `json:"enhanced_transaction,omitempty"`
	// The date that at which this hold is no longer valid.
	ExpiresAt time.Time `json:"expires_at"`
	// an unstructured json blob representing additional transaction information supplied by the integrator.
	ExternalData map[string]interface{} `json:"external_data"`
	// Whether or not the hold was forced (spending controls ignored)
	ForcePost bool `json:"force_post"`
	// A short note to the recipient
	Memo string `json:"memo"`
	// The network this transaction is associated with
	Network string `json:"network"`
	// The description of the offset transaction
	OffsetDescription *string `json:"offset_description,omitempty"`
	// An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers.
	ReferenceId NullableString `json:"reference_id"`
	// The status of the hold.
	Status string `json:"status"`
	// The specific transaction type. For example, for `ach`, this may be \"outgoing_debit\".
	Subtype string `json:"subtype"`
	// The time the transaction occurred.
	TransactionTime time.Time `json:"transaction_time"`
	// The general type of transaction. For example, \"card\" or \"ach\".
	Type string `json:"type"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data"`
}

// NewHoldCreateRequest instantiates a new HoldCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldCreateRequest(accountNo string, allowPartial bool, amount int64, autoPostAt time.Time, currency string, dcSign DcSign, declineReason string, effectiveDate time.Time, expiresAt time.Time, externalData map[string]interface{}, forcePost bool, memo string, network string, referenceId NullableString, status string, subtype string, transactionTime time.Time, type_ string, userData map[string]interface{}) *HoldCreateRequest {
	this := HoldCreateRequest{}
	this.AccountNo = accountNo
	this.AllowPartial = allowPartial
	this.Amount = amount
	this.AutoPostAt = autoPostAt
	this.Currency = currency
	this.DcSign = dcSign
	this.DeclineReason = declineReason
	this.EffectiveDate = effectiveDate
	this.ExpiresAt = expiresAt
	this.ExternalData = externalData
	this.ForcePost = forcePost
	this.Memo = memo
	this.Network = network
	this.ReferenceId = referenceId
	this.Status = status
	this.Subtype = subtype
	this.TransactionTime = transactionTime
	this.Type = type_
	this.UserData = userData
	return &this
}

// NewHoldCreateRequestWithDefaults instantiates a new HoldCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldCreateRequestWithDefaults() *HoldCreateRequest {
	this := HoldCreateRequest{}
	return &this
}

// GetAccountNo returns the AccountNo field value
func (o *HoldCreateRequest) GetAccountNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetAccountNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNo, true
}

// SetAccountNo sets field value
func (o *HoldCreateRequest) SetAccountNo(v string) {
	o.AccountNo = v
}

// GetAllowPartial returns the AllowPartial field value
func (o *HoldCreateRequest) GetAllowPartial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowPartial
}

// GetAllowPartialOk returns a tuple with the AllowPartial field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetAllowPartialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowPartial, true
}

// SetAllowPartial sets field value
func (o *HoldCreateRequest) SetAllowPartial(v bool) {
	o.AllowPartial = v
}

// GetAmount returns the Amount field value
func (o *HoldCreateRequest) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *HoldCreateRequest) SetAmount(v int64) {
	o.Amount = v
}

// GetAutoPostAt returns the AutoPostAt field value
func (o *HoldCreateRequest) GetAutoPostAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AutoPostAt
}

// GetAutoPostAtOk returns a tuple with the AutoPostAt field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetAutoPostAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoPostAt, true
}

// SetAutoPostAt sets field value
func (o *HoldCreateRequest) SetAutoPostAt(v time.Time) {
	o.AutoPostAt = v
}

// GetCurrency returns the Currency field value
func (o *HoldCreateRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *HoldCreateRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetDcSign returns the DcSign field value
func (o *HoldCreateRequest) GetDcSign() DcSign {
	if o == nil {
		var ret DcSign
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetDcSignOk() (*DcSign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *HoldCreateRequest) SetDcSign(v DcSign) {
	o.DcSign = v
}

// GetDeclineReason returns the DeclineReason field value
func (o *HoldCreateRequest) GetDeclineReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeclineReason
}

// GetDeclineReasonOk returns a tuple with the DeclineReason field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetDeclineReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeclineReason, true
}

// SetDeclineReason sets field value
func (o *HoldCreateRequest) SetDeclineReason(v string) {
	o.DeclineReason = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HoldCreateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HoldCreateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HoldCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *HoldCreateRequest) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *HoldCreateRequest) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetEnhancedTransaction returns the EnhancedTransaction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HoldCreateRequest) GetEnhancedTransaction() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EnhancedTransaction
}

// GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldCreateRequest) GetEnhancedTransactionOk() (map[string]interface{}, bool) {
	if o == nil || o.EnhancedTransaction == nil {
		return nil, false
	}
	return o.EnhancedTransaction, true
}

// HasEnhancedTransaction returns a boolean if a field has been set.
func (o *HoldCreateRequest) HasEnhancedTransaction() bool {
	if o != nil && o.EnhancedTransaction != nil {
		return true
	}

	return false
}

// SetEnhancedTransaction gets a reference to the given map[string]interface{} and assigns it to the EnhancedTransaction field.
func (o *HoldCreateRequest) SetEnhancedTransaction(v map[string]interface{}) {
	o.EnhancedTransaction = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *HoldCreateRequest) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *HoldCreateRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetExternalData returns the ExternalData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *HoldCreateRequest) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldCreateRequest) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ExternalData == nil {
		return nil, false
	}
	return o.ExternalData, true
}

// SetExternalData sets field value
func (o *HoldCreateRequest) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetForcePost returns the ForcePost field value
func (o *HoldCreateRequest) GetForcePost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetForcePostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForcePost, true
}

// SetForcePost sets field value
func (o *HoldCreateRequest) SetForcePost(v bool) {
	o.ForcePost = v
}

// GetMemo returns the Memo field value
func (o *HoldCreateRequest) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *HoldCreateRequest) SetMemo(v string) {
	o.Memo = v
}

// GetNetwork returns the Network field value
func (o *HoldCreateRequest) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *HoldCreateRequest) SetNetwork(v string) {
	o.Network = v
}

// GetOffsetDescription returns the OffsetDescription field value if set, zero value otherwise.
func (o *HoldCreateRequest) GetOffsetDescription() string {
	if o == nil || o.OffsetDescription == nil {
		var ret string
		return ret
	}
	return *o.OffsetDescription
}

// GetOffsetDescriptionOk returns a tuple with the OffsetDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetOffsetDescriptionOk() (*string, bool) {
	if o == nil || o.OffsetDescription == nil {
		return nil, false
	}
	return o.OffsetDescription, true
}

// HasOffsetDescription returns a boolean if a field has been set.
func (o *HoldCreateRequest) HasOffsetDescription() bool {
	if o != nil && o.OffsetDescription != nil {
		return true
	}

	return false
}

// SetOffsetDescription gets a reference to the given string and assigns it to the OffsetDescription field.
func (o *HoldCreateRequest) SetOffsetDescription(v string) {
	o.OffsetDescription = &v
}

// GetReferenceId returns the ReferenceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HoldCreateRequest) GetReferenceId() string {
	if o == nil || o.ReferenceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldCreateRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// SetReferenceId sets field value
func (o *HoldCreateRequest) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}

// GetStatus returns the Status field value
func (o *HoldCreateRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *HoldCreateRequest) SetStatus(v string) {
	o.Status = v
}

// GetSubtype returns the Subtype field value
func (o *HoldCreateRequest) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *HoldCreateRequest) SetSubtype(v string) {
	o.Subtype = v
}

// GetTransactionTime returns the TransactionTime field value
func (o *HoldCreateRequest) GetTransactionTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetTransactionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTime, true
}

// SetTransactionTime sets field value
func (o *HoldCreateRequest) SetTransactionTime(v time.Time) {
	o.TransactionTime = v
}

// GetType returns the Type field value
func (o *HoldCreateRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HoldCreateRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HoldCreateRequest) SetType(v string) {
	o.Type = v
}

// GetUserData returns the UserData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *HoldCreateRequest) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldCreateRequest) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// SetUserData sets field value
func (o *HoldCreateRequest) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

func (o HoldCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_no"] = o.AccountNo
	}
	if true {
		toSerialize["allow_partial"] = o.AllowPartial
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["auto_post_at"] = o.AutoPostAt
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["dc_sign"] = o.DcSign
	}
	if true {
		toSerialize["decline_reason"] = o.DeclineReason
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["effective_date"] = o.EffectiveDate
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
		toSerialize["force_post"] = o.ForcePost
	}
	if true {
		toSerialize["memo"] = o.Memo
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if o.OffsetDescription != nil {
		toSerialize["offset_description"] = o.OffsetDescription
	}
	if true {
		toSerialize["reference_id"] = o.ReferenceId.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["transaction_time"] = o.TransactionTime
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UserData != nil {
		toSerialize["user_data"] = o.UserData
	}
	return json.Marshal(toSerialize)
}

type NullableHoldCreateRequest struct {
	value *HoldCreateRequest
	isSet bool
}

func (v NullableHoldCreateRequest) Get() *HoldCreateRequest {
	return v.value
}

func (v *NullableHoldCreateRequest) Set(val *HoldCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldCreateRequest(val *HoldCreateRequest) *NullableHoldCreateRequest {
	return &NullableHoldCreateRequest{value: val, isSet: true}
}

func (v NullableHoldCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

