# \AutopaysAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutopay**](AutopaysAPI.md#GetAutopay) | **Get** /autopays/{autopay_id} | Get autopay
[**ListAutopays**](AutopaysAPI.md#ListAutopays) | **Get** /autopays | List autopays
[**PatchAutopay**](AutopaysAPI.md#PatchAutopay) | **Patch** /autopays/{autopay_id} | Update autopay



## GetAutopay

> Autopay GetAutopay(ctx, autopayId).Execute()

Get autopay



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
	autopayId := "2d0f7601-ecc6-48b4-b82f-5d64fa84628b" // string | Autopay ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutopaysAPI.GetAutopay(context.Background(), autopayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutopaysAPI.GetAutopay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutopay`: Autopay
	fmt.Fprintf(os.Stdout, "Response from `AutopaysAPI.GetAutopay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autopayId** | **string** | Autopay ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutopayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Autopay**](Autopay.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutopays

> AutopayList ListAutopays(ctx).Limit(limit).PageToken(pageToken).AutopayConfigId(autopayConfigId).BillingPeriodId(billingPeriodId).FromCreationTime(fromCreationTime).FromLastUpdatedTime(fromLastUpdatedTime).FromScheduledDate(fromScheduledDate).Id(id).LendingAccountId(lendingAccountId).RenderedDescription(renderedDescription).RenderedDescriptionLike(renderedDescriptionLike).ScheduledDate(scheduledDate).SortBy(sortBy).StatementId(statementId).Status(status).Tenant(tenant).ToCreationTime(toCreationTime).ToLastUpdatedTime(toLastUpdatedTime).ToScheduledDate(toScheduledDate).Execute()

List autopays



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)
	autopayConfigId := []string{"Inner_example"} // []string |  (optional)
	billingPeriodId := []string{"Inner_example"} // []string |  (optional)
	fromCreationTime := time.Now() // time.Time |  (optional)
	fromLastUpdatedTime := time.Now() // time.Time |  (optional)
	fromScheduledDate := time.Now() // string |  (optional)
	id := []string{"Inner_example"} // []string |  (optional)
	lendingAccountId := []string{"Inner_example"} // []string |  (optional)
	renderedDescription := []string{"Inner_example"} // []string |  (optional)
	renderedDescriptionLike := "renderedDescriptionLike_example" // string | Case insensitive wildcard search for rendered_description, wildcards can be specified with '*'. Wildcards at both the start and the end of the input is assumed. (optional)
	scheduledDate := time.Now() // string |  (optional)
	sortBy := []string{"SortBy_example"} // []string |  (optional)
	statementId := []string{"Inner_example"} // []string |  (optional)
	status := openapiclient.autopay_status("PENDING") // AutopayStatus |  (optional)
	tenant := []string{"abcdef_ghijkl"} // []string |  (optional)
	toCreationTime := time.Now() // time.Time |  (optional)
	toLastUpdatedTime := time.Now() // time.Time |  (optional)
	toScheduledDate := time.Now() // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutopaysAPI.ListAutopays(context.Background()).Limit(limit).PageToken(pageToken).AutopayConfigId(autopayConfigId).BillingPeriodId(billingPeriodId).FromCreationTime(fromCreationTime).FromLastUpdatedTime(fromLastUpdatedTime).FromScheduledDate(fromScheduledDate).Id(id).LendingAccountId(lendingAccountId).RenderedDescription(renderedDescription).RenderedDescriptionLike(renderedDescriptionLike).ScheduledDate(scheduledDate).SortBy(sortBy).StatementId(statementId).Status(status).Tenant(tenant).ToCreationTime(toCreationTime).ToLastUpdatedTime(toLastUpdatedTime).ToScheduledDate(toScheduledDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutopaysAPI.ListAutopays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutopays`: AutopayList
	fmt.Fprintf(os.Stdout, "Response from `AutopaysAPI.ListAutopays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAutopaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **autopayConfigId** | **[]string** |  | 
 **billingPeriodId** | **[]string** |  | 
 **fromCreationTime** | **time.Time** |  | 
 **fromLastUpdatedTime** | **time.Time** |  | 
 **fromScheduledDate** | **string** |  | 
 **id** | **[]string** |  | 
 **lendingAccountId** | **[]string** |  | 
 **renderedDescription** | **[]string** |  | 
 **renderedDescriptionLike** | **string** | Case insensitive wildcard search for rendered_description, wildcards can be specified with &#39;*&#39;. Wildcards at both the start and the end of the input is assumed. | 
 **scheduledDate** | **string** |  | 
 **sortBy** | **[]string** |  | 
 **statementId** | **[]string** |  | 
 **status** | [**AutopayStatus**](AutopayStatus.md) |  | 
 **tenant** | **[]string** |  | 
 **toCreationTime** | **time.Time** |  | 
 **toLastUpdatedTime** | **time.Time** |  | 
 **toScheduledDate** | **string** |  | 

### Return type

[**AutopayList**](AutopayList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAutopay

> Autopay PatchAutopay(ctx, autopayId).AutopayUpdateRequest(autopayUpdateRequest).Execute()

Update autopay



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
	autopayId := "2d0f7601-ecc6-48b4-b82f-5d64fa84628b" // string | Autopay ID
	autopayUpdateRequest := *openapiclient.NewAutopayUpdateRequest() // AutopayUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutopaysAPI.PatchAutopay(context.Background(), autopayId).AutopayUpdateRequest(autopayUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutopaysAPI.PatchAutopay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAutopay`: Autopay
	fmt.Fprintf(os.Stdout, "Response from `AutopaysAPI.PatchAutopay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autopayId** | **string** | Autopay ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAutopayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autopayUpdateRequest** | [**AutopayUpdateRequest**](AutopayUpdateRequest.md) |  | 

### Return type

[**Autopay**](Autopay.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

