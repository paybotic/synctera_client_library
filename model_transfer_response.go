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

// TransferResponse struct for TransferResponse
type TransferResponse struct {
	// The ID of the account to which the card will be linked
	AccountId string `json:"account_id"`
	// Amount of the transfer in cents
	Amount int32 `json:"amount"`
	// ISO-3166-1 Alpha-2 country code
	CountryCode string `json:"country_code"`
	CreatedTime time.Time `json:"created_time"`
	// ISO 4217  Alpha-3 currency code
	Currency string `json:"currency"`
	// The ID of the customer for which card is being activated
	CustomerId string `json:"customer_id"`
	// The id of the transfer
	Id string `json:"id"`
	LastModifiedTime time.Time `json:"last_modified_time"`
	Merchant Merchant `json:"merchant"`
	// The status of the transfer
	Status string `json:"status"`
	// The transaction ID
	TransactionId *string `json:"transaction_id,omitempty"`
	Type TransferType `json:"type"`
}

// NewTransferResponse instantiates a new TransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferResponse(accountId string, amount int32, countryCode string, createdTime time.Time, currency string, customerId string, id string, lastModifiedTime time.Time, merchant Merchant, status string, type_ TransferType) *TransferResponse {
	this := TransferResponse{}
	this.AccountId = accountId
	this.Amount = amount
	this.CountryCode = countryCode
	this.CreatedTime = createdTime
	this.Currency = currency
	this.CustomerId = customerId
	this.Id = id
	this.LastModifiedTime = lastModifiedTime
	this.Merchant = merchant
	this.Status = status
	this.Type = type_
	return &this
}

// NewTransferResponseWithDefaults instantiates a new TransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferResponseWithDefaults() *TransferResponse {
	this := TransferResponse{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *TransferResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransferResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmount returns the Amount field value
func (o *TransferResponse) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferResponse) SetAmount(v int32) {
	o.Amount = v
}

// GetCountryCode returns the CountryCode field value
func (o *TransferResponse) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *TransferResponse) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *TransferResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *TransferResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrency returns the Currency field value
func (o *TransferResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransferResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value
func (o *TransferResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *TransferResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetId returns the Id field value
func (o *TransferResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferResponse) SetId(v string) {
	o.Id = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *TransferResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *TransferResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetMerchant returns the Merchant field value
func (o *TransferResponse) GetMerchant() Merchant {
	if o == nil {
		var ret Merchant
		return ret
	}

	return o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetMerchantOk() (*Merchant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Merchant, true
}

// SetMerchant sets field value
func (o *TransferResponse) SetMerchant(v Merchant) {
	o.Merchant = v
}

// GetStatus returns the Status field value
func (o *TransferResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferResponse) SetStatus(v string) {
	o.Status = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *TransferResponse) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *TransferResponse) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *TransferResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetType returns the Type field value
func (o *TransferResponse) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferResponse) GetTypeOk() (*TransferType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferResponse) SetType(v TransferType) {
	o.Type = v
}

func (o TransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["country_code"] = o.CountryCode
	}
	if true {
		toSerialize["created_time"] = o.CreatedTime
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["merchant"] = o.Merchant
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTransferResponse struct {
	value *TransferResponse
	isSet bool
}

func (v NullableTransferResponse) Get() *TransferResponse {
	return v.value
}

func (v *NullableTransferResponse) Set(val *TransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferResponse(val *TransferResponse) *NullableTransferResponse {
	return &NullableTransferResponse{value: val, isSet: true}
}

func (v NullableTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


