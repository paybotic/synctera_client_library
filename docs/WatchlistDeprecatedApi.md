# \WatchlistDeprecatedAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListWatchlistSubscriptions**](WatchlistDeprecatedAPI.md#ListWatchlistSubscriptions) | **Get** /customers/{customer_id}/watchlists/subscriptions | List watchlist monitoring subscriptions for a customer
[**UpdateWatchlistAlert**](WatchlistDeprecatedAPI.md#UpdateWatchlistAlert) | **Put** /customers/{customer_id}/watchlists/alerts/{alert_id} | Update watchlist alert
[**UpdateWatchlistSubscription**](WatchlistDeprecatedAPI.md#UpdateWatchlistSubscription) | **Put** /customers/{customer_id}/watchlists/subscriptions/{subscription_id} | Update watchlist monitoring subscription



## ListWatchlistSubscriptions

> WatchlistSubscriptionList ListWatchlistSubscriptions(ctx, customerId).Execute()

List watchlist monitoring subscriptions for a customer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "./openapi"
)

func main() {
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchlistDeprecatedAPI.ListWatchlistSubscriptions(context.Background(), customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchlistDeprecatedAPI.ListWatchlistSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWatchlistSubscriptions`: WatchlistSubscriptionList
	fmt.Fprintf(os.Stdout, "Response from `WatchlistDeprecatedAPI.ListWatchlistSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWatchlistSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WatchlistSubscriptionList**](WatchlistSubscriptionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWatchlistAlert

> UpdateWatchlistAlert(ctx, alertId, customerId).WatchlistAlert(watchlistAlert).Execute()

Update watchlist alert

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "./openapi"
)

func main() {
	alertId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Unique identifier for this watchlist alert.
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
	watchlistAlert := *openapiclient.NewWatchlistAlert("Status_example") // WatchlistAlert | A watchlist body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WatchlistDeprecatedAPI.UpdateWatchlistAlert(context.Background(), alertId, customerId).WatchlistAlert(watchlistAlert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchlistDeprecatedAPI.UpdateWatchlistAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | Unique identifier for this watchlist alert. | 
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWatchlistAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **watchlistAlert** | [**WatchlistAlert**](WatchlistAlert.md) | A watchlist body | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWatchlistSubscription

> WatchlistSubscription UpdateWatchlistSubscription(ctx, subscriptionId, customerId).WatchlistSubscription(watchlistSubscription).Execute()

Update watchlist monitoring subscription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "./openapi"
)

func main() {
	subscriptionId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Watchlist monitoring subscription ID
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
	watchlistSubscription := *openapiclient.NewWatchlistSubscription(false) // WatchlistSubscription | Watchlist monitoring subscription to be updated. The only field that matters is `status`; all other fields are ignored. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchlistDeprecatedAPI.UpdateWatchlistSubscription(context.Background(), subscriptionId, customerId).WatchlistSubscription(watchlistSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchlistDeprecatedAPI.UpdateWatchlistSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWatchlistSubscription`: WatchlistSubscription
	fmt.Fprintf(os.Stdout, "Response from `WatchlistDeprecatedAPI.UpdateWatchlistSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Watchlist monitoring subscription ID | 
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWatchlistSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **watchlistSubscription** | [**WatchlistSubscription**](WatchlistSubscription.md) | Watchlist monitoring subscription to be updated. The only field that matters is &#x60;status&#x60;; all other fields are ignored.  | 

### Return type

[**WatchlistSubscription**](WatchlistSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

