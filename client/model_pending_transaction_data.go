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
	"time"
)

// checks if the PendingTransactionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingTransactionData{}

// PendingTransactionData struct for PendingTransactionData
type PendingTransactionData struct {
	// The amount of the hold.
	Amount int64 `json:"amount"`
	// The account \"available balance\" at the time this hold was created
	AutoPostAt time.Time `json:"auto_post_at"`
	// The account \"available balance\" at the time this hold was created (to be deprecated)
	AvailBalance int64 `json:"avail_balance"`
	// The account \"available balance\" at the time this hold was created
	AvailableBalance int64 `json:"available_balance"`
	// The account balance at the time this hold was created
	Balance int64 `json:"balance"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	DcSign   DcSign `json:"dc_sign"`
	// The effective date of the transaction once it gets posted
	EffectiveDate time.Time `json:"effective_date"`
	// The date that at which this hold is no longer valid.
	ExpiresAt time.Time `json:"expires_at"`
	// an unstructured json blob representing additional transaction information supplied by the integrator.
	ExternalData map[string]interface{} `json:"external_data,omitempty"`
	// Whether or not the hold was forced (spending controls ignored)
	ForcePost bool `json:"force_post"`
	// An array representing any previous states of the hold, if it has been modified (For example, increasing or decreasing the hold amount).
	History []PendingTransactionHistory `json:"history"`
	// The idempotency key used when initially creating this hold.
	Idemkey string `json:"idemkey"`
	// A short note to the recipient
	Memo string `json:"memo"`
	// The network this transaction is associated with
	Network   string `json:"network"`
	Operation string `json:"operation"`
	// If a hold has been declined or modified, this will include the reason.
	Reason string `json:"reason"`
	// The requested amount, in the case of hold modifications.
	ReqAmount int64 `json:"req_amount"`
	// Information received by the transaction risk/fraud service related to this transaction
	RiskInfo map[string]interface{} `json:"risk_info,omitempty"`
	// The status of the hold.
	Status string `json:"status"`
	// The specific transaction type. For example, for `ach`, this may be \"outgoing_debit\".
	Subtype string `json:"subtype"`
	// The total amount of the hold. This may be different than `amount` in the case where a hold increase or decrease was requested.
	TotalAmount int64 `json:"total_amount"`
	// The uuid of the transaction that this pending transaction originated from, if any. This is primary used when a transaction \"posts\", but a subset of the amount reserved until a future settlement date.
	TransactionId *string `json:"transaction_id,omitempty"`
	// The time that the transaction was created
	TransactionTime time.Time `json:"transaction_time"`
	// The general type of transaction. For example, \"card\" or \"ach\".
	Type string `json:"type"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data,omitempty"`
	// Does this hold represent a partial debit (or credit)?
	WasPartial           bool `json:"was_partial"`
	AdditionalProperties map[string]interface{}
}

type _PendingTransactionData PendingTransactionData

// NewPendingTransactionData instantiates a new PendingTransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingTransactionData(amount int64, autoPostAt time.Time, availBalance int64, availableBalance int64, balance int64, currency string, dcSign DcSign, effectiveDate time.Time, expiresAt time.Time, forcePost bool, history []PendingTransactionHistory, idemkey string, memo string, network string, operation string, reason string, reqAmount int64, status string, subtype string, totalAmount int64, transactionTime time.Time, type_ string, wasPartial bool) *PendingTransactionData {
	this := PendingTransactionData{}
	this.Amount = amount
	this.AutoPostAt = autoPostAt
	this.AvailBalance = availBalance
	this.AvailableBalance = availableBalance
	this.Balance = balance
	this.Currency = currency
	this.DcSign = dcSign
	this.EffectiveDate = effectiveDate
	this.ExpiresAt = expiresAt
	this.ForcePost = forcePost
	this.History = history
	this.Idemkey = idemkey
	this.Memo = memo
	this.Network = network
	this.Operation = operation
	this.Reason = reason
	this.ReqAmount = reqAmount
	this.Status = status
	this.Subtype = subtype
	this.TotalAmount = totalAmount
	this.TransactionTime = transactionTime
	this.Type = type_
	this.WasPartial = wasPartial
	return &this
}

// NewPendingTransactionDataWithDefaults instantiates a new PendingTransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingTransactionDataWithDefaults() *PendingTransactionData {
	this := PendingTransactionData{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PendingTransactionData) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PendingTransactionData) SetAmount(v int64) {
	o.Amount = v
}

// GetAutoPostAt returns the AutoPostAt field value
func (o *PendingTransactionData) GetAutoPostAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AutoPostAt
}

// GetAutoPostAtOk returns a tuple with the AutoPostAt field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetAutoPostAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoPostAt, true
}

// SetAutoPostAt sets field value
func (o *PendingTransactionData) SetAutoPostAt(v time.Time) {
	o.AutoPostAt = v
}

// GetAvailBalance returns the AvailBalance field value
func (o *PendingTransactionData) GetAvailBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AvailBalance
}

// GetAvailBalanceOk returns a tuple with the AvailBalance field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetAvailBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailBalance, true
}

// SetAvailBalance sets field value
func (o *PendingTransactionData) SetAvailBalance(v int64) {
	o.AvailBalance = v
}

// GetAvailableBalance returns the AvailableBalance field value
func (o *PendingTransactionData) GetAvailableBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetAvailableBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableBalance, true
}

// SetAvailableBalance sets field value
func (o *PendingTransactionData) SetAvailableBalance(v int64) {
	o.AvailableBalance = v
}

// GetBalance returns the Balance field value
func (o *PendingTransactionData) GetBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *PendingTransactionData) SetBalance(v int64) {
	o.Balance = v
}

// GetCurrency returns the Currency field value
func (o *PendingTransactionData) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PendingTransactionData) SetCurrency(v string) {
	o.Currency = v
}

// GetDcSign returns the DcSign field value
func (o *PendingTransactionData) GetDcSign() DcSign {
	if o == nil {
		var ret DcSign
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetDcSignOk() (*DcSign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *PendingTransactionData) SetDcSign(v DcSign) {
	o.DcSign = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *PendingTransactionData) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *PendingTransactionData) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *PendingTransactionData) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *PendingTransactionData) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetExternalData returns the ExternalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingTransactionData) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingTransactionData) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExternalData) {
		return map[string]interface{}{}, false
	}
	return o.ExternalData, true
}

// HasExternalData returns a boolean if a field has been set.
func (o *PendingTransactionData) HasExternalData() bool {
	if o != nil && !IsNil(o.ExternalData) {
		return true
	}

	return false
}

// SetExternalData gets a reference to the given map[string]interface{} and assigns it to the ExternalData field.
func (o *PendingTransactionData) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetForcePost returns the ForcePost field value
func (o *PendingTransactionData) GetForcePost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetForcePostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForcePost, true
}

// SetForcePost sets field value
func (o *PendingTransactionData) SetForcePost(v bool) {
	o.ForcePost = v
}

// GetHistory returns the History field value
func (o *PendingTransactionData) GetHistory() []PendingTransactionHistory {
	if o == nil {
		var ret []PendingTransactionHistory
		return ret
	}

	return o.History
}

// GetHistoryOk returns a tuple with the History field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetHistoryOk() ([]PendingTransactionHistory, bool) {
	if o == nil {
		return nil, false
	}
	return o.History, true
}

// SetHistory sets field value
func (o *PendingTransactionData) SetHistory(v []PendingTransactionHistory) {
	o.History = v
}

// GetIdemkey returns the Idemkey field value
func (o *PendingTransactionData) GetIdemkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Idemkey
}

// GetIdemkeyOk returns a tuple with the Idemkey field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetIdemkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idemkey, true
}

// SetIdemkey sets field value
func (o *PendingTransactionData) SetIdemkey(v string) {
	o.Idemkey = v
}

// GetMemo returns the Memo field value
func (o *PendingTransactionData) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *PendingTransactionData) SetMemo(v string) {
	o.Memo = v
}

// GetNetwork returns the Network field value
func (o *PendingTransactionData) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *PendingTransactionData) SetNetwork(v string) {
	o.Network = v
}

// GetOperation returns the Operation field value
func (o *PendingTransactionData) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *PendingTransactionData) SetOperation(v string) {
	o.Operation = v
}

// GetReason returns the Reason field value
func (o *PendingTransactionData) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PendingTransactionData) SetReason(v string) {
	o.Reason = v
}

// GetReqAmount returns the ReqAmount field value
func (o *PendingTransactionData) GetReqAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReqAmount
}

// GetReqAmountOk returns a tuple with the ReqAmount field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetReqAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqAmount, true
}

// SetReqAmount sets field value
func (o *PendingTransactionData) SetReqAmount(v int64) {
	o.ReqAmount = v
}

// GetRiskInfo returns the RiskInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingTransactionData) GetRiskInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RiskInfo
}

// GetRiskInfoOk returns a tuple with the RiskInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingTransactionData) GetRiskInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RiskInfo) {
		return map[string]interface{}{}, false
	}
	return o.RiskInfo, true
}

// HasRiskInfo returns a boolean if a field has been set.
func (o *PendingTransactionData) HasRiskInfo() bool {
	if o != nil && !IsNil(o.RiskInfo) {
		return true
	}

	return false
}

// SetRiskInfo gets a reference to the given map[string]interface{} and assigns it to the RiskInfo field.
func (o *PendingTransactionData) SetRiskInfo(v map[string]interface{}) {
	o.RiskInfo = v
}

// GetStatus returns the Status field value
func (o *PendingTransactionData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PendingTransactionData) SetStatus(v string) {
	o.Status = v
}

// GetSubtype returns the Subtype field value
func (o *PendingTransactionData) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *PendingTransactionData) SetSubtype(v string) {
	o.Subtype = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *PendingTransactionData) GetTotalAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetTotalAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *PendingTransactionData) SetTotalAmount(v int64) {
	o.TotalAmount = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *PendingTransactionData) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *PendingTransactionData) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *PendingTransactionData) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionTime returns the TransactionTime field value
func (o *PendingTransactionData) GetTransactionTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetTransactionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTime, true
}

// SetTransactionTime sets field value
func (o *PendingTransactionData) SetTransactionTime(v time.Time) {
	o.TransactionTime = v
}

// GetType returns the Type field value
func (o *PendingTransactionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PendingTransactionData) SetType(v string) {
	o.Type = v
}

// GetUserData returns the UserData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingTransactionData) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingTransactionData) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserData) {
		return map[string]interface{}{}, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *PendingTransactionData) HasUserData() bool {
	if o != nil && !IsNil(o.UserData) {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]interface{} and assigns it to the UserData field.
func (o *PendingTransactionData) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

// GetWasPartial returns the WasPartial field value
func (o *PendingTransactionData) GetWasPartial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WasPartial
}

// GetWasPartialOk returns a tuple with the WasPartial field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetWasPartialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WasPartial, true
}

// SetWasPartial sets field value
func (o *PendingTransactionData) SetWasPartial(v bool) {
	o.WasPartial = v
}

func (o PendingTransactionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingTransactionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["auto_post_at"] = o.AutoPostAt
	toSerialize["avail_balance"] = o.AvailBalance
	toSerialize["available_balance"] = o.AvailableBalance
	toSerialize["balance"] = o.Balance
	toSerialize["currency"] = o.Currency
	toSerialize["dc_sign"] = o.DcSign
	toSerialize["effective_date"] = o.EffectiveDate
	toSerialize["expires_at"] = o.ExpiresAt
	if o.ExternalData != nil {
		toSerialize["external_data"] = o.ExternalData
	}
	toSerialize["force_post"] = o.ForcePost
	toSerialize["history"] = o.History
	toSerialize["idemkey"] = o.Idemkey
	toSerialize["memo"] = o.Memo
	toSerialize["network"] = o.Network
	toSerialize["operation"] = o.Operation
	toSerialize["reason"] = o.Reason
	toSerialize["req_amount"] = o.ReqAmount
	if o.RiskInfo != nil {
		toSerialize["risk_info"] = o.RiskInfo
	}
	toSerialize["status"] = o.Status
	toSerialize["subtype"] = o.Subtype
	toSerialize["total_amount"] = o.TotalAmount
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	toSerialize["transaction_time"] = o.TransactionTime
	toSerialize["type"] = o.Type
	if o.UserData != nil {
		toSerialize["user_data"] = o.UserData
	}
	toSerialize["was_partial"] = o.WasPartial

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PendingTransactionData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"auto_post_at",
		"avail_balance",
		"available_balance",
		"balance",
		"currency",
		"dc_sign",
		"effective_date",
		"expires_at",
		"force_post",
		"history",
		"idemkey",
		"memo",
		"network",
		"operation",
		"reason",
		"req_amount",
		"status",
		"subtype",
		"total_amount",
		"transaction_time",
		"type",
		"was_partial",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPendingTransactionData := _PendingTransactionData{}

	err = json.Unmarshal(data, &varPendingTransactionData)

	if err != nil {
		return err
	}

	*o = PendingTransactionData(varPendingTransactionData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "auto_post_at")
		delete(additionalProperties, "avail_balance")
		delete(additionalProperties, "available_balance")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "dc_sign")
		delete(additionalProperties, "effective_date")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "external_data")
		delete(additionalProperties, "force_post")
		delete(additionalProperties, "history")
		delete(additionalProperties, "idemkey")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "network")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "req_amount")
		delete(additionalProperties, "risk_info")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "total_amount")
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "transaction_time")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user_data")
		delete(additionalProperties, "was_partial")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePendingTransactionData struct {
	value *PendingTransactionData
	isSet bool
}

func (v NullablePendingTransactionData) Get() *PendingTransactionData {
	return v.value
}

func (v *NullablePendingTransactionData) Set(val *PendingTransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingTransactionData(val *PendingTransactionData) *NullablePendingTransactionData {
	return &NullablePendingTransactionData{value: val, isSet: true}
}

func (v NullablePendingTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
