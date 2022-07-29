# ExternalAccountTransaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transaction amount. Number in cents. E.g. 1000 represents $10.00 | [optional] [default to null]
**AuthorizedDate** | **string** | Date that the transaction is authorized. ISO 8601 format ( YYYY-MM-DD ). | [optional] [default to null]
**Category** | **[]string** | Category of the transaction | [optional] [default to null]
**CheckNumber** | **string** | Check number of the transaction. This field will be null if not a check transaction. | [optional] [default to null]
**Currency** | **string** | ISO 4217 alphabetic currency code | [optional] [default to null]
**Date** | **string** | For pending transactions, this represents the date of the transaction occurred; for posted transactions, this represents the date of the transaction posted. ISO 8601 format ( YYYY-MM-DD ).  | [optional] [default to null]
**IsPending** | **bool** | Indicates the transaction is pending or unsettled if true. | [optional] [default to null]
**MerchantName** | **string** | Merchant name of the transaction | [optional] [default to null]
**PaymentChannel** | **string** | channel used to make a payment | [optional] [default to null]
**PaymentMethod** | **string** | Transfer type of the transaction, e.g. ACH | [optional] [default to null]
**TransactionId** | **string** | case-sensitive transaction ID | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

