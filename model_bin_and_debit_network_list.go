/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BinAndDebitNetworkList struct for BinAndDebitNetworkList
type BinAndDebitNetworkList struct {
	// Array of BINs and debit networks
	BinAndDebitNetworks []BinAndDebitNetwork `json:"bin_and_debit_networks"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewBinAndDebitNetworkList instantiates a new BinAndDebitNetworkList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinAndDebitNetworkList(binAndDebitNetworks []BinAndDebitNetwork) *BinAndDebitNetworkList {
	this := BinAndDebitNetworkList{}
	this.BinAndDebitNetworks = binAndDebitNetworks
	return &this
}

// NewBinAndDebitNetworkListWithDefaults instantiates a new BinAndDebitNetworkList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinAndDebitNetworkListWithDefaults() *BinAndDebitNetworkList {
	this := BinAndDebitNetworkList{}
	return &this
}

// GetBinAndDebitNetworks returns the BinAndDebitNetworks field value
func (o *BinAndDebitNetworkList) GetBinAndDebitNetworks() []BinAndDebitNetwork {
	if o == nil {
		var ret []BinAndDebitNetwork
		return ret
	}

	return o.BinAndDebitNetworks
}

// GetBinAndDebitNetworksOk returns a tuple with the BinAndDebitNetworks field value
// and a boolean to check if the value has been set.
func (o *BinAndDebitNetworkList) GetBinAndDebitNetworksOk() ([]BinAndDebitNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return o.BinAndDebitNetworks, true
}

// SetBinAndDebitNetworks sets field value
func (o *BinAndDebitNetworkList) SetBinAndDebitNetworks(v []BinAndDebitNetwork) {
	o.BinAndDebitNetworks = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *BinAndDebitNetworkList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinAndDebitNetworkList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *BinAndDebitNetworkList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *BinAndDebitNetworkList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o BinAndDebitNetworkList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bin_and_debit_networks"] = o.BinAndDebitNetworks
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableBinAndDebitNetworkList struct {
	value *BinAndDebitNetworkList
	isSet bool
}

func (v NullableBinAndDebitNetworkList) Get() *BinAndDebitNetworkList {
	return v.value
}

func (v *NullableBinAndDebitNetworkList) Set(val *BinAndDebitNetworkList) {
	v.value = val
	v.isSet = true
}

func (v NullableBinAndDebitNetworkList) IsSet() bool {
	return v.isSet
}

func (v *NullableBinAndDebitNetworkList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinAndDebitNetworkList(val *BinAndDebitNetworkList) *NullableBinAndDebitNetworkList {
	return &NullableBinAndDebitNetworkList{value: val, isSet: true}
}

func (v NullableBinAndDebitNetworkList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinAndDebitNetworkList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
