# \TransactionsAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPendingTransactionByID**](TransactionsAPI.md#GetPendingTransactionByID) | **Get** /transactions/pending/{id} | Get a pending transaction
[**GetPostedTransactionByID**](TransactionsAPI.md#GetPostedTransactionByID) | **Get** /transactions/posted/{id} | Get a posted transaction
[**GetTransactionsBatchPayments**](TransactionsAPI.md#GetTransactionsBatchPayments) | **Get** /transactions/batchable | Get Transactions From Batch Payments Templates
[**ListPendingTransactions**](TransactionsAPI.md#ListPendingTransactions) | **Get** /transactions/pending | List pending transactions
[**ListPostedTransactions**](TransactionsAPI.md#ListPostedTransactions) | **Get** /transactions/posted | List posted transactions



## GetPendingTransactionByID

> PendingTransaction GetPendingTransactionByID(ctx).Id(id).Execute()

Get a pending transaction



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
	id := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string | Unique resource identifier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetPendingTransactionByID(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetPendingTransactionByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPendingTransactionByID`: PendingTransaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetPendingTransactionByID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingTransactionByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Unique resource identifier | 

### Return type

[**PendingTransaction**](PendingTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostedTransactionByID

> PostedTransaction GetPostedTransactionByID(ctx).Id(id).Execute()

Get a posted transaction



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
	id := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string | Unique resource identifier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetPostedTransactionByID(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetPostedTransactionByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostedTransactionByID`: PostedTransaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetPostedTransactionByID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostedTransactionByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Unique resource identifier | 

### Return type

[**PostedTransaction**](PostedTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsBatchPayments

> PostedTransactions GetTransactionsBatchPayments(ctx).Enabled(enabled).Execute()

Get Transactions From Batch Payments Templates



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
	enabled := true // bool | Whether the resource is enabled.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsBatchPayments(context.Background()).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsBatchPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsBatchPayments`: PostedTransactions
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsBatchPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsBatchPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** | Whether the resource is enabled.  | 

### Return type

[**PostedTransactions**](PostedTransactions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPendingTransactions

> PendingTransactions ListPendingTransactions(ctx).Uuid(uuid).UuidNin(uuidNin).IdempotencyKey(idempotencyKey).AccountNo(accountNo).ReferenceId(referenceId).AccountId(accountId).CardId(cardId).DigitalWalletTokenId(digitalWalletTokenId).FromDate(fromDate).ToDate(toDate).Status(status).TransactionId(transactionId).Type_(type_).Subtype(subtype).Cashback(cashback).IncludeChildTransactions(includeChildTransactions).ExcludeJitTransactions(excludeJitTransactions).Amount(amount).TotalAmount(totalAmount).Limit(limit).PageToken(pageToken).Execute()

List pending transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	uuid := []string{"Inner_example"} // []string | Transaction UUID(s). Multiple UUIDs can be provided as a comma-separated list. (optional)
	uuidNin := []string{"Inner_example"} // []string | Transaction UUID(s) to exclude from the query. Multiple UUIDs can be provided as a comma-separated list. (optional)
	idempotencyKey := []string{"Inner_example"} // []string | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. (optional)
	accountNo := "accountNo_example" // string | Account number (optional)
	referenceId := "referenceId_example" // string | Reference ID (optional)
	accountId := []string{"Inner_example"} // []string | Account ID (optional)
	cardId := "6dc0397d-1aba-4be9-9582-895a7a887d49" // string | Card ID (optional)
	digitalWalletTokenId := []string{"Inner_example"} // []string | Digital Wallet token ID(s). Multiple IDs can be provided as a comma-separated list. (optional)
	fromDate := time.Now() // string | Only display transactions with a posting date greater than from_date (optional)
	toDate := time.Now() // string | Only display transactions with a posting date less than or equal to to_date (optional)
	status := []string{"Inner_example"} // []string | The status of the transaction (optional)
	transactionId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Only display holds linked to the provided transaction id (optional)
	type_ := []openapiclient.DocumentType{openapiclient.document_type("ADDRESS_VERIFICATION")} // []DocumentType | The type of documents. Multiple types can be provided as a comma-separated list. (optional)
	subtype := "subtype_example" // string | Only display transactions matching the given subtype (optional)
	cashback := true // bool | If true, only transactions with cashback will be returned. If false, only transactions without cashback will be returned. If not provided, all transactions will be returned. (optional)
	includeChildTransactions := true // bool | Include transactions from sub-accounts when listing transactions for a given account (optional)
	excludeJitTransactions := true // bool | Hide \"JIT funding\" transactions from results (optional)
	amount := int32(56) // int32 | Only return transactions equal to the given `amount` (optional)
	totalAmount := int64(1200) // int64 | Only return transactions equal to the given total amount (in minor currency units, e.g., cents). This may differ from amount for holds with increases/decreases. (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListPendingTransactions(context.Background()).Uuid(uuid).UuidNin(uuidNin).IdempotencyKey(idempotencyKey).AccountNo(accountNo).ReferenceId(referenceId).AccountId(accountId).CardId(cardId).DigitalWalletTokenId(digitalWalletTokenId).FromDate(fromDate).ToDate(toDate).Status(status).TransactionId(transactionId).Type_(type_).Subtype(subtype).Cashback(cashback).IncludeChildTransactions(includeChildTransactions).ExcludeJitTransactions(excludeJitTransactions).Amount(amount).TotalAmount(totalAmount).Limit(limit).PageToken(pageToken).Execute()
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
 **uuid** | **[]string** | Transaction UUID(s). Multiple UUIDs can be provided as a comma-separated list. | 
 **uuidNin** | **[]string** | Transaction UUID(s) to exclude from the query. Multiple UUIDs can be provided as a comma-separated list. | 
 **idempotencyKey** | **[]string** | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. | 
 **accountNo** | **string** | Account number | 
 **referenceId** | **string** | Reference ID | 
 **accountId** | **[]string** | Account ID | 
 **cardId** | **string** | Card ID | 
 **digitalWalletTokenId** | **[]string** | Digital Wallet token ID(s). Multiple IDs can be provided as a comma-separated list. | 
 **fromDate** | **string** | Only display transactions with a posting date greater than from_date | 
 **toDate** | **string** | Only display transactions with a posting date less than or equal to to_date | 
 **status** | **[]string** | The status of the transaction | 
 **transactionId** | **string** | Only display holds linked to the provided transaction id | 
 **type_** | [**[]DocumentType**](DocumentType.md) | The type of documents. Multiple types can be provided as a comma-separated list. | 
 **subtype** | **string** | Only display transactions matching the given subtype | 
 **cashback** | **bool** | If true, only transactions with cashback will be returned. If false, only transactions without cashback will be returned. If not provided, all transactions will be returned. | 
 **includeChildTransactions** | **bool** | Include transactions from sub-accounts when listing transactions for a given account | 
 **excludeJitTransactions** | **bool** | Hide \&quot;JIT funding\&quot; transactions from results | 
 **amount** | **int32** | Only return transactions equal to the given &#x60;amount&#x60; | 
 **totalAmount** | **int64** | Only return transactions equal to the given total amount (in minor currency units, e.g., cents). This may differ from amount for holds with increases/decreases. | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**PendingTransactions**](PendingTransactions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostedTransactions

> PostedTransactions ListPostedTransactions(ctx).Uuid(uuid).UuidNin(uuidNin).BatchIds(batchIds).IsBatched(isBatched).IdempotencyKey(idempotencyKey).AccountNo(accountNo).AccountId(accountId).CardId(cardId).DigitalWalletTokenId(digitalWalletTokenId).FromDate(fromDate).ToDate(toDate).Type_(type_).Subtype(subtype).Cashback(cashback).ReferenceId(referenceId).IncludeChildTransactions(includeChildTransactions).ExcludeJitTransactions(excludeJitTransactions).DcSign(dcSign).Amount(amount).Limit(limit).PageToken(pageToken).Execute()

List posted transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	uuid := []string{"Inner_example"} // []string | Transaction UUID(s). Multiple UUIDs can be provided as a comma-separated list. (optional)
	uuidNin := []string{"Inner_example"} // []string | Transaction UUID(s) to exclude from the query. Multiple UUIDs can be provided as a comma-separated list. (optional)
	batchIds := []string{"Inner_example"} // []string | Batch UUID(s). Multiple UUIDs can be provided as a comma-separated list. Cannot be used with is_batched. (optional)
	isBatched := true // bool | If true, only display transactions that are part of a batch payment. Cannot be used with batch_ids. (optional)
	idempotencyKey := []string{"Inner_example"} // []string | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. (optional)
	accountNo := "accountNo_example" // string | Account number (optional)
	accountId := []string{"Inner_example"} // []string | Account ID (optional)
	cardId := "6dc0397d-1aba-4be9-9582-895a7a887d49" // string | Card ID (optional)
	digitalWalletTokenId := []string{"Inner_example"} // []string | Digital Wallet token ID(s). Multiple IDs can be provided as a comma-separated list. (optional)
	fromDate := time.Now() // string | Only display transactions with a posting date greater than from_date (optional)
	toDate := time.Now() // string | Only display transactions with a posting date less than or equal to to_date (optional)
	type_ := []openapiclient.DocumentType{openapiclient.document_type("ADDRESS_VERIFICATION")} // []DocumentType | The type of documents. Multiple types can be provided as a comma-separated list. (optional)
	subtype := "subtype_example" // string | Only display transactions matching the given subtype (optional)
	cashback := true // bool | If true, only transactions with cashback will be returned. If false, only transactions without cashback will be returned. If not provided, all transactions will be returned. (optional)
	referenceId := "referenceId_example" // string | Reference ID (optional)
	includeChildTransactions := true // bool | Include transactions from sub-accounts when listing transactions for a given account (optional)
	excludeJitTransactions := true // bool | Hide \"JIT funding\" transactions from results (optional)
	dcSign := "dcSign_example" // string | Debit/Credit sign (optional)
	amount := int32(56) // int32 | Only return transactions equal to the given `amount` (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListPostedTransactions(context.Background()).Uuid(uuid).UuidNin(uuidNin).BatchIds(batchIds).IsBatched(isBatched).IdempotencyKey(idempotencyKey).AccountNo(accountNo).AccountId(accountId).CardId(cardId).DigitalWalletTokenId(digitalWalletTokenId).FromDate(fromDate).ToDate(toDate).Type_(type_).Subtype(subtype).Cashback(cashback).ReferenceId(referenceId).IncludeChildTransactions(includeChildTransactions).ExcludeJitTransactions(excludeJitTransactions).DcSign(dcSign).Amount(amount).Limit(limit).PageToken(pageToken).Execute()
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
 **uuid** | **[]string** | Transaction UUID(s). Multiple UUIDs can be provided as a comma-separated list. | 
 **uuidNin** | **[]string** | Transaction UUID(s) to exclude from the query. Multiple UUIDs can be provided as a comma-separated list. | 
 **batchIds** | **[]string** | Batch UUID(s). Multiple UUIDs can be provided as a comma-separated list. Cannot be used with is_batched. | 
 **isBatched** | **bool** | If true, only display transactions that are part of a batch payment. Cannot be used with batch_ids. | 
 **idempotencyKey** | **[]string** | Transaction Idempotency Key(s). Multiple keys can be provided as a comma-separated list. | 
 **accountNo** | **string** | Account number | 
 **accountId** | **[]string** | Account ID | 
 **cardId** | **string** | Card ID | 
 **digitalWalletTokenId** | **[]string** | Digital Wallet token ID(s). Multiple IDs can be provided as a comma-separated list. | 
 **fromDate** | **string** | Only display transactions with a posting date greater than from_date | 
 **toDate** | **string** | Only display transactions with a posting date less than or equal to to_date | 
 **type_** | [**[]DocumentType**](DocumentType.md) | The type of documents. Multiple types can be provided as a comma-separated list. | 
 **subtype** | **string** | Only display transactions matching the given subtype | 
 **cashback** | **bool** | If true, only transactions with cashback will be returned. If false, only transactions without cashback will be returned. If not provided, all transactions will be returned. | 
 **referenceId** | **string** | Reference ID | 
 **includeChildTransactions** | **bool** | Include transactions from sub-accounts when listing transactions for a given account | 
 **excludeJitTransactions** | **bool** | Hide \&quot;JIT funding\&quot; transactions from results | 
 **dcSign** | **string** | Debit/Credit sign | 
 **amount** | **int32** | Only return transactions equal to the given &#x60;amount&#x60; | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**PostedTransactions**](PostedTransactions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

