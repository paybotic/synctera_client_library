/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TemplateFieldsGeneralLedger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateFieldsGeneralLedger{}

// TemplateFieldsGeneralLedger In production, general ledger account templates can only be created or updated by a Synctera administrator. General ledger account templates are in Alpha status, and cannot yet be created. We may make breaking changes.
type TemplateFieldsGeneralLedger struct {
	AccountType AccountType `json:"account_type"`
	// The bank account ID for this account. This is a unique identifier for the bank side account that this Synctera account belongs to. This field can be auto filled if only one bank account of the appropriate type exist for the tenant of concern.
	BankAccountId *string `json:"bank_account_id,omitempty"`
	// Bank country of the account. ISO 3166-1 Alpha-2 or Alpha-3 country code.
	BankCountry string `json:"bank_country" validate:"regexp=^[A-Z]{2,3}$"`
	// Account currency. ISO 4217 alphabetic currency code
	Currency string `json:"currency" validate:"regexp=^[A-Z]{3}$"`
}

type _TemplateFieldsGeneralLedger TemplateFieldsGeneralLedger

// NewTemplateFieldsGeneralLedger instantiates a new TemplateFieldsGeneralLedger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateFieldsGeneralLedger(accountType AccountType, bankCountry string, currency string) *TemplateFieldsGeneralLedger {
	this := TemplateFieldsGeneralLedger{}
	this.AccountType = accountType
	this.BankCountry = bankCountry
	this.Currency = currency
	return &this
}

// NewTemplateFieldsGeneralLedgerWithDefaults instantiates a new TemplateFieldsGeneralLedger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateFieldsGeneralLedgerWithDefaults() *TemplateFieldsGeneralLedger {
	this := TemplateFieldsGeneralLedger{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *TemplateFieldsGeneralLedger) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneralLedger) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *TemplateFieldsGeneralLedger) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetBankAccountId returns the BankAccountId field value if set, zero value otherwise.
func (o *TemplateFieldsGeneralLedger) GetBankAccountId() string {
	if o == nil || IsNil(o.BankAccountId) {
		var ret string
		return ret
	}
	return *o.BankAccountId
}

// GetBankAccountIdOk returns a tuple with the BankAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneralLedger) GetBankAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountId) {
		return nil, false
	}
	return o.BankAccountId, true
}

// HasBankAccountId returns a boolean if a field has been set.
func (o *TemplateFieldsGeneralLedger) HasBankAccountId() bool {
	if o != nil && !IsNil(o.BankAccountId) {
		return true
	}

	return false
}

// SetBankAccountId gets a reference to the given string and assigns it to the BankAccountId field.
func (o *TemplateFieldsGeneralLedger) SetBankAccountId(v string) {
	o.BankAccountId = &v
}

// GetBankCountry returns the BankCountry field value
func (o *TemplateFieldsGeneralLedger) GetBankCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCountry
}

// GetBankCountryOk returns a tuple with the BankCountry field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneralLedger) GetBankCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCountry, true
}

// SetBankCountry sets field value
func (o *TemplateFieldsGeneralLedger) SetBankCountry(v string) {
	o.BankCountry = v
}

// GetCurrency returns the Currency field value
func (o *TemplateFieldsGeneralLedger) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneralLedger) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TemplateFieldsGeneralLedger) SetCurrency(v string) {
	o.Currency = v
}

func (o TemplateFieldsGeneralLedger) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateFieldsGeneralLedger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_type"] = o.AccountType
	if !IsNil(o.BankAccountId) {
		toSerialize["bank_account_id"] = o.BankAccountId
	}
	toSerialize["bank_country"] = o.BankCountry
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *TemplateFieldsGeneralLedger) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_type",
		"bank_country",
		"currency",
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

	varTemplateFieldsGeneralLedger := _TemplateFieldsGeneralLedger{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplateFieldsGeneralLedger)

	if err != nil {
		return err
	}

	*o = TemplateFieldsGeneralLedger(varTemplateFieldsGeneralLedger)

	return err
}

type NullableTemplateFieldsGeneralLedger struct {
	value *TemplateFieldsGeneralLedger
	isSet bool
}

func (v NullableTemplateFieldsGeneralLedger) Get() *TemplateFieldsGeneralLedger {
	return v.value
}

func (v *NullableTemplateFieldsGeneralLedger) Set(val *TemplateFieldsGeneralLedger) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFieldsGeneralLedger) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFieldsGeneralLedger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFieldsGeneralLedger(val *TemplateFieldsGeneralLedger) *NullableTemplateFieldsGeneralLedger {
	return &NullableTemplateFieldsGeneralLedger{value: val, isSet: true}
}

func (v NullableTemplateFieldsGeneralLedger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFieldsGeneralLedger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}