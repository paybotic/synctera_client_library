/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Lookup3dsRequestSdk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Lookup3dsRequestSdk{}

// Lookup3dsRequestSdk struct for Lookup3dsRequestSdk
type Lookup3dsRequestSdk struct {
	AuthenticationIndicator string `json:"authentication_indicator"`
	// Channel through which Device Data Collection was performed  Channel | Description --- | --- `BROWSER` | Internet browser `SDK` | Mobile app
	DeviceChannel string `json:"device_channel"`
	// The unique identifier of the 3DS authentication
	Id              string `json:"id"`
	TransactionMode string `json:"transaction_mode"`
}

type _Lookup3dsRequestSdk Lookup3dsRequestSdk

// NewLookup3dsRequestSdk instantiates a new Lookup3dsRequestSdk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookup3dsRequestSdk(authenticationIndicator string, deviceChannel string, id string, transactionMode string) *Lookup3dsRequestSdk {
	this := Lookup3dsRequestSdk{}
	this.AuthenticationIndicator = authenticationIndicator
	this.DeviceChannel = deviceChannel
	this.Id = id
	this.TransactionMode = transactionMode
	return &this
}

// NewLookup3dsRequestSdkWithDefaults instantiates a new Lookup3dsRequestSdk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookup3dsRequestSdkWithDefaults() *Lookup3dsRequestSdk {
	this := Lookup3dsRequestSdk{}
	return &this
}

// GetAuthenticationIndicator returns the AuthenticationIndicator field value
func (o *Lookup3dsRequestSdk) GetAuthenticationIndicator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationIndicator
}

// GetAuthenticationIndicatorOk returns a tuple with the AuthenticationIndicator field value
// and a boolean to check if the value has been set.
func (o *Lookup3dsRequestSdk) GetAuthenticationIndicatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationIndicator, true
}

// SetAuthenticationIndicator sets field value
func (o *Lookup3dsRequestSdk) SetAuthenticationIndicator(v string) {
	o.AuthenticationIndicator = v
}

// GetDeviceChannel returns the DeviceChannel field value
func (o *Lookup3dsRequestSdk) GetDeviceChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceChannel
}

// GetDeviceChannelOk returns a tuple with the DeviceChannel field value
// and a boolean to check if the value has been set.
func (o *Lookup3dsRequestSdk) GetDeviceChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceChannel, true
}

// SetDeviceChannel sets field value
func (o *Lookup3dsRequestSdk) SetDeviceChannel(v string) {
	o.DeviceChannel = v
}

// GetId returns the Id field value
func (o *Lookup3dsRequestSdk) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Lookup3dsRequestSdk) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Lookup3dsRequestSdk) SetId(v string) {
	o.Id = v
}

// GetTransactionMode returns the TransactionMode field value
func (o *Lookup3dsRequestSdk) GetTransactionMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionMode
}

// GetTransactionModeOk returns a tuple with the TransactionMode field value
// and a boolean to check if the value has been set.
func (o *Lookup3dsRequestSdk) GetTransactionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionMode, true
}

// SetTransactionMode sets field value
func (o *Lookup3dsRequestSdk) SetTransactionMode(v string) {
	o.TransactionMode = v
}

func (o Lookup3dsRequestSdk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Lookup3dsRequestSdk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authentication_indicator"] = o.AuthenticationIndicator
	toSerialize["device_channel"] = o.DeviceChannel
	toSerialize["id"] = o.Id
	toSerialize["transaction_mode"] = o.TransactionMode
	return toSerialize, nil
}

func (o *Lookup3dsRequestSdk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authentication_indicator",
		"device_channel",
		"id",
		"transaction_mode",
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

	varLookup3dsRequestSdk := _Lookup3dsRequestSdk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLookup3dsRequestSdk)

	if err != nil {
		return err
	}

	*o = Lookup3dsRequestSdk(varLookup3dsRequestSdk)

	return err
}

type NullableLookup3dsRequestSdk struct {
	value *Lookup3dsRequestSdk
	isSet bool
}

func (v NullableLookup3dsRequestSdk) Get() *Lookup3dsRequestSdk {
	return v.value
}

func (v *NullableLookup3dsRequestSdk) Set(val *Lookup3dsRequestSdk) {
	v.value = val
	v.isSet = true
}

func (v NullableLookup3dsRequestSdk) IsSet() bool {
	return v.isSet
}

func (v *NullableLookup3dsRequestSdk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookup3dsRequestSdk(val *Lookup3dsRequestSdk) *NullableLookup3dsRequestSdk {
	return &NullableLookup3dsRequestSdk{value: val, isSet: true}
}

func (v NullableLookup3dsRequestSdk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookup3dsRequestSdk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
