# \RemoteCheckDepositBetaAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRdcDeposit**](RemoteCheckDepositBetaAPI.md#CreateRdcDeposit) | **Post** /rdc/deposits | Create a Remote Check Deposit
[**GetRdcDeposit**](RemoteCheckDepositBetaAPI.md#GetRdcDeposit) | **Get** /rdc/deposits/{deposit_id} | Get Remote Check Deposit
[**ListRdcDeposits**](RemoteCheckDepositBetaAPI.md#ListRdcDeposits) | **Get** /rdc/deposits | List Remote Check Deposits



## CreateRdcDeposit

> DepositGet CreateRdcDeposit(ctx).DepositPost(depositPost).Execute()

Create a Remote Check Deposit



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
	depositPost := *openapiclient.NewDepositPost("b01db9c7-78f2-4a99-8aca-1231d32f9b96", "b01db9c7-78f2-4a99-8aca-1231d32f9b96", int32(12345), "USD", "b01db9c7-78f2-4a99-8aca-1231d32f9b96") // DepositPost | Attributes of the Remote Check Deposit to create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteCheckDepositBetaAPI.CreateRdcDeposit(context.Background()).DepositPost(depositPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteCheckDepositBetaAPI.CreateRdcDeposit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRdcDeposit`: DepositGet
	fmt.Fprintf(os.Stdout, "Response from `RemoteCheckDepositBetaAPI.CreateRdcDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRdcDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositPost** | [**DepositPost**](DepositPost.md) | Attributes of the Remote Check Deposit to create | 

### Return type

[**DepositGet**](DepositGet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRdcDeposit

> DepositGet GetRdcDeposit(ctx, depositId).Execute()

Get Remote Check Deposit



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
	depositId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | ID of a deposit for a remote deposit capture

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteCheckDepositBetaAPI.GetRdcDeposit(context.Background(), depositId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteCheckDepositBetaAPI.GetRdcDeposit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRdcDeposit`: DepositGet
	fmt.Fprintf(os.Stdout, "Response from `RemoteCheckDepositBetaAPI.GetRdcDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**depositId** | **string** | ID of a deposit for a remote deposit capture | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRdcDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DepositGet**](DepositGet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRdcDeposits

> DepositList ListRdcDeposits(ctx).AccountId(accountId).PageToken(pageToken).Limit(limit).Execute()

List Remote Check Deposits



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
	accountId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Unique identifier for the account. (optional)
	pageToken := "a8937a0d" // string |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteCheckDepositBetaAPI.ListRdcDeposits(context.Background()).AccountId(accountId).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteCheckDepositBetaAPI.ListRdcDeposits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRdcDeposits`: DepositList
	fmt.Fprintf(os.Stdout, "Response from `RemoteCheckDepositBetaAPI.ListRdcDeposits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRdcDepositsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Unique identifier for the account. | 
 **pageToken** | **string** |  | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**DepositList**](DepositList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

