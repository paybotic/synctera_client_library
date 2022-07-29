# AccountSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Account number | [optional] [default to null]
**AccountStatus** | **string** | Account Status | [optional] [default to null]
**AccountType** | **string** | The type of the account. In lead mode, this always takes the value of the template. If not specified in shadow mode, CHECKING will be assumed.  | [optional] [default to null]
**BalanceCeiling** | [***AccountSummaryBalanceCeiling**](account_summary_balance_ceiling.md) |  | [optional] [default to null]
**BalanceFloor** | [***AccountSummaryBalanceFloor**](account_summary_balance_floor.md) |  | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | Account creation time | [optional] [default to null]
**Currency** | **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] [default to null]
**CustomerType** | **string** | Customer type | [optional] [default to null]
**FinancialInstitution** | [***FinancialInstitution**](financial_institution.md) |  | [optional] [default to null]
**Id** | **string** | The unique identifier of the account the statement belongs to | [optional] [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) | Account last modification time | [optional] [default to null]
**Nickname** | **string** | User provided account nickname | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

