# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimulateCardFulfillmentEvent**](CardWebhookSimulationsApi.md#SimulateCardFulfillmentEvent) | **Post** /cards/{card_id}/webhook_simulations/fulfillment | Simulate Card Fulfillment Event

# **SimulateCardFulfillmentEvent**
> SimulateCardFulfillment SimulateCardFulfillmentEvent(ctx, body, cardId)
Simulate Card Fulfillment Event

This endpoint is for testing environment only to trigger a simulated change in card fulfillment status event 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimulateCardFulfillment**](SimulateCardFulfillment.md)| Desired simulated fulfillment status change value | 
  **cardId** | [**string**](.md)|  | 

### Return type

[**SimulateCardFulfillment**](simulate_card_fulfillment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

