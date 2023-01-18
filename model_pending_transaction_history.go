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

// PendingTransactionHistory struct for PendingTransactionHistory
type PendingTransactionHistory struct {
	// The account id associated with the hold
	AccountId string `json:"account_id"`
	// The account number associated with the hold
	AccountNo string `json:"account_no"`
	// The creation date of the hold
	Created time.Time `json:"created"`
	Data *PendingTransactionHistoryData `json:"data,omitempty"`
	Id int64 `json:"id"`
	// The idempotency key used when initially creating this transaction.
	Idemkey string `json:"idemkey"`
	// The offset account id associated with the hold
	OffsetAccountId *string `json:"offset_account_id,omitempty"`
	// The offset account number associated with the hold
	OffsetAccountNo *string `json:"offset_account_no,omitempty"`
	// An external ID provided by the payment network to represent this transaction.
	ReferenceId NullableString `json:"reference_id"`
	// The tenant associated with this transaction, in the form \"<bankid>_<partnerid>\"
	Tenant string `json:"tenant"`
	// The date the hold was last update
	Updated time.Time `json:"updated"`
	// The unique identifier of the hold transaction.
	Uuid string `json:"uuid"`
}

// NewPendingTransactionHistory instantiates a new PendingTransactionHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingTransactionHistory(accountId string, accountNo string, created time.Time, id int64, idemkey string, referenceId NullableString, tenant string, updated time.Time, uuid string) *PendingTransactionHistory {
	this := PendingTransactionHistory{}
	this.AccountId = accountId
	this.AccountNo = accountNo
	this.Created = created
	this.Id = id
	this.Idemkey = idemkey
	this.ReferenceId = referenceId
	this.Tenant = tenant
	this.Updated = updated
	this.Uuid = uuid
	return &this
}

// NewPendingTransactionHistoryWithDefaults instantiates a new PendingTransactionHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingTransactionHistoryWithDefaults() *PendingTransactionHistory {
	this := PendingTransactionHistory{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *PendingTransactionHistory) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *PendingTransactionHistory) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountNo returns the AccountNo field value
func (o *PendingTransactionHistory) GetAccountNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetAccountNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNo, true
}

// SetAccountNo sets field value
func (o *PendingTransactionHistory) SetAccountNo(v string) {
	o.AccountNo = v
}

// GetCreated returns the Created field value
func (o *PendingTransactionHistory) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PendingTransactionHistory) SetCreated(v time.Time) {
	o.Created = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PendingTransactionHistory) GetData() PendingTransactionHistoryData {
	if o == nil || o.Data == nil {
		var ret PendingTransactionHistoryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetDataOk() (*PendingTransactionHistoryData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PendingTransactionHistory) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PendingTransactionHistoryData and assigns it to the Data field.
func (o *PendingTransactionHistory) SetData(v PendingTransactionHistoryData) {
	o.Data = &v
}

// GetId returns the Id field value
func (o *PendingTransactionHistory) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PendingTransactionHistory) SetId(v int64) {
	o.Id = v
}

// GetIdemkey returns the Idemkey field value
func (o *PendingTransactionHistory) GetIdemkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Idemkey
}

// GetIdemkeyOk returns a tuple with the Idemkey field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetIdemkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idemkey, true
}

// SetIdemkey sets field value
func (o *PendingTransactionHistory) SetIdemkey(v string) {
	o.Idemkey = v
}

// GetOffsetAccountId returns the OffsetAccountId field value if set, zero value otherwise.
func (o *PendingTransactionHistory) GetOffsetAccountId() string {
	if o == nil || o.OffsetAccountId == nil {
		var ret string
		return ret
	}
	return *o.OffsetAccountId
}

// GetOffsetAccountIdOk returns a tuple with the OffsetAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetOffsetAccountIdOk() (*string, bool) {
	if o == nil || o.OffsetAccountId == nil {
		return nil, false
	}
	return o.OffsetAccountId, true
}

// HasOffsetAccountId returns a boolean if a field has been set.
func (o *PendingTransactionHistory) HasOffsetAccountId() bool {
	if o != nil && o.OffsetAccountId != nil {
		return true
	}

	return false
}

// SetOffsetAccountId gets a reference to the given string and assigns it to the OffsetAccountId field.
func (o *PendingTransactionHistory) SetOffsetAccountId(v string) {
	o.OffsetAccountId = &v
}

// GetOffsetAccountNo returns the OffsetAccountNo field value if set, zero value otherwise.
func (o *PendingTransactionHistory) GetOffsetAccountNo() string {
	if o == nil || o.OffsetAccountNo == nil {
		var ret string
		return ret
	}
	return *o.OffsetAccountNo
}

// GetOffsetAccountNoOk returns a tuple with the OffsetAccountNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetOffsetAccountNoOk() (*string, bool) {
	if o == nil || o.OffsetAccountNo == nil {
		return nil, false
	}
	return o.OffsetAccountNo, true
}

// HasOffsetAccountNo returns a boolean if a field has been set.
func (o *PendingTransactionHistory) HasOffsetAccountNo() bool {
	if o != nil && o.OffsetAccountNo != nil {
		return true
	}

	return false
}

// SetOffsetAccountNo gets a reference to the given string and assigns it to the OffsetAccountNo field.
func (o *PendingTransactionHistory) SetOffsetAccountNo(v string) {
	o.OffsetAccountNo = &v
}

// GetReferenceId returns the ReferenceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PendingTransactionHistory) GetReferenceId() string {
	if o == nil || o.ReferenceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingTransactionHistory) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// SetReferenceId sets field value
func (o *PendingTransactionHistory) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}

// GetTenant returns the Tenant field value
func (o *PendingTransactionHistory) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *PendingTransactionHistory) SetTenant(v string) {
	o.Tenant = v
}

// GetUpdated returns the Updated field value
func (o *PendingTransactionHistory) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *PendingTransactionHistory) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetUuid returns the Uuid field value
func (o *PendingTransactionHistory) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionHistory) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *PendingTransactionHistory) SetUuid(v string) {
	o.Uuid = v
}

func (o PendingTransactionHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account_no"] = o.AccountNo
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["idemkey"] = o.Idemkey
	}
	if o.OffsetAccountId != nil {
		toSerialize["offset_account_id"] = o.OffsetAccountId
	}
	if o.OffsetAccountNo != nil {
		toSerialize["offset_account_no"] = o.OffsetAccountNo
	}
	if true {
		toSerialize["reference_id"] = o.ReferenceId.Get()
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullablePendingTransactionHistory struct {
	value *PendingTransactionHistory
	isSet bool
}

func (v NullablePendingTransactionHistory) Get() *PendingTransactionHistory {
	return v.value
}

func (v *NullablePendingTransactionHistory) Set(val *PendingTransactionHistory) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingTransactionHistory) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingTransactionHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingTransactionHistory(val *PendingTransactionHistory) *NullablePendingTransactionHistory {
	return &NullablePendingTransactionHistory{value: val, isSet: true}
}

func (v NullablePendingTransactionHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingTransactionHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


