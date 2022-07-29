# PendingTransaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account id associated with the hold | [default to null]
**AccountNo** | **string** | The account number associated with the hold | [default to null]
**Created** | [**time.Time**](time.Time.md) | The creation date of the hold | [default to null]
**Data** | [***PendingTransactionData**](pending_transaction_data.md) |  | [default to null]
**Id** | **int32** |  | [default to null]
**Idemkey** | **string** | The idempotency key used when initially creating this hold. | [default to null]
**ReferenceId** | **string** | An external ID provided by the payment network to represent this transaction. | [default to null]
**Tenant** | **string** | The tenant associated with this hold, in the form \&quot;&lt;bankid&gt;_&lt;partnerid&gt;\&quot; | [default to null]
**Updated** | [**time.Time**](time.Time.md) | The date the hold was last update | [default to null]
**Uuid** | **string** | The unique identifier of the hold transaction. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

