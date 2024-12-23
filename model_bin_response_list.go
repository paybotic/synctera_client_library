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

// BinResponseList struct for BinResponseList
type BinResponseList struct {
	// Array of Bins
	Bins []BinResponse `json:"bins"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewBinResponseList instantiates a new BinResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinResponseList(bins []BinResponse) *BinResponseList {
	this := BinResponseList{}
	this.Bins = bins
	return &this
}

// NewBinResponseListWithDefaults instantiates a new BinResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinResponseListWithDefaults() *BinResponseList {
	this := BinResponseList{}
	return &this
}

// GetBins returns the Bins field value
func (o *BinResponseList) GetBins() []BinResponse {
	if o == nil {
		var ret []BinResponse
		return ret
	}

	return o.Bins
}

// GetBinsOk returns a tuple with the Bins field value
// and a boolean to check if the value has been set.
func (o *BinResponseList) GetBinsOk() ([]BinResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bins, true
}

// SetBins sets field value
func (o *BinResponseList) SetBins(v []BinResponse) {
	o.Bins = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *BinResponseList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinResponseList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *BinResponseList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *BinResponseList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o BinResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bins"] = o.Bins
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableBinResponseList struct {
	value *BinResponseList
	isSet bool
}

func (v NullableBinResponseList) Get() *BinResponseList {
	return v.value
}

func (v *NullableBinResponseList) Set(val *BinResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableBinResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableBinResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinResponseList(val *BinResponseList) *NullableBinResponseList {
	return &NullableBinResponseList{value: val, isSet: true}
}

func (v NullableBinResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
