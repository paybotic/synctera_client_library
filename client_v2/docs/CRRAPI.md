# \CRRAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCRR**](CRRAPI.md#ListCRR) | **Get** /crr | List CRRs



## ListCRR

> CrrList ListCRR(ctx).Limit(limit).PageToken(pageToken).CustomerId(customerId).BusinessId(businessId).IncludeHistory(includeHistory).Execute()

List CRRs



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
	customerId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of a customer. (optional)
	businessId := "edebece0-4a06-4806-b062-7de25a848be8" // string | ID of the business that the risk score applies to (optional)
	includeHistory := true // bool | If true, include old risk scores in the response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CRRAPI.ListCRR(context.Background()).Limit(limit).PageToken(pageToken).CustomerId(customerId).BusinessId(businessId).IncludeHistory(includeHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CRRAPI.ListCRR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCRR`: CrrList
	fmt.Fprintf(os.Stdout, "Response from `CRRAPI.ListCRR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCRRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **customerId** | **string** | The unique identifier of a customer. | 
 **businessId** | **string** | ID of the business that the risk score applies to | 
 **includeHistory** | **bool** | If true, include old risk scores in the response | [default to false]

### Return type

[**CrrList**](CrrList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

