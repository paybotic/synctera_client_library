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

// checks if the UpdateCardImageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCardImageRequest{}

// UpdateCardImageRequest struct for UpdateCardImageRequest
type UpdateCardImageRequest struct {
	RejectionMemo        *string                   `json:"rejection_memo,omitempty"`
	RejectionReason      *CardImageRejectionReason `json:"rejection_reason,omitempty"`
	Status               CardImageStatus           `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _UpdateCardImageRequest UpdateCardImageRequest

// NewUpdateCardImageRequest instantiates a new UpdateCardImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCardImageRequest(status CardImageStatus) *UpdateCardImageRequest {
	this := UpdateCardImageRequest{}
	this.Status = status
	return &this
}

// NewUpdateCardImageRequestWithDefaults instantiates a new UpdateCardImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCardImageRequestWithDefaults() *UpdateCardImageRequest {
	this := UpdateCardImageRequest{}
	return &this
}

// GetRejectionMemo returns the RejectionMemo field value if set, zero value otherwise.
func (o *UpdateCardImageRequest) GetRejectionMemo() string {
	if o == nil || IsNil(o.RejectionMemo) {
		var ret string
		return ret
	}
	return *o.RejectionMemo
}

// GetRejectionMemoOk returns a tuple with the RejectionMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCardImageRequest) GetRejectionMemoOk() (*string, bool) {
	if o == nil || IsNil(o.RejectionMemo) {
		return nil, false
	}
	return o.RejectionMemo, true
}

// HasRejectionMemo returns a boolean if a field has been set.
func (o *UpdateCardImageRequest) HasRejectionMemo() bool {
	if o != nil && !IsNil(o.RejectionMemo) {
		return true
	}

	return false
}

// SetRejectionMemo gets a reference to the given string and assigns it to the RejectionMemo field.
func (o *UpdateCardImageRequest) SetRejectionMemo(v string) {
	o.RejectionMemo = &v
}

// GetRejectionReason returns the RejectionReason field value if set, zero value otherwise.
func (o *UpdateCardImageRequest) GetRejectionReason() CardImageRejectionReason {
	if o == nil || IsNil(o.RejectionReason) {
		var ret CardImageRejectionReason
		return ret
	}
	return *o.RejectionReason
}

// GetRejectionReasonOk returns a tuple with the RejectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCardImageRequest) GetRejectionReasonOk() (*CardImageRejectionReason, bool) {
	if o == nil || IsNil(o.RejectionReason) {
		return nil, false
	}
	return o.RejectionReason, true
}

// HasRejectionReason returns a boolean if a field has been set.
func (o *UpdateCardImageRequest) HasRejectionReason() bool {
	if o != nil && !IsNil(o.RejectionReason) {
		return true
	}

	return false
}

// SetRejectionReason gets a reference to the given CardImageRejectionReason and assigns it to the RejectionReason field.
func (o *UpdateCardImageRequest) SetRejectionReason(v CardImageRejectionReason) {
	o.RejectionReason = &v
}

// GetStatus returns the Status field value
func (o *UpdateCardImageRequest) GetStatus() CardImageStatus {
	if o == nil {
		var ret CardImageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateCardImageRequest) GetStatusOk() (*CardImageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateCardImageRequest) SetStatus(v CardImageStatus) {
	o.Status = v
}

func (o UpdateCardImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCardImageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RejectionMemo) {
		toSerialize["rejection_memo"] = o.RejectionMemo
	}
	if !IsNil(o.RejectionReason) {
		toSerialize["rejection_reason"] = o.RejectionReason
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateCardImageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varUpdateCardImageRequest := _UpdateCardImageRequest{}

	err = json.Unmarshal(data, &varUpdateCardImageRequest)

	if err != nil {
		return err
	}

	*o = UpdateCardImageRequest(varUpdateCardImageRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rejection_memo")
		delete(additionalProperties, "rejection_reason")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateCardImageRequest struct {
	value *UpdateCardImageRequest
	isSet bool
}

func (v NullableUpdateCardImageRequest) Get() *UpdateCardImageRequest {
	return v.value
}

func (v *NullableUpdateCardImageRequest) Set(val *UpdateCardImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCardImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCardImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCardImageRequest(val *UpdateCardImageRequest) *NullableUpdateCardImageRequest {
	return &NullableUpdateCardImageRequest{value: val, isSet: true}
}

func (v NullableUpdateCardImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCardImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
