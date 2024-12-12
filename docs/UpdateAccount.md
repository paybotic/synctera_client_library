# UpdateAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceCeiling** | Pointer to [**BalanceCeiling**](BalanceCeiling.md) |  | [optional] 
**BalanceFloor** | Pointer to [**BalanceFloor**](BalanceFloor.md) |  | [optional] 
**FeeProductIds** | Pointer to **[]string** | A list of fee account products that the current account associates with. | [optional] 
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with.  | [optional] 
**Note** | Pointer to **string** | Add an optional note when creating or updating an account. A note is required when updating the status to or from SUSPENDED | [optional] 
**OverdraftLimit** | Pointer to **int64** | This field is unused and will be removed in a future API version.  | [optional] 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 
**SpendingLimits** | Pointer to [**SpendingLimits**](SpendingLimits.md) |  | [optional] 
**IsAchEnabled** | Pointer to **bool** | A flag to indicate whether ACH transactions are enabled. | [optional] 
**IsCardEnabled** | Pointer to **bool** | A flag to indicate whether card transactions are enabled. | [optional] 
**IsEftCaEnabled** | Pointer to **bool** | A flag to indicate whether EFT Canada transactions are enabled. | [optional] 
**IsExternalCardEnabled** | Pointer to **bool** | A flag to indicate whether external card transactions are enabled. | [optional] 
**IsP2pEnabled** | Pointer to **bool** | A flag to indicate whether P2P transactions are enabled. | [optional] 
**IsSyncteraPayEnabled** | Pointer to **bool** | A flag to indicate whether Synctera Pay transactions are enabled. | [optional] 
**IsWireEnabled** | Pointer to **bool** | A flag to indicate whether wire transactions are enabled. | [optional] 
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

## Methods

### NewUpdateAccount

`func NewUpdateAccount() *UpdateAccount`

NewUpdateAccount instantiates a new UpdateAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountWithDefaults

`func NewUpdateAccountWithDefaults() *UpdateAccount`

NewUpdateAccountWithDefaults instantiates a new UpdateAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceCeiling

`func (o *UpdateAccount) GetBalanceCeiling() BalanceCeiling`

GetBalanceCeiling returns the BalanceCeiling field if non-nil, zero value otherwise.

### GetBalanceCeilingOk

`func (o *UpdateAccount) GetBalanceCeilingOk() (*BalanceCeiling, bool)`

GetBalanceCeilingOk returns a tuple with the BalanceCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCeiling

`func (o *UpdateAccount) SetBalanceCeiling(v BalanceCeiling)`

SetBalanceCeiling sets BalanceCeiling field to given value.

### HasBalanceCeiling

`func (o *UpdateAccount) HasBalanceCeiling() bool`

HasBalanceCeiling returns a boolean if a field has been set.

### GetBalanceFloor

`func (o *UpdateAccount) GetBalanceFloor() BalanceFloor`

GetBalanceFloor returns the BalanceFloor field if non-nil, zero value otherwise.

### GetBalanceFloorOk

`func (o *UpdateAccount) GetBalanceFloorOk() (*BalanceFloor, bool)`

GetBalanceFloorOk returns a tuple with the BalanceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloor

`func (o *UpdateAccount) SetBalanceFloor(v BalanceFloor)`

SetBalanceFloor sets BalanceFloor field to given value.

### HasBalanceFloor

`func (o *UpdateAccount) HasBalanceFloor() bool`

HasBalanceFloor returns a boolean if a field has been set.

### GetFeeProductIds

`func (o *UpdateAccount) GetFeeProductIds() []string`

GetFeeProductIds returns the FeeProductIds field if non-nil, zero value otherwise.

### GetFeeProductIdsOk

`func (o *UpdateAccount) GetFeeProductIdsOk() (*[]string, bool)`

GetFeeProductIdsOk returns a tuple with the FeeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeProductIds

`func (o *UpdateAccount) SetFeeProductIds(v []string)`

SetFeeProductIds sets FeeProductIds field to given value.

### HasFeeProductIds

`func (o *UpdateAccount) HasFeeProductIds() bool`

HasFeeProductIds returns a boolean if a field has been set.

### GetInterestProductId

`func (o *UpdateAccount) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *UpdateAccount) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *UpdateAccount) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *UpdateAccount) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetNote

`func (o *UpdateAccount) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UpdateAccount) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UpdateAccount) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UpdateAccount) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOverdraftLimit

`func (o *UpdateAccount) GetOverdraftLimit() int64`

GetOverdraftLimit returns the OverdraftLimit field if non-nil, zero value otherwise.

### GetOverdraftLimitOk

`func (o *UpdateAccount) GetOverdraftLimitOk() (*int64, bool)`

GetOverdraftLimitOk returns a tuple with the OverdraftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLimit

`func (o *UpdateAccount) SetOverdraftLimit(v int64)`

SetOverdraftLimit sets OverdraftLimit field to given value.

### HasOverdraftLimit

`func (o *UpdateAccount) HasOverdraftLimit() bool`

HasOverdraftLimit returns a boolean if a field has been set.

### GetSpendControlIds

`func (o *UpdateAccount) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *UpdateAccount) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *UpdateAccount) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *UpdateAccount) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.

### GetSpendingLimits

`func (o *UpdateAccount) GetSpendingLimits() SpendingLimits`

GetSpendingLimits returns the SpendingLimits field if non-nil, zero value otherwise.

### GetSpendingLimitsOk

`func (o *UpdateAccount) GetSpendingLimitsOk() (*SpendingLimits, bool)`

GetSpendingLimitsOk returns a tuple with the SpendingLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingLimits

`func (o *UpdateAccount) SetSpendingLimits(v SpendingLimits)`

SetSpendingLimits sets SpendingLimits field to given value.

### HasSpendingLimits

`func (o *UpdateAccount) HasSpendingLimits() bool`

HasSpendingLimits returns a boolean if a field has been set.

### GetIsAchEnabled

`func (o *UpdateAccount) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *UpdateAccount) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *UpdateAccount) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *UpdateAccount) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *UpdateAccount) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *UpdateAccount) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *UpdateAccount) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *UpdateAccount) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsEftCaEnabled

`func (o *UpdateAccount) GetIsEftCaEnabled() bool`

GetIsEftCaEnabled returns the IsEftCaEnabled field if non-nil, zero value otherwise.

### GetIsEftCaEnabledOk

`func (o *UpdateAccount) GetIsEftCaEnabledOk() (*bool, bool)`

GetIsEftCaEnabledOk returns a tuple with the IsEftCaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEftCaEnabled

`func (o *UpdateAccount) SetIsEftCaEnabled(v bool)`

SetIsEftCaEnabled sets IsEftCaEnabled field to given value.

### HasIsEftCaEnabled

`func (o *UpdateAccount) HasIsEftCaEnabled() bool`

HasIsEftCaEnabled returns a boolean if a field has been set.

### GetIsExternalCardEnabled

`func (o *UpdateAccount) GetIsExternalCardEnabled() bool`

GetIsExternalCardEnabled returns the IsExternalCardEnabled field if non-nil, zero value otherwise.

### GetIsExternalCardEnabledOk

`func (o *UpdateAccount) GetIsExternalCardEnabledOk() (*bool, bool)`

GetIsExternalCardEnabledOk returns a tuple with the IsExternalCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalCardEnabled

`func (o *UpdateAccount) SetIsExternalCardEnabled(v bool)`

SetIsExternalCardEnabled sets IsExternalCardEnabled field to given value.

### HasIsExternalCardEnabled

`func (o *UpdateAccount) HasIsExternalCardEnabled() bool`

HasIsExternalCardEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *UpdateAccount) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *UpdateAccount) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *UpdateAccount) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *UpdateAccount) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsSyncteraPayEnabled

`func (o *UpdateAccount) GetIsSyncteraPayEnabled() bool`

GetIsSyncteraPayEnabled returns the IsSyncteraPayEnabled field if non-nil, zero value otherwise.

### GetIsSyncteraPayEnabledOk

`func (o *UpdateAccount) GetIsSyncteraPayEnabledOk() (*bool, bool)`

GetIsSyncteraPayEnabledOk returns a tuple with the IsSyncteraPayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncteraPayEnabled

`func (o *UpdateAccount) SetIsSyncteraPayEnabled(v bool)`

SetIsSyncteraPayEnabled sets IsSyncteraPayEnabled field to given value.

### HasIsSyncteraPayEnabled

`func (o *UpdateAccount) HasIsSyncteraPayEnabled() bool`

HasIsSyncteraPayEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *UpdateAccount) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *UpdateAccount) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *UpdateAccount) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *UpdateAccount) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetAccessStatus

`func (o *UpdateAccount) GetAccessStatus() AccountAccessStatus`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *UpdateAccount) GetAccessStatusOk() (*AccountAccessStatus, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *UpdateAccount) SetAccessStatus(v AccountAccessStatus)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *UpdateAccount) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetAccountNumber

`func (o *UpdateAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *UpdateAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *UpdateAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *UpdateAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountNumberMasked

`func (o *UpdateAccount) GetAccountNumberMasked() string`

GetAccountNumberMasked returns the AccountNumberMasked field if non-nil, zero value otherwise.

### GetAccountNumberMaskedOk

`func (o *UpdateAccount) GetAccountNumberMaskedOk() (*string, bool)`

GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberMasked

`func (o *UpdateAccount) SetAccountNumberMasked(v string)`

SetAccountNumberMasked sets AccountNumberMasked field to given value.

### HasAccountNumberMasked

`func (o *UpdateAccount) HasAccountNumberMasked() bool`

HasAccountNumberMasked returns a boolean if a field has been set.

### GetAccountPurpose

`func (o *UpdateAccount) GetAccountPurpose() string`

GetAccountPurpose returns the AccountPurpose field if non-nil, zero value otherwise.

### GetAccountPurposeOk

`func (o *UpdateAccount) GetAccountPurposeOk() (*string, bool)`

GetAccountPurposeOk returns a tuple with the AccountPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPurpose

`func (o *UpdateAccount) SetAccountPurpose(v string)`

SetAccountPurpose sets AccountPurpose field to given value.

### HasAccountPurpose

`func (o *UpdateAccount) HasAccountPurpose() bool`

HasAccountPurpose returns a boolean if a field has been set.

### GetAccountType

`func (o *UpdateAccount) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UpdateAccount) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UpdateAccount) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UpdateAccount) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetApplicationId

`func (o *UpdateAccount) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *UpdateAccount) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *UpdateAccount) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *UpdateAccount) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBalances

`func (o *UpdateAccount) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *UpdateAccount) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *UpdateAccount) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *UpdateAccount) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankRouting

`func (o *UpdateAccount) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *UpdateAccount) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *UpdateAccount) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *UpdateAccount) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetCreationTime

`func (o *UpdateAccount) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *UpdateAccount) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *UpdateAccount) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *UpdateAccount) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerIds

`func (o *UpdateAccount) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *UpdateAccount) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *UpdateAccount) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *UpdateAccount) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetCustomerType

`func (o *UpdateAccount) GetCustomerType() CustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *UpdateAccount) GetCustomerTypeOk() (*CustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *UpdateAccount) SetCustomerType(v CustomerType)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *UpdateAccount) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *UpdateAccount) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *UpdateAccount) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *UpdateAccount) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *UpdateAccount) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetIban

`func (o *UpdateAccount) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *UpdateAccount) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *UpdateAccount) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *UpdateAccount) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *UpdateAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *UpdateAccount) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *UpdateAccount) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *UpdateAccount) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *UpdateAccount) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetIsSarEnabled

`func (o *UpdateAccount) GetIsSarEnabled() bool`

GetIsSarEnabled returns the IsSarEnabled field if non-nil, zero value otherwise.

### GetIsSarEnabledOk

`func (o *UpdateAccount) GetIsSarEnabledOk() (*bool, bool)`

GetIsSarEnabledOk returns a tuple with the IsSarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSarEnabled

`func (o *UpdateAccount) SetIsSarEnabled(v bool)`

SetIsSarEnabled sets IsSarEnabled field to given value.

### HasIsSarEnabled

`func (o *UpdateAccount) HasIsSarEnabled() bool`

HasIsSarEnabled returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *UpdateAccount) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *UpdateAccount) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *UpdateAccount) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *UpdateAccount) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateAccount) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateAccount) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateAccount) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateAccount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNickname

`func (o *UpdateAccount) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UpdateAccount) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UpdateAccount) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UpdateAccount) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateAccount) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAccount) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAccount) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwiftCode

`func (o *UpdateAccount) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *UpdateAccount) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *UpdateAccount) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *UpdateAccount) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


