/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// GoogleDigitalWalletProvisionResponse struct for GoogleDigitalWalletProvisionResponse
type GoogleDigitalWalletProvisionResponse struct {
	// The unique identifier of a card
	CardId *string `json:"card_id,omitempty"`
	CreatedTime *time.Time `json:"created_time,omitempty"`
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	PushTokenizeRequestData *PushTokenizeRequestData `json:"push_tokenize_request_data,omitempty"`
}

// NewGoogleDigitalWalletProvisionResponse instantiates a new GoogleDigitalWalletProvisionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleDigitalWalletProvisionResponse() *GoogleDigitalWalletProvisionResponse {
	this := GoogleDigitalWalletProvisionResponse{}
	return &this
}

// NewGoogleDigitalWalletProvisionResponseWithDefaults instantiates a new GoogleDigitalWalletProvisionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleDigitalWalletProvisionResponseWithDefaults() *GoogleDigitalWalletProvisionResponse {
	this := GoogleDigitalWalletProvisionResponse{}
	return &this
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *GoogleDigitalWalletProvisionResponse) GetCardId() string {
	if o == nil || o.CardId == nil {
		var ret string
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleDigitalWalletProvisionResponse) GetCardIdOk() (*string, bool) {
	if o == nil || o.CardId == nil {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *GoogleDigitalWalletProvisionResponse) HasCardId() bool {
	if o != nil && o.CardId != nil {
		return true
	}

	return false
}

// SetCardId gets a reference to the given string and assigns it to the CardId field.
func (o *GoogleDigitalWalletProvisionResponse) SetCardId(v string) {
	o.CardId = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *GoogleDigitalWalletProvisionResponse) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleDigitalWalletProvisionResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *GoogleDigitalWalletProvisionResponse) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *GoogleDigitalWalletProvisionResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *GoogleDigitalWalletProvisionResponse) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleDigitalWalletProvisionResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *GoogleDigitalWalletProvisionResponse) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *GoogleDigitalWalletProvisionResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetPushTokenizeRequestData returns the PushTokenizeRequestData field value if set, zero value otherwise.
func (o *GoogleDigitalWalletProvisionResponse) GetPushTokenizeRequestData() PushTokenizeRequestData {
	if o == nil || o.PushTokenizeRequestData == nil {
		var ret PushTokenizeRequestData
		return ret
	}
	return *o.PushTokenizeRequestData
}

// GetPushTokenizeRequestDataOk returns a tuple with the PushTokenizeRequestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleDigitalWalletProvisionResponse) GetPushTokenizeRequestDataOk() (*PushTokenizeRequestData, bool) {
	if o == nil || o.PushTokenizeRequestData == nil {
		return nil, false
	}
	return o.PushTokenizeRequestData, true
}

// HasPushTokenizeRequestData returns a boolean if a field has been set.
func (o *GoogleDigitalWalletProvisionResponse) HasPushTokenizeRequestData() bool {
	if o != nil && o.PushTokenizeRequestData != nil {
		return true
	}

	return false
}

// SetPushTokenizeRequestData gets a reference to the given PushTokenizeRequestData and assigns it to the PushTokenizeRequestData field.
func (o *GoogleDigitalWalletProvisionResponse) SetPushTokenizeRequestData(v PushTokenizeRequestData) {
	o.PushTokenizeRequestData = &v
}

func (o GoogleDigitalWalletProvisionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CardId != nil {
		toSerialize["card_id"] = o.CardId
	}
	if o.CreatedTime != nil {
		toSerialize["created_time"] = o.CreatedTime
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if o.PushTokenizeRequestData != nil {
		toSerialize["push_tokenize_request_data"] = o.PushTokenizeRequestData
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleDigitalWalletProvisionResponse struct {
	value *GoogleDigitalWalletProvisionResponse
	isSet bool
}

func (v NullableGoogleDigitalWalletProvisionResponse) Get() *GoogleDigitalWalletProvisionResponse {
	return v.value
}

func (v *NullableGoogleDigitalWalletProvisionResponse) Set(val *GoogleDigitalWalletProvisionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleDigitalWalletProvisionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleDigitalWalletProvisionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleDigitalWalletProvisionResponse(val *GoogleDigitalWalletProvisionResponse) *NullableGoogleDigitalWalletProvisionResponse {
	return &NullableGoogleDigitalWalletProvisionResponse{value: val, isSet: true}
}

func (v NullableGoogleDigitalWalletProvisionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleDigitalWalletProvisionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


