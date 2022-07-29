# PostedTransaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**time.Time**](time.Time.md) | The creation date of the transaction | [default to null]
**Data** | [***PostedTransactionData**](posted_transaction_data.md) |  | [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | The \&quot;effective date\&quot; of a transaction. This may be earlier than posted_date in some cases (for example, a transaction that occurs on a Saturday may not be posted until the following Monday, but would have an effective date of Saturday) | [default to null]
**Id** | **int32** |  | [default to null]
**Idemkey** | **string** | The idempotency key used when initially creating this transaction. | [default to null]
**InfoOnly** | **bool** | Whether or not this transaction represents a purely informational operation or an actual money movement | [default to null]
**LeadMode** | **bool** | Whether or not this transaction was created operating in \&quot;lead ledger\&quot; mode | [default to null]
**PostedDate** | [**time.Time**](time.Time.md) | The date the transaction was posted. This is the date any money is considered to be added or removed from an account. | [default to null]
**ReferenceId** | **string** | An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers. | [default to null]
**Status** | **string** |  | [default to null]
**Subtype** | **string** | The specific transaction type. For example, for &#x60;ach&#x60;, this may be \&quot;outgoing_debit\&quot;. | [default to null]
**Tenant** | **string** | The tenant associated with this transaction, in the form \&quot;&lt;bankid&gt;_&lt;partnerid&gt;\&quot; | [default to null]
**Type_** | **string** | The general type of transaction. For example, \&quot;card\&quot; or \&quot;ach\&quot;. | [default to null]
**Updated** | [**time.Time**](time.Time.md) | The date the transaction was last updated | [default to null]
**Uuid** | **string** | The unique identifier of the transaction. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

