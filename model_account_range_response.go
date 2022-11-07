/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// AccountRangeResponse struct for AccountRangeResponse
type AccountRangeResponse struct {
	AccountRange []int64 `json:"account_range"`
	// The bank ID
	BankId int32 `json:"bank_id"`
	// The ID of the BIN this account range belogns to
	BinId string `json:"bin_id"`
	// The timestamp representing when the account range was created
	CreationTime time.Time `json:"creation_time"`
	// The time when account range becomes inactive
	EndDate *time.Time `json:"end_date,omitempty"`
	// Account Range Id
	Id string `json:"id"`
	// Controls whether account range allows tokenization
	IsTokenizationEnabled bool `json:"is_tokenization_enabled"`
	// The timestamp representing when the account range was last modified
	LastModifiedTime time.Time `json:"last_modified_time"`
	// The partner ID
	PartnerId int32 `json:"partner_id"`
	PhysicalCardFormat *PhysicalCardFormat `json:"physical_card_format,omitempty"`
	// The time when account range becomes active
	StartDate *time.Time `json:"start_date,omitempty"`
}

// NewAccountRangeResponse instantiates a new AccountRangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRangeResponse(accountRange []int64, bankId int32, binId string, creationTime time.Time, id string, isTokenizationEnabled bool, lastModifiedTime time.Time, partnerId int32) *AccountRangeResponse {
	this := AccountRangeResponse{}
	this.AccountRange = accountRange
	this.BankId = bankId
	this.BinId = binId
	this.CreationTime = creationTime
	this.Id = id
	this.IsTokenizationEnabled = isTokenizationEnabled
	this.LastModifiedTime = lastModifiedTime
	this.PartnerId = partnerId
	return &this
}

// NewAccountRangeResponseWithDefaults instantiates a new AccountRangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRangeResponseWithDefaults() *AccountRangeResponse {
	this := AccountRangeResponse{}
	var isTokenizationEnabled bool = false
	this.IsTokenizationEnabled = isTokenizationEnabled
	return &this
}

// GetAccountRange returns the AccountRange field value
func (o *AccountRangeResponse) GetAccountRange() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.AccountRange
}

// GetAccountRangeOk returns a tuple with the AccountRange field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetAccountRangeOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountRange, true
}

// SetAccountRange sets field value
func (o *AccountRangeResponse) SetAccountRange(v []int64) {
	o.AccountRange = v
}

// GetBankId returns the BankId field value
func (o *AccountRangeResponse) GetBankId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetBankIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankId, true
}

// SetBankId sets field value
func (o *AccountRangeResponse) SetBankId(v int32) {
	o.BankId = v
}

// GetBinId returns the BinId field value
func (o *AccountRangeResponse) GetBinId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinId
}

// GetBinIdOk returns a tuple with the BinId field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetBinIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinId, true
}

// SetBinId sets field value
func (o *AccountRangeResponse) SetBinId(v string) {
	o.BinId = v
}

// GetCreationTime returns the CreationTime field value
func (o *AccountRangeResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *AccountRangeResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *AccountRangeResponse) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *AccountRangeResponse) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *AccountRangeResponse) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetId returns the Id field value
func (o *AccountRangeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountRangeResponse) SetId(v string) {
	o.Id = v
}

// GetIsTokenizationEnabled returns the IsTokenizationEnabled field value
func (o *AccountRangeResponse) GetIsTokenizationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTokenizationEnabled
}

// GetIsTokenizationEnabledOk returns a tuple with the IsTokenizationEnabled field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetIsTokenizationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTokenizationEnabled, true
}

// SetIsTokenizationEnabled sets field value
func (o *AccountRangeResponse) SetIsTokenizationEnabled(v bool) {
	o.IsTokenizationEnabled = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *AccountRangeResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *AccountRangeResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetPartnerId returns the PartnerId field value
func (o *AccountRangeResponse) GetPartnerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetPartnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerId, true
}

// SetPartnerId sets field value
func (o *AccountRangeResponse) SetPartnerId(v int32) {
	o.PartnerId = v
}

// GetPhysicalCardFormat returns the PhysicalCardFormat field value if set, zero value otherwise.
func (o *AccountRangeResponse) GetPhysicalCardFormat() PhysicalCardFormat {
	if o == nil || o.PhysicalCardFormat == nil {
		var ret PhysicalCardFormat
		return ret
	}
	return *o.PhysicalCardFormat
}

// GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool) {
	if o == nil || o.PhysicalCardFormat == nil {
		return nil, false
	}
	return o.PhysicalCardFormat, true
}

// HasPhysicalCardFormat returns a boolean if a field has been set.
func (o *AccountRangeResponse) HasPhysicalCardFormat() bool {
	if o != nil && o.PhysicalCardFormat != nil {
		return true
	}

	return false
}

// SetPhysicalCardFormat gets a reference to the given PhysicalCardFormat and assigns it to the PhysicalCardFormat field.
func (o *AccountRangeResponse) SetPhysicalCardFormat(v PhysicalCardFormat) {
	o.PhysicalCardFormat = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *AccountRangeResponse) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRangeResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *AccountRangeResponse) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *AccountRangeResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o AccountRangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_range"] = o.AccountRange
	}
	if true {
		toSerialize["bank_id"] = o.BankId
	}
	if true {
		toSerialize["bin_id"] = o.BinId
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["is_tokenization_enabled"] = o.IsTokenizationEnabled
	}
	if true {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId
	}
	if o.PhysicalCardFormat != nil {
		toSerialize["physical_card_format"] = o.PhysicalCardFormat
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableAccountRangeResponse struct {
	value *AccountRangeResponse
	isSet bool
}

func (v NullableAccountRangeResponse) Get() *AccountRangeResponse {
	return v.value
}

func (v *NullableAccountRangeResponse) Set(val *AccountRangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRangeResponse(val *AccountRangeResponse) *NullableAccountRangeResponse {
	return &NullableAccountRangeResponse{value: val, isSet: true}
}

func (v NullableAccountRangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


