# \InstitutionsBetaAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstitutions**](InstitutionsBetaAPI.md#GetInstitutions) | **Get** /institutions | Retrieve a list of institutions



## GetInstitutions

> InstitutionList GetInstitutions(ctx).CountryCodes(countryCodes).RoutingNumbers(routingNumbers).PageToken(pageToken).Execute()

Retrieve a list of institutions

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
	countryCodes := []string{"CountryCodes_example"} // []string | The countries of operation of the financial institutions
	routingNumbers := []string{"Inner_example"} // []string | An array of routing numbers to filter institutions. The response will only return institutions that match all of the routing numbers in the array
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsBetaAPI.GetInstitutions(context.Background()).CountryCodes(countryCodes).RoutingNumbers(routingNumbers).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsBetaAPI.GetInstitutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstitutions`: InstitutionList
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsBetaAPI.GetInstitutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstitutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodes** | **[]string** | The countries of operation of the financial institutions | 
 **routingNumbers** | **[]string** | An array of routing numbers to filter institutions. The response will only return institutions that match all of the routing numbers in the array | 
 **pageToken** | **string** |  | 

### Return type

[**InstitutionList**](InstitutionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

