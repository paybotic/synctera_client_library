/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CardTransactionData Transaction metadata specific to transactions with type `CARD`
type CardTransactionData struct {
	// Debit Network ID
	CardId             *string                                `json:"card_id,omitempty"`
	CurrencyCode       *string                                `json:"currency_code,omitempty"`
	CurrencyConversion *CardTransactionDataCurrencyConversion `json:"currency_conversion,omitempty"`
	Merchant           *CardTransactionDataMerchant           `json:"merchant,omitempty"`
	// The network used for the transaction
	Network *string `json:"network,omitempty"`
	// The ID provided by he network to represent this transaction
	NetworkReferenceId *string                 `json:"network_reference_id,omitempty"`
	Pos                *CardTransactionDataPos `json:"pos,omitempty"`
}

// NewCardTransactionData instantiates a new CardTransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardTransactionData() *CardTransactionData {
	this := CardTransactionData{}
	return &this
}

// NewCardTransactionDataWithDefaults instantiates a new CardTransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardTransactionDataWithDefaults() *CardTransactionData {
	this := CardTransactionData{}
	return &this
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *CardTransactionData) GetCardId() string {
	if o == nil || o.CardId == nil {
		var ret string
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionData) GetCardIdOk() (*string, bool) {
	if o == nil || o.CardId == nil {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *CardTransactionData) HasCardId() bool {
	if o != nil && o.CardId != nil {
		return true
	}

	return false
}

// SetCardId gets a reference to the given string and assigns it to the CardId field.
func (o *CardTransactionData) SetCardId(v string) {
	o.CardId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *CardTransactionData) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionData) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *CardTransactionData) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *CardTransactionData) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencyConversion returns the CurrencyConversion field value if set, zero value otherwise.
func (o *CardTransactionData) GetCurrencyConversion() CardTransactionDataCurrencyConversion {
	if o == nil || o.CurrencyConversion == nil {
		var ret CardTransactionDataCurrencyConversion
		return ret
	}
	return *o.CurrencyConversion
}

// GetCurrencyConversionOk returns a tuple with the CurrencyConversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionData) GetCurrencyConversionOk() (*CardTransactionDataCurrencyConversion, bool) {
	if o == nil || o.CurrencyConversion == nil {
		return nil, false
	}
	return o.CurrencyConversion, true
}

// HasCurrencyConversion returns a boolean if a field has been set.
func (o *CardTransactionData) HasCurrencyConversion() bool {
	if o != nil && o.CurrencyConversion != nil {
		return true
	}

	return false
}

// SetCurrencyConversion gets a reference to the given CardTransactionDataCurrencyConversion and assigns it to the CurrencyConversion field.
func (o *CardTransactionData) SetCurrencyConversion(v CardTransactionDataCurrencyConversion) {
	o.CurrencyConversion = &v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *CardTransactionData) GetMerchant() CardTransactionDataMerchant {
	if o == nil || o.Merchant == nil {
		var ret CardTransactionDataMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionData) GetMerchantOk() (*CardTransactionDataMerchant, bool) {
	if o == nil || o.Merchant == nil {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *CardTransactionData) HasMerchant() bool {
	if o != nil && o.Merchant != nil {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given CardTransactionDataMerchant and assigns it to the Merchant field.
func (o *CardTransactionData) SetMerchant(v CardTransactionDataMerchant) {
	o.Merchant = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CardTransactionData) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionData) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CardTransactionData) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *CardTransactionData) SetNetwork(v string) {
	o.Network = &v
}

// GetNetworkReferenceId returns the NetworkReferenceId field value if set, zero value otherwise.
func (o *CardTransactionData) GetNetworkReferenceId() string {
	if o == nil || o.NetworkReferenceId == nil {
		var ret string
		return ret
	}
	return *o.NetworkReferenceId
}

// GetNetworkReferenceIdOk returns a tuple with the NetworkReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionData) GetNetworkReferenceIdOk() (*string, bool) {
	if o == nil || o.NetworkReferenceId == nil {
		return nil, false
	}
	return o.NetworkReferenceId, true
}

// HasNetworkReferenceId returns a boolean if a field has been set.
func (o *CardTransactionData) HasNetworkReferenceId() bool {
	if o != nil && o.NetworkReferenceId != nil {
		return true
	}

	return false
}

// SetNetworkReferenceId gets a reference to the given string and assigns it to the NetworkReferenceId field.
func (o *CardTransactionData) SetNetworkReferenceId(v string) {
	o.NetworkReferenceId = &v
}

// GetPos returns the Pos field value if set, zero value otherwise.
func (o *CardTransactionData) GetPos() CardTransactionDataPos {
	if o == nil || o.Pos == nil {
		var ret CardTransactionDataPos
		return ret
	}
	return *o.Pos
}

// GetPosOk returns a tuple with the Pos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionData) GetPosOk() (*CardTransactionDataPos, bool) {
	if o == nil || o.Pos == nil {
		return nil, false
	}
	return o.Pos, true
}

// HasPos returns a boolean if a field has been set.
func (o *CardTransactionData) HasPos() bool {
	if o != nil && o.Pos != nil {
		return true
	}

	return false
}

// SetPos gets a reference to the given CardTransactionDataPos and assigns it to the Pos field.
func (o *CardTransactionData) SetPos(v CardTransactionDataPos) {
	o.Pos = &v
}

func (o CardTransactionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CardId != nil {
		toSerialize["card_id"] = o.CardId
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.CurrencyConversion != nil {
		toSerialize["currency_conversion"] = o.CurrencyConversion
	}
	if o.Merchant != nil {
		toSerialize["merchant"] = o.Merchant
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.NetworkReferenceId != nil {
		toSerialize["network_reference_id"] = o.NetworkReferenceId
	}
	if o.Pos != nil {
		toSerialize["pos"] = o.Pos
	}
	return json.Marshal(toSerialize)
}

type NullableCardTransactionData struct {
	value *CardTransactionData
	isSet bool
}

func (v NullableCardTransactionData) Get() *CardTransactionData {
	return v.value
}

func (v *NullableCardTransactionData) Set(val *CardTransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullableCardTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullableCardTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardTransactionData(val *CardTransactionData) *NullableCardTransactionData {
	return &NullableCardTransactionData{value: val, isSet: true}
}

func (v NullableCardTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
