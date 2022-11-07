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

// GoogleDigitalWalletProvisionRequest struct for GoogleDigitalWalletProvisionRequest
type GoogleDigitalWalletProvisionRequest struct {
	// The user’s Android device ID; the device’s unique identifier.
	DeviceId string `json:"device_id"`
	DeviceType DeviceType `json:"device_type"`
	// Version of the application making the provisioning request.
	ProvisioningAppVersion string `json:"provisioning_app_version"`
	// The user’s Google wallet account ID.
	WalletAccountId string `json:"wallet_account_id"`
}

// NewGoogleDigitalWalletProvisionRequest instantiates a new GoogleDigitalWalletProvisionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleDigitalWalletProvisionRequest(deviceId string, deviceType DeviceType, provisioningAppVersion string, walletAccountId string) *GoogleDigitalWalletProvisionRequest {
	this := GoogleDigitalWalletProvisionRequest{}
	this.DeviceId = deviceId
	this.DeviceType = deviceType
	this.ProvisioningAppVersion = provisioningAppVersion
	this.WalletAccountId = walletAccountId
	return &this
}

// NewGoogleDigitalWalletProvisionRequestWithDefaults instantiates a new GoogleDigitalWalletProvisionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleDigitalWalletProvisionRequestWithDefaults() *GoogleDigitalWalletProvisionRequest {
	this := GoogleDigitalWalletProvisionRequest{}
	return &this
}

// GetDeviceId returns the DeviceId field value
func (o *GoogleDigitalWalletProvisionRequest) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *GoogleDigitalWalletProvisionRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *GoogleDigitalWalletProvisionRequest) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetDeviceType returns the DeviceType field value
func (o *GoogleDigitalWalletProvisionRequest) GetDeviceType() DeviceType {
	if o == nil {
		var ret DeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *GoogleDigitalWalletProvisionRequest) GetDeviceTypeOk() (*DeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *GoogleDigitalWalletProvisionRequest) SetDeviceType(v DeviceType) {
	o.DeviceType = v
}

// GetProvisioningAppVersion returns the ProvisioningAppVersion field value
func (o *GoogleDigitalWalletProvisionRequest) GetProvisioningAppVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisioningAppVersion
}

// GetProvisioningAppVersionOk returns a tuple with the ProvisioningAppVersion field value
// and a boolean to check if the value has been set.
func (o *GoogleDigitalWalletProvisionRequest) GetProvisioningAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningAppVersion, true
}

// SetProvisioningAppVersion sets field value
func (o *GoogleDigitalWalletProvisionRequest) SetProvisioningAppVersion(v string) {
	o.ProvisioningAppVersion = v
}

// GetWalletAccountId returns the WalletAccountId field value
func (o *GoogleDigitalWalletProvisionRequest) GetWalletAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletAccountId
}

// GetWalletAccountIdOk returns a tuple with the WalletAccountId field value
// and a boolean to check if the value has been set.
func (o *GoogleDigitalWalletProvisionRequest) GetWalletAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletAccountId, true
}

// SetWalletAccountId sets field value
func (o *GoogleDigitalWalletProvisionRequest) SetWalletAccountId(v string) {
	o.WalletAccountId = v
}

func (o GoogleDigitalWalletProvisionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["device_id"] = o.DeviceId
	}
	if true {
		toSerialize["device_type"] = o.DeviceType
	}
	if true {
		toSerialize["provisioning_app_version"] = o.ProvisioningAppVersion
	}
	if true {
		toSerialize["wallet_account_id"] = o.WalletAccountId
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleDigitalWalletProvisionRequest struct {
	value *GoogleDigitalWalletProvisionRequest
	isSet bool
}

func (v NullableGoogleDigitalWalletProvisionRequest) Get() *GoogleDigitalWalletProvisionRequest {
	return v.value
}

func (v *NullableGoogleDigitalWalletProvisionRequest) Set(val *GoogleDigitalWalletProvisionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleDigitalWalletProvisionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleDigitalWalletProvisionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleDigitalWalletProvisionRequest(val *GoogleDigitalWalletProvisionRequest) *NullableGoogleDigitalWalletProvisionRequest {
	return &NullableGoogleDigitalWalletProvisionRequest{value: val, isSet: true}
}

func (v NullableGoogleDigitalWalletProvisionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleDigitalWalletProvisionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


