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

// checks if the TransferResponsePull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferResponsePull{}

// TransferResponsePull struct for TransferResponsePull
type TransferResponsePull struct {
	// The ID of the account to which the card will be linked
	AccountId string `json:"account_id"`
	// Amount of the transfer in cents
	Amount int32 `json:"amount"`
	// ISO-3166-1 Alpha-2 country code
	CountryCode string `json:"country_code"`
	CreatedTime time.Time `json:"created_time"`
	// ISO 4217  Alpha-3 currency code
	Currency string `json:"currency"`
	// The ID of the customer to whom the card belongs
	CustomerId string `json:"customer_id"`
	// The ID of the external card from/to which the transfer was initiated/received
	ExternalCardId string `json:"external_card_id"`
	// The ID of the transfer
	Id string `json:"id"`
	LastModifiedTime time.Time `json:"last_modified_time"`
	Merchant Merchant `json:"merchant"`
	// The status of the transfer
	Status string `json:"status"`
	// The transaction ID
	TransactionId *string `json:"transaction_id,omitempty"`
	Type TransferType `json:"type"`
	// Unique identifier of an External Card Transfer 3-D Secure Authorization - conditionally required according to your program's 3DS policy
	ThreeDsId *string `json:"three_ds_id,omitempty"`
}

// NewTransferResponsePull instantiates a new TransferResponsePull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferResponsePull(accountId string, amount int32, countryCode string, createdTime time.Time, currency string, customerId string, externalCardId string, id string, lastModifiedTime time.Time, merchant Merchant, status string, type_ TransferType) *TransferResponsePull {
	this := TransferResponsePull{}
	this.AccountId = accountId
	this.Amount = amount
	this.CountryCode = countryCode
	this.CreatedTime = createdTime
	this.Currency = currency
	this.CustomerId = customerId
	this.ExternalCardId = externalCardId
	this.Id = id
	this.LastModifiedTime = lastModifiedTime
	this.Merchant = merchant
	this.Status = status
	this.Type = type_
	return &this
}

// NewTransferResponsePullWithDefaults instantiates a new TransferResponsePull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferResponsePullWithDefaults() *TransferResponsePull {
	this := TransferResponsePull{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *TransferResponsePull) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransferResponsePull) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmount returns the Amount field value
func (o *TransferResponsePull) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferResponsePull) SetAmount(v int32) {
	o.Amount = v
}

// GetCountryCode returns the CountryCode field value
func (o *TransferResponsePull) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *TransferResponsePull) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *TransferResponsePull) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *TransferResponsePull) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrency returns the Currency field value
func (o *TransferResponsePull) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransferResponsePull) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value
func (o *TransferResponsePull) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *TransferResponsePull) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetExternalCardId returns the ExternalCardId field value
func (o *TransferResponsePull) GetExternalCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalCardId
}

// GetExternalCardIdOk returns a tuple with the ExternalCardId field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetExternalCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCardId, true
}

// SetExternalCardId sets field value
func (o *TransferResponsePull) SetExternalCardId(v string) {
	o.ExternalCardId = v
}

// GetId returns the Id field value
func (o *TransferResponsePull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferResponsePull) SetId(v string) {
	o.Id = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *TransferResponsePull) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *TransferResponsePull) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetMerchant returns the Merchant field value
func (o *TransferResponsePull) GetMerchant() Merchant {
	if o == nil {
		var ret Merchant
		return ret
	}

	return o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetMerchantOk() (*Merchant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Merchant, true
}

// SetMerchant sets field value
func (o *TransferResponsePull) SetMerchant(v Merchant) {
	o.Merchant = v
}

// GetStatus returns the Status field value
func (o *TransferResponsePull) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferResponsePull) SetStatus(v string) {
	o.Status = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *TransferResponsePull) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *TransferResponsePull) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *TransferResponsePull) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetType returns the Type field value
func (o *TransferResponsePull) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetTypeOk() (*TransferType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferResponsePull) SetType(v TransferType) {
	o.Type = v
}

// GetThreeDsId returns the ThreeDsId field value if set, zero value otherwise.
func (o *TransferResponsePull) GetThreeDsId() string {
	if o == nil || IsNil(o.ThreeDsId) {
		var ret string
		return ret
	}
	return *o.ThreeDsId
}

// GetThreeDsIdOk returns a tuple with the ThreeDsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferResponsePull) GetThreeDsIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreeDsId) {
		return nil, false
	}
	return o.ThreeDsId, true
}

// HasThreeDsId returns a boolean if a field has been set.
func (o *TransferResponsePull) HasThreeDsId() bool {
	if o != nil && !IsNil(o.ThreeDsId) {
		return true
	}

	return false
}

// SetThreeDsId gets a reference to the given string and assigns it to the ThreeDsId field.
func (o *TransferResponsePull) SetThreeDsId(v string) {
	o.ThreeDsId = &v
}

func (o TransferResponsePull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferResponsePull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["amount"] = o.Amount
	toSerialize["country_code"] = o.CountryCode
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["currency"] = o.Currency
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["external_card_id"] = o.ExternalCardId
	toSerialize["id"] = o.Id
	toSerialize["last_modified_time"] = o.LastModifiedTime
	toSerialize["merchant"] = o.Merchant
	toSerialize["status"] = o.Status
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.ThreeDsId) {
		toSerialize["three_ds_id"] = o.ThreeDsId
	}
	return toSerialize, nil
}

type NullableTransferResponsePull struct {
	value *TransferResponsePull
	isSet bool
}

func (v NullableTransferResponsePull) Get() *TransferResponsePull {
	return v.value
}

func (v *NullableTransferResponsePull) Set(val *TransferResponsePull) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferResponsePull) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferResponsePull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferResponsePull(val *TransferResponsePull) *NullableTransferResponsePull {
	return &NullableTransferResponsePull{value: val, isSet: true}
}

func (v NullableTransferResponsePull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferResponsePull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

