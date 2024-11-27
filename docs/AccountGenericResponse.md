# AccountGenericResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessStatus** | Pointer to [**AccountAccessStatus**](AccountAccessStatus.md) |  | [optional] 
**AccessStatusLastUpdatedTime** | Pointer to **time.Time** | Timestamp of the last modification of the access_status. RFC3339 format. | [optional] [readonly] 
**AccountClosure** | Pointer to [**AccountClosure**](AccountClosure.md) |  | [optional] 
**AccountNumber** | Pointer to **string** | Account number | [optional] [readonly] 
**AccountNumberMasked** | Pointer to **string** | The response will contain the bank fintech ID (3 or 6 digits) plus the last 4 digits, with the digits in between replaced with * characters. Shadow mode account numbers will not be masked. | [optional] [readonly] 
**AccountPurpose** | Pointer to **string** | Purpose of the account | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**ApplicationId** | Pointer to **string** | The application ID for this account.  | [optional] 
**AutoPaymentPeriod** | Pointer to **int32** | The number of days past the billing period to initiate an auto payment. Only applicable for accounts with type &#x60;CHARGE_SECURED&#x60;, where the account holder has opted in for auto payment functionality. This value must be lower than or equal the &#x60;grace_period&#x60; setting on the account. If this value is 0, the auto payment will happen on the same day as the statement is generated. Auto payment only occurs if regular payments are not received on time.  | [optional] 
**BalanceCeiling** | Pointer to [**BalanceCeiling**](BalanceCeiling.md) |  | [optional] 
**BalanceFloor** | Pointer to [**BalanceFloor**](BalanceFloor.md) |  | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for account based on different type | [optional] [readonly] 
**BankAccountId** | Pointer to **string** | Identifier of the bank side account that this account is a part of | [optional] [readonly] 
**BankRouting** | Pointer to **string** | Bank routing number | [optional] [readonly] 
**BillingPeriod** | Pointer to [**BillingPeriod**](BillingPeriod.md) |  | [optional] 
**BusinessIds** | Pointer to **[]string** | A list of the business IDs of the account holders. | [optional] [readonly] 
**CloseDate** | Pointer to **string** | The account close date. This is the bank&#39;s posting date when the account resource&#39;s status was changed to CLOSED or CHARGED_OFF. | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** | Account creation timestamp in RFC3339 format | [optional] [readonly] 
**CreditLimit** | Pointer to **int64** | The credit limit for this line of credit account in cents. Minimum is 0.  | [optional] 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**CustomerIds** | Pointer to **[]string** | A list of the customer IDs of the account holders. | [optional] [readonly] 
**CustomerType** | Pointer to [**CustomerType**](CustomerType.md) |  | [optional] 
**DaysPastDue** | Pointer to **int32** | The number of days since the account went past due on their minimum payments. | [optional] 
**ExchangeRateType** | Pointer to **string** | Exchange rate type | [optional] 
**FeeProductIds** | Pointer to **[]string** | A list of fee account products that the current account associates with. | [optional] 
**FundsOwnership** | Pointer to [**FundsOwnership**](FundsOwnership.md) |  | [optional] 
**GeneralLedgerCategory** | Pointer to [**GeneralLedgerCategory**](GeneralLedgerCategory.md) |  | [optional] 
**GeneralLedgerType** | Pointer to [**GeneralLedgerType**](GeneralLedgerType.md) |  | [optional] 
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment. The default will be set to 21 days.  | [optional] 
**Iban** | Pointer to **string** | International bank account number | [optional] 
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with. | [optional] 
**IsAccountPool** | Pointer to **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] 
**IsAchEnabled** | Pointer to **bool** | A flag to indicate whether ACH transactions are enabled. | [optional] [readonly] 
**IsCardEnabled** | Pointer to **bool** | A flag to indicate whether card transactions are enabled. | [optional] [readonly] 
**IsEftCaEnabled** | Pointer to **bool** | A flag to indicate whether EFT Canada transactions are enabled. | [optional] [readonly] 
**IsExternalCardEnabled** | Pointer to **bool** | A flag to indicate whether external card transactions are enabled. | [optional] [readonly] 
**IsP2pEnabled** | Pointer to **bool** | A flag to indicate whether P2P transactions are enabled. | [optional] [readonly] 
**IsSarEnabled** | Pointer to **bool** | A flag to indicate whether SAR generation is enabled. | [optional] [readonly] 
**IsSecurity** | Pointer to **bool** | A flag to indicate whether this account is being used as security for another account. | [optional] [readonly] 
**IsSyncteraPayEnabled** | Pointer to **bool** | A flag to indicate whether Synctera Pay transactions are enabled. | [optional] [readonly] 
**IsSystemAutoPayEnabled** | Pointer to **bool** | A flag to indicate whether auto pay feature is enabled. | [optional] 
**IsWireEnabled** | Pointer to **bool** | A flag to indicate whether wire transactions are enabled. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | Timestamp of the last account modification in RFC3339 format | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | User provided account metadata | [optional] 
**MinimumPayment** | Pointer to [**MinimumPayment**](MinimumPayment.md) |  | [optional] 
**Nickname** | Pointer to **string** | User provided account nickname | [optional] 
**OpenDate** | Pointer to **string** | The account open date. This is the bank&#39;s posting date when the account resource was created. | [optional] [readonly] 
**OverdraftLimit** | Pointer to **int64** | This field is unused and will be removed in a future API version.  | [optional] 
**Restrictions** | Pointer to [**AccountRestrictions**](AccountRestrictions.md) |  | [optional] 
**Security** | Pointer to [**Security**](Security.md) |  | [optional] 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 
**SpendingLimits** | Pointer to [**SpendingLimits**](SpendingLimits.md) |  | [optional] 
**Status** | Pointer to [**AccountStatus**](AccountStatus.md) |  | [optional] 
**StopPayments** | Pointer to [**[]StopPayment**](StopPayment.md) |  | [optional] [readonly] 
**SwiftCode** | Pointer to **string** | SWIFT code | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 

## Methods

### NewAccountGenericResponse

`func NewAccountGenericResponse() *AccountGenericResponse`

NewAccountGenericResponse instantiates a new AccountGenericResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGenericResponseWithDefaults

`func NewAccountGenericResponseWithDefaults() *AccountGenericResponse`

NewAccountGenericResponseWithDefaults instantiates a new AccountGenericResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessStatus

`func (o *AccountGenericResponse) GetAccessStatus() AccountAccessStatus`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *AccountGenericResponse) GetAccessStatusOk() (*AccountAccessStatus, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *AccountGenericResponse) SetAccessStatus(v AccountAccessStatus)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *AccountGenericResponse) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetAccessStatusLastUpdatedTime

`func (o *AccountGenericResponse) GetAccessStatusLastUpdatedTime() time.Time`

GetAccessStatusLastUpdatedTime returns the AccessStatusLastUpdatedTime field if non-nil, zero value otherwise.

### GetAccessStatusLastUpdatedTimeOk

`func (o *AccountGenericResponse) GetAccessStatusLastUpdatedTimeOk() (*time.Time, bool)`

GetAccessStatusLastUpdatedTimeOk returns a tuple with the AccessStatusLastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatusLastUpdatedTime

`func (o *AccountGenericResponse) SetAccessStatusLastUpdatedTime(v time.Time)`

SetAccessStatusLastUpdatedTime sets AccessStatusLastUpdatedTime field to given value.

### HasAccessStatusLastUpdatedTime

`func (o *AccountGenericResponse) HasAccessStatusLastUpdatedTime() bool`

HasAccessStatusLastUpdatedTime returns a boolean if a field has been set.

### GetAccountClosure

`func (o *AccountGenericResponse) GetAccountClosure() AccountClosure`

GetAccountClosure returns the AccountClosure field if non-nil, zero value otherwise.

### GetAccountClosureOk

`func (o *AccountGenericResponse) GetAccountClosureOk() (*AccountClosure, bool)`

GetAccountClosureOk returns a tuple with the AccountClosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountClosure

`func (o *AccountGenericResponse) SetAccountClosure(v AccountClosure)`

SetAccountClosure sets AccountClosure field to given value.

### HasAccountClosure

`func (o *AccountGenericResponse) HasAccountClosure() bool`

HasAccountClosure returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AccountGenericResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountGenericResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountGenericResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountGenericResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountNumberMasked

`func (o *AccountGenericResponse) GetAccountNumberMasked() string`

GetAccountNumberMasked returns the AccountNumberMasked field if non-nil, zero value otherwise.

### GetAccountNumberMaskedOk

`func (o *AccountGenericResponse) GetAccountNumberMaskedOk() (*string, bool)`

GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberMasked

`func (o *AccountGenericResponse) SetAccountNumberMasked(v string)`

SetAccountNumberMasked sets AccountNumberMasked field to given value.

### HasAccountNumberMasked

`func (o *AccountGenericResponse) HasAccountNumberMasked() bool`

HasAccountNumberMasked returns a boolean if a field has been set.

### GetAccountPurpose

`func (o *AccountGenericResponse) GetAccountPurpose() string`

GetAccountPurpose returns the AccountPurpose field if non-nil, zero value otherwise.

### GetAccountPurposeOk

`func (o *AccountGenericResponse) GetAccountPurposeOk() (*string, bool)`

GetAccountPurposeOk returns a tuple with the AccountPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPurpose

`func (o *AccountGenericResponse) SetAccountPurpose(v string)`

SetAccountPurpose sets AccountPurpose field to given value.

### HasAccountPurpose

`func (o *AccountGenericResponse) HasAccountPurpose() bool`

HasAccountPurpose returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountGenericResponse) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountGenericResponse) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountGenericResponse) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountGenericResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetApplicationId

`func (o *AccountGenericResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AccountGenericResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AccountGenericResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AccountGenericResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAutoPaymentPeriod

`func (o *AccountGenericResponse) GetAutoPaymentPeriod() int32`

GetAutoPaymentPeriod returns the AutoPaymentPeriod field if non-nil, zero value otherwise.

### GetAutoPaymentPeriodOk

`func (o *AccountGenericResponse) GetAutoPaymentPeriodOk() (*int32, bool)`

GetAutoPaymentPeriodOk returns a tuple with the AutoPaymentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPaymentPeriod

`func (o *AccountGenericResponse) SetAutoPaymentPeriod(v int32)`

SetAutoPaymentPeriod sets AutoPaymentPeriod field to given value.

### HasAutoPaymentPeriod

`func (o *AccountGenericResponse) HasAutoPaymentPeriod() bool`

HasAutoPaymentPeriod returns a boolean if a field has been set.

### GetBalanceCeiling

`func (o *AccountGenericResponse) GetBalanceCeiling() BalanceCeiling`

GetBalanceCeiling returns the BalanceCeiling field if non-nil, zero value otherwise.

### GetBalanceCeilingOk

`func (o *AccountGenericResponse) GetBalanceCeilingOk() (*BalanceCeiling, bool)`

GetBalanceCeilingOk returns a tuple with the BalanceCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCeiling

`func (o *AccountGenericResponse) SetBalanceCeiling(v BalanceCeiling)`

SetBalanceCeiling sets BalanceCeiling field to given value.

### HasBalanceCeiling

`func (o *AccountGenericResponse) HasBalanceCeiling() bool`

HasBalanceCeiling returns a boolean if a field has been set.

### GetBalanceFloor

`func (o *AccountGenericResponse) GetBalanceFloor() BalanceFloor`

GetBalanceFloor returns the BalanceFloor field if non-nil, zero value otherwise.

### GetBalanceFloorOk

`func (o *AccountGenericResponse) GetBalanceFloorOk() (*BalanceFloor, bool)`

GetBalanceFloorOk returns a tuple with the BalanceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloor

`func (o *AccountGenericResponse) SetBalanceFloor(v BalanceFloor)`

SetBalanceFloor sets BalanceFloor field to given value.

### HasBalanceFloor

`func (o *AccountGenericResponse) HasBalanceFloor() bool`

HasBalanceFloor returns a boolean if a field has been set.

### GetBalances

`func (o *AccountGenericResponse) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountGenericResponse) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountGenericResponse) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *AccountGenericResponse) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankAccountId

`func (o *AccountGenericResponse) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *AccountGenericResponse) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *AccountGenericResponse) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *AccountGenericResponse) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### GetBankRouting

`func (o *AccountGenericResponse) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *AccountGenericResponse) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *AccountGenericResponse) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *AccountGenericResponse) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *AccountGenericResponse) GetBillingPeriod() BillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *AccountGenericResponse) GetBillingPeriodOk() (*BillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *AccountGenericResponse) SetBillingPeriod(v BillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *AccountGenericResponse) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetBusinessIds

`func (o *AccountGenericResponse) GetBusinessIds() []string`

GetBusinessIds returns the BusinessIds field if non-nil, zero value otherwise.

### GetBusinessIdsOk

`func (o *AccountGenericResponse) GetBusinessIdsOk() (*[]string, bool)`

GetBusinessIdsOk returns a tuple with the BusinessIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessIds

`func (o *AccountGenericResponse) SetBusinessIds(v []string)`

SetBusinessIds sets BusinessIds field to given value.

### HasBusinessIds

`func (o *AccountGenericResponse) HasBusinessIds() bool`

HasBusinessIds returns a boolean if a field has been set.

### GetCloseDate

`func (o *AccountGenericResponse) GetCloseDate() string`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *AccountGenericResponse) GetCloseDateOk() (*string, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *AccountGenericResponse) SetCloseDate(v string)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *AccountGenericResponse) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCreationTime

`func (o *AccountGenericResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountGenericResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountGenericResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountGenericResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreditLimit

`func (o *AccountGenericResponse) GetCreditLimit() int64`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountGenericResponse) GetCreditLimitOk() (*int64, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountGenericResponse) SetCreditLimit(v int64)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *AccountGenericResponse) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountGenericResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountGenericResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountGenericResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountGenericResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerIds

`func (o *AccountGenericResponse) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *AccountGenericResponse) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *AccountGenericResponse) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *AccountGenericResponse) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetCustomerType

`func (o *AccountGenericResponse) GetCustomerType() CustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AccountGenericResponse) GetCustomerTypeOk() (*CustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AccountGenericResponse) SetCustomerType(v CustomerType)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *AccountGenericResponse) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetDaysPastDue

`func (o *AccountGenericResponse) GetDaysPastDue() int32`

GetDaysPastDue returns the DaysPastDue field if non-nil, zero value otherwise.

### GetDaysPastDueOk

`func (o *AccountGenericResponse) GetDaysPastDueOk() (*int32, bool)`

GetDaysPastDueOk returns a tuple with the DaysPastDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPastDue

`func (o *AccountGenericResponse) SetDaysPastDue(v int32)`

SetDaysPastDue sets DaysPastDue field to given value.

### HasDaysPastDue

`func (o *AccountGenericResponse) HasDaysPastDue() bool`

HasDaysPastDue returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *AccountGenericResponse) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *AccountGenericResponse) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *AccountGenericResponse) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *AccountGenericResponse) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetFeeProductIds

`func (o *AccountGenericResponse) GetFeeProductIds() []string`

GetFeeProductIds returns the FeeProductIds field if non-nil, zero value otherwise.

### GetFeeProductIdsOk

`func (o *AccountGenericResponse) GetFeeProductIdsOk() (*[]string, bool)`

GetFeeProductIdsOk returns a tuple with the FeeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeProductIds

`func (o *AccountGenericResponse) SetFeeProductIds(v []string)`

SetFeeProductIds sets FeeProductIds field to given value.

### HasFeeProductIds

`func (o *AccountGenericResponse) HasFeeProductIds() bool`

HasFeeProductIds returns a boolean if a field has been set.

### GetFundsOwnership

`func (o *AccountGenericResponse) GetFundsOwnership() FundsOwnership`

GetFundsOwnership returns the FundsOwnership field if non-nil, zero value otherwise.

### GetFundsOwnershipOk

`func (o *AccountGenericResponse) GetFundsOwnershipOk() (*FundsOwnership, bool)`

GetFundsOwnershipOk returns a tuple with the FundsOwnership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsOwnership

`func (o *AccountGenericResponse) SetFundsOwnership(v FundsOwnership)`

SetFundsOwnership sets FundsOwnership field to given value.

### HasFundsOwnership

`func (o *AccountGenericResponse) HasFundsOwnership() bool`

HasFundsOwnership returns a boolean if a field has been set.

### GetGeneralLedgerCategory

`func (o *AccountGenericResponse) GetGeneralLedgerCategory() GeneralLedgerCategory`

GetGeneralLedgerCategory returns the GeneralLedgerCategory field if non-nil, zero value otherwise.

### GetGeneralLedgerCategoryOk

`func (o *AccountGenericResponse) GetGeneralLedgerCategoryOk() (*GeneralLedgerCategory, bool)`

GetGeneralLedgerCategoryOk returns a tuple with the GeneralLedgerCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralLedgerCategory

`func (o *AccountGenericResponse) SetGeneralLedgerCategory(v GeneralLedgerCategory)`

SetGeneralLedgerCategory sets GeneralLedgerCategory field to given value.

### HasGeneralLedgerCategory

`func (o *AccountGenericResponse) HasGeneralLedgerCategory() bool`

HasGeneralLedgerCategory returns a boolean if a field has been set.

### GetGeneralLedgerType

`func (o *AccountGenericResponse) GetGeneralLedgerType() GeneralLedgerType`

GetGeneralLedgerType returns the GeneralLedgerType field if non-nil, zero value otherwise.

### GetGeneralLedgerTypeOk

`func (o *AccountGenericResponse) GetGeneralLedgerTypeOk() (*GeneralLedgerType, bool)`

GetGeneralLedgerTypeOk returns a tuple with the GeneralLedgerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralLedgerType

`func (o *AccountGenericResponse) SetGeneralLedgerType(v GeneralLedgerType)`

SetGeneralLedgerType sets GeneralLedgerType field to given value.

### HasGeneralLedgerType

`func (o *AccountGenericResponse) HasGeneralLedgerType() bool`

HasGeneralLedgerType returns a boolean if a field has been set.

### GetGracePeriod

`func (o *AccountGenericResponse) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *AccountGenericResponse) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *AccountGenericResponse) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *AccountGenericResponse) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetIban

`func (o *AccountGenericResponse) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountGenericResponse) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountGenericResponse) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountGenericResponse) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *AccountGenericResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountGenericResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountGenericResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountGenericResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestProductId

`func (o *AccountGenericResponse) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *AccountGenericResponse) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *AccountGenericResponse) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *AccountGenericResponse) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *AccountGenericResponse) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *AccountGenericResponse) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *AccountGenericResponse) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *AccountGenericResponse) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetIsAchEnabled

`func (o *AccountGenericResponse) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *AccountGenericResponse) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *AccountGenericResponse) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *AccountGenericResponse) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *AccountGenericResponse) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *AccountGenericResponse) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *AccountGenericResponse) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *AccountGenericResponse) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsEftCaEnabled

`func (o *AccountGenericResponse) GetIsEftCaEnabled() bool`

GetIsEftCaEnabled returns the IsEftCaEnabled field if non-nil, zero value otherwise.

### GetIsEftCaEnabledOk

`func (o *AccountGenericResponse) GetIsEftCaEnabledOk() (*bool, bool)`

GetIsEftCaEnabledOk returns a tuple with the IsEftCaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEftCaEnabled

`func (o *AccountGenericResponse) SetIsEftCaEnabled(v bool)`

SetIsEftCaEnabled sets IsEftCaEnabled field to given value.

### HasIsEftCaEnabled

`func (o *AccountGenericResponse) HasIsEftCaEnabled() bool`

HasIsEftCaEnabled returns a boolean if a field has been set.

### GetIsExternalCardEnabled

`func (o *AccountGenericResponse) GetIsExternalCardEnabled() bool`

GetIsExternalCardEnabled returns the IsExternalCardEnabled field if non-nil, zero value otherwise.

### GetIsExternalCardEnabledOk

`func (o *AccountGenericResponse) GetIsExternalCardEnabledOk() (*bool, bool)`

GetIsExternalCardEnabledOk returns a tuple with the IsExternalCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalCardEnabled

`func (o *AccountGenericResponse) SetIsExternalCardEnabled(v bool)`

SetIsExternalCardEnabled sets IsExternalCardEnabled field to given value.

### HasIsExternalCardEnabled

`func (o *AccountGenericResponse) HasIsExternalCardEnabled() bool`

HasIsExternalCardEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *AccountGenericResponse) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *AccountGenericResponse) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *AccountGenericResponse) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *AccountGenericResponse) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsSarEnabled

`func (o *AccountGenericResponse) GetIsSarEnabled() bool`

GetIsSarEnabled returns the IsSarEnabled field if non-nil, zero value otherwise.

### GetIsSarEnabledOk

`func (o *AccountGenericResponse) GetIsSarEnabledOk() (*bool, bool)`

GetIsSarEnabledOk returns a tuple with the IsSarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSarEnabled

`func (o *AccountGenericResponse) SetIsSarEnabled(v bool)`

SetIsSarEnabled sets IsSarEnabled field to given value.

### HasIsSarEnabled

`func (o *AccountGenericResponse) HasIsSarEnabled() bool`

HasIsSarEnabled returns a boolean if a field has been set.

### GetIsSecurity

`func (o *AccountGenericResponse) GetIsSecurity() bool`

GetIsSecurity returns the IsSecurity field if non-nil, zero value otherwise.

### GetIsSecurityOk

`func (o *AccountGenericResponse) GetIsSecurityOk() (*bool, bool)`

GetIsSecurityOk returns a tuple with the IsSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurity

`func (o *AccountGenericResponse) SetIsSecurity(v bool)`

SetIsSecurity sets IsSecurity field to given value.

### HasIsSecurity

`func (o *AccountGenericResponse) HasIsSecurity() bool`

HasIsSecurity returns a boolean if a field has been set.

### GetIsSyncteraPayEnabled

`func (o *AccountGenericResponse) GetIsSyncteraPayEnabled() bool`

GetIsSyncteraPayEnabled returns the IsSyncteraPayEnabled field if non-nil, zero value otherwise.

### GetIsSyncteraPayEnabledOk

`func (o *AccountGenericResponse) GetIsSyncteraPayEnabledOk() (*bool, bool)`

GetIsSyncteraPayEnabledOk returns a tuple with the IsSyncteraPayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncteraPayEnabled

`func (o *AccountGenericResponse) SetIsSyncteraPayEnabled(v bool)`

SetIsSyncteraPayEnabled sets IsSyncteraPayEnabled field to given value.

### HasIsSyncteraPayEnabled

`func (o *AccountGenericResponse) HasIsSyncteraPayEnabled() bool`

HasIsSyncteraPayEnabled returns a boolean if a field has been set.

### GetIsSystemAutoPayEnabled

`func (o *AccountGenericResponse) GetIsSystemAutoPayEnabled() bool`

GetIsSystemAutoPayEnabled returns the IsSystemAutoPayEnabled field if non-nil, zero value otherwise.

### GetIsSystemAutoPayEnabledOk

`func (o *AccountGenericResponse) GetIsSystemAutoPayEnabledOk() (*bool, bool)`

GetIsSystemAutoPayEnabledOk returns a tuple with the IsSystemAutoPayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAutoPayEnabled

`func (o *AccountGenericResponse) SetIsSystemAutoPayEnabled(v bool)`

SetIsSystemAutoPayEnabled sets IsSystemAutoPayEnabled field to given value.

### HasIsSystemAutoPayEnabled

`func (o *AccountGenericResponse) HasIsSystemAutoPayEnabled() bool`

HasIsSystemAutoPayEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *AccountGenericResponse) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *AccountGenericResponse) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *AccountGenericResponse) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *AccountGenericResponse) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AccountGenericResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AccountGenericResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AccountGenericResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AccountGenericResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountGenericResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountGenericResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountGenericResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountGenericResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *AccountGenericResponse) GetMinimumPayment() MinimumPayment`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *AccountGenericResponse) GetMinimumPaymentOk() (*MinimumPayment, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *AccountGenericResponse) SetMinimumPayment(v MinimumPayment)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *AccountGenericResponse) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.

### GetNickname

`func (o *AccountGenericResponse) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountGenericResponse) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountGenericResponse) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountGenericResponse) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetOpenDate

`func (o *AccountGenericResponse) GetOpenDate() string`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *AccountGenericResponse) GetOpenDateOk() (*string, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *AccountGenericResponse) SetOpenDate(v string)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *AccountGenericResponse) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetOverdraftLimit

`func (o *AccountGenericResponse) GetOverdraftLimit() int64`

GetOverdraftLimit returns the OverdraftLimit field if non-nil, zero value otherwise.

### GetOverdraftLimitOk

`func (o *AccountGenericResponse) GetOverdraftLimitOk() (*int64, bool)`

GetOverdraftLimitOk returns a tuple with the OverdraftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLimit

`func (o *AccountGenericResponse) SetOverdraftLimit(v int64)`

SetOverdraftLimit sets OverdraftLimit field to given value.

### HasOverdraftLimit

`func (o *AccountGenericResponse) HasOverdraftLimit() bool`

HasOverdraftLimit returns a boolean if a field has been set.

### GetRestrictions

`func (o *AccountGenericResponse) GetRestrictions() AccountRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *AccountGenericResponse) GetRestrictionsOk() (*AccountRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *AccountGenericResponse) SetRestrictions(v AccountRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *AccountGenericResponse) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetSecurity

`func (o *AccountGenericResponse) GetSecurity() Security`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *AccountGenericResponse) GetSecurityOk() (*Security, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *AccountGenericResponse) SetSecurity(v Security)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *AccountGenericResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSpendControlIds

`func (o *AccountGenericResponse) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *AccountGenericResponse) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *AccountGenericResponse) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *AccountGenericResponse) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.

### GetSpendingLimits

`func (o *AccountGenericResponse) GetSpendingLimits() SpendingLimits`

GetSpendingLimits returns the SpendingLimits field if non-nil, zero value otherwise.

### GetSpendingLimitsOk

`func (o *AccountGenericResponse) GetSpendingLimitsOk() (*SpendingLimits, bool)`

GetSpendingLimitsOk returns a tuple with the SpendingLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingLimits

`func (o *AccountGenericResponse) SetSpendingLimits(v SpendingLimits)`

SetSpendingLimits sets SpendingLimits field to given value.

### HasSpendingLimits

`func (o *AccountGenericResponse) HasSpendingLimits() bool`

HasSpendingLimits returns a boolean if a field has been set.

### GetStatus

`func (o *AccountGenericResponse) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountGenericResponse) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountGenericResponse) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountGenericResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPayments

`func (o *AccountGenericResponse) GetStopPayments() []StopPayment`

GetStopPayments returns the StopPayments field if non-nil, zero value otherwise.

### GetStopPaymentsOk

`func (o *AccountGenericResponse) GetStopPaymentsOk() (*[]StopPayment, bool)`

GetStopPaymentsOk returns a tuple with the StopPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPayments

`func (o *AccountGenericResponse) SetStopPayments(v []StopPayment)`

SetStopPayments sets StopPayments field to given value.

### HasStopPayments

`func (o *AccountGenericResponse) HasStopPayments() bool`

HasStopPayments returns a boolean if a field has been set.

### GetSwiftCode

`func (o *AccountGenericResponse) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountGenericResponse) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountGenericResponse) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *AccountGenericResponse) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetTenant

`func (o *AccountGenericResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountGenericResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountGenericResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AccountGenericResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


