# BinNetworkMapping

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | indicates whether mapping is active | [default to null]
**BankNetworkId** | **string** | ID debit network uses to identify a bank | [default to null]
**BinId** | **string** | The ID of the bank&#x27;s BIN that uses this debit network | [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | The timestamp representing when BIN network mapping was created | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The time when mapping becomes inactive | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | The timestamp representing when the BIN network mapping was last modified | [optional] [default to null]
**NetworkId** | **string** | The ID of the debit_network associated with the BIN of the bank | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The time when mapping becomes active | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

