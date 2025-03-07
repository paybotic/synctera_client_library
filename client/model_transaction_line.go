/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the TransactionLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionLine{}

// TransactionLine struct for TransactionLine
type TransactionLine struct {
	// The account uuid associated with this transaction line
	AccountId string `json:"account_id"`
	// The account number associated with this transaction line
	AccountNo string `json:"account_no"`
	// The amount (in cents) of the transaction
	Amount int32 `json:"amount"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency             string `json:"currency"`
	DcSign               DcSign `json:"dc_sign"`
	Uuid                 string `json:"uuid"`
	AdditionalProperties map[string]interface{}
}

type _TransactionLine TransactionLine

// NewTransactionLine instantiates a new TransactionLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionLine(accountId string, accountNo string, amount int32, currency string, dcSign DcSign, uuid string) *TransactionLine {
	this := TransactionLine{}
	this.AccountId = accountId
	this.AccountNo = accountNo
	this.Amount = amount
	this.Currency = currency
	this.DcSign = dcSign
	this.Uuid = uuid
	return &this
}

// NewTransactionLineWithDefaults instantiates a new TransactionLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionLineWithDefaults() *TransactionLine {
	this := TransactionLine{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *TransactionLine) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransactionLine) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransactionLine) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountNo returns the AccountNo field value
func (o *TransactionLine) GetAccountNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value
// and a boolean to check if the value has been set.
func (o *TransactionLine) GetAccountNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNo, true
}

// SetAccountNo sets field value
func (o *TransactionLine) SetAccountNo(v string) {
	o.AccountNo = v
}

// GetAmount returns the Amount field value
func (o *TransactionLine) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionLine) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionLine) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *TransactionLine) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionLine) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionLine) SetCurrency(v string) {
	o.Currency = v
}

// GetDcSign returns the DcSign field value
func (o *TransactionLine) GetDcSign() DcSign {
	if o == nil {
		var ret DcSign
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *TransactionLine) GetDcSignOk() (*DcSign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *TransactionLine) SetDcSign(v DcSign) {
	o.DcSign = v
}

// GetUuid returns the Uuid field value
func (o *TransactionLine) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *TransactionLine) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *TransactionLine) SetUuid(v string) {
	o.Uuid = v
}

func (o TransactionLine) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["account_no"] = o.AccountNo
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	toSerialize["dc_sign"] = o.DcSign
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransactionLine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"account_no",
		"amount",
		"currency",
		"dc_sign",
		"uuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTransactionLine := _TransactionLine{}

	err = json.Unmarshal(data, &varTransactionLine)

	if err != nil {
		return err
	}

	*o = TransactionLine(varTransactionLine)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "account_no")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "dc_sign")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionLine struct {
	value *TransactionLine
	isSet bool
}

func (v NullableTransactionLine) Get() *TransactionLine {
	return v.value
}

func (v *NullableTransactionLine) Set(val *TransactionLine) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionLine) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionLine(val *TransactionLine) *NullableTransactionLine {
	return &NullableTransactionLine{value: val, isSet: true}
}

func (v NullableTransactionLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
