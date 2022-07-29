# PatchAccountsRequestRoutingIdentifiers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchRoutingNumber** | **string** | The routing number used for US ACH payments. On write, Synctera will store the entire routing number; on read, we only return the last 4 characters.  | [optional] [default to null]
**BankName** | **string** | The name of the bank managing the account | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

