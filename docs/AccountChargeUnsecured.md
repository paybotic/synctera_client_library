# AccountChargeUnsecured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessStatus** | Pointer to [**AccountAccessStatus**](AccountAccessStatus.md) |  | [optional] 
**AccountNumber** | Pointer to **string** | Account number | [optional] [readonly] 
**AccountNumberMasked** | Pointer to **string** | The response will contain the bank fintech ID (3 or 6 digits) plus the last 4 digits, with the digits in between replaced with * characters. Shadow mode account numbers will not be masked. | [optional] [readonly] 
**AccountPurpose** | Pointer to **string** | Purpose of the account | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**ApplicationId** | Pointer to **string** | The application ID for this account.  | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for account based on different type | [optional] [readonly] 
**BankRouting** | Pointer to **string** | Bank routing number | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** | Account creation timestamp in RFC3339 format | [optional] [readonly] 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**CustomerIds** | Pointer to **[]string** | A list of the customer IDs of the account holders. | [optional] [readonly] 
**CustomerType** | Pointer to [**CustomerType**](CustomerType.md) |  | [optional] 
**ExchangeRateType** | Pointer to **string** | Exchange rate type | [optional] 
**Iban** | Pointer to **string** | International bank account number | [optional] 
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**IsAccountPool** | Pointer to **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] 
**IsSarEnabled** | Pointer to **bool** | A flag to indicate whether SAR generation is enabled. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | Timestamp of the last account modification in RFC3339 format | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | User provided account metadata | [optional] 
**Nickname** | Pointer to **string** | User provided account nickname | [optional] 
**Status** | Pointer to [**AccountStatus**](AccountStatus.md) |  | [optional] 
**SwiftCode** | Pointer to **string** | SWIFT code | [optional] 
**IsAchEnabled** | Pointer to **bool** | A flag to indicate whether ACH transactions are enabled. | [optional] 
**IsCardEnabled** | Pointer to **bool** | A flag to indicate whether card transactions are enabled. | [optional] 
**IsEftCaEnabled** | Pointer to **bool** | A flag to indicate whether EFT Canada transactions are enabled. | [optional] 
**IsP2pEnabled** | Pointer to **bool** | A flag to indicate whether P2P transactions are enabled. | [optional] 
**IsSyncteraPayEnabled** | Pointer to **bool** | A flag to indicate whether Synctera Pay transactions are enabled. | [optional] 
**IsWireEnabled** | Pointer to **bool** | A flag to indicate whether wire transactions are enabled. | [optional] 
**BalanceCeiling** | Pointer to [**BalanceCeiling**](BalanceCeiling.md) |  | [optional] 
**BalanceFloor** | Pointer to [**BalanceFloor**](BalanceFloor.md) |  | [optional] 
**CreditLimit** | Pointer to **int64** | The credit limit for this charge unsecured account in cents. Minimum is 0.  | [optional] 
**MinimumPayment** | Pointer to [**MinimumPaymentFull**](MinimumPaymentFull.md) |  | [optional] 
**Note** | Pointer to **string** | Add an optional note when creating or updating an account. A note is required when updating the status to or from SUSPENDED | [optional] 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 

## Methods

### NewAccountChargeUnsecured

`func NewAccountChargeUnsecured() *AccountChargeUnsecured`

NewAccountChargeUnsecured instantiates a new AccountChargeUnsecured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountChargeUnsecuredWithDefaults

`func NewAccountChargeUnsecuredWithDefaults() *AccountChargeUnsecured`

NewAccountChargeUnsecuredWithDefaults instantiates a new AccountChargeUnsecured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessStatus

`func (o *AccountChargeUnsecured) GetAccessStatus() AccountAccessStatus`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *AccountChargeUnsecured) GetAccessStatusOk() (*AccountAccessStatus, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *AccountChargeUnsecured) SetAccessStatus(v AccountAccessStatus)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *AccountChargeUnsecured) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AccountChargeUnsecured) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountChargeUnsecured) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountChargeUnsecured) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountChargeUnsecured) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountNumberMasked

`func (o *AccountChargeUnsecured) GetAccountNumberMasked() string`

GetAccountNumberMasked returns the AccountNumberMasked field if non-nil, zero value otherwise.

### GetAccountNumberMaskedOk

`func (o *AccountChargeUnsecured) GetAccountNumberMaskedOk() (*string, bool)`

GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberMasked

`func (o *AccountChargeUnsecured) SetAccountNumberMasked(v string)`

SetAccountNumberMasked sets AccountNumberMasked field to given value.

### HasAccountNumberMasked

`func (o *AccountChargeUnsecured) HasAccountNumberMasked() bool`

HasAccountNumberMasked returns a boolean if a field has been set.

### GetAccountPurpose

`func (o *AccountChargeUnsecured) GetAccountPurpose() string`

GetAccountPurpose returns the AccountPurpose field if non-nil, zero value otherwise.

### GetAccountPurposeOk

`func (o *AccountChargeUnsecured) GetAccountPurposeOk() (*string, bool)`

GetAccountPurposeOk returns a tuple with the AccountPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPurpose

`func (o *AccountChargeUnsecured) SetAccountPurpose(v string)`

SetAccountPurpose sets AccountPurpose field to given value.

### HasAccountPurpose

`func (o *AccountChargeUnsecured) HasAccountPurpose() bool`

HasAccountPurpose returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountChargeUnsecured) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountChargeUnsecured) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountChargeUnsecured) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountChargeUnsecured) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetApplicationId

`func (o *AccountChargeUnsecured) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AccountChargeUnsecured) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AccountChargeUnsecured) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AccountChargeUnsecured) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBalances

`func (o *AccountChargeUnsecured) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountChargeUnsecured) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountChargeUnsecured) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *AccountChargeUnsecured) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankRouting

`func (o *AccountChargeUnsecured) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *AccountChargeUnsecured) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *AccountChargeUnsecured) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *AccountChargeUnsecured) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetCreationTime

`func (o *AccountChargeUnsecured) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountChargeUnsecured) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountChargeUnsecured) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountChargeUnsecured) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountChargeUnsecured) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountChargeUnsecured) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountChargeUnsecured) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountChargeUnsecured) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerIds

`func (o *AccountChargeUnsecured) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *AccountChargeUnsecured) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *AccountChargeUnsecured) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *AccountChargeUnsecured) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetCustomerType

`func (o *AccountChargeUnsecured) GetCustomerType() CustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AccountChargeUnsecured) GetCustomerTypeOk() (*CustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AccountChargeUnsecured) SetCustomerType(v CustomerType)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *AccountChargeUnsecured) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *AccountChargeUnsecured) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *AccountChargeUnsecured) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *AccountChargeUnsecured) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *AccountChargeUnsecured) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetIban

`func (o *AccountChargeUnsecured) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountChargeUnsecured) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountChargeUnsecured) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountChargeUnsecured) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *AccountChargeUnsecured) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountChargeUnsecured) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountChargeUnsecured) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountChargeUnsecured) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *AccountChargeUnsecured) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *AccountChargeUnsecured) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *AccountChargeUnsecured) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *AccountChargeUnsecured) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetIsSarEnabled

`func (o *AccountChargeUnsecured) GetIsSarEnabled() bool`

GetIsSarEnabled returns the IsSarEnabled field if non-nil, zero value otherwise.

### GetIsSarEnabledOk

`func (o *AccountChargeUnsecured) GetIsSarEnabledOk() (*bool, bool)`

GetIsSarEnabledOk returns a tuple with the IsSarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSarEnabled

`func (o *AccountChargeUnsecured) SetIsSarEnabled(v bool)`

SetIsSarEnabled sets IsSarEnabled field to given value.

### HasIsSarEnabled

`func (o *AccountChargeUnsecured) HasIsSarEnabled() bool`

HasIsSarEnabled returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AccountChargeUnsecured) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AccountChargeUnsecured) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AccountChargeUnsecured) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AccountChargeUnsecured) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountChargeUnsecured) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountChargeUnsecured) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountChargeUnsecured) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountChargeUnsecured) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNickname

`func (o *AccountChargeUnsecured) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountChargeUnsecured) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountChargeUnsecured) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountChargeUnsecured) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetStatus

`func (o *AccountChargeUnsecured) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountChargeUnsecured) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountChargeUnsecured) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountChargeUnsecured) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwiftCode

`func (o *AccountChargeUnsecured) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountChargeUnsecured) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountChargeUnsecured) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *AccountChargeUnsecured) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetIsAchEnabled

`func (o *AccountChargeUnsecured) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *AccountChargeUnsecured) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *AccountChargeUnsecured) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *AccountChargeUnsecured) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *AccountChargeUnsecured) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *AccountChargeUnsecured) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *AccountChargeUnsecured) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *AccountChargeUnsecured) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsEftCaEnabled

`func (o *AccountChargeUnsecured) GetIsEftCaEnabled() bool`

GetIsEftCaEnabled returns the IsEftCaEnabled field if non-nil, zero value otherwise.

### GetIsEftCaEnabledOk

`func (o *AccountChargeUnsecured) GetIsEftCaEnabledOk() (*bool, bool)`

GetIsEftCaEnabledOk returns a tuple with the IsEftCaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEftCaEnabled

`func (o *AccountChargeUnsecured) SetIsEftCaEnabled(v bool)`

SetIsEftCaEnabled sets IsEftCaEnabled field to given value.

### HasIsEftCaEnabled

`func (o *AccountChargeUnsecured) HasIsEftCaEnabled() bool`

HasIsEftCaEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *AccountChargeUnsecured) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *AccountChargeUnsecured) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *AccountChargeUnsecured) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *AccountChargeUnsecured) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsSyncteraPayEnabled

`func (o *AccountChargeUnsecured) GetIsSyncteraPayEnabled() bool`

GetIsSyncteraPayEnabled returns the IsSyncteraPayEnabled field if non-nil, zero value otherwise.

### GetIsSyncteraPayEnabledOk

`func (o *AccountChargeUnsecured) GetIsSyncteraPayEnabledOk() (*bool, bool)`

GetIsSyncteraPayEnabledOk returns a tuple with the IsSyncteraPayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncteraPayEnabled

`func (o *AccountChargeUnsecured) SetIsSyncteraPayEnabled(v bool)`

SetIsSyncteraPayEnabled sets IsSyncteraPayEnabled field to given value.

### HasIsSyncteraPayEnabled

`func (o *AccountChargeUnsecured) HasIsSyncteraPayEnabled() bool`

HasIsSyncteraPayEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *AccountChargeUnsecured) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *AccountChargeUnsecured) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *AccountChargeUnsecured) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *AccountChargeUnsecured) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetBalanceCeiling

`func (o *AccountChargeUnsecured) GetBalanceCeiling() BalanceCeiling`

GetBalanceCeiling returns the BalanceCeiling field if non-nil, zero value otherwise.

### GetBalanceCeilingOk

`func (o *AccountChargeUnsecured) GetBalanceCeilingOk() (*BalanceCeiling, bool)`

GetBalanceCeilingOk returns a tuple with the BalanceCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCeiling

`func (o *AccountChargeUnsecured) SetBalanceCeiling(v BalanceCeiling)`

SetBalanceCeiling sets BalanceCeiling field to given value.

### HasBalanceCeiling

`func (o *AccountChargeUnsecured) HasBalanceCeiling() bool`

HasBalanceCeiling returns a boolean if a field has been set.

### GetBalanceFloor

`func (o *AccountChargeUnsecured) GetBalanceFloor() BalanceFloor`

GetBalanceFloor returns the BalanceFloor field if non-nil, zero value otherwise.

### GetBalanceFloorOk

`func (o *AccountChargeUnsecured) GetBalanceFloorOk() (*BalanceFloor, bool)`

GetBalanceFloorOk returns a tuple with the BalanceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloor

`func (o *AccountChargeUnsecured) SetBalanceFloor(v BalanceFloor)`

SetBalanceFloor sets BalanceFloor field to given value.

### HasBalanceFloor

`func (o *AccountChargeUnsecured) HasBalanceFloor() bool`

HasBalanceFloor returns a boolean if a field has been set.

### GetCreditLimit

`func (o *AccountChargeUnsecured) GetCreditLimit() int64`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountChargeUnsecured) GetCreditLimitOk() (*int64, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountChargeUnsecured) SetCreditLimit(v int64)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *AccountChargeUnsecured) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *AccountChargeUnsecured) GetMinimumPayment() MinimumPaymentFull`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *AccountChargeUnsecured) GetMinimumPaymentOk() (*MinimumPaymentFull, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *AccountChargeUnsecured) SetMinimumPayment(v MinimumPaymentFull)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *AccountChargeUnsecured) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.

### GetNote

`func (o *AccountChargeUnsecured) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AccountChargeUnsecured) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AccountChargeUnsecured) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AccountChargeUnsecured) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetSpendControlIds

`func (o *AccountChargeUnsecured) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *AccountChargeUnsecured) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *AccountChargeUnsecured) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *AccountChargeUnsecured) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


