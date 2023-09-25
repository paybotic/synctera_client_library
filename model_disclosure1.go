/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Disclosure1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Disclosure1{}

// Disclosure1 Represents a disclosure
type Disclosure1 struct {
	CreationTime *time.Time `json:"creation_time,omitempty"`
	EventType string `json:"event_type"`
	// Disclosure ID
	Id *string `json:"id,omitempty"`
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Date of disclosure
	Timestamp time.Time `json:"timestamp"`
	// Disclosure Type
	Type string `json:"type"`
	// Disclosure Version
	Version string `json:"version"`
}

// NewDisclosure1 instantiates a new Disclosure1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisclosure1(eventType string, timestamp time.Time, type_ string, version string) *Disclosure1 {
	this := Disclosure1{}
	this.EventType = eventType
	this.Timestamp = timestamp
	this.Type = type_
	this.Version = version
	return &this
}

// NewDisclosure1WithDefaults instantiates a new Disclosure1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisclosure1WithDefaults() *Disclosure1 {
	this := Disclosure1{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Disclosure1) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosure1) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Disclosure1) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Disclosure1) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEventType returns the EventType field value
func (o *Disclosure1) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *Disclosure1) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *Disclosure1) SetEventType(v string) {
	o.EventType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Disclosure1) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosure1) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Disclosure1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Disclosure1) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Disclosure1) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosure1) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Disclosure1) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Disclosure1) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetTimestamp returns the Timestamp field value
func (o *Disclosure1) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Disclosure1) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Disclosure1) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetType returns the Type field value
func (o *Disclosure1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Disclosure1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Disclosure1) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *Disclosure1) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Disclosure1) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Disclosure1) SetVersion(v string) {
	o.Version = v
}

func (o Disclosure1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Disclosure1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	toSerialize["event_type"] = o.EventType
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableDisclosure1 struct {
	value *Disclosure1
	isSet bool
}

func (v NullableDisclosure1) Get() *Disclosure1 {
	return v.value
}

func (v *NullableDisclosure1) Set(val *Disclosure1) {
	v.value = val
	v.isSet = true
}

func (v NullableDisclosure1) IsSet() bool {
	return v.isSet
}

func (v *NullableDisclosure1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisclosure1(val *Disclosure1) *NullableDisclosure1 {
	return &NullableDisclosure1{value: val, isSet: true}
}

func (v NullableDisclosure1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisclosure1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


