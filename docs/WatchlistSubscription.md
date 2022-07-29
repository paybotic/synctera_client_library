# WatchlistSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRenew** | **bool** | Whether this subscription should automatically renew when the subscription period is over (default: vendor-dependent).  | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) | When this subscription was created | [optional] [default to null]
**CustomerConsent** | **bool** | Whether this customer has consented to being enrolled for watchlist monitoring  | [default to null]
**Id** | **string** | Unique identifier for this subscription | [optional] [default to null]
**PeriodEnd** | **string** | The date when monitoring of this individual should end. | [optional] [default to null]
**PeriodStart** | **string** | The date when monitoring of this individual should begin (default: today). | [optional] [default to null]
**ProviderSubscriptionId** | **string** | External provider subscription id | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

