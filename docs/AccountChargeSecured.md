# AccountChargeSecured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAchEnabled** | Pointer to **bool** | A flag to indicate whether ACH transactions are enabled. | [optional] 
**IsCardEnabled** | Pointer to **bool** | A flag to indicate whether card transactions are enabled. | [optional] 
**IsEftCaEnabled** | Pointer to **bool** | A flag to indicate whether EFT Canada transactions are enabled. | [optional] 
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
**IsSystemAutoPayEnabled** | Pointer to **bool** | A flag to indicate whether auto payments are enabled. | [optional] [default to false]
**Note** | Pointer to **string** | Add an optional note when creating or updating an account. A note is required when updating the status to or from SUSPENDED | [optional] 
**Security** | Pointer to [**Security**](Security.md) |  | [optional] 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 

## Methods

### NewAccountChargeSecured

`func NewAccountChargeSecured() *AccountChargeSecured`

NewAccountChargeSecured instantiates a new AccountChargeSecured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountChargeSecuredWithDefaults

`func NewAccountChargeSecuredWithDefaults() *AccountChargeSecured`

NewAccountChargeSecuredWithDefaults instantiates a new AccountChargeSecured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAchEnabled

`func (o *AccountChargeSecured) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *AccountChargeSecured) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *AccountChargeSecured) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *AccountChargeSecured) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *AccountChargeSecured) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *AccountChargeSecured) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *AccountChargeSecured) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *AccountChargeSecured) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsEftCaEnabled

`func (o *AccountChargeSecured) GetIsEftCaEnabled() bool`

GetIsEftCaEnabled returns the IsEftCaEnabled field if non-nil, zero value otherwise.

### GetIsEftCaEnabledOk

`func (o *AccountChargeSecured) GetIsEftCaEnabledOk() (*bool, bool)`

GetIsEftCaEnabledOk returns a tuple with the IsEftCaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEftCaEnabled

`func (o *AccountChargeSecured) SetIsEftCaEnabled(v bool)`

SetIsEftCaEnabled sets IsEftCaEnabled field to given value.

### HasIsEftCaEnabled

`func (o *AccountChargeSecured) HasIsEftCaEnabled() bool`

HasIsEftCaEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *AccountChargeSecured) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *AccountChargeSecured) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *AccountChargeSecured) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *AccountChargeSecured) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsSyncteraPayEnabled

`func (o *AccountChargeSecured) GetIsSyncteraPayEnabled() bool`

GetIsSyncteraPayEnabled returns the IsSyncteraPayEnabled field if non-nil, zero value otherwise.

### GetIsSyncteraPayEnabledOk

`func (o *AccountChargeSecured) GetIsSyncteraPayEnabledOk() (*bool, bool)`

GetIsSyncteraPayEnabledOk returns a tuple with the IsSyncteraPayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncteraPayEnabled

`func (o *AccountChargeSecured) SetIsSyncteraPayEnabled(v bool)`

SetIsSyncteraPayEnabled sets IsSyncteraPayEnabled field to given value.

### HasIsSyncteraPayEnabled

`func (o *AccountChargeSecured) HasIsSyncteraPayEnabled() bool`

HasIsSyncteraPayEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *AccountChargeSecured) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *AccountChargeSecured) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *AccountChargeSecured) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *AccountChargeSecured) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetAccessStatus

`func (o *AccountChargeSecured) GetAccessStatus() AccountAccessStatus`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *AccountChargeSecured) GetAccessStatusOk() (*AccountAccessStatus, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *AccountChargeSecured) SetAccessStatus(v AccountAccessStatus)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *AccountChargeSecured) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AccountChargeSecured) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountChargeSecured) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountChargeSecured) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountChargeSecured) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountNumberMasked

`func (o *AccountChargeSecured) GetAccountNumberMasked() string`

GetAccountNumberMasked returns the AccountNumberMasked field if non-nil, zero value otherwise.

### GetAccountNumberMaskedOk

`func (o *AccountChargeSecured) GetAccountNumberMaskedOk() (*string, bool)`

GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberMasked

`func (o *AccountChargeSecured) SetAccountNumberMasked(v string)`

SetAccountNumberMasked sets AccountNumberMasked field to given value.

### HasAccountNumberMasked

`func (o *AccountChargeSecured) HasAccountNumberMasked() bool`

HasAccountNumberMasked returns a boolean if a field has been set.

### GetAccountPurpose

`func (o *AccountChargeSecured) GetAccountPurpose() string`

GetAccountPurpose returns the AccountPurpose field if non-nil, zero value otherwise.

### GetAccountPurposeOk

`func (o *AccountChargeSecured) GetAccountPurposeOk() (*string, bool)`

GetAccountPurposeOk returns a tuple with the AccountPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPurpose

`func (o *AccountChargeSecured) SetAccountPurpose(v string)`

SetAccountPurpose sets AccountPurpose field to given value.

### HasAccountPurpose

`func (o *AccountChargeSecured) HasAccountPurpose() bool`

HasAccountPurpose returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountChargeSecured) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountChargeSecured) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountChargeSecured) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountChargeSecured) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetApplicationId

`func (o *AccountChargeSecured) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AccountChargeSecured) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AccountChargeSecured) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AccountChargeSecured) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBalances

`func (o *AccountChargeSecured) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountChargeSecured) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountChargeSecured) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *AccountChargeSecured) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankRouting

`func (o *AccountChargeSecured) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *AccountChargeSecured) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *AccountChargeSecured) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *AccountChargeSecured) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetCreationTime

`func (o *AccountChargeSecured) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountChargeSecured) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountChargeSecured) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountChargeSecured) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountChargeSecured) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountChargeSecured) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountChargeSecured) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountChargeSecured) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerIds

`func (o *AccountChargeSecured) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *AccountChargeSecured) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *AccountChargeSecured) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *AccountChargeSecured) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetCustomerType

`func (o *AccountChargeSecured) GetCustomerType() CustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AccountChargeSecured) GetCustomerTypeOk() (*CustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AccountChargeSecured) SetCustomerType(v CustomerType)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *AccountChargeSecured) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *AccountChargeSecured) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *AccountChargeSecured) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *AccountChargeSecured) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *AccountChargeSecured) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetIban

`func (o *AccountChargeSecured) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountChargeSecured) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountChargeSecured) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountChargeSecured) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *AccountChargeSecured) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountChargeSecured) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountChargeSecured) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountChargeSecured) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *AccountChargeSecured) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *AccountChargeSecured) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *AccountChargeSecured) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *AccountChargeSecured) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetIsSarEnabled

`func (o *AccountChargeSecured) GetIsSarEnabled() bool`

GetIsSarEnabled returns the IsSarEnabled field if non-nil, zero value otherwise.

### GetIsSarEnabledOk

`func (o *AccountChargeSecured) GetIsSarEnabledOk() (*bool, bool)`

GetIsSarEnabledOk returns a tuple with the IsSarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSarEnabled

`func (o *AccountChargeSecured) SetIsSarEnabled(v bool)`

SetIsSarEnabled sets IsSarEnabled field to given value.

### HasIsSarEnabled

`func (o *AccountChargeSecured) HasIsSarEnabled() bool`

HasIsSarEnabled returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AccountChargeSecured) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AccountChargeSecured) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AccountChargeSecured) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AccountChargeSecured) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountChargeSecured) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountChargeSecured) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountChargeSecured) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountChargeSecured) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNickname

`func (o *AccountChargeSecured) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountChargeSecured) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountChargeSecured) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountChargeSecured) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetStatus

`func (o *AccountChargeSecured) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountChargeSecured) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountChargeSecured) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountChargeSecured) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwiftCode

`func (o *AccountChargeSecured) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountChargeSecured) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountChargeSecured) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *AccountChargeSecured) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetIsSystemAutoPayEnabled

`func (o *AccountChargeSecured) GetIsSystemAutoPayEnabled() bool`

GetIsSystemAutoPayEnabled returns the IsSystemAutoPayEnabled field if non-nil, zero value otherwise.

### GetIsSystemAutoPayEnabledOk

`func (o *AccountChargeSecured) GetIsSystemAutoPayEnabledOk() (*bool, bool)`

GetIsSystemAutoPayEnabledOk returns a tuple with the IsSystemAutoPayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAutoPayEnabled

`func (o *AccountChargeSecured) SetIsSystemAutoPayEnabled(v bool)`

SetIsSystemAutoPayEnabled sets IsSystemAutoPayEnabled field to given value.

### HasIsSystemAutoPayEnabled

`func (o *AccountChargeSecured) HasIsSystemAutoPayEnabled() bool`

HasIsSystemAutoPayEnabled returns a boolean if a field has been set.

### GetNote

`func (o *AccountChargeSecured) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AccountChargeSecured) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AccountChargeSecured) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AccountChargeSecured) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetSecurity

`func (o *AccountChargeSecured) GetSecurity() Security`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *AccountChargeSecured) GetSecurityOk() (*Security, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *AccountChargeSecured) SetSecurity(v Security)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *AccountChargeSecured) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSpendControlIds

`func (o *AccountChargeSecured) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *AccountChargeSecured) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *AccountChargeSecured) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *AccountChargeSecured) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


