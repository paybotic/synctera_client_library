# \RemoteCheckDepositBetaApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRdcDeposit**](RemoteCheckDepositBetaApi.md#CreateRdcDeposit) | **Post** /rdc/deposits | Create a Remote Check Deposit
[**GetRdcDeposit**](RemoteCheckDepositBetaApi.md#GetRdcDeposit) | **Get** /rdc/deposits/{deposit_id} | Get Remote Check Deposit
[**ListRdcDeposits**](RemoteCheckDepositBetaApi.md#ListRdcDeposits) | **Get** /rdc/deposits | List Remote Check Deposits



## CreateRdcDeposit

> Deposit CreateRdcDeposit(ctx).Deposit(deposit).Execute()

Create a Remote Check Deposit



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
    deposit := *openapiclient.NewDeposit("a2b17b47-49a3-4f69-88bf-b7a0abb555f2", "b6d8065d-7122-40df-922e-6f2ba88a5d50", int32(5005), time.Now(), time.Now(), int32(2979), "USD", "350bba5f-27e1-4eec-ad0a-94b322d7fc59", "38ebff87-fb06-44a0-9a06-72f0175dfc62", "Status_example", "50f24553-6017-4d36-82b9-43a251ee4358", openapiclient.vendor_info1{VendorJson: openapiclient.NewVendorJson("ContentType_example", map[string]interface{}(123), "SOCURE")}) // Deposit | Attributes of the Remote Check Deposit to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteCheckDepositBetaApi.CreateRdcDeposit(context.Background()).Deposit(deposit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteCheckDepositBetaApi.CreateRdcDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRdcDeposit`: Deposit
    fmt.Fprintf(os.Stdout, "Response from `RemoteCheckDepositBetaApi.CreateRdcDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRdcDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deposit** | [**Deposit**](Deposit.md) | Attributes of the Remote Check Deposit to create | 

### Return type

[**Deposit**](Deposit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRdcDeposit

> Deposit GetRdcDeposit(ctx, depositId).Execute()

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
    depositId := "4d463562-87b0-40bf-ba3f-7b52833f4898" // string | ID of a deposit for a remote deposit capture

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteCheckDepositBetaApi.GetRdcDeposit(context.Background(), depositId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteCheckDepositBetaApi.GetRdcDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRdcDeposit`: Deposit
    fmt.Fprintf(os.Stdout, "Response from `RemoteCheckDepositBetaApi.GetRdcDeposit`: %v\n", resp)
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

[**Deposit**](Deposit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRdcDeposits

> DepositList ListRdcDeposits(ctx).AccountId(accountId).Limit(limit).PageToken(pageToken).Execute()

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
    accountId := "12dbc24c-3cfc-46dc-a164-0977434afabf" // string | Unique identifier for the account. (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "a8937a0d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteCheckDepositBetaApi.ListRdcDeposits(context.Background()).AccountId(accountId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteCheckDepositBetaApi.ListRdcDeposits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRdcDeposits`: DepositList
    fmt.Fprintf(os.Stdout, "Response from `RemoteCheckDepositBetaApi.ListRdcDeposits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRdcDepositsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Unique identifier for the account. | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

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

