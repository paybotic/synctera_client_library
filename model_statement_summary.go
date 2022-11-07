/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatementSummary struct for StatementSummary
type StatementSummary struct {
	// The unique identifier of the account the statement belongs to
	AccountId *string `json:"account_id,omitempty"`
	// The limit date when the due amount indicated on the statement should be paid
	DueDate *string `json:"due_date,omitempty"`
	// The date indicating the ending of the time interval covered by the statement
	EndDate *string `json:"end_date,omitempty"`
	// statement ID
	Id *string `json:"id,omitempty"`
	// The date when the statement has been issued
	IssueDate *string `json:"issue_date,omitempty"`
	// The date indicating the beginning of the time interval covered by the statement
	StartDate *string `json:"start_date,omitempty"`
}

// NewStatementSummary instantiates a new StatementSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementSummary() *StatementSummary {
	this := StatementSummary{}
	return &this
}

// NewStatementSummaryWithDefaults instantiates a new StatementSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementSummaryWithDefaults() *StatementSummary {
	this := StatementSummary{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *StatementSummary) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementSummary) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *StatementSummary) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *StatementSummary) SetAccountId(v string) {
	o.AccountId = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *StatementSummary) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementSummary) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *StatementSummary) HasDueDate() bool {
	if o != nil && o.DueDate != nil {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *StatementSummary) SetDueDate(v string) {
	o.DueDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *StatementSummary) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementSummary) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *StatementSummary) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *StatementSummary) SetEndDate(v string) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StatementSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementSummary) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StatementSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StatementSummary) SetId(v string) {
	o.Id = &v
}

// GetIssueDate returns the IssueDate field value if set, zero value otherwise.
func (o *StatementSummary) GetIssueDate() string {
	if o == nil || o.IssueDate == nil {
		var ret string
		return ret
	}
	return *o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementSummary) GetIssueDateOk() (*string, bool) {
	if o == nil || o.IssueDate == nil {
		return nil, false
	}
	return o.IssueDate, true
}

// HasIssueDate returns a boolean if a field has been set.
func (o *StatementSummary) HasIssueDate() bool {
	if o != nil && o.IssueDate != nil {
		return true
	}

	return false
}

// SetIssueDate gets a reference to the given string and assigns it to the IssueDate field.
func (o *StatementSummary) SetIssueDate(v string) {
	o.IssueDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *StatementSummary) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementSummary) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *StatementSummary) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *StatementSummary) SetStartDate(v string) {
	o.StartDate = &v
}

func (o StatementSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.DueDate != nil {
		toSerialize["due_date"] = o.DueDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IssueDate != nil {
		toSerialize["issue_date"] = o.IssueDate
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableStatementSummary struct {
	value *StatementSummary
	isSet bool
}

func (v NullableStatementSummary) Get() *StatementSummary {
	return v.value
}

func (v *NullableStatementSummary) Set(val *StatementSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementSummary(val *StatementSummary) *NullableStatementSummary {
	return &NullableStatementSummary{value: val, isSet: true}
}

func (v NullableStatementSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


