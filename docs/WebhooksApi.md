# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecret**](WebhooksApi.md#CreateSecret) | **Post** /webhooks/secret | Create a secret
[**CreateWebhook1**](WebhooksApi.md#CreateWebhook1) | **Post** /webhooks | Create a webhook
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /webhooks/{webhook_id} | Delete a webhook
[**GetEvent**](WebhooksApi.md#GetEvent) | **Get** /webhooks/{webhook_id}/events/{event_id} | Get webhook event
[**GetWebhook1**](WebhooksApi.md#GetWebhook1) | **Get** /webhooks/{webhook_id} | Get a webhook
[**ListEvents**](WebhooksApi.md#ListEvents) | **Get** /webhooks/{webhook_id}/events | List webhook events
[**ListWebhooks1**](WebhooksApi.md#ListWebhooks1) | **Get** /webhooks | List webhooks
[**ReplaceSecret**](WebhooksApi.md#ReplaceSecret) | **Put** /webhooks/secret | Replace an existing secret
[**ResendEvent**](WebhooksApi.md#ResendEvent) | **Post** /webhooks/{webhook_id}/events/{event_id}/resend | Resend an event
[**RevokeSecret**](WebhooksApi.md#RevokeSecret) | **Delete** /webhooks/secret | Revoke the secret
[**TriggerEvent**](WebhooksApi.md#TriggerEvent) | **Post** /webhooks/trigger | Trigger an event
[**UpdateWebhook**](WebhooksApi.md#UpdateWebhook) | **Put** /webhooks/{webhook_id} | Update a webhook

# **CreateSecret**
> InlineResponse201 CreateSecret(ctx, )
Create a secret

Create a webhook secret. The secret will be used to verify all subsequent webhook request signature.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWebhook1**
> Webhook CreateWebhook1(ctx, body)
Create a webhook

Create a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Webhook**](Webhook.md)| Webhook to create | 

### Return type

[**Webhook**](webhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> DeleteResponse DeleteWebhook(ctx, webhookId)
Delete a webhook

Delete a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | [**string**](.md)| Webhook ID | 

### Return type

[**DeleteResponse**](delete_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvent**
> Event GetEvent(ctx, webhookId, eventId)
Get webhook event

Get webhook event by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | [**string**](.md)| Webhook ID | 
  **eventId** | [**string**](.md)| Webhook event ID | 

### Return type

[**Event**](event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook1**
> Webhook GetWebhook1(ctx, webhookId)
Get a webhook

Get a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | [**string**](.md)| Webhook ID | 

### Return type

[**Webhook**](webhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEvents**
> EventList ListEvents(ctx, webhookId, optional)
List webhook events

List webhook events. This response will not associate with the event response history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | [**string**](.md)| Webhook ID | 
 **optional** | ***WebhooksApiListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiListEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.String**| Start date of date range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00.. start_date is alias of start_time and is deprecated. Please use start_time instead. | 
 **endDate** | **optional.String**| End date of date range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00.. end_date is alias of end_time and is deprecated. Please use end_time instead. | 
 **startTime** | **optional.Time**| Start time of date-time range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00. | 
 **endTime** | **optional.Time**| End time of date-time range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00. | 
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 

### Return type

[**EventList**](event_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhooks1**
> WebhookList ListWebhooks1(ctx, optional)
List webhooks

List all webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiListWebhooks1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiListWebhooks1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 
 **isEnabledOnly** | **optional.Bool**|  | 

### Return type

[**WebhookList**](webhook_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSecret**
> InlineResponse2001 ReplaceSecret(ctx, body)
Replace an existing secret

Replace an existing webhook secret immediately or as part of rotation. This new secret will be used to verify all subsequent webhook request signature.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksSecretBody**](WebhooksSecretBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResendEvent**
> Event ResendEvent(ctx, webhookId, eventId, optional)
Resend an event

Resend a webhook event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | [**string**](.md)| Webhook ID | 
  **eventId** | [**string**](.md)| Webhook event ID | 
 **optional** | ***WebhooksApiResendEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiResendEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **delay** | **optional.Int32**| Delay the event triggering in seconds | 

### Return type

[**Event**](event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeSecret**
> RevokeSecret(ctx, optional)
Revoke the secret

Revoke the existing webhook secret. If this is called at the rolling secret time, then both old and new secrets will be revoked

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiRevokeSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiRevokeSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oldSecretOnly** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerEvent**
> EventTrigger TriggerEvent(ctx, body)
Trigger an event

Trigger an specific event for webhook testing purpose

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksTriggerBody**](WebhooksTriggerBody.md)| Provide an event type to trigger | 

### Return type

[**EventTrigger**](event_trigger.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> Webhook UpdateWebhook(ctx, body, webhookId)
Update a webhook

Update a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Webhook**](Webhook.md)| Webhook to update | 
  **webhookId** | [**string**](.md)| Webhook ID | 

### Return type

[**Webhook**](webhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

