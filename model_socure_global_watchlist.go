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

// SocureGlobalWatchlist struct for SocureGlobalWatchlist
type SocureGlobalWatchlist struct {
	// Contains key-value pair of the Source list name and an array of details about that match.
	Matches map[string][]SocureMatch `json:"matches"`
	// Array of reason codes
	ReasonCodes []string `json:"reasonCodes,omitempty"`
}

// NewSocureGlobalWatchlist instantiates a new SocureGlobalWatchlist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocureGlobalWatchlist(matches map[string][]SocureMatch) *SocureGlobalWatchlist {
	this := SocureGlobalWatchlist{}
	this.Matches = matches
	return &this
}

// NewSocureGlobalWatchlistWithDefaults instantiates a new SocureGlobalWatchlist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocureGlobalWatchlistWithDefaults() *SocureGlobalWatchlist {
	this := SocureGlobalWatchlist{}
	return &this
}

// GetMatches returns the Matches field value
func (o *SocureGlobalWatchlist) GetMatches() map[string][]SocureMatch {
	if o == nil {
		var ret map[string][]SocureMatch
		return ret
	}

	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value
// and a boolean to check if the value has been set.
func (o *SocureGlobalWatchlist) GetMatchesOk() (*map[string][]SocureMatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Matches, true
}

// SetMatches sets field value
func (o *SocureGlobalWatchlist) SetMatches(v map[string][]SocureMatch) {
	o.Matches = v
}

// GetReasonCodes returns the ReasonCodes field value if set, zero value otherwise.
func (o *SocureGlobalWatchlist) GetReasonCodes() []string {
	if o == nil || o.ReasonCodes == nil {
		var ret []string
		return ret
	}
	return o.ReasonCodes
}

// GetReasonCodesOk returns a tuple with the ReasonCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocureGlobalWatchlist) GetReasonCodesOk() ([]string, bool) {
	if o == nil || o.ReasonCodes == nil {
		return nil, false
	}
	return o.ReasonCodes, true
}

// HasReasonCodes returns a boolean if a field has been set.
func (o *SocureGlobalWatchlist) HasReasonCodes() bool {
	if o != nil && o.ReasonCodes != nil {
		return true
	}

	return false
}

// SetReasonCodes gets a reference to the given []string and assigns it to the ReasonCodes field.
func (o *SocureGlobalWatchlist) SetReasonCodes(v []string) {
	o.ReasonCodes = v
}

func (o SocureGlobalWatchlist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["matches"] = o.Matches
	}
	if o.ReasonCodes != nil {
		toSerialize["reasonCodes"] = o.ReasonCodes
	}
	return json.Marshal(toSerialize)
}

type NullableSocureGlobalWatchlist struct {
	value *SocureGlobalWatchlist
	isSet bool
}

func (v NullableSocureGlobalWatchlist) Get() *SocureGlobalWatchlist {
	return v.value
}

func (v *NullableSocureGlobalWatchlist) Set(val *SocureGlobalWatchlist) {
	v.value = val
	v.isSet = true
}

func (v NullableSocureGlobalWatchlist) IsSet() bool {
	return v.isSet
}

func (v *NullableSocureGlobalWatchlist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocureGlobalWatchlist(val *SocureGlobalWatchlist) *NullableSocureGlobalWatchlist {
	return &NullableSocureGlobalWatchlist{value: val, isSet: true}
}

func (v NullableSocureGlobalWatchlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocureGlobalWatchlist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


