/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// DebitNetworkResponse struct for DebitNetworkResponse
type DebitNetworkResponse struct {
	// indicates whether debit network is active
	Active bool `json:"active"`
	// The timestamp representing when the debit network was created
	CreationTime time.Time `json:"creation_time"`
	// The time when debit network became inactive
	EndDate *time.Time `json:"end_date,omitempty"`
	// Debit Network ID
	Id string `json:"id"`
	// The timestamp representing when the debit network was last modified
	LastModifiedTime time.Time `json:"last_modified_time"`
	// The name describing the debit network
	Name string `json:"name"`
	// The time when debit network goes live
	StartDate *time.Time `json:"start_date,omitempty"`
}

// NewDebitNetworkResponse instantiates a new DebitNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebitNetworkResponse(active bool, creationTime time.Time, id string, lastModifiedTime time.Time, name string) *DebitNetworkResponse {
	this := DebitNetworkResponse{}
	this.Active = active
	this.CreationTime = creationTime
	this.Id = id
	this.LastModifiedTime = lastModifiedTime
	this.Name = name
	return &this
}

// NewDebitNetworkResponseWithDefaults instantiates a new DebitNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebitNetworkResponseWithDefaults() *DebitNetworkResponse {
	this := DebitNetworkResponse{}
	return &this
}

// GetActive returns the Active field value
func (o *DebitNetworkResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *DebitNetworkResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *DebitNetworkResponse) SetActive(v bool) {
	o.Active = v
}

// GetCreationTime returns the CreationTime field value
func (o *DebitNetworkResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *DebitNetworkResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *DebitNetworkResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DebitNetworkResponse) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebitNetworkResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DebitNetworkResponse) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *DebitNetworkResponse) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetId returns the Id field value
func (o *DebitNetworkResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DebitNetworkResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DebitNetworkResponse) SetId(v string) {
	o.Id = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *DebitNetworkResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *DebitNetworkResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *DebitNetworkResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetName returns the Name field value
func (o *DebitNetworkResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DebitNetworkResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DebitNetworkResponse) SetName(v string) {
	o.Name = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DebitNetworkResponse) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebitNetworkResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DebitNetworkResponse) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *DebitNetworkResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o DebitNetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableDebitNetworkResponse struct {
	value *DebitNetworkResponse
	isSet bool
}

func (v NullableDebitNetworkResponse) Get() *DebitNetworkResponse {
	return v.value
}

func (v *NullableDebitNetworkResponse) Set(val *DebitNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDebitNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDebitNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebitNetworkResponse(val *DebitNetworkResponse) *NullableDebitNetworkResponse {
	return &NullableDebitNetworkResponse{value: val, isSet: true}
}

func (v NullableDebitNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebitNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
