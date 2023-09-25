# \MerchantsAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MxReadMerchant**](MerchantsAPI.md#MxReadMerchant) | **Get** /mx/merchants/{merchant_guid} | Get merchant from MX



## MxReadMerchant

> MerchantResponse MxReadMerchant(ctx, merchantGuid).Execute()

Get merchant from MX



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
    merchantGuid := "MCH-7ed79542-884d-2b1b-dd74-501c5cc9d25b" // string | The unique id for a `merchant`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerchantsAPI.MxReadMerchant(context.Background(), merchantGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsAPI.MxReadMerchant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MxReadMerchant`: MerchantResponse
    fmt.Fprintf(os.Stdout, "Response from `MerchantsAPI.MxReadMerchant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantGuid** | **string** | The unique id for a &#x60;merchant&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMxReadMerchantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MerchantResponse**](MerchantResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

