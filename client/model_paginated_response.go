/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// checks if the PaginatedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponse{}

// PaginatedResponse struct for PaginatedResponse
type PaginatedResponse struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken        *string `json:"next_page_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedResponse PaginatedResponse

// NewPaginatedResponse instantiates a new PaginatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponse() *PaginatedResponse {
	this := PaginatedResponse{}
	return &this
}

// NewPaginatedResponseWithDefaults instantiates a new PaginatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseWithDefaults() *PaginatedResponse {
	this := PaginatedResponse{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *PaginatedResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *PaginatedResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *PaginatedResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o PaginatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedResponse) UnmarshalJSON(data []byte) (err error) {
	varPaginatedResponse := _PaginatedResponse{}

	err = json.Unmarshal(data, &varPaginatedResponse)

	if err != nil {
		return err
	}

	*o = PaginatedResponse(varPaginatedResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedResponse struct {
	value *PaginatedResponse
	isSet bool
}

func (v NullablePaginatedResponse) Get() *PaginatedResponse {
	return v.value
}

func (v *NullablePaginatedResponse) Set(val *PaginatedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponse(val *PaginatedResponse) *NullablePaginatedResponse {
	return &NullablePaginatedResponse{value: val, isSet: true}
}

func (v NullablePaginatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
