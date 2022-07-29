# TransactionData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalData** | [***interface{}**](interface{}.md) | an unstructured json blob representing additional transaction information supplied by the integrator. | [optional] [default to null]
**Lines** | [**[]TransactionLine**](transaction_line.md) | The set of accounting entries associated with this transaction. For example, a debit to a customer account will have a corresponding credit in a general ledger account. | [default to null]
**Memo** | **string** |  | [default to null]
**Metadata** | [***interface{}**](interface{}.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

