# TemplateFieldsGenericResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [***AccountType**](account_type.md) |  | [default to null]
**BalanceCeiling** | [***BalanceCeiling**](balance_ceiling.md) |  | [optional] [default to null]
**BalanceFloor** | [***BalanceFloor**](balance_floor.md) |  | [optional] [default to null]
**BankCountry** | **string** | Bank country of the account | [default to null]
**BillingPeriod** | [***BillingPeriod**](billing_period.md) |  | [optional] [default to null]
**ChargeoffPeriod** | **int32** | The number of days an account can stay delinquent before marking an account as charged-off.  | [optional] [default to 90]
**Currency** | **string** | Account currency. ISO 4217 alphabetic currency code | [default to null]
**DelinquencyPeriod** | **int32** | The number of days past the due date to wait for a minimum payment before marking an account as delinquent.  | [optional] [default to 30]
**FeeProductIds** | **[]string** | A list of fee resources from account product that new accounts will associate with | [optional] [default to null]
**GracePeriod** | **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment.  | [optional] [default to 30]
**InterestProductId** | **string** | Interest from account product that new accounts will associate with | [optional] [default to null]
**IsAchEnabled** | **bool** | Enable ACH transaction on ledger. | [optional] [default to false]
**IsCardEnabled** | **bool** | Enable card transaction on ledger. | [optional] [default to false]
**IsP2pEnabled** | **bool** | Enable P2P transaction on ledger. | [optional] [default to false]
**IsWireEnabled** | **bool** | Enable wire transaction on ledger. | [optional] [default to false]
**MinimumPayment** | [***MinimumPayment**](minimum_payment.md) |  | [optional] [default to null]
**OverdraftLimit** | **int64** | Account&#x27;s overdraft limit. Default is 0. Unit in cents. | [optional] [default to null]
**SpendingLimits** | [***SpendingLimits**](spending_limits.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

