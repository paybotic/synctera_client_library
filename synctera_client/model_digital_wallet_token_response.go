/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// DigitalWalletTokenResponse struct for DigitalWalletTokenResponse
type DigitalWalletTokenResponse struct {
	ApprovedTime *time.Time `json:"approved_time,omitempty"`
	// Card ID of the Digital wallet Token
	CardId *string `json:"card_id,omitempty"`
	// The user’s Android device ID; the device’s unique identifier.
	DeviceId *string `json:"device_id,omitempty"`
	// Type of the device where the Digital Wallet Token is used in
	DeviceType *string `json:"device_type,omitempty"`
	// Digital Wallet Token ID
	Id *string `json:"id,omitempty"`
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	RequestedTime *time.Time `json:"requested_time,omitempty"`
	State *DigitalWalletTokenState `json:"state,omitempty"`
	// Type of the Digital Wallet
	Type *string `json:"type,omitempty"`
}

// NewDigitalWalletTokenResponse instantiates a new DigitalWalletTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletTokenResponse() *DigitalWalletTokenResponse {
	this := DigitalWalletTokenResponse{}
	return &this
}

// NewDigitalWalletTokenResponseWithDefaults instantiates a new DigitalWalletTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletTokenResponseWithDefaults() *DigitalWalletTokenResponse {
	this := DigitalWalletTokenResponse{}
	return &this
}

// GetApprovedTime returns the ApprovedTime field value if set, zero value otherwise.
func (o *DigitalWalletTokenResponse) GetApprovedTime() time.Time {
	if o == nil || o.ApprovedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ApprovedTime
}

// GetApprovedTimeOk returns a tuple with the ApprovedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenResponse) GetApprovedTimeOk() (*time.Time, bool) {
	if o == nil || o.ApprovedTime == nil {
		return nil, false
	}
	return o.ApprovedTime, true
}

// HasApprovedTime returns a boolean if a field has been set.
func (o *DigitalWalletTokenResponse) HasApprovedTime() bool {
	if o != nil && o.ApprovedTime != nil {
		return true
	}

	return false
}

// SetApprovedTime gets a reference to the given time.Time and assigns it to the ApprovedTime field.
func (o *DigitalWalletTokenResponse) SetApprovedTime(v time.Time) {
	o.ApprovedTime = &v
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *DigitalWalletTokenResponse) GetCardId() string {
	if o == nil || o.CardId == nil {
		var ret string
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenResponse) GetCardIdOk() (*string, bool) {
	if o == nil || o.CardId == nil {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *DigitalWalletTokenResponse) HasCardId() bool {
	if o != nil && o.CardId != nil {
		return true
	}

	return false
}

// SetCardId gets a reference to the given string and assigns it to the CardId field.
func (o *DigitalWalletTokenResponse) SetCardId(v string) {
	o.CardId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DigitalWalletTokenResponse) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenResponse) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DigitalWalletTokenResponse) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DigitalWalletTokenResponse) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *DigitalWalletTokenResponse) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenResponse) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *DigitalWalletTokenResponse) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *DigitalWalletTokenResponse) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DigitalWalletTokenResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DigitalWalletTokenResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DigitalWalletTokenResponse) SetId(v string) {
	o.Id = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *DigitalWalletTokenResponse) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *DigitalWalletTokenResponse) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *DigitalWalletTokenResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetRequestedTime returns the RequestedTime field value if set, zero value otherwise.
func (o *DigitalWalletTokenResponse) GetRequestedTime() time.Time {
	if o == nil || o.RequestedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedTime
}

// GetRequestedTimeOk returns a tuple with the RequestedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenResponse) GetRequestedTimeOk() (*time.Time, bool) {
	if o == nil || o.RequestedTime == nil {
		return nil, false
	}
	return o.RequestedTime, true
}

// HasRequestedTime returns a boolean if a field has been set.
func (o *DigitalWalletTokenResponse) HasRequestedTime() bool {
	if o != nil && o.RequestedTime != nil {
		return true
	}

	return false
}

// SetRequestedTime gets a reference to the given time.Time and assigns it to the RequestedTime field.
func (o *DigitalWalletTokenResponse) SetRequestedTime(v time.Time) {
	o.RequestedTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DigitalWalletTokenResponse) GetState() DigitalWalletTokenState {
	if o == nil || o.State == nil {
		var ret DigitalWalletTokenState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenResponse) GetStateOk() (*DigitalWalletTokenState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DigitalWalletTokenResponse) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given DigitalWalletTokenState and assigns it to the State field.
func (o *DigitalWalletTokenResponse) SetState(v DigitalWalletTokenState) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DigitalWalletTokenResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DigitalWalletTokenResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DigitalWalletTokenResponse) SetType(v string) {
	o.Type = &v
}

func (o DigitalWalletTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApprovedTime != nil {
		toSerialize["approved_time"] = o.ApprovedTime
	}
	if o.CardId != nil {
		toSerialize["card_id"] = o.CardId
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.DeviceType != nil {
		toSerialize["device_type"] = o.DeviceType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if o.RequestedTime != nil {
		toSerialize["requested_time"] = o.RequestedTime
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDigitalWalletTokenResponse struct {
	value *DigitalWalletTokenResponse
	isSet bool
}

func (v NullableDigitalWalletTokenResponse) Get() *DigitalWalletTokenResponse {
	return v.value
}

func (v *NullableDigitalWalletTokenResponse) Set(val *DigitalWalletTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletTokenResponse(val *DigitalWalletTokenResponse) *NullableDigitalWalletTokenResponse {
	return &NullableDigitalWalletTokenResponse{value: val, isSet: true}
}

func (v NullableDigitalWalletTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


