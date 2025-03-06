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

// checks if the InternalTransferResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalTransferResponse{}

// InternalTransferResponse struct for InternalTransferResponse
type InternalTransferResponse struct {
	// The amount (in cents) to transfer from originating account to receiving account.
	Amount int64 `json:"amount"`
	// Controls when the transfer will take effect. A value of `IMMEDIATE` (the default) means that the transfer will be completed immediately. A value of `MANUAL` means that the transaction will remain in a \"pending\" state until explicitly completed or cancelled (or the auth expires).
	CaptureMode *string `json:"capture_mode,omitempty"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	// When `capture_mode` is `MANUAL`, this field describes when the pending transaction should expire.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// The customer id of the international customer that receives the final remittance transfer (required for remittance payments).
	FinalCustomerId *string `json:"final_customer_id,omitempty"`
	// A short note to the recipient
	Memo *string `json:"memo,omitempty"`
	// Arbitrary key-value metadata to associate with the transaction
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// An alias representing a GL account to debit. This is alternative to specifying by account id
	OriginatingAccountAlias *string `json:"originating_account_alias,omitempty"`
	// The customer id of the owner of the originating account.
	OriginatingAccountCustomerId *string `json:"originating_account_customer_id,omitempty"`
	// The UUID of the account being debited
	OriginatingAccountId *string `json:"originating_account_id,omitempty"`
	// An alias representing a GL account to credit. This is an alternative to specifying by account id
	ReceivingAccountAlias *string `json:"receiving_account_alias,omitempty"`
	// The customer id of the owner of the receiving account.
	ReceivingAccountCustomerId *string `json:"receiving_account_customer_id,omitempty"`
	// The UUID of the account being credited
	ReceivingAccountId *string `json:"receiving_account_id,omitempty"`
	// Network reference id
	ReferenceId *string `json:"reference_id,omitempty"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant *string `json:"tenant,omitempty"`
	// The desired transaction type to use for this transfer
	Type string `json:"type"`
	// The transaction id associated with the transfer
	Id                   string                         `json:"id"`
	Status               InternalTransferResponseStatus `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _InternalTransferResponse InternalTransferResponse

// NewInternalTransferResponse instantiates a new InternalTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalTransferResponse(amount int64, currency string, type_ string, id string, status InternalTransferResponseStatus) *InternalTransferResponse {
	this := InternalTransferResponse{}
	this.Amount = amount
	var captureMode string = "IMMEDIATE"
	this.CaptureMode = &captureMode
	this.Currency = currency
	this.Type = type_
	this.Id = id
	this.Status = status
	return &this
}

// NewInternalTransferResponseWithDefaults instantiates a new InternalTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalTransferResponseWithDefaults() *InternalTransferResponse {
	this := InternalTransferResponse{}
	var captureMode string = "IMMEDIATE"
	this.CaptureMode = &captureMode
	return &this
}

// GetAmount returns the Amount field value
func (o *InternalTransferResponse) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InternalTransferResponse) SetAmount(v int64) {
	o.Amount = v
}

// GetCaptureMode returns the CaptureMode field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetCaptureMode() string {
	if o == nil || IsNil(o.CaptureMode) {
		var ret string
		return ret
	}
	return *o.CaptureMode
}

// GetCaptureModeOk returns a tuple with the CaptureMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetCaptureModeOk() (*string, bool) {
	if o == nil || IsNil(o.CaptureMode) {
		return nil, false
	}
	return o.CaptureMode, true
}

// HasCaptureMode returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasCaptureMode() bool {
	if o != nil && !IsNil(o.CaptureMode) {
		return true
	}

	return false
}

// SetCaptureMode gets a reference to the given string and assigns it to the CaptureMode field.
func (o *InternalTransferResponse) SetCaptureMode(v string) {
	o.CaptureMode = &v
}

// GetCurrency returns the Currency field value
func (o *InternalTransferResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InternalTransferResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *InternalTransferResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFinalCustomerId returns the FinalCustomerId field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetFinalCustomerId() string {
	if o == nil || IsNil(o.FinalCustomerId) {
		var ret string
		return ret
	}
	return *o.FinalCustomerId
}

// GetFinalCustomerIdOk returns a tuple with the FinalCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetFinalCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.FinalCustomerId) {
		return nil, false
	}
	return o.FinalCustomerId, true
}

// HasFinalCustomerId returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasFinalCustomerId() bool {
	if o != nil && !IsNil(o.FinalCustomerId) {
		return true
	}

	return false
}

// SetFinalCustomerId gets a reference to the given string and assigns it to the FinalCustomerId field.
func (o *InternalTransferResponse) SetFinalCustomerId(v string) {
	o.FinalCustomerId = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *InternalTransferResponse) SetMemo(v string) {
	o.Memo = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *InternalTransferResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetOriginatingAccountAlias returns the OriginatingAccountAlias field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetOriginatingAccountAlias() string {
	if o == nil || IsNil(o.OriginatingAccountAlias) {
		var ret string
		return ret
	}
	return *o.OriginatingAccountAlias
}

// GetOriginatingAccountAliasOk returns a tuple with the OriginatingAccountAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetOriginatingAccountAliasOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingAccountAlias) {
		return nil, false
	}
	return o.OriginatingAccountAlias, true
}

// HasOriginatingAccountAlias returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasOriginatingAccountAlias() bool {
	if o != nil && !IsNil(o.OriginatingAccountAlias) {
		return true
	}

	return false
}

// SetOriginatingAccountAlias gets a reference to the given string and assigns it to the OriginatingAccountAlias field.
func (o *InternalTransferResponse) SetOriginatingAccountAlias(v string) {
	o.OriginatingAccountAlias = &v
}

// GetOriginatingAccountCustomerId returns the OriginatingAccountCustomerId field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetOriginatingAccountCustomerId() string {
	if o == nil || IsNil(o.OriginatingAccountCustomerId) {
		var ret string
		return ret
	}
	return *o.OriginatingAccountCustomerId
}

// GetOriginatingAccountCustomerIdOk returns a tuple with the OriginatingAccountCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetOriginatingAccountCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingAccountCustomerId) {
		return nil, false
	}
	return o.OriginatingAccountCustomerId, true
}

// HasOriginatingAccountCustomerId returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasOriginatingAccountCustomerId() bool {
	if o != nil && !IsNil(o.OriginatingAccountCustomerId) {
		return true
	}

	return false
}

// SetOriginatingAccountCustomerId gets a reference to the given string and assigns it to the OriginatingAccountCustomerId field.
func (o *InternalTransferResponse) SetOriginatingAccountCustomerId(v string) {
	o.OriginatingAccountCustomerId = &v
}

// GetOriginatingAccountId returns the OriginatingAccountId field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetOriginatingAccountId() string {
	if o == nil || IsNil(o.OriginatingAccountId) {
		var ret string
		return ret
	}
	return *o.OriginatingAccountId
}

// GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetOriginatingAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingAccountId) {
		return nil, false
	}
	return o.OriginatingAccountId, true
}

// HasOriginatingAccountId returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasOriginatingAccountId() bool {
	if o != nil && !IsNil(o.OriginatingAccountId) {
		return true
	}

	return false
}

// SetOriginatingAccountId gets a reference to the given string and assigns it to the OriginatingAccountId field.
func (o *InternalTransferResponse) SetOriginatingAccountId(v string) {
	o.OriginatingAccountId = &v
}

// GetReceivingAccountAlias returns the ReceivingAccountAlias field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetReceivingAccountAlias() string {
	if o == nil || IsNil(o.ReceivingAccountAlias) {
		var ret string
		return ret
	}
	return *o.ReceivingAccountAlias
}

// GetReceivingAccountAliasOk returns a tuple with the ReceivingAccountAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetReceivingAccountAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivingAccountAlias) {
		return nil, false
	}
	return o.ReceivingAccountAlias, true
}

// HasReceivingAccountAlias returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasReceivingAccountAlias() bool {
	if o != nil && !IsNil(o.ReceivingAccountAlias) {
		return true
	}

	return false
}

// SetReceivingAccountAlias gets a reference to the given string and assigns it to the ReceivingAccountAlias field.
func (o *InternalTransferResponse) SetReceivingAccountAlias(v string) {
	o.ReceivingAccountAlias = &v
}

// GetReceivingAccountCustomerId returns the ReceivingAccountCustomerId field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetReceivingAccountCustomerId() string {
	if o == nil || IsNil(o.ReceivingAccountCustomerId) {
		var ret string
		return ret
	}
	return *o.ReceivingAccountCustomerId
}

// GetReceivingAccountCustomerIdOk returns a tuple with the ReceivingAccountCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetReceivingAccountCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivingAccountCustomerId) {
		return nil, false
	}
	return o.ReceivingAccountCustomerId, true
}

// HasReceivingAccountCustomerId returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasReceivingAccountCustomerId() bool {
	if o != nil && !IsNil(o.ReceivingAccountCustomerId) {
		return true
	}

	return false
}

// SetReceivingAccountCustomerId gets a reference to the given string and assigns it to the ReceivingAccountCustomerId field.
func (o *InternalTransferResponse) SetReceivingAccountCustomerId(v string) {
	o.ReceivingAccountCustomerId = &v
}

// GetReceivingAccountId returns the ReceivingAccountId field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetReceivingAccountId() string {
	if o == nil || IsNil(o.ReceivingAccountId) {
		var ret string
		return ret
	}
	return *o.ReceivingAccountId
}

// GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetReceivingAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivingAccountId) {
		return nil, false
	}
	return o.ReceivingAccountId, true
}

// HasReceivingAccountId returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasReceivingAccountId() bool {
	if o != nil && !IsNil(o.ReceivingAccountId) {
		return true
	}

	return false
}

// SetReceivingAccountId gets a reference to the given string and assigns it to the ReceivingAccountId field.
func (o *InternalTransferResponse) SetReceivingAccountId(v string) {
	o.ReceivingAccountId = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *InternalTransferResponse) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *InternalTransferResponse) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *InternalTransferResponse) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *InternalTransferResponse) SetTenant(v string) {
	o.Tenant = &v
}

// GetType returns the Type field value
func (o *InternalTransferResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InternalTransferResponse) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InternalTransferResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InternalTransferResponse) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *InternalTransferResponse) GetStatus() InternalTransferResponseStatus {
	if o == nil {
		var ret InternalTransferResponseStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InternalTransferResponse) GetStatusOk() (*InternalTransferResponseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InternalTransferResponse) SetStatus(v InternalTransferResponseStatus) {
	o.Status = v
}

func (o InternalTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CaptureMode) {
		toSerialize["capture_mode"] = o.CaptureMode
	}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.FinalCustomerId) {
		toSerialize["final_customer_id"] = o.FinalCustomerId
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.OriginatingAccountAlias) {
		toSerialize["originating_account_alias"] = o.OriginatingAccountAlias
	}
	if !IsNil(o.OriginatingAccountCustomerId) {
		toSerialize["originating_account_customer_id"] = o.OriginatingAccountCustomerId
	}
	if !IsNil(o.OriginatingAccountId) {
		toSerialize["originating_account_id"] = o.OriginatingAccountId
	}
	if !IsNil(o.ReceivingAccountAlias) {
		toSerialize["receiving_account_alias"] = o.ReceivingAccountAlias
	}
	if !IsNil(o.ReceivingAccountCustomerId) {
		toSerialize["receiving_account_customer_id"] = o.ReceivingAccountCustomerId
	}
	if !IsNil(o.ReceivingAccountId) {
		toSerialize["receiving_account_id"] = o.ReceivingAccountId
	}
	if !IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InternalTransferResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"type",
		"id",
		"status",
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

	varInternalTransferResponse := _InternalTransferResponse{}

	err = json.Unmarshal(data, &varInternalTransferResponse)

	if err != nil {
		return err
	}

	*o = InternalTransferResponse(varInternalTransferResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "capture_mode")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "final_customer_id")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "originating_account_alias")
		delete(additionalProperties, "originating_account_customer_id")
		delete(additionalProperties, "originating_account_id")
		delete(additionalProperties, "receiving_account_alias")
		delete(additionalProperties, "receiving_account_customer_id")
		delete(additionalProperties, "receiving_account_id")
		delete(additionalProperties, "reference_id")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInternalTransferResponse struct {
	value *InternalTransferResponse
	isSet bool
}

func (v NullableInternalTransferResponse) Get() *InternalTransferResponse {
	return v.value
}

func (v *NullableInternalTransferResponse) Set(val *InternalTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalTransferResponse(val *InternalTransferResponse) *NullableInternalTransferResponse {
	return &NullableInternalTransferResponse{value: val, isSet: true}
}

func (v NullableInternalTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
