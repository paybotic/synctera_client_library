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

// BaseCashPickup struct for BaseCashPickup
type BaseCashPickup struct {
	AccountId string `json:"account_id"`
	// The amount (in cents) of the transaction
	Amount             int32              `json:"amount"`
	BusinessId         *string            `json:"business_id,omitempty"`
	CreationTime       *time.Time         `json:"creation_time,omitempty"`
	Currency           *string            `json:"currency,omitempty"`
	EmployeeCustomerId *string            `json:"employee_customer_id,omitempty"`
	Id                 *string            `json:"id,omitempty"`
	LastUpdatedTime    *time.Time         `json:"last_updated_time,omitempty"`
	Metadata           *map[string]string `json:"metadata,omitempty"`
	// the date when the money was actually picked up
	PickupTime   *time.Time `json:"pickup_time,omitempty"`
	PostedAmount *int32     `json:"posted_amount,omitempty"`
	// The date the transaction was posted. This is the date any money is considered to be added or removed from an account.
	PostedDate *string `json:"posted_date,omitempty"`
	// An external ID provided by the partner. This is not guaranteed to be globally unique.
	ReferenceId string `json:"reference_id"`
	// the date when the money was expected to be picked up
	ScheduledPickupDate *string `json:"scheduled_pickup_date,omitempty"`
	TransactionId       *string `json:"transaction_id,omitempty"`
}

// NewBaseCashPickup instantiates a new BaseCashPickup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseCashPickup(accountId string, amount int32, referenceId string) *BaseCashPickup {
	this := BaseCashPickup{}
	this.AccountId = accountId
	this.Amount = amount
	this.ReferenceId = referenceId
	return &this
}

// NewBaseCashPickupWithDefaults instantiates a new BaseCashPickup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseCashPickupWithDefaults() *BaseCashPickup {
	this := BaseCashPickup{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *BaseCashPickup) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *BaseCashPickup) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmount returns the Amount field value
func (o *BaseCashPickup) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BaseCashPickup) SetAmount(v int32) {
	o.Amount = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *BaseCashPickup) GetBusinessId() string {
	if o == nil || o.BusinessId == nil {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetBusinessIdOk() (*string, bool) {
	if o == nil || o.BusinessId == nil {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *BaseCashPickup) HasBusinessId() bool {
	if o != nil && o.BusinessId != nil {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *BaseCashPickup) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BaseCashPickup) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BaseCashPickup) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BaseCashPickup) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BaseCashPickup) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BaseCashPickup) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BaseCashPickup) SetCurrency(v string) {
	o.Currency = &v
}

// GetEmployeeCustomerId returns the EmployeeCustomerId field value if set, zero value otherwise.
func (o *BaseCashPickup) GetEmployeeCustomerId() string {
	if o == nil || o.EmployeeCustomerId == nil {
		var ret string
		return ret
	}
	return *o.EmployeeCustomerId
}

// GetEmployeeCustomerIdOk returns a tuple with the EmployeeCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetEmployeeCustomerIdOk() (*string, bool) {
	if o == nil || o.EmployeeCustomerId == nil {
		return nil, false
	}
	return o.EmployeeCustomerId, true
}

// HasEmployeeCustomerId returns a boolean if a field has been set.
func (o *BaseCashPickup) HasEmployeeCustomerId() bool {
	if o != nil && o.EmployeeCustomerId != nil {
		return true
	}

	return false
}

// SetEmployeeCustomerId gets a reference to the given string and assigns it to the EmployeeCustomerId field.
func (o *BaseCashPickup) SetEmployeeCustomerId(v string) {
	o.EmployeeCustomerId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseCashPickup) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseCashPickup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseCashPickup) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *BaseCashPickup) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *BaseCashPickup) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *BaseCashPickup) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BaseCashPickup) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BaseCashPickup) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *BaseCashPickup) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPickupTime returns the PickupTime field value if set, zero value otherwise.
func (o *BaseCashPickup) GetPickupTime() time.Time {
	if o == nil || o.PickupTime == nil {
		var ret time.Time
		return ret
	}
	return *o.PickupTime
}

// GetPickupTimeOk returns a tuple with the PickupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetPickupTimeOk() (*time.Time, bool) {
	if o == nil || o.PickupTime == nil {
		return nil, false
	}
	return o.PickupTime, true
}

// HasPickupTime returns a boolean if a field has been set.
func (o *BaseCashPickup) HasPickupTime() bool {
	if o != nil && o.PickupTime != nil {
		return true
	}

	return false
}

// SetPickupTime gets a reference to the given time.Time and assigns it to the PickupTime field.
func (o *BaseCashPickup) SetPickupTime(v time.Time) {
	o.PickupTime = &v
}

// GetPostedAmount returns the PostedAmount field value if set, zero value otherwise.
func (o *BaseCashPickup) GetPostedAmount() int32 {
	if o == nil || o.PostedAmount == nil {
		var ret int32
		return ret
	}
	return *o.PostedAmount
}

// GetPostedAmountOk returns a tuple with the PostedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetPostedAmountOk() (*int32, bool) {
	if o == nil || o.PostedAmount == nil {
		return nil, false
	}
	return o.PostedAmount, true
}

// HasPostedAmount returns a boolean if a field has been set.
func (o *BaseCashPickup) HasPostedAmount() bool {
	if o != nil && o.PostedAmount != nil {
		return true
	}

	return false
}

// SetPostedAmount gets a reference to the given int32 and assigns it to the PostedAmount field.
func (o *BaseCashPickup) SetPostedAmount(v int32) {
	o.PostedAmount = &v
}

// GetPostedDate returns the PostedDate field value if set, zero value otherwise.
func (o *BaseCashPickup) GetPostedDate() string {
	if o == nil || o.PostedDate == nil {
		var ret string
		return ret
	}
	return *o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetPostedDateOk() (*string, bool) {
	if o == nil || o.PostedDate == nil {
		return nil, false
	}
	return o.PostedDate, true
}

// HasPostedDate returns a boolean if a field has been set.
func (o *BaseCashPickup) HasPostedDate() bool {
	if o != nil && o.PostedDate != nil {
		return true
	}

	return false
}

// SetPostedDate gets a reference to the given string and assigns it to the PostedDate field.
func (o *BaseCashPickup) SetPostedDate(v string) {
	o.PostedDate = &v
}

// GetReferenceId returns the ReferenceId field value
func (o *BaseCashPickup) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *BaseCashPickup) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetScheduledPickupDate returns the ScheduledPickupDate field value if set, zero value otherwise.
func (o *BaseCashPickup) GetScheduledPickupDate() string {
	if o == nil || o.ScheduledPickupDate == nil {
		var ret string
		return ret
	}
	return *o.ScheduledPickupDate
}

// GetScheduledPickupDateOk returns a tuple with the ScheduledPickupDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetScheduledPickupDateOk() (*string, bool) {
	if o == nil || o.ScheduledPickupDate == nil {
		return nil, false
	}
	return o.ScheduledPickupDate, true
}

// HasScheduledPickupDate returns a boolean if a field has been set.
func (o *BaseCashPickup) HasScheduledPickupDate() bool {
	if o != nil && o.ScheduledPickupDate != nil {
		return true
	}

	return false
}

// SetScheduledPickupDate gets a reference to the given string and assigns it to the ScheduledPickupDate field.
func (o *BaseCashPickup) SetScheduledPickupDate(v string) {
	o.ScheduledPickupDate = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *BaseCashPickup) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCashPickup) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *BaseCashPickup) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *BaseCashPickup) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o BaseCashPickup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.BusinessId != nil {
		toSerialize["business_id"] = o.BusinessId
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.EmployeeCustomerId != nil {
		toSerialize["employee_customer_id"] = o.EmployeeCustomerId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PickupTime != nil {
		toSerialize["pickup_time"] = o.PickupTime
	}
	if o.PostedAmount != nil {
		toSerialize["posted_amount"] = o.PostedAmount
	}
	if o.PostedDate != nil {
		toSerialize["posted_date"] = o.PostedDate
	}
	if true {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if o.ScheduledPickupDate != nil {
		toSerialize["scheduled_pickup_date"] = o.ScheduledPickupDate
	}
	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableBaseCashPickup struct {
	value *BaseCashPickup
	isSet bool
}

func (v NullableBaseCashPickup) Get() *BaseCashPickup {
	return v.value
}

func (v *NullableBaseCashPickup) Set(val *BaseCashPickup) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseCashPickup) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseCashPickup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseCashPickup(val *BaseCashPickup) *NullableBaseCashPickup {
	return &NullableBaseCashPickup{value: val, isSet: true}
}

func (v NullableBaseCashPickup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseCashPickup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
