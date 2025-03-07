/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the TransferListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferListResponse{}

// TransferListResponse struct for TransferListResponse
type TransferListResponse struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of External transfer
	ExternalTransfers    []TransferResponse `json:"external_transfers"`
	AdditionalProperties map[string]interface{}
}

type _TransferListResponse TransferListResponse

// NewTransferListResponse instantiates a new TransferListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferListResponse(externalTransfers []TransferResponse) *TransferListResponse {
	this := TransferListResponse{}
	this.ExternalTransfers = externalTransfers
	return &this
}

// NewTransferListResponseWithDefaults instantiates a new TransferListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferListResponseWithDefaults() *TransferListResponse {
	this := TransferListResponse{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *TransferListResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *TransferListResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *TransferListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetExternalTransfers returns the ExternalTransfers field value
func (o *TransferListResponse) GetExternalTransfers() []TransferResponse {
	if o == nil {
		var ret []TransferResponse
		return ret
	}

	return o.ExternalTransfers
}

// GetExternalTransfersOk returns a tuple with the ExternalTransfers field value
// and a boolean to check if the value has been set.
func (o *TransferListResponse) GetExternalTransfersOk() ([]TransferResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalTransfers, true
}

// SetExternalTransfers sets field value
func (o *TransferListResponse) SetExternalTransfers(v []TransferResponse) {
	o.ExternalTransfers = v
}

func (o TransferListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["external_transfers"] = o.ExternalTransfers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransferListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"external_transfers",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTransferListResponse := _TransferListResponse{}

	err = json.Unmarshal(data, &varTransferListResponse)

	if err != nil {
		return err
	}

	*o = TransferListResponse(varTransferListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "external_transfers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferListResponse struct {
	value *TransferListResponse
	isSet bool
}

func (v NullableTransferListResponse) Get() *TransferListResponse {
	return v.value
}

func (v *NullableTransferListResponse) Set(val *TransferListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferListResponse(val *TransferListResponse) *NullableTransferListResponse {
	return &NullableTransferListResponse{value: val, isSet: true}
}

func (v NullableTransferListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
