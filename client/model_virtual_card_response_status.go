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

// checks if the VirtualCardResponseStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualCardResponseStatus{}

// VirtualCardResponseStatus struct for VirtualCardResponseStatus
type VirtualCardResponseStatus struct {
	CardStatus CardStatus `json:"card_status"`
	// Additional details about the reason for the status change
	Memo           *string                   `json:"memo,omitempty"`
	PendingReasons *CardStatusPendingReasons `json:"pending_reasons,omitempty"`
	StatusReason   CardStatusReasonCode      `json:"status_reason"`
}

type _VirtualCardResponseStatus VirtualCardResponseStatus

// NewVirtualCardResponseStatus instantiates a new VirtualCardResponseStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualCardResponseStatus(cardStatus CardStatus, statusReason CardStatusReasonCode) *VirtualCardResponseStatus {
	this := VirtualCardResponseStatus{}
	this.CardStatus = cardStatus
	this.StatusReason = statusReason
	return &this
}

// NewVirtualCardResponseStatusWithDefaults instantiates a new VirtualCardResponseStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualCardResponseStatusWithDefaults() *VirtualCardResponseStatus {
	this := VirtualCardResponseStatus{}
	return &this
}

// GetCardStatus returns the CardStatus field value
func (o *VirtualCardResponseStatus) GetCardStatus() CardStatus {
	if o == nil {
		var ret CardStatus
		return ret
	}

	return o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponseStatus) GetCardStatusOk() (*CardStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardStatus, true
}

// SetCardStatus sets field value
func (o *VirtualCardResponseStatus) SetCardStatus(v CardStatus) {
	o.CardStatus = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *VirtualCardResponseStatus) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponseStatus) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *VirtualCardResponseStatus) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *VirtualCardResponseStatus) SetMemo(v string) {
	o.Memo = &v
}

// GetPendingReasons returns the PendingReasons field value if set, zero value otherwise.
func (o *VirtualCardResponseStatus) GetPendingReasons() CardStatusPendingReasons {
	if o == nil || IsNil(o.PendingReasons) {
		var ret CardStatusPendingReasons
		return ret
	}
	return *o.PendingReasons
}

// GetPendingReasonsOk returns a tuple with the PendingReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponseStatus) GetPendingReasonsOk() (*CardStatusPendingReasons, bool) {
	if o == nil || IsNil(o.PendingReasons) {
		return nil, false
	}
	return o.PendingReasons, true
}

// HasPendingReasons returns a boolean if a field has been set.
func (o *VirtualCardResponseStatus) HasPendingReasons() bool {
	if o != nil && !IsNil(o.PendingReasons) {
		return true
	}

	return false
}

// SetPendingReasons gets a reference to the given CardStatusPendingReasons and assigns it to the PendingReasons field.
func (o *VirtualCardResponseStatus) SetPendingReasons(v CardStatusPendingReasons) {
	o.PendingReasons = &v
}

// GetStatusReason returns the StatusReason field value
func (o *VirtualCardResponseStatus) GetStatusReason() CardStatusReasonCode {
	if o == nil {
		var ret CardStatusReasonCode
		return ret
	}

	return o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponseStatus) GetStatusReasonOk() (*CardStatusReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusReason, true
}

// SetStatusReason sets field value
func (o *VirtualCardResponseStatus) SetStatusReason(v CardStatusReasonCode) {
	o.StatusReason = v
}

func (o VirtualCardResponseStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualCardResponseStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card_status"] = o.CardStatus
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.PendingReasons) {
		toSerialize["pending_reasons"] = o.PendingReasons
	}
	toSerialize["status_reason"] = o.StatusReason
	return toSerialize, nil
}

func (o *VirtualCardResponseStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"card_status",
		"status_reason",
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

	varVirtualCardResponseStatus := _VirtualCardResponseStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVirtualCardResponseStatus)

	if err != nil {
		return err
	}

	*o = VirtualCardResponseStatus(varVirtualCardResponseStatus)

	return err
}

type NullableVirtualCardResponseStatus struct {
	value *VirtualCardResponseStatus
	isSet bool
}

func (v NullableVirtualCardResponseStatus) Get() *VirtualCardResponseStatus {
	return v.value
}

func (v *NullableVirtualCardResponseStatus) Set(val *VirtualCardResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualCardResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualCardResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualCardResponseStatus(val *VirtualCardResponseStatus) *NullableVirtualCardResponseStatus {
	return &NullableVirtualCardResponseStatus{value: val, isSet: true}
}

func (v NullableVirtualCardResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualCardResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
