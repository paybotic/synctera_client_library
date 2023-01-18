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

// WaitlistAllOf Waitlist
type WaitlistAllOf struct {
	// Creation time
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Waitlist ID
	Id *string `json:"id,omitempty"`
	// Most recent updated time
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Current number of prospects in this waitlist with a status of admitted
	NumAdmitted *int32 `json:"num_admitted,omitempty"`
	// Current number of prospects in this waitlist with a status of created
	NumCreated *int32 `json:"num_created,omitempty"`
	// Current number of prospects in this waitlist, in any state
	NumProspects *int32 `json:"num_prospects,omitempty"`
	// Current number of prospects in this waitlist with a status of verified
	NumVerified *int32 `json:"num_verified,omitempty"`
	// Current number of prospects in this waitlist with a status of withdrawn
	NumWithdrawn *int32 `json:"num_withdrawn,omitempty"`
}

// NewWaitlistAllOf instantiates a new WaitlistAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaitlistAllOf() *WaitlistAllOf {
	this := WaitlistAllOf{}
	return &this
}

// NewWaitlistAllOfWithDefaults instantiates a new WaitlistAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaitlistAllOfWithDefaults() *WaitlistAllOf {
	this := WaitlistAllOf{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *WaitlistAllOf) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistAllOf) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *WaitlistAllOf) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *WaitlistAllOf) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WaitlistAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WaitlistAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WaitlistAllOf) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *WaitlistAllOf) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistAllOf) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *WaitlistAllOf) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *WaitlistAllOf) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetNumAdmitted returns the NumAdmitted field value if set, zero value otherwise.
func (o *WaitlistAllOf) GetNumAdmitted() int32 {
	if o == nil || o.NumAdmitted == nil {
		var ret int32
		return ret
	}
	return *o.NumAdmitted
}

// GetNumAdmittedOk returns a tuple with the NumAdmitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistAllOf) GetNumAdmittedOk() (*int32, bool) {
	if o == nil || o.NumAdmitted == nil {
		return nil, false
	}
	return o.NumAdmitted, true
}

// HasNumAdmitted returns a boolean if a field has been set.
func (o *WaitlistAllOf) HasNumAdmitted() bool {
	if o != nil && o.NumAdmitted != nil {
		return true
	}

	return false
}

// SetNumAdmitted gets a reference to the given int32 and assigns it to the NumAdmitted field.
func (o *WaitlistAllOf) SetNumAdmitted(v int32) {
	o.NumAdmitted = &v
}

// GetNumCreated returns the NumCreated field value if set, zero value otherwise.
func (o *WaitlistAllOf) GetNumCreated() int32 {
	if o == nil || o.NumCreated == nil {
		var ret int32
		return ret
	}
	return *o.NumCreated
}

// GetNumCreatedOk returns a tuple with the NumCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistAllOf) GetNumCreatedOk() (*int32, bool) {
	if o == nil || o.NumCreated == nil {
		return nil, false
	}
	return o.NumCreated, true
}

// HasNumCreated returns a boolean if a field has been set.
func (o *WaitlistAllOf) HasNumCreated() bool {
	if o != nil && o.NumCreated != nil {
		return true
	}

	return false
}

// SetNumCreated gets a reference to the given int32 and assigns it to the NumCreated field.
func (o *WaitlistAllOf) SetNumCreated(v int32) {
	o.NumCreated = &v
}

// GetNumProspects returns the NumProspects field value if set, zero value otherwise.
func (o *WaitlistAllOf) GetNumProspects() int32 {
	if o == nil || o.NumProspects == nil {
		var ret int32
		return ret
	}
	return *o.NumProspects
}

// GetNumProspectsOk returns a tuple with the NumProspects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistAllOf) GetNumProspectsOk() (*int32, bool) {
	if o == nil || o.NumProspects == nil {
		return nil, false
	}
	return o.NumProspects, true
}

// HasNumProspects returns a boolean if a field has been set.
func (o *WaitlistAllOf) HasNumProspects() bool {
	if o != nil && o.NumProspects != nil {
		return true
	}

	return false
}

// SetNumProspects gets a reference to the given int32 and assigns it to the NumProspects field.
func (o *WaitlistAllOf) SetNumProspects(v int32) {
	o.NumProspects = &v
}

// GetNumVerified returns the NumVerified field value if set, zero value otherwise.
func (o *WaitlistAllOf) GetNumVerified() int32 {
	if o == nil || o.NumVerified == nil {
		var ret int32
		return ret
	}
	return *o.NumVerified
}

// GetNumVerifiedOk returns a tuple with the NumVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistAllOf) GetNumVerifiedOk() (*int32, bool) {
	if o == nil || o.NumVerified == nil {
		return nil, false
	}
	return o.NumVerified, true
}

// HasNumVerified returns a boolean if a field has been set.
func (o *WaitlistAllOf) HasNumVerified() bool {
	if o != nil && o.NumVerified != nil {
		return true
	}

	return false
}

// SetNumVerified gets a reference to the given int32 and assigns it to the NumVerified field.
func (o *WaitlistAllOf) SetNumVerified(v int32) {
	o.NumVerified = &v
}

// GetNumWithdrawn returns the NumWithdrawn field value if set, zero value otherwise.
func (o *WaitlistAllOf) GetNumWithdrawn() int32 {
	if o == nil || o.NumWithdrawn == nil {
		var ret int32
		return ret
	}
	return *o.NumWithdrawn
}

// GetNumWithdrawnOk returns a tuple with the NumWithdrawn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistAllOf) GetNumWithdrawnOk() (*int32, bool) {
	if o == nil || o.NumWithdrawn == nil {
		return nil, false
	}
	return o.NumWithdrawn, true
}

// HasNumWithdrawn returns a boolean if a field has been set.
func (o *WaitlistAllOf) HasNumWithdrawn() bool {
	if o != nil && o.NumWithdrawn != nil {
		return true
	}

	return false
}

// SetNumWithdrawn gets a reference to the given int32 and assigns it to the NumWithdrawn field.
func (o *WaitlistAllOf) SetNumWithdrawn(v int32) {
	o.NumWithdrawn = &v
}

func (o WaitlistAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.NumAdmitted != nil {
		toSerialize["num_admitted"] = o.NumAdmitted
	}
	if o.NumCreated != nil {
		toSerialize["num_created"] = o.NumCreated
	}
	if o.NumProspects != nil {
		toSerialize["num_prospects"] = o.NumProspects
	}
	if o.NumVerified != nil {
		toSerialize["num_verified"] = o.NumVerified
	}
	if o.NumWithdrawn != nil {
		toSerialize["num_withdrawn"] = o.NumWithdrawn
	}
	return json.Marshal(toSerialize)
}

type NullableWaitlistAllOf struct {
	value *WaitlistAllOf
	isSet bool
}

func (v NullableWaitlistAllOf) Get() *WaitlistAllOf {
	return v.value
}

func (v *NullableWaitlistAllOf) Set(val *WaitlistAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWaitlistAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWaitlistAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaitlistAllOf(val *WaitlistAllOf) *NullableWaitlistAllOf {
	return &NullableWaitlistAllOf{value: val, isSet: true}
}

func (v NullableWaitlistAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaitlistAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


