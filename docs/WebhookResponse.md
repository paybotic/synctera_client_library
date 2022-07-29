# WebhookResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | indicates whether webhook is active | [optional] [default to null]
**Config** | [***WebhookConfig**](webhook_config.md) |  | [optional] [default to null]
**Events** | **[]string** | list of webhook events, use * to receive all notifications | [optional] [default to null]
**Id** | **string** | id of the webhook | [optional] [default to null]
**Name** | **string** | name of the webhook | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | The timestamp representing when the webhook request was made | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | The timestamp representing when the webhook was last modified | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

