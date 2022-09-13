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

// ClearingModel struct for ClearingModel
type ClearingModel struct {
	// The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents)
	Amount int32 `json:"amount"`
	CardAcceptor *CardAcceptorModel `json:"card_acceptor,omitempty"`
	ForcePost *bool `json:"force_post,omitempty"`
	IsRefund *bool `json:"is_refund,omitempty"`
	Mid *string `json:"mid,omitempty"`
	NetworkFees []NetworkFeeModel `json:"network_fees,omitempty"`
	OriginalTransactionId string `json:"original_transaction_id"`
}

// NewClearingModel instantiates a new ClearingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClearingModel(amount int32, originalTransactionId string) *ClearingModel {
	this := ClearingModel{}
	this.Amount = amount
	var forcePost bool = false
	this.ForcePost = &forcePost
	var isRefund bool = false
	this.IsRefund = &isRefund
	this.OriginalTransactionId = originalTransactionId
	return &this
}

// NewClearingModelWithDefaults instantiates a new ClearingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClearingModelWithDefaults() *ClearingModel {
	this := ClearingModel{}
	var forcePost bool = false
	this.ForcePost = &forcePost
	var isRefund bool = false
	this.IsRefund = &isRefund
	return &this
}

// GetAmount returns the Amount field value
func (o *ClearingModel) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ClearingModel) SetAmount(v int32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *ClearingModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil || o.CardAcceptor == nil {
		var ret CardAcceptorModel
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil || o.CardAcceptor == nil {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *ClearingModel) HasCardAcceptor() bool {
	if o != nil && o.CardAcceptor != nil {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given CardAcceptorModel and assigns it to the CardAcceptor field.
func (o *ClearingModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = &v
}

// GetForcePost returns the ForcePost field value if set, zero value otherwise.
func (o *ClearingModel) GetForcePost() bool {
	if o == nil || o.ForcePost == nil {
		var ret bool
		return ret
	}
	return *o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetForcePostOk() (*bool, bool) {
	if o == nil || o.ForcePost == nil {
		return nil, false
	}
	return o.ForcePost, true
}

// HasForcePost returns a boolean if a field has been set.
func (o *ClearingModel) HasForcePost() bool {
	if o != nil && o.ForcePost != nil {
		return true
	}

	return false
}

// SetForcePost gets a reference to the given bool and assigns it to the ForcePost field.
func (o *ClearingModel) SetForcePost(v bool) {
	o.ForcePost = &v
}

// GetIsRefund returns the IsRefund field value if set, zero value otherwise.
func (o *ClearingModel) GetIsRefund() bool {
	if o == nil || o.IsRefund == nil {
		var ret bool
		return ret
	}
	return *o.IsRefund
}

// GetIsRefundOk returns a tuple with the IsRefund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetIsRefundOk() (*bool, bool) {
	if o == nil || o.IsRefund == nil {
		return nil, false
	}
	return o.IsRefund, true
}

// HasIsRefund returns a boolean if a field has been set.
func (o *ClearingModel) HasIsRefund() bool {
	if o != nil && o.IsRefund != nil {
		return true
	}

	return false
}

// SetIsRefund gets a reference to the given bool and assigns it to the IsRefund field.
func (o *ClearingModel) SetIsRefund(v bool) {
	o.IsRefund = &v
}

// GetMid returns the Mid field value if set, zero value otherwise.
func (o *ClearingModel) GetMid() string {
	if o == nil || o.Mid == nil {
		var ret string
		return ret
	}
	return *o.Mid
}

// GetMidOk returns a tuple with the Mid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetMidOk() (*string, bool) {
	if o == nil || o.Mid == nil {
		return nil, false
	}
	return o.Mid, true
}

// HasMid returns a boolean if a field has been set.
func (o *ClearingModel) HasMid() bool {
	if o != nil && o.Mid != nil {
		return true
	}

	return false
}

// SetMid gets a reference to the given string and assigns it to the Mid field.
func (o *ClearingModel) SetMid(v string) {
	o.Mid = &v
}

// GetNetworkFees returns the NetworkFees field value if set, zero value otherwise.
func (o *ClearingModel) GetNetworkFees() []NetworkFeeModel {
	if o == nil || o.NetworkFees == nil {
		var ret []NetworkFeeModel
		return ret
	}
	return o.NetworkFees
}

// GetNetworkFeesOk returns a tuple with the NetworkFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetNetworkFeesOk() ([]NetworkFeeModel, bool) {
	if o == nil || o.NetworkFees == nil {
		return nil, false
	}
	return o.NetworkFees, true
}

// HasNetworkFees returns a boolean if a field has been set.
func (o *ClearingModel) HasNetworkFees() bool {
	if o != nil && o.NetworkFees != nil {
		return true
	}

	return false
}

// SetNetworkFees gets a reference to the given []NetworkFeeModel and assigns it to the NetworkFees field.
func (o *ClearingModel) SetNetworkFees(v []NetworkFeeModel) {
	o.NetworkFees = v
}

// GetOriginalTransactionId returns the OriginalTransactionId field value
func (o *ClearingModel) GetOriginalTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalTransactionId
}

// GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field value
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetOriginalTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalTransactionId, true
}

// SetOriginalTransactionId sets field value
func (o *ClearingModel) SetOriginalTransactionId(v string) {
	o.OriginalTransactionId = v
}

func (o ClearingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.CardAcceptor != nil {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	if o.ForcePost != nil {
		toSerialize["force_post"] = o.ForcePost
	}
	if o.IsRefund != nil {
		toSerialize["is_refund"] = o.IsRefund
	}
	if o.Mid != nil {
		toSerialize["mid"] = o.Mid
	}
	if o.NetworkFees != nil {
		toSerialize["network_fees"] = o.NetworkFees
	}
	if true {
		toSerialize["original_transaction_id"] = o.OriginalTransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableClearingModel struct {
	value *ClearingModel
	isSet bool
}

func (v NullableClearingModel) Get() *ClearingModel {
	return v.value
}

func (v *NullableClearingModel) Set(val *ClearingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableClearingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableClearingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClearingModel(val *ClearingModel) *NullableClearingModel {
	return &NullableClearingModel{value: val, isSet: true}
}

func (v NullableClearingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClearingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


