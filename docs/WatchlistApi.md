# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWatchlistAlert**](WatchlistApi.md#GetWatchlistAlert) | **Get** /customers/{customer_id}/watchlists/alerts/{alert_id} | Retrieve watchlist monitoring alert
[**GetWatchlistSubscription**](WatchlistApi.md#GetWatchlistSubscription) | **Get** /customers/{customer_id}/watchlists/subscriptions/{subscription_id} | Retrieve watchlist monitoring subscription
[**ListWatchlistAlerts**](WatchlistApi.md#ListWatchlistAlerts) | **Get** /customers/{customer_id}/watchlists/alerts | List watchlist monitoring alerts for a customer
[**ListWatchlistSubscriptions**](WatchlistApi.md#ListWatchlistSubscriptions) | **Get** /customers/{customer_id}/watchlists/subscriptions | List watchlist monitoring subscriptions for a customer
[**SuppressWatchlistEntityAlert**](WatchlistApi.md#SuppressWatchlistEntityAlert) | **Post** /customers/{customer_id}/watchlists/suppressions | Suppress entity alert
[**UpdateWatchlistAlert**](WatchlistApi.md#UpdateWatchlistAlert) | **Put** /customers/{customer_id}/watchlists/alerts/{alert_id} | Update watchlist alert
[**UpdateWatchlistSubscription**](WatchlistApi.md#UpdateWatchlistSubscription) | **Put** /customers/{customer_id}/watchlists/subscriptions/{subscription_id} | Update watchlist monitoring subscription
[**WatchlistSubscribe**](WatchlistApi.md#WatchlistSubscribe) | **Post** /customers/{customer_id}/watchlists/subscriptions | Subscribe a customer to watchlist monitoring

# **GetWatchlistAlert**
> WatchlistAlert GetWatchlistAlert(ctx, customerId, alertId)
Retrieve watchlist monitoring alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | [**string**](.md)| The customer&#x27;s unique identifier | 
  **alertId** | [**string**](.md)| Unique identifier for this watchlist alert. | 

### Return type

[**WatchlistAlert**](watchlist_alert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWatchlistSubscription**
> WatchlistSubscription GetWatchlistSubscription(ctx, customerId, subscriptionId)
Retrieve watchlist monitoring subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | [**string**](.md)| The customer&#x27;s unique identifier | 
  **subscriptionId** | [**string**](.md)| Watchlist monitoring subscription ID | 

### Return type

[**WatchlistSubscription**](watchlist_subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWatchlistAlerts**
> WatchlistAlertList ListWatchlistAlerts(ctx, customerId)
List watchlist monitoring alerts for a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | [**string**](.md)| The customer&#x27;s unique identifier | 

### Return type

[**WatchlistAlertList**](watchlist_alert_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWatchlistSubscriptions**
> WatchlistSubscriptionList ListWatchlistSubscriptions(ctx, customerId)
List watchlist monitoring subscriptions for a customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | [**string**](.md)| The customer&#x27;s unique identifier | 

### Return type

[**WatchlistSubscriptionList**](watchlist_subscription_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuppressWatchlistEntityAlert**
> SuppressWatchlistEntityAlert(ctx, body, customerId, optional)
Suppress entity alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WatchlistSuppress**](WatchlistSuppress.md)| A watchlist suppression object | 
  **customerId** | [**string**](.md)| The customer&#x27;s unique identifier | 
 **optional** | ***WatchlistApiSuppressWatchlistEntityAlertOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WatchlistApiSuppressWatchlistEntityAlertOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **optional.**| An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWatchlistAlert**
> UpdateWatchlistAlert(ctx, body, customerId, alertId)
Update watchlist alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WatchlistAlert**](WatchlistAlert.md)| A watchlist body | 
  **customerId** | [**string**](.md)| The customer&#x27;s unique identifier | 
  **alertId** | [**string**](.md)| Unique identifier for this watchlist alert. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWatchlistSubscription**
> WatchlistSubscription UpdateWatchlistSubscription(ctx, body, customerId, subscriptionId)
Update watchlist monitoring subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WatchlistSubscription**](WatchlistSubscription.md)| Watchlist monitoring subscription to be updated. The only field that matters is &#x60;status&#x60;; all other fields are ignored.
 | 
  **customerId** | [**string**](.md)| The customer&#x27;s unique identifier | 
  **subscriptionId** | [**string**](.md)| Watchlist monitoring subscription ID | 

### Return type

[**WatchlistSubscription**](watchlist_subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchlistSubscribe**
> WatchlistSubscription WatchlistSubscribe(ctx, body, customerId, optional)
Subscribe a customer to watchlist monitoring

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WatchlistSubscription**](WatchlistSubscription.md)| A watchlist subscription | 
  **customerId** | [**string**](.md)| The customer&#x27;s unique identifier | 
 **optional** | ***WatchlistApiWatchlistSubscribeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WatchlistApiWatchlistSubscribeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **optional.**| An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

[**WatchlistSubscription**](watchlist_subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

