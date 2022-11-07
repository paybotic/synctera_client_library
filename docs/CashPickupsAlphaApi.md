# \CashPickupsAlphaApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCashPickup**](CashPickupsAlphaApi.md#CreateCashPickup) | **Post** /cash_pickups | Create a cash pickup
[**GetCashPickup**](CashPickupsAlphaApi.md#GetCashPickup) | **Get** /cash_pickups/{cash_pickup_id} | Get a cash pickup
[**ListCashPickups**](CashPickupsAlphaApi.md#ListCashPickups) | **Get** /cash_pickups | List cash pickups
[**PatchCashPickup**](CashPickupsAlphaApi.md#PatchCashPickup) | **Patch** /cash_pickups/{cash_pickup_id} | Update a cash pickup



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
    openapiclient "./openapi"
)

func main() {
    cashPickupPostRequest := *openapiclient.NewCashPickupPostRequest("23f71110-3b25-4f3d-a1c3-915d699d8db6", int32(10000), "ReferenceId_example") // CashPickupPostRequest | cash pickup to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CashPickupsAlphaApi.CreateCashPickup(context.Background()).CashPickupPostRequest(cashPickupPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashPickupsAlphaApi.CreateCashPickup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCashPickup`: CashPickup
    fmt.Fprintf(os.Stdout, "Response from `CashPickupsAlphaApi.CreateCashPickup`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    cashPickupId := "1d2a39f5-e39b-4cd4-96a1-8727187469a9" // string | The unique identifier of a cash pickup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CashPickupsAlphaApi.GetCashPickup(context.Background(), cashPickupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashPickupsAlphaApi.GetCashPickup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCashPickup`: CashPickup
    fmt.Fprintf(os.Stdout, "Response from `CashPickupsAlphaApi.GetCashPickup`: %v\n", resp)
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

> CashPickupList ListCashPickups(ctx).Limit(limit).PageToken(pageToken).Execute()

List cash pickups



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
    pageToken := "a8937a0d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CashPickupsAlphaApi.ListCashPickups(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashPickupsAlphaApi.ListCashPickups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCashPickups`: CashPickupList
    fmt.Fprintf(os.Stdout, "Response from `CashPickupsAlphaApi.ListCashPickups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCashPickupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

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
    openapiclient "./openapi"
)

func main() {
    cashPickupId := "1d2a39f5-e39b-4cd4-96a1-8727187469a9" // string | The unique identifier of a cash pickup
    cashPickupPatchRequest := *openapiclient.NewCashPickupPatchRequest("23f71110-3b25-4f3d-a1c3-915d699d8db6", int32(10000), "ReferenceId_example") // CashPickupPatchRequest | cash pickup to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CashPickupsAlphaApi.PatchCashPickup(context.Background(), cashPickupId).CashPickupPatchRequest(cashPickupPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashPickupsAlphaApi.PatchCashPickup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCashPickup`: CashPickup
    fmt.Fprintf(os.Stdout, "Response from `CashPickupsAlphaApi.PatchCashPickup`: %v\n", resp)
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

