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

// AdhocVerificationResponse struct for AdhocVerificationResponse
type AdhocVerificationResponse struct {
	// Unique ID for this verification result.
	Id string `json:"id"`
	// list of watchlists that the subject of the request matched 
	MatchingWatchlists []string `json:"matching_watchlists"`
	Result VerificationResult `json:"result"`
	VendorInfo *VendorInfo `json:"vendor_info,omitempty"`
}

// NewAdhocVerificationResponse instantiates a new AdhocVerificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdhocVerificationResponse(id string, matchingWatchlists []string, result VerificationResult) *AdhocVerificationResponse {
	this := AdhocVerificationResponse{}
	this.Id = id
	this.MatchingWatchlists = matchingWatchlists
	this.Result = result
	return &this
}

// NewAdhocVerificationResponseWithDefaults instantiates a new AdhocVerificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdhocVerificationResponseWithDefaults() *AdhocVerificationResponse {
	this := AdhocVerificationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AdhocVerificationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdhocVerificationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdhocVerificationResponse) SetId(v string) {
	o.Id = v
}

// GetMatchingWatchlists returns the MatchingWatchlists field value
func (o *AdhocVerificationResponse) GetMatchingWatchlists() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MatchingWatchlists
}

// GetMatchingWatchlistsOk returns a tuple with the MatchingWatchlists field value
// and a boolean to check if the value has been set.
func (o *AdhocVerificationResponse) GetMatchingWatchlistsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MatchingWatchlists, true
}

// SetMatchingWatchlists sets field value
func (o *AdhocVerificationResponse) SetMatchingWatchlists(v []string) {
	o.MatchingWatchlists = v
}

// GetResult returns the Result field value
func (o *AdhocVerificationResponse) GetResult() VerificationResult {
	if o == nil {
		var ret VerificationResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *AdhocVerificationResponse) GetResultOk() (*VerificationResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *AdhocVerificationResponse) SetResult(v VerificationResult) {
	o.Result = v
}

// GetVendorInfo returns the VendorInfo field value if set, zero value otherwise.
func (o *AdhocVerificationResponse) GetVendorInfo() VendorInfo {
	if o == nil || o.VendorInfo == nil {
		var ret VendorInfo
		return ret
	}
	return *o.VendorInfo
}

// GetVendorInfoOk returns a tuple with the VendorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdhocVerificationResponse) GetVendorInfoOk() (*VendorInfo, bool) {
	if o == nil || o.VendorInfo == nil {
		return nil, false
	}
	return o.VendorInfo, true
}

// HasVendorInfo returns a boolean if a field has been set.
func (o *AdhocVerificationResponse) HasVendorInfo() bool {
	if o != nil && o.VendorInfo != nil {
		return true
	}

	return false
}

// SetVendorInfo gets a reference to the given VendorInfo and assigns it to the VendorInfo field.
func (o *AdhocVerificationResponse) SetVendorInfo(v VendorInfo) {
	o.VendorInfo = &v
}

func (o AdhocVerificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["matching_watchlists"] = o.MatchingWatchlists
	}
	if true {
		toSerialize["result"] = o.Result
	}
	if o.VendorInfo != nil {
		toSerialize["vendor_info"] = o.VendorInfo
	}
	return json.Marshal(toSerialize)
}

type NullableAdhocVerificationResponse struct {
	value *AdhocVerificationResponse
	isSet bool
}

func (v NullableAdhocVerificationResponse) Get() *AdhocVerificationResponse {
	return v.value
}

func (v *NullableAdhocVerificationResponse) Set(val *AdhocVerificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdhocVerificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdhocVerificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdhocVerificationResponse(val *AdhocVerificationResponse) *NullableAdhocVerificationResponse {
	return &NullableAdhocVerificationResponse{value: val, isSet: true}
}

func (v NullableAdhocVerificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdhocVerificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


