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
	"time"
)

// checks if the FinicityAccountVerification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinicityAccountVerification{}

// FinicityAccountVerification struct for FinicityAccountVerification
type FinicityAccountVerification struct {
	// The time at which verification was first completed.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The time at which verification was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// The status of verification
	Status string `json:"status"`
	// The vendor used for verifying the account
	Vendor string `json:"vendor"`
}

type _FinicityAccountVerification FinicityAccountVerification

// NewFinicityAccountVerification instantiates a new FinicityAccountVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinicityAccountVerification(status string, vendor string) *FinicityAccountVerification {
	this := FinicityAccountVerification{}
	this.Status = status
	this.Vendor = vendor
	return &this
}

// NewFinicityAccountVerificationWithDefaults instantiates a new FinicityAccountVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinicityAccountVerificationWithDefaults() *FinicityAccountVerification {
	this := FinicityAccountVerification{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *FinicityAccountVerification) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinicityAccountVerification) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *FinicityAccountVerification) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *FinicityAccountVerification) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *FinicityAccountVerification) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinicityAccountVerification) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *FinicityAccountVerification) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *FinicityAccountVerification) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetStatus returns the Status field value
func (o *FinicityAccountVerification) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FinicityAccountVerification) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FinicityAccountVerification) SetStatus(v string) {
	o.Status = v
}

// GetVendor returns the Vendor field value
func (o *FinicityAccountVerification) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *FinicityAccountVerification) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *FinicityAccountVerification) SetVendor(v string) {
	o.Vendor = v
}

func (o FinicityAccountVerification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinicityAccountVerification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	toSerialize["status"] = o.Status
	toSerialize["vendor"] = o.Vendor
	return toSerialize, nil
}

func (o *FinicityAccountVerification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"vendor",
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

	varFinicityAccountVerification := _FinicityAccountVerification{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinicityAccountVerification)

	if err != nil {
		return err
	}

	*o = FinicityAccountVerification(varFinicityAccountVerification)

	return err
}

type NullableFinicityAccountVerification struct {
	value *FinicityAccountVerification
	isSet bool
}

func (v NullableFinicityAccountVerification) Get() *FinicityAccountVerification {
	return v.value
}

func (v *NullableFinicityAccountVerification) Set(val *FinicityAccountVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableFinicityAccountVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableFinicityAccountVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinicityAccountVerification(val *FinicityAccountVerification) *NullableFinicityAccountVerification {
	return &NullableFinicityAccountVerification{value: val, isSet: true}
}

func (v NullableFinicityAccountVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinicityAccountVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
