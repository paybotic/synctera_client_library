# \IncomingWiresAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIncomingWire**](IncomingWiresAPI.md#GetIncomingWire) | **Get** /wires/incoming/{wire_id} | Get incoming wire by id
[**ListIncomingWires**](IncomingWiresAPI.md#ListIncomingWires) | **Get** /wires/incoming | List incoming wires
[**PatchIncomingWire**](IncomingWiresAPI.md#PatchIncomingWire) | **Patch** /wires/incoming/{wire_id} | Update an incoming wire by id



## GetIncomingWire

> IncomingWire GetIncomingWire(ctx, wireId).Execute()

Get incoming wire by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	wireId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of a wire transfer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncomingWiresAPI.GetIncomingWire(context.Background(), wireId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncomingWiresAPI.GetIncomingWire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomingWire`: IncomingWire
	fmt.Fprintf(os.Stdout, "Response from `IncomingWiresAPI.GetIncomingWire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wireId** | **string** | The unique identifier of a wire transfer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomingWireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IncomingWire**](IncomingWire.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingWires

> IncomingWireList ListIncomingWires(ctx).ToDate(toDate).FromDate(fromDate).Status(status).PageToken(pageToken).Id(id).Limit(limit).Execute()

List incoming wires



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	toDate := time.Now() // string | Only display transactions with an effective date less than or equal to to_date (optional)
	fromDate := time.Now() // string | Only display transactions with an effective date greater than from_date (optional)
	status := "READY" // string |  (optional)
	pageToken := "a8937a0d" // string |  (optional)
	id := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string | Unique resource identifier (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncomingWiresAPI.ListIncomingWires(context.Background()).ToDate(toDate).FromDate(fromDate).Status(status).PageToken(pageToken).Id(id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncomingWiresAPI.ListIncomingWires``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIncomingWires`: IncomingWireList
	fmt.Fprintf(os.Stdout, "Response from `IncomingWiresAPI.ListIncomingWires`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIncomingWiresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toDate** | **string** | Only display transactions with an effective date less than or equal to to_date | 
 **fromDate** | **string** | Only display transactions with an effective date greater than from_date | 
 **status** | **string** |  | 
 **pageToken** | **string** |  | 
 **id** | **[]string** | Unique resource identifier | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**IncomingWireList**](IncomingWireList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIncomingWire

> IncomingWire PatchIncomingWire(ctx, wireId).IncomingWirePatch(incomingWirePatch).Execute()

Update an incoming wire by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	wireId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of a wire transfer.
	incomingWirePatch := *openapiclient.NewIncomingWirePatch() // IncomingWirePatch | Patch incoming wire request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncomingWiresAPI.PatchIncomingWire(context.Background(), wireId).IncomingWirePatch(incomingWirePatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncomingWiresAPI.PatchIncomingWire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchIncomingWire`: IncomingWire
	fmt.Fprintf(os.Stdout, "Response from `IncomingWiresAPI.PatchIncomingWire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wireId** | **string** | The unique identifier of a wire transfer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIncomingWireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **incomingWirePatch** | [**IncomingWirePatch**](IncomingWirePatch.md) | Patch incoming wire request | 

### Return type

[**IncomingWire**](IncomingWire.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

