/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AuthorizationAdviceModel struct for AuthorizationAdviceModel
type AuthorizationAdviceModel struct {
	// The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents)
	Amount int32 `json:"amount"`
	NetworkFees []NetworkFeeModel `json:"network_fees,omitempty"`
	OriginalTransactionId string `json:"original_transaction_id"`
	TransactionOptions *TransactionOptions `json:"transaction_options,omitempty"`
}

// NewAuthorizationAdviceModel instantiates a new AuthorizationAdviceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationAdviceModel(amount int32, originalTransactionId string) *AuthorizationAdviceModel {
	this := AuthorizationAdviceModel{}
	this.Amount = amount
	this.OriginalTransactionId = originalTransactionId
	return &this
}

// NewAuthorizationAdviceModelWithDefaults instantiates a new AuthorizationAdviceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationAdviceModelWithDefaults() *AuthorizationAdviceModel {
	this := AuthorizationAdviceModel{}
	return &this
}

// GetAmount returns the Amount field value
func (o *AuthorizationAdviceModel) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AuthorizationAdviceModel) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AuthorizationAdviceModel) SetAmount(v int32) {
	o.Amount = v
}

// GetNetworkFees returns the NetworkFees field value if set, zero value otherwise.
func (o *AuthorizationAdviceModel) GetNetworkFees() []NetworkFeeModel {
	if o == nil || o.NetworkFees == nil {
		var ret []NetworkFeeModel
		return ret
	}
	return o.NetworkFees
}

// GetNetworkFeesOk returns a tuple with the NetworkFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationAdviceModel) GetNetworkFeesOk() ([]NetworkFeeModel, bool) {
	if o == nil || o.NetworkFees == nil {
		return nil, false
	}
	return o.NetworkFees, true
}

// HasNetworkFees returns a boolean if a field has been set.
func (o *AuthorizationAdviceModel) HasNetworkFees() bool {
	if o != nil && o.NetworkFees != nil {
		return true
	}

	return false
}

// SetNetworkFees gets a reference to the given []NetworkFeeModel and assigns it to the NetworkFees field.
func (o *AuthorizationAdviceModel) SetNetworkFees(v []NetworkFeeModel) {
	o.NetworkFees = v
}

// GetOriginalTransactionId returns the OriginalTransactionId field value
func (o *AuthorizationAdviceModel) GetOriginalTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalTransactionId
}

// GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field value
// and a boolean to check if the value has been set.
func (o *AuthorizationAdviceModel) GetOriginalTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalTransactionId, true
}

// SetOriginalTransactionId sets field value
func (o *AuthorizationAdviceModel) SetOriginalTransactionId(v string) {
	o.OriginalTransactionId = v
}

// GetTransactionOptions returns the TransactionOptions field value if set, zero value otherwise.
func (o *AuthorizationAdviceModel) GetTransactionOptions() TransactionOptions {
	if o == nil || o.TransactionOptions == nil {
		var ret TransactionOptions
		return ret
	}
	return *o.TransactionOptions
}

// GetTransactionOptionsOk returns a tuple with the TransactionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationAdviceModel) GetTransactionOptionsOk() (*TransactionOptions, bool) {
	if o == nil || o.TransactionOptions == nil {
		return nil, false
	}
	return o.TransactionOptions, true
}

// HasTransactionOptions returns a boolean if a field has been set.
func (o *AuthorizationAdviceModel) HasTransactionOptions() bool {
	if o != nil && o.TransactionOptions != nil {
		return true
	}

	return false
}

// SetTransactionOptions gets a reference to the given TransactionOptions and assigns it to the TransactionOptions field.
func (o *AuthorizationAdviceModel) SetTransactionOptions(v TransactionOptions) {
	o.TransactionOptions = &v
}

func (o AuthorizationAdviceModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.NetworkFees != nil {
		toSerialize["network_fees"] = o.NetworkFees
	}
	if true {
		toSerialize["original_transaction_id"] = o.OriginalTransactionId
	}
	if o.TransactionOptions != nil {
		toSerialize["transaction_options"] = o.TransactionOptions
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationAdviceModel struct {
	value *AuthorizationAdviceModel
	isSet bool
}

func (v NullableAuthorizationAdviceModel) Get() *AuthorizationAdviceModel {
	return v.value
}

func (v *NullableAuthorizationAdviceModel) Set(val *AuthorizationAdviceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationAdviceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationAdviceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationAdviceModel(val *AuthorizationAdviceModel) *NullableAuthorizationAdviceModel {
	return &NullableAuthorizationAdviceModel{value: val, isSet: true}
}

func (v NullableAuthorizationAdviceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationAdviceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


