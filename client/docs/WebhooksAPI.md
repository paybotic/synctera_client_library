# \WebhooksAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecret**](WebhooksAPI.md#CreateSecret) | **Post** /webhook_secrets | Create a secret
[**CreateWebhook**](WebhooksAPI.md#CreateWebhook) | **Post** /webhooks | Create a webhook
[**DeleteWebhook**](WebhooksAPI.md#DeleteWebhook) | **Delete** /webhooks/{webhook_id} | Delete a webhook
[**GetEvent**](WebhooksAPI.md#GetEvent) | **Get** /webhooks/{webhook_id}/events/{event_id} | Get webhook event
[**GetWebhook**](WebhooksAPI.md#GetWebhook) | **Get** /webhooks/{webhook_id} | Get a webhook
[**ListEvents**](WebhooksAPI.md#ListEvents) | **Get** /webhooks/{webhook_id}/events | List webhook events
[**ListWebhooks**](WebhooksAPI.md#ListWebhooks) | **Get** /webhooks | List webhooks
[**ReplaceSecret**](WebhooksAPI.md#ReplaceSecret) | **Put** /webhook_secrets | Replace an existing secret
[**ResendEvent**](WebhooksAPI.md#ResendEvent) | **Post** /webhooks/{webhook_id}/events/{event_id}/resend | Resend an event
[**RevokeSecret**](WebhooksAPI.md#RevokeSecret) | **Delete** /webhook_secrets | Revoke the secret
[**TriggerEvent**](WebhooksAPI.md#TriggerEvent) | **Post** /webhooks/trigger | Trigger an event
[**UpdateWebhook**](WebhooksAPI.md#UpdateWebhook) | **Put** /webhooks/{webhook_id} | Update a webhook



## CreateSecret

> CreateSecret201Response CreateSecret(ctx).Execute()

Create a secret



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.CreateSecret(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecret`: CreateSecret201Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CreateSecret`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretRequest struct via the builder pattern


### Return type

[**CreateSecret201Response**](CreateSecret201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> Webhook CreateWebhook(ctx).Webhook(webhook).Execute()

Create a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	webhook := *openapiclient.NewWebhook([]openapiclient.EventType{openapiclient.event_type{EventTypeExplicit: openapiclient.event_type_explicit("ACCOUNT.CREATED")}}, false, "Url_example") // Webhook | Webhook to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.CreateWebhook(context.Background()).Webhook(webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhook** | [**Webhook**](Webhook.md) | Webhook to create | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteResponse DeleteWebhook(ctx, webhookId).Execute()

Delete a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	webhookId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.DeleteWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> Event GetEvent(ctx, eventId, webhookId).Execute()

Get webhook event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	eventId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Webhook event ID
	webhookId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetEvent(context.Background(), eventId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Webhook event ID | 
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, webhookId).Execute()

Get a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	webhookId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> EventList ListEvents(ctx, webhookId).StartTime(startTime).EndDate(endDate).EndTime(endTime).PageToken(pageToken).ResourceId(resourceId).Limit(limit).StartDate(startDate).Execute()

List webhook events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	webhookId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Webhook ID
	startTime := time.Now() // time.Time | Start time of date-time range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00. (optional)
	endDate := time.Now() // string | End date of date range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00.. end_date is alias of end_time and is deprecated. Please use end_time instead. (optional)
	endTime := time.Now() // time.Time | End time of date-time range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00. (optional)
	pageToken := "a8937a0d" // string |  (optional)
	resourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit returned events to those that occurred on the specified resource. (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	startDate := time.Now() // string | Start date of date range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00.. start_date is alias of start_time and is deprecated. Please use start_time instead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ListEvents(context.Background(), webhookId).StartTime(startTime).EndDate(endDate).EndTime(endTime).PageToken(pageToken).ResourceId(resourceId).Limit(limit).StartDate(startDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvents`: EventList
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **time.Time** | Start time of date-time range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00. | 
 **endDate** | **string** | End date of date range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00.. end_date is alias of end_time and is deprecated. Please use end_time instead. | 
 **endTime** | **time.Time** | End time of date-time range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00. | 
 **pageToken** | **string** |  | 
 **resourceId** | **string** | Limit returned events to those that occurred on the specified resource. | 
 **limit** | **int32** |  | [default to 100]
 **startDate** | **string** | Start date of date range filtering for events. Date is inclusive and should be in UTC timezone 00:00:00.. start_date is alias of start_time and is deprecated. Please use start_time instead. | 

### Return type

[**EventList**](EventList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooks

> WebhookList ListWebhooks(ctx).PageToken(pageToken).IsEnabledOnly(isEnabledOnly).Limit(limit).Execute()

List webhooks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	pageToken := "a8937a0d" // string |  (optional)
	isEnabledOnly := true // bool |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ListWebhooks(context.Background()).PageToken(pageToken).IsEnabledOnly(isEnabledOnly).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhooks`: WebhookList
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageToken** | **string** |  | 
 **isEnabledOnly** | **bool** |  | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**WebhookList**](WebhookList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSecret

> ReplaceSecret200Response ReplaceSecret(ctx).ReplaceSecretRequest(replaceSecretRequest).Execute()

Replace an existing secret



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	replaceSecretRequest := *openapiclient.NewReplaceSecretRequest() // ReplaceSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ReplaceSecret(context.Background()).ReplaceSecretRequest(replaceSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ReplaceSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceSecret`: ReplaceSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ReplaceSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replaceSecretRequest** | [**ReplaceSecretRequest**](ReplaceSecretRequest.md) |  | 

### Return type

[**ReplaceSecret200Response**](ReplaceSecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEvent

> Event ResendEvent(ctx, eventId, webhookId).Delay(delay).Execute()

Resend an event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	eventId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Webhook event ID
	webhookId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Webhook ID
	delay := int32(56) // int32 | Delay the event triggering in seconds. Events are checked once a minute, so a short delay may not result in an immediate resend. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ResendEvent(context.Background(), eventId, webhookId).Delay(delay).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ResendEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ResendEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Webhook event ID | 
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **delay** | **int32** | Delay the event triggering in seconds. Events are checked once a minute, so a short delay may not result in an immediate resend. | 

### Return type

[**Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeSecret

> RevokeSecret(ctx).OldSecretOnly(oldSecretOnly).Execute()

Revoke the secret



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	oldSecretOnly := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.RevokeSecret(context.Background()).OldSecretOnly(oldSecretOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.RevokeSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oldSecretOnly** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerEvent

> EventTrigger TriggerEvent(ctx).TriggerEventRequest(triggerEventRequest).Execute()

Trigger an event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	triggerEventRequest := *openapiclient.NewTriggerEventRequest() // TriggerEventRequest | Provide an event type to trigger

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.TriggerEvent(context.Background()).TriggerEventRequest(triggerEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.TriggerEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerEvent`: EventTrigger
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.TriggerEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggerEventRequest** | [**TriggerEventRequest**](TriggerEventRequest.md) | Provide an event type to trigger | 

### Return type

[**EventTrigger**](EventTrigger.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> Webhook UpdateWebhook(ctx, webhookId).Webhook(webhook).Execute()

Update a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	webhookId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Webhook ID
	webhook := *openapiclient.NewWebhook([]openapiclient.EventType{openapiclient.event_type{EventTypeExplicit: openapiclient.event_type_explicit("ACCOUNT.CREATED")}}, false, "Url_example") // Webhook | Webhook to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.UpdateWebhook(context.Background(), webhookId).Webhook(webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhook** | [**Webhook**](Webhook.md) | Webhook to update | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

