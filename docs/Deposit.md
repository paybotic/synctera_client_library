# Deposit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account | [default to null]
**BackImageId** | **string** | ID of the uploaded image of the back of the check | [default to null]
**BusinessId** | **string** |  | [optional] [default to null]
**CheckAmount** | **int32** | Amount on check in ISO 4217 minor currency units | [default to null]
**CreationTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**DateCaptured** | [**time.Time**](time.Time.md) | Date the deposit was captured, in RFC 3339 format | [default to null]
**DateProcessed** | [**time.Time**](time.Time.md) | Date the deposit was processed, in RFC 3339 format | [default to null]
**DepositAmount** | **int32** | Amount deposited in ISO 4217 minor currency units | [default to null]
**DepositCurrency** | **string** | ISO 4217 currency code for the deposit amount | [default to null]
**FrontImageId** | **string** | ID of the uploaded image of the front of the check | [default to null]
**Id** | **string** | Remote Check Deposit ID | [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Metadata** | [***Metadata**](metadata.md) |  | [optional] [default to null]
**PersonId** | **string** |  | [optional] [default to null]
**Status** | **string** | The status of the deposit | [default to null]
**TransactionId** | **string** | The ID of the transaction associated with this deposit | [default to null]
**VendorInfo** | [***VendorInfo1**](vendor_info1.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

