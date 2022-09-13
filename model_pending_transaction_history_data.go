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

// PendingTransactionHistoryData struct for PendingTransactionHistoryData
type PendingTransactionHistoryData struct {
	// The amount of the hold.
	Amount int64 `json:"amount"`
	// The time the transaction will be automatically posted.
	AutoPostAt time.Time `json:"auto_post_at"`
	// The account \"available balance\" at the time this hold was created (to be deprecated)
	AvailBalance int64 `json:"avail_balance"`
	// The account \"available balance\" at the time this hold was created
	AvailableBalance int64 `json:"available_balance"`
	// The account balance at the time this hold was created
	Balance int64 `json:"balance"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	DcSign DcSign `json:"dc_sign"`
	// The effective date of the transaction once it gets posted
	EffectiveDate time.Time `json:"effective_date"`
	// The date that at which this hold is no longer valid.
	ExpiresAt time.Time `json:"expires_at"`
	// an unstructured json blob representing additional transaction information supplied by the integrator.
	ExternalData map[string]interface{} `json:"external_data"`
	// Whether or not the hold was forced (spending controls ignored)
	ForcePost bool `json:"force_post"`
	// The idempotency key used when initially creating this hold.
	Idemkey string `json:"idemkey"`
	// A short note to the recipient
	Memo string `json:"memo"`
	// The network this transaction is associated with
	Network string `json:"network"`
	Operation string `json:"operation"`
	// If a hold has been declined or modified, this will include the reason.
	Reason string `json:"reason"`
	// The requested amount, in the case of hold modifications.
	ReqAmount int64 `json:"req_amount"`
	// Information received by the transaction risk/fraud service related to this transaction
	RiskInfo map[string]interface{} `json:"risk_info"`
	// The status of the hold.
	Status string `json:"status"`
	// The specific transaction type. For example, for `ach`, this may be \"outgoing_debit\".
	Subtype string `json:"subtype"`
	// The total amount of the hold. This may be different than `amount` in the case where a hold increase or decrease was requested.
	TotalAmount int64 `json:"total_amount"`
	// The uuid of the transaction that this pending transaction originated from, if any. This is primary used when a transaction \"posts\", but a subset of the amount reserved until a future settlement date.
	TransactionId *string `json:"transaction_id,omitempty"`
	// The time the transaction occurred.
	TransactionTime time.Time `json:"transaction_time"`
	// The general type of transaction. For example, \"card\" or \"ach\".
	Type string `json:"type"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data"`
	// Does this hold represent a partial debit (or credit)?
	WasPartial bool `json:"was_partial"`
}

// NewPendingTransactionHistoryData instantiates a new PendingTransactionHistoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingTransactionHistoryData(amount int64, autoPostAt time.Time, availBalance int64, availableBalance int64, balance int64, currency string, dcSign DcSign, effectiveDate time.Time, expiresAt time.Time, externalData map[string]interface{}, forcePost bool, idemkey string, memo string, network string, operation string, reason string, reqAmount int64, riskInfo map[string]interface{}, status string, subtype string, totalAmount int64, transactionTime time.Time, type_ string, userData map[string]interface{}, wasPartial bool) *PendingTransactionHistoryData {
	this := PendingTransactionHistoryData{}
	this.Amount = amount
	this.AutoPostAt = autoPostAt
	this.AvailBalance = availBalance
	this.AvailableBalance = availableBalance
	this.Balance = balance
	this.Currency = currency
	this.DcSign = dcSign
	this.EffectiveDate = effectiveDate
	this.ExpiresAt = expiresAt
	this.ExternalData = externalData
	this.ForcePost = forcePost
	this.Idemkey = idemkey
	this.Memo = memo
	this.Network = network
	this.Operation = operation
	this.Reason = reason
	this.ReqAmount = reqAmount
	this.RiskInfo = riskInfo
	this.Status = status
	this.Subtype = subtype
	this.TotalAmount = totalAmount
	this.TransactionTime = transactionTime
	this.Type = type_
	this.UserData = userData
	this.WasPartial = wasPartial
	return &this
}

// NewPendingTransactionHistoryDataWithDefaults instantiates a new PendingTransactionHistoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingTransactionHistoryDataWithDefaults() *PendingTransactionHistoryData {
	this := PendingTransactionHistoryData{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PendingTransactionHistoryData) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PendingTransactionHistoryData) SetAmount(v int64) {
	o.Amount = v
}

// GetAutoPostAt returns the AutoPostAt field value
func (o *PendingTransactionHistoryData) GetAutoPostAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AutoPostAt
}

// GetAutoPostAtOk returns a tuple with the AutoPostAt field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetAutoPostAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoPostAt, true
}

// SetAutoPostAt sets field value
func (o *PendingTransactionHistoryData) SetAutoPostAt(v time.Time) {
	o.AutoPostAt = v
}

// GetAvailBalance returns the AvailBalance field value
func (o *PendingTransactionHistoryData) GetAvailBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AvailBalance
}

// GetAvailBalanceOk returns a tuple with the AvailBalance field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetAvailBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailBalance, true
}

// SetAvailBalance sets field value
func (o *PendingTransactionHistoryData) SetAvailBalance(v int64) {
	o.AvailBalance = v
}

// GetAvailableBalance returns the AvailableBalance field value
func (o *PendingTransactionHistoryData) GetAvailableBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetAvailableBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableBalance, true
}

// SetAvailableBalance sets field value
func (o *PendingTransactionHistoryData) SetAvailableBalance(v int64) {
	o.AvailableBalance = v
}

// GetBalance returns the Balance field value
func (o *PendingTransactionHistoryData) GetBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *PendingTransactionHistoryData) SetBalance(v int64) {
	o.Balance = v
}

// GetCurrency returns the Currency field value
func (o *PendingTransactionHistoryData) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PendingTransactionHistoryData) SetCurrency(v string) {
	o.Currency = v
}

// GetDcSign returns the DcSign field value
func (o *PendingTransactionHistoryData) GetDcSign() DcSign {
	if o == nil {
		var ret DcSign
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetDcSignOk() (*DcSign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *PendingTransactionHistoryData) SetDcSign(v DcSign) {
	o.DcSign = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *PendingTransactionHistoryData) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *PendingTransactionHistoryData) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *PendingTransactionHistoryData) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *PendingTransactionHistoryData) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetExternalData returns the ExternalData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PendingTransactionHistoryData) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingTransactionHistoryData) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ExternalData == nil {
		return nil, false
	}
	return o.ExternalData, true
}

// SetExternalData sets field value
func (o *PendingTransactionHistoryData) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetForcePost returns the ForcePost field value
func (o *PendingTransactionHistoryData) GetForcePost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetForcePostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForcePost, true
}

// SetForcePost sets field value
func (o *PendingTransactionHistoryData) SetForcePost(v bool) {
	o.ForcePost = v
}

// GetIdemkey returns the Idemkey field value
func (o *PendingTransactionHistoryData) GetIdemkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Idemkey
}

// GetIdemkeyOk returns a tuple with the Idemkey field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetIdemkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idemkey, true
}

// SetIdemkey sets field value
func (o *PendingTransactionHistoryData) SetIdemkey(v string) {
	o.Idemkey = v
}

// GetMemo returns the Memo field value
func (o *PendingTransactionHistoryData) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *PendingTransactionHistoryData) SetMemo(v string) {
	o.Memo = v
}

// GetNetwork returns the Network field value
func (o *PendingTransactionHistoryData) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *PendingTransactionHistoryData) SetNetwork(v string) {
	o.Network = v
}

// GetOperation returns the Operation field value
func (o *PendingTransactionHistoryData) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *PendingTransactionHistoryData) SetOperation(v string) {
	o.Operation = v
}

// GetReason returns the Reason field value
func (o *PendingTransactionHistoryData) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PendingTransactionHistoryData) SetReason(v string) {
	o.Reason = v
}

// GetReqAmount returns the ReqAmount field value
func (o *PendingTransactionHistoryData) GetReqAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReqAmount
}

// GetReqAmountOk returns a tuple with the ReqAmount field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetReqAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqAmount, true
}

// SetReqAmount sets field value
func (o *PendingTransactionHistoryData) SetReqAmount(v int64) {
	o.ReqAmount = v
}

// GetRiskInfo returns the RiskInfo field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PendingTransactionHistoryData) GetRiskInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.RiskInfo
}

// GetRiskInfoOk returns a tuple with the RiskInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingTransactionHistoryData) GetRiskInfoOk() (map[string]interface{}, bool) {
	if o == nil || o.RiskInfo == nil {
		return nil, false
	}
	return o.RiskInfo, true
}

// SetRiskInfo sets field value
func (o *PendingTransactionHistoryData) SetRiskInfo(v map[string]interface{}) {
	o.RiskInfo = v
}

// GetStatus returns the Status field value
func (o *PendingTransactionHistoryData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PendingTransactionHistoryData) SetStatus(v string) {
	o.Status = v
}

// GetSubtype returns the Subtype field value
func (o *PendingTransactionHistoryData) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *PendingTransactionHistoryData) SetSubtype(v string) {
	o.Subtype = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *PendingTransactionHistoryData) GetTotalAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetTotalAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *PendingTransactionHistoryData) SetTotalAmount(v int64) {
	o.TotalAmount = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *PendingTransactionHistoryData) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *PendingTransactionHistoryData) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *PendingTransactionHistoryData) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionTime returns the TransactionTime field value
func (o *PendingTransactionHistoryData) GetTransactionTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetTransactionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTime, true
}

// SetTransactionTime sets field value
func (o *PendingTransactionHistoryData) SetTransactionTime(v time.Time) {
	o.TransactionTime = v
}

// GetType returns the Type field value
func (o *PendingTransactionHistoryData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PendingTransactionHistoryData) SetType(v string) {
	o.Type = v
}

// GetUserData returns the UserData field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PendingTransactionHistoryData) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingTransactionHistoryData) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// SetUserData sets field value
func (o *PendingTransactionHistoryData) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

// GetWasPartial returns the WasPartial field value
func (o *PendingTransactionHistoryData) GetWasPartial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WasPartial
}

// GetWasPartialOk returns a tuple with the WasPartial field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistoryData) GetWasPartialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WasPartial, true
}

// SetWasPartial sets field value
func (o *PendingTransactionHistoryData) SetWasPartial(v bool) {
	o.WasPartial = v
}

func (o PendingTransactionHistoryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["auto_post_at"] = o.AutoPostAt
	}
	if true {
		toSerialize["avail_balance"] = o.AvailBalance
	}
	if true {
		toSerialize["available_balance"] = o.AvailableBalance
	}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["dc_sign"] = o.DcSign
	}
	if true {
		toSerialize["effective_date"] = o.EffectiveDate
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
		toSerialize["idemkey"] = o.Idemkey
	}
	if true {
		toSerialize["memo"] = o.Memo
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["req_amount"] = o.ReqAmount
	}
	if o.RiskInfo != nil {
		toSerialize["risk_info"] = o.RiskInfo
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["total_amount"] = o.TotalAmount
	}
	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
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
	if true {
		toSerialize["was_partial"] = o.WasPartial
	}
	return json.Marshal(toSerialize)
}

type NullablePendingTransactionHistoryData struct {
	value *PendingTransactionHistoryData
	isSet bool
}

func (v NullablePendingTransactionHistoryData) Get() *PendingTransactionHistoryData {
	return v.value
}

func (v *NullablePendingTransactionHistoryData) Set(val *PendingTransactionHistoryData) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingTransactionHistoryData) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingTransactionHistoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingTransactionHistoryData(val *PendingTransactionHistoryData) *NullablePendingTransactionHistoryData {
	return &NullablePendingTransactionHistoryData{value: val, isSet: true}
}

func (v NullablePendingTransactionHistoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingTransactionHistoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


