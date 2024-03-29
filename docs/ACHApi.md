# \ACHApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTransactionOut**](ACHApi.md#AddTransactionOut) | **Post** /ach | Send an ACH
[**GetTransactionOut**](ACHApi.md#GetTransactionOut) | **Get** /ach/{transaction_id} | Get a sent ACH transaction
[**ListTransactionsOut**](ACHApi.md#ListTransactionsOut) | **Get** /ach | List sent ACH transactions
[**PatchTransactionOut**](ACHApi.md#PatchTransactionOut) | **Patch** /ach/{transaction_id} | Update a sent ACH transaction



## AddTransactionOut

> OutgoingAch AddTransactionOut(ctx).OutgoingAchRequest(outgoingAchRequest).IdempotencyKey(idempotencyKey).Execute()

Send an ACH



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
    outgoingAchRequest := *openapiclient.NewOutgoingAchRequest(int32(607), "USD", "e9061cf1-b472-4066-bf4b-2491b33fd182", "debit", "b58b33b8-9629-4414-91ee-62c0b200c7b4", "3110c3f4-7937-4010-bded-86917767f50b") // OutgoingAchRequest | Send ACH request
    idempotencyKey := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACHApi.AddTransactionOut(context.Background()).OutgoingAchRequest(outgoingAchRequest).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACHApi.AddTransactionOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTransactionOut`: OutgoingAch
    fmt.Fprintf(os.Stdout, "Response from `ACHApi.AddTransactionOut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTransactionOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outgoingAchRequest** | [**OutgoingAchRequest**](OutgoingAchRequest.md) | Send ACH request | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. | 

### Return type

[**OutgoingAch**](OutgoingAch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionOut

> OutgoingAch GetTransactionOut(ctx, transactionId).Execute()

Get a sent ACH transaction



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
    transactionId := "1985f769-dd31-4128-95db-f765355e6631" // string | Transaction ID in the ledger

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACHApi.GetTransactionOut(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACHApi.GetTransactionOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionOut`: OutgoingAch
    fmt.Fprintf(os.Stdout, "Response from `ACHApi.GetTransactionOut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID in the ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutgoingAch**](OutgoingAch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsOut

> OutgoingAchList ListTransactionsOut(ctx).Limit(limit).PageToken(pageToken).Execute()

List sent ACH transactions



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
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "a8937a0d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACHApi.ListTransactionsOut(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACHApi.ListTransactionsOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsOut`: OutgoingAchList
    fmt.Fprintf(os.Stdout, "Response from `ACHApi.ListTransactionsOut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**OutgoingAchList**](OutgoingAchList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTransactionOut

> OutgoingAch PatchTransactionOut(ctx, transactionId).OutgoingAchPatch(outgoingAchPatch).Execute()

Update a sent ACH transaction



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
    transactionId := "1985f769-dd31-4128-95db-f765355e6631" // string | Transaction ID in the ledger
    outgoingAchPatch := *openapiclient.NewOutgoingAchPatch() // OutgoingAchPatch | Update sent ach transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACHApi.PatchTransactionOut(context.Background(), transactionId).OutgoingAchPatch(outgoingAchPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACHApi.PatchTransactionOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchTransactionOut`: OutgoingAch
    fmt.Fprintf(os.Stdout, "Response from `ACHApi.PatchTransactionOut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID in the ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTransactionOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **outgoingAchPatch** | [**OutgoingAchPatch**](OutgoingAchPatch.md) | Update sent ach transaction | 

### Return type

[**OutgoingAch**](OutgoingAch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

