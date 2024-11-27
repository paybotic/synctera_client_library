# \StatementsAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatement**](StatementsAPI.md#GetStatement) | **Get** /statements/{statement_id} | Get a statement
[**GetStatementTransactions**](StatementsAPI.md#GetStatementTransactions) | **Get** /statements/{statement_id}/transactions | Get a statement&#39;s transactions
[**ListStatements**](StatementsAPI.md#ListStatements) | **Get** /statements | List statements



## GetStatement

> Statement GetStatement(ctx, statementId).Execute()

Get a statement



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
	statementId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique identifier of a statement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.GetStatement(context.Background(), statementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.GetStatement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatement`: Statement
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.GetStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statementId** | **string** | The unique identifier of a statement | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Statement**](Statement.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatementTransactions

> TransactionList GetStatementTransactions(ctx, statementId).PageToken(pageToken).Limit(limit).Execute()

Get a statement's transactions



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
	statementId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique identifier of a statement
	pageToken := "a8937a0d" // string |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.GetStatementTransactions(context.Background(), statementId).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.GetStatementTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatementTransactions`: TransactionList
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.GetStatementTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statementId** | **string** | The unique identifier of a statement | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** |  | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**TransactionList**](TransactionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStatements

> StatementList ListStatements(ctx).AccountId(accountId).PageToken(pageToken).Limit(limit).Execute()

List statements



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
	accountId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The account's unique identifier provided by Synctera
	pageToken := "a8937a0d" // string |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.ListStatements(context.Background()).AccountId(accountId).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.ListStatements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStatements`: StatementList
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.ListStatements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStatementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The account&#39;s unique identifier provided by Synctera | 
 **pageToken** | **string** |  | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**StatementList**](StatementList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

