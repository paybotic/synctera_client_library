# \ApplicationsBetaAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](ApplicationsBetaAPI.md#CreateApplication) | **Post** /applications | Create an application
[**GetApplication**](ApplicationsBetaAPI.md#GetApplication) | **Get** /applications/{application_id} | Get an application
[**ListApplications**](ApplicationsBetaAPI.md#ListApplications) | **Get** /applications | List applications
[**PatchApplication**](ApplicationsBetaAPI.md#PatchApplication) | **Patch** /applications/{application_id} | Modify an application



## CreateApplication

> ApplicationResponse CreateApplication(ctx).Application(application).Execute()

Create an application



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
	application := openapiclient.application{CreditApplication: openapiclient.NewCreditApplication(openapiclient.application_account_type("LINE_OF_CREDIT"), []openapiclient.Applicant{*openapiclient.NewApplicant(false)}, openapiclient.credit_application_purpose("ACCOUNT_OPENING"), openapiclient.credit_application_status("SUBMITTED"), openapiclient.application_type("CREDIT"))} // Application | Application model.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsBetaAPI.CreateApplication(context.Background()).Application(application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaAPI.CreateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplication`: ApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsBetaAPI.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | [**Application**](Application.md) | Application model. | 

### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> ApplicationResponse GetApplication(ctx, applicationId).Execute()

Get an application



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
	applicationId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Unique identifier for the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsBetaAPI.GetApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaAPI.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: ApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsBetaAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Unique identifier for the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> ApplicationList ListApplications(ctx).CustomerId(customerId).BusinessId(businessId).Status(status).Limit(limit).PageToken(pageToken).Type_(type_).SortBy(sortBy).Execute()

List applications



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
	customerId := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string | A list of customer unique identifiers, with a comma separating any values. (optional)
	businessId := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string | A list of business unique identifiers, with a comma separating any values. Returns applications that are linked to the specified Business IDs. (optional)
	status := openapiclient.all_application_status("APPROVED") // AllApplicationStatus | Application status values for all types of applications  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)
	type_ := openapiclient.application_type("CREDIT") // ApplicationType | Type of application  (optional)
	sortBy := []string{"SortBy_example"} // []string | Specifies the sort order for the returned applications.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsBetaAPI.ListApplications(context.Background()).CustomerId(customerId).BusinessId(businessId).Status(status).Limit(limit).PageToken(pageToken).Type_(type_).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaAPI.ListApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplications`: ApplicationList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsBetaAPI.ListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **[]string** | A list of customer unique identifiers, with a comma separating any values. | 
 **businessId** | **[]string** | A list of business unique identifiers, with a comma separating any values. Returns applications that are linked to the specified Business IDs. | 
 **status** | [**AllApplicationStatus**](AllApplicationStatus.md) | Application status values for all types of applications  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **type_** | [**ApplicationType**](ApplicationType.md) | Type of application  | 
 **sortBy** | **[]string** | Specifies the sort order for the returned applications.  | 

### Return type

[**ApplicationList**](ApplicationList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApplication

> ApplicationResponse PatchApplication(ctx, applicationId).ApplicationPatch(applicationPatch).Execute()

Modify an application



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
	applicationId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | Unique identifier for the application.
	applicationPatch := *openapiclient.NewApplicationPatch() // ApplicationPatch | Application fields to be patched.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsBetaAPI.PatchApplication(context.Background(), applicationId).ApplicationPatch(applicationPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaAPI.PatchApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApplication`: ApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsBetaAPI.PatchApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Unique identifier for the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationPatch** | [**ApplicationPatch**](ApplicationPatch.md) | Application fields to be patched. | 

### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

