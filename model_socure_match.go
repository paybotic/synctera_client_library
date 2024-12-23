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

// SocureMatch struct for SocureMatch
type SocureMatch struct {
	Comments    *SocureMatchComments `json:"comments,omitempty"`
	EntityId    string               `json:"entityId"`
	MatchFields []string             `json:"matchFields,omitempty"`
	SourceUrls  []string             `json:"sourceUrls,omitempty"`
	Status      string               `json:"status"`
}

// NewSocureMatch instantiates a new SocureMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocureMatch(entityId string, status string) *SocureMatch {
	this := SocureMatch{}
	this.EntityId = entityId
	this.Status = status
	return &this
}

// NewSocureMatchWithDefaults instantiates a new SocureMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocureMatchWithDefaults() *SocureMatch {
	this := SocureMatch{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *SocureMatch) GetComments() SocureMatchComments {
	if o == nil || o.Comments == nil {
		var ret SocureMatchComments
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocureMatch) GetCommentsOk() (*SocureMatchComments, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *SocureMatch) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given SocureMatchComments and assigns it to the Comments field.
func (o *SocureMatch) SetComments(v SocureMatchComments) {
	o.Comments = &v
}

// GetEntityId returns the EntityId field value
func (o *SocureMatch) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *SocureMatch) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *SocureMatch) SetEntityId(v string) {
	o.EntityId = v
}

// GetMatchFields returns the MatchFields field value if set, zero value otherwise.
func (o *SocureMatch) GetMatchFields() []string {
	if o == nil || o.MatchFields == nil {
		var ret []string
		return ret
	}
	return o.MatchFields
}

// GetMatchFieldsOk returns a tuple with the MatchFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocureMatch) GetMatchFieldsOk() ([]string, bool) {
	if o == nil || o.MatchFields == nil {
		return nil, false
	}
	return o.MatchFields, true
}

// HasMatchFields returns a boolean if a field has been set.
func (o *SocureMatch) HasMatchFields() bool {
	if o != nil && o.MatchFields != nil {
		return true
	}

	return false
}

// SetMatchFields gets a reference to the given []string and assigns it to the MatchFields field.
func (o *SocureMatch) SetMatchFields(v []string) {
	o.MatchFields = v
}

// GetSourceUrls returns the SourceUrls field value if set, zero value otherwise.
func (o *SocureMatch) GetSourceUrls() []string {
	if o == nil || o.SourceUrls == nil {
		var ret []string
		return ret
	}
	return o.SourceUrls
}

// GetSourceUrlsOk returns a tuple with the SourceUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocureMatch) GetSourceUrlsOk() ([]string, bool) {
	if o == nil || o.SourceUrls == nil {
		return nil, false
	}
	return o.SourceUrls, true
}

// HasSourceUrls returns a boolean if a field has been set.
func (o *SocureMatch) HasSourceUrls() bool {
	if o != nil && o.SourceUrls != nil {
		return true
	}

	return false
}

// SetSourceUrls gets a reference to the given []string and assigns it to the SourceUrls field.
func (o *SocureMatch) SetSourceUrls(v []string) {
	o.SourceUrls = v
}

// GetStatus returns the Status field value
func (o *SocureMatch) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SocureMatch) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SocureMatch) SetStatus(v string) {
	o.Status = v
}

func (o SocureMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if true {
		toSerialize["entityId"] = o.EntityId
	}
	if o.MatchFields != nil {
		toSerialize["matchFields"] = o.MatchFields
	}
	if o.SourceUrls != nil {
		toSerialize["sourceUrls"] = o.SourceUrls
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSocureMatch struct {
	value *SocureMatch
	isSet bool
}

func (v NullableSocureMatch) Get() *SocureMatch {
	return v.value
}

func (v *NullableSocureMatch) Set(val *SocureMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableSocureMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableSocureMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocureMatch(val *SocureMatch) *NullableSocureMatch {
	return &NullableSocureMatch{value: val, isSet: true}
}

func (v NullableSocureMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocureMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
