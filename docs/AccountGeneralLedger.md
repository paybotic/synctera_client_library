# AccountGeneralLedger

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
**GeneralLedgerType** | Pointer to [**GeneralLedgerType**](GeneralLedgerType.md) |  | [optional] 
**Note** | Pointer to **string** | Add an optional note when creating or updating an account. A note is required when updating the status to or from SUSPENDED | [optional] 

## Methods

### NewAccountGeneralLedger

`func NewAccountGeneralLedger() *AccountGeneralLedger`

NewAccountGeneralLedger instantiates a new AccountGeneralLedger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGeneralLedgerWithDefaults

`func NewAccountGeneralLedgerWithDefaults() *AccountGeneralLedger`

NewAccountGeneralLedgerWithDefaults instantiates a new AccountGeneralLedger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAchEnabled

`func (o *AccountGeneralLedger) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *AccountGeneralLedger) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *AccountGeneralLedger) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *AccountGeneralLedger) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *AccountGeneralLedger) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *AccountGeneralLedger) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *AccountGeneralLedger) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *AccountGeneralLedger) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsEftCaEnabled

`func (o *AccountGeneralLedger) GetIsEftCaEnabled() bool`

GetIsEftCaEnabled returns the IsEftCaEnabled field if non-nil, zero value otherwise.

### GetIsEftCaEnabledOk

`func (o *AccountGeneralLedger) GetIsEftCaEnabledOk() (*bool, bool)`

GetIsEftCaEnabledOk returns a tuple with the IsEftCaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEftCaEnabled

`func (o *AccountGeneralLedger) SetIsEftCaEnabled(v bool)`

SetIsEftCaEnabled sets IsEftCaEnabled field to given value.

### HasIsEftCaEnabled

`func (o *AccountGeneralLedger) HasIsEftCaEnabled() bool`

HasIsEftCaEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *AccountGeneralLedger) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *AccountGeneralLedger) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *AccountGeneralLedger) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *AccountGeneralLedger) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsSyncteraPayEnabled

`func (o *AccountGeneralLedger) GetIsSyncteraPayEnabled() bool`

GetIsSyncteraPayEnabled returns the IsSyncteraPayEnabled field if non-nil, zero value otherwise.

### GetIsSyncteraPayEnabledOk

`func (o *AccountGeneralLedger) GetIsSyncteraPayEnabledOk() (*bool, bool)`

GetIsSyncteraPayEnabledOk returns a tuple with the IsSyncteraPayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncteraPayEnabled

`func (o *AccountGeneralLedger) SetIsSyncteraPayEnabled(v bool)`

SetIsSyncteraPayEnabled sets IsSyncteraPayEnabled field to given value.

### HasIsSyncteraPayEnabled

`func (o *AccountGeneralLedger) HasIsSyncteraPayEnabled() bool`

HasIsSyncteraPayEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *AccountGeneralLedger) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *AccountGeneralLedger) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *AccountGeneralLedger) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *AccountGeneralLedger) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetAccessStatus

`func (o *AccountGeneralLedger) GetAccessStatus() AccountAccessStatus`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *AccountGeneralLedger) GetAccessStatusOk() (*AccountAccessStatus, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *AccountGeneralLedger) SetAccessStatus(v AccountAccessStatus)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *AccountGeneralLedger) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AccountGeneralLedger) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountGeneralLedger) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountGeneralLedger) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountGeneralLedger) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountNumberMasked

`func (o *AccountGeneralLedger) GetAccountNumberMasked() string`

GetAccountNumberMasked returns the AccountNumberMasked field if non-nil, zero value otherwise.

### GetAccountNumberMaskedOk

`func (o *AccountGeneralLedger) GetAccountNumberMaskedOk() (*string, bool)`

GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberMasked

`func (o *AccountGeneralLedger) SetAccountNumberMasked(v string)`

SetAccountNumberMasked sets AccountNumberMasked field to given value.

### HasAccountNumberMasked

`func (o *AccountGeneralLedger) HasAccountNumberMasked() bool`

HasAccountNumberMasked returns a boolean if a field has been set.

### GetAccountPurpose

`func (o *AccountGeneralLedger) GetAccountPurpose() string`

GetAccountPurpose returns the AccountPurpose field if non-nil, zero value otherwise.

### GetAccountPurposeOk

`func (o *AccountGeneralLedger) GetAccountPurposeOk() (*string, bool)`

GetAccountPurposeOk returns a tuple with the AccountPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPurpose

`func (o *AccountGeneralLedger) SetAccountPurpose(v string)`

SetAccountPurpose sets AccountPurpose field to given value.

### HasAccountPurpose

`func (o *AccountGeneralLedger) HasAccountPurpose() bool`

HasAccountPurpose returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountGeneralLedger) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountGeneralLedger) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountGeneralLedger) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountGeneralLedger) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetApplicationId

`func (o *AccountGeneralLedger) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AccountGeneralLedger) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AccountGeneralLedger) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AccountGeneralLedger) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBalances

`func (o *AccountGeneralLedger) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountGeneralLedger) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountGeneralLedger) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *AccountGeneralLedger) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankRouting

`func (o *AccountGeneralLedger) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *AccountGeneralLedger) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *AccountGeneralLedger) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *AccountGeneralLedger) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetCreationTime

`func (o *AccountGeneralLedger) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountGeneralLedger) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountGeneralLedger) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountGeneralLedger) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountGeneralLedger) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountGeneralLedger) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountGeneralLedger) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountGeneralLedger) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerIds

`func (o *AccountGeneralLedger) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *AccountGeneralLedger) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *AccountGeneralLedger) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *AccountGeneralLedger) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetCustomerType

`func (o *AccountGeneralLedger) GetCustomerType() CustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AccountGeneralLedger) GetCustomerTypeOk() (*CustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AccountGeneralLedger) SetCustomerType(v CustomerType)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *AccountGeneralLedger) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *AccountGeneralLedger) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *AccountGeneralLedger) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *AccountGeneralLedger) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *AccountGeneralLedger) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetIban

`func (o *AccountGeneralLedger) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountGeneralLedger) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountGeneralLedger) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountGeneralLedger) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *AccountGeneralLedger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountGeneralLedger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountGeneralLedger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountGeneralLedger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *AccountGeneralLedger) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *AccountGeneralLedger) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *AccountGeneralLedger) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *AccountGeneralLedger) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetIsSarEnabled

`func (o *AccountGeneralLedger) GetIsSarEnabled() bool`

GetIsSarEnabled returns the IsSarEnabled field if non-nil, zero value otherwise.

### GetIsSarEnabledOk

`func (o *AccountGeneralLedger) GetIsSarEnabledOk() (*bool, bool)`

GetIsSarEnabledOk returns a tuple with the IsSarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSarEnabled

`func (o *AccountGeneralLedger) SetIsSarEnabled(v bool)`

SetIsSarEnabled sets IsSarEnabled field to given value.

### HasIsSarEnabled

`func (o *AccountGeneralLedger) HasIsSarEnabled() bool`

HasIsSarEnabled returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AccountGeneralLedger) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AccountGeneralLedger) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AccountGeneralLedger) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AccountGeneralLedger) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountGeneralLedger) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountGeneralLedger) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountGeneralLedger) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountGeneralLedger) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNickname

`func (o *AccountGeneralLedger) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountGeneralLedger) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountGeneralLedger) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountGeneralLedger) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetStatus

`func (o *AccountGeneralLedger) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountGeneralLedger) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountGeneralLedger) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountGeneralLedger) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwiftCode

`func (o *AccountGeneralLedger) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountGeneralLedger) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountGeneralLedger) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *AccountGeneralLedger) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetGeneralLedgerType

`func (o *AccountGeneralLedger) GetGeneralLedgerType() GeneralLedgerType`

GetGeneralLedgerType returns the GeneralLedgerType field if non-nil, zero value otherwise.

### GetGeneralLedgerTypeOk

`func (o *AccountGeneralLedger) GetGeneralLedgerTypeOk() (*GeneralLedgerType, bool)`

GetGeneralLedgerTypeOk returns a tuple with the GeneralLedgerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralLedgerType

`func (o *AccountGeneralLedger) SetGeneralLedgerType(v GeneralLedgerType)`

SetGeneralLedgerType sets GeneralLedgerType field to given value.

### HasGeneralLedgerType

`func (o *AccountGeneralLedger) HasGeneralLedgerType() bool`

HasGeneralLedgerType returns a boolean if a field has been set.

### GetNote

`func (o *AccountGeneralLedger) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AccountGeneralLedger) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AccountGeneralLedger) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AccountGeneralLedger) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


