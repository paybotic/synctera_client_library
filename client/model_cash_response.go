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
)

// checks if the CashResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashResponse{}

// CashResponse Cash transfer
type CashResponse struct {
	// Transfer amount in cents
	Amount int64 `json:"amount"`
	// Debit or credit sign
	DcSign string `json:"dc_sign"`
	// Additional information to be added to the transfer
	SourceData map[string]interface{} `json:"source_data,omitempty"`
	// The bank account of the client.
	ClientBankAccount *string `json:"client_bank_account,omitempty"`
	// The name of the client.
	ClientName *string `json:"client_name,omitempty"`
	// The UUID of the Synctera customer resource that is the originator of the transfer.
	CustomerId *string `json:"customer_id,omitempty"`
	// The UUID of the Synctera account that is the destination of the transfer. For a transfer originated by the Synctera platform, this will be an external account resource, while for a transfer originated by the external account, this account will be an account resource.
	DestinationAccountId *string `json:"destination_account_id,omitempty"`
	// The account number of the destination account.
	DestinationAccountNumber *string `json:"destination_account_number,omitempty"`
	// The account owner name of the destination account.
	DestinationAccountOwnerName *string `json:"destination_account_owner_name,omitempty"`
	// The effective date of the transaction once it gets posted
	EffectiveDate string `json:"effective_date"`
	// Whether the transfer failed or not.
	Failed  *bool    `json:"failed,omitempty"`
	History []Action `json:"history,omitempty"`
	// ID of the transfer
	Id string `json:"id"`
	// Send the same day (use only is_same_day without specific effective_date).
	IsSameDay     bool               `json:"is_same_day"`
	NetworkStatus *CashNetworkStatus `json:"network_status,omitempty"`
	// The original reference id of the transfer if it's a return.
	OriginalReferenceId *string `json:"original_reference_id,omitempty"`
	// The UUID of the Synctera account that is the origination of the transfer. For a transfer originated by the Synctera platform, this will be an account resource, while for a transfer originated by the external account, this will be an external account resource.
	OriginatingAccountId *string `json:"originating_account_id,omitempty"`
	// The account number of the originating account.
	OriginatingAccountNumber *string `json:"originating_account_number,omitempty"`
	// The account owner name of the origination account.
	OriginatingAccountOwnerName *string `json:"originating_account_owner_name,omitempty"`
	// The posting date of the transaction once it gets posted
	PostingDate *string `json:"posting_date,omitempty"`
	// The reference id of the transfer.
	ReferenceId string     `json:"reference_id"`
	Status      CashStatus `json:"status"`
	// The subtype of the transfer
	Subtype string `json:"subtype"`
	// Whether the transfer is suspended or not.
	Suspended *bool `json:"suspended,omitempty"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	TenantId string `json:"tenant_id"`
	// The related transaction id of the transfer.
	TransactionId        *string `json:"transaction_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CashResponse CashResponse

// NewCashResponse instantiates a new CashResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashResponse(amount int64, dcSign string, effectiveDate string, id string, isSameDay bool, referenceId string, status CashStatus, subtype string, tenantId string) *CashResponse {
	this := CashResponse{}
	this.Amount = amount
	this.DcSign = dcSign
	this.EffectiveDate = effectiveDate
	this.Id = id
	this.IsSameDay = isSameDay
	this.ReferenceId = referenceId
	this.Status = status
	this.Subtype = subtype
	this.TenantId = tenantId
	return &this
}

// NewCashResponseWithDefaults instantiates a new CashResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashResponseWithDefaults() *CashResponse {
	this := CashResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CashResponse) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CashResponse) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CashResponse) SetAmount(v int64) {
	o.Amount = v
}

// GetDcSign returns the DcSign field value
func (o *CashResponse) GetDcSign() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *CashResponse) GetDcSignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *CashResponse) SetDcSign(v string) {
	o.DcSign = v
}

// GetSourceData returns the SourceData field value if set, zero value otherwise.
func (o *CashResponse) GetSourceData() map[string]interface{} {
	if o == nil || IsNil(o.SourceData) {
		var ret map[string]interface{}
		return ret
	}
	return o.SourceData
}

// GetSourceDataOk returns a tuple with the SourceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetSourceDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SourceData) {
		return map[string]interface{}{}, false
	}
	return o.SourceData, true
}

// HasSourceData returns a boolean if a field has been set.
func (o *CashResponse) HasSourceData() bool {
	if o != nil && !IsNil(o.SourceData) {
		return true
	}

	return false
}

// SetSourceData gets a reference to the given map[string]interface{} and assigns it to the SourceData field.
func (o *CashResponse) SetSourceData(v map[string]interface{}) {
	o.SourceData = v
}

// GetClientBankAccount returns the ClientBankAccount field value if set, zero value otherwise.
func (o *CashResponse) GetClientBankAccount() string {
	if o == nil || IsNil(o.ClientBankAccount) {
		var ret string
		return ret
	}
	return *o.ClientBankAccount
}

// GetClientBankAccountOk returns a tuple with the ClientBankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetClientBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ClientBankAccount) {
		return nil, false
	}
	return o.ClientBankAccount, true
}

// HasClientBankAccount returns a boolean if a field has been set.
func (o *CashResponse) HasClientBankAccount() bool {
	if o != nil && !IsNil(o.ClientBankAccount) {
		return true
	}

	return false
}

// SetClientBankAccount gets a reference to the given string and assigns it to the ClientBankAccount field.
func (o *CashResponse) SetClientBankAccount(v string) {
	o.ClientBankAccount = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *CashResponse) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *CashResponse) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *CashResponse) SetClientName(v string) {
	o.ClientName = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CashResponse) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CashResponse) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CashResponse) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDestinationAccountId returns the DestinationAccountId field value if set, zero value otherwise.
func (o *CashResponse) GetDestinationAccountId() string {
	if o == nil || IsNil(o.DestinationAccountId) {
		var ret string
		return ret
	}
	return *o.DestinationAccountId
}

// GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetDestinationAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationAccountId) {
		return nil, false
	}
	return o.DestinationAccountId, true
}

// HasDestinationAccountId returns a boolean if a field has been set.
func (o *CashResponse) HasDestinationAccountId() bool {
	if o != nil && !IsNil(o.DestinationAccountId) {
		return true
	}

	return false
}

// SetDestinationAccountId gets a reference to the given string and assigns it to the DestinationAccountId field.
func (o *CashResponse) SetDestinationAccountId(v string) {
	o.DestinationAccountId = &v
}

// GetDestinationAccountNumber returns the DestinationAccountNumber field value if set, zero value otherwise.
func (o *CashResponse) GetDestinationAccountNumber() string {
	if o == nil || IsNil(o.DestinationAccountNumber) {
		var ret string
		return ret
	}
	return *o.DestinationAccountNumber
}

// GetDestinationAccountNumberOk returns a tuple with the DestinationAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetDestinationAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationAccountNumber) {
		return nil, false
	}
	return o.DestinationAccountNumber, true
}

// HasDestinationAccountNumber returns a boolean if a field has been set.
func (o *CashResponse) HasDestinationAccountNumber() bool {
	if o != nil && !IsNil(o.DestinationAccountNumber) {
		return true
	}

	return false
}

// SetDestinationAccountNumber gets a reference to the given string and assigns it to the DestinationAccountNumber field.
func (o *CashResponse) SetDestinationAccountNumber(v string) {
	o.DestinationAccountNumber = &v
}

// GetDestinationAccountOwnerName returns the DestinationAccountOwnerName field value if set, zero value otherwise.
func (o *CashResponse) GetDestinationAccountOwnerName() string {
	if o == nil || IsNil(o.DestinationAccountOwnerName) {
		var ret string
		return ret
	}
	return *o.DestinationAccountOwnerName
}

// GetDestinationAccountOwnerNameOk returns a tuple with the DestinationAccountOwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetDestinationAccountOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationAccountOwnerName) {
		return nil, false
	}
	return o.DestinationAccountOwnerName, true
}

// HasDestinationAccountOwnerName returns a boolean if a field has been set.
func (o *CashResponse) HasDestinationAccountOwnerName() bool {
	if o != nil && !IsNil(o.DestinationAccountOwnerName) {
		return true
	}

	return false
}

// SetDestinationAccountOwnerName gets a reference to the given string and assigns it to the DestinationAccountOwnerName field.
func (o *CashResponse) SetDestinationAccountOwnerName(v string) {
	o.DestinationAccountOwnerName = &v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *CashResponse) GetEffectiveDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *CashResponse) GetEffectiveDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *CashResponse) SetEffectiveDate(v string) {
	o.EffectiveDate = v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *CashResponse) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *CashResponse) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *CashResponse) SetFailed(v bool) {
	o.Failed = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *CashResponse) GetHistory() []Action {
	if o == nil || IsNil(o.History) {
		var ret []Action
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetHistoryOk() ([]Action, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *CashResponse) HasHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []Action and assigns it to the History field.
func (o *CashResponse) SetHistory(v []Action) {
	o.History = v
}

// GetId returns the Id field value
func (o *CashResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CashResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CashResponse) SetId(v string) {
	o.Id = v
}

// GetIsSameDay returns the IsSameDay field value
func (o *CashResponse) GetIsSameDay() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSameDay
}

// GetIsSameDayOk returns a tuple with the IsSameDay field value
// and a boolean to check if the value has been set.
func (o *CashResponse) GetIsSameDayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSameDay, true
}

// SetIsSameDay sets field value
func (o *CashResponse) SetIsSameDay(v bool) {
	o.IsSameDay = v
}

// GetNetworkStatus returns the NetworkStatus field value if set, zero value otherwise.
func (o *CashResponse) GetNetworkStatus() CashNetworkStatus {
	if o == nil || IsNil(o.NetworkStatus) {
		var ret CashNetworkStatus
		return ret
	}
	return *o.NetworkStatus
}

// GetNetworkStatusOk returns a tuple with the NetworkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetNetworkStatusOk() (*CashNetworkStatus, bool) {
	if o == nil || IsNil(o.NetworkStatus) {
		return nil, false
	}
	return o.NetworkStatus, true
}

// HasNetworkStatus returns a boolean if a field has been set.
func (o *CashResponse) HasNetworkStatus() bool {
	if o != nil && !IsNil(o.NetworkStatus) {
		return true
	}

	return false
}

// SetNetworkStatus gets a reference to the given CashNetworkStatus and assigns it to the NetworkStatus field.
func (o *CashResponse) SetNetworkStatus(v CashNetworkStatus) {
	o.NetworkStatus = &v
}

// GetOriginalReferenceId returns the OriginalReferenceId field value if set, zero value otherwise.
func (o *CashResponse) GetOriginalReferenceId() string {
	if o == nil || IsNil(o.OriginalReferenceId) {
		var ret string
		return ret
	}
	return *o.OriginalReferenceId
}

// GetOriginalReferenceIdOk returns a tuple with the OriginalReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetOriginalReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalReferenceId) {
		return nil, false
	}
	return o.OriginalReferenceId, true
}

// HasOriginalReferenceId returns a boolean if a field has been set.
func (o *CashResponse) HasOriginalReferenceId() bool {
	if o != nil && !IsNil(o.OriginalReferenceId) {
		return true
	}

	return false
}

// SetOriginalReferenceId gets a reference to the given string and assigns it to the OriginalReferenceId field.
func (o *CashResponse) SetOriginalReferenceId(v string) {
	o.OriginalReferenceId = &v
}

// GetOriginatingAccountId returns the OriginatingAccountId field value if set, zero value otherwise.
func (o *CashResponse) GetOriginatingAccountId() string {
	if o == nil || IsNil(o.OriginatingAccountId) {
		var ret string
		return ret
	}
	return *o.OriginatingAccountId
}

// GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetOriginatingAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingAccountId) {
		return nil, false
	}
	return o.OriginatingAccountId, true
}

// HasOriginatingAccountId returns a boolean if a field has been set.
func (o *CashResponse) HasOriginatingAccountId() bool {
	if o != nil && !IsNil(o.OriginatingAccountId) {
		return true
	}

	return false
}

// SetOriginatingAccountId gets a reference to the given string and assigns it to the OriginatingAccountId field.
func (o *CashResponse) SetOriginatingAccountId(v string) {
	o.OriginatingAccountId = &v
}

// GetOriginatingAccountNumber returns the OriginatingAccountNumber field value if set, zero value otherwise.
func (o *CashResponse) GetOriginatingAccountNumber() string {
	if o == nil || IsNil(o.OriginatingAccountNumber) {
		var ret string
		return ret
	}
	return *o.OriginatingAccountNumber
}

// GetOriginatingAccountNumberOk returns a tuple with the OriginatingAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetOriginatingAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingAccountNumber) {
		return nil, false
	}
	return o.OriginatingAccountNumber, true
}

// HasOriginatingAccountNumber returns a boolean if a field has been set.
func (o *CashResponse) HasOriginatingAccountNumber() bool {
	if o != nil && !IsNil(o.OriginatingAccountNumber) {
		return true
	}

	return false
}

// SetOriginatingAccountNumber gets a reference to the given string and assigns it to the OriginatingAccountNumber field.
func (o *CashResponse) SetOriginatingAccountNumber(v string) {
	o.OriginatingAccountNumber = &v
}

// GetOriginatingAccountOwnerName returns the OriginatingAccountOwnerName field value if set, zero value otherwise.
func (o *CashResponse) GetOriginatingAccountOwnerName() string {
	if o == nil || IsNil(o.OriginatingAccountOwnerName) {
		var ret string
		return ret
	}
	return *o.OriginatingAccountOwnerName
}

// GetOriginatingAccountOwnerNameOk returns a tuple with the OriginatingAccountOwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetOriginatingAccountOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingAccountOwnerName) {
		return nil, false
	}
	return o.OriginatingAccountOwnerName, true
}

// HasOriginatingAccountOwnerName returns a boolean if a field has been set.
func (o *CashResponse) HasOriginatingAccountOwnerName() bool {
	if o != nil && !IsNil(o.OriginatingAccountOwnerName) {
		return true
	}

	return false
}

// SetOriginatingAccountOwnerName gets a reference to the given string and assigns it to the OriginatingAccountOwnerName field.
func (o *CashResponse) SetOriginatingAccountOwnerName(v string) {
	o.OriginatingAccountOwnerName = &v
}

// GetPostingDate returns the PostingDate field value if set, zero value otherwise.
func (o *CashResponse) GetPostingDate() string {
	if o == nil || IsNil(o.PostingDate) {
		var ret string
		return ret
	}
	return *o.PostingDate
}

// GetPostingDateOk returns a tuple with the PostingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetPostingDateOk() (*string, bool) {
	if o == nil || IsNil(o.PostingDate) {
		return nil, false
	}
	return o.PostingDate, true
}

// HasPostingDate returns a boolean if a field has been set.
func (o *CashResponse) HasPostingDate() bool {
	if o != nil && !IsNil(o.PostingDate) {
		return true
	}

	return false
}

// SetPostingDate gets a reference to the given string and assigns it to the PostingDate field.
func (o *CashResponse) SetPostingDate(v string) {
	o.PostingDate = &v
}

// GetReferenceId returns the ReferenceId field value
func (o *CashResponse) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CashResponse) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CashResponse) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetStatus returns the Status field value
func (o *CashResponse) GetStatus() CashStatus {
	if o == nil {
		var ret CashStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CashResponse) GetStatusOk() (*CashStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CashResponse) SetStatus(v CashStatus) {
	o.Status = v
}

// GetSubtype returns the Subtype field value
func (o *CashResponse) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *CashResponse) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *CashResponse) SetSubtype(v string) {
	o.Subtype = v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *CashResponse) GetSuspended() bool {
	if o == nil || IsNil(o.Suspended) {
		var ret bool
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetSuspendedOk() (*bool, bool) {
	if o == nil || IsNil(o.Suspended) {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *CashResponse) HasSuspended() bool {
	if o != nil && !IsNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given bool and assigns it to the Suspended field.
func (o *CashResponse) SetSuspended(v bool) {
	o.Suspended = &v
}

// GetTenantId returns the TenantId field value
func (o *CashResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *CashResponse) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *CashResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *CashResponse) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *CashResponse) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *CashResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o CashResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["dc_sign"] = o.DcSign
	if !IsNil(o.SourceData) {
		toSerialize["source_data"] = o.SourceData
	}
	if !IsNil(o.ClientBankAccount) {
		toSerialize["client_bank_account"] = o.ClientBankAccount
	}
	if !IsNil(o.ClientName) {
		toSerialize["client_name"] = o.ClientName
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.DestinationAccountId) {
		toSerialize["destination_account_id"] = o.DestinationAccountId
	}
	if !IsNil(o.DestinationAccountNumber) {
		toSerialize["destination_account_number"] = o.DestinationAccountNumber
	}
	if !IsNil(o.DestinationAccountOwnerName) {
		toSerialize["destination_account_owner_name"] = o.DestinationAccountOwnerName
	}
	toSerialize["effective_date"] = o.EffectiveDate
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	toSerialize["id"] = o.Id
	toSerialize["is_same_day"] = o.IsSameDay
	if !IsNil(o.NetworkStatus) {
		toSerialize["network_status"] = o.NetworkStatus
	}
	if !IsNil(o.OriginalReferenceId) {
		toSerialize["original_reference_id"] = o.OriginalReferenceId
	}
	if !IsNil(o.OriginatingAccountId) {
		toSerialize["originating_account_id"] = o.OriginatingAccountId
	}
	if !IsNil(o.OriginatingAccountNumber) {
		toSerialize["originating_account_number"] = o.OriginatingAccountNumber
	}
	if !IsNil(o.OriginatingAccountOwnerName) {
		toSerialize["originating_account_owner_name"] = o.OriginatingAccountOwnerName
	}
	if !IsNil(o.PostingDate) {
		toSerialize["posting_date"] = o.PostingDate
	}
	toSerialize["reference_id"] = o.ReferenceId
	toSerialize["status"] = o.Status
	toSerialize["subtype"] = o.Subtype
	if !IsNil(o.Suspended) {
		toSerialize["suspended"] = o.Suspended
	}
	toSerialize["tenant_id"] = o.TenantId
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CashResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"dc_sign",
		"effective_date",
		"id",
		"is_same_day",
		"reference_id",
		"status",
		"subtype",
		"tenant_id",
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

	varCashResponse := _CashResponse{}

	err = json.Unmarshal(data, &varCashResponse)

	if err != nil {
		return err
	}

	*o = CashResponse(varCashResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "dc_sign")
		delete(additionalProperties, "source_data")
		delete(additionalProperties, "client_bank_account")
		delete(additionalProperties, "client_name")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "destination_account_id")
		delete(additionalProperties, "destination_account_number")
		delete(additionalProperties, "destination_account_owner_name")
		delete(additionalProperties, "effective_date")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "history")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_same_day")
		delete(additionalProperties, "network_status")
		delete(additionalProperties, "original_reference_id")
		delete(additionalProperties, "originating_account_id")
		delete(additionalProperties, "originating_account_number")
		delete(additionalProperties, "originating_account_owner_name")
		delete(additionalProperties, "posting_date")
		delete(additionalProperties, "reference_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "suspended")
		delete(additionalProperties, "tenant_id")
		delete(additionalProperties, "transaction_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCashResponse struct {
	value *CashResponse
	isSet bool
}

func (v NullableCashResponse) Get() *CashResponse {
	return v.value
}

func (v *NullableCashResponse) Set(val *CashResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCashResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCashResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashResponse(val *CashResponse) *NullableCashResponse {
	return &NullableCashResponse{value: val, isSet: true}
}

func (v NullableCashResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
