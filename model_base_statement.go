/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BaseStatement struct for BaseStatement
type BaseStatement struct {
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
	AccountSummary *AccountSummary `json:"account_summary,omitempty"`
	AuthorizedSigner []Person1 `json:"authorized_signer,omitempty"`
	Disclosure *string `json:"disclosure,omitempty"`
	JointAccountHolders []Person1 `json:"joint_account_holders,omitempty"`
	PrimaryAccountHolderBusiness *Business1 `json:"primary_account_holder_business,omitempty"`
	PrimaryAccountHolderPersonal *Person1 `json:"primary_account_holder_personal,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
}

// NewBaseStatement instantiates a new BaseStatement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseStatement() *BaseStatement {
	this := BaseStatement{}
	return &this
}

// NewBaseStatementWithDefaults instantiates a new BaseStatement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseStatementWithDefaults() *BaseStatement {
	this := BaseStatement{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *BaseStatement) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *BaseStatement) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *BaseStatement) SetAccountId(v string) {
	o.AccountId = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *BaseStatement) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *BaseStatement) HasDueDate() bool {
	if o != nil && o.DueDate != nil {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *BaseStatement) SetDueDate(v string) {
	o.DueDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BaseStatement) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BaseStatement) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *BaseStatement) SetEndDate(v string) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseStatement) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseStatement) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseStatement) SetId(v string) {
	o.Id = &v
}

// GetIssueDate returns the IssueDate field value if set, zero value otherwise.
func (o *BaseStatement) GetIssueDate() string {
	if o == nil || o.IssueDate == nil {
		var ret string
		return ret
	}
	return *o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetIssueDateOk() (*string, bool) {
	if o == nil || o.IssueDate == nil {
		return nil, false
	}
	return o.IssueDate, true
}

// HasIssueDate returns a boolean if a field has been set.
func (o *BaseStatement) HasIssueDate() bool {
	if o != nil && o.IssueDate != nil {
		return true
	}

	return false
}

// SetIssueDate gets a reference to the given string and assigns it to the IssueDate field.
func (o *BaseStatement) SetIssueDate(v string) {
	o.IssueDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BaseStatement) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BaseStatement) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *BaseStatement) SetStartDate(v string) {
	o.StartDate = &v
}

// GetAccountSummary returns the AccountSummary field value if set, zero value otherwise.
func (o *BaseStatement) GetAccountSummary() AccountSummary {
	if o == nil || o.AccountSummary == nil {
		var ret AccountSummary
		return ret
	}
	return *o.AccountSummary
}

// GetAccountSummaryOk returns a tuple with the AccountSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetAccountSummaryOk() (*AccountSummary, bool) {
	if o == nil || o.AccountSummary == nil {
		return nil, false
	}
	return o.AccountSummary, true
}

// HasAccountSummary returns a boolean if a field has been set.
func (o *BaseStatement) HasAccountSummary() bool {
	if o != nil && o.AccountSummary != nil {
		return true
	}

	return false
}

// SetAccountSummary gets a reference to the given AccountSummary and assigns it to the AccountSummary field.
func (o *BaseStatement) SetAccountSummary(v AccountSummary) {
	o.AccountSummary = &v
}

// GetAuthorizedSigner returns the AuthorizedSigner field value if set, zero value otherwise.
func (o *BaseStatement) GetAuthorizedSigner() []Person1 {
	if o == nil || o.AuthorizedSigner == nil {
		var ret []Person1
		return ret
	}
	return o.AuthorizedSigner
}

// GetAuthorizedSignerOk returns a tuple with the AuthorizedSigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetAuthorizedSignerOk() ([]Person1, bool) {
	if o == nil || o.AuthorizedSigner == nil {
		return nil, false
	}
	return o.AuthorizedSigner, true
}

// HasAuthorizedSigner returns a boolean if a field has been set.
func (o *BaseStatement) HasAuthorizedSigner() bool {
	if o != nil && o.AuthorizedSigner != nil {
		return true
	}

	return false
}

// SetAuthorizedSigner gets a reference to the given []Person1 and assigns it to the AuthorizedSigner field.
func (o *BaseStatement) SetAuthorizedSigner(v []Person1) {
	o.AuthorizedSigner = v
}

// GetDisclosure returns the Disclosure field value if set, zero value otherwise.
func (o *BaseStatement) GetDisclosure() string {
	if o == nil || o.Disclosure == nil {
		var ret string
		return ret
	}
	return *o.Disclosure
}

// GetDisclosureOk returns a tuple with the Disclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetDisclosureOk() (*string, bool) {
	if o == nil || o.Disclosure == nil {
		return nil, false
	}
	return o.Disclosure, true
}

// HasDisclosure returns a boolean if a field has been set.
func (o *BaseStatement) HasDisclosure() bool {
	if o != nil && o.Disclosure != nil {
		return true
	}

	return false
}

// SetDisclosure gets a reference to the given string and assigns it to the Disclosure field.
func (o *BaseStatement) SetDisclosure(v string) {
	o.Disclosure = &v
}

// GetJointAccountHolders returns the JointAccountHolders field value if set, zero value otherwise.
func (o *BaseStatement) GetJointAccountHolders() []Person1 {
	if o == nil || o.JointAccountHolders == nil {
		var ret []Person1
		return ret
	}
	return o.JointAccountHolders
}

// GetJointAccountHoldersOk returns a tuple with the JointAccountHolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetJointAccountHoldersOk() ([]Person1, bool) {
	if o == nil || o.JointAccountHolders == nil {
		return nil, false
	}
	return o.JointAccountHolders, true
}

// HasJointAccountHolders returns a boolean if a field has been set.
func (o *BaseStatement) HasJointAccountHolders() bool {
	if o != nil && o.JointAccountHolders != nil {
		return true
	}

	return false
}

// SetJointAccountHolders gets a reference to the given []Person1 and assigns it to the JointAccountHolders field.
func (o *BaseStatement) SetJointAccountHolders(v []Person1) {
	o.JointAccountHolders = v
}

// GetPrimaryAccountHolderBusiness returns the PrimaryAccountHolderBusiness field value if set, zero value otherwise.
func (o *BaseStatement) GetPrimaryAccountHolderBusiness() Business1 {
	if o == nil || o.PrimaryAccountHolderBusiness == nil {
		var ret Business1
		return ret
	}
	return *o.PrimaryAccountHolderBusiness
}

// GetPrimaryAccountHolderBusinessOk returns a tuple with the PrimaryAccountHolderBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetPrimaryAccountHolderBusinessOk() (*Business1, bool) {
	if o == nil || o.PrimaryAccountHolderBusiness == nil {
		return nil, false
	}
	return o.PrimaryAccountHolderBusiness, true
}

// HasPrimaryAccountHolderBusiness returns a boolean if a field has been set.
func (o *BaseStatement) HasPrimaryAccountHolderBusiness() bool {
	if o != nil && o.PrimaryAccountHolderBusiness != nil {
		return true
	}

	return false
}

// SetPrimaryAccountHolderBusiness gets a reference to the given Business1 and assigns it to the PrimaryAccountHolderBusiness field.
func (o *BaseStatement) SetPrimaryAccountHolderBusiness(v Business1) {
	o.PrimaryAccountHolderBusiness = &v
}

// GetPrimaryAccountHolderPersonal returns the PrimaryAccountHolderPersonal field value if set, zero value otherwise.
func (o *BaseStatement) GetPrimaryAccountHolderPersonal() Person1 {
	if o == nil || o.PrimaryAccountHolderPersonal == nil {
		var ret Person1
		return ret
	}
	return *o.PrimaryAccountHolderPersonal
}

// GetPrimaryAccountHolderPersonalOk returns a tuple with the PrimaryAccountHolderPersonal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetPrimaryAccountHolderPersonalOk() (*Person1, bool) {
	if o == nil || o.PrimaryAccountHolderPersonal == nil {
		return nil, false
	}
	return o.PrimaryAccountHolderPersonal, true
}

// HasPrimaryAccountHolderPersonal returns a boolean if a field has been set.
func (o *BaseStatement) HasPrimaryAccountHolderPersonal() bool {
	if o != nil && o.PrimaryAccountHolderPersonal != nil {
		return true
	}

	return false
}

// SetPrimaryAccountHolderPersonal gets a reference to the given Person1 and assigns it to the PrimaryAccountHolderPersonal field.
func (o *BaseStatement) SetPrimaryAccountHolderPersonal(v Person1) {
	o.PrimaryAccountHolderPersonal = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *BaseStatement) GetTransactions() []Transaction {
	if o == nil || o.Transactions == nil {
		var ret []Transaction
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseStatement) GetTransactionsOk() ([]Transaction, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *BaseStatement) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []Transaction and assigns it to the Transactions field.
func (o *BaseStatement) SetTransactions(v []Transaction) {
	o.Transactions = v
}

func (o BaseStatement) MarshalJSON() ([]byte, error) {
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
	if o.AccountSummary != nil {
		toSerialize["account_summary"] = o.AccountSummary
	}
	if o.AuthorizedSigner != nil {
		toSerialize["authorized_signer"] = o.AuthorizedSigner
	}
	if o.Disclosure != nil {
		toSerialize["disclosure"] = o.Disclosure
	}
	if o.JointAccountHolders != nil {
		toSerialize["joint_account_holders"] = o.JointAccountHolders
	}
	if o.PrimaryAccountHolderBusiness != nil {
		toSerialize["primary_account_holder_business"] = o.PrimaryAccountHolderBusiness
	}
	if o.PrimaryAccountHolderPersonal != nil {
		toSerialize["primary_account_holder_personal"] = o.PrimaryAccountHolderPersonal
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableBaseStatement struct {
	value *BaseStatement
	isSet bool
}

func (v NullableBaseStatement) Get() *BaseStatement {
	return v.value
}

func (v *NullableBaseStatement) Set(val *BaseStatement) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseStatement(val *BaseStatement) *NullableBaseStatement {
	return &NullableBaseStatement{value: val, isSet: true}
}

func (v NullableBaseStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


