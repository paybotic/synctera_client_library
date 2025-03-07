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

// checks if the OutgoingAchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutgoingAchRequest{}

// OutgoingAchRequest Send an ACH
type OutgoingAchRequest struct {
	// Amount to transfer in ISO 4217 minor currency units
	Amount int32 `json:"amount"`
	// Company Entry Description field in ACH batch header. Originator inserts this field's value to provide the Receiver with a description of the entry's purpose.
	CompanyEntryDescription NullableString `json:"company_entry_description,omitempty"`
	// Overrides the 'Company Name' field in ACH batch header, which otherwise defaults to the configured partner name. The provided name will be prepended with the Bank's configured prefix and a *. It will then be truncated to 16 characters.
	CompanyName *string `json:"company_name,omitempty"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	// The customer's unique identifier
	CustomerId string `json:"customer_id"`
	// The type of transaction (debit or credit). A debit is a transfer in and a credit is a transfer out of the originating account
	DcSign string `json:"dc_sign"`
	// Effective date transaction proccesses (is_same_day needs to be false or not present at all)
	EffectiveDate NullableString `json:"effective_date,omitempty"`
	// Additional transfer metadata structured as key-value pairs
	ExternalData map[string]interface{} `json:"external_data,omitempty"`
	// ID of the international customer that receives the final remittance transfer (required for OFAC enabled payments)
	FinalCustomerId NullableString             `json:"final_customer_id,omitempty"`
	Hold            NullableAchRequestHoldData `json:"hold,omitempty"`
	// Send as same day ACH transaction (use only is_same_day without specific effective_date)
	IsSameDay NullableBool `json:"is_same_day,omitempty"`
	// Memo for the payment
	Memo NullableString `json:"memo,omitempty"`
	// The unique identifier for an originating account
	OriginatingAccountId string `json:"originating_account_id"`
	// The unique identifier for an receiving account
	ReceivingAccountId string `json:"receiving_account_id"`
	// Will be sent to the ACH network and maps to Addenda record 05 - the recipient bank will receive this info
	ReferenceInfo NullableString   `json:"reference_info,omitempty"`
	Risk          NullableRiskData `json:"risk,omitempty"`
	// Standard Entry Class Code: * WEB: Internet initiated / Mobile Entry (default if empty) * CCD: Corporate Credit or Debit * PPD: Pre-arranged Payment or Deposit (only deposits currently supported)
	SecCode              *string `json:"sec_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OutgoingAchRequest OutgoingAchRequest

// NewOutgoingAchRequest instantiates a new OutgoingAchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutgoingAchRequest(amount int32, currency string, customerId string, dcSign string, originatingAccountId string, receivingAccountId string) *OutgoingAchRequest {
	this := OutgoingAchRequest{}
	this.Amount = amount
	this.Currency = currency
	this.CustomerId = customerId
	this.DcSign = dcSign
	this.OriginatingAccountId = originatingAccountId
	this.ReceivingAccountId = receivingAccountId
	var secCode string = "WEB"
	this.SecCode = &secCode
	return &this
}

// NewOutgoingAchRequestWithDefaults instantiates a new OutgoingAchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutgoingAchRequestWithDefaults() *OutgoingAchRequest {
	this := OutgoingAchRequest{}
	var secCode string = "WEB"
	this.SecCode = &secCode
	return &this
}

// GetAmount returns the Amount field value
func (o *OutgoingAchRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OutgoingAchRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OutgoingAchRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetCompanyEntryDescription returns the CompanyEntryDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingAchRequest) GetCompanyEntryDescription() string {
	if o == nil || IsNil(o.CompanyEntryDescription.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyEntryDescription.Get()
}

// GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingAchRequest) GetCompanyEntryDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyEntryDescription.Get(), o.CompanyEntryDescription.IsSet()
}

// HasCompanyEntryDescription returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasCompanyEntryDescription() bool {
	if o != nil && o.CompanyEntryDescription.IsSet() {
		return true
	}

	return false
}

// SetCompanyEntryDescription gets a reference to the given NullableString and assigns it to the CompanyEntryDescription field.
func (o *OutgoingAchRequest) SetCompanyEntryDescription(v string) {
	o.CompanyEntryDescription.Set(&v)
}

// SetCompanyEntryDescriptionNil sets the value for CompanyEntryDescription to be an explicit nil
func (o *OutgoingAchRequest) SetCompanyEntryDescriptionNil() {
	o.CompanyEntryDescription.Set(nil)
}

// UnsetCompanyEntryDescription ensures that no value is present for CompanyEntryDescription, not even an explicit nil
func (o *OutgoingAchRequest) UnsetCompanyEntryDescription() {
	o.CompanyEntryDescription.Unset()
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *OutgoingAchRequest) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAchRequest) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *OutgoingAchRequest) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCurrency returns the Currency field value
func (o *OutgoingAchRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *OutgoingAchRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *OutgoingAchRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value
func (o *OutgoingAchRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *OutgoingAchRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *OutgoingAchRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetDcSign returns the DcSign field value
func (o *OutgoingAchRequest) GetDcSign() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *OutgoingAchRequest) GetDcSignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *OutgoingAchRequest) SetDcSign(v string) {
	o.DcSign = v
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingAchRequest) GetEffectiveDate() string {
	if o == nil || IsNil(o.EffectiveDate.Get()) {
		var ret string
		return ret
	}
	return *o.EffectiveDate.Get()
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingAchRequest) GetEffectiveDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EffectiveDate.Get(), o.EffectiveDate.IsSet()
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasEffectiveDate() bool {
	if o != nil && o.EffectiveDate.IsSet() {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given NullableString and assigns it to the EffectiveDate field.
func (o *OutgoingAchRequest) SetEffectiveDate(v string) {
	o.EffectiveDate.Set(&v)
}

// SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil
func (o *OutgoingAchRequest) SetEffectiveDateNil() {
	o.EffectiveDate.Set(nil)
}

// UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
func (o *OutgoingAchRequest) UnsetEffectiveDate() {
	o.EffectiveDate.Unset()
}

// GetExternalData returns the ExternalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingAchRequest) GetExternalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingAchRequest) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExternalData) {
		return map[string]interface{}{}, false
	}
	return o.ExternalData, true
}

// HasExternalData returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasExternalData() bool {
	if o != nil && !IsNil(o.ExternalData) {
		return true
	}

	return false
}

// SetExternalData gets a reference to the given map[string]interface{} and assigns it to the ExternalData field.
func (o *OutgoingAchRequest) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetFinalCustomerId returns the FinalCustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingAchRequest) GetFinalCustomerId() string {
	if o == nil || IsNil(o.FinalCustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.FinalCustomerId.Get()
}

// GetFinalCustomerIdOk returns a tuple with the FinalCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingAchRequest) GetFinalCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinalCustomerId.Get(), o.FinalCustomerId.IsSet()
}

// HasFinalCustomerId returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasFinalCustomerId() bool {
	if o != nil && o.FinalCustomerId.IsSet() {
		return true
	}

	return false
}

// SetFinalCustomerId gets a reference to the given NullableString and assigns it to the FinalCustomerId field.
func (o *OutgoingAchRequest) SetFinalCustomerId(v string) {
	o.FinalCustomerId.Set(&v)
}

// SetFinalCustomerIdNil sets the value for FinalCustomerId to be an explicit nil
func (o *OutgoingAchRequest) SetFinalCustomerIdNil() {
	o.FinalCustomerId.Set(nil)
}

// UnsetFinalCustomerId ensures that no value is present for FinalCustomerId, not even an explicit nil
func (o *OutgoingAchRequest) UnsetFinalCustomerId() {
	o.FinalCustomerId.Unset()
}

// GetHold returns the Hold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingAchRequest) GetHold() AchRequestHoldData {
	if o == nil || IsNil(o.Hold.Get()) {
		var ret AchRequestHoldData
		return ret
	}
	return *o.Hold.Get()
}

// GetHoldOk returns a tuple with the Hold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingAchRequest) GetHoldOk() (*AchRequestHoldData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hold.Get(), o.Hold.IsSet()
}

// HasHold returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasHold() bool {
	if o != nil && o.Hold.IsSet() {
		return true
	}

	return false
}

// SetHold gets a reference to the given NullableAchRequestHoldData and assigns it to the Hold field.
func (o *OutgoingAchRequest) SetHold(v AchRequestHoldData) {
	o.Hold.Set(&v)
}

// SetHoldNil sets the value for Hold to be an explicit nil
func (o *OutgoingAchRequest) SetHoldNil() {
	o.Hold.Set(nil)
}

// UnsetHold ensures that no value is present for Hold, not even an explicit nil
func (o *OutgoingAchRequest) UnsetHold() {
	o.Hold.Unset()
}

// GetIsSameDay returns the IsSameDay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingAchRequest) GetIsSameDay() bool {
	if o == nil || IsNil(o.IsSameDay.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSameDay.Get()
}

// GetIsSameDayOk returns a tuple with the IsSameDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingAchRequest) GetIsSameDayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSameDay.Get(), o.IsSameDay.IsSet()
}

// HasIsSameDay returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasIsSameDay() bool {
	if o != nil && o.IsSameDay.IsSet() {
		return true
	}

	return false
}

// SetIsSameDay gets a reference to the given NullableBool and assigns it to the IsSameDay field.
func (o *OutgoingAchRequest) SetIsSameDay(v bool) {
	o.IsSameDay.Set(&v)
}

// SetIsSameDayNil sets the value for IsSameDay to be an explicit nil
func (o *OutgoingAchRequest) SetIsSameDayNil() {
	o.IsSameDay.Set(nil)
}

// UnsetIsSameDay ensures that no value is present for IsSameDay, not even an explicit nil
func (o *OutgoingAchRequest) UnsetIsSameDay() {
	o.IsSameDay.Unset()
}

// GetMemo returns the Memo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingAchRequest) GetMemo() string {
	if o == nil || IsNil(o.Memo.Get()) {
		var ret string
		return ret
	}
	return *o.Memo.Get()
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingAchRequest) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memo.Get(), o.Memo.IsSet()
}

// HasMemo returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasMemo() bool {
	if o != nil && o.Memo.IsSet() {
		return true
	}

	return false
}

// SetMemo gets a reference to the given NullableString and assigns it to the Memo field.
func (o *OutgoingAchRequest) SetMemo(v string) {
	o.Memo.Set(&v)
}

// SetMemoNil sets the value for Memo to be an explicit nil
func (o *OutgoingAchRequest) SetMemoNil() {
	o.Memo.Set(nil)
}

// UnsetMemo ensures that no value is present for Memo, not even an explicit nil
func (o *OutgoingAchRequest) UnsetMemo() {
	o.Memo.Unset()
}

// GetOriginatingAccountId returns the OriginatingAccountId field value
func (o *OutgoingAchRequest) GetOriginatingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatingAccountId
}

// GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field value
// and a boolean to check if the value has been set.
func (o *OutgoingAchRequest) GetOriginatingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatingAccountId, true
}

// SetOriginatingAccountId sets field value
func (o *OutgoingAchRequest) SetOriginatingAccountId(v string) {
	o.OriginatingAccountId = v
}

// GetReceivingAccountId returns the ReceivingAccountId field value
func (o *OutgoingAchRequest) GetReceivingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceivingAccountId
}

// GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field value
// and a boolean to check if the value has been set.
func (o *OutgoingAchRequest) GetReceivingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivingAccountId, true
}

// SetReceivingAccountId sets field value
func (o *OutgoingAchRequest) SetReceivingAccountId(v string) {
	o.ReceivingAccountId = v
}

// GetReferenceInfo returns the ReferenceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingAchRequest) GetReferenceInfo() string {
	if o == nil || IsNil(o.ReferenceInfo.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceInfo.Get()
}

// GetReferenceInfoOk returns a tuple with the ReferenceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingAchRequest) GetReferenceInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceInfo.Get(), o.ReferenceInfo.IsSet()
}

// HasReferenceInfo returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasReferenceInfo() bool {
	if o != nil && o.ReferenceInfo.IsSet() {
		return true
	}

	return false
}

// SetReferenceInfo gets a reference to the given NullableString and assigns it to the ReferenceInfo field.
func (o *OutgoingAchRequest) SetReferenceInfo(v string) {
	o.ReferenceInfo.Set(&v)
}

// SetReferenceInfoNil sets the value for ReferenceInfo to be an explicit nil
func (o *OutgoingAchRequest) SetReferenceInfoNil() {
	o.ReferenceInfo.Set(nil)
}

// UnsetReferenceInfo ensures that no value is present for ReferenceInfo, not even an explicit nil
func (o *OutgoingAchRequest) UnsetReferenceInfo() {
	o.ReferenceInfo.Unset()
}

// GetRisk returns the Risk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingAchRequest) GetRisk() RiskData {
	if o == nil || IsNil(o.Risk.Get()) {
		var ret RiskData
		return ret
	}
	return *o.Risk.Get()
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingAchRequest) GetRiskOk() (*RiskData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Risk.Get(), o.Risk.IsSet()
}

// HasRisk returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasRisk() bool {
	if o != nil && o.Risk.IsSet() {
		return true
	}

	return false
}

// SetRisk gets a reference to the given NullableRiskData and assigns it to the Risk field.
func (o *OutgoingAchRequest) SetRisk(v RiskData) {
	o.Risk.Set(&v)
}

// SetRiskNil sets the value for Risk to be an explicit nil
func (o *OutgoingAchRequest) SetRiskNil() {
	o.Risk.Set(nil)
}

// UnsetRisk ensures that no value is present for Risk, not even an explicit nil
func (o *OutgoingAchRequest) UnsetRisk() {
	o.Risk.Unset()
}

// GetSecCode returns the SecCode field value if set, zero value otherwise.
func (o *OutgoingAchRequest) GetSecCode() string {
	if o == nil || IsNil(o.SecCode) {
		var ret string
		return ret
	}
	return *o.SecCode
}

// GetSecCodeOk returns a tuple with the SecCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAchRequest) GetSecCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SecCode) {
		return nil, false
	}
	return o.SecCode, true
}

// HasSecCode returns a boolean if a field has been set.
func (o *OutgoingAchRequest) HasSecCode() bool {
	if o != nil && !IsNil(o.SecCode) {
		return true
	}

	return false
}

// SetSecCode gets a reference to the given string and assigns it to the SecCode field.
func (o *OutgoingAchRequest) SetSecCode(v string) {
	o.SecCode = &v
}

func (o OutgoingAchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutgoingAchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if o.CompanyEntryDescription.IsSet() {
		toSerialize["company_entry_description"] = o.CompanyEntryDescription.Get()
	}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	toSerialize["currency"] = o.Currency
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["dc_sign"] = o.DcSign
	if o.EffectiveDate.IsSet() {
		toSerialize["effective_date"] = o.EffectiveDate.Get()
	}
	if o.ExternalData != nil {
		toSerialize["external_data"] = o.ExternalData
	}
	if o.FinalCustomerId.IsSet() {
		toSerialize["final_customer_id"] = o.FinalCustomerId.Get()
	}
	if o.Hold.IsSet() {
		toSerialize["hold"] = o.Hold.Get()
	}
	if o.IsSameDay.IsSet() {
		toSerialize["is_same_day"] = o.IsSameDay.Get()
	}
	if o.Memo.IsSet() {
		toSerialize["memo"] = o.Memo.Get()
	}
	toSerialize["originating_account_id"] = o.OriginatingAccountId
	toSerialize["receiving_account_id"] = o.ReceivingAccountId
	if o.ReferenceInfo.IsSet() {
		toSerialize["reference_info"] = o.ReferenceInfo.Get()
	}
	if o.Risk.IsSet() {
		toSerialize["risk"] = o.Risk.Get()
	}
	if !IsNil(o.SecCode) {
		toSerialize["sec_code"] = o.SecCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OutgoingAchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"customer_id",
		"dc_sign",
		"originating_account_id",
		"receiving_account_id",
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

	varOutgoingAchRequest := _OutgoingAchRequest{}

	err = json.Unmarshal(data, &varOutgoingAchRequest)

	if err != nil {
		return err
	}

	*o = OutgoingAchRequest(varOutgoingAchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "company_entry_description")
		delete(additionalProperties, "company_name")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "dc_sign")
		delete(additionalProperties, "effective_date")
		delete(additionalProperties, "external_data")
		delete(additionalProperties, "final_customer_id")
		delete(additionalProperties, "hold")
		delete(additionalProperties, "is_same_day")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "originating_account_id")
		delete(additionalProperties, "receiving_account_id")
		delete(additionalProperties, "reference_info")
		delete(additionalProperties, "risk")
		delete(additionalProperties, "sec_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutgoingAchRequest struct {
	value *OutgoingAchRequest
	isSet bool
}

func (v NullableOutgoingAchRequest) Get() *OutgoingAchRequest {
	return v.value
}

func (v *NullableOutgoingAchRequest) Set(val *OutgoingAchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOutgoingAchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOutgoingAchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutgoingAchRequest(val *OutgoingAchRequest) *NullableOutgoingAchRequest {
	return &NullableOutgoingAchRequest{value: val, isSet: true}
}

func (v NullableOutgoingAchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutgoingAchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
