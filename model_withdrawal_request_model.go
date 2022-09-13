/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WithdrawalRequestModel struct for WithdrawalRequestModel
type WithdrawalRequestModel struct {
	AccountType *string `json:"account_type,omitempty"`
	// The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents)
	Amount int32 `json:"amount"`
	CardAcceptor *CardAcceptorModel `json:"card_acceptor,omitempty"`
	CardId string `json:"card_id"`
	Mid string `json:"mid"`
	Pin *string `json:"pin,omitempty"`
}

// NewWithdrawalRequestModel instantiates a new WithdrawalRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawalRequestModel(amount int32, cardId string, mid string) *WithdrawalRequestModel {
	this := WithdrawalRequestModel{}
	this.Amount = amount
	this.CardId = cardId
	this.Mid = mid
	return &this
}

// NewWithdrawalRequestModelWithDefaults instantiates a new WithdrawalRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawalRequestModelWithDefaults() *WithdrawalRequestModel {
	this := WithdrawalRequestModel{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *WithdrawalRequestModel) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *WithdrawalRequestModel) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *WithdrawalRequestModel) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAmount returns the Amount field value
func (o *WithdrawalRequestModel) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *WithdrawalRequestModel) SetAmount(v int32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *WithdrawalRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil || o.CardAcceptor == nil {
		var ret CardAcceptorModel
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil || o.CardAcceptor == nil {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *WithdrawalRequestModel) HasCardAcceptor() bool {
	if o != nil && o.CardAcceptor != nil {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given CardAcceptorModel and assigns it to the CardAcceptor field.
func (o *WithdrawalRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = &v
}

// GetCardId returns the CardId field value
func (o *WithdrawalRequestModel) GetCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardId, true
}

// SetCardId sets field value
func (o *WithdrawalRequestModel) SetCardId(v string) {
	o.CardId = v
}

// GetMid returns the Mid field value
func (o *WithdrawalRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *WithdrawalRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *WithdrawalRequestModel) GetPin() string {
	if o == nil || o.Pin == nil {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetPinOk() (*string, bool) {
	if o == nil || o.Pin == nil {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *WithdrawalRequestModel) HasPin() bool {
	if o != nil && o.Pin != nil {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *WithdrawalRequestModel) SetPin(v string) {
	o.Pin = &v
}

func (o WithdrawalRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountType != nil {
		toSerialize["account_type"] = o.AccountType
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.CardAcceptor != nil {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	if true {
		toSerialize["card_id"] = o.CardId
	}
	if true {
		toSerialize["mid"] = o.Mid
	}
	if o.Pin != nil {
		toSerialize["pin"] = o.Pin
	}
	return json.Marshal(toSerialize)
}

type NullableWithdrawalRequestModel struct {
	value *WithdrawalRequestModel
	isSet bool
}

func (v NullableWithdrawalRequestModel) Get() *WithdrawalRequestModel {
	return v.value
}

func (v *NullableWithdrawalRequestModel) Set(val *WithdrawalRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawalRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawalRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawalRequestModel(val *WithdrawalRequestModel) *NullableWithdrawalRequestModel {
	return &NullableWithdrawalRequestModel{value: val, isSet: true}
}

func (v NullableWithdrawalRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawalRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


