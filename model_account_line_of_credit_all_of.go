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

// AccountLineOfCreditAllOf struct for AccountLineOfCreditAllOf
type AccountLineOfCreditAllOf struct {
	// The number of days an account can stay delinquent before marking an account as charged-off. 
	ChargeoffPeriod *int32 `json:"chargeoff_period,omitempty"`
	// The credit limit for this line of credit account in cents. Minimum is 0. 
	CreditLimit *int64 `json:"credit_limit,omitempty"`
	// The number of days past the due date to wait for a minimum payment before marking an account as delinquent. 
	DelinquencyPeriod *int32 `json:"delinquency_period,omitempty"`
	// The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment. 
	GracePeriod *int32 `json:"grace_period,omitempty"`
	// An interest account product that the current account associates with. The account product must have its calculation_method set to COMPOUNDED_DAILY. 
	InterestProductId *string `json:"interest_product_id,omitempty"`
	MinimumPayment *MinimumPaymentPartial `json:"minimum_payment,omitempty"`
}

// NewAccountLineOfCreditAllOf instantiates a new AccountLineOfCreditAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountLineOfCreditAllOf() *AccountLineOfCreditAllOf {
	this := AccountLineOfCreditAllOf{}
	var chargeoffPeriod int32 = 90
	this.ChargeoffPeriod = &chargeoffPeriod
	var delinquencyPeriod int32 = 30
	this.DelinquencyPeriod = &delinquencyPeriod
	return &this
}

// NewAccountLineOfCreditAllOfWithDefaults instantiates a new AccountLineOfCreditAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountLineOfCreditAllOfWithDefaults() *AccountLineOfCreditAllOf {
	this := AccountLineOfCreditAllOf{}
	var chargeoffPeriod int32 = 90
	this.ChargeoffPeriod = &chargeoffPeriod
	var delinquencyPeriod int32 = 30
	this.DelinquencyPeriod = &delinquencyPeriod
	return &this
}

// GetChargeoffPeriod returns the ChargeoffPeriod field value if set, zero value otherwise.
func (o *AccountLineOfCreditAllOf) GetChargeoffPeriod() int32 {
	if o == nil || o.ChargeoffPeriod == nil {
		var ret int32
		return ret
	}
	return *o.ChargeoffPeriod
}

// GetChargeoffPeriodOk returns a tuple with the ChargeoffPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLineOfCreditAllOf) GetChargeoffPeriodOk() (*int32, bool) {
	if o == nil || o.ChargeoffPeriod == nil {
		return nil, false
	}
	return o.ChargeoffPeriod, true
}

// HasChargeoffPeriod returns a boolean if a field has been set.
func (o *AccountLineOfCreditAllOf) HasChargeoffPeriod() bool {
	if o != nil && o.ChargeoffPeriod != nil {
		return true
	}

	return false
}

// SetChargeoffPeriod gets a reference to the given int32 and assigns it to the ChargeoffPeriod field.
func (o *AccountLineOfCreditAllOf) SetChargeoffPeriod(v int32) {
	o.ChargeoffPeriod = &v
}

// GetCreditLimit returns the CreditLimit field value if set, zero value otherwise.
func (o *AccountLineOfCreditAllOf) GetCreditLimit() int64 {
	if o == nil || o.CreditLimit == nil {
		var ret int64
		return ret
	}
	return *o.CreditLimit
}

// GetCreditLimitOk returns a tuple with the CreditLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLineOfCreditAllOf) GetCreditLimitOk() (*int64, bool) {
	if o == nil || o.CreditLimit == nil {
		return nil, false
	}
	return o.CreditLimit, true
}

// HasCreditLimit returns a boolean if a field has been set.
func (o *AccountLineOfCreditAllOf) HasCreditLimit() bool {
	if o != nil && o.CreditLimit != nil {
		return true
	}

	return false
}

// SetCreditLimit gets a reference to the given int64 and assigns it to the CreditLimit field.
func (o *AccountLineOfCreditAllOf) SetCreditLimit(v int64) {
	o.CreditLimit = &v
}

// GetDelinquencyPeriod returns the DelinquencyPeriod field value if set, zero value otherwise.
func (o *AccountLineOfCreditAllOf) GetDelinquencyPeriod() int32 {
	if o == nil || o.DelinquencyPeriod == nil {
		var ret int32
		return ret
	}
	return *o.DelinquencyPeriod
}

// GetDelinquencyPeriodOk returns a tuple with the DelinquencyPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLineOfCreditAllOf) GetDelinquencyPeriodOk() (*int32, bool) {
	if o == nil || o.DelinquencyPeriod == nil {
		return nil, false
	}
	return o.DelinquencyPeriod, true
}

// HasDelinquencyPeriod returns a boolean if a field has been set.
func (o *AccountLineOfCreditAllOf) HasDelinquencyPeriod() bool {
	if o != nil && o.DelinquencyPeriod != nil {
		return true
	}

	return false
}

// SetDelinquencyPeriod gets a reference to the given int32 and assigns it to the DelinquencyPeriod field.
func (o *AccountLineOfCreditAllOf) SetDelinquencyPeriod(v int32) {
	o.DelinquencyPeriod = &v
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *AccountLineOfCreditAllOf) GetGracePeriod() int32 {
	if o == nil || o.GracePeriod == nil {
		var ret int32
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLineOfCreditAllOf) GetGracePeriodOk() (*int32, bool) {
	if o == nil || o.GracePeriod == nil {
		return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *AccountLineOfCreditAllOf) HasGracePeriod() bool {
	if o != nil && o.GracePeriod != nil {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given int32 and assigns it to the GracePeriod field.
func (o *AccountLineOfCreditAllOf) SetGracePeriod(v int32) {
	o.GracePeriod = &v
}

// GetInterestProductId returns the InterestProductId field value if set, zero value otherwise.
func (o *AccountLineOfCreditAllOf) GetInterestProductId() string {
	if o == nil || o.InterestProductId == nil {
		var ret string
		return ret
	}
	return *o.InterestProductId
}

// GetInterestProductIdOk returns a tuple with the InterestProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLineOfCreditAllOf) GetInterestProductIdOk() (*string, bool) {
	if o == nil || o.InterestProductId == nil {
		return nil, false
	}
	return o.InterestProductId, true
}

// HasInterestProductId returns a boolean if a field has been set.
func (o *AccountLineOfCreditAllOf) HasInterestProductId() bool {
	if o != nil && o.InterestProductId != nil {
		return true
	}

	return false
}

// SetInterestProductId gets a reference to the given string and assigns it to the InterestProductId field.
func (o *AccountLineOfCreditAllOf) SetInterestProductId(v string) {
	o.InterestProductId = &v
}

// GetMinimumPayment returns the MinimumPayment field value if set, zero value otherwise.
func (o *AccountLineOfCreditAllOf) GetMinimumPayment() MinimumPaymentPartial {
	if o == nil || o.MinimumPayment == nil {
		var ret MinimumPaymentPartial
		return ret
	}
	return *o.MinimumPayment
}

// GetMinimumPaymentOk returns a tuple with the MinimumPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLineOfCreditAllOf) GetMinimumPaymentOk() (*MinimumPaymentPartial, bool) {
	if o == nil || o.MinimumPayment == nil {
		return nil, false
	}
	return o.MinimumPayment, true
}

// HasMinimumPayment returns a boolean if a field has been set.
func (o *AccountLineOfCreditAllOf) HasMinimumPayment() bool {
	if o != nil && o.MinimumPayment != nil {
		return true
	}

	return false
}

// SetMinimumPayment gets a reference to the given MinimumPaymentPartial and assigns it to the MinimumPayment field.
func (o *AccountLineOfCreditAllOf) SetMinimumPayment(v MinimumPaymentPartial) {
	o.MinimumPayment = &v
}

func (o AccountLineOfCreditAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChargeoffPeriod != nil {
		toSerialize["chargeoff_period"] = o.ChargeoffPeriod
	}
	if o.CreditLimit != nil {
		toSerialize["credit_limit"] = o.CreditLimit
	}
	if o.DelinquencyPeriod != nil {
		toSerialize["delinquency_period"] = o.DelinquencyPeriod
	}
	if o.GracePeriod != nil {
		toSerialize["grace_period"] = o.GracePeriod
	}
	if o.InterestProductId != nil {
		toSerialize["interest_product_id"] = o.InterestProductId
	}
	if o.MinimumPayment != nil {
		toSerialize["minimum_payment"] = o.MinimumPayment
	}
	return json.Marshal(toSerialize)
}

type NullableAccountLineOfCreditAllOf struct {
	value *AccountLineOfCreditAllOf
	isSet bool
}

func (v NullableAccountLineOfCreditAllOf) Get() *AccountLineOfCreditAllOf {
	return v.value
}

func (v *NullableAccountLineOfCreditAllOf) Set(val *AccountLineOfCreditAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountLineOfCreditAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountLineOfCreditAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountLineOfCreditAllOf(val *AccountLineOfCreditAllOf) *NullableAccountLineOfCreditAllOf {
	return &NullableAccountLineOfCreditAllOf{value: val, isSet: true}
}

func (v NullableAccountLineOfCreditAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountLineOfCreditAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


