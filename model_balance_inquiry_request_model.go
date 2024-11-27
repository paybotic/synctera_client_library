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

// checks if the BalanceInquiryRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceInquiryRequestModel{}

// BalanceInquiryRequestModel struct for BalanceInquiryRequestModel
type BalanceInquiryRequestModel struct {
	AccountType  string            `json:"account_type"`
	CardAcceptor CardAcceptorModel `json:"card_acceptor"`
	CardId       string            `json:"card_id"`
	Mid          string            `json:"mid"`
	NetworkFees  []NetworkFeeModel `json:"network_fees,omitempty"`
	Pin          *string           `json:"pin,omitempty"`
}

type _BalanceInquiryRequestModel BalanceInquiryRequestModel

// NewBalanceInquiryRequestModel instantiates a new BalanceInquiryRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceInquiryRequestModel(accountType string, cardAcceptor CardAcceptorModel, cardId string, mid string) *BalanceInquiryRequestModel {
	this := BalanceInquiryRequestModel{}
	this.AccountType = accountType
	this.CardAcceptor = cardAcceptor
	this.CardId = cardId
	this.Mid = mid
	return &this
}

// NewBalanceInquiryRequestModelWithDefaults instantiates a new BalanceInquiryRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceInquiryRequestModelWithDefaults() *BalanceInquiryRequestModel {
	this := BalanceInquiryRequestModel{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *BalanceInquiryRequestModel) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *BalanceInquiryRequestModel) SetAccountType(v string) {
	o.AccountType = v
}

// GetCardAcceptor returns the CardAcceptor field value
func (o *BalanceInquiryRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil {
		var ret CardAcceptorModel
		return ret
	}

	return o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardAcceptor, true
}

// SetCardAcceptor sets field value
func (o *BalanceInquiryRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = v
}

// GetCardId returns the CardId field value
func (o *BalanceInquiryRequestModel) GetCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardId, true
}

// SetCardId sets field value
func (o *BalanceInquiryRequestModel) SetCardId(v string) {
	o.CardId = v
}

// GetMid returns the Mid field value
func (o *BalanceInquiryRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *BalanceInquiryRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetNetworkFees returns the NetworkFees field value if set, zero value otherwise.
func (o *BalanceInquiryRequestModel) GetNetworkFees() []NetworkFeeModel {
	if o == nil || IsNil(o.NetworkFees) {
		var ret []NetworkFeeModel
		return ret
	}
	return o.NetworkFees
}

// GetNetworkFeesOk returns a tuple with the NetworkFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetNetworkFeesOk() ([]NetworkFeeModel, bool) {
	if o == nil || IsNil(o.NetworkFees) {
		return nil, false
	}
	return o.NetworkFees, true
}

// HasNetworkFees returns a boolean if a field has been set.
func (o *BalanceInquiryRequestModel) HasNetworkFees() bool {
	if o != nil && !IsNil(o.NetworkFees) {
		return true
	}

	return false
}

// SetNetworkFees gets a reference to the given []NetworkFeeModel and assigns it to the NetworkFees field.
func (o *BalanceInquiryRequestModel) SetNetworkFees(v []NetworkFeeModel) {
	o.NetworkFees = v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *BalanceInquiryRequestModel) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *BalanceInquiryRequestModel) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *BalanceInquiryRequestModel) SetPin(v string) {
	o.Pin = &v
}

func (o BalanceInquiryRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceInquiryRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_type"] = o.AccountType
	toSerialize["card_acceptor"] = o.CardAcceptor
	toSerialize["card_id"] = o.CardId
	toSerialize["mid"] = o.Mid
	if !IsNil(o.NetworkFees) {
		toSerialize["network_fees"] = o.NetworkFees
	}
	if !IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	return toSerialize, nil
}

func (o *BalanceInquiryRequestModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_type",
		"card_acceptor",
		"card_id",
		"mid",
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

	varBalanceInquiryRequestModel := _BalanceInquiryRequestModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBalanceInquiryRequestModel)

	if err != nil {
		return err
	}

	*o = BalanceInquiryRequestModel(varBalanceInquiryRequestModel)

	return err
}

type NullableBalanceInquiryRequestModel struct {
	value *BalanceInquiryRequestModel
	isSet bool
}

func (v NullableBalanceInquiryRequestModel) Get() *BalanceInquiryRequestModel {
	return v.value
}

func (v *NullableBalanceInquiryRequestModel) Set(val *BalanceInquiryRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceInquiryRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceInquiryRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceInquiryRequestModel(val *BalanceInquiryRequestModel) *NullableBalanceInquiryRequestModel {
	return &NullableBalanceInquiryRequestModel{value: val, isSet: true}
}

func (v NullableBalanceInquiryRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceInquiryRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
