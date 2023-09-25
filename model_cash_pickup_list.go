/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CashPickupList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashPickupList{}

// CashPickupList struct for CashPickupList
type CashPickupList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of cash pickups
	CashPickups []CashPickup `json:"cash_pickups"`
}

// NewCashPickupList instantiates a new CashPickupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashPickupList(cashPickups []CashPickup) *CashPickupList {
	this := CashPickupList{}
	this.CashPickups = cashPickups
	return &this
}

// NewCashPickupListWithDefaults instantiates a new CashPickupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashPickupListWithDefaults() *CashPickupList {
	this := CashPickupList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CashPickupList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashPickupList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CashPickupList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CashPickupList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetCashPickups returns the CashPickups field value
func (o *CashPickupList) GetCashPickups() []CashPickup {
	if o == nil {
		var ret []CashPickup
		return ret
	}

	return o.CashPickups
}

// GetCashPickupsOk returns a tuple with the CashPickups field value
// and a boolean to check if the value has been set.
func (o *CashPickupList) GetCashPickupsOk() ([]CashPickup, bool) {
	if o == nil {
		return nil, false
	}
	return o.CashPickups, true
}

// SetCashPickups sets field value
func (o *CashPickupList) SetCashPickups(v []CashPickup) {
	o.CashPickups = v
}

func (o CashPickupList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashPickupList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["cash_pickups"] = o.CashPickups
	return toSerialize, nil
}

type NullableCashPickupList struct {
	value *CashPickupList
	isSet bool
}

func (v NullableCashPickupList) Get() *CashPickupList {
	return v.value
}

func (v *NullableCashPickupList) Set(val *CashPickupList) {
	v.value = val
	v.isSet = true
}

func (v NullableCashPickupList) IsSet() bool {
	return v.isSet
}

func (v *NullableCashPickupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashPickupList(val *CashPickupList) *NullableCashPickupList {
	return &NullableCashPickupList{value: val, isSet: true}
}

func (v NullableCashPickupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashPickupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


