# \ACHTransactionSimulationsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AchReturnSimulation**](ACHTransactionSimulationsApi.md#AchReturnSimulation) | **Post** /ach/transaction_simulations/receiving_return | Simulate receiving ACH return
[**AchTransactionSimulation**](ACHTransactionSimulationsApi.md#AchTransactionSimulation) | **Post** /ach/transaction_simulations/receiving_transaction | Simulate receiving ACH transaction



## AchReturnSimulation

> map[string]interface{} AchReturnSimulation(ctx).AchReturnSimulationRequest(achReturnSimulationRequest).Execute()

Simulate receiving ACH return



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
    achReturnSimulationRequest := *openapiclient.NewAchReturnSimulationRequest("23a37f14-16eb-461d-9331-b78182adbad4") // AchReturnSimulationRequest | Incoming ACH return simulation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACHTransactionSimulationsApi.AchReturnSimulation(context.Background()).AchReturnSimulationRequest(achReturnSimulationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACHTransactionSimulationsApi.AchReturnSimulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AchReturnSimulation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ACHTransactionSimulationsApi.AchReturnSimulation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAchReturnSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **achReturnSimulationRequest** | [**AchReturnSimulationRequest**](AchReturnSimulationRequest.md) | Incoming ACH return simulation | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AchTransactionSimulation

> map[string]interface{} AchTransactionSimulation(ctx).AchTransactionSimulationRequest(achTransactionSimulationRequest).Execute()

Simulate receiving ACH transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    achTransactionSimulationRequest := *openapiclient.NewAchTransactionSimulationRequest("123638791329", int32(607), "debit", time.Now()) // AchTransactionSimulationRequest | Sent ACH request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACHTransactionSimulationsApi.AchTransactionSimulation(context.Background()).AchTransactionSimulationRequest(achTransactionSimulationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACHTransactionSimulationsApi.AchTransactionSimulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AchTransactionSimulation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ACHTransactionSimulationsApi.AchTransactionSimulation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAchTransactionSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **achTransactionSimulationRequest** | [**AchTransactionSimulationRequest**](AchTransactionSimulationRequest.md) | Sent ACH request | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

