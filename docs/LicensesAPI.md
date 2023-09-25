# \LicensesAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLicenses**](LicensesAPI.md#CreateLicenses) | **Post** /licenses | Create a license and associate it with a business
[**GetLicense**](LicensesAPI.md#GetLicense) | **Get** /licenses/{license_id} | Get license
[**ListLicenses**](LicensesAPI.md#ListLicenses) | **Get** /licenses | list licenses
[**PatchLicense**](LicensesAPI.md#PatchLicense) | **Patch** /licenses/{license_id} | Patch License



## CreateLicenses

> ResponseLicense CreateLicenses(ctx).PostLicense(postLicense).Execute()

Create a license and associate it with a business



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
    postLicense := *openapiclient.NewPostLicense() // PostLicense | license to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesAPI.CreateLicenses(context.Background()).PostLicense(postLicense).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.CreateLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLicenses`: ResponseLicense
    fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.CreateLicenses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postLicense** | [**PostLicense**](PostLicense.md) | license to be created | 

### Return type

[**ResponseLicense**](ResponseLicense.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicense

> ResponseLicense GetLicense(ctx, licenseId).Execute()

Get license



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
    licenseId := "ba4f84d7-fa20-4010-a573-0bbca57ab589" // string | License record identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesAPI.GetLicense(context.Background(), licenseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.GetLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicense`: ResponseLicense
    fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.GetLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseId** | **string** | License record identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseLicense**](ResponseLicense.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLicenses

> LicenseList ListLicenses(ctx).BusinessId(businessId).CustomerId(customerId).Execute()

list licenses

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
    businessId := []string{"Inner_example"} // []string | Unique identifier for the business. Multiple IDs can be provided as a comma-separated list.  (optional)
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesAPI.ListLicenses(context.Background()).BusinessId(businessId).CustomerId(customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.ListLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLicenses`: LicenseList
    fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.ListLicenses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessId** | **[]string** | Unique identifier for the business. Multiple IDs can be provided as a comma-separated list.  | 
 **customerId** | **string** | The customer&#39;s unique identifier | 

### Return type

[**LicenseList**](LicenseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLicense

> ResponseLicense PatchLicense(ctx, licenseId).LicensePatch(licensePatch).Execute()

Patch License



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
    licenseId := "ba4f84d7-fa20-4010-a573-0bbca57ab589" // string | License record identifier
    licensePatch := *openapiclient.NewLicensePatch("10321") // LicensePatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesAPI.PatchLicense(context.Background(), licenseId).LicensePatch(licensePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.PatchLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchLicense`: ResponseLicense
    fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.PatchLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseId** | **string** | License record identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licensePatch** | [**LicensePatch**](LicensePatch.md) |  | 

### Return type

[**ResponseLicense**](ResponseLicense.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

