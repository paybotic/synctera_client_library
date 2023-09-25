# \TransactionsAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPendingTransactionByID**](TransactionsAPI.md#GetPendingTransactionByID) | **Get** /transactions/pending/{id} | Get a pending transaction
[**GetPostedTransactionByID**](TransactionsAPI.md#GetPostedTransactionByID) | **Get** /transactions/posted/{id} | Get a posted transaction
[**ListPendingTransactions**](TransactionsAPI.md#ListPendingTransactions) | **Get** /transactions/pending | List pending transactions
[**ListPostedTransactions**](TransactionsAPI.md#ListPostedTransactions) | **Get** /transactions/posted | List posted transactions



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "472341e0-ea3e-41a1-96bc-fd0185e1eac8" // string | Transaction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.GetPendingTransactionByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetPendingTransactionByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingTransactionByID`: PendingTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetPendingTransactionByID`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "472341e0-ea3e-41a1-96bc-fd0185e1eac8" // string | Transaction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.GetPostedTransactionByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetPostedTransactionByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostedTransactionByID`: PostedTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetPostedTransactionByID`: %v\n", resp)
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

> PendingTransactions ListPendingTransactions(ctx).IncludeChildTransactions(includeChildTransactions).CardId(cardId).Type_(type_).FromDate(fromDate).ToDate(toDate).TransactionId(transactionId).Status(status).AccountId(accountId).IdempotencyKey(idempotencyKey).AccountNo(accountNo).ExcludeJitTransactions(excludeJitTransactions).PageToken(pageToken).ReferenceId(referenceId).Limit(limit).Subtype(subtype).Execute()

List pending transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    includeChildTransactions := true // bool | Include transactions from sub-accounts when listing transactions for a given account (optional)
    cardId := "6dc0397d-1aba-4be9-9582-895a7a887d49" // string | Card ID (optional)
    type_ := "type__example" // string | Only display transactions matching the given type (optional)
    fromDate := time.Now() // string | Only display transactions with a posting date greater than from_date (optional)
    toDate := time.Now() // string | Only display transactions with a posting date less than or equal to to_date (optional)
    transactionId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Only display holds linked to the provided transaction id (optional)
    status := []string{"Inner_example"} // []string | The status of the transaction (optional)
    accountId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Account ID (optional)
    idempotencyKey := []string{"Inner_example"} // []string | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. (optional)
    accountNo := "accountNo_example" // string | Account number (optional)
    excludeJitTransactions := true // bool | Hide \"JIT funding\" transactions from results (optional)
    pageToken := "a8937a0d" // string |  (optional)
    referenceId := "referenceId_example" // string | Reference ID (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    subtype := "subtype_example" // string | Only display transactions matching the given subtype (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.ListPendingTransactions(context.Background()).IncludeChildTransactions(includeChildTransactions).CardId(cardId).Type_(type_).FromDate(fromDate).ToDate(toDate).TransactionId(transactionId).Status(status).AccountId(accountId).IdempotencyKey(idempotencyKey).AccountNo(accountNo).ExcludeJitTransactions(excludeJitTransactions).PageToken(pageToken).ReferenceId(referenceId).Limit(limit).Subtype(subtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListPendingTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPendingTransactions`: PendingTransactions
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListPendingTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPendingTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeChildTransactions** | **bool** | Include transactions from sub-accounts when listing transactions for a given account | 
 **cardId** | **string** | Card ID | 
 **type_** | **string** | Only display transactions matching the given type | 
 **fromDate** | **string** | Only display transactions with a posting date greater than from_date | 
 **toDate** | **string** | Only display transactions with a posting date less than or equal to to_date | 
 **transactionId** | **string** | Only display holds linked to the provided transaction id | 
 **status** | **[]string** | The status of the transaction | 
 **accountId** | **string** | Account ID | 
 **idempotencyKey** | **[]string** | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. | 
 **accountNo** | **string** | Account number | 
 **excludeJitTransactions** | **bool** | Hide \&quot;JIT funding\&quot; transactions from results | 
 **pageToken** | **string** |  | 
 **referenceId** | **string** | Reference ID | 
 **limit** | **int32** |  | [default to 100]
 **subtype** | **string** | Only display transactions matching the given subtype | 

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

> PostedTransactions ListPostedTransactions(ctx).IncludeChildTransactions(includeChildTransactions).Uuid(uuid).CardId(cardId).Type_(type_).FromDate(fromDate).ToDate(toDate).AccountId(accountId).IdempotencyKey(idempotencyKey).AccountNo(accountNo).ExcludeJitTransactions(excludeJitTransactions).PageToken(pageToken).ReferenceId(referenceId).Limit(limit).Subtype(subtype).Execute()

List posted transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    includeChildTransactions := true // bool | Include transactions from sub-accounts when listing transactions for a given account (optional)
    uuid := []string{"Inner_example"} // []string | Transaction UUID(s). Multiple UUIDs can be provided as a comma-separated list. (optional)
    cardId := "6dc0397d-1aba-4be9-9582-895a7a887d49" // string | Card ID (optional)
    type_ := "type__example" // string | Only display transactions matching the given type (optional)
    fromDate := time.Now() // string | Only display transactions with a posting date greater than from_date (optional)
    toDate := time.Now() // string | Only display transactions with a posting date less than or equal to to_date (optional)
    accountId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Account ID (optional)
    idempotencyKey := []string{"Inner_example"} // []string | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. (optional)
    accountNo := "accountNo_example" // string | Account number (optional)
    excludeJitTransactions := true // bool | Hide \"JIT funding\" transactions from results (optional)
    pageToken := "a8937a0d" // string |  (optional)
    referenceId := "referenceId_example" // string | Reference ID (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    subtype := "subtype_example" // string | Only display transactions matching the given subtype (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.ListPostedTransactions(context.Background()).IncludeChildTransactions(includeChildTransactions).Uuid(uuid).CardId(cardId).Type_(type_).FromDate(fromDate).ToDate(toDate).AccountId(accountId).IdempotencyKey(idempotencyKey).AccountNo(accountNo).ExcludeJitTransactions(excludeJitTransactions).PageToken(pageToken).ReferenceId(referenceId).Limit(limit).Subtype(subtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListPostedTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPostedTransactions`: PostedTransactions
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListPostedTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPostedTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeChildTransactions** | **bool** | Include transactions from sub-accounts when listing transactions for a given account | 
 **uuid** | **[]string** | Transaction UUID(s). Multiple UUIDs can be provided as a comma-separated list. | 
 **cardId** | **string** | Card ID | 
 **type_** | **string** | Only display transactions matching the given type | 
 **fromDate** | **string** | Only display transactions with a posting date greater than from_date | 
 **toDate** | **string** | Only display transactions with a posting date less than or equal to to_date | 
 **accountId** | **string** | Account ID | 
 **idempotencyKey** | **[]string** | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. | 
 **accountNo** | **string** | Account number | 
 **excludeJitTransactions** | **bool** | Hide \&quot;JIT funding\&quot; transactions from results | 
 **pageToken** | **string** |  | 
 **referenceId** | **string** | Reference ID | 
 **limit** | **int32** |  | [default to 100]
 **subtype** | **string** | Only display transactions matching the given subtype | 

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

