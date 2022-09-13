/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AddAccountsRequestRoutingIdentifiers struct for AddAccountsRequestRoutingIdentifiers
type AddAccountsRequestRoutingIdentifiers struct {
	// The routing number used for US ACH payments. 
	AchRoutingNumber string `json:"ach_routing_number"`
	// The countries that this bank operates the account in
	BankCountries []string `json:"bank_countries"`
	// The name of the bank managing the account
	BankName string `json:"bank_name"`
	// The routing number used for US wire payments. 
	WireRoutingNumber *string `json:"wire_routing_number,omitempty"`
}

// NewAddAccountsRequestRoutingIdentifiers instantiates a new AddAccountsRequestRoutingIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAccountsRequestRoutingIdentifiers(achRoutingNumber string, bankCountries []string, bankName string) *AddAccountsRequestRoutingIdentifiers {
	this := AddAccountsRequestRoutingIdentifiers{}
	this.AchRoutingNumber = achRoutingNumber
	this.BankCountries = bankCountries
	this.BankName = bankName
	return &this
}

// NewAddAccountsRequestRoutingIdentifiersWithDefaults instantiates a new AddAccountsRequestRoutingIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAccountsRequestRoutingIdentifiersWithDefaults() *AddAccountsRequestRoutingIdentifiers {
	this := AddAccountsRequestRoutingIdentifiers{}
	return &this
}

// GetAchRoutingNumber returns the AchRoutingNumber field value
func (o *AddAccountsRequestRoutingIdentifiers) GetAchRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AchRoutingNumber
}

// GetAchRoutingNumberOk returns a tuple with the AchRoutingNumber field value
// and a boolean to check if the value has been set.
func (o *AddAccountsRequestRoutingIdentifiers) GetAchRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AchRoutingNumber, true
}

// SetAchRoutingNumber sets field value
func (o *AddAccountsRequestRoutingIdentifiers) SetAchRoutingNumber(v string) {
	o.AchRoutingNumber = v
}

// GetBankCountries returns the BankCountries field value
func (o *AddAccountsRequestRoutingIdentifiers) GetBankCountries() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BankCountries
}

// GetBankCountriesOk returns a tuple with the BankCountries field value
// and a boolean to check if the value has been set.
func (o *AddAccountsRequestRoutingIdentifiers) GetBankCountriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankCountries, true
}

// SetBankCountries sets field value
func (o *AddAccountsRequestRoutingIdentifiers) SetBankCountries(v []string) {
	o.BankCountries = v
}

// GetBankName returns the BankName field value
func (o *AddAccountsRequestRoutingIdentifiers) GetBankName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
func (o *AddAccountsRequestRoutingIdentifiers) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankName, true
}

// SetBankName sets field value
func (o *AddAccountsRequestRoutingIdentifiers) SetBankName(v string) {
	o.BankName = v
}

// GetWireRoutingNumber returns the WireRoutingNumber field value if set, zero value otherwise.
func (o *AddAccountsRequestRoutingIdentifiers) GetWireRoutingNumber() string {
	if o == nil || o.WireRoutingNumber == nil {
		var ret string
		return ret
	}
	return *o.WireRoutingNumber
}

// GetWireRoutingNumberOk returns a tuple with the WireRoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAccountsRequestRoutingIdentifiers) GetWireRoutingNumberOk() (*string, bool) {
	if o == nil || o.WireRoutingNumber == nil {
		return nil, false
	}
	return o.WireRoutingNumber, true
}

// HasWireRoutingNumber returns a boolean if a field has been set.
func (o *AddAccountsRequestRoutingIdentifiers) HasWireRoutingNumber() bool {
	if o != nil && o.WireRoutingNumber != nil {
		return true
	}

	return false
}

// SetWireRoutingNumber gets a reference to the given string and assigns it to the WireRoutingNumber field.
func (o *AddAccountsRequestRoutingIdentifiers) SetWireRoutingNumber(v string) {
	o.WireRoutingNumber = &v
}

func (o AddAccountsRequestRoutingIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ach_routing_number"] = o.AchRoutingNumber
	}
	if true {
		toSerialize["bank_countries"] = o.BankCountries
	}
	if true {
		toSerialize["bank_name"] = o.BankName
	}
	if o.WireRoutingNumber != nil {
		toSerialize["wire_routing_number"] = o.WireRoutingNumber
	}
	return json.Marshal(toSerialize)
}

type NullableAddAccountsRequestRoutingIdentifiers struct {
	value *AddAccountsRequestRoutingIdentifiers
	isSet bool
}

func (v NullableAddAccountsRequestRoutingIdentifiers) Get() *AddAccountsRequestRoutingIdentifiers {
	return v.value
}

func (v *NullableAddAccountsRequestRoutingIdentifiers) Set(val *AddAccountsRequestRoutingIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAccountsRequestRoutingIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAccountsRequestRoutingIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAccountsRequestRoutingIdentifiers(val *AddAccountsRequestRoutingIdentifiers) *NullableAddAccountsRequestRoutingIdentifiers {
	return &NullableAddAccountsRequestRoutingIdentifiers{value: val, isSet: true}
}

func (v NullableAddAccountsRequestRoutingIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAccountsRequestRoutingIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


