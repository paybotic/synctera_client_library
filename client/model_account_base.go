/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// checks if the AccountBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountBase{}

// AccountBase struct for AccountBase
type AccountBase struct {
	AccessStatus *AccountAccessStatus `json:"access_status,omitempty"`
	// Account number
	AccountNumber *string `json:"account_number,omitempty"`
	// The response will contain the bank fintech ID (3 or 6 digits) plus the last 4 digits, with the digits in between replaced with * characters. Shadow mode account numbers will not be masked.
	AccountNumberMasked *string `json:"account_number_masked,omitempty"`
	// Purpose of the account
	AccountPurpose *string      `json:"account_purpose,omitempty"`
	AccountType    *AccountType `json:"account_type,omitempty"`
	// The application ID for this account.
	ApplicationId *string `json:"application_id,omitempty"`
	// A list of balances for account based on different type
	Balances []Balance `json:"balances,omitempty"`
	// Bank routing number
	BankRouting *string `json:"bank_routing,omitempty"`
	// Account creation timestamp in RFC3339 format
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD
	Currency *string `json:"currency,omitempty" validate:"regexp=^[A-Z]{3}$"`
	// A list of the customer IDs of the account holders.
	CustomerIds  []string      `json:"customer_ids,omitempty"`
	CustomerType *CustomerType `json:"customer_type,omitempty"`
	// Exchange rate type
	ExchangeRateType *string `json:"exchange_rate_type,omitempty"`
	// International bank account number
	Iban *string `json:"iban,omitempty"`
	// Account ID
	Id *string `json:"id,omitempty"`
	// Account is investment (variable balance) account or a multi-balance account pool. Default false
	IsAccountPool *bool `json:"is_account_pool,omitempty"`
	// A flag to indicate whether SAR generation is enabled.
	IsSarEnabled *bool `json:"is_sar_enabled,omitempty"`
	// Timestamp of the last account modification in RFC3339 format
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// User provided account metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// User provided account nickname
	Nickname *string        `json:"nickname,omitempty"`
	Status   *AccountStatus `json:"status,omitempty"`
	// SWIFT code
	SwiftCode *string `json:"swift_code,omitempty"`
}

// NewAccountBase instantiates a new AccountBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBase() *AccountBase {
	this := AccountBase{}
	return &this
}

// NewAccountBaseWithDefaults instantiates a new AccountBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBaseWithDefaults() *AccountBase {
	this := AccountBase{}
	return &this
}

// GetAccessStatus returns the AccessStatus field value if set, zero value otherwise.
func (o *AccountBase) GetAccessStatus() AccountAccessStatus {
	if o == nil || IsNil(o.AccessStatus) {
		var ret AccountAccessStatus
		return ret
	}
	return *o.AccessStatus
}

// GetAccessStatusOk returns a tuple with the AccessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccessStatusOk() (*AccountAccessStatus, bool) {
	if o == nil || IsNil(o.AccessStatus) {
		return nil, false
	}
	return o.AccessStatus, true
}

// HasAccessStatus returns a boolean if a field has been set.
func (o *AccountBase) HasAccessStatus() bool {
	if o != nil && !IsNil(o.AccessStatus) {
		return true
	}

	return false
}

// SetAccessStatus gets a reference to the given AccountAccessStatus and assigns it to the AccessStatus field.
func (o *AccountBase) SetAccessStatus(v AccountAccessStatus) {
	o.AccessStatus = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountBase) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountBase) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountBase) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountNumberMasked returns the AccountNumberMasked field value if set, zero value otherwise.
func (o *AccountBase) GetAccountNumberMasked() string {
	if o == nil || IsNil(o.AccountNumberMasked) {
		var ret string
		return ret
	}
	return *o.AccountNumberMasked
}

// GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccountNumberMaskedOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumberMasked) {
		return nil, false
	}
	return o.AccountNumberMasked, true
}

// HasAccountNumberMasked returns a boolean if a field has been set.
func (o *AccountBase) HasAccountNumberMasked() bool {
	if o != nil && !IsNil(o.AccountNumberMasked) {
		return true
	}

	return false
}

// SetAccountNumberMasked gets a reference to the given string and assigns it to the AccountNumberMasked field.
func (o *AccountBase) SetAccountNumberMasked(v string) {
	o.AccountNumberMasked = &v
}

// GetAccountPurpose returns the AccountPurpose field value if set, zero value otherwise.
func (o *AccountBase) GetAccountPurpose() string {
	if o == nil || IsNil(o.AccountPurpose) {
		var ret string
		return ret
	}
	return *o.AccountPurpose
}

// GetAccountPurposeOk returns a tuple with the AccountPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccountPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountPurpose) {
		return nil, false
	}
	return o.AccountPurpose, true
}

// HasAccountPurpose returns a boolean if a field has been set.
func (o *AccountBase) HasAccountPurpose() bool {
	if o != nil && !IsNil(o.AccountPurpose) {
		return true
	}

	return false
}

// SetAccountPurpose gets a reference to the given string and assigns it to the AccountPurpose field.
func (o *AccountBase) SetAccountPurpose(v string) {
	o.AccountPurpose = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *AccountBase) GetAccountType() AccountType {
	if o == nil || IsNil(o.AccountType) {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *AccountBase) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *AccountBase) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *AccountBase) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *AccountBase) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *AccountBase) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *AccountBase) GetBalances() []Balance {
	if o == nil || IsNil(o.Balances) {
		var ret []Balance
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetBalancesOk() ([]Balance, bool) {
	if o == nil || IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *AccountBase) HasBalances() bool {
	if o != nil && !IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []Balance and assigns it to the Balances field.
func (o *AccountBase) SetBalances(v []Balance) {
	o.Balances = v
}

// GetBankRouting returns the BankRouting field value if set, zero value otherwise.
func (o *AccountBase) GetBankRouting() string {
	if o == nil || IsNil(o.BankRouting) {
		var ret string
		return ret
	}
	return *o.BankRouting
}

// GetBankRoutingOk returns a tuple with the BankRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetBankRoutingOk() (*string, bool) {
	if o == nil || IsNil(o.BankRouting) {
		return nil, false
	}
	return o.BankRouting, true
}

// HasBankRouting returns a boolean if a field has been set.
func (o *AccountBase) HasBankRouting() bool {
	if o != nil && !IsNil(o.BankRouting) {
		return true
	}

	return false
}

// SetBankRouting gets a reference to the given string and assigns it to the BankRouting field.
func (o *AccountBase) SetBankRouting(v string) {
	o.BankRouting = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AccountBase) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AccountBase) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *AccountBase) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountBase) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountBase) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountBase) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerIds returns the CustomerIds field value if set, zero value otherwise.
func (o *AccountBase) GetCustomerIds() []string {
	if o == nil || IsNil(o.CustomerIds) {
		var ret []string
		return ret
	}
	return o.CustomerIds
}

// GetCustomerIdsOk returns a tuple with the CustomerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCustomerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomerIds) {
		return nil, false
	}
	return o.CustomerIds, true
}

// HasCustomerIds returns a boolean if a field has been set.
func (o *AccountBase) HasCustomerIds() bool {
	if o != nil && !IsNil(o.CustomerIds) {
		return true
	}

	return false
}

// SetCustomerIds gets a reference to the given []string and assigns it to the CustomerIds field.
func (o *AccountBase) SetCustomerIds(v []string) {
	o.CustomerIds = v
}

// GetCustomerType returns the CustomerType field value if set, zero value otherwise.
func (o *AccountBase) GetCustomerType() CustomerType {
	if o == nil || IsNil(o.CustomerType) {
		var ret CustomerType
		return ret
	}
	return *o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCustomerTypeOk() (*CustomerType, bool) {
	if o == nil || IsNil(o.CustomerType) {
		return nil, false
	}
	return o.CustomerType, true
}

// HasCustomerType returns a boolean if a field has been set.
func (o *AccountBase) HasCustomerType() bool {
	if o != nil && !IsNil(o.CustomerType) {
		return true
	}

	return false
}

// SetCustomerType gets a reference to the given CustomerType and assigns it to the CustomerType field.
func (o *AccountBase) SetCustomerType(v CustomerType) {
	o.CustomerType = &v
}

// GetExchangeRateType returns the ExchangeRateType field value if set, zero value otherwise.
func (o *AccountBase) GetExchangeRateType() string {
	if o == nil || IsNil(o.ExchangeRateType) {
		var ret string
		return ret
	}
	return *o.ExchangeRateType
}

// GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetExchangeRateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeRateType) {
		return nil, false
	}
	return o.ExchangeRateType, true
}

// HasExchangeRateType returns a boolean if a field has been set.
func (o *AccountBase) HasExchangeRateType() bool {
	if o != nil && !IsNil(o.ExchangeRateType) {
		return true
	}

	return false
}

// SetExchangeRateType gets a reference to the given string and assigns it to the ExchangeRateType field.
func (o *AccountBase) SetExchangeRateType(v string) {
	o.ExchangeRateType = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *AccountBase) GetIban() string {
	if o == nil || IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIbanOk() (*string, bool) {
	if o == nil || IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *AccountBase) HasIban() bool {
	if o != nil && !IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *AccountBase) SetIban(v string) {
	o.Iban = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountBase) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountBase) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountBase) SetId(v string) {
	o.Id = &v
}

// GetIsAccountPool returns the IsAccountPool field value if set, zero value otherwise.
func (o *AccountBase) GetIsAccountPool() bool {
	if o == nil || IsNil(o.IsAccountPool) {
		var ret bool
		return ret
	}
	return *o.IsAccountPool
}

// GetIsAccountPoolOk returns a tuple with the IsAccountPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIsAccountPoolOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAccountPool) {
		return nil, false
	}
	return o.IsAccountPool, true
}

// HasIsAccountPool returns a boolean if a field has been set.
func (o *AccountBase) HasIsAccountPool() bool {
	if o != nil && !IsNil(o.IsAccountPool) {
		return true
	}

	return false
}

// SetIsAccountPool gets a reference to the given bool and assigns it to the IsAccountPool field.
func (o *AccountBase) SetIsAccountPool(v bool) {
	o.IsAccountPool = &v
}

// GetIsSarEnabled returns the IsSarEnabled field value if set, zero value otherwise.
func (o *AccountBase) GetIsSarEnabled() bool {
	if o == nil || IsNil(o.IsSarEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSarEnabled
}

// GetIsSarEnabledOk returns a tuple with the IsSarEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIsSarEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSarEnabled) {
		return nil, false
	}
	return o.IsSarEnabled, true
}

// HasIsSarEnabled returns a boolean if a field has been set.
func (o *AccountBase) HasIsSarEnabled() bool {
	if o != nil && !IsNil(o.IsSarEnabled) {
		return true
	}

	return false
}

// SetIsSarEnabled gets a reference to the given bool and assigns it to the IsSarEnabled field.
func (o *AccountBase) SetIsSarEnabled(v bool) {
	o.IsSarEnabled = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *AccountBase) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *AccountBase) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *AccountBase) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AccountBase) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AccountBase) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AccountBase) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *AccountBase) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *AccountBase) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *AccountBase) SetNickname(v string) {
	o.Nickname = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountBase) GetStatus() AccountStatus {
	if o == nil || IsNil(o.Status) {
		var ret AccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetStatusOk() (*AccountStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountBase) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AccountStatus and assigns it to the Status field.
func (o *AccountBase) SetStatus(v AccountStatus) {
	o.Status = &v
}

// GetSwiftCode returns the SwiftCode field value if set, zero value otherwise.
func (o *AccountBase) GetSwiftCode() string {
	if o == nil || IsNil(o.SwiftCode) {
		var ret string
		return ret
	}
	return *o.SwiftCode
}

// GetSwiftCodeOk returns a tuple with the SwiftCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetSwiftCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SwiftCode) {
		return nil, false
	}
	return o.SwiftCode, true
}

// HasSwiftCode returns a boolean if a field has been set.
func (o *AccountBase) HasSwiftCode() bool {
	if o != nil && !IsNil(o.SwiftCode) {
		return true
	}

	return false
}

// SetSwiftCode gets a reference to the given string and assigns it to the SwiftCode field.
func (o *AccountBase) SetSwiftCode(v string) {
	o.SwiftCode = &v
}

func (o AccountBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessStatus) {
		toSerialize["access_status"] = o.AccessStatus
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.AccountNumberMasked) {
		toSerialize["account_number_masked"] = o.AccountNumberMasked
	}
	if !IsNil(o.AccountPurpose) {
		toSerialize["account_purpose"] = o.AccountPurpose
	}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	if !IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	if !IsNil(o.BankRouting) {
		toSerialize["bank_routing"] = o.BankRouting
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CustomerIds) {
		toSerialize["customer_ids"] = o.CustomerIds
	}
	if !IsNil(o.CustomerType) {
		toSerialize["customer_type"] = o.CustomerType
	}
	if !IsNil(o.ExchangeRateType) {
		toSerialize["exchange_rate_type"] = o.ExchangeRateType
	}
	if !IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsAccountPool) {
		toSerialize["is_account_pool"] = o.IsAccountPool
	}
	if !IsNil(o.IsSarEnabled) {
		toSerialize["is_sar_enabled"] = o.IsSarEnabled
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SwiftCode) {
		toSerialize["swift_code"] = o.SwiftCode
	}
	return toSerialize, nil
}

type NullableAccountBase struct {
	value *AccountBase
	isSet bool
}

func (v NullableAccountBase) Get() *AccountBase {
	return v.value
}

func (v *NullableAccountBase) Set(val *AccountBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBase(val *AccountBase) *NullableAccountBase {
	return &NullableAccountBase{value: val, isSet: true}
}

func (v NullableAccountBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
