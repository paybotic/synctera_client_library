# PendingTransactionHistory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | **string** | The account number associated with the hold | [default to null]
**Created** | [**time.Time**](time.Time.md) | The creation date of the hold | [default to null]
**Data** | [***PendingTransactionHistoryData**](pending_transaction_history_data.md) |  | [default to null]
**Id** | **int32** |  | [default to null]
**Idemkey** | **string** | The idempotency key used when initially creating this transaction. | [default to null]
**Tenant** | **string** | The tenant associated with this transaction, in the form \&quot;&lt;bankid&gt;_&lt;partnerid&gt;\&quot; | [default to null]
**Updated** | [**time.Time**](time.Time.md) | The date the hold was last update | [default to null]
**Uuid** | **string** | The unique identifier of the hold. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

