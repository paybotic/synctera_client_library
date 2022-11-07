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

// BanRuleList struct for BanRuleList
type BanRuleList struct {
	// Array of ban rules
	BanRules []BanRule `json:"ban_rules"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewBanRuleList instantiates a new BanRuleList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanRuleList(banRules []BanRule) *BanRuleList {
	this := BanRuleList{}
	this.BanRules = banRules
	return &this
}

// NewBanRuleListWithDefaults instantiates a new BanRuleList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanRuleListWithDefaults() *BanRuleList {
	this := BanRuleList{}
	return &this
}

// GetBanRules returns the BanRules field value
func (o *BanRuleList) GetBanRules() []BanRule {
	if o == nil {
		var ret []BanRule
		return ret
	}

	return o.BanRules
}

// GetBanRulesOk returns a tuple with the BanRules field value
// and a boolean to check if the value has been set.
func (o *BanRuleList) GetBanRulesOk() ([]BanRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.BanRules, true
}

// SetBanRules sets field value
func (o *BanRuleList) SetBanRules(v []BanRule) {
	o.BanRules = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *BanRuleList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BanRuleList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *BanRuleList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *BanRuleList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o BanRuleList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ban_rules"] = o.BanRules
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableBanRuleList struct {
	value *BanRuleList
	isSet bool
}

func (v NullableBanRuleList) Get() *BanRuleList {
	return v.value
}

func (v *NullableBanRuleList) Set(val *BanRuleList) {
	v.value = val
	v.isSet = true
}

func (v NullableBanRuleList) IsSet() bool {
	return v.isSet
}

func (v *NullableBanRuleList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanRuleList(val *BanRuleList) *NullableBanRuleList {
	return &NullableBanRuleList{value: val, isSet: true}
}

func (v NullableBanRuleList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanRuleList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


