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

// TransactionDirectPostRequest struct for TransactionDirectPostRequest
type TransactionDirectPostRequest struct {
	AccountAlias string `json:"account_alias"`
	// The account number associated with the transaction
	AccountNo string `json:"account_no"`
	// The amount of the hold.
	Amount int64 `json:"amount"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	DcSign DcSign `json:"dc_sign"`
	// The description of the transaction
	Description *string `json:"description,omitempty"`
	// The effective date of the transaction once it gets posted
	EffectiveDate time.Time `json:"effective_date"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	EnhancedTransaction map[string]interface{} `json:"enhanced_transaction,omitempty"`
	// an unstructured json blob representing additional transaction information supplied by the integrator.
	ExternalData map[string]interface{} `json:"external_data,omitempty"`
	// Whether or not the hold was forced (spending controls ignored)
	ForcePost bool `json:"force_post"`
	// Whether or not this transaction represents a purely informational operation or an actual money movement
	InfoOnly bool `json:"info_only"`
	// A short note to the recipient
	Memo string `json:"memo"`
	// The network this transaction is associated with
	Network string `json:"network"`
	// The description of the offset transaction
	OffsetDescription *string `json:"offset_description,omitempty"`
	// An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers.
	ReferenceId string `json:"reference_id"`
	// Information received by the transaction risk/fraud service related to this transaction
	RiskInfo map[string]interface{} `json:"risk_info,omitempty"`
	// The specific transaction type. For example, for `ach`, this may be \"outgoing_debit\".
	Subtype string `json:"subtype"`
	// The time the transaction occurred.
	TransactionTime time.Time `json:"transaction_time"`
	// The general type of transaction. For example, \"card\" or \"ach\".
	Type string `json:"type"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data,omitempty"`
	// The unique identifier of the transaction.
	Uuid *string `json:"uuid,omitempty"`
}

// NewTransactionDirectPostRequest instantiates a new TransactionDirectPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDirectPostRequest(accountAlias string, accountNo string, amount int64, currency string, dcSign DcSign, effectiveDate time.Time, forcePost bool, infoOnly bool, memo string, network string, referenceId string, subtype string, transactionTime time.Time, type_ string) *TransactionDirectPostRequest {
	this := TransactionDirectPostRequest{}
	this.AccountAlias = accountAlias
	this.AccountNo = accountNo
	this.Amount = amount
	this.Currency = currency
	this.DcSign = dcSign
	this.EffectiveDate = effectiveDate
	this.ForcePost = forcePost
	this.InfoOnly = infoOnly
	this.Memo = memo
	this.Network = network
	this.ReferenceId = referenceId
	this.Subtype = subtype
	this.TransactionTime = transactionTime
	this.Type = type_
	return &this
}

// NewTransactionDirectPostRequestWithDefaults instantiates a new TransactionDirectPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDirectPostRequestWithDefaults() *TransactionDirectPostRequest {
	this := TransactionDirectPostRequest{}
	return &this
}

// GetAccountAlias returns the AccountAlias field value
func (o *TransactionDirectPostRequest) GetAccountAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountAlias
}

// GetAccountAliasOk returns a tuple with the AccountAlias field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetAccountAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountAlias, true
}

// SetAccountAlias sets field value
func (o *TransactionDirectPostRequest) SetAccountAlias(v string) {
	o.AccountAlias = v
}

// GetAccountNo returns the AccountNo field value
func (o *TransactionDirectPostRequest) GetAccountNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetAccountNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNo, true
}

// SetAccountNo sets field value
func (o *TransactionDirectPostRequest) SetAccountNo(v string) {
	o.AccountNo = v
}

// GetAmount returns the Amount field value
func (o *TransactionDirectPostRequest) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionDirectPostRequest) SetAmount(v int64) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *TransactionDirectPostRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionDirectPostRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetDcSign returns the DcSign field value
func (o *TransactionDirectPostRequest) GetDcSign() DcSign {
	if o == nil {
		var ret DcSign
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetDcSignOk() (*DcSign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *TransactionDirectPostRequest) SetDcSign(v DcSign) {
	o.DcSign = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransactionDirectPostRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransactionDirectPostRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransactionDirectPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *TransactionDirectPostRequest) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *TransactionDirectPostRequest) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetEnhancedTransaction returns the EnhancedTransaction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionDirectPostRequest) GetEnhancedTransaction() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EnhancedTransaction
}

// GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionDirectPostRequest) GetEnhancedTransactionOk() (map[string]interface{}, bool) {
	if o == nil || o.EnhancedTransaction == nil {
		return nil, false
	}
	return o.EnhancedTransaction, true
}

// HasEnhancedTransaction returns a boolean if a field has been set.
func (o *TransactionDirectPostRequest) HasEnhancedTransaction() bool {
	if o != nil && o.EnhancedTransaction != nil {
		return true
	}

	return false
}

// SetEnhancedTransaction gets a reference to the given map[string]interface{} and assigns it to the EnhancedTransaction field.
func (o *TransactionDirectPostRequest) SetEnhancedTransaction(v map[string]interface{}) {
	o.EnhancedTransaction = v
}

// GetExternalData returns the ExternalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionDirectPostRequest) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionDirectPostRequest) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ExternalData == nil {
		return nil, false
	}
	return o.ExternalData, true
}

// HasExternalData returns a boolean if a field has been set.
func (o *TransactionDirectPostRequest) HasExternalData() bool {
	if o != nil && o.ExternalData != nil {
		return true
	}

	return false
}

// SetExternalData gets a reference to the given map[string]interface{} and assigns it to the ExternalData field.
func (o *TransactionDirectPostRequest) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetForcePost returns the ForcePost field value
func (o *TransactionDirectPostRequest) GetForcePost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetForcePostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForcePost, true
}

// SetForcePost sets field value
func (o *TransactionDirectPostRequest) SetForcePost(v bool) {
	o.ForcePost = v
}

// GetInfoOnly returns the InfoOnly field value
func (o *TransactionDirectPostRequest) GetInfoOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InfoOnly
}

// GetInfoOnlyOk returns a tuple with the InfoOnly field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetInfoOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfoOnly, true
}

// SetInfoOnly sets field value
func (o *TransactionDirectPostRequest) SetInfoOnly(v bool) {
	o.InfoOnly = v
}

// GetMemo returns the Memo field value
func (o *TransactionDirectPostRequest) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *TransactionDirectPostRequest) SetMemo(v string) {
	o.Memo = v
}

// GetNetwork returns the Network field value
func (o *TransactionDirectPostRequest) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TransactionDirectPostRequest) SetNetwork(v string) {
	o.Network = v
}

// GetOffsetDescription returns the OffsetDescription field value if set, zero value otherwise.
func (o *TransactionDirectPostRequest) GetOffsetDescription() string {
	if o == nil || o.OffsetDescription == nil {
		var ret string
		return ret
	}
	return *o.OffsetDescription
}

// GetOffsetDescriptionOk returns a tuple with the OffsetDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetOffsetDescriptionOk() (*string, bool) {
	if o == nil || o.OffsetDescription == nil {
		return nil, false
	}
	return o.OffsetDescription, true
}

// HasOffsetDescription returns a boolean if a field has been set.
func (o *TransactionDirectPostRequest) HasOffsetDescription() bool {
	if o != nil && o.OffsetDescription != nil {
		return true
	}

	return false
}

// SetOffsetDescription gets a reference to the given string and assigns it to the OffsetDescription field.
func (o *TransactionDirectPostRequest) SetOffsetDescription(v string) {
	o.OffsetDescription = &v
}

// GetReferenceId returns the ReferenceId field value
func (o *TransactionDirectPostRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *TransactionDirectPostRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetRiskInfo returns the RiskInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionDirectPostRequest) GetRiskInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RiskInfo
}

// GetRiskInfoOk returns a tuple with the RiskInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionDirectPostRequest) GetRiskInfoOk() (map[string]interface{}, bool) {
	if o == nil || o.RiskInfo == nil {
		return nil, false
	}
	return o.RiskInfo, true
}

// HasRiskInfo returns a boolean if a field has been set.
func (o *TransactionDirectPostRequest) HasRiskInfo() bool {
	if o != nil && o.RiskInfo != nil {
		return true
	}

	return false
}

// SetRiskInfo gets a reference to the given map[string]interface{} and assigns it to the RiskInfo field.
func (o *TransactionDirectPostRequest) SetRiskInfo(v map[string]interface{}) {
	o.RiskInfo = v
}

// GetSubtype returns the Subtype field value
func (o *TransactionDirectPostRequest) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *TransactionDirectPostRequest) SetSubtype(v string) {
	o.Subtype = v
}

// GetTransactionTime returns the TransactionTime field value
func (o *TransactionDirectPostRequest) GetTransactionTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetTransactionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTime, true
}

// SetTransactionTime sets field value
func (o *TransactionDirectPostRequest) SetTransactionTime(v time.Time) {
	o.TransactionTime = v
}

// GetType returns the Type field value
func (o *TransactionDirectPostRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionDirectPostRequest) SetType(v string) {
	o.Type = v
}

// GetUserData returns the UserData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionDirectPostRequest) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionDirectPostRequest) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *TransactionDirectPostRequest) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]interface{} and assigns it to the UserData field.
func (o *TransactionDirectPostRequest) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *TransactionDirectPostRequest) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDirectPostRequest) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *TransactionDirectPostRequest) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *TransactionDirectPostRequest) SetUuid(v string) {
	o.Uuid = &v
}

func (o TransactionDirectPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_alias"] = o.AccountAlias
	}
	if true {
		toSerialize["account_no"] = o.AccountNo
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["dc_sign"] = o.DcSign
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
	if o.ExternalData != nil {
		toSerialize["external_data"] = o.ExternalData
	}
	if true {
		toSerialize["force_post"] = o.ForcePost
	}
	if true {
		toSerialize["info_only"] = o.InfoOnly
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
		toSerialize["reference_id"] = o.ReferenceId
	}
	if o.RiskInfo != nil {
		toSerialize["risk_info"] = o.RiskInfo
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
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionDirectPostRequest struct {
	value *TransactionDirectPostRequest
	isSet bool
}

func (v NullableTransactionDirectPostRequest) Get() *TransactionDirectPostRequest {
	return v.value
}

func (v *NullableTransactionDirectPostRequest) Set(val *TransactionDirectPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDirectPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDirectPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDirectPostRequest(val *TransactionDirectPostRequest) *NullableTransactionDirectPostRequest {
	return &NullableTransactionDirectPostRequest{value: val, isSet: true}
}

func (v NullableTransactionDirectPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDirectPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


