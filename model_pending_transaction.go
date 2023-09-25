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

// checks if the PendingTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingTransaction{}

// PendingTransaction struct for PendingTransaction
type PendingTransaction struct {
	// The account id associated with the hold
	AccountId string `json:"account_id"`
	// The account number associated with the hold
	AccountNo string `json:"account_no"`
	// The creation date of the hold
	Created time.Time `json:"created"`
	Data PendingTransactionData `json:"data"`
	Id int64 `json:"id"`
	// The idempotency key used when initially creating this hold.
	Idemkey string `json:"idemkey"`
	// The offset account id associated with the hold
	OffsetAccountId *string `json:"offset_account_id,omitempty"`
	// The offset account number associated with the hold
	OffsetAccountNo *string `json:"offset_account_no,omitempty"`
	// An external ID provided by the payment network to represent this transaction.
	ReferenceId NullableString `json:"reference_id"`
	// The tenant associated with this hold, in the form \"<bankid>_<partnerid>\"
	Tenant string `json:"tenant"`
	// The date the hold was last update
	Updated time.Time `json:"updated"`
	// The unique identifier of the hold transaction.
	Uuid string `json:"uuid"`
}

// NewPendingTransaction instantiates a new PendingTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingTransaction(accountId string, accountNo string, created time.Time, data PendingTransactionData, id int64, idemkey string, referenceId NullableString, tenant string, updated time.Time, uuid string) *PendingTransaction {
	this := PendingTransaction{}
	this.AccountId = accountId
	this.AccountNo = accountNo
	this.Created = created
	this.Data = data
	this.Id = id
	this.Idemkey = idemkey
	this.ReferenceId = referenceId
	this.Tenant = tenant
	this.Updated = updated
	this.Uuid = uuid
	return &this
}

// NewPendingTransactionWithDefaults instantiates a new PendingTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingTransactionWithDefaults() *PendingTransaction {
	this := PendingTransaction{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *PendingTransaction) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *PendingTransaction) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountNo returns the AccountNo field value
func (o *PendingTransaction) GetAccountNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetAccountNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNo, true
}

// SetAccountNo sets field value
func (o *PendingTransaction) SetAccountNo(v string) {
	o.AccountNo = v
}

// GetCreated returns the Created field value
func (o *PendingTransaction) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PendingTransaction) SetCreated(v time.Time) {
	o.Created = v
}

// GetData returns the Data field value
func (o *PendingTransaction) GetData() PendingTransactionData {
	if o == nil {
		var ret PendingTransactionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetDataOk() (*PendingTransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PendingTransaction) SetData(v PendingTransactionData) {
	o.Data = v
}

// GetId returns the Id field value
func (o *PendingTransaction) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PendingTransaction) SetId(v int64) {
	o.Id = v
}

// GetIdemkey returns the Idemkey field value
func (o *PendingTransaction) GetIdemkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Idemkey
}

// GetIdemkeyOk returns a tuple with the Idemkey field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetIdemkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idemkey, true
}

// SetIdemkey sets field value
func (o *PendingTransaction) SetIdemkey(v string) {
	o.Idemkey = v
}

// GetOffsetAccountId returns the OffsetAccountId field value if set, zero value otherwise.
func (o *PendingTransaction) GetOffsetAccountId() string {
	if o == nil || IsNil(o.OffsetAccountId) {
		var ret string
		return ret
	}
	return *o.OffsetAccountId
}

// GetOffsetAccountIdOk returns a tuple with the OffsetAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetOffsetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.OffsetAccountId) {
		return nil, false
	}
	return o.OffsetAccountId, true
}

// HasOffsetAccountId returns a boolean if a field has been set.
func (o *PendingTransaction) HasOffsetAccountId() bool {
	if o != nil && !IsNil(o.OffsetAccountId) {
		return true
	}

	return false
}

// SetOffsetAccountId gets a reference to the given string and assigns it to the OffsetAccountId field.
func (o *PendingTransaction) SetOffsetAccountId(v string) {
	o.OffsetAccountId = &v
}

// GetOffsetAccountNo returns the OffsetAccountNo field value if set, zero value otherwise.
func (o *PendingTransaction) GetOffsetAccountNo() string {
	if o == nil || IsNil(o.OffsetAccountNo) {
		var ret string
		return ret
	}
	return *o.OffsetAccountNo
}

// GetOffsetAccountNoOk returns a tuple with the OffsetAccountNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetOffsetAccountNoOk() (*string, bool) {
	if o == nil || IsNil(o.OffsetAccountNo) {
		return nil, false
	}
	return o.OffsetAccountNo, true
}

// HasOffsetAccountNo returns a boolean if a field has been set.
func (o *PendingTransaction) HasOffsetAccountNo() bool {
	if o != nil && !IsNil(o.OffsetAccountNo) {
		return true
	}

	return false
}

// SetOffsetAccountNo gets a reference to the given string and assigns it to the OffsetAccountNo field.
func (o *PendingTransaction) SetOffsetAccountNo(v string) {
	o.OffsetAccountNo = &v
}

// GetReferenceId returns the ReferenceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PendingTransaction) GetReferenceId() string {
	if o == nil || o.ReferenceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingTransaction) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// SetReferenceId sets field value
func (o *PendingTransaction) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}

// GetTenant returns the Tenant field value
func (o *PendingTransaction) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *PendingTransaction) SetTenant(v string) {
	o.Tenant = v
}

// GetUpdated returns the Updated field value
func (o *PendingTransaction) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *PendingTransaction) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetUuid returns the Uuid field value
func (o *PendingTransaction) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *PendingTransaction) SetUuid(v string) {
	o.Uuid = v
}

func (o PendingTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["account_no"] = o.AccountNo
	toSerialize["created"] = o.Created
	toSerialize["data"] = o.Data
	toSerialize["id"] = o.Id
	toSerialize["idemkey"] = o.Idemkey
	if !IsNil(o.OffsetAccountId) {
		toSerialize["offset_account_id"] = o.OffsetAccountId
	}
	if !IsNil(o.OffsetAccountNo) {
		toSerialize["offset_account_no"] = o.OffsetAccountNo
	}
	toSerialize["reference_id"] = o.ReferenceId.Get()
	toSerialize["tenant"] = o.Tenant
	toSerialize["updated"] = o.Updated
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

type NullablePendingTransaction struct {
	value *PendingTransaction
	isSet bool
}

func (v NullablePendingTransaction) Get() *PendingTransaction {
	return v.value
}

func (v *NullablePendingTransaction) Set(val *PendingTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingTransaction(val *PendingTransaction) *NullablePendingTransaction {
	return &NullablePendingTransaction{value: val, isSet: true}
}

func (v NullablePendingTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


