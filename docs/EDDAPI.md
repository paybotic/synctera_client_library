# \EDDAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEDD**](EDDAPI.md#CreateEDD) | **Post** /edd | Create a EDD
[**DeleteEDD**](EDDAPI.md#DeleteEDD) | **Delete** /edd/{edd_id} | Delete a EDD
[**GetEDD**](EDDAPI.md#GetEDD) | **Get** /edd/{edd_id} | Get a EDD
[**ListEDD**](EDDAPI.md#ListEDD) | **Get** /edd | List EDD



## CreateEDD

> CreateEddResponse CreateEDD(ctx).CreateEddRequest(createEddRequest).Execute()

Create a EDD



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
	createEddRequest := openapiclient.create_edd_request{EddAccount: openapiclient.NewEddAccount("Reason_example", "RelatedResourceId_example", openapiclient.related_resource_type1("ACCOUNT"))} // CreateEddRequest | EDD to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EDDAPI.CreateEDD(context.Background()).CreateEddRequest(createEddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EDDAPI.CreateEDD``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEDD`: CreateEddResponse
	fmt.Fprintf(os.Stdout, "Response from `EDDAPI.CreateEDD`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEDDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEddRequest** | [**CreateEddRequest**](CreateEddRequest.md) | EDD to create | 

### Return type

[**CreateEddResponse**](CreateEddResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEDD

> DeleteResponse DeleteEDD(ctx, eddId).Execute()

Delete a EDD



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
	eddId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of a edd

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EDDAPI.DeleteEDD(context.Background(), eddId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EDDAPI.DeleteEDD``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEDD`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `EDDAPI.DeleteEDD`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eddId** | **string** | The unique identifier of a edd | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEDDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEDD

> CreateEddResponse GetEDD(ctx, eddId).Execute()

Get a EDD



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
	eddId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of a edd

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EDDAPI.GetEDD(context.Background(), eddId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EDDAPI.GetEDD``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEDD`: CreateEddResponse
	fmt.Fprintf(os.Stdout, "Response from `EDDAPI.GetEDD`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eddId** | **string** | The unique identifier of a edd | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEDDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateEddResponse**](CreateEddResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEDD

> EddList ListEDD(ctx).RelatedResourceType(relatedResourceType).RelatedResourceId(relatedResourceId).PageToken(pageToken).Limit(limit).Execute()

List EDD



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
	relatedResourceType := []openapiclient.RelatedResourceType1{openapiclient.related_resource_type1("ACCOUNT")} // []RelatedResourceType1 | Type of related resource. Multiple values can be provided as a comma-separated list. 
	relatedResourceId := []string{"Inner_example"} // []string | Unique identifier for the related resource. Multiple IDs can be provided as a comma-separated list.  (optional)
	pageToken := "a8937a0d" // string |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EDDAPI.ListEDD(context.Background()).RelatedResourceType(relatedResourceType).RelatedResourceId(relatedResourceId).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EDDAPI.ListEDD``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEDD`: EddList
	fmt.Fprintf(os.Stdout, "Response from `EDDAPI.ListEDD`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEDDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relatedResourceType** | [**[]RelatedResourceType1**](RelatedResourceType1.md) | Type of related resource. Multiple values can be provided as a comma-separated list.  | 
 **relatedResourceId** | **[]string** | Unique identifier for the related resource. Multiple IDs can be provided as a comma-separated list.  | 
 **pageToken** | **string** |  | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**EddList**](EddList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

