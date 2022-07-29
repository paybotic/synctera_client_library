# TemplateFieldsDepository

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [***AccountType**](account_type.md) |  | [default to null]
**BankCountry** | **string** | Bank country of the account. ISO 3166-1 Alpha-2 or Alpha-3 country code. | [default to null]
**Currency** | **string** | Account currency. ISO 4217 alphabetic currency code | [default to null]
**BalanceCeiling** | [***BalanceCeiling**](balance_ceiling.md) |  | [optional] [default to null]
**BalanceFloor** | [***BalanceFloor**](balance_floor.md) |  | [optional] [default to null]
**FeeProductIds** | **[]string** | A list of fee resources from account product that new accounts will associate with | [optional] [default to null]
**InterestProductId** | **string** | Interest from account product that new accounts will associate with | [optional] [default to null]
**IsAchEnabled** | **bool** | Enable ACH transaction. | [optional] [default to false]
**IsCardEnabled** | **bool** | Enable card transaction. | [optional] [default to false]
**IsP2pEnabled** | **bool** | Enable P2P transaction. | [optional] [default to false]
**IsWireEnabled** | **bool** | Enable wire transaction. | [optional] [default to false]
**OverdraftLimit** | **int64** | Account&#x27;s overdraft limit. Default is 0. Unit in cents. | [optional] [default to null]
**SpendingLimits** | [***SpendingLimits**](spending_limits.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

