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
)

// checks if the EddAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EddAccount{}

// EddAccount struct for EddAccount
type EddAccount struct {
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
	Tenant *string `json:"tenant,omitempty"`
}

type _EddAccount EddAccount

// NewEddAccount instantiates a new EddAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEddAccount(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1) *EddAccount {
	this := EddAccount{}
	this.Reason = reason
	this.RelatedResourceId = relatedResourceId
	this.RelatedResourceType = relatedResourceType
	return &this
}

// NewEddAccountWithDefaults instantiates a new EddAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEddAccountWithDefaults() *EddAccount {
	this := EddAccount{}
	return &this
}

// GetAdditionalQuestions returns the AdditionalQuestions field value if set, zero value otherwise.
func (o *EddAccount) GetAdditionalQuestions() []Question {
	if o == nil || IsNil(o.AdditionalQuestions) {
		var ret []Question
		return ret
	}
	return o.AdditionalQuestions
}

// GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddAccount) GetAdditionalQuestionsOk() ([]Question, bool) {
	if o == nil || IsNil(o.AdditionalQuestions) {
		return nil, false
	}
	return o.AdditionalQuestions, true
}

// HasAdditionalQuestions returns a boolean if a field has been set.
func (o *EddAccount) HasAdditionalQuestions() bool {
	if o != nil && !IsNil(o.AdditionalQuestions) {
		return true
	}

	return false
}

// SetAdditionalQuestions gets a reference to the given []Question and assigns it to the AdditionalQuestions field.
func (o *EddAccount) SetAdditionalQuestions(v []Question) {
	o.AdditionalQuestions = v
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *EddAccount) GetCaseId() int32 {
	if o == nil || IsNil(o.CaseId) {
		var ret int32
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddAccount) GetCaseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *EddAccount) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given int32 and assigns it to the CaseId field.
func (o *EddAccount) SetCaseId(v int32) {
	o.CaseId = &v
}

// GetReason returns the Reason field value
func (o *EddAccount) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *EddAccount) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *EddAccount) SetReason(v string) {
	o.Reason = v
}

// GetRelatedResourceId returns the RelatedResourceId field value
func (o *EddAccount) GetRelatedResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelatedResourceId
}

// GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field value
// and a boolean to check if the value has been set.
func (o *EddAccount) GetRelatedResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceId, true
}

// SetRelatedResourceId sets field value
func (o *EddAccount) SetRelatedResourceId(v string) {
	o.RelatedResourceId = v
}

// GetRelatedResourceType returns the RelatedResourceType field value
func (o *EddAccount) GetRelatedResourceType() RelatedResourceType1 {
	if o == nil {
		var ret RelatedResourceType1
		return ret
	}

	return o.RelatedResourceType
}

// GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field value
// and a boolean to check if the value has been set.
func (o *EddAccount) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceType, true
}

// SetRelatedResourceType sets field value
func (o *EddAccount) SetRelatedResourceType(v RelatedResourceType1) {
	o.RelatedResourceType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *EddAccount) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddAccount) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *EddAccount) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *EddAccount) SetTenant(v string) {
	o.Tenant = &v
}

func (o EddAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EddAccount) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *EddAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reason",
		"related_resource_id",
		"related_resource_type",
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

	varEddAccount := _EddAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEddAccount)

	if err != nil {
		return err
	}

	*o = EddAccount(varEddAccount)

	return err
}

type NullableEddAccount struct {
	value *EddAccount
	isSet bool
}

func (v NullableEddAccount) Get() *EddAccount {
	return v.value
}

func (v *NullableEddAccount) Set(val *EddAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableEddAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableEddAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEddAccount(val *EddAccount) *NullableEddAccount {
	return &NullableEddAccount{value: val, isSet: true}
}

func (v NullableEddAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEddAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
