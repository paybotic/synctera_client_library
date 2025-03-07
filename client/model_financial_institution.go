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

// checks if the FinancialInstitution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialInstitution{}

// FinancialInstitution struct for FinancialInstitution
type FinancialInstitution struct {
	LegalAddress         *Address2 `json:"legal_address,omitempty"`
	Name                 *string   `json:"name,omitempty"`
	PhoneNumber          *string   `json:"phone_number,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FinancialInstitution FinancialInstitution

// NewFinancialInstitution instantiates a new FinancialInstitution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialInstitution() *FinancialInstitution {
	this := FinancialInstitution{}
	return &this
}

// NewFinancialInstitutionWithDefaults instantiates a new FinancialInstitution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialInstitutionWithDefaults() *FinancialInstitution {
	this := FinancialInstitution{}
	return &this
}

// GetLegalAddress returns the LegalAddress field value if set, zero value otherwise.
func (o *FinancialInstitution) GetLegalAddress() Address2 {
	if o == nil || IsNil(o.LegalAddress) {
		var ret Address2
		return ret
	}
	return *o.LegalAddress
}

// GetLegalAddressOk returns a tuple with the LegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialInstitution) GetLegalAddressOk() (*Address2, bool) {
	if o == nil || IsNil(o.LegalAddress) {
		return nil, false
	}
	return o.LegalAddress, true
}

// HasLegalAddress returns a boolean if a field has been set.
func (o *FinancialInstitution) HasLegalAddress() bool {
	if o != nil && !IsNil(o.LegalAddress) {
		return true
	}

	return false
}

// SetLegalAddress gets a reference to the given Address2 and assigns it to the LegalAddress field.
func (o *FinancialInstitution) SetLegalAddress(v Address2) {
	o.LegalAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FinancialInstitution) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialInstitution) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FinancialInstitution) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FinancialInstitution) SetName(v string) {
	o.Name = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *FinancialInstitution) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialInstitution) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *FinancialInstitution) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *FinancialInstitution) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o FinancialInstitution) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialInstitution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LegalAddress) {
		toSerialize["legal_address"] = o.LegalAddress
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FinancialInstitution) UnmarshalJSON(data []byte) (err error) {
	varFinancialInstitution := _FinancialInstitution{}

	err = json.Unmarshal(data, &varFinancialInstitution)

	if err != nil {
		return err
	}

	*o = FinancialInstitution(varFinancialInstitution)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "legal_address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "phone_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFinancialInstitution struct {
	value *FinancialInstitution
	isSet bool
}

func (v NullableFinancialInstitution) Get() *FinancialInstitution {
	return v.value
}

func (v *NullableFinancialInstitution) Set(val *FinancialInstitution) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialInstitution) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialInstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialInstitution(val *FinancialInstitution) *NullableFinancialInstitution {
	return &NullableFinancialInstitution{value: val, isSet: true}
}

func (v NullableFinancialInstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialInstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
