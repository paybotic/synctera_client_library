/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ScheduleConfig Payment schedule recurrence configuration
type ScheduleConfig struct {
	// Number of times to recur. Exactly one of end_date or count must be provided
	Count *int32 `json:"count,omitempty"`
	// End date of the schedule (exclusive). Exactly one of end_date or count must be provided
	EndDate *string `json:"end_date,omitempty"`
	Frequency string `json:"frequency"`
	// Interval between recurrences, e.g. interval = 2 with frequency = WEEKLY means every other week.
	Interval int32 `json:"interval"`
	// Start date of the schedule (inclusive)
	StartDate string `json:"start_date"`
}

// NewScheduleConfig instantiates a new ScheduleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleConfig(frequency string, interval int32, startDate string) *ScheduleConfig {
	this := ScheduleConfig{}
	this.Frequency = frequency
	this.Interval = interval
	this.StartDate = startDate
	return &this
}

// NewScheduleConfigWithDefaults instantiates a new ScheduleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleConfigWithDefaults() *ScheduleConfig {
	this := ScheduleConfig{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ScheduleConfig) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ScheduleConfig) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ScheduleConfig) SetCount(v int32) {
	o.Count = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ScheduleConfig) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ScheduleConfig) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ScheduleConfig) SetEndDate(v string) {
	o.EndDate = &v
}

// GetFrequency returns the Frequency field value
func (o *ScheduleConfig) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *ScheduleConfig) SetFrequency(v string) {
	o.Frequency = v
}

// GetInterval returns the Interval field value
func (o *ScheduleConfig) GetInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ScheduleConfig) SetInterval(v int32) {
	o.Interval = v
}

// GetStartDate returns the StartDate field value
func (o *ScheduleConfig) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *ScheduleConfig) SetStartDate(v string) {
	o.StartDate = v
}

func (o ScheduleConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["frequency"] = o.Frequency
	}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleConfig struct {
	value *ScheduleConfig
	isSet bool
}

func (v NullableScheduleConfig) Get() *ScheduleConfig {
	return v.value
}

func (v *NullableScheduleConfig) Set(val *ScheduleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleConfig(val *ScheduleConfig) *NullableScheduleConfig {
	return &NullableScheduleConfig{value: val, isSet: true}
}

func (v NullableScheduleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


