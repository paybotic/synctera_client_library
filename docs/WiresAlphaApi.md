# \WiresAlphaApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelWire**](WiresAlphaApi.md#CancelWire) | **Patch** /wires/{wire_id} | Cancel an outgoing wire
[**CreateWire**](WiresAlphaApi.md#CreateWire) | **Post** /wires | Send a wire
[**GetWire**](WiresAlphaApi.md#GetWire) | **Get** /wires/{wire_id} | Get a wire by id
[**ListWires**](WiresAlphaApi.md#ListWires) | **Get** /wires | List wires



## CancelWire

> Wire CancelWire(ctx, wireId).UpdateTransfer(updateTransfer).Execute()

Cancel an outgoing wire



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
    wireId := "63ea71cb-7f86-44c1-956b-77f6cbd6ecf7" // string | The unique identifier of a wire
    updateTransfer := *openapiclient.NewUpdateTransfer() // UpdateTransfer | wire to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WiresAlphaApi.CancelWire(context.Background(), wireId).UpdateTransfer(updateTransfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WiresAlphaApi.CancelWire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelWire`: Wire
    fmt.Fprintf(os.Stdout, "Response from `WiresAlphaApi.CancelWire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wireId** | **string** | The unique identifier of a wire | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelWireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransfer** | [**UpdateTransfer**](UpdateTransfer.md) | wire to update | 

### Return type

[**Wire**](Wire.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWire

> Wire CreateWire(ctx).WireRequest(wireRequest).Execute()

Send a wire



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
    wireRequest := *openapiclient.NewWireRequest(int32(10000), "USD", "8281b015-d958-430a-b12e-b1b7767fe7dc", "d3d505cf-c439-42bb-be83-050653b5afba", "5e95eef3-8194-4142-86d6-11b0cb707c39") // WireRequest | Wire transfer request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WiresAlphaApi.CreateWire(context.Background()).WireRequest(wireRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WiresAlphaApi.CreateWire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWire`: Wire
    fmt.Fprintf(os.Stdout, "Response from `WiresAlphaApi.CreateWire`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wireRequest** | [**WireRequest**](WireRequest.md) | Wire transfer request | 

### Return type

[**Wire**](Wire.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWire

> Wire GetWire(ctx, wireId).Execute()

Get a wire by id



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
    wireId := "63ea71cb-7f86-44c1-956b-77f6cbd6ecf7" // string | The unique identifier of a wire

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WiresAlphaApi.GetWire(context.Background(), wireId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WiresAlphaApi.GetWire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWire`: Wire
    fmt.Fprintf(os.Stdout, "Response from `WiresAlphaApi.GetWire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wireId** | **string** | The unique identifier of a wire | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wire**](Wire.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWires

> WireList ListWires(ctx).Limit(limit).PageToken(pageToken).Status(status).CustomerId(customerId).OriginatingAccountId(originatingAccountId).ReceivingAccountId(receivingAccountId).Execute()

List wires



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
    status := "PENDING" // string |  (optional)
    customerId := "ff23b9d0-4e64-4b98-9f4a-3591ed08121a" // string |  (optional)
    originatingAccountId := "21b74a02-b7c7-49f4-a497-98ba5cbedb27" // string |  (optional)
    receivingAccountId := "3b601be1-1c91-4136-bca8-6754cc78f713" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WiresAlphaApi.ListWires(context.Background()).Limit(limit).PageToken(pageToken).Status(status).CustomerId(customerId).OriginatingAccountId(originatingAccountId).ReceivingAccountId(receivingAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WiresAlphaApi.ListWires``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWires`: WireList
    fmt.Fprintf(os.Stdout, "Response from `WiresAlphaApi.ListWires`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWiresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **status** | **string** |  | 
 **customerId** | **string** |  | 
 **originatingAccountId** | **string** |  | 
 **receivingAccountId** | **string** |  | 

### Return type

[**WireList**](WireList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

