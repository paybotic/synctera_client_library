# \SpendControlsBetaAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpendControl**](SpendControlsBetaAPI.md#CreateSpendControl) | **Post** /spend_controls | Create Spend Control
[**GetSpendControl**](SpendControlsBetaAPI.md#GetSpendControl) | **Get** /spend_controls/{spend_control_id} | Get Spend Control
[**ListSpendControls**](SpendControlsBetaAPI.md#ListSpendControls) | **Get** /spend_controls | List Spend Controls
[**UpdateSpendControl**](SpendControlsBetaAPI.md#UpdateSpendControl) | **Patch** /spend_controls/{spend_control_id} | Update Spend Control



## CreateSpendControl

> SpendControlResponse CreateSpendControl(ctx).SpendControlCreationRequest(spendControlCreationRequest).Execute()

Create Spend Control



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
	spendControlCreationRequest := *openapiclient.NewSpendControlCreationRequest(false, false, openapiclient.spend_control_direction("DEBITS"), true, "Name_example", openapiclient.spend_control_time_range{SpendControlForever: openapiclient.NewSpendControlForever(openapiclient.spend_control_time_range_type("ROLLING_WINDOW_DAYS"))}) // SpendControlCreationRequest | Details of the spend control to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendControlsBetaAPI.CreateSpendControl(context.Background()).SpendControlCreationRequest(spendControlCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsBetaAPI.CreateSpendControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpendControl`: SpendControlResponse
	fmt.Fprintf(os.Stdout, "Response from `SpendControlsBetaAPI.CreateSpendControl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpendControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spendControlCreationRequest** | [**SpendControlCreationRequest**](SpendControlCreationRequest.md) | Details of the spend control to create | 

### Return type

[**SpendControlResponse**](SpendControlResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	spendControlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendControlsBetaAPI.GetSpendControl(context.Background(), spendControlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsBetaAPI.GetSpendControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpendControl`: SpendControlResponse
	fmt.Fprintf(os.Stdout, "Response from `SpendControlsBetaAPI.GetSpendControl`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpendControls

> SpendControlResponseList ListSpendControls(ctx).AccountId(accountId).PaymentType(paymentType).AmountLimit(amountLimit).AmountLimitGte(amountLimitGte).AmountLimitLte(amountLimitLte).TransactionCountLimit(transactionCountLimit).TransactionCountLimitGte(transactionCountLimitGte).TransactionCountLimitLte(transactionCountLimitLte).NumRelatedAccounts(numRelatedAccounts).NumRelatedAccountsGte(numRelatedAccountsGte).NumRelatedAccountsLte(numRelatedAccountsLte).IsActive(isActive).Name(name).Direction(direction).Limit(limit).Id(id).SortBy(sortBy).TimeRangeType(timeRangeType).Execute()

List Spend Controls



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
	accountId := []string{"Inner_example"} // []string | Originating account IDs. Multiple IDs can be provided as a comma-separated list. (optional)
	paymentType := openapiclient.payment_type("ACH") // PaymentType |  (optional)
	amountLimit := int64(789) // int64 | return results matching this amount limit (optional)
	amountLimitGte := int32(56) // int32 | return results with an amount limit greater than or equal to this (optional)
	amountLimitLte := int32(56) // int32 | return results with an amount limit less than or equal to this (optional)
	transactionCountLimit := int32(56) // int32 | return results matching this transaction count limit (optional)
	transactionCountLimitGte := int32(56) // int32 | return results with a transaction count limit greater than or equal to this (optional)
	transactionCountLimitLte := int32(56) // int32 | return results with a transaction count limit less than or equal to this (optional)
	numRelatedAccounts := int32(56) // int32 | return results that are associated with this many accounts (optional)
	numRelatedAccountsGte := int32(56) // int32 | return results that are associated with at least this many accounts (optional)
	numRelatedAccountsLte := int32(56) // int32 | return results that are associated with at most this many accounts (optional)
	isActive := true // bool | return results that match this active status (optional)
	name := "name_example" // string | return results that match this name (optional)
	direction := openapiclient.spend_control_direction("DEBITS") // SpendControlDirection |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	id := []string{"Inner_example"} // []string | return results with these comma-separated IDs (optional)
	sortBy := []string{"SortBy_example"} // []string | Specifies the sort order for returned Spend Controls.  (optional)
	timeRangeType := []string{"TimeRangeType_example"} // []string | return results with this time range type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendControlsBetaAPI.ListSpendControls(context.Background()).AccountId(accountId).PaymentType(paymentType).AmountLimit(amountLimit).AmountLimitGte(amountLimitGte).AmountLimitLte(amountLimitLte).TransactionCountLimit(transactionCountLimit).TransactionCountLimitGte(transactionCountLimitGte).TransactionCountLimitLte(transactionCountLimitLte).NumRelatedAccounts(numRelatedAccounts).NumRelatedAccountsGte(numRelatedAccountsGte).NumRelatedAccountsLte(numRelatedAccountsLte).IsActive(isActive).Name(name).Direction(direction).Limit(limit).Id(id).SortBy(sortBy).TimeRangeType(timeRangeType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsBetaAPI.ListSpendControls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpendControls`: SpendControlResponseList
	fmt.Fprintf(os.Stdout, "Response from `SpendControlsBetaAPI.ListSpendControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpendControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **[]string** | Originating account IDs. Multiple IDs can be provided as a comma-separated list. | 
 **paymentType** | [**PaymentType**](PaymentType.md) |  | 
 **amountLimit** | **int64** | return results matching this amount limit | 
 **amountLimitGte** | **int32** | return results with an amount limit greater than or equal to this | 
 **amountLimitLte** | **int32** | return results with an amount limit less than or equal to this | 
 **transactionCountLimit** | **int32** | return results matching this transaction count limit | 
 **transactionCountLimitGte** | **int32** | return results with a transaction count limit greater than or equal to this | 
 **transactionCountLimitLte** | **int32** | return results with a transaction count limit less than or equal to this | 
 **numRelatedAccounts** | **int32** | return results that are associated with this many accounts | 
 **numRelatedAccountsGte** | **int32** | return results that are associated with at least this many accounts | 
 **numRelatedAccountsLte** | **int32** | return results that are associated with at most this many accounts | 
 **isActive** | **bool** | return results that match this active status | 
 **name** | **string** | return results that match this name | 
 **direction** | [**SpendControlDirection**](SpendControlDirection.md) |  | 
 **limit** | **int32** |  | [default to 100]
 **id** | **[]string** | return results with these comma-separated IDs | 
 **sortBy** | **[]string** | Specifies the sort order for returned Spend Controls.  | 
 **timeRangeType** | **[]string** | return results with this time range type | 

### Return type

[**SpendControlResponseList**](SpendControlResponseList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	spendControlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	spendControlUpdateRequest := *openapiclient.NewSpendControlUpdateRequest() // SpendControlUpdateRequest | Fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendControlsBetaAPI.UpdateSpendControl(context.Background(), spendControlId).SpendControlUpdateRequest(spendControlUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendControlsBetaAPI.UpdateSpendControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSpendControl`: SpendControlResponse
	fmt.Fprintf(os.Stdout, "Response from `SpendControlsBetaAPI.UpdateSpendControl`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

