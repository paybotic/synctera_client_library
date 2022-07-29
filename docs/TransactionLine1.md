# TransactionLine1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account uuid associated with this transaction line | [default to null]
**AccountNo** | **string** | The account number associated with this transaction line | [default to null]
**Amount** | **int32** | The amount (in cents) of the transaction | [default to null]
**AvailableBalance** | **int32** | The account \&quot;available balance\&quot; at the point in time this transaction was posted | [default to null]
**Balance** | **int32** | The account balance at the point in time this transaction was posted | [default to null]
**Created** | [**time.Time**](time.Time.md) | The creation date of the transaction | [default to null]
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | [default to null]
**DcSign** | [***DcSign**](dc_sign.md) |  | [default to null]
**IsFee** | **bool** | Whether or not this line is considered a fee | [default to null]
**IsGlAcc** | **bool** | Whether or not this line represents a GL account | [default to null]
**IsOffset** | **bool** | Whether or not this line is considered the \&quot;offset\&quot; line | [default to null]
**IsPrimary** | **bool** | Whether or not this line is considered the \&quot;primary\&quot; line | [default to null]
**Meta** | [***interface{}**](interface{}.md) |  | [default to null]
**Network** | **string** | The network this transaction is associated with | [default to null]
**RelatedLine** | **int32** |  | [default to null]
**Seq** | **int32** |  | [default to null]
**Tenant** | **string** | The tenant associated with this transaction, in the form \&quot;&lt;bankid&gt;_&lt;partnerid&gt;\&quot; | [default to null]
**Updated** | [**time.Time**](time.Time.md) | The creation date of the transaction | [default to null]
**Uuid** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

