# TemplateFieldsLineOfCredit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [***AccountType**](account_type.md) |  | [default to null]
**BankCountry** | **string** | Bank country of the account. ISO 3166-1 Alpha-2 or Alpha-3 country code. | [default to null]
**Currency** | **string** | Account currency. ISO 4217 alphabetic currency code | [default to null]
**ChargeoffPeriod** | **int32** | The number of days an account can stay delinquent before marking an account as charged-off.  | [optional] [default to 90]
**DelinquencyPeriod** | **int32** | The number of days past the due date to wait for a minimum payment before marking an account as delinquent.  | [optional] [default to 30]
**GracePeriod** | **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment.  | [optional] [default to 30]
**MinimumPayment** | [***MinimumPayment**](minimum_payment.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

