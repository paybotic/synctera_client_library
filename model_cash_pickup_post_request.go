/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the CashPickupPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashPickupPostRequest{}

// CashPickupPostRequest struct for CashPickupPostRequest
type CashPickupPostRequest struct {
	AccountId string `json:"account_id"`
	// The amount (in cents) of the transaction
	Amount int32 `json:"amount"`
	BusinessId *string `json:"business_id,omitempty"`
	CreationTime *time.Time `json:"creation_time,omitempty"`
	Currency *string `json:"currency,omitempty"`
	EmployeeCustomerId *string `json:"employee_customer_id,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	// the date when the money was actually picked up
	PickupTime *time.Time `json:"pickup_time,omitempty"`
	PostedAmount *int32 `json:"posted_amount,omitempty"`
	// The date the transaction was posted. This is the date any money is considered to be added or removed from an account.
	PostedDate *string `json:"posted_date,omitempty"`
	// An external ID provided by the partner. This is not guaranteed to be globally unique.
	ReferenceId string `json:"reference_id"`
	// the date when the money was expected to be picked up
	ScheduledPickupDate *string `json:"scheduled_pickup_date,omitempty"`
	TransactionId *string `json:"transaction_id,omitempty"`
}

// NewCashPickupPostRequest instantiates a new CashPickupPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashPickupPostRequest(accountId string, amount int32, referenceId string) *CashPickupPostRequest {
	this := CashPickupPostRequest{}
	this.AccountId = accountId
	this.Amount = amount
	this.ReferenceId = referenceId
	return &this
}

// NewCashPickupPostRequestWithDefaults instantiates a new CashPickupPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashPickupPostRequestWithDefaults() *CashPickupPostRequest {
	this := CashPickupPostRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *CashPickupPostRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CashPickupPostRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmount returns the Amount field value
func (o *CashPickupPostRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CashPickupPostRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *CashPickupPostRequest) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *CashPickupPostRequest) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CashPickupPostRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetEmployeeCustomerId returns the EmployeeCustomerId field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetEmployeeCustomerId() string {
	if o == nil || IsNil(o.EmployeeCustomerId) {
		var ret string
		return ret
	}
	return *o.EmployeeCustomerId
}

// GetEmployeeCustomerIdOk returns a tuple with the EmployeeCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetEmployeeCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeCustomerId) {
		return nil, false
	}
	return o.EmployeeCustomerId, true
}

// HasEmployeeCustomerId returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasEmployeeCustomerId() bool {
	if o != nil && !IsNil(o.EmployeeCustomerId) {
		return true
	}

	return false
}

// SetEmployeeCustomerId gets a reference to the given string and assigns it to the EmployeeCustomerId field.
func (o *CashPickupPostRequest) SetEmployeeCustomerId(v string) {
	o.EmployeeCustomerId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CashPickupPostRequest) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *CashPickupPostRequest) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CashPickupPostRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPickupTime returns the PickupTime field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetPickupTime() time.Time {
	if o == nil || IsNil(o.PickupTime) {
		var ret time.Time
		return ret
	}
	return *o.PickupTime
}

// GetPickupTimeOk returns a tuple with the PickupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetPickupTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PickupTime) {
		return nil, false
	}
	return o.PickupTime, true
}

// HasPickupTime returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasPickupTime() bool {
	if o != nil && !IsNil(o.PickupTime) {
		return true
	}

	return false
}

// SetPickupTime gets a reference to the given time.Time and assigns it to the PickupTime field.
func (o *CashPickupPostRequest) SetPickupTime(v time.Time) {
	o.PickupTime = &v
}

// GetPostedAmount returns the PostedAmount field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetPostedAmount() int32 {
	if o == nil || IsNil(o.PostedAmount) {
		var ret int32
		return ret
	}
	return *o.PostedAmount
}

// GetPostedAmountOk returns a tuple with the PostedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetPostedAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.PostedAmount) {
		return nil, false
	}
	return o.PostedAmount, true
}

// HasPostedAmount returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasPostedAmount() bool {
	if o != nil && !IsNil(o.PostedAmount) {
		return true
	}

	return false
}

// SetPostedAmount gets a reference to the given int32 and assigns it to the PostedAmount field.
func (o *CashPickupPostRequest) SetPostedAmount(v int32) {
	o.PostedAmount = &v
}

// GetPostedDate returns the PostedDate field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetPostedDate() string {
	if o == nil || IsNil(o.PostedDate) {
		var ret string
		return ret
	}
	return *o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetPostedDateOk() (*string, bool) {
	if o == nil || IsNil(o.PostedDate) {
		return nil, false
	}
	return o.PostedDate, true
}

// HasPostedDate returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasPostedDate() bool {
	if o != nil && !IsNil(o.PostedDate) {
		return true
	}

	return false
}

// SetPostedDate gets a reference to the given string and assigns it to the PostedDate field.
func (o *CashPickupPostRequest) SetPostedDate(v string) {
	o.PostedDate = &v
}

// GetReferenceId returns the ReferenceId field value
func (o *CashPickupPostRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CashPickupPostRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetScheduledPickupDate returns the ScheduledPickupDate field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetScheduledPickupDate() string {
	if o == nil || IsNil(o.ScheduledPickupDate) {
		var ret string
		return ret
	}
	return *o.ScheduledPickupDate
}

// GetScheduledPickupDateOk returns a tuple with the ScheduledPickupDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetScheduledPickupDateOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduledPickupDate) {
		return nil, false
	}
	return o.ScheduledPickupDate, true
}

// HasScheduledPickupDate returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasScheduledPickupDate() bool {
	if o != nil && !IsNil(o.ScheduledPickupDate) {
		return true
	}

	return false
}

// SetScheduledPickupDate gets a reference to the given string and assigns it to the ScheduledPickupDate field.
func (o *CashPickupPostRequest) SetScheduledPickupDate(v string) {
	o.ScheduledPickupDate = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *CashPickupPostRequest) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupPostRequest) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *CashPickupPostRequest) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *CashPickupPostRequest) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o CashPickupPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashPickupPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["amount"] = o.Amount
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.EmployeeCustomerId) {
		toSerialize["employee_customer_id"] = o.EmployeeCustomerId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PickupTime) {
		toSerialize["pickup_time"] = o.PickupTime
	}
	if !IsNil(o.PostedAmount) {
		toSerialize["posted_amount"] = o.PostedAmount
	}
	if !IsNil(o.PostedDate) {
		toSerialize["posted_date"] = o.PostedDate
	}
	toSerialize["reference_id"] = o.ReferenceId
	if !IsNil(o.ScheduledPickupDate) {
		toSerialize["scheduled_pickup_date"] = o.ScheduledPickupDate
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return toSerialize, nil
}

type NullableCashPickupPostRequest struct {
	value *CashPickupPostRequest
	isSet bool
}

func (v NullableCashPickupPostRequest) Get() *CashPickupPostRequest {
	return v.value
}

func (v *NullableCashPickupPostRequest) Set(val *CashPickupPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCashPickupPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCashPickupPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashPickupPostRequest(val *CashPickupPostRequest) *NullableCashPickupPostRequest {
	return &NullableCashPickupPostRequest{value: val, isSet: true}
}

func (v NullableCashPickupPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashPickupPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


