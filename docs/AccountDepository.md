# AccountDepository

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessStatus** | [***AccountAccessStatus**](account_access_status.md) |  | [optional] [default to null]
**AccountNumber** | **string** | Account number | [optional] [default to null]
**AccountPurpose** | **string** | Purpose of the account | [optional] [default to null]
**AccountType** | [***AccountType**](account_type.md) |  | [optional] [default to null]
**Balances** | [**[]Balance**](balance.md) | A list of balances for account based on different type | [optional] [default to null]
**BankRouting** | **string** | Bank routing number | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | Account creation timestamp in RFC3339 format | [optional] [default to null]
**Currency** | **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] [default to null]
**CustomerIds** | **[]string** | A list of the customer IDs of the account holders. | [optional] [default to null]
**CustomerType** | [***CustomerType**](customer_type.md) |  | [optional] [default to null]
**ExchangeRateType** | **string** | Exchange rate type | [optional] [default to null]
**FeeProductIds** | **[]string** | A list of fee resources from account product that the current account associate with | [optional] [default to null]
**Iban** | **string** | International bank account number | [optional] [default to null]
**Id** | **string** | Account ID | [optional] [default to null]
**InterestProductId** | **string** | An interest from account product that the current account associate with | [optional] [default to null]
**IsAccountPool** | **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] [default to null]
**IsAchEnabled** | **bool** | A flag to indicate whether ACH transactions are enabled. | [optional] [default to null]
**IsCardEnabled** | **bool** | A flag to indicate whether card transactions are enabled. | [optional] [default to null]
**IsP2pEnabled** | **bool** | A flag to indicate whether P2P transactions are enabled. | [optional] [default to null]
**IsWireEnabled** | **bool** | A flag to indicate whether wire transactions are enabled. | [optional] [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) | Timestamp of the last account modification in RFC3339 format | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | User provided account metadata | [optional] [default to null]
**Nickname** | **string** | User provided account nickname | [optional] [default to null]
**Status** | [***Status**](status.md) |  | [optional] [default to null]
**SwiftCode** | **string** | SWIFT code | [optional] [default to null]
**BalanceCeiling** | [***BalanceCeiling**](balance_ceiling.md) |  | [optional] [default to null]
**BalanceFloor** | [***BalanceFloor**](balance_floor.md) |  | [optional] [default to null]
**OverdraftLimit** | **int64** | Account&#x27;s overdraft limit | [optional] [default to null]
**SpendingLimits** | [***SpendingLimits**](spending_limits.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

