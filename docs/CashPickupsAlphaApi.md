# \CashPickupsAlphaAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCashPickup**](CashPickupsAlphaAPI.md#CreateCashPickup) | **Post** /cash_pickups | Create a cash pickup
[**GetCashPickup**](CashPickupsAlphaAPI.md#GetCashPickup) | **Get** /cash_pickups/{cash_pickup_id} | Get a cash pickup
[**ListCashPickups**](CashPickupsAlphaAPI.md#ListCashPickups) | **Get** /cash_pickups | List cash pickups
[**PatchCashPickup**](CashPickupsAlphaAPI.md#PatchCashPickup) | **Patch** /cash_pickups/{cash_pickup_id} | Update a cash pickup



## CreateCashPickup

> CashPickup CreateCashPickup(ctx).CashPickupPostRequest(cashPickupPostRequest).Execute()

Create a cash pickup



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
    cashPickupPostRequest := *openapiclient.NewCashPickupPostRequest("23f71110-3b25-4f3d-a1c3-915d699d8db6", int32(10000), "ReferenceId_example") // CashPickupPostRequest | cash pickup to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CashPickupsAlphaAPI.CreateCashPickup(context.Background()).CashPickupPostRequest(cashPickupPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashPickupsAlphaAPI.CreateCashPickup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCashPickup`: CashPickup
    fmt.Fprintf(os.Stdout, "Response from `CashPickupsAlphaAPI.CreateCashPickup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCashPickupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cashPickupPostRequest** | [**CashPickupPostRequest**](CashPickupPostRequest.md) | cash pickup to create | 

### Return type

[**CashPickup**](CashPickup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCashPickup

> CashPickup GetCashPickup(ctx, cashPickupId).Execute()

Get a cash pickup



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
    cashPickupId := "1d2a39f5-e39b-4cd4-96a1-8727187469a9" // string | The unique identifier of a cash pickup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CashPickupsAlphaAPI.GetCashPickup(context.Background(), cashPickupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashPickupsAlphaAPI.GetCashPickup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCashPickup`: CashPickup
    fmt.Fprintf(os.Stdout, "Response from `CashPickupsAlphaAPI.GetCashPickup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cashPickupId** | **string** | The unique identifier of a cash pickup | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCashPickupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CashPickup**](CashPickup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCashPickups

> CashPickupList ListCashPickups(ctx).PageToken(pageToken).Limit(limit).Execute()

List cash pickups



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
    pageToken := "a8937a0d" // string |  (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CashPickupsAlphaAPI.ListCashPickups(context.Background()).PageToken(pageToken).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashPickupsAlphaAPI.ListCashPickups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCashPickups`: CashPickupList
    fmt.Fprintf(os.Stdout, "Response from `CashPickupsAlphaAPI.ListCashPickups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCashPickupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageToken** | **string** |  | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**CashPickupList**](CashPickupList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCashPickup

> CashPickup PatchCashPickup(ctx, cashPickupId).CashPickupPatchRequest(cashPickupPatchRequest).Execute()

Update a cash pickup



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
    cashPickupId := "1d2a39f5-e39b-4cd4-96a1-8727187469a9" // string | The unique identifier of a cash pickup
    cashPickupPatchRequest := *openapiclient.NewCashPickupPatchRequest() // CashPickupPatchRequest | cash pickup to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CashPickupsAlphaAPI.PatchCashPickup(context.Background(), cashPickupId).CashPickupPatchRequest(cashPickupPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashPickupsAlphaAPI.PatchCashPickup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCashPickup`: CashPickup
    fmt.Fprintf(os.Stdout, "Response from `CashPickupsAlphaAPI.PatchCashPickup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cashPickupId** | **string** | The unique identifier of a cash pickup | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCashPickupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cashPickupPatchRequest** | [**CashPickupPatchRequest**](CashPickupPatchRequest.md) | cash pickup to update | 

### Return type

[**CashPickup**](CashPickup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

