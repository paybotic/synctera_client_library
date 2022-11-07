/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CustomerList struct for CustomerList
type CustomerList struct {
	// Array of Customers
	Customers []CustomerInBody `json:"customers"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewCustomerList instantiates a new CustomerList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerList(customers []CustomerInBody) *CustomerList {
	this := CustomerList{}
	this.Customers = customers
	return &this
}

// NewCustomerListWithDefaults instantiates a new CustomerList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerListWithDefaults() *CustomerList {
	this := CustomerList{}
	return &this
}

// GetCustomers returns the Customers field value
func (o *CustomerList) GetCustomers() []CustomerInBody {
	if o == nil {
		var ret []CustomerInBody
		return ret
	}

	return o.Customers
}

// GetCustomersOk returns a tuple with the Customers field value
// and a boolean to check if the value has been set.
func (o *CustomerList) GetCustomersOk() ([]CustomerInBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.Customers, true
}

// SetCustomers sets field value
func (o *CustomerList) SetCustomers(v []CustomerInBody) {
	o.Customers = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CustomerList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CustomerList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CustomerList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o CustomerList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customers"] = o.Customers
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerList struct {
	value *CustomerList
	isSet bool
}

func (v NullableCustomerList) Get() *CustomerList {
	return v.value
}

func (v *NullableCustomerList) Set(val *CustomerList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerList(val *CustomerList) *NullableCustomerList {
	return &NullableCustomerList{value: val, isSet: true}
}

func (v NullableCustomerList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


