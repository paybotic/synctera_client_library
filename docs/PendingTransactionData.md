# PendingTransactionData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of the hold. | [default to null]
**AvailBalance** | **int32** | The account \&quot;available balance\&quot; at the time this hold was created | [default to null]
**Balance** | **int32** | The account balance at the time this hold was created | [default to null]
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | [default to null]
**DcSign** | [***DcSign**](dc_sign.md) |  | [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | The effective date of the transaction once it gets posted | [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | The date that at which this hold is no longer valid. | [default to null]
**ExternalData** | [***interface{}**](interface{}.md) | an unstructured json blob representing additional transaction information supplied by the integrator. | [default to null]
**ForcePost** | **bool** | Whether or not the hold was forced (spending controls ignored) | [default to null]
**History** | [**[]PendingTransactionHistory**](pending_transaction_history.md) | An array representing any previous states of the hold, if it has been modified (For example, increasing or decreasing the hold amount). | [default to null]
**Memo** | **string** |  | [default to null]
**Network** | **string** | The network this transaction is associated with | [default to null]
**Operation** | **string** |  | [default to null]
**Reason** | **string** | If a hold has been declined or modified, this will include the reason. | [default to null]
**ReqAmount** | **int32** | The requested amount, in the case of hold modifications. | [default to null]
**RiskInfo** | [***interface{}**](interface{}.md) | Information received by the transaction risk/fraud service related to this transaction | [default to null]
**Status** | **string** | The status of the hold. | [default to null]
**Subtype** | **string** | The specific transaction type. For example, for &#x60;ach&#x60;, this may be \&quot;outgoing_debit\&quot;. | [default to null]
**TotalAmount** | **int32** | The total amount of the hold. This may be different than &#x60;amount&#x60; in the case where a hold increase or decrease was requested. | [default to null]
**TransactionId** | **string** | The uuid of the transaction that this pending transaction originated from, if any. This is primary used when a transaction \&quot;posts\&quot;, but a subset of the amount reserved until a future settlement date. | [optional] [default to null]
**Type_** | **string** | The general type of transaction. For example, \&quot;card\&quot; or \&quot;ach\&quot;. | [default to null]
**UserData** | [***interface{}**](interface{}.md) | An unstructured JSON blob representing additional transaction information specific to each payment rail. | [default to null]
**WasPartial** | **bool** | Does this hold represent a partial debit (or credit)? | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

