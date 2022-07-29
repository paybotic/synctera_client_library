# Wire

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Transfer amount in cents ($100 would be 10000) | [default to null]
**BankMessage** | **string** | Instructions intended for the financial institutions that are processing the wire. | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) |  | [default to null]
**Currency** | **string** | 3-character currency code | [default to null]
**CustomerId** | **string** | The customer UUID representing the person initiating the Wire transfer | [default to null]
**Id** | **string** | wire ID | [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) |  | [default to null]
**OriginatingAccountId** | **string** | Sender account ID | [default to null]
**ReceivingAccountId** | **string** | The external account uuid representing the recipient of the wire. | [default to null]
**RecipientMessage** | **string** | Information from the originator to the beneficiary (recipient). | [optional] [default to null]
**SenderReferenceId** | **string** | Sender&#x27;s id associated with fedwire transfer | [default to null]
**Status** | **string** | The current status of the transfer | [default to null]
**TransactionId** | **string** | ID of the resulting transaction resource | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

