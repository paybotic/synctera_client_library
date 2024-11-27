/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Party type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Party{}

// Party Information about a party in a wire transfer
type Party struct {
	// account number of the person
	AccountNumber *string       `json:"account_number,omitempty"`
	Address       *Address3     `json:"address,omitempty"`
	AddressLines  *AddressLines `json:"address_lines,omitempty"`
	// alternate identifier of the party, used in place of account number
	AlternateIdentifier *string `json:"alternate_identifier,omitempty"`
	// name of the bank this party is a member of
	BankName *string `json:"bank_name,omitempty"`
	// type of value used to identify the party
	IdentifierType *string `json:"identifier_type,omitempty"`
	// name of the person
	Name *string `json:"name,omitempty"`
	// routing number of the bank this person is a member of
	RoutingNumber *string `json:"routing_number,omitempty"`
}

// NewParty instantiates a new Party object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParty() *Party {
	this := Party{}
	return &this
}

// NewPartyWithDefaults instantiates a new Party object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyWithDefaults() *Party {
	this := Party{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *Party) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Party) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *Party) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *Party) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Party) GetAddress() Address3 {
	if o == nil || IsNil(o.Address) {
		var ret Address3
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Party) GetAddressOk() (*Address3, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Party) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address3 and assigns it to the Address field.
func (o *Party) SetAddress(v Address3) {
	o.Address = &v
}

// GetAddressLines returns the AddressLines field value if set, zero value otherwise.
func (o *Party) GetAddressLines() AddressLines {
	if o == nil || IsNil(o.AddressLines) {
		var ret AddressLines
		return ret
	}
	return *o.AddressLines
}

// GetAddressLinesOk returns a tuple with the AddressLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Party) GetAddressLinesOk() (*AddressLines, bool) {
	if o == nil || IsNil(o.AddressLines) {
		return nil, false
	}
	return o.AddressLines, true
}

// HasAddressLines returns a boolean if a field has been set.
func (o *Party) HasAddressLines() bool {
	if o != nil && !IsNil(o.AddressLines) {
		return true
	}

	return false
}

// SetAddressLines gets a reference to the given AddressLines and assigns it to the AddressLines field.
func (o *Party) SetAddressLines(v AddressLines) {
	o.AddressLines = &v
}

// GetAlternateIdentifier returns the AlternateIdentifier field value if set, zero value otherwise.
func (o *Party) GetAlternateIdentifier() string {
	if o == nil || IsNil(o.AlternateIdentifier) {
		var ret string
		return ret
	}
	return *o.AlternateIdentifier
}

// GetAlternateIdentifierOk returns a tuple with the AlternateIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Party) GetAlternateIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateIdentifier) {
		return nil, false
	}
	return o.AlternateIdentifier, true
}

// HasAlternateIdentifier returns a boolean if a field has been set.
func (o *Party) HasAlternateIdentifier() bool {
	if o != nil && !IsNil(o.AlternateIdentifier) {
		return true
	}

	return false
}

// SetAlternateIdentifier gets a reference to the given string and assigns it to the AlternateIdentifier field.
func (o *Party) SetAlternateIdentifier(v string) {
	o.AlternateIdentifier = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *Party) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Party) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *Party) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *Party) SetBankName(v string) {
	o.BankName = &v
}

// GetIdentifierType returns the IdentifierType field value if set, zero value otherwise.
func (o *Party) GetIdentifierType() string {
	if o == nil || IsNil(o.IdentifierType) {
		var ret string
		return ret
	}
	return *o.IdentifierType
}

// GetIdentifierTypeOk returns a tuple with the IdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Party) GetIdentifierTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentifierType) {
		return nil, false
	}
	return o.IdentifierType, true
}

// HasIdentifierType returns a boolean if a field has been set.
func (o *Party) HasIdentifierType() bool {
	if o != nil && !IsNil(o.IdentifierType) {
		return true
	}

	return false
}

// SetIdentifierType gets a reference to the given string and assigns it to the IdentifierType field.
func (o *Party) SetIdentifierType(v string) {
	o.IdentifierType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Party) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Party) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Party) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Party) SetName(v string) {
	o.Name = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *Party) GetRoutingNumber() string {
	if o == nil || IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Party) GetRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *Party) HasRoutingNumber() bool {
	if o != nil && !IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *Party) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

func (o Party) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Party) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AddressLines) {
		toSerialize["address_lines"] = o.AddressLines
	}
	if !IsNil(o.AlternateIdentifier) {
		toSerialize["alternate_identifier"] = o.AlternateIdentifier
	}
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	if !IsNil(o.IdentifierType) {
		toSerialize["identifier_type"] = o.IdentifierType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RoutingNumber) {
		toSerialize["routing_number"] = o.RoutingNumber
	}
	return toSerialize, nil
}

type NullableParty struct {
	value *Party
	isSet bool
}

func (v NullableParty) Get() *Party {
	return v.value
}

func (v *NullableParty) Set(val *Party) {
	v.value = val
	v.isSet = true
}

func (v NullableParty) IsSet() bool {
	return v.isSet
}

func (v *NullableParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParty(val *Party) *NullableParty {
	return &NullableParty{value: val, isSet: true}
}

func (v NullableParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
