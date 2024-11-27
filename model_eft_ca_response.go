/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EftCaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EftCaResponse{}

// EftCaResponse EFT transfer specific to Canada
type EftCaResponse struct {
	// Transfer amount in cents
	Amount int64 `json:"amount"`
	// The UUID of the Synctera customer resource that is the originator of the transfer.
	CustomerId string `json:"customer_id"`
	// Debit or credit sign
	DcSign string `json:"dc_sign"`
	// Additional information to be added to the transfer
	SourceData map[string]interface{} `json:"source_data,omitempty"`
	// The three digit transaction code that identifies the type of transaction. More information can be found here: https://www.payments.ca/sites/default/files/standard007eng.pdf.
	TransactionCode string `json:"transaction_code"`
	// The UUID of the Synctera account that is the destination of the transfer. For a transfer originated by the Synctera platform, this will be an external account resource, while for a transfer originated by the external account, this account will be an account resource.
	DestinationAccountId string `json:"destination_account_id"`
	// The account number of the destination account.
	DestinationAccountNumber string `json:"destination_account_number"`
	// The account owner name of the destination account.
	DestinationAccountOwnerName string `json:"destination_account_owner_name"`
	// The effective date of the transaction once it gets posted
	EffectiveDate string `json:"effective_date"`
	// Whether the transfer failed or not.
	Failed  *bool    `json:"failed,omitempty"`
	History []Action `json:"history,omitempty"`
	// ID of the transfer
	Id string `json:"id"`
	// Send the same day (use only is_same_day without specific effective_date).
	IsSameDay bool `json:"is_same_day"`
	// The network status of the transfer.
	NetworkStatus *string `json:"network_status,omitempty"`
	// The UUID of the Synctera account that is the origination of the transfer. For a transfer originated by the Synctera platform, this will be an account resource, while for a transfer originated by the external account, this will be an external account resource.
	OriginatingAccountId string `json:"originating_account_id"`
	// The account number of the originating account.
	OriginatingAccountNumber string `json:"originating_account_number"`
	// The account owner name of the origination account.
	OriginatingAccountOwnerName string `json:"originating_account_owner_name"`
	// The posting date of the transaction once it gets posted
	PostingDate *string `json:"posting_date,omitempty"`
	// The reference id of the transfer.
	ReferenceId string      `json:"reference_id"`
	Status      EftCaStatus `json:"status"`
	// The subtype of the transfer
	Subtype string `json:"subtype"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	TenantId string `json:"tenant_id"`
	// The related transaction id of the transfer.
	TransactionId *string `json:"transaction_id,omitempty"`
}

type _EftCaResponse EftCaResponse

// NewEftCaResponse instantiates a new EftCaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEftCaResponse(amount int64, customerId string, dcSign string, transactionCode string, destinationAccountId string, destinationAccountNumber string, destinationAccountOwnerName string, effectiveDate string, id string, isSameDay bool, originatingAccountId string, originatingAccountNumber string, originatingAccountOwnerName string, referenceId string, status EftCaStatus, subtype string, tenantId string) *EftCaResponse {
	this := EftCaResponse{}
	this.Amount = amount
	this.CustomerId = customerId
	this.DcSign = dcSign
	this.TransactionCode = transactionCode
	this.DestinationAccountId = destinationAccountId
	this.DestinationAccountNumber = destinationAccountNumber
	this.DestinationAccountOwnerName = destinationAccountOwnerName
	this.EffectiveDate = effectiveDate
	this.Id = id
	this.IsSameDay = isSameDay
	this.OriginatingAccountId = originatingAccountId
	this.OriginatingAccountNumber = originatingAccountNumber
	this.OriginatingAccountOwnerName = originatingAccountOwnerName
	this.ReferenceId = referenceId
	this.Status = status
	this.Subtype = subtype
	this.TenantId = tenantId
	return &this
}

// NewEftCaResponseWithDefaults instantiates a new EftCaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEftCaResponseWithDefaults() *EftCaResponse {
	this := EftCaResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *EftCaResponse) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *EftCaResponse) SetAmount(v int64) {
	o.Amount = v
}

// GetCustomerId returns the CustomerId field value
func (o *EftCaResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *EftCaResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetDcSign returns the DcSign field value
func (o *EftCaResponse) GetDcSign() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetDcSignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *EftCaResponse) SetDcSign(v string) {
	o.DcSign = v
}

// GetSourceData returns the SourceData field value if set, zero value otherwise.
func (o *EftCaResponse) GetSourceData() map[string]interface{} {
	if o == nil || IsNil(o.SourceData) {
		var ret map[string]interface{}
		return ret
	}
	return o.SourceData
}

// GetSourceDataOk returns a tuple with the SourceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetSourceDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SourceData) {
		return map[string]interface{}{}, false
	}
	return o.SourceData, true
}

// HasSourceData returns a boolean if a field has been set.
func (o *EftCaResponse) HasSourceData() bool {
	if o != nil && !IsNil(o.SourceData) {
		return true
	}

	return false
}

// SetSourceData gets a reference to the given map[string]interface{} and assigns it to the SourceData field.
func (o *EftCaResponse) SetSourceData(v map[string]interface{}) {
	o.SourceData = v
}

// GetTransactionCode returns the TransactionCode field value
func (o *EftCaResponse) GetTransactionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionCode
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetTransactionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionCode, true
}

// SetTransactionCode sets field value
func (o *EftCaResponse) SetTransactionCode(v string) {
	o.TransactionCode = v
}

// GetDestinationAccountId returns the DestinationAccountId field value
func (o *EftCaResponse) GetDestinationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationAccountId
}

// GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetDestinationAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationAccountId, true
}

// SetDestinationAccountId sets field value
func (o *EftCaResponse) SetDestinationAccountId(v string) {
	o.DestinationAccountId = v
}

// GetDestinationAccountNumber returns the DestinationAccountNumber field value
func (o *EftCaResponse) GetDestinationAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationAccountNumber
}

// GetDestinationAccountNumberOk returns a tuple with the DestinationAccountNumber field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetDestinationAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationAccountNumber, true
}

// SetDestinationAccountNumber sets field value
func (o *EftCaResponse) SetDestinationAccountNumber(v string) {
	o.DestinationAccountNumber = v
}

// GetDestinationAccountOwnerName returns the DestinationAccountOwnerName field value
func (o *EftCaResponse) GetDestinationAccountOwnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationAccountOwnerName
}

// GetDestinationAccountOwnerNameOk returns a tuple with the DestinationAccountOwnerName field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetDestinationAccountOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationAccountOwnerName, true
}

// SetDestinationAccountOwnerName sets field value
func (o *EftCaResponse) SetDestinationAccountOwnerName(v string) {
	o.DestinationAccountOwnerName = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *EftCaResponse) GetEffectiveDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetEffectiveDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *EftCaResponse) SetEffectiveDate(v string) {
	o.EffectiveDate = v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *EftCaResponse) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *EftCaResponse) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *EftCaResponse) SetFailed(v bool) {
	o.Failed = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *EftCaResponse) GetHistory() []Action {
	if o == nil || IsNil(o.History) {
		var ret []Action
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetHistoryOk() ([]Action, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *EftCaResponse) HasHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []Action and assigns it to the History field.
func (o *EftCaResponse) SetHistory(v []Action) {
	o.History = v
}

// GetId returns the Id field value
func (o *EftCaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EftCaResponse) SetId(v string) {
	o.Id = v
}

// GetIsSameDay returns the IsSameDay field value
func (o *EftCaResponse) GetIsSameDay() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSameDay
}

// GetIsSameDayOk returns a tuple with the IsSameDay field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetIsSameDayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSameDay, true
}

// SetIsSameDay sets field value
func (o *EftCaResponse) SetIsSameDay(v bool) {
	o.IsSameDay = v
}

// GetNetworkStatus returns the NetworkStatus field value if set, zero value otherwise.
func (o *EftCaResponse) GetNetworkStatus() string {
	if o == nil || IsNil(o.NetworkStatus) {
		var ret string
		return ret
	}
	return *o.NetworkStatus
}

// GetNetworkStatusOk returns a tuple with the NetworkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetNetworkStatusOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkStatus) {
		return nil, false
	}
	return o.NetworkStatus, true
}

// HasNetworkStatus returns a boolean if a field has been set.
func (o *EftCaResponse) HasNetworkStatus() bool {
	if o != nil && !IsNil(o.NetworkStatus) {
		return true
	}

	return false
}

// SetNetworkStatus gets a reference to the given string and assigns it to the NetworkStatus field.
func (o *EftCaResponse) SetNetworkStatus(v string) {
	o.NetworkStatus = &v
}

// GetOriginatingAccountId returns the OriginatingAccountId field value
func (o *EftCaResponse) GetOriginatingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatingAccountId
}

// GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetOriginatingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatingAccountId, true
}

// SetOriginatingAccountId sets field value
func (o *EftCaResponse) SetOriginatingAccountId(v string) {
	o.OriginatingAccountId = v
}

// GetOriginatingAccountNumber returns the OriginatingAccountNumber field value
func (o *EftCaResponse) GetOriginatingAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatingAccountNumber
}

// GetOriginatingAccountNumberOk returns a tuple with the OriginatingAccountNumber field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetOriginatingAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatingAccountNumber, true
}

// SetOriginatingAccountNumber sets field value
func (o *EftCaResponse) SetOriginatingAccountNumber(v string) {
	o.OriginatingAccountNumber = v
}

// GetOriginatingAccountOwnerName returns the OriginatingAccountOwnerName field value
func (o *EftCaResponse) GetOriginatingAccountOwnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatingAccountOwnerName
}

// GetOriginatingAccountOwnerNameOk returns a tuple with the OriginatingAccountOwnerName field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetOriginatingAccountOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatingAccountOwnerName, true
}

// SetOriginatingAccountOwnerName sets field value
func (o *EftCaResponse) SetOriginatingAccountOwnerName(v string) {
	o.OriginatingAccountOwnerName = v
}

// GetPostingDate returns the PostingDate field value if set, zero value otherwise.
func (o *EftCaResponse) GetPostingDate() string {
	if o == nil || IsNil(o.PostingDate) {
		var ret string
		return ret
	}
	return *o.PostingDate
}

// GetPostingDateOk returns a tuple with the PostingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetPostingDateOk() (*string, bool) {
	if o == nil || IsNil(o.PostingDate) {
		return nil, false
	}
	return o.PostingDate, true
}

// HasPostingDate returns a boolean if a field has been set.
func (o *EftCaResponse) HasPostingDate() bool {
	if o != nil && !IsNil(o.PostingDate) {
		return true
	}

	return false
}

// SetPostingDate gets a reference to the given string and assigns it to the PostingDate field.
func (o *EftCaResponse) SetPostingDate(v string) {
	o.PostingDate = &v
}

// GetReferenceId returns the ReferenceId field value
func (o *EftCaResponse) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *EftCaResponse) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetStatus returns the Status field value
func (o *EftCaResponse) GetStatus() EftCaStatus {
	if o == nil {
		var ret EftCaStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetStatusOk() (*EftCaStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EftCaResponse) SetStatus(v EftCaStatus) {
	o.Status = v
}

// GetSubtype returns the Subtype field value
func (o *EftCaResponse) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *EftCaResponse) SetSubtype(v string) {
	o.Subtype = v
}

// GetTenantId returns the TenantId field value
func (o *EftCaResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *EftCaResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *EftCaResponse) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftCaResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *EftCaResponse) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *EftCaResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o EftCaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EftCaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["dc_sign"] = o.DcSign
	if !IsNil(o.SourceData) {
		toSerialize["source_data"] = o.SourceData
	}
	toSerialize["transaction_code"] = o.TransactionCode
	toSerialize["destination_account_id"] = o.DestinationAccountId
	toSerialize["destination_account_number"] = o.DestinationAccountNumber
	toSerialize["destination_account_owner_name"] = o.DestinationAccountOwnerName
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
	toSerialize["originating_account_id"] = o.OriginatingAccountId
	toSerialize["originating_account_number"] = o.OriginatingAccountNumber
	toSerialize["originating_account_owner_name"] = o.OriginatingAccountOwnerName
	if !IsNil(o.PostingDate) {
		toSerialize["posting_date"] = o.PostingDate
	}
	toSerialize["reference_id"] = o.ReferenceId
	toSerialize["status"] = o.Status
	toSerialize["subtype"] = o.Subtype
	toSerialize["tenant_id"] = o.TenantId
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return toSerialize, nil
}

func (o *EftCaResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"customer_id",
		"dc_sign",
		"transaction_code",
		"destination_account_id",
		"destination_account_number",
		"destination_account_owner_name",
		"effective_date",
		"id",
		"is_same_day",
		"originating_account_id",
		"originating_account_number",
		"originating_account_owner_name",
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

	varEftCaResponse := _EftCaResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEftCaResponse)

	if err != nil {
		return err
	}

	*o = EftCaResponse(varEftCaResponse)

	return err
}

type NullableEftCaResponse struct {
	value *EftCaResponse
	isSet bool
}

func (v NullableEftCaResponse) Get() *EftCaResponse {
	return v.value
}

func (v *NullableEftCaResponse) Set(val *EftCaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEftCaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEftCaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEftCaResponse(val *EftCaResponse) *NullableEftCaResponse {
	return &NullableEftCaResponse{value: val, isSet: true}
}

func (v NullableEftCaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEftCaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
