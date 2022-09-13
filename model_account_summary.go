/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// AccountSummary struct for AccountSummary
type AccountSummary struct {
	// Account number
	AccountNumber *string `json:"account_number,omitempty"`
	// Account Status
	AccountStatus *string `json:"account_status,omitempty"`
	// The type of the account. In lead mode, this always takes the value of the template. If not specified in shadow mode, CHECKING will be assumed. 
	AccountType *string `json:"account_type,omitempty"`
	BalanceCeiling *AccountSummaryBalanceCeiling `json:"balance_ceiling,omitempty"`
	BalanceFloor *AccountSummaryBalanceFloor `json:"balance_floor,omitempty"`
	// Account creation time
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD
	Currency *string `json:"currency,omitempty"`
	// Customer type
	CustomerType *string `json:"customer_type,omitempty"`
	FinancialInstitution *FinancialInstitution `json:"financial_institution,omitempty"`
	// The unique identifier of the account the statement belongs to
	Id *string `json:"id,omitempty"`
	// Account last modification time
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// User provided account nickname
	Nickname *string `json:"nickname,omitempty"`
}

// NewAccountSummary instantiates a new AccountSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSummary() *AccountSummary {
	this := AccountSummary{}
	return &this
}

// NewAccountSummaryWithDefaults instantiates a new AccountSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSummaryWithDefaults() *AccountSummary {
	this := AccountSummary{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountSummary) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountSummary) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountSummary) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *AccountSummary) GetAccountStatus() string {
	if o == nil || o.AccountStatus == nil {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetAccountStatusOk() (*string, bool) {
	if o == nil || o.AccountStatus == nil {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *AccountSummary) HasAccountStatus() bool {
	if o != nil && o.AccountStatus != nil {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *AccountSummary) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *AccountSummary) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *AccountSummary) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *AccountSummary) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBalanceCeiling returns the BalanceCeiling field value if set, zero value otherwise.
func (o *AccountSummary) GetBalanceCeiling() AccountSummaryBalanceCeiling {
	if o == nil || o.BalanceCeiling == nil {
		var ret AccountSummaryBalanceCeiling
		return ret
	}
	return *o.BalanceCeiling
}

// GetBalanceCeilingOk returns a tuple with the BalanceCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetBalanceCeilingOk() (*AccountSummaryBalanceCeiling, bool) {
	if o == nil || o.BalanceCeiling == nil {
		return nil, false
	}
	return o.BalanceCeiling, true
}

// HasBalanceCeiling returns a boolean if a field has been set.
func (o *AccountSummary) HasBalanceCeiling() bool {
	if o != nil && o.BalanceCeiling != nil {
		return true
	}

	return false
}

// SetBalanceCeiling gets a reference to the given AccountSummaryBalanceCeiling and assigns it to the BalanceCeiling field.
func (o *AccountSummary) SetBalanceCeiling(v AccountSummaryBalanceCeiling) {
	o.BalanceCeiling = &v
}

// GetBalanceFloor returns the BalanceFloor field value if set, zero value otherwise.
func (o *AccountSummary) GetBalanceFloor() AccountSummaryBalanceFloor {
	if o == nil || o.BalanceFloor == nil {
		var ret AccountSummaryBalanceFloor
		return ret
	}
	return *o.BalanceFloor
}

// GetBalanceFloorOk returns a tuple with the BalanceFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetBalanceFloorOk() (*AccountSummaryBalanceFloor, bool) {
	if o == nil || o.BalanceFloor == nil {
		return nil, false
	}
	return o.BalanceFloor, true
}

// HasBalanceFloor returns a boolean if a field has been set.
func (o *AccountSummary) HasBalanceFloor() bool {
	if o != nil && o.BalanceFloor != nil {
		return true
	}

	return false
}

// SetBalanceFloor gets a reference to the given AccountSummaryBalanceFloor and assigns it to the BalanceFloor field.
func (o *AccountSummary) SetBalanceFloor(v AccountSummaryBalanceFloor) {
	o.BalanceFloor = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AccountSummary) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AccountSummary) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *AccountSummary) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountSummary) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountSummary) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountSummary) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerType returns the CustomerType field value if set, zero value otherwise.
func (o *AccountSummary) GetCustomerType() string {
	if o == nil || o.CustomerType == nil {
		var ret string
		return ret
	}
	return *o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetCustomerTypeOk() (*string, bool) {
	if o == nil || o.CustomerType == nil {
		return nil, false
	}
	return o.CustomerType, true
}

// HasCustomerType returns a boolean if a field has been set.
func (o *AccountSummary) HasCustomerType() bool {
	if o != nil && o.CustomerType != nil {
		return true
	}

	return false
}

// SetCustomerType gets a reference to the given string and assigns it to the CustomerType field.
func (o *AccountSummary) SetCustomerType(v string) {
	o.CustomerType = &v
}

// GetFinancialInstitution returns the FinancialInstitution field value if set, zero value otherwise.
func (o *AccountSummary) GetFinancialInstitution() FinancialInstitution {
	if o == nil || o.FinancialInstitution == nil {
		var ret FinancialInstitution
		return ret
	}
	return *o.FinancialInstitution
}

// GetFinancialInstitutionOk returns a tuple with the FinancialInstitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetFinancialInstitutionOk() (*FinancialInstitution, bool) {
	if o == nil || o.FinancialInstitution == nil {
		return nil, false
	}
	return o.FinancialInstitution, true
}

// HasFinancialInstitution returns a boolean if a field has been set.
func (o *AccountSummary) HasFinancialInstitution() bool {
	if o != nil && o.FinancialInstitution != nil {
		return true
	}

	return false
}

// SetFinancialInstitution gets a reference to the given FinancialInstitution and assigns it to the FinancialInstitution field.
func (o *AccountSummary) SetFinancialInstitution(v FinancialInstitution) {
	o.FinancialInstitution = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountSummary) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *AccountSummary) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *AccountSummary) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *AccountSummary) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *AccountSummary) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSummary) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *AccountSummary) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *AccountSummary) SetNickname(v string) {
	o.Nickname = &v
}

func (o AccountSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountNumber != nil {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.AccountStatus != nil {
		toSerialize["account_status"] = o.AccountStatus
	}
	if o.AccountType != nil {
		toSerialize["account_type"] = o.AccountType
	}
	if o.BalanceCeiling != nil {
		toSerialize["balance_ceiling"] = o.BalanceCeiling
	}
	if o.BalanceFloor != nil {
		toSerialize["balance_floor"] = o.BalanceFloor
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.CustomerType != nil {
		toSerialize["customer_type"] = o.CustomerType
	}
	if o.FinancialInstitution != nil {
		toSerialize["financial_institution"] = o.FinancialInstitution
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	return json.Marshal(toSerialize)
}

type NullableAccountSummary struct {
	value *AccountSummary
	isSet bool
}

func (v NullableAccountSummary) Get() *AccountSummary {
	return v.value
}

func (v *NullableAccountSummary) Set(val *AccountSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSummary(val *AccountSummary) *NullableAccountSummary {
	return &NullableAccountSummary{value: val, isSet: true}
}

func (v NullableAccountSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


