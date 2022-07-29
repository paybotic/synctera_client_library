# Webhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description of what the webhook is used for | [optional] [default to null]
**EnabledEvents** | [**[]EventType1**](event_type1.md) | A list of the events that will trigger the webhook | [default to null]
**Id** | **string** | The unique ID of the webhook | [optional] [default to null]
**IsEnabled** | **bool** | Set the webhook to be enabled or disabled | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | Timestamp that this webhook was created or the last time any field was changed | [optional] [default to null]
**Metadata** | **string** | Additional information stored to the webhook | [optional] [default to null]
**Url** | **string** | URL that the webhook will send request to | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

