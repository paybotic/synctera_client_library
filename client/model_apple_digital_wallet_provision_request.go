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

// checks if the AppleDigitalWalletProvisionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppleDigitalWalletProvisionRequest{}

// AppleDigitalWalletProvisionRequest struct for AppleDigitalWalletProvisionRequest
type AppleDigitalWalletProvisionRequest struct {
	// Leaf and sub-CA certificates provided by Apple
	Certificates []string   `json:"certificates"`
	DeviceType   DeviceType `json:"device_type"`
	// One-time-use nonce provided by Apple for security purposes.
	Nonce string `json:"nonce"`
	// Apple-provided signature to the nonce.
	NonceSignature string `json:"nonce_signature"`
	// Version of the application making the provisioning request.
	ProvisioningAppVersion string `json:"provisioning_app_version"`
	AdditionalProperties   map[string]interface{}
}

type _AppleDigitalWalletProvisionRequest AppleDigitalWalletProvisionRequest

// NewAppleDigitalWalletProvisionRequest instantiates a new AppleDigitalWalletProvisionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppleDigitalWalletProvisionRequest(certificates []string, deviceType DeviceType, nonce string, nonceSignature string, provisioningAppVersion string) *AppleDigitalWalletProvisionRequest {
	this := AppleDigitalWalletProvisionRequest{}
	this.Certificates = certificates
	this.DeviceType = deviceType
	this.Nonce = nonce
	this.NonceSignature = nonceSignature
	this.ProvisioningAppVersion = provisioningAppVersion
	return &this
}

// NewAppleDigitalWalletProvisionRequestWithDefaults instantiates a new AppleDigitalWalletProvisionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleDigitalWalletProvisionRequestWithDefaults() *AppleDigitalWalletProvisionRequest {
	this := AppleDigitalWalletProvisionRequest{}
	return &this
}

// GetCertificates returns the Certificates field value
func (o *AppleDigitalWalletProvisionRequest) GetCertificates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionRequest) GetCertificatesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificates, true
}

// SetCertificates sets field value
func (o *AppleDigitalWalletProvisionRequest) SetCertificates(v []string) {
	o.Certificates = v
}

// GetDeviceType returns the DeviceType field value
func (o *AppleDigitalWalletProvisionRequest) GetDeviceType() DeviceType {
	if o == nil {
		var ret DeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionRequest) GetDeviceTypeOk() (*DeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *AppleDigitalWalletProvisionRequest) SetDeviceType(v DeviceType) {
	o.DeviceType = v
}

// GetNonce returns the Nonce field value
func (o *AppleDigitalWalletProvisionRequest) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionRequest) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *AppleDigitalWalletProvisionRequest) SetNonce(v string) {
	o.Nonce = v
}

// GetNonceSignature returns the NonceSignature field value
func (o *AppleDigitalWalletProvisionRequest) GetNonceSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NonceSignature
}

// GetNonceSignatureOk returns a tuple with the NonceSignature field value
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionRequest) GetNonceSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NonceSignature, true
}

// SetNonceSignature sets field value
func (o *AppleDigitalWalletProvisionRequest) SetNonceSignature(v string) {
	o.NonceSignature = v
}

// GetProvisioningAppVersion returns the ProvisioningAppVersion field value
func (o *AppleDigitalWalletProvisionRequest) GetProvisioningAppVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisioningAppVersion
}

// GetProvisioningAppVersionOk returns a tuple with the ProvisioningAppVersion field value
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionRequest) GetProvisioningAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningAppVersion, true
}

// SetProvisioningAppVersion sets field value
func (o *AppleDigitalWalletProvisionRequest) SetProvisioningAppVersion(v string) {
	o.ProvisioningAppVersion = v
}

func (o AppleDigitalWalletProvisionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppleDigitalWalletProvisionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certificates"] = o.Certificates
	toSerialize["device_type"] = o.DeviceType
	toSerialize["nonce"] = o.Nonce
	toSerialize["nonce_signature"] = o.NonceSignature
	toSerialize["provisioning_app_version"] = o.ProvisioningAppVersion

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppleDigitalWalletProvisionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificates",
		"device_type",
		"nonce",
		"nonce_signature",
		"provisioning_app_version",
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

	varAppleDigitalWalletProvisionRequest := _AppleDigitalWalletProvisionRequest{}

	err = json.Unmarshal(data, &varAppleDigitalWalletProvisionRequest)

	if err != nil {
		return err
	}

	*o = AppleDigitalWalletProvisionRequest(varAppleDigitalWalletProvisionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificates")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "nonce")
		delete(additionalProperties, "nonce_signature")
		delete(additionalProperties, "provisioning_app_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppleDigitalWalletProvisionRequest struct {
	value *AppleDigitalWalletProvisionRequest
	isSet bool
}

func (v NullableAppleDigitalWalletProvisionRequest) Get() *AppleDigitalWalletProvisionRequest {
	return v.value
}

func (v *NullableAppleDigitalWalletProvisionRequest) Set(val *AppleDigitalWalletProvisionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppleDigitalWalletProvisionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppleDigitalWalletProvisionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppleDigitalWalletProvisionRequest(val *AppleDigitalWalletProvisionRequest) *NullableAppleDigitalWalletProvisionRequest {
	return &NullableAppleDigitalWalletProvisionRequest{value: val, isSet: true}
}

func (v NullableAppleDigitalWalletProvisionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppleDigitalWalletProvisionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
