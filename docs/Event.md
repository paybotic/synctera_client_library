# Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventResource** | **string** | Json string of object associated with the event. For example, if your event is ACCOUNT.CREATED, You can refer to Acccount to parse the account event to obtain the ID, status etc.  | [optional] [default to null]
**EventResourceChangedFields** | **string** | Json string of object associated with the event related to a resource change. This only contains those fields that have value changed on the event, and the field values are prior to the resource change event.  | [optional] [default to null]
**EventTime** | [**time.Time**](time.Time.md) | Timestamp of the current event raised | [optional] [default to null]
**Id** | **string** | Unique event ID of the webhook request. Use event endpoints to get more event summary data | [optional] [default to null]
**Metadata** | **string** | Metadata that stored in the webhook subscription | [optional] [default to null]
**ResponseHistory** | [**[]ResponseHistoryItem**](response_history_item.md) | Response history of the webhook request | [optional] [default to null]
**Status** | **string** | Current event status. Failing event will keep retry until it is purged. | [optional] [default to null]
**Type_** | [***EventTypeExplicit**](event_type_explicit.md) |  | [optional] [default to null]
**Url** | **string** | URL that the current event will be sent to | [optional] [default to null]
**WebhookId** | **string** | Webhook the current event belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

