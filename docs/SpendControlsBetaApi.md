# \SpendControlsBetaApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpendControl**](SpendControlsBetaApi.md#CreateSpendControl) | **Post** /spend_controls | Create Spend Control
[**GetSpendControl**](SpendControlsBetaApi.md#GetSpendControl) | **Get** /spend_controls/{spend_control_id} | Get Spend Control
[**ListSpendControls**](SpendControlsBetaApi.md#ListSpendControls) | **Get** /spend_controls | List Spend Controls
[**UpdateSpendControl**](SpendControlsBetaApi.md#UpdateSpendControl) | **Patch** /spend_controls/{spend_control_id} | Update Spend Control



## CreateSpendControl

> SpendControlResponse CreateSpendControl(ctx).SpendControl(spendControl).Execute()

Create Spend Control



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
    spendControl := *openapiclient.NewSpendControl(false, false, int64(123), true, "Name_example", openapiclient.spend_control_time_range{SpendControlRollingWindowDays: openapiclient.NewSpendControlRollingWindowDays("TimeRangeType_example", int32(123))}) // SpendControl | Details of the spend control to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendControlsBetaApi.CreateSpendControl(context.Background()).SpendControl(spendControl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsBetaApi.CreateSpendControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpendControl`: SpendControlResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendControlsBetaApi.CreateSpendControl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpendControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spendControl** | [**SpendControl**](SpendControl.md) | Details of the spend control to create | 

### Return type

[**SpendControlResponse**](SpendControlResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpendControl

> SpendControlResponse GetSpendControl(ctx, spendControlId).Execute()

Get Spend Control



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
    spendControlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendControlsBetaApi.GetSpendControl(context.Background(), spendControlId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsBetaApi.GetSpendControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpendControl`: SpendControlResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendControlsBetaApi.GetSpendControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spendControlId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpendControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpendControlResponse**](SpendControlResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpendControls

> SpendControlResponseList ListSpendControls(ctx).AccountId(accountId).PaymentType(paymentType).AmountLimit(amountLimit).AmountLimitGte(amountLimitGte).AmountLimitLte(amountLimitLte).NumRelatedAccounts(numRelatedAccounts).NumRelatedAccountsGte(numRelatedAccountsGte).NumRelatedAccountsLte(numRelatedAccountsLte).IsActive(isActive).Name(name).Direction(direction).Id(id).SortBy(sortBy).Execute()

List Spend Controls



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
    accountId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | return results matching this account ID (optional)
    paymentType := openapiclient.payment_type("CARD") // PaymentType |  (optional)
    amountLimit := int64(789) // int64 | return results matching this amount limit (optional)
    amountLimitGte := int32(56) // int32 | return results with an amount limit greater than or equal to this (optional)
    amountLimitLte := int32(56) // int32 | return results with an amount limit less than or equal to this (optional)
    numRelatedAccounts := int32(56) // int32 | return results that are associated with this many accounts (optional)
    numRelatedAccountsGte := int32(56) // int32 | return results that are associated with at least this many accounts (optional)
    numRelatedAccountsLte := int32(56) // int32 | return results that are associated with at most this many accounts (optional)
    isActive := true // bool | return results that match this active status (optional)
    name := "name_example" // string | return results that match this name (optional)
    direction := openapiclient.spend_control_direction("DEBITS") // SpendControlDirection |  (optional)
    id := []string{"Inner_example"} // []string | return results with these comma-separated IDs (optional)
    sortBy := []string{"SortBy_example"} // []string | Specifies the sort order for returned Spend Controls.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendControlsBetaApi.ListSpendControls(context.Background()).AccountId(accountId).PaymentType(paymentType).AmountLimit(amountLimit).AmountLimitGte(amountLimitGte).AmountLimitLte(amountLimitLte).NumRelatedAccounts(numRelatedAccounts).NumRelatedAccountsGte(numRelatedAccountsGte).NumRelatedAccountsLte(numRelatedAccountsLte).IsActive(isActive).Name(name).Direction(direction).Id(id).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsBetaApi.ListSpendControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpendControls`: SpendControlResponseList
    fmt.Fprintf(os.Stdout, "Response from `SpendControlsBetaApi.ListSpendControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpendControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | return results matching this account ID | 
 **paymentType** | [**PaymentType**](PaymentType.md) |  | 
 **amountLimit** | **int64** | return results matching this amount limit | 
 **amountLimitGte** | **int32** | return results with an amount limit greater than or equal to this | 
 **amountLimitLte** | **int32** | return results with an amount limit less than or equal to this | 
 **numRelatedAccounts** | **int32** | return results that are associated with this many accounts | 
 **numRelatedAccountsGte** | **int32** | return results that are associated with at least this many accounts | 
 **numRelatedAccountsLte** | **int32** | return results that are associated with at most this many accounts | 
 **isActive** | **bool** | return results that match this active status | 
 **name** | **string** | return results that match this name | 
 **direction** | [**SpendControlDirection**](SpendControlDirection.md) |  | 
 **id** | **[]string** | return results with these comma-separated IDs | 
 **sortBy** | **[]string** | Specifies the sort order for returned Spend Controls.  | 

### Return type

[**SpendControlResponseList**](SpendControlResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpendControl

> SpendControlResponse UpdateSpendControl(ctx, spendControlId).SpendControlUpdateRequest(spendControlUpdateRequest).Execute()

Update Spend Control



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
    spendControlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    spendControlUpdateRequest := *openapiclient.NewSpendControlUpdateRequest() // SpendControlUpdateRequest | Fields to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendControlsBetaApi.UpdateSpendControl(context.Background(), spendControlId).SpendControlUpdateRequest(spendControlUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsBetaApi.UpdateSpendControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpendControl`: SpendControlResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendControlsBetaApi.UpdateSpendControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spendControlId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpendControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spendControlUpdateRequest** | [**SpendControlUpdateRequest**](SpendControlUpdateRequest.md) | Fields to update | 

### Return type

[**SpendControlResponse**](SpendControlResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

