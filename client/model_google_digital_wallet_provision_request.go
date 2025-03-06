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

// checks if the GoogleDigitalWalletProvisionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleDigitalWalletProvisionRequest{}

// GoogleDigitalWalletProvisionRequest struct for GoogleDigitalWalletProvisionRequest
type GoogleDigitalWalletProvisionRequest struct {
	// The user’s Android device ID; the device’s unique identifier.
	DeviceId   string     `json:"device_id"`
	DeviceType DeviceType `json:"device_type"`
	// Version of the application making the provisioning request.
	ProvisioningAppVersion string `json:"provisioning_app_version"`
	// The user’s Google wallet account ID.
	WalletAccountId      string `json:"wallet_account_id"`
	AdditionalProperties map[string]interface{}
}

type _GoogleDigitalWalletProvisionRequest GoogleDigitalWalletProvisionRequest

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleDigitalWalletProvisionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device_id"] = o.DeviceId
	toSerialize["device_type"] = o.DeviceType
	toSerialize["provisioning_app_version"] = o.ProvisioningAppVersion
	toSerialize["wallet_account_id"] = o.WalletAccountId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GoogleDigitalWalletProvisionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device_id",
		"device_type",
		"provisioning_app_version",
		"wallet_account_id",
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

	varGoogleDigitalWalletProvisionRequest := _GoogleDigitalWalletProvisionRequest{}

	err = json.Unmarshal(data, &varGoogleDigitalWalletProvisionRequest)

	if err != nil {
		return err
	}

	*o = GoogleDigitalWalletProvisionRequest(varGoogleDigitalWalletProvisionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_id")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "provisioning_app_version")
		delete(additionalProperties, "wallet_account_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
