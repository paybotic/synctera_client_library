# \WireTransactionSimulationsAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WireReturnSimulation**](WireTransactionSimulationsAPI.md#WireReturnSimulation) | **Post** /wires/transaction_simulations/receiving_return | Simulate receiving Wire transfer return
[**WireTransactionSimulation**](WireTransactionSimulationsAPI.md#WireTransactionSimulation) | **Post** /wires/transaction_simulations/receiving_transaction | Simulate receiving Wire transaction



## WireReturnSimulation

> WireSimulationResponse WireReturnSimulation(ctx).WireReturnSimulationRequest(wireReturnSimulationRequest).Execute()

Simulate receiving Wire transfer return



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
	wireReturnSimulationRequest := *openapiclient.NewWireReturnSimulationRequest("23a37f14-16eb-461d-9331-b78182adbad4") // WireReturnSimulationRequest | Incoming Wire return simulation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireTransactionSimulationsAPI.WireReturnSimulation(context.Background()).WireReturnSimulationRequest(wireReturnSimulationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireTransactionSimulationsAPI.WireReturnSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WireReturnSimulation`: WireSimulationResponse
	fmt.Fprintf(os.Stdout, "Response from `WireTransactionSimulationsAPI.WireReturnSimulation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWireReturnSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wireReturnSimulationRequest** | [**WireReturnSimulationRequest**](WireReturnSimulationRequest.md) | Incoming Wire return simulation request | 

### Return type

[**WireSimulationResponse**](WireSimulationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WireTransactionSimulation

> WireSimulationResponse WireTransactionSimulation(ctx).WireTransactionSimulationRequest(wireTransactionSimulationRequest).Execute()

Simulate receiving Wire transaction



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
	wireTransactionSimulationRequest := *openapiclient.NewWireTransactionSimulationRequest("123638791329", int32(607)) // WireTransactionSimulationRequest | Incoming Wire simulation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireTransactionSimulationsAPI.WireTransactionSimulation(context.Background()).WireTransactionSimulationRequest(wireTransactionSimulationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireTransactionSimulationsAPI.WireTransactionSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WireTransactionSimulation`: WireSimulationResponse
	fmt.Fprintf(os.Stdout, "Response from `WireTransactionSimulationsAPI.WireTransactionSimulation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWireTransactionSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wireTransactionSimulationRequest** | [**WireTransactionSimulationRequest**](WireTransactionSimulationRequest.md) | Incoming Wire simulation request | 

### Return type

[**WireSimulationResponse**](WireSimulationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

