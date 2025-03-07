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

// checks if the PushTokenizeRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushTokenizeRequestData{}

// PushTokenizeRequestData struct for PushTokenizeRequestData
type PushTokenizeRequestData struct {
	DisplayName          *string   `json:"display_name,omitempty"`
	LastDigits           *string   `json:"last_digits,omitempty"`
	Network              *string   `json:"network,omitempty"`
	OpaquePaymentCard    *string   `json:"opaque_payment_card,omitempty"`
	TokenServiceProvider *string   `json:"token_service_provider,omitempty"`
	UserAddress          *Address1 `json:"user_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PushTokenizeRequestData PushTokenizeRequestData

// NewPushTokenizeRequestData instantiates a new PushTokenizeRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushTokenizeRequestData() *PushTokenizeRequestData {
	this := PushTokenizeRequestData{}
	return &this
}

// NewPushTokenizeRequestDataWithDefaults instantiates a new PushTokenizeRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushTokenizeRequestDataWithDefaults() *PushTokenizeRequestData {
	this := PushTokenizeRequestData{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PushTokenizeRequestData) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTokenizeRequestData) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PushTokenizeRequestData) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PushTokenizeRequestData) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLastDigits returns the LastDigits field value if set, zero value otherwise.
func (o *PushTokenizeRequestData) GetLastDigits() string {
	if o == nil || IsNil(o.LastDigits) {
		var ret string
		return ret
	}
	return *o.LastDigits
}

// GetLastDigitsOk returns a tuple with the LastDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTokenizeRequestData) GetLastDigitsOk() (*string, bool) {
	if o == nil || IsNil(o.LastDigits) {
		return nil, false
	}
	return o.LastDigits, true
}

// HasLastDigits returns a boolean if a field has been set.
func (o *PushTokenizeRequestData) HasLastDigits() bool {
	if o != nil && !IsNil(o.LastDigits) {
		return true
	}

	return false
}

// SetLastDigits gets a reference to the given string and assigns it to the LastDigits field.
func (o *PushTokenizeRequestData) SetLastDigits(v string) {
	o.LastDigits = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *PushTokenizeRequestData) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTokenizeRequestData) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *PushTokenizeRequestData) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *PushTokenizeRequestData) SetNetwork(v string) {
	o.Network = &v
}

// GetOpaquePaymentCard returns the OpaquePaymentCard field value if set, zero value otherwise.
func (o *PushTokenizeRequestData) GetOpaquePaymentCard() string {
	if o == nil || IsNil(o.OpaquePaymentCard) {
		var ret string
		return ret
	}
	return *o.OpaquePaymentCard
}

// GetOpaquePaymentCardOk returns a tuple with the OpaquePaymentCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTokenizeRequestData) GetOpaquePaymentCardOk() (*string, bool) {
	if o == nil || IsNil(o.OpaquePaymentCard) {
		return nil, false
	}
	return o.OpaquePaymentCard, true
}

// HasOpaquePaymentCard returns a boolean if a field has been set.
func (o *PushTokenizeRequestData) HasOpaquePaymentCard() bool {
	if o != nil && !IsNil(o.OpaquePaymentCard) {
		return true
	}

	return false
}

// SetOpaquePaymentCard gets a reference to the given string and assigns it to the OpaquePaymentCard field.
func (o *PushTokenizeRequestData) SetOpaquePaymentCard(v string) {
	o.OpaquePaymentCard = &v
}

// GetTokenServiceProvider returns the TokenServiceProvider field value if set, zero value otherwise.
func (o *PushTokenizeRequestData) GetTokenServiceProvider() string {
	if o == nil || IsNil(o.TokenServiceProvider) {
		var ret string
		return ret
	}
	return *o.TokenServiceProvider
}

// GetTokenServiceProviderOk returns a tuple with the TokenServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTokenizeRequestData) GetTokenServiceProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TokenServiceProvider) {
		return nil, false
	}
	return o.TokenServiceProvider, true
}

// HasTokenServiceProvider returns a boolean if a field has been set.
func (o *PushTokenizeRequestData) HasTokenServiceProvider() bool {
	if o != nil && !IsNil(o.TokenServiceProvider) {
		return true
	}

	return false
}

// SetTokenServiceProvider gets a reference to the given string and assigns it to the TokenServiceProvider field.
func (o *PushTokenizeRequestData) SetTokenServiceProvider(v string) {
	o.TokenServiceProvider = &v
}

// GetUserAddress returns the UserAddress field value if set, zero value otherwise.
func (o *PushTokenizeRequestData) GetUserAddress() Address1 {
	if o == nil || IsNil(o.UserAddress) {
		var ret Address1
		return ret
	}
	return *o.UserAddress
}

// GetUserAddressOk returns a tuple with the UserAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTokenizeRequestData) GetUserAddressOk() (*Address1, bool) {
	if o == nil || IsNil(o.UserAddress) {
		return nil, false
	}
	return o.UserAddress, true
}

// HasUserAddress returns a boolean if a field has been set.
func (o *PushTokenizeRequestData) HasUserAddress() bool {
	if o != nil && !IsNil(o.UserAddress) {
		return true
	}

	return false
}

// SetUserAddress gets a reference to the given Address1 and assigns it to the UserAddress field.
func (o *PushTokenizeRequestData) SetUserAddress(v Address1) {
	o.UserAddress = &v
}

func (o PushTokenizeRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushTokenizeRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.LastDigits) {
		toSerialize["last_digits"] = o.LastDigits
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.OpaquePaymentCard) {
		toSerialize["opaque_payment_card"] = o.OpaquePaymentCard
	}
	if !IsNil(o.TokenServiceProvider) {
		toSerialize["token_service_provider"] = o.TokenServiceProvider
	}
	if !IsNil(o.UserAddress) {
		toSerialize["user_address"] = o.UserAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PushTokenizeRequestData) UnmarshalJSON(data []byte) (err error) {
	varPushTokenizeRequestData := _PushTokenizeRequestData{}

	err = json.Unmarshal(data, &varPushTokenizeRequestData)

	if err != nil {
		return err
	}

	*o = PushTokenizeRequestData(varPushTokenizeRequestData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "last_digits")
		delete(additionalProperties, "network")
		delete(additionalProperties, "opaque_payment_card")
		delete(additionalProperties, "token_service_provider")
		delete(additionalProperties, "user_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePushTokenizeRequestData struct {
	value *PushTokenizeRequestData
	isSet bool
}

func (v NullablePushTokenizeRequestData) Get() *PushTokenizeRequestData {
	return v.value
}

func (v *NullablePushTokenizeRequestData) Set(val *PushTokenizeRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePushTokenizeRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePushTokenizeRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushTokenizeRequestData(val *PushTokenizeRequestData) *NullablePushTokenizeRequestData {
	return &NullablePushTokenizeRequestData{value: val, isSet: true}
}

func (v NullablePushTokenizeRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushTokenizeRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
