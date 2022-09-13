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

// ApplicationList struct for ApplicationList
type ApplicationList struct {
	// Array of credit applications.
	Applications []ApplicationResponse `json:"applications"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewApplicationList instantiates a new ApplicationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationList(applications []ApplicationResponse) *ApplicationList {
	this := ApplicationList{}
	this.Applications = applications
	return &this
}

// NewApplicationListWithDefaults instantiates a new ApplicationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationListWithDefaults() *ApplicationList {
	this := ApplicationList{}
	return &this
}

// GetApplications returns the Applications field value
func (o *ApplicationList) GetApplications() []ApplicationResponse {
	if o == nil {
		var ret []ApplicationResponse
		return ret
	}

	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value
// and a boolean to check if the value has been set.
func (o *ApplicationList) GetApplicationsOk() ([]ApplicationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Applications, true
}

// SetApplications sets field value
func (o *ApplicationList) SetApplications(v []ApplicationResponse) {
	o.Applications = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ApplicationList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ApplicationList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ApplicationList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ApplicationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["applications"] = o.Applications
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationList struct {
	value *ApplicationList
	isSet bool
}

func (v NullableApplicationList) Get() *ApplicationList {
	return v.value
}

func (v *NullableApplicationList) Set(val *ApplicationList) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationList) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationList(val *ApplicationList) *NullableApplicationList {
	return &NullableApplicationList{value: val, isSet: true}
}

func (v NullableApplicationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


