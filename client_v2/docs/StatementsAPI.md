# \StatementsAPI

All URIs are relative to *https://api.synctera.com/v2*

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
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatementTransactions

> TransactionList GetStatementTransactions(ctx, statementId).Limit(limit).PageToken(pageToken).Execute()

Get a statement's transactions



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
	statementId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique identifier of a statement
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.GetStatementTransactions(context.Background(), statementId).Limit(limit).PageToken(pageToken).Execute()
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

 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**TransactionList**](TransactionList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStatements

> StatementList ListStatements(ctx).AccountId(accountId).Limit(limit).PageToken(pageToken).Execute()

List statements



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
	accountId := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string | Account ID (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.ListStatements(context.Background()).AccountId(accountId).Limit(limit).PageToken(pageToken).Execute()
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
 **accountId** | **[]string** | Account ID | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**StatementList**](StatementList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

