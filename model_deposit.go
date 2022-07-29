/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Deposit Deposit using remote deposit capture
type Deposit struct {
	// The ID of the account
	AccountId string `json:"account_id"`
	// ID of the uploaded image of the back of the check
	BackImageId string `json:"back_image_id"`
	// Unique ID for the business. Exactly one of `business_id` or `person_id` must be set. 
	BusinessId *string `json:"business_id,omitempty"`
	// Amount on check in ISO 4217 minor currency units
	CheckAmount int32 `json:"check_amount"`
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Date the deposit was captured, in RFC 3339 format
	DateCaptured time.Time `json:"date_captured"`
	// Date the deposit was processed, in RFC 3339 format
	DateProcessed time.Time `json:"date_processed"`
	// Amount deposited in ISO 4217 minor currency units
	DepositAmount int32 `json:"deposit_amount"`
	// ISO 4217 currency code for the deposit amount
	DepositCurrency string `json:"deposit_currency"`
	// ID of the uploaded image of the front of the check
	FrontImageId string `json:"front_image_id"`
	// Remote Check Deposit ID
	Id string `json:"id"`
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Unique ID for the person. Exactly one of `person_id` or `business_id` must be set. 
	PersonId *string `json:"person_id,omitempty"`
	// The status of the deposit
	Status string `json:"status"`
	// The ID of the transaction associated with this deposit
	TransactionId string `json:"transaction_id"`
	VendorInfo VendorInfo1 `json:"vendor_info"`
}

// NewDeposit instantiates a new Deposit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeposit(accountId string, backImageId string, checkAmount int32, dateCaptured time.Time, dateProcessed time.Time, depositAmount int32, depositCurrency string, frontImageId string, id string, status string, transactionId string, vendorInfo VendorInfo1) *Deposit {
	this := Deposit{}
	this.AccountId = accountId
	this.BackImageId = backImageId
	this.CheckAmount = checkAmount
	this.DateCaptured = dateCaptured
	this.DateProcessed = dateProcessed
	this.DepositAmount = depositAmount
	this.DepositCurrency = depositCurrency
	this.FrontImageId = frontImageId
	this.Id = id
	this.Status = status
	this.TransactionId = transactionId
	this.VendorInfo = vendorInfo
	return &this
}

// NewDepositWithDefaults instantiates a new Deposit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositWithDefaults() *Deposit {
	this := Deposit{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *Deposit) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Deposit) SetAccountId(v string) {
	o.AccountId = v
}

// GetBackImageId returns the BackImageId field value
func (o *Deposit) GetBackImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackImageId
}

// GetBackImageIdOk returns a tuple with the BackImageId field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetBackImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackImageId, true
}

// SetBackImageId sets field value
func (o *Deposit) SetBackImageId(v string) {
	o.BackImageId = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *Deposit) GetBusinessId() string {
	if o == nil || o.BusinessId == nil {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetBusinessIdOk() (*string, bool) {
	if o == nil || o.BusinessId == nil {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *Deposit) HasBusinessId() bool {
	if o != nil && o.BusinessId != nil {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *Deposit) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCheckAmount returns the CheckAmount field value
func (o *Deposit) GetCheckAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CheckAmount
}

// GetCheckAmountOk returns a tuple with the CheckAmount field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetCheckAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckAmount, true
}

// SetCheckAmount sets field value
func (o *Deposit) SetCheckAmount(v int32) {
	o.CheckAmount = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Deposit) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Deposit) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Deposit) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDateCaptured returns the DateCaptured field value
func (o *Deposit) GetDateCaptured() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateCaptured
}

// GetDateCapturedOk returns a tuple with the DateCaptured field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetDateCapturedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCaptured, true
}

// SetDateCaptured sets field value
func (o *Deposit) SetDateCaptured(v time.Time) {
	o.DateCaptured = v
}

// GetDateProcessed returns the DateProcessed field value
func (o *Deposit) GetDateProcessed() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateProcessed
}

// GetDateProcessedOk returns a tuple with the DateProcessed field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetDateProcessedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateProcessed, true
}

// SetDateProcessed sets field value
func (o *Deposit) SetDateProcessed(v time.Time) {
	o.DateProcessed = v
}

// GetDepositAmount returns the DepositAmount field value
func (o *Deposit) GetDepositAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DepositAmount
}

// GetDepositAmountOk returns a tuple with the DepositAmount field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetDepositAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositAmount, true
}

// SetDepositAmount sets field value
func (o *Deposit) SetDepositAmount(v int32) {
	o.DepositAmount = v
}

// GetDepositCurrency returns the DepositCurrency field value
func (o *Deposit) GetDepositCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositCurrency
}

// GetDepositCurrencyOk returns a tuple with the DepositCurrency field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetDepositCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositCurrency, true
}

// SetDepositCurrency sets field value
func (o *Deposit) SetDepositCurrency(v string) {
	o.DepositCurrency = v
}

// GetFrontImageId returns the FrontImageId field value
func (o *Deposit) GetFrontImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FrontImageId
}

// GetFrontImageIdOk returns a tuple with the FrontImageId field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetFrontImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrontImageId, true
}

// SetFrontImageId sets field value
func (o *Deposit) SetFrontImageId(v string) {
	o.FrontImageId = v
}

// GetId returns the Id field value
func (o *Deposit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Deposit) SetId(v string) {
	o.Id = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Deposit) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Deposit) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Deposit) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Deposit) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Deposit) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
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
	if o == nil || o.PersonId == nil {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deposit) GetPersonIdOk() (*string, bool) {
	if o == nil || o.PersonId == nil {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *Deposit) HasPersonId() bool {
	if o != nil && o.PersonId != nil {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
func (o *Deposit) SetPersonId(v string) {
	o.PersonId = &v
}

// GetStatus returns the Status field value
func (o *Deposit) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Deposit) SetStatus(v string) {
	o.Status = v
}

// GetTransactionId returns the TransactionId field value
func (o *Deposit) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *Deposit) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetVendorInfo returns the VendorInfo field value
func (o *Deposit) GetVendorInfo() VendorInfo1 {
	if o == nil {
		var ret VendorInfo1
		return ret
	}

	return o.VendorInfo
}

// GetVendorInfoOk returns a tuple with the VendorInfo field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetVendorInfoOk() (*VendorInfo1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorInfo, true
}

// SetVendorInfo sets field value
func (o *Deposit) SetVendorInfo(v VendorInfo1) {
	o.VendorInfo = v
}

func (o Deposit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["back_image_id"] = o.BackImageId
	}
	if o.BusinessId != nil {
		toSerialize["business_id"] = o.BusinessId
	}
	if true {
		toSerialize["check_amount"] = o.CheckAmount
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["date_captured"] = o.DateCaptured
	}
	if true {
		toSerialize["date_processed"] = o.DateProcessed
	}
	if true {
		toSerialize["deposit_amount"] = o.DepositAmount
	}
	if true {
		toSerialize["deposit_currency"] = o.DepositCurrency
	}
	if true {
		toSerialize["front_image_id"] = o.FrontImageId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PersonId != nil {
		toSerialize["person_id"] = o.PersonId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if true {
		toSerialize["vendor_info"] = o.VendorInfo
	}
	return json.Marshal(toSerialize)
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


