/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TemplateFieldsLineOfCreditAllOf struct for TemplateFieldsLineOfCreditAllOf
type TemplateFieldsLineOfCreditAllOf struct {
	// The number of days an account can stay delinquent before marking an account as charged-off. 
	ChargeoffPeriod *int32 `json:"chargeoff_period,omitempty"`
	// The number of days past the due date to wait for a minimum payment before marking an account as delinquent. 
	DelinquencyPeriod *int32 `json:"delinquency_period,omitempty"`
	// The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment. 
	GracePeriod *int32 `json:"grace_period,omitempty"`
	MinimumPayment MinimumPayment `json:"minimum_payment"`
}

// NewTemplateFieldsLineOfCreditAllOf instantiates a new TemplateFieldsLineOfCreditAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateFieldsLineOfCreditAllOf(minimumPayment MinimumPayment) *TemplateFieldsLineOfCreditAllOf {
	this := TemplateFieldsLineOfCreditAllOf{}
	var chargeoffPeriod int32 = 90
	this.ChargeoffPeriod = &chargeoffPeriod
	var delinquencyPeriod int32 = 30
	this.DelinquencyPeriod = &delinquencyPeriod
	var gracePeriod int32 = 30
	this.GracePeriod = &gracePeriod
	this.MinimumPayment = minimumPayment
	return &this
}

// NewTemplateFieldsLineOfCreditAllOfWithDefaults instantiates a new TemplateFieldsLineOfCreditAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateFieldsLineOfCreditAllOfWithDefaults() *TemplateFieldsLineOfCreditAllOf {
	this := TemplateFieldsLineOfCreditAllOf{}
	var chargeoffPeriod int32 = 90
	this.ChargeoffPeriod = &chargeoffPeriod
	var delinquencyPeriod int32 = 30
	this.DelinquencyPeriod = &delinquencyPeriod
	var gracePeriod int32 = 30
	this.GracePeriod = &gracePeriod
	return &this
}

// GetChargeoffPeriod returns the ChargeoffPeriod field value if set, zero value otherwise.
func (o *TemplateFieldsLineOfCreditAllOf) GetChargeoffPeriod() int32 {
	if o == nil || o.ChargeoffPeriod == nil {
		var ret int32
		return ret
	}
	return *o.ChargeoffPeriod
}

// GetChargeoffPeriodOk returns a tuple with the ChargeoffPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsLineOfCreditAllOf) GetChargeoffPeriodOk() (*int32, bool) {
	if o == nil || o.ChargeoffPeriod == nil {
		return nil, false
	}
	return o.ChargeoffPeriod, true
}

// HasChargeoffPeriod returns a boolean if a field has been set.
func (o *TemplateFieldsLineOfCreditAllOf) HasChargeoffPeriod() bool {
	if o != nil && o.ChargeoffPeriod != nil {
		return true
	}

	return false
}

// SetChargeoffPeriod gets a reference to the given int32 and assigns it to the ChargeoffPeriod field.
func (o *TemplateFieldsLineOfCreditAllOf) SetChargeoffPeriod(v int32) {
	o.ChargeoffPeriod = &v
}

// GetDelinquencyPeriod returns the DelinquencyPeriod field value if set, zero value otherwise.
func (o *TemplateFieldsLineOfCreditAllOf) GetDelinquencyPeriod() int32 {
	if o == nil || o.DelinquencyPeriod == nil {
		var ret int32
		return ret
	}
	return *o.DelinquencyPeriod
}

// GetDelinquencyPeriodOk returns a tuple with the DelinquencyPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsLineOfCreditAllOf) GetDelinquencyPeriodOk() (*int32, bool) {
	if o == nil || o.DelinquencyPeriod == nil {
		return nil, false
	}
	return o.DelinquencyPeriod, true
}

// HasDelinquencyPeriod returns a boolean if a field has been set.
func (o *TemplateFieldsLineOfCreditAllOf) HasDelinquencyPeriod() bool {
	if o != nil && o.DelinquencyPeriod != nil {
		return true
	}

	return false
}

// SetDelinquencyPeriod gets a reference to the given int32 and assigns it to the DelinquencyPeriod field.
func (o *TemplateFieldsLineOfCreditAllOf) SetDelinquencyPeriod(v int32) {
	o.DelinquencyPeriod = &v
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *TemplateFieldsLineOfCreditAllOf) GetGracePeriod() int32 {
	if o == nil || o.GracePeriod == nil {
		var ret int32
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsLineOfCreditAllOf) GetGracePeriodOk() (*int32, bool) {
	if o == nil || o.GracePeriod == nil {
		return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *TemplateFieldsLineOfCreditAllOf) HasGracePeriod() bool {
	if o != nil && o.GracePeriod != nil {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given int32 and assigns it to the GracePeriod field.
func (o *TemplateFieldsLineOfCreditAllOf) SetGracePeriod(v int32) {
	o.GracePeriod = &v
}

// GetMinimumPayment returns the MinimumPayment field value
func (o *TemplateFieldsLineOfCreditAllOf) GetMinimumPayment() MinimumPayment {
	if o == nil {
		var ret MinimumPayment
		return ret
	}

	return o.MinimumPayment
}

// GetMinimumPaymentOk returns a tuple with the MinimumPayment field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsLineOfCreditAllOf) GetMinimumPaymentOk() (*MinimumPayment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumPayment, true
}

// SetMinimumPayment sets field value
func (o *TemplateFieldsLineOfCreditAllOf) SetMinimumPayment(v MinimumPayment) {
	o.MinimumPayment = v
}

func (o TemplateFieldsLineOfCreditAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChargeoffPeriod != nil {
		toSerialize["chargeoff_period"] = o.ChargeoffPeriod
	}
	if o.DelinquencyPeriod != nil {
		toSerialize["delinquency_period"] = o.DelinquencyPeriod
	}
	if o.GracePeriod != nil {
		toSerialize["grace_period"] = o.GracePeriod
	}
	if true {
		toSerialize["minimum_payment"] = o.MinimumPayment
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateFieldsLineOfCreditAllOf struct {
	value *TemplateFieldsLineOfCreditAllOf
	isSet bool
}

func (v NullableTemplateFieldsLineOfCreditAllOf) Get() *TemplateFieldsLineOfCreditAllOf {
	return v.value
}

func (v *NullableTemplateFieldsLineOfCreditAllOf) Set(val *TemplateFieldsLineOfCreditAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFieldsLineOfCreditAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFieldsLineOfCreditAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFieldsLineOfCreditAllOf(val *TemplateFieldsLineOfCreditAllOf) *NullableTemplateFieldsLineOfCreditAllOf {
	return &NullableTemplateFieldsLineOfCreditAllOf{value: val, isSet: true}
}

func (v NullableTemplateFieldsLineOfCreditAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFieldsLineOfCreditAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


