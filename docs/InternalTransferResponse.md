# InternalTransferResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount (in cents) to transfer from originating account to receiving account. | [default to null]
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | [default to null]
**Memo** | **string** | A short note to the recipient | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | Arbitrary key-value metadata to associate with the transaction | [optional] [default to null]
**OriginatingAccountAlias** | **string** | An alias representing a GL account to debit. This is alternative to specifying by account id | [optional] [default to null]
**OriginatingAccountCustomerId** | **string** | The customer id of the owner of the originating account. | [optional] [default to null]
**OriginatingAccountId** | **string** | The UUID of the account being debited | [optional] [default to null]
**ReceivingAccountAlias** | **string** | An alias representing a GL account to credit. This is an alternative to specifying by account id | [optional] [default to null]
**ReceivingAccountCustomerId** | **string** | The customer id of the owner of the receiving account. Only required when type is \&quot;outgoing_remittance\&quot; | [optional] [default to null]
**ReceivingAccountId** | **string** | The UUID of the account being credited | [optional] [default to null]
**Type_** | **string** | The desired transaction type to use for this transfer | [default to null]
**Id** | **string** | The transaction id associated with the transfer | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

