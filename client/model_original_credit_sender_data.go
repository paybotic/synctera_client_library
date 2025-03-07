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

// checks if the OriginalCreditSenderData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginalCreditSenderData{}

// OriginalCreditSenderData struct for OriginalCreditSenderData
type OriginalCreditSenderData struct {
	FundingSource                    string  `json:"funding_source"`
	SenderAccountNumber              *string `json:"sender_account_number,omitempty"`
	SenderAccountType                *string `json:"sender_account_type,omitempty"`
	SenderAddress                    *string `json:"sender_address,omitempty"`
	SenderCity                       *string `json:"sender_city,omitempty"`
	SenderCountry                    *string `json:"sender_country,omitempty"`
	SenderName                       *string `json:"sender_name,omitempty"`
	SenderReferenceNumber            *string `json:"sender_reference_number,omitempty"`
	SenderState                      *string `json:"sender_state,omitempty"`
	TransactionPurpose               *string `json:"transaction_purpose,omitempty"`
	UniqueTransactionReferenceNumber *string `json:"unique_transaction_reference_number,omitempty"`
}

type _OriginalCreditSenderData OriginalCreditSenderData

// NewOriginalCreditSenderData instantiates a new OriginalCreditSenderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginalCreditSenderData(fundingSource string) *OriginalCreditSenderData {
	this := OriginalCreditSenderData{}
	this.FundingSource = fundingSource
	return &this
}

// NewOriginalCreditSenderDataWithDefaults instantiates a new OriginalCreditSenderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginalCreditSenderDataWithDefaults() *OriginalCreditSenderData {
	this := OriginalCreditSenderData{}
	return &this
}

// GetFundingSource returns the FundingSource field value
func (o *OriginalCreditSenderData) GetFundingSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetFundingSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingSource, true
}

// SetFundingSource sets field value
func (o *OriginalCreditSenderData) SetFundingSource(v string) {
	o.FundingSource = v
}

// GetSenderAccountNumber returns the SenderAccountNumber field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderAccountNumber() string {
	if o == nil || IsNil(o.SenderAccountNumber) {
		var ret string
		return ret
	}
	return *o.SenderAccountNumber
}

// GetSenderAccountNumberOk returns a tuple with the SenderAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SenderAccountNumber) {
		return nil, false
	}
	return o.SenderAccountNumber, true
}

// HasSenderAccountNumber returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderAccountNumber() bool {
	if o != nil && !IsNil(o.SenderAccountNumber) {
		return true
	}

	return false
}

// SetSenderAccountNumber gets a reference to the given string and assigns it to the SenderAccountNumber field.
func (o *OriginalCreditSenderData) SetSenderAccountNumber(v string) {
	o.SenderAccountNumber = &v
}

// GetSenderAccountType returns the SenderAccountType field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderAccountType() string {
	if o == nil || IsNil(o.SenderAccountType) {
		var ret string
		return ret
	}
	return *o.SenderAccountType
}

// GetSenderAccountTypeOk returns a tuple with the SenderAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SenderAccountType) {
		return nil, false
	}
	return o.SenderAccountType, true
}

// HasSenderAccountType returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderAccountType() bool {
	if o != nil && !IsNil(o.SenderAccountType) {
		return true
	}

	return false
}

// SetSenderAccountType gets a reference to the given string and assigns it to the SenderAccountType field.
func (o *OriginalCreditSenderData) SetSenderAccountType(v string) {
	o.SenderAccountType = &v
}

// GetSenderAddress returns the SenderAddress field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderAddress() string {
	if o == nil || IsNil(o.SenderAddress) {
		var ret string
		return ret
	}
	return *o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SenderAddress) {
		return nil, false
	}
	return o.SenderAddress, true
}

// HasSenderAddress returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderAddress() bool {
	if o != nil && !IsNil(o.SenderAddress) {
		return true
	}

	return false
}

// SetSenderAddress gets a reference to the given string and assigns it to the SenderAddress field.
func (o *OriginalCreditSenderData) SetSenderAddress(v string) {
	o.SenderAddress = &v
}

// GetSenderCity returns the SenderCity field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderCity() string {
	if o == nil || IsNil(o.SenderCity) {
		var ret string
		return ret
	}
	return *o.SenderCity
}

// GetSenderCityOk returns a tuple with the SenderCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderCityOk() (*string, bool) {
	if o == nil || IsNil(o.SenderCity) {
		return nil, false
	}
	return o.SenderCity, true
}

// HasSenderCity returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderCity() bool {
	if o != nil && !IsNil(o.SenderCity) {
		return true
	}

	return false
}

// SetSenderCity gets a reference to the given string and assigns it to the SenderCity field.
func (o *OriginalCreditSenderData) SetSenderCity(v string) {
	o.SenderCity = &v
}

// GetSenderCountry returns the SenderCountry field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderCountry() string {
	if o == nil || IsNil(o.SenderCountry) {
		var ret string
		return ret
	}
	return *o.SenderCountry
}

// GetSenderCountryOk returns a tuple with the SenderCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderCountryOk() (*string, bool) {
	if o == nil || IsNil(o.SenderCountry) {
		return nil, false
	}
	return o.SenderCountry, true
}

// HasSenderCountry returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderCountry() bool {
	if o != nil && !IsNil(o.SenderCountry) {
		return true
	}

	return false
}

// SetSenderCountry gets a reference to the given string and assigns it to the SenderCountry field.
func (o *OriginalCreditSenderData) SetSenderCountry(v string) {
	o.SenderCountry = &v
}

// GetSenderName returns the SenderName field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderName() string {
	if o == nil || IsNil(o.SenderName) {
		var ret string
		return ret
	}
	return *o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SenderName) {
		return nil, false
	}
	return o.SenderName, true
}

// HasSenderName returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderName() bool {
	if o != nil && !IsNil(o.SenderName) {
		return true
	}

	return false
}

// SetSenderName gets a reference to the given string and assigns it to the SenderName field.
func (o *OriginalCreditSenderData) SetSenderName(v string) {
	o.SenderName = &v
}

// GetSenderReferenceNumber returns the SenderReferenceNumber field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderReferenceNumber() string {
	if o == nil || IsNil(o.SenderReferenceNumber) {
		var ret string
		return ret
	}
	return *o.SenderReferenceNumber
}

// GetSenderReferenceNumberOk returns a tuple with the SenderReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderReferenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SenderReferenceNumber) {
		return nil, false
	}
	return o.SenderReferenceNumber, true
}

// HasSenderReferenceNumber returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderReferenceNumber() bool {
	if o != nil && !IsNil(o.SenderReferenceNumber) {
		return true
	}

	return false
}

// SetSenderReferenceNumber gets a reference to the given string and assigns it to the SenderReferenceNumber field.
func (o *OriginalCreditSenderData) SetSenderReferenceNumber(v string) {
	o.SenderReferenceNumber = &v
}

// GetSenderState returns the SenderState field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderState() string {
	if o == nil || IsNil(o.SenderState) {
		var ret string
		return ret
	}
	return *o.SenderState
}

// GetSenderStateOk returns a tuple with the SenderState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderStateOk() (*string, bool) {
	if o == nil || IsNil(o.SenderState) {
		return nil, false
	}
	return o.SenderState, true
}

// HasSenderState returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderState() bool {
	if o != nil && !IsNil(o.SenderState) {
		return true
	}

	return false
}

// SetSenderState gets a reference to the given string and assigns it to the SenderState field.
func (o *OriginalCreditSenderData) SetSenderState(v string) {
	o.SenderState = &v
}

// GetTransactionPurpose returns the TransactionPurpose field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetTransactionPurpose() string {
	if o == nil || IsNil(o.TransactionPurpose) {
		var ret string
		return ret
	}
	return *o.TransactionPurpose
}

// GetTransactionPurposeOk returns a tuple with the TransactionPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetTransactionPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionPurpose) {
		return nil, false
	}
	return o.TransactionPurpose, true
}

// HasTransactionPurpose returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasTransactionPurpose() bool {
	if o != nil && !IsNil(o.TransactionPurpose) {
		return true
	}

	return false
}

// SetTransactionPurpose gets a reference to the given string and assigns it to the TransactionPurpose field.
func (o *OriginalCreditSenderData) SetTransactionPurpose(v string) {
	o.TransactionPurpose = &v
}

// GetUniqueTransactionReferenceNumber returns the UniqueTransactionReferenceNumber field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetUniqueTransactionReferenceNumber() string {
	if o == nil || IsNil(o.UniqueTransactionReferenceNumber) {
		var ret string
		return ret
	}
	return *o.UniqueTransactionReferenceNumber
}

// GetUniqueTransactionReferenceNumberOk returns a tuple with the UniqueTransactionReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetUniqueTransactionReferenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueTransactionReferenceNumber) {
		return nil, false
	}
	return o.UniqueTransactionReferenceNumber, true
}

// HasUniqueTransactionReferenceNumber returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasUniqueTransactionReferenceNumber() bool {
	if o != nil && !IsNil(o.UniqueTransactionReferenceNumber) {
		return true
	}

	return false
}

// SetUniqueTransactionReferenceNumber gets a reference to the given string and assigns it to the UniqueTransactionReferenceNumber field.
func (o *OriginalCreditSenderData) SetUniqueTransactionReferenceNumber(v string) {
	o.UniqueTransactionReferenceNumber = &v
}

func (o OriginalCreditSenderData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginalCreditSenderData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["funding_source"] = o.FundingSource
	if !IsNil(o.SenderAccountNumber) {
		toSerialize["sender_account_number"] = o.SenderAccountNumber
	}
	if !IsNil(o.SenderAccountType) {
		toSerialize["sender_account_type"] = o.SenderAccountType
	}
	if !IsNil(o.SenderAddress) {
		toSerialize["sender_address"] = o.SenderAddress
	}
	if !IsNil(o.SenderCity) {
		toSerialize["sender_city"] = o.SenderCity
	}
	if !IsNil(o.SenderCountry) {
		toSerialize["sender_country"] = o.SenderCountry
	}
	if !IsNil(o.SenderName) {
		toSerialize["sender_name"] = o.SenderName
	}
	if !IsNil(o.SenderReferenceNumber) {
		toSerialize["sender_reference_number"] = o.SenderReferenceNumber
	}
	if !IsNil(o.SenderState) {
		toSerialize["sender_state"] = o.SenderState
	}
	if !IsNil(o.TransactionPurpose) {
		toSerialize["transaction_purpose"] = o.TransactionPurpose
	}
	if !IsNil(o.UniqueTransactionReferenceNumber) {
		toSerialize["unique_transaction_reference_number"] = o.UniqueTransactionReferenceNumber
	}
	return toSerialize, nil
}

func (o *OriginalCreditSenderData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"funding_source",
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

	varOriginalCreditSenderData := _OriginalCreditSenderData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOriginalCreditSenderData)

	if err != nil {
		return err
	}

	*o = OriginalCreditSenderData(varOriginalCreditSenderData)

	return err
}

type NullableOriginalCreditSenderData struct {
	value *OriginalCreditSenderData
	isSet bool
}

func (v NullableOriginalCreditSenderData) Get() *OriginalCreditSenderData {
	return v.value
}

func (v *NullableOriginalCreditSenderData) Set(val *OriginalCreditSenderData) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginalCreditSenderData) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginalCreditSenderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginalCreditSenderData(val *OriginalCreditSenderData) *NullableOriginalCreditSenderData {
	return &NullableOriginalCreditSenderData{value: val, isSet: true}
}

func (v NullableOriginalCreditSenderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginalCreditSenderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
