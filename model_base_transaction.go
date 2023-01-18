/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// BaseTransaction struct for BaseTransaction
type BaseTransaction struct {
	// The account uuid associated with the transaction. `account_id` and `internal_account_id` are mutually exclusive
	AccountId *string `json:"account_id,omitempty"`
	// The total amount of the transaction including both pending and already posted amounts. The value is represented as the smallest denomination of the applicable currency.
	Amount int64 `json:"amount"`
	// The exact time the transaction was recorded in the ledger
	CreationTime time.Time `json:"creation_time"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	// The uuid of the customer that initiated the transaction (if any) 
	CustomerId *string `json:"customer_id,omitempty"`
	// The `dc_sign` represents the direction money was moved. A value of `DEBIT` is money moving out of an account, a value of `CREDIT` is money moving into an account
	DcSign string `json:"dc_sign"`
	Decline *BaseTransactionDecline `json:"decline,omitempty"`
	// A human-friendly description of the transaction, provided by the integrator
	Description *string `json:"description,omitempty"`
	// The effective date of the transaction. This usually aligns with network settlement date, which differs between transaction types. The effective date is also used to determine effective daily balances for the purposes of interest calculation.
	EffectiveDate string `json:"effective_date"`
	EnhancedTransaction *EnhancedTransactionData `json:"enhanced_transaction,omitempty"`
	// Determines whether or not a transaction or auth was \"forced\" or not. A forced transaction skips any account balance checks
	ForcePost bool `json:"force_post"`
	// The group id of the transaction. Every transaction in the ledger is one entry in a double-entry system and the primary and offset transactions share the same `group_id`
	GroupId string `json:"group_id"`
	// The date and time any pending amount is expected to be released back to the account.
	HoldExpirationTime *time.Time `json:"hold_expiration_time,omitempty"`
	// The unique identifier of the transaction
	Id string `json:"id"`
	// The internal account uuid associated with the transaction. `account_id` and `internal_account_id` are mutually exclusive
	InternalAccountId *string `json:"internal_account_id,omitempty"`
	// The date and time the transaction was last modified
	LastUpdatedTime time.Time `json:"last_updated_time"`
	// Determines whether or not the funds on hold were the result of a partial auth or not. If `true` the `pending_amount` of the transaction will be less than the requested amount. This is primarily used for certain types of card transactions.
	PartialHold bool `json:"partial_hold"`
	// The amount amount of the transaction currently authorized or on hold
	PendingAmount int64 `json:"pending_amount"`
	// The amount of the transaction that has been fully posted to the account
	PostedAmount int64 `json:"posted_amount"`
	// The date the transaction was posted (based on the bank calendar and end-of-day). For transaction with multiple postings, this is the date of the earliest posting. This will be omitted for transactions with a `posted_amount` of `0`.
	PostedDate *string `json:"posted_date,omitempty"`
	// An external ID provided by the payment network to represent this transaction. This is not guaranteed to be globally unique. This will always be omitted for internal transfers.
	ReferenceId *string `json:"reference_id,omitempty"`
	// The status of the transaction
	Status string `json:"status"`
	// A human-friendly description of the transaction, provided by the Synctera platform
	SystemDescription *string `json:"system_description,omitempty"`
	// The time the transaction occurred. In most cases this will be roughly identical to creation_time, but it can differ in some situations if the payment doesn't appear in the ledger in real-time.
	TransactionTime time.Time `json:"transaction_time"`
	// The type of the transaction. This typically represents the \"payment\" rail that is used. For example, for ACH payments this will be `ach`, while debit card transactions will use `card`.
	Type string `json:"type"`
}

// NewBaseTransaction instantiates a new BaseTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTransaction(amount int64, creationTime time.Time, currency string, dcSign string, effectiveDate string, forcePost bool, groupId string, id string, lastUpdatedTime time.Time, partialHold bool, pendingAmount int64, postedAmount int64, status string, transactionTime time.Time, type_ string) *BaseTransaction {
	this := BaseTransaction{}
	this.Amount = amount
	this.CreationTime = creationTime
	this.Currency = currency
	this.DcSign = dcSign
	this.EffectiveDate = effectiveDate
	this.ForcePost = forcePost
	this.GroupId = groupId
	this.Id = id
	this.LastUpdatedTime = lastUpdatedTime
	this.PartialHold = partialHold
	this.PendingAmount = pendingAmount
	this.PostedAmount = postedAmount
	this.Status = status
	this.TransactionTime = transactionTime
	this.Type = type_
	return &this
}

// NewBaseTransactionWithDefaults instantiates a new BaseTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTransactionWithDefaults() *BaseTransaction {
	this := BaseTransaction{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *BaseTransaction) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *BaseTransaction) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *BaseTransaction) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAmount returns the Amount field value
func (o *BaseTransaction) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BaseTransaction) SetAmount(v int64) {
	o.Amount = v
}

// GetCreationTime returns the CreationTime field value
func (o *BaseTransaction) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *BaseTransaction) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCurrency returns the Currency field value
func (o *BaseTransaction) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *BaseTransaction) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *BaseTransaction) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *BaseTransaction) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *BaseTransaction) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDcSign returns the DcSign field value
func (o *BaseTransaction) GetDcSign() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetDcSignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *BaseTransaction) SetDcSign(v string) {
	o.DcSign = v
}

// GetDecline returns the Decline field value if set, zero value otherwise.
func (o *BaseTransaction) GetDecline() BaseTransactionDecline {
	if o == nil || o.Decline == nil {
		var ret BaseTransactionDecline
		return ret
	}
	return *o.Decline
}

// GetDeclineOk returns a tuple with the Decline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetDeclineOk() (*BaseTransactionDecline, bool) {
	if o == nil || o.Decline == nil {
		return nil, false
	}
	return o.Decline, true
}

// HasDecline returns a boolean if a field has been set.
func (o *BaseTransaction) HasDecline() bool {
	if o != nil && o.Decline != nil {
		return true
	}

	return false
}

// SetDecline gets a reference to the given BaseTransactionDecline and assigns it to the Decline field.
func (o *BaseTransaction) SetDecline(v BaseTransactionDecline) {
	o.Decline = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseTransaction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseTransaction) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseTransaction) SetDescription(v string) {
	o.Description = &v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *BaseTransaction) GetEffectiveDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetEffectiveDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *BaseTransaction) SetEffectiveDate(v string) {
	o.EffectiveDate = v
}

// GetEnhancedTransaction returns the EnhancedTransaction field value if set, zero value otherwise.
func (o *BaseTransaction) GetEnhancedTransaction() EnhancedTransactionData {
	if o == nil || o.EnhancedTransaction == nil {
		var ret EnhancedTransactionData
		return ret
	}
	return *o.EnhancedTransaction
}

// GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetEnhancedTransactionOk() (*EnhancedTransactionData, bool) {
	if o == nil || o.EnhancedTransaction == nil {
		return nil, false
	}
	return o.EnhancedTransaction, true
}

// HasEnhancedTransaction returns a boolean if a field has been set.
func (o *BaseTransaction) HasEnhancedTransaction() bool {
	if o != nil && o.EnhancedTransaction != nil {
		return true
	}

	return false
}

// SetEnhancedTransaction gets a reference to the given EnhancedTransactionData and assigns it to the EnhancedTransaction field.
func (o *BaseTransaction) SetEnhancedTransaction(v EnhancedTransactionData) {
	o.EnhancedTransaction = &v
}

// GetForcePost returns the ForcePost field value
func (o *BaseTransaction) GetForcePost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetForcePostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForcePost, true
}

// SetForcePost sets field value
func (o *BaseTransaction) SetForcePost(v bool) {
	o.ForcePost = v
}

// GetGroupId returns the GroupId field value
func (o *BaseTransaction) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *BaseTransaction) SetGroupId(v string) {
	o.GroupId = v
}

// GetHoldExpirationTime returns the HoldExpirationTime field value if set, zero value otherwise.
func (o *BaseTransaction) GetHoldExpirationTime() time.Time {
	if o == nil || o.HoldExpirationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.HoldExpirationTime
}

// GetHoldExpirationTimeOk returns a tuple with the HoldExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetHoldExpirationTimeOk() (*time.Time, bool) {
	if o == nil || o.HoldExpirationTime == nil {
		return nil, false
	}
	return o.HoldExpirationTime, true
}

// HasHoldExpirationTime returns a boolean if a field has been set.
func (o *BaseTransaction) HasHoldExpirationTime() bool {
	if o != nil && o.HoldExpirationTime != nil {
		return true
	}

	return false
}

// SetHoldExpirationTime gets a reference to the given time.Time and assigns it to the HoldExpirationTime field.
func (o *BaseTransaction) SetHoldExpirationTime(v time.Time) {
	o.HoldExpirationTime = &v
}

// GetId returns the Id field value
func (o *BaseTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BaseTransaction) SetId(v string) {
	o.Id = v
}

// GetInternalAccountId returns the InternalAccountId field value if set, zero value otherwise.
func (o *BaseTransaction) GetInternalAccountId() string {
	if o == nil || o.InternalAccountId == nil {
		var ret string
		return ret
	}
	return *o.InternalAccountId
}

// GetInternalAccountIdOk returns a tuple with the InternalAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetInternalAccountIdOk() (*string, bool) {
	if o == nil || o.InternalAccountId == nil {
		return nil, false
	}
	return o.InternalAccountId, true
}

// HasInternalAccountId returns a boolean if a field has been set.
func (o *BaseTransaction) HasInternalAccountId() bool {
	if o != nil && o.InternalAccountId != nil {
		return true
	}

	return false
}

// SetInternalAccountId gets a reference to the given string and assigns it to the InternalAccountId field.
func (o *BaseTransaction) SetInternalAccountId(v string) {
	o.InternalAccountId = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *BaseTransaction) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *BaseTransaction) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetPartialHold returns the PartialHold field value
func (o *BaseTransaction) GetPartialHold() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PartialHold
}

// GetPartialHoldOk returns a tuple with the PartialHold field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetPartialHoldOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartialHold, true
}

// SetPartialHold sets field value
func (o *BaseTransaction) SetPartialHold(v bool) {
	o.PartialHold = v
}

// GetPendingAmount returns the PendingAmount field value
func (o *BaseTransaction) GetPendingAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PendingAmount
}

// GetPendingAmountOk returns a tuple with the PendingAmount field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetPendingAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingAmount, true
}

// SetPendingAmount sets field value
func (o *BaseTransaction) SetPendingAmount(v int64) {
	o.PendingAmount = v
}

// GetPostedAmount returns the PostedAmount field value
func (o *BaseTransaction) GetPostedAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PostedAmount
}

// GetPostedAmountOk returns a tuple with the PostedAmount field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetPostedAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostedAmount, true
}

// SetPostedAmount sets field value
func (o *BaseTransaction) SetPostedAmount(v int64) {
	o.PostedAmount = v
}

// GetPostedDate returns the PostedDate field value if set, zero value otherwise.
func (o *BaseTransaction) GetPostedDate() string {
	if o == nil || o.PostedDate == nil {
		var ret string
		return ret
	}
	return *o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetPostedDateOk() (*string, bool) {
	if o == nil || o.PostedDate == nil {
		return nil, false
	}
	return o.PostedDate, true
}

// HasPostedDate returns a boolean if a field has been set.
func (o *BaseTransaction) HasPostedDate() bool {
	if o != nil && o.PostedDate != nil {
		return true
	}

	return false
}

// SetPostedDate gets a reference to the given string and assigns it to the PostedDate field.
func (o *BaseTransaction) SetPostedDate(v string) {
	o.PostedDate = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *BaseTransaction) GetReferenceId() string {
	if o == nil || o.ReferenceId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetReferenceIdOk() (*string, bool) {
	if o == nil || o.ReferenceId == nil {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *BaseTransaction) HasReferenceId() bool {
	if o != nil && o.ReferenceId != nil {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *BaseTransaction) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetStatus returns the Status field value
func (o *BaseTransaction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BaseTransaction) SetStatus(v string) {
	o.Status = v
}

// GetSystemDescription returns the SystemDescription field value if set, zero value otherwise.
func (o *BaseTransaction) GetSystemDescription() string {
	if o == nil || o.SystemDescription == nil {
		var ret string
		return ret
	}
	return *o.SystemDescription
}

// GetSystemDescriptionOk returns a tuple with the SystemDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetSystemDescriptionOk() (*string, bool) {
	if o == nil || o.SystemDescription == nil {
		return nil, false
	}
	return o.SystemDescription, true
}

// HasSystemDescription returns a boolean if a field has been set.
func (o *BaseTransaction) HasSystemDescription() bool {
	if o != nil && o.SystemDescription != nil {
		return true
	}

	return false
}

// SetSystemDescription gets a reference to the given string and assigns it to the SystemDescription field.
func (o *BaseTransaction) SetSystemDescription(v string) {
	o.SystemDescription = &v
}

// GetTransactionTime returns the TransactionTime field value
func (o *BaseTransaction) GetTransactionTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetTransactionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTime, true
}

// SetTransactionTime sets field value
func (o *BaseTransaction) SetTransactionTime(v time.Time) {
	o.TransactionTime = v
}

// GetType returns the Type field value
func (o *BaseTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BaseTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BaseTransaction) SetType(v string) {
	o.Type = v
}

func (o BaseTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["dc_sign"] = o.DcSign
	}
	if o.Decline != nil {
		toSerialize["decline"] = o.Decline
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
		toSerialize["force_post"] = o.ForcePost
	}
	if true {
		toSerialize["group_id"] = o.GroupId
	}
	if o.HoldExpirationTime != nil {
		toSerialize["hold_expiration_time"] = o.HoldExpirationTime
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.InternalAccountId != nil {
		toSerialize["internal_account_id"] = o.InternalAccountId
	}
	if true {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if true {
		toSerialize["partial_hold"] = o.PartialHold
	}
	if true {
		toSerialize["pending_amount"] = o.PendingAmount
	}
	if true {
		toSerialize["posted_amount"] = o.PostedAmount
	}
	if o.PostedDate != nil {
		toSerialize["posted_date"] = o.PostedDate
	}
	if o.ReferenceId != nil {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.SystemDescription != nil {
		toSerialize["system_description"] = o.SystemDescription
	}
	if true {
		toSerialize["transaction_time"] = o.TransactionTime
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBaseTransaction struct {
	value *BaseTransaction
	isSet bool
}

func (v NullableBaseTransaction) Get() *BaseTransaction {
	return v.value
}

func (v *NullableBaseTransaction) Set(val *BaseTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTransaction(val *BaseTransaction) *NullableBaseTransaction {
	return &NullableBaseTransaction{value: val, isSet: true}
}

func (v NullableBaseTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

