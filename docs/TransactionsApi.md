# \TransactionsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPendingTransactionByID**](TransactionsApi.md#GetPendingTransactionByID) | **Get** /transactions/pending/{id} | Get a pending transaction
[**GetPostedTransactionByID**](TransactionsApi.md#GetPostedTransactionByID) | **Get** /transactions/posted/{id} | Get a posted transaction
[**ListPendingTransactions**](TransactionsApi.md#ListPendingTransactions) | **Get** /transactions/pending | List pending transactions
[**ListPostedTransactions**](TransactionsApi.md#ListPostedTransactions) | **Get** /transactions/posted | List posted transactions



## GetPendingTransactionByID

> PendingTransaction GetPendingTransactionByID(ctx, id).Execute()

Get a pending transaction



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
    id := "472341e0-ea3e-41a1-96bc-fd0185e1eac8" // string | Transaction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetPendingTransactionByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetPendingTransactionByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingTransactionByID`: PendingTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetPendingTransactionByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingTransactionByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PendingTransaction**](PendingTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostedTransactionByID

> PostedTransaction GetPostedTransactionByID(ctx, id).Execute()

Get a posted transaction



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
    id := "472341e0-ea3e-41a1-96bc-fd0185e1eac8" // string | Transaction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetPostedTransactionByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetPostedTransactionByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostedTransactionByID`: PostedTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetPostedTransactionByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostedTransactionByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostedTransaction**](PostedTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPendingTransactions

> PendingTransactions ListPendingTransactions(ctx).AccountNo(accountNo).AccountId(accountId).FromDate(fromDate).ToDate(toDate).Status(status).TransactionId(transactionId).Type_(type_).Subtype(subtype).Limit(limit).PageToken(pageToken).Execute()

List pending transactions



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
    accountNo := "accountNo_example" // string | Account number (optional)
    accountId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Account ID (optional)
    fromDate := time.Now() // string | Only display transactions with a posting date greater than from_date (optional)
    toDate := time.Now() // string | Only display transactions with a posting date less than or equal to to_date (optional)
    status := "status_example" // string | The status of the transaction (optional)
    transactionId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Only display holds linked to the provided transaction id (optional)
    type_ := "type__example" // string | Only display transactions matching the given type (optional)
    subtype := "subtype_example" // string | Only display transactions matching the given subtype (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "a8937a0d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.ListPendingTransactions(context.Background()).AccountNo(accountNo).AccountId(accountId).FromDate(fromDate).ToDate(toDate).Status(status).TransactionId(transactionId).Type_(type_).Subtype(subtype).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListPendingTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPendingTransactions`: PendingTransactions
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListPendingTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPendingTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNo** | **string** | Account number | 
 **accountId** | **string** | Account ID | 
 **fromDate** | **string** | Only display transactions with a posting date greater than from_date | 
 **toDate** | **string** | Only display transactions with a posting date less than or equal to to_date | 
 **status** | **string** | The status of the transaction | 
 **transactionId** | **string** | Only display holds linked to the provided transaction id | 
 **type_** | **string** | Only display transactions matching the given type | 
 **subtype** | **string** | Only display transactions matching the given subtype | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**PendingTransactions**](PendingTransactions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostedTransactions

> PostedTransactions ListPostedTransactions(ctx).AccountNo(accountNo).AccountId(accountId).FromDate(fromDate).ToDate(toDate).Type_(type_).Subtype(subtype).Limit(limit).PageToken(pageToken).Execute()

List posted transactions



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
    accountNo := "accountNo_example" // string | Account number (optional)
    accountId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Account ID (optional)
    fromDate := time.Now() // string | Only display transactions with a posting date greater than from_date (optional)
    toDate := time.Now() // string | Only display transactions with a posting date less than or equal to to_date (optional)
    type_ := "type__example" // string | Only display transactions matching the given type (optional)
    subtype := "subtype_example" // string | Only display transactions matching the given subtype (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "a8937a0d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.ListPostedTransactions(context.Background()).AccountNo(accountNo).AccountId(accountId).FromDate(fromDate).ToDate(toDate).Type_(type_).Subtype(subtype).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListPostedTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPostedTransactions`: PostedTransactions
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListPostedTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPostedTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNo** | **string** | Account number | 
 **accountId** | **string** | Account ID | 
 **fromDate** | **string** | Only display transactions with a posting date greater than from_date | 
 **toDate** | **string** | Only display transactions with a posting date less than or equal to to_date | 
 **type_** | **string** | Only display transactions matching the given type | 
 **subtype** | **string** | Only display transactions matching the given subtype | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**PostedTransactions**](PostedTransactions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

