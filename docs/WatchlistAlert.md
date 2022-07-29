# WatchlistAlert

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**time.Time**](time.Time.md) | When this alert was created | [optional] [default to null]
**Id** | **string** | Unique identifier for this alert | [optional] [default to null]
**ProviderInfo** | [***interface{}**](interface{}.md) | The information provided to Synctera that triggered this alert, as an arbitrary JSON object. Interpretation of this object is up to the client.  | [optional] [default to null]
**ProviderSubjectId** | **string** | The id of the provider subject for this alert | [optional] [default to null]
**ProviderSubscriptionId** | **string** | The id of the provider subscription for this alert | [optional] [default to null]
**ProviderWatchlistName** | **string** | The name of the provider for this alert | [optional] [default to null]
**Status** | **string** | The status of this alert | [default to null]
**Urls** | **[]string** | Where to get more information about this alert (according to our third-party data provider).  | [optional] [default to null]
**VendorInfo** | [***VendorInfo**](vendor_info.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

