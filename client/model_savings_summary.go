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

// checks if the SavingsSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavingsSummary{}

// SavingsSummary A summary of the accrued interest for the saving account in the current period
type SavingsSummary struct {
	// The annual percentage yield (APY) for this account for this statement period, rounded to two decimal points. For example, an APY of 5.5% will display as 5.50.
	Apy *float32 `json:"apy,omitempty"`
	// The total interest earned by the depository account for this statement period in ISO 4217 minor currency units. For example, $1.50 USD of interest will be displayed as 150.
	InterestEarned *int64 `json:"interest_earned,omitempty"`
	// The total interest earned by the depository account in the previous statement period in ISO 4217 minor currency units. For example, $1.50 USD of interest will be displayed as 150.
	InterestEarnedPreviousMonth *int64 `json:"interest_earned_previous_month,omitempty"`
	// The total interest earned by the depository account in the previous year in ISO 4217 minor currency units. For example, $100 USD of interest will be displayed as 10000.
	InterestEarnedPreviousYear *int64 `json:"interest_earned_previous_year,omitempty"`
	// The total interest earned by the depository account for this year to date in ISO 4217 minor currency units. For example, $100 USD of interest will be displayed as 10000.
	InterestEarnedYtd    *int64 `json:"interest_earned_ytd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SavingsSummary SavingsSummary

// NewSavingsSummary instantiates a new SavingsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavingsSummary() *SavingsSummary {
	this := SavingsSummary{}
	return &this
}

// NewSavingsSummaryWithDefaults instantiates a new SavingsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavingsSummaryWithDefaults() *SavingsSummary {
	this := SavingsSummary{}
	return &this
}

// GetApy returns the Apy field value if set, zero value otherwise.
func (o *SavingsSummary) GetApy() float32 {
	if o == nil || IsNil(o.Apy) {
		var ret float32
		return ret
	}
	return *o.Apy
}

// GetApyOk returns a tuple with the Apy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavingsSummary) GetApyOk() (*float32, bool) {
	if o == nil || IsNil(o.Apy) {
		return nil, false
	}
	return o.Apy, true
}

// HasApy returns a boolean if a field has been set.
func (o *SavingsSummary) HasApy() bool {
	if o != nil && !IsNil(o.Apy) {
		return true
	}

	return false
}

// SetApy gets a reference to the given float32 and assigns it to the Apy field.
func (o *SavingsSummary) SetApy(v float32) {
	o.Apy = &v
}

// GetInterestEarned returns the InterestEarned field value if set, zero value otherwise.
func (o *SavingsSummary) GetInterestEarned() int64 {
	if o == nil || IsNil(o.InterestEarned) {
		var ret int64
		return ret
	}
	return *o.InterestEarned
}

// GetInterestEarnedOk returns a tuple with the InterestEarned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavingsSummary) GetInterestEarnedOk() (*int64, bool) {
	if o == nil || IsNil(o.InterestEarned) {
		return nil, false
	}
	return o.InterestEarned, true
}

// HasInterestEarned returns a boolean if a field has been set.
func (o *SavingsSummary) HasInterestEarned() bool {
	if o != nil && !IsNil(o.InterestEarned) {
		return true
	}

	return false
}

// SetInterestEarned gets a reference to the given int64 and assigns it to the InterestEarned field.
func (o *SavingsSummary) SetInterestEarned(v int64) {
	o.InterestEarned = &v
}

// GetInterestEarnedPreviousMonth returns the InterestEarnedPreviousMonth field value if set, zero value otherwise.
func (o *SavingsSummary) GetInterestEarnedPreviousMonth() int64 {
	if o == nil || IsNil(o.InterestEarnedPreviousMonth) {
		var ret int64
		return ret
	}
	return *o.InterestEarnedPreviousMonth
}

// GetInterestEarnedPreviousMonthOk returns a tuple with the InterestEarnedPreviousMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavingsSummary) GetInterestEarnedPreviousMonthOk() (*int64, bool) {
	if o == nil || IsNil(o.InterestEarnedPreviousMonth) {
		return nil, false
	}
	return o.InterestEarnedPreviousMonth, true
}

// HasInterestEarnedPreviousMonth returns a boolean if a field has been set.
func (o *SavingsSummary) HasInterestEarnedPreviousMonth() bool {
	if o != nil && !IsNil(o.InterestEarnedPreviousMonth) {
		return true
	}

	return false
}

// SetInterestEarnedPreviousMonth gets a reference to the given int64 and assigns it to the InterestEarnedPreviousMonth field.
func (o *SavingsSummary) SetInterestEarnedPreviousMonth(v int64) {
	o.InterestEarnedPreviousMonth = &v
}

// GetInterestEarnedPreviousYear returns the InterestEarnedPreviousYear field value if set, zero value otherwise.
func (o *SavingsSummary) GetInterestEarnedPreviousYear() int64 {
	if o == nil || IsNil(o.InterestEarnedPreviousYear) {
		var ret int64
		return ret
	}
	return *o.InterestEarnedPreviousYear
}

// GetInterestEarnedPreviousYearOk returns a tuple with the InterestEarnedPreviousYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavingsSummary) GetInterestEarnedPreviousYearOk() (*int64, bool) {
	if o == nil || IsNil(o.InterestEarnedPreviousYear) {
		return nil, false
	}
	return o.InterestEarnedPreviousYear, true
}

// HasInterestEarnedPreviousYear returns a boolean if a field has been set.
func (o *SavingsSummary) HasInterestEarnedPreviousYear() bool {
	if o != nil && !IsNil(o.InterestEarnedPreviousYear) {
		return true
	}

	return false
}

// SetInterestEarnedPreviousYear gets a reference to the given int64 and assigns it to the InterestEarnedPreviousYear field.
func (o *SavingsSummary) SetInterestEarnedPreviousYear(v int64) {
	o.InterestEarnedPreviousYear = &v
}

// GetInterestEarnedYtd returns the InterestEarnedYtd field value if set, zero value otherwise.
func (o *SavingsSummary) GetInterestEarnedYtd() int64 {
	if o == nil || IsNil(o.InterestEarnedYtd) {
		var ret int64
		return ret
	}
	return *o.InterestEarnedYtd
}

// GetInterestEarnedYtdOk returns a tuple with the InterestEarnedYtd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavingsSummary) GetInterestEarnedYtdOk() (*int64, bool) {
	if o == nil || IsNil(o.InterestEarnedYtd) {
		return nil, false
	}
	return o.InterestEarnedYtd, true
}

// HasInterestEarnedYtd returns a boolean if a field has been set.
func (o *SavingsSummary) HasInterestEarnedYtd() bool {
	if o != nil && !IsNil(o.InterestEarnedYtd) {
		return true
	}

	return false
}

// SetInterestEarnedYtd gets a reference to the given int64 and assigns it to the InterestEarnedYtd field.
func (o *SavingsSummary) SetInterestEarnedYtd(v int64) {
	o.InterestEarnedYtd = &v
}

func (o SavingsSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavingsSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apy) {
		toSerialize["apy"] = o.Apy
	}
	if !IsNil(o.InterestEarned) {
		toSerialize["interest_earned"] = o.InterestEarned
	}
	if !IsNil(o.InterestEarnedPreviousMonth) {
		toSerialize["interest_earned_previous_month"] = o.InterestEarnedPreviousMonth
	}
	if !IsNil(o.InterestEarnedPreviousYear) {
		toSerialize["interest_earned_previous_year"] = o.InterestEarnedPreviousYear
	}
	if !IsNil(o.InterestEarnedYtd) {
		toSerialize["interest_earned_ytd"] = o.InterestEarnedYtd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SavingsSummary) UnmarshalJSON(data []byte) (err error) {
	varSavingsSummary := _SavingsSummary{}

	err = json.Unmarshal(data, &varSavingsSummary)

	if err != nil {
		return err
	}

	*o = SavingsSummary(varSavingsSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apy")
		delete(additionalProperties, "interest_earned")
		delete(additionalProperties, "interest_earned_previous_month")
		delete(additionalProperties, "interest_earned_previous_year")
		delete(additionalProperties, "interest_earned_ytd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSavingsSummary struct {
	value *SavingsSummary
	isSet bool
}

func (v NullableSavingsSummary) Get() *SavingsSummary {
	return v.value
}

func (v *NullableSavingsSummary) Set(val *SavingsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableSavingsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSavingsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavingsSummary(val *SavingsSummary) *NullableSavingsSummary {
	return &NullableSavingsSummary{value: val, isSet: true}
}

func (v NullableSavingsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavingsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
