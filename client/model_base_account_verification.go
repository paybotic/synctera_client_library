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
	"time"
)

// checks if the BaseAccountVerification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseAccountVerification{}

// BaseAccountVerification struct for BaseAccountVerification
type BaseAccountVerification struct {
	// The time at which verification was first completed.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The time at which verification was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// The status of verification
	Status string `json:"status"`
	// The vendor used for verifying the account
	Vendor               string `json:"vendor"`
	AdditionalProperties map[string]interface{}
}

type _BaseAccountVerification BaseAccountVerification

// NewBaseAccountVerification instantiates a new BaseAccountVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseAccountVerification(status string, vendor string) *BaseAccountVerification {
	this := BaseAccountVerification{}
	this.Status = status
	this.Vendor = vendor
	return &this
}

// NewBaseAccountVerificationWithDefaults instantiates a new BaseAccountVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseAccountVerificationWithDefaults() *BaseAccountVerification {
	this := BaseAccountVerification{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BaseAccountVerification) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccountVerification) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BaseAccountVerification) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BaseAccountVerification) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *BaseAccountVerification) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAccountVerification) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *BaseAccountVerification) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *BaseAccountVerification) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetStatus returns the Status field value
func (o *BaseAccountVerification) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BaseAccountVerification) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BaseAccountVerification) SetStatus(v string) {
	o.Status = v
}

// GetVendor returns the Vendor field value
func (o *BaseAccountVerification) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *BaseAccountVerification) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *BaseAccountVerification) SetVendor(v string) {
	o.Vendor = v
}

func (o BaseAccountVerification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseAccountVerification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	toSerialize["status"] = o.Status
	toSerialize["vendor"] = o.Vendor

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseAccountVerification) UnmarshalJSON(data []byte) (err error) {
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

	varBaseAccountVerification := _BaseAccountVerification{}

	err = json.Unmarshal(data, &varBaseAccountVerification)

	if err != nil {
		return err
	}

	*o = BaseAccountVerification(varBaseAccountVerification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "creation_time")
		delete(additionalProperties, "last_updated_time")
		delete(additionalProperties, "status")
		delete(additionalProperties, "vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseAccountVerification struct {
	value *BaseAccountVerification
	isSet bool
}

func (v NullableBaseAccountVerification) Get() *BaseAccountVerification {
	return v.value
}

func (v *NullableBaseAccountVerification) Set(val *BaseAccountVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseAccountVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseAccountVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseAccountVerification(val *BaseAccountVerification) *NullableBaseAccountVerification {
	return &NullableBaseAccountVerification{value: val, isSet: true}
}

func (v NullableBaseAccountVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseAccountVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
