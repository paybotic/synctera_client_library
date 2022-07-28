# \PaymentsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTransactionOut**](PaymentsApi.md#AddTransactionOut) | **Post** /ach | Create an outgoing ACH
[**GetTransactionOut**](PaymentsApi.md#GetTransactionOut) | **Get** /ach/{transaction_id} | Get an outgoing ACH transaction
[**ListTransactionsOut**](PaymentsApi.md#ListTransactionsOut) | **Get** /ach | List outgoing ACH transactions
[**PatchTransactionOut**](PaymentsApi.md#PatchTransactionOut) | **Patch** /ach/{transaction_id} | Update outgoing ACH transaction



## AddTransactionOut

> OutgoingAch AddTransactionOut(ctx).OutgoingAchRequest(outgoingAchRequest).Execute()

Create an outgoing ACH



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
    outgoingAchRequest := *openapiclient.NewOutgoingAchRequest(int32(607), "USD", "bd6fb05d-a0ba-4105-b280-51afdbc09e02", "debit", "ced6ad36-d1cd-4e73-90f5-11a07331406b", "f48a96a4-3921-4e11-a062-270697d24bec") // OutgoingAchRequest | Outgoing ACH

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.AddTransactionOut(context.Background()).OutgoingAchRequest(outgoingAchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.AddTransactionOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTransactionOut`: OutgoingAch
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.AddTransactionOut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTransactionOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outgoingAchRequest** | [**OutgoingAchRequest**](OutgoingAchRequest.md) | Outgoing ACH | 

### Return type

[**OutgoingAch**](OutgoingAch.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## GetTransactionOut

> OutgoingAch GetTransactionOut(ctx, transactionId).Execute()

Get an outgoing ACH transaction



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
    transactionId := "25ae35db-92b9-4b74-82d9-140a07eece71" // string | Transaction ID in the ledger

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.GetTransactionOut(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.GetTransactionOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionOut`: OutgoingAch
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.GetTransactionOut`: %v\n", resp)
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

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## ListTransactionsOut

> OutgoingAchList ListTransactionsOut(ctx).Limit(limit).PageToken(pageToken).Execute()

List outgoing ACH transactions



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
    pageToken := "xwsfu1mkaq" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.ListTransactionsOut(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.ListTransactionsOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsOut`: OutgoingAchList
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.ListTransactionsOut`: %v\n", resp)
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

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## PatchTransactionOut

> PatchTransactionOut(ctx, transactionId).OutgoingAchPatch(outgoingAchPatch).Execute()

Update outgoing ACH transaction



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
    transactionId := "25ae35db-92b9-4b74-82d9-140a07eece71" // string | Transaction ID in the ledger
    outgoingAchPatch := *openapiclient.NewOutgoingAchPatch() // OutgoingAchPatch | Outgoing ACH update request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PatchTransactionOut(context.Background(), transactionId).OutgoingAchPatch(outgoingAchPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PatchTransactionOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

 **outgoingAchPatch** | [**OutgoingAchPatch**](OutgoingAchPatch.md) | Outgoing ACH update request | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)

