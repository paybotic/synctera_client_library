/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the EddAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EddAccountResponse{}

// EddAccountResponse struct for EddAccountResponse
type EddAccountResponse struct {
	// Additional questions regarding the related resource
	AdditionalQuestions []Question `json:"additional_questions,omitempty"`
	// The ID of the case related to this EDD record
	CaseId *int32 `json:"case_id,omitempty"`
	// The reason for this EDD record to be requested
	Reason string `json:"reason"`
	// related resource UUID
	RelatedResourceId   string               `json:"related_resource_id"`
	RelatedResourceType RelatedResourceType1 `json:"related_resource_type"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant       *string      `json:"tenant,omitempty"`
	CreationTime time.Time    `json:"creation_time"`
	DeletionTime NullableTime `json:"deletion_time"`
	// EDD record unique identifier
	Id string `json:"id"`
}

type _EddAccountResponse EddAccountResponse

// NewEddAccountResponse instantiates a new EddAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEddAccountResponse(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, creationTime time.Time, deletionTime NullableTime, id string) *EddAccountResponse {
	this := EddAccountResponse{}
	this.Reason = reason
	this.RelatedResourceId = relatedResourceId
	this.RelatedResourceType = relatedResourceType
	this.CreationTime = creationTime
	this.DeletionTime = deletionTime
	this.Id = id
	return &this
}

// NewEddAccountResponseWithDefaults instantiates a new EddAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEddAccountResponseWithDefaults() *EddAccountResponse {
	this := EddAccountResponse{}
	return &this
}

// GetAdditionalQuestions returns the AdditionalQuestions field value if set, zero value otherwise.
func (o *EddAccountResponse) GetAdditionalQuestions() []Question {
	if o == nil || IsNil(o.AdditionalQuestions) {
		var ret []Question
		return ret
	}
	return o.AdditionalQuestions
}

// GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddAccountResponse) GetAdditionalQuestionsOk() ([]Question, bool) {
	if o == nil || IsNil(o.AdditionalQuestions) {
		return nil, false
	}
	return o.AdditionalQuestions, true
}

// HasAdditionalQuestions returns a boolean if a field has been set.
func (o *EddAccountResponse) HasAdditionalQuestions() bool {
	if o != nil && !IsNil(o.AdditionalQuestions) {
		return true
	}

	return false
}

// SetAdditionalQuestions gets a reference to the given []Question and assigns it to the AdditionalQuestions field.
func (o *EddAccountResponse) SetAdditionalQuestions(v []Question) {
	o.AdditionalQuestions = v
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *EddAccountResponse) GetCaseId() int32 {
	if o == nil || IsNil(o.CaseId) {
		var ret int32
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddAccountResponse) GetCaseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *EddAccountResponse) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given int32 and assigns it to the CaseId field.
func (o *EddAccountResponse) SetCaseId(v int32) {
	o.CaseId = &v
}

// GetReason returns the Reason field value
func (o *EddAccountResponse) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *EddAccountResponse) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *EddAccountResponse) SetReason(v string) {
	o.Reason = v
}

// GetRelatedResourceId returns the RelatedResourceId field value
func (o *EddAccountResponse) GetRelatedResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelatedResourceId
}

// GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field value
// and a boolean to check if the value has been set.
func (o *EddAccountResponse) GetRelatedResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceId, true
}

// SetRelatedResourceId sets field value
func (o *EddAccountResponse) SetRelatedResourceId(v string) {
	o.RelatedResourceId = v
}

// GetRelatedResourceType returns the RelatedResourceType field value
func (o *EddAccountResponse) GetRelatedResourceType() RelatedResourceType1 {
	if o == nil {
		var ret RelatedResourceType1
		return ret
	}

	return o.RelatedResourceType
}

// GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field value
// and a boolean to check if the value has been set.
func (o *EddAccountResponse) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceType, true
}

// SetRelatedResourceType sets field value
func (o *EddAccountResponse) SetRelatedResourceType(v RelatedResourceType1) {
	o.RelatedResourceType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *EddAccountResponse) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddAccountResponse) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *EddAccountResponse) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *EddAccountResponse) SetTenant(v string) {
	o.Tenant = &v
}

// GetCreationTime returns the CreationTime field value
func (o *EddAccountResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *EddAccountResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *EddAccountResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetDeletionTime returns the DeletionTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *EddAccountResponse) GetDeletionTime() time.Time {
	if o == nil || o.DeletionTime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DeletionTime.Get()
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EddAccountResponse) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletionTime.Get(), o.DeletionTime.IsSet()
}

// SetDeletionTime sets field value
func (o *EddAccountResponse) SetDeletionTime(v time.Time) {
	o.DeletionTime.Set(&v)
}

// GetId returns the Id field value
func (o *EddAccountResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EddAccountResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EddAccountResponse) SetId(v string) {
	o.Id = v
}

func (o EddAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EddAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalQuestions) {
		toSerialize["additional_questions"] = o.AdditionalQuestions
	}
	if !IsNil(o.CaseId) {
		toSerialize["case_id"] = o.CaseId
	}
	toSerialize["reason"] = o.Reason
	toSerialize["related_resource_id"] = o.RelatedResourceId
	toSerialize["related_resource_type"] = o.RelatedResourceType
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	toSerialize["creation_time"] = o.CreationTime
	toSerialize["deletion_time"] = o.DeletionTime.Get()
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *EddAccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reason",
		"related_resource_id",
		"related_resource_type",
		"creation_time",
		"deletion_time",
		"id",
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

	varEddAccountResponse := _EddAccountResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEddAccountResponse)

	if err != nil {
		return err
	}

	*o = EddAccountResponse(varEddAccountResponse)

	return err
}

type NullableEddAccountResponse struct {
	value *EddAccountResponse
	isSet bool
}

func (v NullableEddAccountResponse) Get() *EddAccountResponse {
	return v.value
}

func (v *NullableEddAccountResponse) Set(val *EddAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEddAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEddAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEddAccountResponse(val *EddAccountResponse) *NullableEddAccountResponse {
	return &NullableEddAccountResponse{value: val, isSet: true}
}

func (v NullableEddAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEddAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
