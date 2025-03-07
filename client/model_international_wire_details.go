/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the InternationalWireDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternationalWireDetails{}

// InternationalWireDetails struct for InternationalWireDetails
type InternationalWireDetails struct {
	BankAddress Address `json:"bank_address"`
	// Correspondent banks details used for international payments.
	CorrespondentBanksDetails []CorrespondentBankDetails `json:"correspondent_banks_details,omitempty"`
	// The SWIFT code (also known as BIC code) used for international payments.
	SwiftCode            string `json:"swift_code" validate:"regexp=^[A-Z]{6}[A-Z0-9]{2}([A-Z0-9]{3})?$"`
	AdditionalProperties map[string]interface{}
}

type _InternationalWireDetails InternationalWireDetails

// NewInternationalWireDetails instantiates a new InternationalWireDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternationalWireDetails(bankAddress Address, swiftCode string) *InternationalWireDetails {
	this := InternationalWireDetails{}
	this.BankAddress = bankAddress
	this.SwiftCode = swiftCode
	return &this
}

// NewInternationalWireDetailsWithDefaults instantiates a new InternationalWireDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternationalWireDetailsWithDefaults() *InternationalWireDetails {
	this := InternationalWireDetails{}
	return &this
}

// GetBankAddress returns the BankAddress field value
func (o *InternationalWireDetails) GetBankAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.BankAddress
}

// GetBankAddressOk returns a tuple with the BankAddress field value
// and a boolean to check if the value has been set.
func (o *InternationalWireDetails) GetBankAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAddress, true
}

// SetBankAddress sets field value
func (o *InternationalWireDetails) SetBankAddress(v Address) {
	o.BankAddress = v
}

// GetCorrespondentBanksDetails returns the CorrespondentBanksDetails field value if set, zero value otherwise.
func (o *InternationalWireDetails) GetCorrespondentBanksDetails() []CorrespondentBankDetails {
	if o == nil || IsNil(o.CorrespondentBanksDetails) {
		var ret []CorrespondentBankDetails
		return ret
	}
	return o.CorrespondentBanksDetails
}

// GetCorrespondentBanksDetailsOk returns a tuple with the CorrespondentBanksDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalWireDetails) GetCorrespondentBanksDetailsOk() ([]CorrespondentBankDetails, bool) {
	if o == nil || IsNil(o.CorrespondentBanksDetails) {
		return nil, false
	}
	return o.CorrespondentBanksDetails, true
}

// HasCorrespondentBanksDetails returns a boolean if a field has been set.
func (o *InternationalWireDetails) HasCorrespondentBanksDetails() bool {
	if o != nil && !IsNil(o.CorrespondentBanksDetails) {
		return true
	}

	return false
}

// SetCorrespondentBanksDetails gets a reference to the given []CorrespondentBankDetails and assigns it to the CorrespondentBanksDetails field.
func (o *InternationalWireDetails) SetCorrespondentBanksDetails(v []CorrespondentBankDetails) {
	o.CorrespondentBanksDetails = v
}

// GetSwiftCode returns the SwiftCode field value
func (o *InternationalWireDetails) GetSwiftCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SwiftCode
}

// GetSwiftCodeOk returns a tuple with the SwiftCode field value
// and a boolean to check if the value has been set.
func (o *InternationalWireDetails) GetSwiftCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SwiftCode, true
}

// SetSwiftCode sets field value
func (o *InternationalWireDetails) SetSwiftCode(v string) {
	o.SwiftCode = v
}

func (o InternationalWireDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternationalWireDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bank_address"] = o.BankAddress
	if !IsNil(o.CorrespondentBanksDetails) {
		toSerialize["correspondent_banks_details"] = o.CorrespondentBanksDetails
	}
	toSerialize["swift_code"] = o.SwiftCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InternationalWireDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bank_address",
		"swift_code",
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

	varInternationalWireDetails := _InternationalWireDetails{}

	err = json.Unmarshal(data, &varInternationalWireDetails)

	if err != nil {
		return err
	}

	*o = InternationalWireDetails(varInternationalWireDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bank_address")
		delete(additionalProperties, "correspondent_banks_details")
		delete(additionalProperties, "swift_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInternationalWireDetails struct {
	value *InternationalWireDetails
	isSet bool
}

func (v NullableInternationalWireDetails) Get() *InternationalWireDetails {
	return v.value
}

func (v *NullableInternationalWireDetails) Set(val *InternationalWireDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableInternationalWireDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableInternationalWireDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternationalWireDetails(val *InternationalWireDetails) *NullableInternationalWireDetails {
	return &NullableInternationalWireDetails{value: val, isSet: true}
}

func (v NullableInternationalWireDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternationalWireDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
