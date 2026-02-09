# \BatchPaymentsAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBatchPaymentTemplate**](BatchPaymentsAPI.md#CreateBatchPaymentTemplate) | **Post** /batch_templates | Create Batch Payment Template
[**CreateBatchPayments**](BatchPaymentsAPI.md#CreateBatchPayments) | **Post** /batches | Create Batch Payments
[**GetBatchPayment**](BatchPaymentsAPI.md#GetBatchPayment) | **Get** /batches/{id} | Get BatchPayment
[**GetBatchPaymentTemplate**](BatchPaymentsAPI.md#GetBatchPaymentTemplate) | **Get** /batch_templates/{id} | Get Batch Payment Template
[**GetBatchPaymentTemplates**](BatchPaymentsAPI.md#GetBatchPaymentTemplates) | **Get** /batch_templates | Get Batch Payment Templates
[**GetBatchPayments**](BatchPaymentsAPI.md#GetBatchPayments) | **Get** /batches | Get Batch Payments
[**UpdateBatchPayment**](BatchPaymentsAPI.md#UpdateBatchPayment) | **Patch** /batches/{id} | Update Batch Payment
[**UpdateBatchPaymentTemplate**](BatchPaymentsAPI.md#UpdateBatchPaymentTemplate) | **Patch** /batch_templates/{id} | Update Batch Payment Template



## CreateBatchPaymentTemplate

> BatchPaymentTemplate CreateBatchPaymentTemplate(ctx).BatchPaymentTemplate(batchPaymentTemplate).Execute()

Create Batch Payment Template



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
	batchPaymentTemplate := *openapiclient.NewBatchPaymentTemplate(*openapiclient.NewBatchPaymentTemplateConfig("ExternalAccountId_example", "SettlementAccountId_example", "SettlementCustomerId_example", []string{"Subtypes_example"}, "Type_example"), "International Remittance", *openapiclient.NewBatchPaymentTemplateRules(), "abcdef_ghijkl") // BatchPaymentTemplate | Attributes of the Batch Payment template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchPaymentsAPI.CreateBatchPaymentTemplate(context.Background()).BatchPaymentTemplate(batchPaymentTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPaymentsAPI.CreateBatchPaymentTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchPaymentTemplate`: BatchPaymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `BatchPaymentsAPI.CreateBatchPaymentTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchPaymentTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchPaymentTemplate** | [**BatchPaymentTemplate**](BatchPaymentTemplate.md) | Attributes of the Batch Payment template | 

### Return type

[**BatchPaymentTemplate**](BatchPaymentTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchPayments

> CreateBatchPaymentsResponse CreateBatchPayments(ctx).CreateBatchPaymentsRequest(createBatchPaymentsRequest).Execute()

Create Batch Payments



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
	createBatchPaymentsRequest := *openapiclient.NewCreateBatchPaymentsRequest([]string{"TransactionIds_example"}) // CreateBatchPaymentsRequest | Attributes of the Batch Payment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchPaymentsAPI.CreateBatchPayments(context.Background()).CreateBatchPaymentsRequest(createBatchPaymentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPaymentsAPI.CreateBatchPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchPayments`: CreateBatchPaymentsResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchPaymentsAPI.CreateBatchPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBatchPaymentsRequest** | [**CreateBatchPaymentsRequest**](CreateBatchPaymentsRequest.md) | Attributes of the Batch Payment | 

### Return type

[**CreateBatchPaymentsResponse**](CreateBatchPaymentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPayment

> BatchPayment GetBatchPayment(ctx, id).Execute()

Get BatchPayment



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
	id := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique resource identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchPaymentsAPI.GetBatchPayment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPaymentsAPI.GetBatchPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchPayment`: BatchPayment
	fmt.Fprintf(os.Stdout, "Response from `BatchPaymentsAPI.GetBatchPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BatchPayment**](BatchPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPaymentTemplate

> BatchPaymentTemplate GetBatchPaymentTemplate(ctx, id).Execute()

Get Batch Payment Template



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
	id := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique resource identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchPaymentsAPI.GetBatchPaymentTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPaymentsAPI.GetBatchPaymentTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchPaymentTemplate`: BatchPaymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `BatchPaymentsAPI.GetBatchPaymentTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchPaymentTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BatchPaymentTemplate**](BatchPaymentTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPaymentTemplates

> []BatchPaymentTemplate GetBatchPaymentTemplates(ctx).PageToken(pageToken).Limit(limit).Id(id).Enabled(enabled).Name(name).Description(description).Execute()

Get Batch Payment Templates



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
	pageToken := "a8937a0d" // string |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	id := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string | Unique resource identifier (optional)
	enabled := true // bool | Whether or not the template is enabled. If the template is not enabled, it will not be used when creating a batch transfer. (optional)
	name := "International Remittance" // string | Name of the Batch Template (optional)
	description := "This template is used for international remittance" // string | Description of the Batch Template (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchPaymentsAPI.GetBatchPaymentTemplates(context.Background()).PageToken(pageToken).Limit(limit).Id(id).Enabled(enabled).Name(name).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPaymentsAPI.GetBatchPaymentTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchPaymentTemplates`: []BatchPaymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `BatchPaymentsAPI.GetBatchPaymentTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchPaymentTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageToken** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **id** | **[]string** | Unique resource identifier | 
 **enabled** | **bool** | Whether or not the template is enabled. If the template is not enabled, it will not be used when creating a batch transfer. | 
 **name** | **string** | Name of the Batch Template | 
 **description** | **string** | Description of the Batch Template | 

### Return type

[**[]BatchPaymentTemplate**](BatchPaymentTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPayments

> []BatchPayment GetBatchPayments(ctx).PageToken(pageToken).Limit(limit).Id(id).Status(status).TemplateId(templateId).TransactionId(transactionId).Execute()

Get Batch Payments



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
	pageToken := "a8937a0d" // string |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the Batch Payment (optional)
	status := "status_example" // string | Status of the Batch Payment (optional)
	templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the Batch Payment Template (optional)
	transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the Batch Payment Transaction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchPaymentsAPI.GetBatchPayments(context.Background()).PageToken(pageToken).Limit(limit).Id(id).Status(status).TemplateId(templateId).TransactionId(transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPaymentsAPI.GetBatchPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchPayments`: []BatchPayment
	fmt.Fprintf(os.Stdout, "Response from `BatchPaymentsAPI.GetBatchPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageToken** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **id** | **string** | ID of the Batch Payment | 
 **status** | **string** | Status of the Batch Payment | 
 **templateId** | **string** | ID of the Batch Payment Template | 
 **transactionId** | **string** | ID of the Batch Payment Transaction | 

### Return type

[**[]BatchPayment**](BatchPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBatchPayment

> BatchPayment UpdateBatchPayment(ctx, id).BatchPaymentPatchRequest(batchPaymentPatchRequest).Execute()

Update Batch Payment



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
	id := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique resource identifier
	batchPaymentPatchRequest := *openapiclient.NewBatchPaymentPatchRequest(time.Now(), time.Now()) // BatchPaymentPatchRequest | Attributes of the Batch Payment for udpate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchPaymentsAPI.UpdateBatchPayment(context.Background(), id).BatchPaymentPatchRequest(batchPaymentPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPaymentsAPI.UpdateBatchPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBatchPayment`: BatchPayment
	fmt.Fprintf(os.Stdout, "Response from `BatchPaymentsAPI.UpdateBatchPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBatchPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchPaymentPatchRequest** | [**BatchPaymentPatchRequest**](BatchPaymentPatchRequest.md) | Attributes of the Batch Payment for udpate | 

### Return type

[**BatchPayment**](BatchPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBatchPaymentTemplate

> BatchPaymentTemplate UpdateBatchPaymentTemplate(ctx, id).BatchPaymentTemplatePatch(batchPaymentTemplatePatch).Execute()

Update Batch Payment Template



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
	id := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique resource identifier
	batchPaymentTemplatePatch := *openapiclient.NewBatchPaymentTemplatePatch() // BatchPaymentTemplatePatch | Attributes of the Batch Payment template (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchPaymentsAPI.UpdateBatchPaymentTemplate(context.Background(), id).BatchPaymentTemplatePatch(batchPaymentTemplatePatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPaymentsAPI.UpdateBatchPaymentTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBatchPaymentTemplate`: BatchPaymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `BatchPaymentsAPI.UpdateBatchPaymentTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBatchPaymentTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchPaymentTemplatePatch** | [**BatchPaymentTemplatePatch**](BatchPaymentTemplatePatch.md) | Attributes of the Batch Payment template | 

### Return type

[**BatchPaymentTemplate**](BatchPaymentTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

