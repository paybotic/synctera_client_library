/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Deposit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Deposit{}

// Deposit Deposit using remote deposit capture
type Deposit struct {
	// The ID of the account
	AccountId *string `json:"account_id,omitempty"`
	// ID of the uploaded image of the back of the check
	BackImageId *string `json:"back_image_id,omitempty"`
	// Unique ID for the business. Exactly one of `business_id` or `person_id` must be set.
	BusinessId *string `json:"business_id,omitempty"`
	// Amount on check in ISO 4217 minor currency units
	CheckAmount *int32 `json:"check_amount,omitempty"`
	// ISO 4217 currency code for the deposit amount
	DepositCurrency *string `json:"deposit_currency,omitempty"`
	// ID of the uploaded image of the front of the check
	FrontImageId *string `json:"front_image_id,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Unique ID for the person. Exactly one of `person_id` or `business_id` must be set.
	PersonId     *string    `json:"person_id,omitempty"`
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Date the deposit was captured, in RFC 3339 format
	DateCaptured *time.Time `json:"date_captured,omitempty"`
	// Date the deposit was processed, in RFC 3339 format
	DateProcessed *time.Time `json:"date_processed,omitempty"`
	// Amount deposited in ISO 4217 minor currency units
	DepositAmount *int32 `json:"deposit_amount,omitempty"`
	// Remote Check Deposit ID
	Id              *string    `json:"id,omitempty"`
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Account number of the issuer of the check, included if OCR is successful
	OcrAccountNumber *string `json:"ocr_account_number,omitempty"`
	// The unique check number for this check in the checkbook, included if OCR is successful and there is a check number on the check
	OcrCheckNumber *string `json:"ocr_check_number,omitempty"`
	// Routing number of the issuing bank, included if OCR is successful
	OcrRoutingNumber *string `json:"ocr_routing_number,omitempty"`
	// The status of the deposit
	Status *string `json:"status,omitempty"`
	// The ID of the transaction associated with this deposit
	TransactionId *string     `json:"transaction_id,omitempty"`
	VendorInfo    *VendorInfo `json:"vendor_info,omitempty"`
}

// NewDeposit instantiates a new Deposit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeposit() *Deposit {
	this := Deposit{}
	return &this
}

// NewDepositWithDefaults instantiates a new Deposit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositWithDefaults() *Deposit {
	this := Deposit{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Deposit) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Deposit) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Deposit) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBackImageId returns the BackImageId field value if set, zero value otherwise.
func (o *Deposit) GetBackImageId() string {
	if o == nil || IsNil(o.BackImageId) {
		var ret string
		return ret
	}
	return *o.BackImageId
}

// GetBackImageIdOk returns a tuple with the BackImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetBackImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackImageId) {
		return nil, false
	}
	return o.BackImageId, true
}

// HasBackImageId returns a boolean if a field has been set.
func (o *Deposit) HasBackImageId() bool {
	if o != nil && !IsNil(o.BackImageId) {
		return true
	}

	return false
}

// SetBackImageId gets a reference to the given string and assigns it to the BackImageId field.
func (o *Deposit) SetBackImageId(v string) {
	o.BackImageId = &v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *Deposit) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *Deposit) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *Deposit) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCheckAmount returns the CheckAmount field value if set, zero value otherwise.
func (o *Deposit) GetCheckAmount() int32 {
	if o == nil || IsNil(o.CheckAmount) {
		var ret int32
		return ret
	}
	return *o.CheckAmount
}

// GetCheckAmountOk returns a tuple with the CheckAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetCheckAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckAmount) {
		return nil, false
	}
	return o.CheckAmount, true
}

// HasCheckAmount returns a boolean if a field has been set.
func (o *Deposit) HasCheckAmount() bool {
	if o != nil && !IsNil(o.CheckAmount) {
		return true
	}

	return false
}

// SetCheckAmount gets a reference to the given int32 and assigns it to the CheckAmount field.
func (o *Deposit) SetCheckAmount(v int32) {
	o.CheckAmount = &v
}

// GetDepositCurrency returns the DepositCurrency field value if set, zero value otherwise.
func (o *Deposit) GetDepositCurrency() string {
	if o == nil || IsNil(o.DepositCurrency) {
		var ret string
		return ret
	}
	return *o.DepositCurrency
}

// GetDepositCurrencyOk returns a tuple with the DepositCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetDepositCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.DepositCurrency) {
		return nil, false
	}
	return o.DepositCurrency, true
}

// HasDepositCurrency returns a boolean if a field has been set.
func (o *Deposit) HasDepositCurrency() bool {
	if o != nil && !IsNil(o.DepositCurrency) {
		return true
	}

	return false
}

// SetDepositCurrency gets a reference to the given string and assigns it to the DepositCurrency field.
func (o *Deposit) SetDepositCurrency(v string) {
	o.DepositCurrency = &v
}

// GetFrontImageId returns the FrontImageId field value if set, zero value otherwise.
func (o *Deposit) GetFrontImageId() string {
	if o == nil || IsNil(o.FrontImageId) {
		var ret string
		return ret
	}
	return *o.FrontImageId
}

// GetFrontImageIdOk returns a tuple with the FrontImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetFrontImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.FrontImageId) {
		return nil, false
	}
	return o.FrontImageId, true
}

// HasFrontImageId returns a boolean if a field has been set.
func (o *Deposit) HasFrontImageId() bool {
	if o != nil && !IsNil(o.FrontImageId) {
		return true
	}

	return false
}

// SetFrontImageId gets a reference to the given string and assigns it to the FrontImageId field.
func (o *Deposit) SetFrontImageId(v string) {
	o.FrontImageId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Deposit) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Deposit) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Deposit) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *Deposit) GetPersonId() string {
	if o == nil || IsNil(o.PersonId) {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetPersonIdOk() (*string, bool) {
	if o == nil || IsNil(o.PersonId) {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *Deposit) HasPersonId() bool {
	if o != nil && !IsNil(o.PersonId) {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
func (o *Deposit) SetPersonId(v string) {
	o.PersonId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Deposit) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Deposit) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Deposit) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDateCaptured returns the DateCaptured field value if set, zero value otherwise.
func (o *Deposit) GetDateCaptured() time.Time {
	if o == nil || IsNil(o.DateCaptured) {
		var ret time.Time
		return ret
	}
	return *o.DateCaptured
}

// GetDateCapturedOk returns a tuple with the DateCaptured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetDateCapturedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCaptured) {
		return nil, false
	}
	return o.DateCaptured, true
}

// HasDateCaptured returns a boolean if a field has been set.
func (o *Deposit) HasDateCaptured() bool {
	if o != nil && !IsNil(o.DateCaptured) {
		return true
	}

	return false
}

// SetDateCaptured gets a reference to the given time.Time and assigns it to the DateCaptured field.
func (o *Deposit) SetDateCaptured(v time.Time) {
	o.DateCaptured = &v
}

// GetDateProcessed returns the DateProcessed field value if set, zero value otherwise.
func (o *Deposit) GetDateProcessed() time.Time {
	if o == nil || IsNil(o.DateProcessed) {
		var ret time.Time
		return ret
	}
	return *o.DateProcessed
}

// GetDateProcessedOk returns a tuple with the DateProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetDateProcessedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateProcessed) {
		return nil, false
	}
	return o.DateProcessed, true
}

// HasDateProcessed returns a boolean if a field has been set.
func (o *Deposit) HasDateProcessed() bool {
	if o != nil && !IsNil(o.DateProcessed) {
		return true
	}

	return false
}

// SetDateProcessed gets a reference to the given time.Time and assigns it to the DateProcessed field.
func (o *Deposit) SetDateProcessed(v time.Time) {
	o.DateProcessed = &v
}

// GetDepositAmount returns the DepositAmount field value if set, zero value otherwise.
func (o *Deposit) GetDepositAmount() int32 {
	if o == nil || IsNil(o.DepositAmount) {
		var ret int32
		return ret
	}
	return *o.DepositAmount
}

// GetDepositAmountOk returns a tuple with the DepositAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetDepositAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.DepositAmount) {
		return nil, false
	}
	return o.DepositAmount, true
}

// HasDepositAmount returns a boolean if a field has been set.
func (o *Deposit) HasDepositAmount() bool {
	if o != nil && !IsNil(o.DepositAmount) {
		return true
	}

	return false
}

// SetDepositAmount gets a reference to the given int32 and assigns it to the DepositAmount field.
func (o *Deposit) SetDepositAmount(v int32) {
	o.DepositAmount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Deposit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Deposit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Deposit) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Deposit) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Deposit) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Deposit) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetOcrAccountNumber returns the OcrAccountNumber field value if set, zero value otherwise.
func (o *Deposit) GetOcrAccountNumber() string {
	if o == nil || IsNil(o.OcrAccountNumber) {
		var ret string
		return ret
	}
	return *o.OcrAccountNumber
}

// GetOcrAccountNumberOk returns a tuple with the OcrAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetOcrAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OcrAccountNumber) {
		return nil, false
	}
	return o.OcrAccountNumber, true
}

// HasOcrAccountNumber returns a boolean if a field has been set.
func (o *Deposit) HasOcrAccountNumber() bool {
	if o != nil && !IsNil(o.OcrAccountNumber) {
		return true
	}

	return false
}

// SetOcrAccountNumber gets a reference to the given string and assigns it to the OcrAccountNumber field.
func (o *Deposit) SetOcrAccountNumber(v string) {
	o.OcrAccountNumber = &v
}

// GetOcrCheckNumber returns the OcrCheckNumber field value if set, zero value otherwise.
func (o *Deposit) GetOcrCheckNumber() string {
	if o == nil || IsNil(o.OcrCheckNumber) {
		var ret string
		return ret
	}
	return *o.OcrCheckNumber
}

// GetOcrCheckNumberOk returns a tuple with the OcrCheckNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetOcrCheckNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OcrCheckNumber) {
		return nil, false
	}
	return o.OcrCheckNumber, true
}

// HasOcrCheckNumber returns a boolean if a field has been set.
func (o *Deposit) HasOcrCheckNumber() bool {
	if o != nil && !IsNil(o.OcrCheckNumber) {
		return true
	}

	return false
}

// SetOcrCheckNumber gets a reference to the given string and assigns it to the OcrCheckNumber field.
func (o *Deposit) SetOcrCheckNumber(v string) {
	o.OcrCheckNumber = &v
}

// GetOcrRoutingNumber returns the OcrRoutingNumber field value if set, zero value otherwise.
func (o *Deposit) GetOcrRoutingNumber() string {
	if o == nil || IsNil(o.OcrRoutingNumber) {
		var ret string
		return ret
	}
	return *o.OcrRoutingNumber
}

// GetOcrRoutingNumberOk returns a tuple with the OcrRoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetOcrRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OcrRoutingNumber) {
		return nil, false
	}
	return o.OcrRoutingNumber, true
}

// HasOcrRoutingNumber returns a boolean if a field has been set.
func (o *Deposit) HasOcrRoutingNumber() bool {
	if o != nil && !IsNil(o.OcrRoutingNumber) {
		return true
	}

	return false
}

// SetOcrRoutingNumber gets a reference to the given string and assigns it to the OcrRoutingNumber field.
func (o *Deposit) SetOcrRoutingNumber(v string) {
	o.OcrRoutingNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Deposit) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Deposit) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Deposit) SetStatus(v string) {
	o.Status = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *Deposit) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *Deposit) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *Deposit) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetVendorInfo returns the VendorInfo field value if set, zero value otherwise.
func (o *Deposit) GetVendorInfo() VendorInfo {
	if o == nil || IsNil(o.VendorInfo) {
		var ret VendorInfo
		return ret
	}
	return *o.VendorInfo
}

// GetVendorInfoOk returns a tuple with the VendorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetVendorInfoOk() (*VendorInfo, bool) {
	if o == nil || IsNil(o.VendorInfo) {
		return nil, false
	}
	return o.VendorInfo, true
}

// HasVendorInfo returns a boolean if a field has been set.
func (o *Deposit) HasVendorInfo() bool {
	if o != nil && !IsNil(o.VendorInfo) {
		return true
	}

	return false
}

// SetVendorInfo gets a reference to the given VendorInfo and assigns it to the VendorInfo field.
func (o *Deposit) SetVendorInfo(v VendorInfo) {
	o.VendorInfo = &v
}

func (o Deposit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Deposit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.BackImageId) {
		toSerialize["back_image_id"] = o.BackImageId
	}
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	if !IsNil(o.CheckAmount) {
		toSerialize["check_amount"] = o.CheckAmount
	}
	if !IsNil(o.DepositCurrency) {
		toSerialize["deposit_currency"] = o.DepositCurrency
	}
	if !IsNil(o.FrontImageId) {
		toSerialize["front_image_id"] = o.FrontImageId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PersonId) {
		toSerialize["person_id"] = o.PersonId
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.DateCaptured) {
		toSerialize["date_captured"] = o.DateCaptured
	}
	if !IsNil(o.DateProcessed) {
		toSerialize["date_processed"] = o.DateProcessed
	}
	if !IsNil(o.DepositAmount) {
		toSerialize["deposit_amount"] = o.DepositAmount
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if !IsNil(o.OcrAccountNumber) {
		toSerialize["ocr_account_number"] = o.OcrAccountNumber
	}
	if !IsNil(o.OcrCheckNumber) {
		toSerialize["ocr_check_number"] = o.OcrCheckNumber
	}
	if !IsNil(o.OcrRoutingNumber) {
		toSerialize["ocr_routing_number"] = o.OcrRoutingNumber
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !IsNil(o.VendorInfo) {
		toSerialize["vendor_info"] = o.VendorInfo
	}
	return toSerialize, nil
}

type NullableDeposit struct {
	value *Deposit
	isSet bool
}

func (v NullableDeposit) Get() *Deposit {
	return v.value
}

func (v *NullableDeposit) Set(val *Deposit) {
	v.value = val
	v.isSet = true
}

func (v NullableDeposit) IsSet() bool {
	return v.isSet
}

func (v *NullableDeposit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeposit(val *Deposit) *NullableDeposit {
	return &NullableDeposit{value: val, isSet: true}
}

func (v NullableDeposit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeposit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
