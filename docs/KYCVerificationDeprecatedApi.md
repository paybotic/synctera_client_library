# \KYCVerificationDeprecatedApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerVerificationResult**](KYCVerificationDeprecatedApi.md#CreateCustomerVerificationResult) | **Post** /customers/{customer_id}/verifications | Create a customer verification result
[**GetVerification**](KYCVerificationDeprecatedApi.md#GetVerification) | **Get** /customers/{customer_id}/verifications/{verification_id} | Get verification result
[**ListVerifications**](KYCVerificationDeprecatedApi.md#ListVerifications) | **Get** /customers/{customer_id}/verifications | List verification results
[**VerifyCustomer**](KYCVerificationDeprecatedApi.md#VerifyCustomer) | **Post** /customers/{customer_id}/verify | Verify a customer&#39;s identity



## CreateCustomerVerificationResult

> CustomerVerificationResult CreateCustomerVerificationResult(ctx, customerId).CustomerVerificationResult(customerVerificationResult).IdempotencyKey(idempotencyKey).Execute()

Create a customer verification result



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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    customerVerificationResult := *openapiclient.NewCustomerVerificationResult("ACCEPTED", time.Now(), openapiclient.verification_type("fraud")) // CustomerVerificationResult | Customer verification result to create.
    idempotencyKey := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KYCVerificationDeprecatedApi.CreateCustomerVerificationResult(context.Background(), customerId).CustomerVerificationResult(customerVerificationResult).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationDeprecatedApi.CreateCustomerVerificationResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerVerificationResult`: CustomerVerificationResult
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationDeprecatedApi.CreateCustomerVerificationResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerVerificationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerVerificationResult** | [**CustomerVerificationResult**](CustomerVerificationResult.md) | Customer verification result to create. | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. | 

### Return type

[**CustomerVerificationResult**](CustomerVerificationResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVerification

> CustomerVerificationResult GetVerification(ctx, customerId, verificationId).Execute()

Get verification result



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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    verificationId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Unique identifier for the verification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KYCVerificationDeprecatedApi.GetVerification(context.Background(), customerId, verificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationDeprecatedApi.GetVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVerification`: CustomerVerificationResult
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationDeprecatedApi.GetVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 
**verificationId** | **string** | Unique identifier for the verification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomerVerificationResult**](CustomerVerificationResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerifications

> CustomerVerificationResultList ListVerifications(ctx, customerId).IncludeHistory(includeHistory).Limit(limit).PageToken(pageToken).Execute()

List verification results



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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    includeHistory := true // bool | If true, include old (inactive) records as well. (optional) (default to false)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "a8937a0d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KYCVerificationDeprecatedApi.ListVerifications(context.Background(), customerId).IncludeHistory(includeHistory).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationDeprecatedApi.ListVerifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVerifications`: CustomerVerificationResultList
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationDeprecatedApi.ListVerifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeHistory** | **bool** | If true, include old (inactive) records as well. | [default to false]
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**CustomerVerificationResultList**](CustomerVerificationResultList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyCustomer

> CustomerVerifyResponse VerifyCustomer(ctx, customerId).CustomerVerification(customerVerification).IdempotencyKey(idempotencyKey).Execute()

Verify a customer's identity



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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    customerVerification := *openapiclient.NewCustomerVerification(false) // CustomerVerification | Customer verification request.
    idempotencyKey := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KYCVerificationDeprecatedApi.VerifyCustomer(context.Background(), customerId).CustomerVerification(customerVerification).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationDeprecatedApi.VerifyCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyCustomer`: CustomerVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationDeprecatedApi.VerifyCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerVerification** | [**CustomerVerification**](CustomerVerification.md) | Customer verification request. | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. | 

### Return type

[**CustomerVerifyResponse**](CustomerVerifyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

