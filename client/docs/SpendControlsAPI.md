# \SpendControlsAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpendControl**](SpendControlsAPI.md#CreateSpendControl) | **Post** /spend_controls | Create Spend Control
[**GetSpendControl**](SpendControlsAPI.md#GetSpendControl) | **Get** /spend_controls/{spend_control_id} | Get Spend Control
[**ListSpendControls**](SpendControlsAPI.md#ListSpendControls) | **Get** /spend_controls | List Spend Controls
[**UpdateSpendControl**](SpendControlsAPI.md#UpdateSpendControl) | **Patch** /spend_controls/{spend_control_id} | Update Spend Control



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
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	spendControl := *openapiclient.NewSpendControl(false, false, int64(123), true, "Name_example", openapiclient.spend_control_time_range{SpendControlRollingWindowDays: openapiclient.NewSpendControlRollingWindowDays("TimeRangeType_example", int32(123))}) // SpendControl | Details of the spend control to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendControlsAPI.CreateSpendControl(context.Background()).SpendControl(spendControl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsAPI.CreateSpendControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpendControl`: SpendControlResponse
	fmt.Fprintf(os.Stdout, "Response from `SpendControlsAPI.CreateSpendControl`: %v\n", resp)
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
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	spendControlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendControlsAPI.GetSpendControl(context.Background(), spendControlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsAPI.GetSpendControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpendControl`: SpendControlResponse
	fmt.Fprintf(os.Stdout, "Response from `SpendControlsAPI.GetSpendControl`: %v\n", resp)
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

> SpendControlResponseList ListSpendControls(ctx).AmountLimitGte(amountLimitGte).AmountLimitLte(amountLimitLte).NumRelatedAccountsGte(numRelatedAccountsGte).NumRelatedAccounts(numRelatedAccounts).NumRelatedAccountsLte(numRelatedAccountsLte).SortBy(sortBy).PaymentType(paymentType).Name(name).Id(id).Direction(direction).IsActive(isActive).Limit(limit).AccountId(accountId).AmountLimit(amountLimit).Execute()

List Spend Controls



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
	amountLimitGte := int32(56) // int32 | return results with an amount limit greater than or equal to this (optional)
	amountLimitLte := int32(56) // int32 | return results with an amount limit less than or equal to this (optional)
	numRelatedAccountsGte := int32(56) // int32 | return results that are associated with at least this many accounts (optional)
	numRelatedAccounts := int32(56) // int32 | return results that are associated with this many accounts (optional)
	numRelatedAccountsLte := int32(56) // int32 | return results that are associated with at most this many accounts (optional)
	sortBy := []string{"SortBy_example"} // []string | Specifies the sort order for returned Spend Controls.  (optional)
	paymentType := openapiclient.payment_type("ACH") // PaymentType |  (optional)
	name := "name_example" // string | return results that match this name (optional)
	id := []string{"Inner_example"} // []string | return results with these comma-separated IDs (optional)
	direction := openapiclient.spend_control_direction("CREDITS") // SpendControlDirection |  (optional)
	isActive := true // bool | return results that match this active status (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	accountId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | return results matching this account ID (optional)
	amountLimit := int64(789) // int64 | return results matching this amount limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendControlsAPI.ListSpendControls(context.Background()).AmountLimitGte(amountLimitGte).AmountLimitLte(amountLimitLte).NumRelatedAccountsGte(numRelatedAccountsGte).NumRelatedAccounts(numRelatedAccounts).NumRelatedAccountsLte(numRelatedAccountsLte).SortBy(sortBy).PaymentType(paymentType).Name(name).Id(id).Direction(direction).IsActive(isActive).Limit(limit).AccountId(accountId).AmountLimit(amountLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsAPI.ListSpendControls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpendControls`: SpendControlResponseList
	fmt.Fprintf(os.Stdout, "Response from `SpendControlsAPI.ListSpendControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpendControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amountLimitGte** | **int32** | return results with an amount limit greater than or equal to this | 
 **amountLimitLte** | **int32** | return results with an amount limit less than or equal to this | 
 **numRelatedAccountsGte** | **int32** | return results that are associated with at least this many accounts | 
 **numRelatedAccounts** | **int32** | return results that are associated with this many accounts | 
 **numRelatedAccountsLte** | **int32** | return results that are associated with at most this many accounts | 
 **sortBy** | **[]string** | Specifies the sort order for returned Spend Controls.  | 
 **paymentType** | [**PaymentType**](PaymentType.md) |  | 
 **name** | **string** | return results that match this name | 
 **id** | **[]string** | return results with these comma-separated IDs | 
 **direction** | [**SpendControlDirection**](SpendControlDirection.md) |  | 
 **isActive** | **bool** | return results that match this active status | 
 **limit** | **int32** |  | [default to 100]
 **accountId** | **string** | return results matching this account ID | 
 **amountLimit** | **int64** | return results matching this amount limit | 

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
	openapiclient "github.com/paybotic/synctera_client_library/client"
)

func main() {
	spendControlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	spendControlUpdateRequest := *openapiclient.NewSpendControlUpdateRequest() // SpendControlUpdateRequest | Fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendControlsAPI.UpdateSpendControl(context.Background(), spendControlId).SpendControlUpdateRequest(spendControlUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsAPI.UpdateSpendControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSpendControl`: SpendControlResponse
	fmt.Fprintf(os.Stdout, "Response from `SpendControlsAPI.UpdateSpendControl`: %v\n", resp)
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

