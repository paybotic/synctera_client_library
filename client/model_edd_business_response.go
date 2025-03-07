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

// checks if the EddBusinessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EddBusinessResponse{}

// EddBusinessResponse struct for EddBusinessResponse
type EddBusinessResponse struct {
	// ISO-3166-1 Alpha-2 country code
	Country          *string           `json:"country,omitempty" validate:"regexp=^[A-Z]{2}$"`
	EstimatedRevenue *EstimatedRevenue `json:"estimated_revenue,omitempty"`
	IndustryType     *IndustryType     `json:"industry_type,omitempty"`
	// The number of negative news findings.
	NegativeNewsFindings *int32 `json:"negative_news_findings,omitempty"`
	// True if the customer is expected to send or receive wire transfers at a regular frequency.
	RecurringWireUsage  *bool                `json:"recurring_wire_usage,omitempty"`
	SpecificInvolvement *SpecificInvolvement `json:"specific_involvement,omitempty"`
	// Array of transaction volumes.
	TransactionVolume []TransactionVolume `json:"transaction_volume,omitempty"`
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

type _EddBusinessResponse EddBusinessResponse

// NewEddBusinessResponse instantiates a new EddBusinessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEddBusinessResponse(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, creationTime time.Time, deletionTime NullableTime, id string) *EddBusinessResponse {
	this := EddBusinessResponse{}
	this.Reason = reason
	this.RelatedResourceId = relatedResourceId
	this.RelatedResourceType = relatedResourceType
	this.CreationTime = creationTime
	this.DeletionTime = deletionTime
	this.Id = id
	return &this
}

// NewEddBusinessResponseWithDefaults instantiates a new EddBusinessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEddBusinessResponseWithDefaults() *EddBusinessResponse {
	this := EddBusinessResponse{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *EddBusinessResponse) SetCountry(v string) {
	o.Country = &v
}

// GetEstimatedRevenue returns the EstimatedRevenue field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetEstimatedRevenue() EstimatedRevenue {
	if o == nil || IsNil(o.EstimatedRevenue) {
		var ret EstimatedRevenue
		return ret
	}
	return *o.EstimatedRevenue
}

// GetEstimatedRevenueOk returns a tuple with the EstimatedRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetEstimatedRevenueOk() (*EstimatedRevenue, bool) {
	if o == nil || IsNil(o.EstimatedRevenue) {
		return nil, false
	}
	return o.EstimatedRevenue, true
}

// HasEstimatedRevenue returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasEstimatedRevenue() bool {
	if o != nil && !IsNil(o.EstimatedRevenue) {
		return true
	}

	return false
}

// SetEstimatedRevenue gets a reference to the given EstimatedRevenue and assigns it to the EstimatedRevenue field.
func (o *EddBusinessResponse) SetEstimatedRevenue(v EstimatedRevenue) {
	o.EstimatedRevenue = &v
}

// GetIndustryType returns the IndustryType field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetIndustryType() IndustryType {
	if o == nil || IsNil(o.IndustryType) {
		var ret IndustryType
		return ret
	}
	return *o.IndustryType
}

// GetIndustryTypeOk returns a tuple with the IndustryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetIndustryTypeOk() (*IndustryType, bool) {
	if o == nil || IsNil(o.IndustryType) {
		return nil, false
	}
	return o.IndustryType, true
}

// HasIndustryType returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasIndustryType() bool {
	if o != nil && !IsNil(o.IndustryType) {
		return true
	}

	return false
}

// SetIndustryType gets a reference to the given IndustryType and assigns it to the IndustryType field.
func (o *EddBusinessResponse) SetIndustryType(v IndustryType) {
	o.IndustryType = &v
}

// GetNegativeNewsFindings returns the NegativeNewsFindings field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetNegativeNewsFindings() int32 {
	if o == nil || IsNil(o.NegativeNewsFindings) {
		var ret int32
		return ret
	}
	return *o.NegativeNewsFindings
}

// GetNegativeNewsFindingsOk returns a tuple with the NegativeNewsFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetNegativeNewsFindingsOk() (*int32, bool) {
	if o == nil || IsNil(o.NegativeNewsFindings) {
		return nil, false
	}
	return o.NegativeNewsFindings, true
}

// HasNegativeNewsFindings returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasNegativeNewsFindings() bool {
	if o != nil && !IsNil(o.NegativeNewsFindings) {
		return true
	}

	return false
}

// SetNegativeNewsFindings gets a reference to the given int32 and assigns it to the NegativeNewsFindings field.
func (o *EddBusinessResponse) SetNegativeNewsFindings(v int32) {
	o.NegativeNewsFindings = &v
}

// GetRecurringWireUsage returns the RecurringWireUsage field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetRecurringWireUsage() bool {
	if o == nil || IsNil(o.RecurringWireUsage) {
		var ret bool
		return ret
	}
	return *o.RecurringWireUsage
}

// GetRecurringWireUsageOk returns a tuple with the RecurringWireUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetRecurringWireUsageOk() (*bool, bool) {
	if o == nil || IsNil(o.RecurringWireUsage) {
		return nil, false
	}
	return o.RecurringWireUsage, true
}

// HasRecurringWireUsage returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasRecurringWireUsage() bool {
	if o != nil && !IsNil(o.RecurringWireUsage) {
		return true
	}

	return false
}

// SetRecurringWireUsage gets a reference to the given bool and assigns it to the RecurringWireUsage field.
func (o *EddBusinessResponse) SetRecurringWireUsage(v bool) {
	o.RecurringWireUsage = &v
}

// GetSpecificInvolvement returns the SpecificInvolvement field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetSpecificInvolvement() SpecificInvolvement {
	if o == nil || IsNil(o.SpecificInvolvement) {
		var ret SpecificInvolvement
		return ret
	}
	return *o.SpecificInvolvement
}

// GetSpecificInvolvementOk returns a tuple with the SpecificInvolvement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetSpecificInvolvementOk() (*SpecificInvolvement, bool) {
	if o == nil || IsNil(o.SpecificInvolvement) {
		return nil, false
	}
	return o.SpecificInvolvement, true
}

// HasSpecificInvolvement returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasSpecificInvolvement() bool {
	if o != nil && !IsNil(o.SpecificInvolvement) {
		return true
	}

	return false
}

// SetSpecificInvolvement gets a reference to the given SpecificInvolvement and assigns it to the SpecificInvolvement field.
func (o *EddBusinessResponse) SetSpecificInvolvement(v SpecificInvolvement) {
	o.SpecificInvolvement = &v
}

// GetTransactionVolume returns the TransactionVolume field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetTransactionVolume() []TransactionVolume {
	if o == nil || IsNil(o.TransactionVolume) {
		var ret []TransactionVolume
		return ret
	}
	return o.TransactionVolume
}

// GetTransactionVolumeOk returns a tuple with the TransactionVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetTransactionVolumeOk() ([]TransactionVolume, bool) {
	if o == nil || IsNil(o.TransactionVolume) {
		return nil, false
	}
	return o.TransactionVolume, true
}

// HasTransactionVolume returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasTransactionVolume() bool {
	if o != nil && !IsNil(o.TransactionVolume) {
		return true
	}

	return false
}

// SetTransactionVolume gets a reference to the given []TransactionVolume and assigns it to the TransactionVolume field.
func (o *EddBusinessResponse) SetTransactionVolume(v []TransactionVolume) {
	o.TransactionVolume = v
}

// GetAdditionalQuestions returns the AdditionalQuestions field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetAdditionalQuestions() []Question {
	if o == nil || IsNil(o.AdditionalQuestions) {
		var ret []Question
		return ret
	}
	return o.AdditionalQuestions
}

// GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetAdditionalQuestionsOk() ([]Question, bool) {
	if o == nil || IsNil(o.AdditionalQuestions) {
		return nil, false
	}
	return o.AdditionalQuestions, true
}

// HasAdditionalQuestions returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasAdditionalQuestions() bool {
	if o != nil && !IsNil(o.AdditionalQuestions) {
		return true
	}

	return false
}

// SetAdditionalQuestions gets a reference to the given []Question and assigns it to the AdditionalQuestions field.
func (o *EddBusinessResponse) SetAdditionalQuestions(v []Question) {
	o.AdditionalQuestions = v
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetCaseId() int32 {
	if o == nil || IsNil(o.CaseId) {
		var ret int32
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetCaseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given int32 and assigns it to the CaseId field.
func (o *EddBusinessResponse) SetCaseId(v int32) {
	o.CaseId = &v
}

// GetReason returns the Reason field value
func (o *EddBusinessResponse) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *EddBusinessResponse) SetReason(v string) {
	o.Reason = v
}

// GetRelatedResourceId returns the RelatedResourceId field value
func (o *EddBusinessResponse) GetRelatedResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelatedResourceId
}

// GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field value
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetRelatedResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceId, true
}

// SetRelatedResourceId sets field value
func (o *EddBusinessResponse) SetRelatedResourceId(v string) {
	o.RelatedResourceId = v
}

// GetRelatedResourceType returns the RelatedResourceType field value
func (o *EddBusinessResponse) GetRelatedResourceType() RelatedResourceType1 {
	if o == nil {
		var ret RelatedResourceType1
		return ret
	}

	return o.RelatedResourceType
}

// GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field value
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceType, true
}

// SetRelatedResourceType sets field value
func (o *EddBusinessResponse) SetRelatedResourceType(v RelatedResourceType1) {
	o.RelatedResourceType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *EddBusinessResponse) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *EddBusinessResponse) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *EddBusinessResponse) SetTenant(v string) {
	o.Tenant = &v
}

// GetCreationTime returns the CreationTime field value
func (o *EddBusinessResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *EddBusinessResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetDeletionTime returns the DeletionTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *EddBusinessResponse) GetDeletionTime() time.Time {
	if o == nil || o.DeletionTime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DeletionTime.Get()
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EddBusinessResponse) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletionTime.Get(), o.DeletionTime.IsSet()
}

// SetDeletionTime sets field value
func (o *EddBusinessResponse) SetDeletionTime(v time.Time) {
	o.DeletionTime.Set(&v)
}

// GetId returns the Id field value
func (o *EddBusinessResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EddBusinessResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EddBusinessResponse) SetId(v string) {
	o.Id = v
}

func (o EddBusinessResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EddBusinessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.EstimatedRevenue) {
		toSerialize["estimated_revenue"] = o.EstimatedRevenue
	}
	if !IsNil(o.IndustryType) {
		toSerialize["industry_type"] = o.IndustryType
	}
	if !IsNil(o.NegativeNewsFindings) {
		toSerialize["negative_news_findings"] = o.NegativeNewsFindings
	}
	if !IsNil(o.RecurringWireUsage) {
		toSerialize["recurring_wire_usage"] = o.RecurringWireUsage
	}
	if !IsNil(o.SpecificInvolvement) {
		toSerialize["specific_involvement"] = o.SpecificInvolvement
	}
	if !IsNil(o.TransactionVolume) {
		toSerialize["transaction_volume"] = o.TransactionVolume
	}
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

func (o *EddBusinessResponse) UnmarshalJSON(data []byte) (err error) {
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

	varEddBusinessResponse := _EddBusinessResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEddBusinessResponse)

	if err != nil {
		return err
	}

	*o = EddBusinessResponse(varEddBusinessResponse)

	return err
}

type NullableEddBusinessResponse struct {
	value *EddBusinessResponse
	isSet bool
}

func (v NullableEddBusinessResponse) Get() *EddBusinessResponse {
	return v.value
}

func (v *NullableEddBusinessResponse) Set(val *EddBusinessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEddBusinessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEddBusinessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEddBusinessResponse(val *EddBusinessResponse) *NullableEddBusinessResponse {
	return &NullableEddBusinessResponse{value: val, isSet: true}
}

func (v NullableEddBusinessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEddBusinessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
