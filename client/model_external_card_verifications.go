/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ExternalCardVerifications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalCardVerifications{}

// ExternalCardVerifications Verify card passed AVS and CVV checks and if it able to perform PUSH/PULL transfers.
type ExternalCardVerifications struct {
	// Address verification results  Status | Description --- | --- VERIFIED | AVS verified NOT_VERIFIED | AVS not verified ADDRESS_MISMATCH | ZIP code match, address no match ZIP_MISMATCH | Address match, ZIP code no match ADDRESS_AND_ZIP_MISMATCH | Address and ZIP code no match
	AddressVerificationResult string `json:"address_verification_result"`
	// Card Verification Value results  Status | Description --- | --- VERIFIED | CVV and expiration date verified NOT_VERIFIED | CVV and expiration date not verified CVV_MISMATCH | Either CVV or expiration date does not match NOT_SUPPORTED | Issuer does not participate in CVV2 service
	Cvv2Result string `json:"cvv2_result"`
	// Issuer cardholder name verification result with Account Name Inquiry (ANI) service The result of verifying the cardholder name against the name on file at the issuing institution. If this fails, it means the issuing institution has a different person's name on file as the cardholder.  Status | Description --- | --- VERIFIED | ANI Name verified NOT_VERIFIED | ANI Name not verified NOT_SUPPORTED | Issuer does not participate in ANI service NAME_MISMATCH | ANI Name does not match
	NameVerificationResult string       `json:"name_verification_result"`
	PullDetails            *PullDetails `json:"pull_details,omitempty"`
	// Indicates if the card is able to perform PULL transfers.
	PullEnabled bool         `json:"pull_enabled"`
	PushDetails *PushDetails `json:"push_details,omitempty"`
	// Indicates if the card is able to perform PUSH transfers.
	PushEnabled bool   `json:"push_enabled"`
	State       string `json:"state"`
}

type _ExternalCardVerifications ExternalCardVerifications

// NewExternalCardVerifications instantiates a new ExternalCardVerifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalCardVerifications(addressVerificationResult string, cvv2Result string, nameVerificationResult string, pullEnabled bool, pushEnabled bool, state string) *ExternalCardVerifications {
	this := ExternalCardVerifications{}
	this.AddressVerificationResult = addressVerificationResult
	this.Cvv2Result = cvv2Result
	this.NameVerificationResult = nameVerificationResult
	this.PullEnabled = pullEnabled
	this.PushEnabled = pushEnabled
	this.State = state
	return &this
}

// NewExternalCardVerificationsWithDefaults instantiates a new ExternalCardVerifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalCardVerificationsWithDefaults() *ExternalCardVerifications {
	this := ExternalCardVerifications{}
	return &this
}

// GetAddressVerificationResult returns the AddressVerificationResult field value
func (o *ExternalCardVerifications) GetAddressVerificationResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressVerificationResult
}

// GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field value
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetAddressVerificationResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressVerificationResult, true
}

// SetAddressVerificationResult sets field value
func (o *ExternalCardVerifications) SetAddressVerificationResult(v string) {
	o.AddressVerificationResult = v
}

// GetCvv2Result returns the Cvv2Result field value
func (o *ExternalCardVerifications) GetCvv2Result() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cvv2Result
}

// GetCvv2ResultOk returns a tuple with the Cvv2Result field value
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetCvv2ResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cvv2Result, true
}

// SetCvv2Result sets field value
func (o *ExternalCardVerifications) SetCvv2Result(v string) {
	o.Cvv2Result = v
}

// GetNameVerificationResult returns the NameVerificationResult field value
func (o *ExternalCardVerifications) GetNameVerificationResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameVerificationResult
}

// GetNameVerificationResultOk returns a tuple with the NameVerificationResult field value
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetNameVerificationResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameVerificationResult, true
}

// SetNameVerificationResult sets field value
func (o *ExternalCardVerifications) SetNameVerificationResult(v string) {
	o.NameVerificationResult = v
}

// GetPullDetails returns the PullDetails field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetPullDetails() PullDetails {
	if o == nil || IsNil(o.PullDetails) {
		var ret PullDetails
		return ret
	}
	return *o.PullDetails
}

// GetPullDetailsOk returns a tuple with the PullDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetPullDetailsOk() (*PullDetails, bool) {
	if o == nil || IsNil(o.PullDetails) {
		return nil, false
	}
	return o.PullDetails, true
}

// HasPullDetails returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasPullDetails() bool {
	if o != nil && !IsNil(o.PullDetails) {
		return true
	}

	return false
}

// SetPullDetails gets a reference to the given PullDetails and assigns it to the PullDetails field.
func (o *ExternalCardVerifications) SetPullDetails(v PullDetails) {
	o.PullDetails = &v
}

// GetPullEnabled returns the PullEnabled field value
func (o *ExternalCardVerifications) GetPullEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PullEnabled
}

// GetPullEnabledOk returns a tuple with the PullEnabled field value
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetPullEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PullEnabled, true
}

// SetPullEnabled sets field value
func (o *ExternalCardVerifications) SetPullEnabled(v bool) {
	o.PullEnabled = v
}

// GetPushDetails returns the PushDetails field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetPushDetails() PushDetails {
	if o == nil || IsNil(o.PushDetails) {
		var ret PushDetails
		return ret
	}
	return *o.PushDetails
}

// GetPushDetailsOk returns a tuple with the PushDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetPushDetailsOk() (*PushDetails, bool) {
	if o == nil || IsNil(o.PushDetails) {
		return nil, false
	}
	return o.PushDetails, true
}

// HasPushDetails returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasPushDetails() bool {
	if o != nil && !IsNil(o.PushDetails) {
		return true
	}

	return false
}

// SetPushDetails gets a reference to the given PushDetails and assigns it to the PushDetails field.
func (o *ExternalCardVerifications) SetPushDetails(v PushDetails) {
	o.PushDetails = &v
}

// GetPushEnabled returns the PushEnabled field value
func (o *ExternalCardVerifications) GetPushEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PushEnabled
}

// GetPushEnabledOk returns a tuple with the PushEnabled field value
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetPushEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PushEnabled, true
}

// SetPushEnabled sets field value
func (o *ExternalCardVerifications) SetPushEnabled(v bool) {
	o.PushEnabled = v
}

// GetState returns the State field value
func (o *ExternalCardVerifications) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ExternalCardVerifications) SetState(v string) {
	o.State = v
}

func (o ExternalCardVerifications) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalCardVerifications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address_verification_result"] = o.AddressVerificationResult
	toSerialize["cvv2_result"] = o.Cvv2Result
	toSerialize["name_verification_result"] = o.NameVerificationResult
	if !IsNil(o.PullDetails) {
		toSerialize["pull_details"] = o.PullDetails
	}
	toSerialize["pull_enabled"] = o.PullEnabled
	if !IsNil(o.PushDetails) {
		toSerialize["push_details"] = o.PushDetails
	}
	toSerialize["push_enabled"] = o.PushEnabled
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *ExternalCardVerifications) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address_verification_result",
		"cvv2_result",
		"name_verification_result",
		"pull_enabled",
		"push_enabled",
		"state",
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

	varExternalCardVerifications := _ExternalCardVerifications{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalCardVerifications)

	if err != nil {
		return err
	}

	*o = ExternalCardVerifications(varExternalCardVerifications)

	return err
}

type NullableExternalCardVerifications struct {
	value *ExternalCardVerifications
	isSet bool
}

func (v NullableExternalCardVerifications) Get() *ExternalCardVerifications {
	return v.value
}

func (v *NullableExternalCardVerifications) Set(val *ExternalCardVerifications) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalCardVerifications) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalCardVerifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalCardVerifications(val *ExternalCardVerifications) *NullableExternalCardVerifications {
	return &NullableExternalCardVerifications{value: val, isSet: true}
}

func (v NullableExternalCardVerifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalCardVerifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
