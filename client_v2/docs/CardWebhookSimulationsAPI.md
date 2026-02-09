# \CardWebhookSimulationsAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimulateCardFulfillmentEvent**](CardWebhookSimulationsAPI.md#SimulateCardFulfillmentEvent) | **Post** /cards/{card_id}/webhook_simulations/fulfillment | Simulate Card Fulfillment Event



## SimulateCardFulfillmentEvent

> SimulateCardFulfillment SimulateCardFulfillmentEvent(ctx).SimulateCardFulfillment(simulateCardFulfillment).CardId(cardId).Execute()

Simulate Card Fulfillment Event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	simulateCardFulfillment := *openapiclient.NewSimulateCardFulfillment(openapiclient.card_fulfillment_status("ISSUED")) // SimulateCardFulfillment | Desired simulated fulfillment status change value
	cardId := "6dc0397d-1aba-4be9-9582-895a7a887d49" // string | Card ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardWebhookSimulationsAPI.SimulateCardFulfillmentEvent(context.Background()).SimulateCardFulfillment(simulateCardFulfillment).CardId(cardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardWebhookSimulationsAPI.SimulateCardFulfillmentEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SimulateCardFulfillmentEvent`: SimulateCardFulfillment
	fmt.Fprintf(os.Stdout, "Response from `CardWebhookSimulationsAPI.SimulateCardFulfillmentEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateCardFulfillmentEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simulateCardFulfillment** | [**SimulateCardFulfillment**](SimulateCardFulfillment.md) | Desired simulated fulfillment status change value | 
 **cardId** | **string** | Card ID | 

### Return type

[**SimulateCardFulfillment**](SimulateCardFulfillment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

