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

// checks if the PatchAccountChargeSecured type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchAccountChargeSecured{}

// PatchAccountChargeSecured struct for PatchAccountChargeSecured
type PatchAccountChargeSecured struct {
	// A flag to indicate whether ACH transactions are enabled.
	IsAchEnabled *bool `json:"is_ach_enabled,omitempty"`
	// A flag to indicate whether card transactions are enabled.
	IsCardEnabled *bool `json:"is_card_enabled,omitempty"`
	// A flag to indicate whether EFT Canada transactions are enabled.
	IsEftCaEnabled *bool `json:"is_eft_ca_enabled,omitempty"`
	// A flag to indicate whether P2P transactions are enabled.
	IsP2pEnabled *bool `json:"is_p2p_enabled,omitempty"`
	// A flag to indicate whether Synctera Pay transactions are enabled.
	IsSyncteraPayEnabled *bool `json:"is_synctera_pay_enabled,omitempty"`
	// A flag to indicate whether wire transactions are enabled.
	IsWireEnabled *bool                `json:"is_wire_enabled,omitempty"`
	AccessStatus  *AccountAccessStatus `json:"access_status,omitempty"`
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
	// A flag to indicate whether auto payments are enabled.
	IsSystemAutoPayEnabled *bool `json:"is_system_auto_pay_enabled,omitempty"`
	// Add an optional note when patching a charged secured account. A note is required when setting the status to or from SUSPENDED.
	Note *string `json:"note,omitempty"`
	// List of spend control IDs to control spending for the account
	SpendControlIds []string `json:"spend_control_ids,omitempty"`
}

// NewPatchAccountChargeSecured instantiates a new PatchAccountChargeSecured object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchAccountChargeSecured() *PatchAccountChargeSecured {
	this := PatchAccountChargeSecured{}
	return &this
}

// NewPatchAccountChargeSecuredWithDefaults instantiates a new PatchAccountChargeSecured object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchAccountChargeSecuredWithDefaults() *PatchAccountChargeSecured {
	this := PatchAccountChargeSecured{}
	return &this
}

// GetIsAchEnabled returns the IsAchEnabled field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIsAchEnabled() bool {
	if o == nil || IsNil(o.IsAchEnabled) {
		var ret bool
		return ret
	}
	return *o.IsAchEnabled
}

// GetIsAchEnabledOk returns a tuple with the IsAchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIsAchEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAchEnabled) {
		return nil, false
	}
	return o.IsAchEnabled, true
}

// HasIsAchEnabled returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIsAchEnabled() bool {
	if o != nil && !IsNil(o.IsAchEnabled) {
		return true
	}

	return false
}

// SetIsAchEnabled gets a reference to the given bool and assigns it to the IsAchEnabled field.
func (o *PatchAccountChargeSecured) SetIsAchEnabled(v bool) {
	o.IsAchEnabled = &v
}

// GetIsCardEnabled returns the IsCardEnabled field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIsCardEnabled() bool {
	if o == nil || IsNil(o.IsCardEnabled) {
		var ret bool
		return ret
	}
	return *o.IsCardEnabled
}

// GetIsCardEnabledOk returns a tuple with the IsCardEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIsCardEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCardEnabled) {
		return nil, false
	}
	return o.IsCardEnabled, true
}

// HasIsCardEnabled returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIsCardEnabled() bool {
	if o != nil && !IsNil(o.IsCardEnabled) {
		return true
	}

	return false
}

// SetIsCardEnabled gets a reference to the given bool and assigns it to the IsCardEnabled field.
func (o *PatchAccountChargeSecured) SetIsCardEnabled(v bool) {
	o.IsCardEnabled = &v
}

// GetIsEftCaEnabled returns the IsEftCaEnabled field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIsEftCaEnabled() bool {
	if o == nil || IsNil(o.IsEftCaEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEftCaEnabled
}

// GetIsEftCaEnabledOk returns a tuple with the IsEftCaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIsEftCaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEftCaEnabled) {
		return nil, false
	}
	return o.IsEftCaEnabled, true
}

// HasIsEftCaEnabled returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIsEftCaEnabled() bool {
	if o != nil && !IsNil(o.IsEftCaEnabled) {
		return true
	}

	return false
}

// SetIsEftCaEnabled gets a reference to the given bool and assigns it to the IsEftCaEnabled field.
func (o *PatchAccountChargeSecured) SetIsEftCaEnabled(v bool) {
	o.IsEftCaEnabled = &v
}

// GetIsP2pEnabled returns the IsP2pEnabled field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIsP2pEnabled() bool {
	if o == nil || IsNil(o.IsP2pEnabled) {
		var ret bool
		return ret
	}
	return *o.IsP2pEnabled
}

// GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIsP2pEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsP2pEnabled) {
		return nil, false
	}
	return o.IsP2pEnabled, true
}

// HasIsP2pEnabled returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIsP2pEnabled() bool {
	if o != nil && !IsNil(o.IsP2pEnabled) {
		return true
	}

	return false
}

// SetIsP2pEnabled gets a reference to the given bool and assigns it to the IsP2pEnabled field.
func (o *PatchAccountChargeSecured) SetIsP2pEnabled(v bool) {
	o.IsP2pEnabled = &v
}

// GetIsSyncteraPayEnabled returns the IsSyncteraPayEnabled field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIsSyncteraPayEnabled() bool {
	if o == nil || IsNil(o.IsSyncteraPayEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSyncteraPayEnabled
}

// GetIsSyncteraPayEnabledOk returns a tuple with the IsSyncteraPayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIsSyncteraPayEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncteraPayEnabled) {
		return nil, false
	}
	return o.IsSyncteraPayEnabled, true
}

// HasIsSyncteraPayEnabled returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIsSyncteraPayEnabled() bool {
	if o != nil && !IsNil(o.IsSyncteraPayEnabled) {
		return true
	}

	return false
}

// SetIsSyncteraPayEnabled gets a reference to the given bool and assigns it to the IsSyncteraPayEnabled field.
func (o *PatchAccountChargeSecured) SetIsSyncteraPayEnabled(v bool) {
	o.IsSyncteraPayEnabled = &v
}

// GetIsWireEnabled returns the IsWireEnabled field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIsWireEnabled() bool {
	if o == nil || IsNil(o.IsWireEnabled) {
		var ret bool
		return ret
	}
	return *o.IsWireEnabled
}

// GetIsWireEnabledOk returns a tuple with the IsWireEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIsWireEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWireEnabled) {
		return nil, false
	}
	return o.IsWireEnabled, true
}

// HasIsWireEnabled returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIsWireEnabled() bool {
	if o != nil && !IsNil(o.IsWireEnabled) {
		return true
	}

	return false
}

// SetIsWireEnabled gets a reference to the given bool and assigns it to the IsWireEnabled field.
func (o *PatchAccountChargeSecured) SetIsWireEnabled(v bool) {
	o.IsWireEnabled = &v
}

// GetAccessStatus returns the AccessStatus field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetAccessStatus() AccountAccessStatus {
	if o == nil || IsNil(o.AccessStatus) {
		var ret AccountAccessStatus
		return ret
	}
	return *o.AccessStatus
}

// GetAccessStatusOk returns a tuple with the AccessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetAccessStatusOk() (*AccountAccessStatus, bool) {
	if o == nil || IsNil(o.AccessStatus) {
		return nil, false
	}
	return o.AccessStatus, true
}

// HasAccessStatus returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasAccessStatus() bool {
	if o != nil && !IsNil(o.AccessStatus) {
		return true
	}

	return false
}

// SetAccessStatus gets a reference to the given AccountAccessStatus and assigns it to the AccessStatus field.
func (o *PatchAccountChargeSecured) SetAccessStatus(v AccountAccessStatus) {
	o.AccessStatus = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *PatchAccountChargeSecured) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountNumberMasked returns the AccountNumberMasked field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetAccountNumberMasked() string {
	if o == nil || IsNil(o.AccountNumberMasked) {
		var ret string
		return ret
	}
	return *o.AccountNumberMasked
}

// GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetAccountNumberMaskedOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumberMasked) {
		return nil, false
	}
	return o.AccountNumberMasked, true
}

// HasAccountNumberMasked returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasAccountNumberMasked() bool {
	if o != nil && !IsNil(o.AccountNumberMasked) {
		return true
	}

	return false
}

// SetAccountNumberMasked gets a reference to the given string and assigns it to the AccountNumberMasked field.
func (o *PatchAccountChargeSecured) SetAccountNumberMasked(v string) {
	o.AccountNumberMasked = &v
}

// GetAccountPurpose returns the AccountPurpose field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetAccountPurpose() string {
	if o == nil || IsNil(o.AccountPurpose) {
		var ret string
		return ret
	}
	return *o.AccountPurpose
}

// GetAccountPurposeOk returns a tuple with the AccountPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetAccountPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountPurpose) {
		return nil, false
	}
	return o.AccountPurpose, true
}

// HasAccountPurpose returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasAccountPurpose() bool {
	if o != nil && !IsNil(o.AccountPurpose) {
		return true
	}

	return false
}

// SetAccountPurpose gets a reference to the given string and assigns it to the AccountPurpose field.
func (o *PatchAccountChargeSecured) SetAccountPurpose(v string) {
	o.AccountPurpose = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetAccountType() AccountType {
	if o == nil || IsNil(o.AccountType) {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *PatchAccountChargeSecured) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *PatchAccountChargeSecured) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetBalances() []Balance {
	if o == nil || IsNil(o.Balances) {
		var ret []Balance
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetBalancesOk() ([]Balance, bool) {
	if o == nil || IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasBalances() bool {
	if o != nil && !IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []Balance and assigns it to the Balances field.
func (o *PatchAccountChargeSecured) SetBalances(v []Balance) {
	o.Balances = v
}

// GetBankRouting returns the BankRouting field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetBankRouting() string {
	if o == nil || IsNil(o.BankRouting) {
		var ret string
		return ret
	}
	return *o.BankRouting
}

// GetBankRoutingOk returns a tuple with the BankRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetBankRoutingOk() (*string, bool) {
	if o == nil || IsNil(o.BankRouting) {
		return nil, false
	}
	return o.BankRouting, true
}

// HasBankRouting returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasBankRouting() bool {
	if o != nil && !IsNil(o.BankRouting) {
		return true
	}

	return false
}

// SetBankRouting gets a reference to the given string and assigns it to the BankRouting field.
func (o *PatchAccountChargeSecured) SetBankRouting(v string) {
	o.BankRouting = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PatchAccountChargeSecured) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PatchAccountChargeSecured) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerIds returns the CustomerIds field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetCustomerIds() []string {
	if o == nil || IsNil(o.CustomerIds) {
		var ret []string
		return ret
	}
	return o.CustomerIds
}

// GetCustomerIdsOk returns a tuple with the CustomerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetCustomerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomerIds) {
		return nil, false
	}
	return o.CustomerIds, true
}

// HasCustomerIds returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasCustomerIds() bool {
	if o != nil && !IsNil(o.CustomerIds) {
		return true
	}

	return false
}

// SetCustomerIds gets a reference to the given []string and assigns it to the CustomerIds field.
func (o *PatchAccountChargeSecured) SetCustomerIds(v []string) {
	o.CustomerIds = v
}

// GetCustomerType returns the CustomerType field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetCustomerType() CustomerType {
	if o == nil || IsNil(o.CustomerType) {
		var ret CustomerType
		return ret
	}
	return *o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetCustomerTypeOk() (*CustomerType, bool) {
	if o == nil || IsNil(o.CustomerType) {
		return nil, false
	}
	return o.CustomerType, true
}

// HasCustomerType returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasCustomerType() bool {
	if o != nil && !IsNil(o.CustomerType) {
		return true
	}

	return false
}

// SetCustomerType gets a reference to the given CustomerType and assigns it to the CustomerType field.
func (o *PatchAccountChargeSecured) SetCustomerType(v CustomerType) {
	o.CustomerType = &v
}

// GetExchangeRateType returns the ExchangeRateType field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetExchangeRateType() string {
	if o == nil || IsNil(o.ExchangeRateType) {
		var ret string
		return ret
	}
	return *o.ExchangeRateType
}

// GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetExchangeRateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeRateType) {
		return nil, false
	}
	return o.ExchangeRateType, true
}

// HasExchangeRateType returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasExchangeRateType() bool {
	if o != nil && !IsNil(o.ExchangeRateType) {
		return true
	}

	return false
}

// SetExchangeRateType gets a reference to the given string and assigns it to the ExchangeRateType field.
func (o *PatchAccountChargeSecured) SetExchangeRateType(v string) {
	o.ExchangeRateType = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIban() string {
	if o == nil || IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIbanOk() (*string, bool) {
	if o == nil || IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIban() bool {
	if o != nil && !IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *PatchAccountChargeSecured) SetIban(v string) {
	o.Iban = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchAccountChargeSecured) SetId(v string) {
	o.Id = &v
}

// GetIsAccountPool returns the IsAccountPool field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIsAccountPool() bool {
	if o == nil || IsNil(o.IsAccountPool) {
		var ret bool
		return ret
	}
	return *o.IsAccountPool
}

// GetIsAccountPoolOk returns a tuple with the IsAccountPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIsAccountPoolOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAccountPool) {
		return nil, false
	}
	return o.IsAccountPool, true
}

// HasIsAccountPool returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIsAccountPool() bool {
	if o != nil && !IsNil(o.IsAccountPool) {
		return true
	}

	return false
}

// SetIsAccountPool gets a reference to the given bool and assigns it to the IsAccountPool field.
func (o *PatchAccountChargeSecured) SetIsAccountPool(v bool) {
	o.IsAccountPool = &v
}

// GetIsSarEnabled returns the IsSarEnabled field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIsSarEnabled() bool {
	if o == nil || IsNil(o.IsSarEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSarEnabled
}

// GetIsSarEnabledOk returns a tuple with the IsSarEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIsSarEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSarEnabled) {
		return nil, false
	}
	return o.IsSarEnabled, true
}

// HasIsSarEnabled returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIsSarEnabled() bool {
	if o != nil && !IsNil(o.IsSarEnabled) {
		return true
	}

	return false
}

// SetIsSarEnabled gets a reference to the given bool and assigns it to the IsSarEnabled field.
func (o *PatchAccountChargeSecured) SetIsSarEnabled(v bool) {
	o.IsSarEnabled = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *PatchAccountChargeSecured) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchAccountChargeSecured) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *PatchAccountChargeSecured) SetNickname(v string) {
	o.Nickname = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetStatus() AccountStatus {
	if o == nil || IsNil(o.Status) {
		var ret AccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetStatusOk() (*AccountStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AccountStatus and assigns it to the Status field.
func (o *PatchAccountChargeSecured) SetStatus(v AccountStatus) {
	o.Status = &v
}

// GetSwiftCode returns the SwiftCode field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetSwiftCode() string {
	if o == nil || IsNil(o.SwiftCode) {
		var ret string
		return ret
	}
	return *o.SwiftCode
}

// GetSwiftCodeOk returns a tuple with the SwiftCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetSwiftCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SwiftCode) {
		return nil, false
	}
	return o.SwiftCode, true
}

// HasSwiftCode returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasSwiftCode() bool {
	if o != nil && !IsNil(o.SwiftCode) {
		return true
	}

	return false
}

// SetSwiftCode gets a reference to the given string and assigns it to the SwiftCode field.
func (o *PatchAccountChargeSecured) SetSwiftCode(v string) {
	o.SwiftCode = &v
}

// GetIsSystemAutoPayEnabled returns the IsSystemAutoPayEnabled field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetIsSystemAutoPayEnabled() bool {
	if o == nil || IsNil(o.IsSystemAutoPayEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSystemAutoPayEnabled
}

// GetIsSystemAutoPayEnabledOk returns a tuple with the IsSystemAutoPayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetIsSystemAutoPayEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSystemAutoPayEnabled) {
		return nil, false
	}
	return o.IsSystemAutoPayEnabled, true
}

// HasIsSystemAutoPayEnabled returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasIsSystemAutoPayEnabled() bool {
	if o != nil && !IsNil(o.IsSystemAutoPayEnabled) {
		return true
	}

	return false
}

// SetIsSystemAutoPayEnabled gets a reference to the given bool and assigns it to the IsSystemAutoPayEnabled field.
func (o *PatchAccountChargeSecured) SetIsSystemAutoPayEnabled(v bool) {
	o.IsSystemAutoPayEnabled = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *PatchAccountChargeSecured) SetNote(v string) {
	o.Note = &v
}

// GetSpendControlIds returns the SpendControlIds field value if set, zero value otherwise.
func (o *PatchAccountChargeSecured) GetSpendControlIds() []string {
	if o == nil || IsNil(o.SpendControlIds) {
		var ret []string
		return ret
	}
	return o.SpendControlIds
}

// GetSpendControlIdsOk returns a tuple with the SpendControlIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountChargeSecured) GetSpendControlIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SpendControlIds) {
		return nil, false
	}
	return o.SpendControlIds, true
}

// HasSpendControlIds returns a boolean if a field has been set.
func (o *PatchAccountChargeSecured) HasSpendControlIds() bool {
	if o != nil && !IsNil(o.SpendControlIds) {
		return true
	}

	return false
}

// SetSpendControlIds gets a reference to the given []string and assigns it to the SpendControlIds field.
func (o *PatchAccountChargeSecured) SetSpendControlIds(v []string) {
	o.SpendControlIds = v
}

func (o PatchAccountChargeSecured) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchAccountChargeSecured) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAchEnabled) {
		toSerialize["is_ach_enabled"] = o.IsAchEnabled
	}
	if !IsNil(o.IsCardEnabled) {
		toSerialize["is_card_enabled"] = o.IsCardEnabled
	}
	if !IsNil(o.IsEftCaEnabled) {
		toSerialize["is_eft_ca_enabled"] = o.IsEftCaEnabled
	}
	if !IsNil(o.IsP2pEnabled) {
		toSerialize["is_p2p_enabled"] = o.IsP2pEnabled
	}
	if !IsNil(o.IsSyncteraPayEnabled) {
		toSerialize["is_synctera_pay_enabled"] = o.IsSyncteraPayEnabled
	}
	if !IsNil(o.IsWireEnabled) {
		toSerialize["is_wire_enabled"] = o.IsWireEnabled
	}
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
	if !IsNil(o.IsSystemAutoPayEnabled) {
		toSerialize["is_system_auto_pay_enabled"] = o.IsSystemAutoPayEnabled
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.SpendControlIds) {
		toSerialize["spend_control_ids"] = o.SpendControlIds
	}
	return toSerialize, nil
}

type NullablePatchAccountChargeSecured struct {
	value *PatchAccountChargeSecured
	isSet bool
}

func (v NullablePatchAccountChargeSecured) Get() *PatchAccountChargeSecured {
	return v.value
}

func (v *NullablePatchAccountChargeSecured) Set(val *PatchAccountChargeSecured) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchAccountChargeSecured) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchAccountChargeSecured) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchAccountChargeSecured(val *PatchAccountChargeSecured) *NullablePatchAccountChargeSecured {
	return &NullablePatchAccountChargeSecured{value: val, isSet: true}
}

func (v NullablePatchAccountChargeSecured) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchAccountChargeSecured) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
