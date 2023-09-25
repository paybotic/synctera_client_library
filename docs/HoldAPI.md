# \HoldAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHold**](HoldAPI.md#GetHold) | **Get** /hold/{id} | Get a pending transaction
[**ListHolds**](HoldAPI.md#ListHolds) | **Get** /hold | List holds



## GetHold

> PendingTransaction GetHold(ctx, id).Execute()

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
    resp, r, err := apiClient.HoldAPI.GetHold(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HoldAPI.GetHold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHold`: PendingTransaction
    fmt.Fprintf(os.Stdout, "Response from `HoldAPI.GetHold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldRequest struct via the builder pattern


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


## ListHolds

> PendingTransactions ListHolds(ctx).IncludeChildTransactions(includeChildTransactions).FromDate(fromDate).ToDate(toDate).TransactionId(transactionId).Status(status).Type_(type_).AccountId(accountId).IdempotencyKey(idempotencyKey).AccountNo(accountNo).ExcludeJitTransactions(excludeJitTransactions).PageToken(pageToken).CardId(cardId).ReferenceId(referenceId).Limit(limit).Subtype(subtype).Execute()

List holds



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
    fromDate := time.Now() // string | Only display transactions with a posting date greater than from_date (optional)
    toDate := time.Now() // string | Only display transactions with a posting date less than or equal to to_date (optional)
    transactionId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Only display holds linked to the provided transaction id (optional)
    status := []string{"Inner_example"} // []string | The status of the transaction (optional)
    type_ := "type__example" // string | Only display transactions matching the given type (optional)
    accountId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Account ID (optional)
    idempotencyKey := []string{"Inner_example"} // []string | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. (optional)
    accountNo := "accountNo_example" // string | Account number (optional)
    excludeJitTransactions := true // bool | Hide \"JIT funding\" transactions from results (optional)
    pageToken := "a8937a0d" // string |  (optional)
    cardId := "6dc0397d-1aba-4be9-9582-895a7a887d49" // string | Card ID (optional)
    referenceId := "referenceId_example" // string | Reference ID (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    subtype := "subtype_example" // string | Only display transactions matching the given subtype (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HoldAPI.ListHolds(context.Background()).IncludeChildTransactions(includeChildTransactions).FromDate(fromDate).ToDate(toDate).TransactionId(transactionId).Status(status).Type_(type_).AccountId(accountId).IdempotencyKey(idempotencyKey).AccountNo(accountNo).ExcludeJitTransactions(excludeJitTransactions).PageToken(pageToken).CardId(cardId).ReferenceId(referenceId).Limit(limit).Subtype(subtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HoldAPI.ListHolds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHolds`: PendingTransactions
    fmt.Fprintf(os.Stdout, "Response from `HoldAPI.ListHolds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHoldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeChildTransactions** | **bool** | Include transactions from sub-accounts when listing transactions for a given account | 
 **fromDate** | **string** | Only display transactions with a posting date greater than from_date | 
 **toDate** | **string** | Only display transactions with a posting date less than or equal to to_date | 
 **transactionId** | **string** | Only display holds linked to the provided transaction id | 
 **status** | **[]string** | The status of the transaction | 
 **type_** | **string** | Only display transactions matching the given type | 
 **accountId** | **string** | Account ID | 
 **idempotencyKey** | **[]string** | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. | 
 **accountNo** | **string** | Account number | 
 **excludeJitTransactions** | **bool** | Hide \&quot;JIT funding\&quot; transactions from results | 
 **pageToken** | **string** |  | 
 **cardId** | **string** | Card ID | 
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

