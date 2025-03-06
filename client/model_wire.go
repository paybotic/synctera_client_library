/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Wire type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Wire{}

// Wire struct for Wire
type Wire struct {
	// Transfer amount in cents ($100 would be 10000)
	Amount int32 `json:"amount"`
	// Instructions intended for the financial institutions that are processing the wire.
	BankMessage *string `json:"bank_message,omitempty"`
	// The batch ID associated with the wire if it was created via the batch payment API.
	BatchId *string `json:"batch_id,omitempty"`
	// The case id associated with the wire.
	CaseId       *int32    `json:"case_id,omitempty"`
	CreationTime time.Time `json:"creation_time"`
	// 3-character currency code
	Currency string `json:"currency"`
	// The customer UUID representing the person initiating the Wire transfer
	CustomerId *string `json:"customer_id,omitempty"`
	// The effective date of the transaction once it gets posted
	EffectiveDate string `json:"effective_date"`
	// wire ID
	Id string `json:"id"`
	// The input message accountability data consists of a 8 character cycle date (CCYYMMDD) an 8 character source and a 6 character sequence number.
	InputMessageAccountabilityData *string `json:"input_message_accountability_data,omitempty"`
	// Whether or not the wire is a \"bulk\" wire created via the batch payment API.
	IsBulk          bool      `json:"is_bulk"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
	// The network used to process the wire
	Network *string `json:"network,omitempty"`
	// Sender account ID
	OriginatingAccountId *string `json:"originating_account_id,omitempty"`
	// The account number representing the sender account. If the outgoing wire is a return, it refers to the sender of the initial wire not the sender of the return.
	OriginatingAccountNumber string `json:"originating_account_number"`
	// The external account uuid representing the recipient of the wire.
	ReceivingAccountId *string `json:"receiving_account_id,omitempty"`
	// The account number representing the recipient account. If the outgoing wire is a return, it refers to the recipient of the initial wire not the destination of the return.
	ReceivingAccountNumber string `json:"receiving_account_number"`
	// Information from the originator to the beneficiary (recipient).
	RecipientMessage *string      `json:"recipient_message,omitempty"`
	ReturnData       *ReturnData1 `json:"return_data,omitempty"`
	// Sender's id associated with fedwire transfer
	SenderReferenceId string `json:"sender_reference_id"`
	// The settlement date of the transaction once it gets posted
	SettlementDate *string `json:"settlement_date,omitempty"`
	// The current status of the transfer
	Status string `json:"status"`
	// Additional details about the status of the transfer
	StatusDetails *string `json:"status_details,omitempty"`
	// ID of the resulting transaction resource
	TransactionId string `json:"transaction_id"`
	// The transaction uuid of the incoming wire that triggered an outgoing return. This is only used if the outgoing wire is a return.
	TransactionInId *string `json:"transaction_in_id,omitempty"`
}

type _Wire Wire

// NewWire instantiates a new Wire object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWire(amount int32, creationTime time.Time, currency string, effectiveDate string, id string, isBulk bool, lastUpdatedTime time.Time, originatingAccountNumber string, receivingAccountNumber string, senderReferenceId string, status string, transactionId string) *Wire {
	this := Wire{}
	this.Amount = amount
	this.CreationTime = creationTime
	this.Currency = currency
	this.EffectiveDate = effectiveDate
	this.Id = id
	this.IsBulk = isBulk
	this.LastUpdatedTime = lastUpdatedTime
	this.OriginatingAccountNumber = originatingAccountNumber
	this.ReceivingAccountNumber = receivingAccountNumber
	this.SenderReferenceId = senderReferenceId
	this.Status = status
	this.TransactionId = transactionId
	return &this
}

// NewWireWithDefaults instantiates a new Wire object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireWithDefaults() *Wire {
	this := Wire{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Wire) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Wire) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Wire) SetAmount(v int32) {
	o.Amount = v
}

// GetBankMessage returns the BankMessage field value if set, zero value otherwise.
func (o *Wire) GetBankMessage() string {
	if o == nil || IsNil(o.BankMessage) {
		var ret string
		return ret
	}
	return *o.BankMessage
}

// GetBankMessageOk returns a tuple with the BankMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetBankMessageOk() (*string, bool) {
	if o == nil || IsNil(o.BankMessage) {
		return nil, false
	}
	return o.BankMessage, true
}

// HasBankMessage returns a boolean if a field has been set.
func (o *Wire) HasBankMessage() bool {
	if o != nil && !IsNil(o.BankMessage) {
		return true
	}

	return false
}

// SetBankMessage gets a reference to the given string and assigns it to the BankMessage field.
func (o *Wire) SetBankMessage(v string) {
	o.BankMessage = &v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *Wire) GetBatchId() string {
	if o == nil || IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *Wire) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *Wire) SetBatchId(v string) {
	o.BatchId = &v
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *Wire) GetCaseId() int32 {
	if o == nil || IsNil(o.CaseId) {
		var ret int32
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetCaseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *Wire) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given int32 and assigns it to the CaseId field.
func (o *Wire) SetCaseId(v int32) {
	o.CaseId = &v
}

// GetCreationTime returns the CreationTime field value
func (o *Wire) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *Wire) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *Wire) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCurrency returns the Currency field value
func (o *Wire) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Wire) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Wire) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *Wire) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Wire) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *Wire) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *Wire) GetEffectiveDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *Wire) GetEffectiveDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *Wire) SetEffectiveDate(v string) {
	o.EffectiveDate = v
}

// GetId returns the Id field value
func (o *Wire) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Wire) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Wire) SetId(v string) {
	o.Id = v
}

// GetInputMessageAccountabilityData returns the InputMessageAccountabilityData field value if set, zero value otherwise.
func (o *Wire) GetInputMessageAccountabilityData() string {
	if o == nil || IsNil(o.InputMessageAccountabilityData) {
		var ret string
		return ret
	}
	return *o.InputMessageAccountabilityData
}

// GetInputMessageAccountabilityDataOk returns a tuple with the InputMessageAccountabilityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetInputMessageAccountabilityDataOk() (*string, bool) {
	if o == nil || IsNil(o.InputMessageAccountabilityData) {
		return nil, false
	}
	return o.InputMessageAccountabilityData, true
}

// HasInputMessageAccountabilityData returns a boolean if a field has been set.
func (o *Wire) HasInputMessageAccountabilityData() bool {
	if o != nil && !IsNil(o.InputMessageAccountabilityData) {
		return true
	}

	return false
}

// SetInputMessageAccountabilityData gets a reference to the given string and assigns it to the InputMessageAccountabilityData field.
func (o *Wire) SetInputMessageAccountabilityData(v string) {
	o.InputMessageAccountabilityData = &v
}

// GetIsBulk returns the IsBulk field value
func (o *Wire) GetIsBulk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBulk
}

// GetIsBulkOk returns a tuple with the IsBulk field value
// and a boolean to check if the value has been set.
func (o *Wire) GetIsBulkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBulk, true
}

// SetIsBulk sets field value
func (o *Wire) SetIsBulk(v bool) {
	o.IsBulk = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *Wire) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *Wire) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *Wire) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *Wire) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *Wire) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *Wire) SetNetwork(v string) {
	o.Network = &v
}

// GetOriginatingAccountId returns the OriginatingAccountId field value if set, zero value otherwise.
func (o *Wire) GetOriginatingAccountId() string {
	if o == nil || IsNil(o.OriginatingAccountId) {
		var ret string
		return ret
	}
	return *o.OriginatingAccountId
}

// GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetOriginatingAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingAccountId) {
		return nil, false
	}
	return o.OriginatingAccountId, true
}

// HasOriginatingAccountId returns a boolean if a field has been set.
func (o *Wire) HasOriginatingAccountId() bool {
	if o != nil && !IsNil(o.OriginatingAccountId) {
		return true
	}

	return false
}

// SetOriginatingAccountId gets a reference to the given string and assigns it to the OriginatingAccountId field.
func (o *Wire) SetOriginatingAccountId(v string) {
	o.OriginatingAccountId = &v
}

// GetOriginatingAccountNumber returns the OriginatingAccountNumber field value
func (o *Wire) GetOriginatingAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatingAccountNumber
}

// GetOriginatingAccountNumberOk returns a tuple with the OriginatingAccountNumber field value
// and a boolean to check if the value has been set.
func (o *Wire) GetOriginatingAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatingAccountNumber, true
}

// SetOriginatingAccountNumber sets field value
func (o *Wire) SetOriginatingAccountNumber(v string) {
	o.OriginatingAccountNumber = v
}

// GetReceivingAccountId returns the ReceivingAccountId field value if set, zero value otherwise.
func (o *Wire) GetReceivingAccountId() string {
	if o == nil || IsNil(o.ReceivingAccountId) {
		var ret string
		return ret
	}
	return *o.ReceivingAccountId
}

// GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetReceivingAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivingAccountId) {
		return nil, false
	}
	return o.ReceivingAccountId, true
}

// HasReceivingAccountId returns a boolean if a field has been set.
func (o *Wire) HasReceivingAccountId() bool {
	if o != nil && !IsNil(o.ReceivingAccountId) {
		return true
	}

	return false
}

// SetReceivingAccountId gets a reference to the given string and assigns it to the ReceivingAccountId field.
func (o *Wire) SetReceivingAccountId(v string) {
	o.ReceivingAccountId = &v
}

// GetReceivingAccountNumber returns the ReceivingAccountNumber field value
func (o *Wire) GetReceivingAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceivingAccountNumber
}

// GetReceivingAccountNumberOk returns a tuple with the ReceivingAccountNumber field value
// and a boolean to check if the value has been set.
func (o *Wire) GetReceivingAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivingAccountNumber, true
}

// SetReceivingAccountNumber sets field value
func (o *Wire) SetReceivingAccountNumber(v string) {
	o.ReceivingAccountNumber = v
}

// GetRecipientMessage returns the RecipientMessage field value if set, zero value otherwise.
func (o *Wire) GetRecipientMessage() string {
	if o == nil || IsNil(o.RecipientMessage) {
		var ret string
		return ret
	}
	return *o.RecipientMessage
}

// GetRecipientMessageOk returns a tuple with the RecipientMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetRecipientMessageOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientMessage) {
		return nil, false
	}
	return o.RecipientMessage, true
}

// HasRecipientMessage returns a boolean if a field has been set.
func (o *Wire) HasRecipientMessage() bool {
	if o != nil && !IsNil(o.RecipientMessage) {
		return true
	}

	return false
}

// SetRecipientMessage gets a reference to the given string and assigns it to the RecipientMessage field.
func (o *Wire) SetRecipientMessage(v string) {
	o.RecipientMessage = &v
}

// GetReturnData returns the ReturnData field value if set, zero value otherwise.
func (o *Wire) GetReturnData() ReturnData1 {
	if o == nil || IsNil(o.ReturnData) {
		var ret ReturnData1
		return ret
	}
	return *o.ReturnData
}

// GetReturnDataOk returns a tuple with the ReturnData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetReturnDataOk() (*ReturnData1, bool) {
	if o == nil || IsNil(o.ReturnData) {
		return nil, false
	}
	return o.ReturnData, true
}

// HasReturnData returns a boolean if a field has been set.
func (o *Wire) HasReturnData() bool {
	if o != nil && !IsNil(o.ReturnData) {
		return true
	}

	return false
}

// SetReturnData gets a reference to the given ReturnData1 and assigns it to the ReturnData field.
func (o *Wire) SetReturnData(v ReturnData1) {
	o.ReturnData = &v
}

// GetSenderReferenceId returns the SenderReferenceId field value
func (o *Wire) GetSenderReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderReferenceId
}

// GetSenderReferenceIdOk returns a tuple with the SenderReferenceId field value
// and a boolean to check if the value has been set.
func (o *Wire) GetSenderReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderReferenceId, true
}

// SetSenderReferenceId sets field value
func (o *Wire) SetSenderReferenceId(v string) {
	o.SenderReferenceId = v
}

// GetSettlementDate returns the SettlementDate field value if set, zero value otherwise.
func (o *Wire) GetSettlementDate() string {
	if o == nil || IsNil(o.SettlementDate) {
		var ret string
		return ret
	}
	return *o.SettlementDate
}

// GetSettlementDateOk returns a tuple with the SettlementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetSettlementDateOk() (*string, bool) {
	if o == nil || IsNil(o.SettlementDate) {
		return nil, false
	}
	return o.SettlementDate, true
}

// HasSettlementDate returns a boolean if a field has been set.
func (o *Wire) HasSettlementDate() bool {
	if o != nil && !IsNil(o.SettlementDate) {
		return true
	}

	return false
}

// SetSettlementDate gets a reference to the given string and assigns it to the SettlementDate field.
func (o *Wire) SetSettlementDate(v string) {
	o.SettlementDate = &v
}

// GetStatus returns the Status field value
func (o *Wire) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Wire) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Wire) SetStatus(v string) {
	o.Status = v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *Wire) GetStatusDetails() string {
	if o == nil || IsNil(o.StatusDetails) {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetStatusDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDetails) {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *Wire) HasStatusDetails() bool {
	if o != nil && !IsNil(o.StatusDetails) {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *Wire) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

// GetTransactionId returns the TransactionId field value
func (o *Wire) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *Wire) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *Wire) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetTransactionInId returns the TransactionInId field value if set, zero value otherwise.
func (o *Wire) GetTransactionInId() string {
	if o == nil || IsNil(o.TransactionInId) {
		var ret string
		return ret
	}
	return *o.TransactionInId
}

// GetTransactionInIdOk returns a tuple with the TransactionInId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetTransactionInIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionInId) {
		return nil, false
	}
	return o.TransactionInId, true
}

// HasTransactionInId returns a boolean if a field has been set.
func (o *Wire) HasTransactionInId() bool {
	if o != nil && !IsNil(o.TransactionInId) {
		return true
	}

	return false
}

// SetTransactionInId gets a reference to the given string and assigns it to the TransactionInId field.
func (o *Wire) SetTransactionInId(v string) {
	o.TransactionInId = &v
}

func (o Wire) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Wire) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.BankMessage) {
		toSerialize["bank_message"] = o.BankMessage
	}
	if !IsNil(o.BatchId) {
		toSerialize["batch_id"] = o.BatchId
	}
	if !IsNil(o.CaseId) {
		toSerialize["case_id"] = o.CaseId
	}
	toSerialize["creation_time"] = o.CreationTime
	toSerialize["currency"] = o.Currency
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	toSerialize["effective_date"] = o.EffectiveDate
	toSerialize["id"] = o.Id
	if !IsNil(o.InputMessageAccountabilityData) {
		toSerialize["input_message_accountability_data"] = o.InputMessageAccountabilityData
	}
	toSerialize["is_bulk"] = o.IsBulk
	toSerialize["last_updated_time"] = o.LastUpdatedTime
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.OriginatingAccountId) {
		toSerialize["originating_account_id"] = o.OriginatingAccountId
	}
	toSerialize["originating_account_number"] = o.OriginatingAccountNumber
	if !IsNil(o.ReceivingAccountId) {
		toSerialize["receiving_account_id"] = o.ReceivingAccountId
	}
	toSerialize["receiving_account_number"] = o.ReceivingAccountNumber
	if !IsNil(o.RecipientMessage) {
		toSerialize["recipient_message"] = o.RecipientMessage
	}
	if !IsNil(o.ReturnData) {
		toSerialize["return_data"] = o.ReturnData
	}
	toSerialize["sender_reference_id"] = o.SenderReferenceId
	if !IsNil(o.SettlementDate) {
		toSerialize["settlement_date"] = o.SettlementDate
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.StatusDetails) {
		toSerialize["status_details"] = o.StatusDetails
	}
	toSerialize["transaction_id"] = o.TransactionId
	if !IsNil(o.TransactionInId) {
		toSerialize["transaction_in_id"] = o.TransactionInId
	}
	return toSerialize, nil
}

func (o *Wire) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"creation_time",
		"currency",
		"effective_date",
		"id",
		"is_bulk",
		"last_updated_time",
		"originating_account_number",
		"receiving_account_number",
		"sender_reference_id",
		"status",
		"transaction_id",
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

	varWire := _Wire{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWire)

	if err != nil {
		return err
	}

	*o = Wire(varWire)

	return err
}

type NullableWire struct {
	value *Wire
	isSet bool
}

func (v NullableWire) Get() *Wire {
	return v.value
}

func (v *NullableWire) Set(val *Wire) {
	v.value = val
	v.isSet = true
}

func (v NullableWire) IsSet() bool {
	return v.isSet
}

func (v *NullableWire) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWire(val *Wire) *NullableWire {
	return &NullableWire{value: val, isSet: true}
}

func (v NullableWire) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWire) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
