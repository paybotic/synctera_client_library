/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// checks if the Financial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Financial{}

// Financial struct for Financial
type Financial struct {
	TaxId                   *string `json:"tax_id,omitempty"`
	TotalTaxAmount          *int64  `json:"total_tax_amount,omitempty"`
	TotalTaxAmountIndicator *string `json:"total_tax_amount_indicator,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _Financial Financial

// NewFinancial instantiates a new Financial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancial() *Financial {
	this := Financial{}
	return &this
}

// NewFinancialWithDefaults instantiates a new Financial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialWithDefaults() *Financial {
	this := Financial{}
	return &this
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *Financial) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Financial) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *Financial) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *Financial) SetTaxId(v string) {
	o.TaxId = &v
}

// GetTotalTaxAmount returns the TotalTaxAmount field value if set, zero value otherwise.
func (o *Financial) GetTotalTaxAmount() int64 {
	if o == nil || IsNil(o.TotalTaxAmount) {
		var ret int64
		return ret
	}
	return *o.TotalTaxAmount
}

// GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Financial) GetTotalTaxAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalTaxAmount) {
		return nil, false
	}
	return o.TotalTaxAmount, true
}

// HasTotalTaxAmount returns a boolean if a field has been set.
func (o *Financial) HasTotalTaxAmount() bool {
	if o != nil && !IsNil(o.TotalTaxAmount) {
		return true
	}

	return false
}

// SetTotalTaxAmount gets a reference to the given int64 and assigns it to the TotalTaxAmount field.
func (o *Financial) SetTotalTaxAmount(v int64) {
	o.TotalTaxAmount = &v
}

// GetTotalTaxAmountIndicator returns the TotalTaxAmountIndicator field value if set, zero value otherwise.
func (o *Financial) GetTotalTaxAmountIndicator() string {
	if o == nil || IsNil(o.TotalTaxAmountIndicator) {
		var ret string
		return ret
	}
	return *o.TotalTaxAmountIndicator
}

// GetTotalTaxAmountIndicatorOk returns a tuple with the TotalTaxAmountIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Financial) GetTotalTaxAmountIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.TotalTaxAmountIndicator) {
		return nil, false
	}
	return o.TotalTaxAmountIndicator, true
}

// HasTotalTaxAmountIndicator returns a boolean if a field has been set.
func (o *Financial) HasTotalTaxAmountIndicator() bool {
	if o != nil && !IsNil(o.TotalTaxAmountIndicator) {
		return true
	}

	return false
}

// SetTotalTaxAmountIndicator gets a reference to the given string and assigns it to the TotalTaxAmountIndicator field.
func (o *Financial) SetTotalTaxAmountIndicator(v string) {
	o.TotalTaxAmountIndicator = &v
}

func (o Financial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Financial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	if !IsNil(o.TotalTaxAmount) {
		toSerialize["total_tax_amount"] = o.TotalTaxAmount
	}
	if !IsNil(o.TotalTaxAmountIndicator) {
		toSerialize["total_tax_amount_indicator"] = o.TotalTaxAmountIndicator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Financial) UnmarshalJSON(data []byte) (err error) {
	varFinancial := _Financial{}

	err = json.Unmarshal(data, &varFinancial)

	if err != nil {
		return err
	}

	*o = Financial(varFinancial)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tax_id")
		delete(additionalProperties, "total_tax_amount")
		delete(additionalProperties, "total_tax_amount_indicator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFinancial struct {
	value *Financial
	isSet bool
}

func (v NullableFinancial) Get() *Financial {
	return v.value
}

func (v *NullableFinancial) Set(val *Financial) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancial) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancial(val *Financial) *NullableFinancial {
	return &NullableFinancial{value: val, isSet: true}
}

func (v NullableFinancial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
