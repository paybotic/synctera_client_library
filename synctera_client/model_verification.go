/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Verification Verification result.
type Verification struct {
	// Unique ID for the business. Exactly one of `business_id` or `person_id` must be set. 
	BusinessId *string `json:"business_id,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// A list of individual checks done as part of the due diligence process  for the verification type. 
	Details []Detail `json:"details,omitempty"`
	// Unique ID for this verification result.
	Id *string `json:"id,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Unique ID for the person. Exactly one of `person_id` or `business_id` must be set. 
	PersonId *string `json:"person_id,omitempty"`
	Result VerificationResult `json:"result"`
	VendorInfo *VendorInfo `json:"vendor_info,omitempty"`
	// The date and time the verification was completed.
	VerificationTime time.Time `json:"verification_time"`
	VerificationType VerificationType1 `json:"verification_type"`
}

// NewVerification instantiates a new Verification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerification(result VerificationResult, verificationTime time.Time, verificationType VerificationType1) *Verification {
	this := Verification{}
	this.Result = result
	this.VerificationTime = verificationTime
	this.VerificationType = verificationType
	return &this
}

// NewVerificationWithDefaults instantiates a new Verification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationWithDefaults() *Verification {
	this := Verification{}
	return &this
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *Verification) GetBusinessId() string {
	if o == nil || o.BusinessId == nil {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Verification) GetBusinessIdOk() (*string, bool) {
	if o == nil || o.BusinessId == nil {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *Verification) HasBusinessId() bool {
	if o != nil && o.BusinessId != nil {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *Verification) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Verification) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Verification) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Verification) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Verification) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Verification) GetDetails() []Detail {
	if o == nil || o.Details == nil {
		var ret []Detail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Verification) GetDetailsOk() ([]Detail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Verification) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []Detail and assigns it to the Details field.
func (o *Verification) SetDetails(v []Detail) {
	o.Details = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Verification) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Verification) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Verification) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Verification) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Verification) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Verification) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Verification) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Verification) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Verification) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Verification) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Verification) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Verification) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *Verification) GetPersonId() string {
	if o == nil || o.PersonId == nil {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Verification) GetPersonIdOk() (*string, bool) {
	if o == nil || o.PersonId == nil {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *Verification) HasPersonId() bool {
	if o != nil && o.PersonId != nil {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
func (o *Verification) SetPersonId(v string) {
	o.PersonId = &v
}

// GetResult returns the Result field value
func (o *Verification) GetResult() VerificationResult {
	if o == nil {
		var ret VerificationResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *Verification) GetResultOk() (*VerificationResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *Verification) SetResult(v VerificationResult) {
	o.Result = v
}

// GetVendorInfo returns the VendorInfo field value if set, zero value otherwise.
func (o *Verification) GetVendorInfo() VendorInfo {
	if o == nil || o.VendorInfo == nil {
		var ret VendorInfo
		return ret
	}
	return *o.VendorInfo
}

// GetVendorInfoOk returns a tuple with the VendorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Verification) GetVendorInfoOk() (*VendorInfo, bool) {
	if o == nil || o.VendorInfo == nil {
		return nil, false
	}
	return o.VendorInfo, true
}

// HasVendorInfo returns a boolean if a field has been set.
func (o *Verification) HasVendorInfo() bool {
	if o != nil && o.VendorInfo != nil {
		return true
	}

	return false
}

// SetVendorInfo gets a reference to the given VendorInfo and assigns it to the VendorInfo field.
func (o *Verification) SetVendorInfo(v VendorInfo) {
	o.VendorInfo = &v
}

// GetVerificationTime returns the VerificationTime field value
func (o *Verification) GetVerificationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.VerificationTime
}

// GetVerificationTimeOk returns a tuple with the VerificationTime field value
// and a boolean to check if the value has been set.
func (o *Verification) GetVerificationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationTime, true
}

// SetVerificationTime sets field value
func (o *Verification) SetVerificationTime(v time.Time) {
	o.VerificationTime = v
}

// GetVerificationType returns the VerificationType field value
func (o *Verification) GetVerificationType() VerificationType1 {
	if o == nil {
		var ret VerificationType1
		return ret
	}

	return o.VerificationType
}

// GetVerificationTypeOk returns a tuple with the VerificationType field value
// and a boolean to check if the value has been set.
func (o *Verification) GetVerificationTypeOk() (*VerificationType1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationType, true
}

// SetVerificationType sets field value
func (o *Verification) SetVerificationType(v VerificationType1) {
	o.VerificationType = v
}

func (o Verification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BusinessId != nil {
		toSerialize["business_id"] = o.BusinessId
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PersonId != nil {
		toSerialize["person_id"] = o.PersonId
	}
	if true {
		toSerialize["result"] = o.Result
	}
	if o.VendorInfo != nil {
		toSerialize["vendor_info"] = o.VendorInfo
	}
	if true {
		toSerialize["verification_time"] = o.VerificationTime
	}
	if true {
		toSerialize["verification_type"] = o.VerificationType
	}
	return json.Marshal(toSerialize)
}

type NullableVerification struct {
	value *Verification
	isSet bool
}

func (v NullableVerification) Get() *Verification {
	return v.value
}

func (v *NullableVerification) Set(val *Verification) {
	v.value = val
	v.isSet = true
}

func (v NullableVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerification(val *Verification) *NullableVerification {
	return &NullableVerification{value: val, isSet: true}
}

func (v NullableVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


