# \PlaidCoreExchangeBetaAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthorizationTest**](PlaidCoreExchangeBetaAPI.md#GetAuthorizationTest) | **Get** /fdx_auth_requests/authorization_test | Authorization Test
[**GetFDXToken**](PlaidCoreExchangeBetaAPI.md#GetFDXToken) | **Get** /fdx_tokens/{fdx_token_id} | Get an FDX token
[**GrantFdxAuthRequest**](PlaidCoreExchangeBetaAPI.md#GrantFdxAuthRequest) | **Post** /fdx_auth_requests/authorize | Grant an FDX authorization request
[**ListFDXAuthRequests**](PlaidCoreExchangeBetaAPI.md#ListFDXAuthRequests) | **Get** /fdx_auth_requests | List FDX authorization requests
[**ListFDXToken**](PlaidCoreExchangeBetaAPI.md#ListFDXToken) | **Get** /fdx_tokens | List tokens



## GetAuthorizationTest

> TestAuthorizeResponse GetAuthorizationTest(ctx).Execute()

Authorization Test



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaidCoreExchangeBetaAPI.GetAuthorizationTest(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaidCoreExchangeBetaAPI.GetAuthorizationTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorizationTest`: TestAuthorizeResponse
	fmt.Fprintf(os.Stdout, "Response from `PlaidCoreExchangeBetaAPI.GetAuthorizationTest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationTestRequest struct via the builder pattern


### Return type

[**TestAuthorizeResponse**](TestAuthorizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFDXToken

> FdxTokenResponse GetFDXToken(ctx, fdxTokenId).Execute()

Get an FDX token



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
	fdxTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of an FDX token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaidCoreExchangeBetaAPI.GetFDXToken(context.Background(), fdxTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaidCoreExchangeBetaAPI.GetFDXToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFDXToken`: FdxTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PlaidCoreExchangeBetaAPI.GetFDXToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fdxTokenId** | **string** | The unique identifier of an FDX token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFDXTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FdxTokenResponse**](FdxTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantFdxAuthRequest

> FdxAuthGrantResponse GrantFdxAuthRequest(ctx).FdxAuthGrantPost(fdxAuthGrantPost).Execute()

Grant an FDX authorization request



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
	fdxAuthGrantPost := *openapiclient.NewFdxAuthGrantPost("AuthRequestId_example", openapiclient.fdx_auth_grant_status("GRANTED")) // FdxAuthGrantPost | FDX authorization grant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaidCoreExchangeBetaAPI.GrantFdxAuthRequest(context.Background()).FdxAuthGrantPost(fdxAuthGrantPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaidCoreExchangeBetaAPI.GrantFdxAuthRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GrantFdxAuthRequest`: FdxAuthGrantResponse
	fmt.Fprintf(os.Stdout, "Response from `PlaidCoreExchangeBetaAPI.GrantFdxAuthRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGrantFdxAuthRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fdxAuthGrantPost** | [**FdxAuthGrantPost**](FdxAuthGrantPost.md) | FDX authorization grant | 

### Return type

[**FdxAuthGrantResponse**](FdxAuthGrantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFDXAuthRequests

> FdxAuthRequestList ListFDXAuthRequests(ctx).Limit(limit).PageToken(pageToken).Execute()

List FDX authorization requests



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
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaidCoreExchangeBetaAPI.ListFDXAuthRequests(context.Background()).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaidCoreExchangeBetaAPI.ListFDXAuthRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFDXAuthRequests`: FdxAuthRequestList
	fmt.Fprintf(os.Stdout, "Response from `PlaidCoreExchangeBetaAPI.ListFDXAuthRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFDXAuthRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**FdxAuthRequestList**](FdxAuthRequestList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFDXToken

> FdxTokenList ListFDXToken(ctx).Limit(limit).PageToken(pageToken).Execute()

List tokens



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
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaidCoreExchangeBetaAPI.ListFDXToken(context.Background()).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaidCoreExchangeBetaAPI.ListFDXToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFDXToken`: FdxTokenList
	fmt.Fprintf(os.Stdout, "Response from `PlaidCoreExchangeBetaAPI.ListFDXToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFDXTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**FdxTokenList**](FdxTokenList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

