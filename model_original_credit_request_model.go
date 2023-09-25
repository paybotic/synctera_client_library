/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OriginalCreditRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginalCreditRequestModel{}

// OriginalCreditRequestModel struct for OriginalCreditRequestModel
type OriginalCreditRequestModel struct {
	// The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents)
	Amount int32 `json:"amount"`
	CardAcceptor *CardAcceptorModel `json:"card_acceptor,omitempty"`
	CardId string `json:"card_id"`
	Mid string `json:"mid"`
	ScreeningScore *string `json:"screening_score,omitempty"`
	SenderData *OriginalCreditSenderData `json:"sender_data,omitempty"`
	TransactionPurpose *string `json:"transactionPurpose,omitempty"`
	Type string `json:"type"`
}

// NewOriginalCreditRequestModel instantiates a new OriginalCreditRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginalCreditRequestModel(amount int32, cardId string, mid string, type_ string) *OriginalCreditRequestModel {
	this := OriginalCreditRequestModel{}
	this.Amount = amount
	this.CardId = cardId
	this.Mid = mid
	this.Type = type_
	return &this
}

// NewOriginalCreditRequestModelWithDefaults instantiates a new OriginalCreditRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginalCreditRequestModelWithDefaults() *OriginalCreditRequestModel {
	this := OriginalCreditRequestModel{}
	return &this
}

// GetAmount returns the Amount field value
func (o *OriginalCreditRequestModel) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OriginalCreditRequestModel) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OriginalCreditRequestModel) SetAmount(v int32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *OriginalCreditRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil || IsNil(o.CardAcceptor) {
		var ret CardAcceptorModel
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil || IsNil(o.CardAcceptor) {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *OriginalCreditRequestModel) HasCardAcceptor() bool {
	if o != nil && !IsNil(o.CardAcceptor) {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given CardAcceptorModel and assigns it to the CardAcceptor field.
func (o *OriginalCreditRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = &v
}

// GetCardId returns the CardId field value
func (o *OriginalCreditRequestModel) GetCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value
// and a boolean to check if the value has been set.
func (o *OriginalCreditRequestModel) GetCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardId, true
}

// SetCardId sets field value
func (o *OriginalCreditRequestModel) SetCardId(v string) {
	o.CardId = v
}

// GetMid returns the Mid field value
func (o *OriginalCreditRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *OriginalCreditRequestModel) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *OriginalCreditRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetScreeningScore returns the ScreeningScore field value if set, zero value otherwise.
func (o *OriginalCreditRequestModel) GetScreeningScore() string {
	if o == nil || IsNil(o.ScreeningScore) {
		var ret string
		return ret
	}
	return *o.ScreeningScore
}

// GetScreeningScoreOk returns a tuple with the ScreeningScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditRequestModel) GetScreeningScoreOk() (*string, bool) {
	if o == nil || IsNil(o.ScreeningScore) {
		return nil, false
	}
	return o.ScreeningScore, true
}

// HasScreeningScore returns a boolean if a field has been set.
func (o *OriginalCreditRequestModel) HasScreeningScore() bool {
	if o != nil && !IsNil(o.ScreeningScore) {
		return true
	}

	return false
}

// SetScreeningScore gets a reference to the given string and assigns it to the ScreeningScore field.
func (o *OriginalCreditRequestModel) SetScreeningScore(v string) {
	o.ScreeningScore = &v
}

// GetSenderData returns the SenderData field value if set, zero value otherwise.
func (o *OriginalCreditRequestModel) GetSenderData() OriginalCreditSenderData {
	if o == nil || IsNil(o.SenderData) {
		var ret OriginalCreditSenderData
		return ret
	}
	return *o.SenderData
}

// GetSenderDataOk returns a tuple with the SenderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditRequestModel) GetSenderDataOk() (*OriginalCreditSenderData, bool) {
	if o == nil || IsNil(o.SenderData) {
		return nil, false
	}
	return o.SenderData, true
}

// HasSenderData returns a boolean if a field has been set.
func (o *OriginalCreditRequestModel) HasSenderData() bool {
	if o != nil && !IsNil(o.SenderData) {
		return true
	}

	return false
}

// SetSenderData gets a reference to the given OriginalCreditSenderData and assigns it to the SenderData field.
func (o *OriginalCreditRequestModel) SetSenderData(v OriginalCreditSenderData) {
	o.SenderData = &v
}

// GetTransactionPurpose returns the TransactionPurpose field value if set, zero value otherwise.
func (o *OriginalCreditRequestModel) GetTransactionPurpose() string {
	if o == nil || IsNil(o.TransactionPurpose) {
		var ret string
		return ret
	}
	return *o.TransactionPurpose
}

// GetTransactionPurposeOk returns a tuple with the TransactionPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditRequestModel) GetTransactionPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionPurpose) {
		return nil, false
	}
	return o.TransactionPurpose, true
}

// HasTransactionPurpose returns a boolean if a field has been set.
func (o *OriginalCreditRequestModel) HasTransactionPurpose() bool {
	if o != nil && !IsNil(o.TransactionPurpose) {
		return true
	}

	return false
}

// SetTransactionPurpose gets a reference to the given string and assigns it to the TransactionPurpose field.
func (o *OriginalCreditRequestModel) SetTransactionPurpose(v string) {
	o.TransactionPurpose = &v
}

// GetType returns the Type field value
func (o *OriginalCreditRequestModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OriginalCreditRequestModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OriginalCreditRequestModel) SetType(v string) {
	o.Type = v
}

func (o OriginalCreditRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginalCreditRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CardAcceptor) {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	toSerialize["card_id"] = o.CardId
	toSerialize["mid"] = o.Mid
	if !IsNil(o.ScreeningScore) {
		toSerialize["screening_score"] = o.ScreeningScore
	}
	if !IsNil(o.SenderData) {
		toSerialize["sender_data"] = o.SenderData
	}
	if !IsNil(o.TransactionPurpose) {
		toSerialize["transactionPurpose"] = o.TransactionPurpose
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableOriginalCreditRequestModel struct {
	value *OriginalCreditRequestModel
	isSet bool
}

func (v NullableOriginalCreditRequestModel) Get() *OriginalCreditRequestModel {
	return v.value
}

func (v *NullableOriginalCreditRequestModel) Set(val *OriginalCreditRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginalCreditRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginalCreditRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginalCreditRequestModel(val *OriginalCreditRequestModel) *NullableOriginalCreditRequestModel {
	return &NullableOriginalCreditRequestModel{value: val, isSet: true}
}

func (v NullableOriginalCreditRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginalCreditRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


