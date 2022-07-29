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

// Interest struct for Interest
type Interest struct {
	AccrualPayoutSchedule AccrualPayoutSchedule `json:"accrual_payout_schedule"`
	CalculationMethod CalculationMethod `json:"calculation_method"`
	// User provided description for the current interest.
	Description *string `json:"description,omitempty"`
	// Interest ID
	Id *string `json:"id,omitempty"`
	ProductType string `json:"product_type"`
	// A list of interest rate. Date intervals between valid_from and valid_to expect to have no overlap. 
	Rates []RateDetails `json:"rates"`
}

// NewInterest instantiates a new Interest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterest(accrualPayoutSchedule AccrualPayoutSchedule, calculationMethod CalculationMethod, productType string, rates []RateDetails) *Interest {
	this := Interest{}
	this.AccrualPayoutSchedule = accrualPayoutSchedule
	this.CalculationMethod = calculationMethod
	this.ProductType = productType
	this.Rates = rates
	return &this
}

// NewInterestWithDefaults instantiates a new Interest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterestWithDefaults() *Interest {
	this := Interest{}
	return &this
}

// GetAccrualPayoutSchedule returns the AccrualPayoutSchedule field value
func (o *Interest) GetAccrualPayoutSchedule() AccrualPayoutSchedule {
	if o == nil {
		var ret AccrualPayoutSchedule
		return ret
	}

	return o.AccrualPayoutSchedule
}

// GetAccrualPayoutScheduleOk returns a tuple with the AccrualPayoutSchedule field value
// and a boolean to check if the value has been set.
func (o *Interest) GetAccrualPayoutScheduleOk() (*AccrualPayoutSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccrualPayoutSchedule, true
}

// SetAccrualPayoutSchedule sets field value
func (o *Interest) SetAccrualPayoutSchedule(v AccrualPayoutSchedule) {
	o.AccrualPayoutSchedule = v
}

// GetCalculationMethod returns the CalculationMethod field value
func (o *Interest) GetCalculationMethod() CalculationMethod {
	if o == nil {
		var ret CalculationMethod
		return ret
	}

	return o.CalculationMethod
}

// GetCalculationMethodOk returns a tuple with the CalculationMethod field value
// and a boolean to check if the value has been set.
func (o *Interest) GetCalculationMethodOk() (*CalculationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CalculationMethod, true
}

// SetCalculationMethod sets field value
func (o *Interest) SetCalculationMethod(v CalculationMethod) {
	o.CalculationMethod = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Interest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Interest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Interest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Interest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Interest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Interest) SetId(v string) {
	o.Id = &v
}

// GetProductType returns the ProductType field value
func (o *Interest) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *Interest) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *Interest) SetProductType(v string) {
	o.ProductType = v
}

// GetRates returns the Rates field value
func (o *Interest) GetRates() []RateDetails {
	if o == nil {
		var ret []RateDetails
		return ret
	}

	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value
// and a boolean to check if the value has been set.
func (o *Interest) GetRatesOk() ([]RateDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rates, true
}

// SetRates sets field value
func (o *Interest) SetRates(v []RateDetails) {
	o.Rates = v
}

func (o Interest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accrual_payout_schedule"] = o.AccrualPayoutSchedule
	}
	if true {
		toSerialize["calculation_method"] = o.CalculationMethod
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["product_type"] = o.ProductType
	}
	if true {
		toSerialize["rates"] = o.Rates
	}
	return json.Marshal(toSerialize)
}

type NullableInterest struct {
	value *Interest
	isSet bool
}

func (v NullableInterest) Get() *Interest {
	return v.value
}

func (v *NullableInterest) Set(val *Interest) {
	v.value = val
	v.isSet = true
}

func (v NullableInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterest(val *Interest) *NullableInterest {
	return &NullableInterest{value: val, isSet: true}
}

func (v NullableInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


