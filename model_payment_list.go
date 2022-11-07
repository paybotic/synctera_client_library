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

// PaymentList struct for PaymentList
type PaymentList struct {
	// Array of payments
	Payments []Payment `json:"payments"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewPaymentList instantiates a new PaymentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentList(payments []Payment) *PaymentList {
	this := PaymentList{}
	this.Payments = payments
	return &this
}

// NewPaymentListWithDefaults instantiates a new PaymentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentListWithDefaults() *PaymentList {
	this := PaymentList{}
	return &this
}

// GetPayments returns the Payments field value
func (o *PaymentList) GetPayments() []Payment {
	if o == nil {
		var ret []Payment
		return ret
	}

	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value
// and a boolean to check if the value has been set.
func (o *PaymentList) GetPaymentsOk() ([]Payment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payments, true
}

// SetPayments sets field value
func (o *PaymentList) SetPayments(v []Payment) {
	o.Payments = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *PaymentList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *PaymentList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *PaymentList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o PaymentList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payments"] = o.Payments
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentList struct {
	value *PaymentList
	isSet bool
}

func (v NullablePaymentList) Get() *PaymentList {
	return v.value
}

func (v *NullablePaymentList) Set(val *PaymentList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentList(val *PaymentList) *NullablePaymentList {
	return &NullablePaymentList{value: val, isSet: true}
}

func (v NullablePaymentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


