/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the EddTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EddTransaction{}

// EddTransaction struct for EddTransaction
type EddTransaction struct {
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
	// The source of funds for the transaction.
	SourceOfFunds *string `json:"source_of_funds,omitempty"`
	// The purpose of the transaction.
	TransactionPurpose   *string `json:"transaction_purpose,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EddTransaction EddTransaction

// NewEddTransaction instantiates a new EddTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEddTransaction(reason string, relatedResourceId string, relatedResourceType RelatedResourceType1) *EddTransaction {
	this := EddTransaction{}
	this.Reason = reason
	this.RelatedResourceId = relatedResourceId
	this.RelatedResourceType = relatedResourceType
	return &this
}

// NewEddTransactionWithDefaults instantiates a new EddTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEddTransactionWithDefaults() *EddTransaction {
	this := EddTransaction{}
	return &this
}

// GetAdditionalQuestions returns the AdditionalQuestions field value if set, zero value otherwise.
func (o *EddTransaction) GetAdditionalQuestions() []Question {
	if o == nil || IsNil(o.AdditionalQuestions) {
		var ret []Question
		return ret
	}
	return o.AdditionalQuestions
}

// GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddTransaction) GetAdditionalQuestionsOk() ([]Question, bool) {
	if o == nil || IsNil(o.AdditionalQuestions) {
		return nil, false
	}
	return o.AdditionalQuestions, true
}

// HasAdditionalQuestions returns a boolean if a field has been set.
func (o *EddTransaction) HasAdditionalQuestions() bool {
	if o != nil && !IsNil(o.AdditionalQuestions) {
		return true
	}

	return false
}

// SetAdditionalQuestions gets a reference to the given []Question and assigns it to the AdditionalQuestions field.
func (o *EddTransaction) SetAdditionalQuestions(v []Question) {
	o.AdditionalQuestions = v
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *EddTransaction) GetCaseId() int32 {
	if o == nil || IsNil(o.CaseId) {
		var ret int32
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddTransaction) GetCaseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *EddTransaction) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given int32 and assigns it to the CaseId field.
func (o *EddTransaction) SetCaseId(v int32) {
	o.CaseId = &v
}

// GetReason returns the Reason field value
func (o *EddTransaction) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *EddTransaction) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *EddTransaction) SetReason(v string) {
	o.Reason = v
}

// GetRelatedResourceId returns the RelatedResourceId field value
func (o *EddTransaction) GetRelatedResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelatedResourceId
}

// GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field value
// and a boolean to check if the value has been set.
func (o *EddTransaction) GetRelatedResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceId, true
}

// SetRelatedResourceId sets field value
func (o *EddTransaction) SetRelatedResourceId(v string) {
	o.RelatedResourceId = v
}

// GetRelatedResourceType returns the RelatedResourceType field value
func (o *EddTransaction) GetRelatedResourceType() RelatedResourceType1 {
	if o == nil {
		var ret RelatedResourceType1
		return ret
	}

	return o.RelatedResourceType
}

// GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field value
// and a boolean to check if the value has been set.
func (o *EddTransaction) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceType, true
}

// SetRelatedResourceType sets field value
func (o *EddTransaction) SetRelatedResourceType(v RelatedResourceType1) {
	o.RelatedResourceType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *EddTransaction) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddTransaction) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *EddTransaction) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *EddTransaction) SetTenant(v string) {
	o.Tenant = &v
}

// GetSourceOfFunds returns the SourceOfFunds field value if set, zero value otherwise.
func (o *EddTransaction) GetSourceOfFunds() string {
	if o == nil || IsNil(o.SourceOfFunds) {
		var ret string
		return ret
	}
	return *o.SourceOfFunds
}

// GetSourceOfFundsOk returns a tuple with the SourceOfFunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddTransaction) GetSourceOfFundsOk() (*string, bool) {
	if o == nil || IsNil(o.SourceOfFunds) {
		return nil, false
	}
	return o.SourceOfFunds, true
}

// HasSourceOfFunds returns a boolean if a field has been set.
func (o *EddTransaction) HasSourceOfFunds() bool {
	if o != nil && !IsNil(o.SourceOfFunds) {
		return true
	}

	return false
}

// SetSourceOfFunds gets a reference to the given string and assigns it to the SourceOfFunds field.
func (o *EddTransaction) SetSourceOfFunds(v string) {
	o.SourceOfFunds = &v
}

// GetTransactionPurpose returns the TransactionPurpose field value if set, zero value otherwise.
func (o *EddTransaction) GetTransactionPurpose() string {
	if o == nil || IsNil(o.TransactionPurpose) {
		var ret string
		return ret
	}
	return *o.TransactionPurpose
}

// GetTransactionPurposeOk returns a tuple with the TransactionPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EddTransaction) GetTransactionPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionPurpose) {
		return nil, false
	}
	return o.TransactionPurpose, true
}

// HasTransactionPurpose returns a boolean if a field has been set.
func (o *EddTransaction) HasTransactionPurpose() bool {
	if o != nil && !IsNil(o.TransactionPurpose) {
		return true
	}

	return false
}

// SetTransactionPurpose gets a reference to the given string and assigns it to the TransactionPurpose field.
func (o *EddTransaction) SetTransactionPurpose(v string) {
	o.TransactionPurpose = &v
}

func (o EddTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EddTransaction) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SourceOfFunds) {
		toSerialize["source_of_funds"] = o.SourceOfFunds
	}
	if !IsNil(o.TransactionPurpose) {
		toSerialize["transaction_purpose"] = o.TransactionPurpose
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EddTransaction) UnmarshalJSON(data []byte) (err error) {
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

	varEddTransaction := _EddTransaction{}

	err = json.Unmarshal(data, &varEddTransaction)

	if err != nil {
		return err
	}

	*o = EddTransaction(varEddTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_questions")
		delete(additionalProperties, "case_id")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "related_resource_id")
		delete(additionalProperties, "related_resource_type")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "source_of_funds")
		delete(additionalProperties, "transaction_purpose")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEddTransaction struct {
	value *EddTransaction
	isSet bool
}

func (v NullableEddTransaction) Get() *EddTransaction {
	return v.value
}

func (v *NullableEddTransaction) Set(val *EddTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableEddTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableEddTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEddTransaction(val *EddTransaction) *NullableEddTransaction {
	return &NullableEddTransaction{value: val, isSet: true}
}

func (v NullableEddTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEddTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
