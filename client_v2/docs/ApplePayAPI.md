# \ApplePayAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplePayCsr**](ApplePayAPI.md#CreateApplePayCsr) | **Post** /certificates/applepay/csr | Create an Apple Pay CSR



## CreateApplePayCsr

> *os.File CreateApplePayCsr(ctx).ApplepayCsrRequest(applepayCsrRequest).Execute()

Create an Apple Pay CSR



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
	applepayCsrRequest := *openapiclient.NewApplepayCsrRequest("merchant.com.example", "Example Inc.") // ApplepayCsrRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplePayAPI.CreateApplePayCsr(context.Background()).ApplepayCsrRequest(applepayCsrRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplePayAPI.CreateApplePayCsr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplePayCsr`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ApplePayAPI.CreateApplePayCsr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplePayCsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applepayCsrRequest** | [**ApplepayCsrRequest**](ApplepayCsrRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

