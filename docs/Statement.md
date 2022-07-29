# Statement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The unique identifier of the account the statement belongs to | [optional] [default to null]
**DueDate** | **string** | The limit date when the due amount indicated on the statement should be paid | [optional] [default to null]
**EndDate** | **string** | The date indicating the ending of the time interval covered by the statement | [optional] [default to null]
**Id** | **string** | statement ID | [optional] [default to null]
**IssueDate** | **string** | The date when the statement has been issued | [optional] [default to null]
**StartDate** | **string** | The date indicating the begining of the time interval covered by the statement | [optional] [default to null]
**AccountSummary** | [***AccountSummary**](account_summary.md) |  | [optional] [default to null]
**AuthorizedSigner** | [**[]Person1**](person1.md) |  | [optional] [default to null]
**Disclosure** | **string** |  | [optional] [default to null]
**JointAccountHolders** | [**[]Person1**](person1.md) |  | [optional] [default to null]
**PrimaryAccountHolderBusiness** | [***Business1**](business1.md) |  | [optional] [default to null]
**PrimaryAccountHolderPersonal** | [***Person1**](person1.md) |  | [optional] [default to null]
**SavingsSummary** | [***SavingsSummary**](savings_summary.md) |  | [optional] [default to null]
**Transactions** | [**[]Transaction**](transaction.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

