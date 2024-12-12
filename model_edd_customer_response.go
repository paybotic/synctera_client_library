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

// checks if the EddCustomerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EddCustomerResponse{}

// EddCustomerResponse struct for EddCustomerResponse
type EddCustomerResponse struct {
	// List of countries where the related customer holds citizenship.
	CitizenshipCountries []string `json:"citizenship_countries,omitempty"`
	// The type of employment.
	EmploymentType *string `json:"employment_type,omitempty"`
	Income         *Income `json:"income,omitempty"`
	// The number of negative news findings.
	NegativeNewsFindings *int32 `json:"negative_news_findings,omitempty"`
	// The occupation of the related resource.
	Occupation         *string       `json:"occupation,omitempty"`
	OccupationIndustry *IndustryType `json:"occupation_industry,omitempty"`
	// True if the customer is expected to use direct deposit at a regular frequency.
	RecurringDirectDeposit *bool `json:"recurring_direct_deposit,omitempty"`
	// The type of residence.
	ResidenceType      *string             `json:"residence_type,omitempty"`
	ResidentialExpense *ResidentialExpense `json:"residential_expense,omitempty"`
	// The sources of wealth for the customer.
	SourceOfWealth []SourceOfWealth `json:"source_of_wealth,omitempty"`
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

type _EddCustomerResponse EddCustomerResponse

// NewEddCustomerResponse instantiates a new EddCustomerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEddCustomerResponse(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1, creationTime time.Time, deletionTime NullableTime, id string) *EddCustomerResponse {
	this := EddCustomerResponse{}
	this.Reason = reason
	this.RelatedResourceId = relatedResourceId
	this.RelatedResourceType = relatedResourceType
	this.CreationTime = creationTime
	this.DeletionTime = deletionTime
	this.Id = id
	return &this
}

// NewEddCustomerResponseWithDefaults instantiates a new EddCustomerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEddCustomerResponseWithDefaults() *EddCustomerResponse {
	this := EddCustomerResponse{}
	return &this
}

// GetCitizenshipCountries returns the CitizenshipCountries field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetCitizenshipCountries() []string {
	if o == nil || IsNil(o.CitizenshipCountries) {
		var ret []string
		return ret
	}
	return o.CitizenshipCountries
}

// GetCitizenshipCountriesOk returns a tuple with the CitizenshipCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetCitizenshipCountriesOk() ([]string, bool) {
	if o == nil || IsNil(o.CitizenshipCountries) {
		return nil, false
	}
	return o.CitizenshipCountries, true
}

// HasCitizenshipCountries returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasCitizenshipCountries() bool {
	if o != nil && !IsNil(o.CitizenshipCountries) {
		return true
	}

	return false
}

// SetCitizenshipCountries gets a reference to the given []string and assigns it to the CitizenshipCountries field.
func (o *EddCustomerResponse) SetCitizenshipCountries(v []string) {
	o.CitizenshipCountries = v
}

// GetEmploymentType returns the EmploymentType field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetEmploymentType() string {
	if o == nil || IsNil(o.EmploymentType) {
		var ret string
		return ret
	}
	return *o.EmploymentType
}

// GetEmploymentTypeOk returns a tuple with the EmploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetEmploymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmploymentType) {
		return nil, false
	}
	return o.EmploymentType, true
}

// HasEmploymentType returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasEmploymentType() bool {
	if o != nil && !IsNil(o.EmploymentType) {
		return true
	}

	return false
}

// SetEmploymentType gets a reference to the given string and assigns it to the EmploymentType field.
func (o *EddCustomerResponse) SetEmploymentType(v string) {
	o.EmploymentType = &v
}

// GetIncome returns the Income field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetIncome() Income {
	if o == nil || IsNil(o.Income) {
		var ret Income
		return ret
	}
	return *o.Income
}

// GetIncomeOk returns a tuple with the Income field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetIncomeOk() (*Income, bool) {
	if o == nil || IsNil(o.Income) {
		return nil, false
	}
	return o.Income, true
}

// HasIncome returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasIncome() bool {
	if o != nil && !IsNil(o.Income) {
		return true
	}

	return false
}

// SetIncome gets a reference to the given Income and assigns it to the Income field.
func (o *EddCustomerResponse) SetIncome(v Income) {
	o.Income = &v
}

// GetNegativeNewsFindings returns the NegativeNewsFindings field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetNegativeNewsFindings() int32 {
	if o == nil || IsNil(o.NegativeNewsFindings) {
		var ret int32
		return ret
	}
	return *o.NegativeNewsFindings
}

// GetNegativeNewsFindingsOk returns a tuple with the NegativeNewsFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetNegativeNewsFindingsOk() (*int32, bool) {
	if o == nil || IsNil(o.NegativeNewsFindings) {
		return nil, false
	}
	return o.NegativeNewsFindings, true
}

// HasNegativeNewsFindings returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasNegativeNewsFindings() bool {
	if o != nil && !IsNil(o.NegativeNewsFindings) {
		return true
	}

	return false
}

// SetNegativeNewsFindings gets a reference to the given int32 and assigns it to the NegativeNewsFindings field.
func (o *EddCustomerResponse) SetNegativeNewsFindings(v int32) {
	o.NegativeNewsFindings = &v
}

// GetOccupation returns the Occupation field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetOccupation() string {
	if o == nil || IsNil(o.Occupation) {
		var ret string
		return ret
	}
	return *o.Occupation
}

// GetOccupationOk returns a tuple with the Occupation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetOccupationOk() (*string, bool) {
	if o == nil || IsNil(o.Occupation) {
		return nil, false
	}
	return o.Occupation, true
}

// HasOccupation returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasOccupation() bool {
	if o != nil && !IsNil(o.Occupation) {
		return true
	}

	return false
}

// SetOccupation gets a reference to the given string and assigns it to the Occupation field.
func (o *EddCustomerResponse) SetOccupation(v string) {
	o.Occupation = &v
}

// GetOccupationIndustry returns the OccupationIndustry field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetOccupationIndustry() IndustryType {
	if o == nil || IsNil(o.OccupationIndustry) {
		var ret IndustryType
		return ret
	}
	return *o.OccupationIndustry
}

// GetOccupationIndustryOk returns a tuple with the OccupationIndustry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetOccupationIndustryOk() (*IndustryType, bool) {
	if o == nil || IsNil(o.OccupationIndustry) {
		return nil, false
	}
	return o.OccupationIndustry, true
}

// HasOccupationIndustry returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasOccupationIndustry() bool {
	if o != nil && !IsNil(o.OccupationIndustry) {
		return true
	}

	return false
}

// SetOccupationIndustry gets a reference to the given IndustryType and assigns it to the OccupationIndustry field.
func (o *EddCustomerResponse) SetOccupationIndustry(v IndustryType) {
	o.OccupationIndustry = &v
}

// GetRecurringDirectDeposit returns the RecurringDirectDeposit field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetRecurringDirectDeposit() bool {
	if o == nil || IsNil(o.RecurringDirectDeposit) {
		var ret bool
		return ret
	}
	return *o.RecurringDirectDeposit
}

// GetRecurringDirectDepositOk returns a tuple with the RecurringDirectDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetRecurringDirectDepositOk() (*bool, bool) {
	if o == nil || IsNil(o.RecurringDirectDeposit) {
		return nil, false
	}
	return o.RecurringDirectDeposit, true
}

// HasRecurringDirectDeposit returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasRecurringDirectDeposit() bool {
	if o != nil && !IsNil(o.RecurringDirectDeposit) {
		return true
	}

	return false
}

// SetRecurringDirectDeposit gets a reference to the given bool and assigns it to the RecurringDirectDeposit field.
func (o *EddCustomerResponse) SetRecurringDirectDeposit(v bool) {
	o.RecurringDirectDeposit = &v
}

// GetResidenceType returns the ResidenceType field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetResidenceType() string {
	if o == nil || IsNil(o.ResidenceType) {
		var ret string
		return ret
	}
	return *o.ResidenceType
}

// GetResidenceTypeOk returns a tuple with the ResidenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetResidenceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResidenceType) {
		return nil, false
	}
	return o.ResidenceType, true
}

// HasResidenceType returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasResidenceType() bool {
	if o != nil && !IsNil(o.ResidenceType) {
		return true
	}

	return false
}

// SetResidenceType gets a reference to the given string and assigns it to the ResidenceType field.
func (o *EddCustomerResponse) SetResidenceType(v string) {
	o.ResidenceType = &v
}

// GetResidentialExpense returns the ResidentialExpense field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetResidentialExpense() ResidentialExpense {
	if o == nil || IsNil(o.ResidentialExpense) {
		var ret ResidentialExpense
		return ret
	}
	return *o.ResidentialExpense
}

// GetResidentialExpenseOk returns a tuple with the ResidentialExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetResidentialExpenseOk() (*ResidentialExpense, bool) {
	if o == nil || IsNil(o.ResidentialExpense) {
		return nil, false
	}
	return o.ResidentialExpense, true
}

// HasResidentialExpense returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasResidentialExpense() bool {
	if o != nil && !IsNil(o.ResidentialExpense) {
		return true
	}

	return false
}

// SetResidentialExpense gets a reference to the given ResidentialExpense and assigns it to the ResidentialExpense field.
func (o *EddCustomerResponse) SetResidentialExpense(v ResidentialExpense) {
	o.ResidentialExpense = &v
}

// GetSourceOfWealth returns the SourceOfWealth field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetSourceOfWealth() []SourceOfWealth {
	if o == nil || IsNil(o.SourceOfWealth) {
		var ret []SourceOfWealth
		return ret
	}
	return o.SourceOfWealth
}

// GetSourceOfWealthOk returns a tuple with the SourceOfWealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetSourceOfWealthOk() ([]SourceOfWealth, bool) {
	if o == nil || IsNil(o.SourceOfWealth) {
		return nil, false
	}
	return o.SourceOfWealth, true
}

// HasSourceOfWealth returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasSourceOfWealth() bool {
	if o != nil && !IsNil(o.SourceOfWealth) {
		return true
	}

	return false
}

// SetSourceOfWealth gets a reference to the given []SourceOfWealth and assigns it to the SourceOfWealth field.
func (o *EddCustomerResponse) SetSourceOfWealth(v []SourceOfWealth) {
	o.SourceOfWealth = v
}

// GetAdditionalQuestions returns the AdditionalQuestions field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetAdditionalQuestions() []Question {
	if o == nil || IsNil(o.AdditionalQuestions) {
		var ret []Question
		return ret
	}
	return o.AdditionalQuestions
}

// GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetAdditionalQuestionsOk() ([]Question, bool) {
	if o == nil || IsNil(o.AdditionalQuestions) {
		return nil, false
	}
	return o.AdditionalQuestions, true
}

// HasAdditionalQuestions returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasAdditionalQuestions() bool {
	if o != nil && !IsNil(o.AdditionalQuestions) {
		return true
	}

	return false
}

// SetAdditionalQuestions gets a reference to the given []Question and assigns it to the AdditionalQuestions field.
func (o *EddCustomerResponse) SetAdditionalQuestions(v []Question) {
	o.AdditionalQuestions = v
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetCaseId() int32 {
	if o == nil || IsNil(o.CaseId) {
		var ret int32
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetCaseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given int32 and assigns it to the CaseId field.
func (o *EddCustomerResponse) SetCaseId(v int32) {
	o.CaseId = &v
}

// GetReason returns the Reason field value
func (o *EddCustomerResponse) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *EddCustomerResponse) SetReason(v string) {
	o.Reason = v
}

// GetRelatedResourceId returns the RelatedResourceId field value
func (o *EddCustomerResponse) GetRelatedResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelatedResourceId
}

// GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field value
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetRelatedResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceId, true
}

// SetRelatedResourceId sets field value
func (o *EddCustomerResponse) SetRelatedResourceId(v string) {
	o.RelatedResourceId = v
}

// GetRelatedResourceType returns the RelatedResourceType field value
func (o *EddCustomerResponse) GetRelatedResourceType() RelatedResourceType1 {
	if o == nil {
		var ret RelatedResourceType1
		return ret
	}

	return o.RelatedResourceType
}

// GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field value
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceType, true
}

// SetRelatedResourceType sets field value
func (o *EddCustomerResponse) SetRelatedResourceType(v RelatedResourceType1) {
	o.RelatedResourceType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *EddCustomerResponse) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *EddCustomerResponse) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *EddCustomerResponse) SetTenant(v string) {
	o.Tenant = &v
}

// GetCreationTime returns the CreationTime field value
func (o *EddCustomerResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *EddCustomerResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetDeletionTime returns the DeletionTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *EddCustomerResponse) GetDeletionTime() time.Time {
	if o == nil || o.DeletionTime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DeletionTime.Get()
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EddCustomerResponse) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletionTime.Get(), o.DeletionTime.IsSet()
}

// SetDeletionTime sets field value
func (o *EddCustomerResponse) SetDeletionTime(v time.Time) {
	o.DeletionTime.Set(&v)
}

// GetId returns the Id field value
func (o *EddCustomerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EddCustomerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EddCustomerResponse) SetId(v string) {
	o.Id = v
}

func (o EddCustomerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EddCustomerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CitizenshipCountries) {
		toSerialize["citizenship_countries"] = o.CitizenshipCountries
	}
	if !IsNil(o.EmploymentType) {
		toSerialize["employment_type"] = o.EmploymentType
	}
	if !IsNil(o.Income) {
		toSerialize["income"] = o.Income
	}
	if !IsNil(o.NegativeNewsFindings) {
		toSerialize["negative_news_findings"] = o.NegativeNewsFindings
	}
	if !IsNil(o.Occupation) {
		toSerialize["occupation"] = o.Occupation
	}
	if !IsNil(o.OccupationIndustry) {
		toSerialize["occupation_industry"] = o.OccupationIndustry
	}
	if !IsNil(o.RecurringDirectDeposit) {
		toSerialize["recurring_direct_deposit"] = o.RecurringDirectDeposit
	}
	if !IsNil(o.ResidenceType) {
		toSerialize["residence_type"] = o.ResidenceType
	}
	if !IsNil(o.ResidentialExpense) {
		toSerialize["residential_expense"] = o.ResidentialExpense
	}
	if !IsNil(o.SourceOfWealth) {
		toSerialize["source_of_wealth"] = o.SourceOfWealth
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

func (o *EddCustomerResponse) UnmarshalJSON(data []byte) (err error) {
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

	varEddCustomerResponse := _EddCustomerResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEddCustomerResponse)

	if err != nil {
		return err
	}

	*o = EddCustomerResponse(varEddCustomerResponse)

	return err
}

type NullableEddCustomerResponse struct {
	value *EddCustomerResponse
	isSet bool
}

func (v NullableEddCustomerResponse) Get() *EddCustomerResponse {
	return v.value
}

func (v *NullableEddCustomerResponse) Set(val *EddCustomerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEddCustomerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEddCustomerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEddCustomerResponse(val *EddCustomerResponse) *NullableEddCustomerResponse {
	return &NullableEddCustomerResponse{value: val, isSet: true}
}

func (v NullableEddCustomerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEddCustomerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
