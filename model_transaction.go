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

// Transaction struct for Transaction
type Transaction struct {
	Data TransactionData `json:"data"`
	// The \"effective date\" of a transaction. This may be earlier than posted_date in some cases (for example, a transaction that occurs on a Saturday may not be posted until the following Monday, but would have an effective date of Saturday)
	EffectiveDate time.Time `json:"effective_date"`
	Id int32 `json:"id"`
	// The date the transaction was posted. This is the date any money is considered to be added or removed from an account.
	PostedDate time.Time `json:"posted_date"`
	Status string `json:"status"`
	// The specific transaction type. For example, for `ach`, this may be \"outgoing_debit\".
	Subtype string `json:"subtype"`
	// The general type of transaction. For example, \"card\" or \"ach\".
	Type string `json:"type"`
	// The unique identifier of the transaction.
	Uuid string `json:"uuid"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(data TransactionData, effectiveDate time.Time, id int32, postedDate time.Time, status string, subtype string, type_ string, uuid string) *Transaction {
	this := Transaction{}
	this.Data = data
	this.EffectiveDate = effectiveDate
	this.Id = id
	this.PostedDate = postedDate
	this.Status = status
	this.Subtype = subtype
	this.Type = type_
	this.Uuid = uuid
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetData returns the Data field value
func (o *Transaction) GetData() TransactionData {
	if o == nil {
		var ret TransactionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetDataOk() (*TransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Transaction) SetData(v TransactionData) {
	o.Data = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *Transaction) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *Transaction) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetId returns the Id field value
func (o *Transaction) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transaction) SetId(v int32) {
	o.Id = v
}

// GetPostedDate returns the PostedDate field value
func (o *Transaction) GetPostedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetPostedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostedDate, true
}

// SetPostedDate sets field value
func (o *Transaction) SetPostedDate(v time.Time) {
	o.PostedDate = v
}

// GetStatus returns the Status field value
func (o *Transaction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transaction) SetStatus(v string) {
	o.Status = v
}

// GetSubtype returns the Subtype field value
func (o *Transaction) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *Transaction) SetSubtype(v string) {
	o.Subtype = v
}

// GetType returns the Type field value
func (o *Transaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Transaction) SetType(v string) {
	o.Type = v
}

// GetUuid returns the Uuid field value
func (o *Transaction) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Transaction) SetUuid(v string) {
	o.Uuid = v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["effective_date"] = o.EffectiveDate
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["posted_date"] = o.PostedDate
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


