/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TransferRequestPush type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferRequestPush{}

// TransferRequestPush struct for TransferRequestPush
type TransferRequestPush struct {
	// Amount of the transfer in cents
	Amount int32 `json:"amount"`
	// ISO 4217  Alpha-3 currency code
	Currency string `json:"currency"`
	// The ID of the external card from/to which the transfer will be initiated/received
	ExternalCardId string    `json:"external_card_id"`
	Merchant       *Merchant `json:"merchant,omitempty"`
	// The ID of the account to which the transfer will be initiated/received
	OriginatingAccountId string              `json:"originating_account_id"`
	Type                 TransferTypeRequest `json:"type"`
	// For person-to-person PUSH transactions, this is the `customer_id` of the sender who must have privileges to access funds in the originating account in order to send funds to the recipient cardholder
	OriginatingCustomerId *string `json:"originating_customer_id,omitempty"`
}

type _TransferRequestPush TransferRequestPush

// NewTransferRequestPush instantiates a new TransferRequestPush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRequestPush(amount int32, currency string, externalCardId string, originatingAccountId string, type_ TransferTypeRequest) *TransferRequestPush {
	this := TransferRequestPush{}
	this.Amount = amount
	this.Currency = currency
	this.ExternalCardId = externalCardId
	this.OriginatingAccountId = originatingAccountId
	this.Type = type_
	return &this
}

// NewTransferRequestPushWithDefaults instantiates a new TransferRequestPush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRequestPushWithDefaults() *TransferRequestPush {
	this := TransferRequestPush{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TransferRequestPush) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferRequestPush) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferRequestPush) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *TransferRequestPush) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransferRequestPush) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransferRequestPush) SetCurrency(v string) {
	o.Currency = v
}

// GetExternalCardId returns the ExternalCardId field value
func (o *TransferRequestPush) GetExternalCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalCardId
}

// GetExternalCardIdOk returns a tuple with the ExternalCardId field value
// and a boolean to check if the value has been set.
func (o *TransferRequestPush) GetExternalCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalCardId, true
}

// SetExternalCardId sets field value
func (o *TransferRequestPush) SetExternalCardId(v string) {
	o.ExternalCardId = v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *TransferRequestPush) GetMerchant() Merchant {
	if o == nil || IsNil(o.Merchant) {
		var ret Merchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequestPush) GetMerchantOk() (*Merchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *TransferRequestPush) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given Merchant and assigns it to the Merchant field.
func (o *TransferRequestPush) SetMerchant(v Merchant) {
	o.Merchant = &v
}

// GetOriginatingAccountId returns the OriginatingAccountId field value
func (o *TransferRequestPush) GetOriginatingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatingAccountId
}

// GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field value
// and a boolean to check if the value has been set.
func (o *TransferRequestPush) GetOriginatingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatingAccountId, true
}

// SetOriginatingAccountId sets field value
func (o *TransferRequestPush) SetOriginatingAccountId(v string) {
	o.OriginatingAccountId = v
}

// GetType returns the Type field value
func (o *TransferRequestPush) GetType() TransferTypeRequest {
	if o == nil {
		var ret TransferTypeRequest
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferRequestPush) GetTypeOk() (*TransferTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferRequestPush) SetType(v TransferTypeRequest) {
	o.Type = v
}

// GetOriginatingCustomerId returns the OriginatingCustomerId field value if set, zero value otherwise.
func (o *TransferRequestPush) GetOriginatingCustomerId() string {
	if o == nil || IsNil(o.OriginatingCustomerId) {
		var ret string
		return ret
	}
	return *o.OriginatingCustomerId
}

// GetOriginatingCustomerIdOk returns a tuple with the OriginatingCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequestPush) GetOriginatingCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingCustomerId) {
		return nil, false
	}
	return o.OriginatingCustomerId, true
}

// HasOriginatingCustomerId returns a boolean if a field has been set.
func (o *TransferRequestPush) HasOriginatingCustomerId() bool {
	if o != nil && !IsNil(o.OriginatingCustomerId) {
		return true
	}

	return false
}

// SetOriginatingCustomerId gets a reference to the given string and assigns it to the OriginatingCustomerId field.
func (o *TransferRequestPush) SetOriginatingCustomerId(v string) {
	o.OriginatingCustomerId = &v
}

func (o TransferRequestPush) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferRequestPush) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	toSerialize["external_card_id"] = o.ExternalCardId
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	toSerialize["originating_account_id"] = o.OriginatingAccountId
	toSerialize["type"] = o.Type
	if !IsNil(o.OriginatingCustomerId) {
		toSerialize["originating_customer_id"] = o.OriginatingCustomerId
	}
	return toSerialize, nil
}

func (o *TransferRequestPush) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"external_card_id",
		"originating_account_id",
		"type",
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

	varTransferRequestPush := _TransferRequestPush{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransferRequestPush)

	if err != nil {
		return err
	}

	*o = TransferRequestPush(varTransferRequestPush)

	return err
}

type NullableTransferRequestPush struct {
	value *TransferRequestPush
	isSet bool
}

func (v NullableTransferRequestPush) Get() *TransferRequestPush {
	return v.value
}

func (v *NullableTransferRequestPush) Set(val *TransferRequestPush) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRequestPush) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRequestPush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRequestPush(val *TransferRequestPush) *NullableTransferRequestPush {
	return &NullableTransferRequestPush{value: val, isSet: true}
}

func (v NullableTransferRequestPush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRequestPush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
