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

// checks if the CashBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashBase{}

// CashBase Cash transfer
type CashBase struct {
	// Transfer amount in cents
	Amount int64 `json:"amount"`
	// Debit or credit sign
	DcSign string `json:"dc_sign"`
	// Additional information to be added to the transfer
	SourceData map[string]interface{} `json:"source_data,omitempty"`
}

type _CashBase CashBase

// NewCashBase instantiates a new CashBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashBase(amount int64, dcSign string) *CashBase {
	this := CashBase{}
	this.Amount = amount
	this.DcSign = dcSign
	return &this
}

// NewCashBaseWithDefaults instantiates a new CashBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashBaseWithDefaults() *CashBase {
	this := CashBase{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CashBase) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CashBase) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CashBase) SetAmount(v int64) {
	o.Amount = v
}

// GetDcSign returns the DcSign field value
func (o *CashBase) GetDcSign() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *CashBase) GetDcSignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *CashBase) SetDcSign(v string) {
	o.DcSign = v
}

// GetSourceData returns the SourceData field value if set, zero value otherwise.
func (o *CashBase) GetSourceData() map[string]interface{} {
	if o == nil || IsNil(o.SourceData) {
		var ret map[string]interface{}
		return ret
	}
	return o.SourceData
}

// GetSourceDataOk returns a tuple with the SourceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashBase) GetSourceDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SourceData) {
		return map[string]interface{}{}, false
	}
	return o.SourceData, true
}

// HasSourceData returns a boolean if a field has been set.
func (o *CashBase) HasSourceData() bool {
	if o != nil && !IsNil(o.SourceData) {
		return true
	}

	return false
}

// SetSourceData gets a reference to the given map[string]interface{} and assigns it to the SourceData field.
func (o *CashBase) SetSourceData(v map[string]interface{}) {
	o.SourceData = v
}

func (o CashBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["dc_sign"] = o.DcSign
	if !IsNil(o.SourceData) {
		toSerialize["source_data"] = o.SourceData
	}
	return toSerialize, nil
}

func (o *CashBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"dc_sign",
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

	varCashBase := _CashBase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCashBase)

	if err != nil {
		return err
	}

	*o = CashBase(varCashBase)

	return err
}

type NullableCashBase struct {
	value *CashBase
	isSet bool
}

func (v NullableCashBase) Get() *CashBase {
	return v.value
}

func (v *NullableCashBase) Set(val *CashBase) {
	v.value = val
	v.isSet = true
}

func (v NullableCashBase) IsSet() bool {
	return v.isSet
}

func (v *NullableCashBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashBase(val *CashBase) *NullableCashBase {
	return &NullableCashBase{value: val, isSet: true}
}

func (v NullableCashBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
