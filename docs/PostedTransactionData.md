# PostedTransactionData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalData** | [***interface{}**](interface{}.md) | an unstructured json blob representing additional transaction information supplied by the integrator. | [default to null]
**HoldId** | **string** | The uuid of the hold (pending transaction) that this transaction originated from, if any. | [optional] [default to null]
**Lines** | [**[]TransactionLine1**](transaction_line1.md) | The set of accounting entries associated with this transaction. For example, a debit to a customer account will have a corresponding credit in a general ledger account. | [default to null]
**Memo** | **string** |  | [default to null]
**Metadata** | [***interface{}**](interface{}.md) |  | [default to null]
**OriginalTrx** | **string** | The \&quot;original\&quot; transaction that this transaction is related to. This is only populated in the case of reversed transactions. | [optional] [default to null]
**RiskInfo** | [***interface{}**](interface{}.md) | Information received by the transaction risk/fraud service related to this transaction | [default to null]
**UserData** | [***interface{}**](interface{}.md) | An unstructured JSON blob representing additional transaction information specific to each payment rail. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

