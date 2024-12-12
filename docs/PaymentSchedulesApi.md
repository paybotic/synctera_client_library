# \PaymentSchedulesAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentSchedule**](PaymentSchedulesAPI.md#CreatePaymentSchedule) | **Post** /payment_schedules | Create a payment schedule
[**ListPaymentSchedules**](PaymentSchedulesAPI.md#ListPaymentSchedules) | **Get** /payment_schedules | List payment schedules
[**ListPayments**](PaymentSchedulesAPI.md#ListPayments) | **Get** /payment_schedules/payments | List payments
[**PatchPaymentSchedule**](PaymentSchedulesAPI.md#PatchPaymentSchedule) | **Patch** /payment_schedules/{payment_schedule_id} | Update a payment schedule



## CreatePaymentSchedule

> PaymentSchedule CreatePaymentSchedule(ctx).PaymentSchedule(paymentSchedule).Execute()

Create a payment schedule



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
	paymentSchedule := *openapiclient.NewPaymentSchedule("Description_example", openapiclient.payment_instruction{AchInstruction: openapiclient.NewAchInstruction(*openapiclient.NewOutgoingAchRequest(int32(607), "USD", "b01db9c7-78f2-4a99-8aca-1231d32f9b96", "debit", "b01db9c7-78f2-4a99-8aca-1231d32f9b96", "b01db9c7-78f2-4a99-8aca-1231d32f9b96"), "Type_example")}, *openapiclient.NewScheduleConfig("Frequency_example", int32(123), time.Now())) // PaymentSchedule | payment schedule to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentSchedulesAPI.CreatePaymentSchedule(context.Background()).PaymentSchedule(paymentSchedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.CreatePaymentSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentSchedule`: PaymentSchedule
	fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.CreatePaymentSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentSchedule** | [**PaymentSchedule**](PaymentSchedule.md) | payment schedule to create | 

### Return type

[**PaymentSchedule**](PaymentSchedule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentSchedules

> PaymentScheduleList ListPaymentSchedules(ctx).Id(id).AccountId(accountId).PageToken(pageToken).Limit(limit).CustomerId(customerId).Execute()

List payment schedules



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
	id := []string{"Inner_example"} // []string | IDs. Multiple IDs can be provided as a comma-separated list. (optional)
	accountId := []string{"Inner_example"} // []string | Originating account IDs. Multiple IDs can be provided as a comma-separated list. (optional)
	pageToken := "a8937a0d" // string |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	customerId := []string{"Inner_example"} // []string | The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentSchedulesAPI.ListPaymentSchedules(context.Background()).Id(id).AccountId(accountId).PageToken(pageToken).Limit(limit).CustomerId(customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.ListPaymentSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentSchedules`: PaymentScheduleList
	fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.ListPaymentSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | IDs. Multiple IDs can be provided as a comma-separated list. | 
 **accountId** | **[]string** | Originating account IDs. Multiple IDs can be provided as a comma-separated list. | 
 **pageToken** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **customerId** | **[]string** | The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. | 

### Return type

[**PaymentScheduleList**](PaymentScheduleList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayments

> PaymentList ListPayments(ctx).Id(id).AccountId(accountId).PageToken(pageToken).ScheduleId(scheduleId).Limit(limit).CustomerId(customerId).Execute()

List payments



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
	id := []string{"Inner_example"} // []string | IDs. Multiple IDs can be provided as a comma-separated list. (optional)
	accountId := []string{"Inner_example"} // []string | Originating account IDs. Multiple IDs can be provided as a comma-separated list. (optional)
	pageToken := "a8937a0d" // string |  (optional)
	scheduleId := []string{"Inner_example"} // []string | Payment schedule IDs. Multiple IDs can be provided as a comma-separated list. (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	customerId := []string{"Inner_example"} // []string | The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentSchedulesAPI.ListPayments(context.Background()).Id(id).AccountId(accountId).PageToken(pageToken).ScheduleId(scheduleId).Limit(limit).CustomerId(customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.ListPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPayments`: PaymentList
	fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.ListPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | IDs. Multiple IDs can be provided as a comma-separated list. | 
 **accountId** | **[]string** | Originating account IDs. Multiple IDs can be provided as a comma-separated list. | 
 **pageToken** | **string** |  | 
 **scheduleId** | **[]string** | Payment schedule IDs. Multiple IDs can be provided as a comma-separated list. | 
 **limit** | **int32** |  | [default to 100]
 **customerId** | **[]string** | The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. | 

### Return type

[**PaymentList**](PaymentList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPaymentSchedule

> PaymentSchedule PatchPaymentSchedule(ctx, paymentScheduleId).PatchPaymentSchedule(patchPaymentSchedule).Execute()

Update a payment schedule



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
	paymentScheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Payment schedule ID
	patchPaymentSchedule := *openapiclient.NewPatchPaymentSchedule() // PatchPaymentSchedule | payment schedule to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentSchedulesAPI.PatchPaymentSchedule(context.Background(), paymentScheduleId).PatchPaymentSchedule(patchPaymentSchedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.PatchPaymentSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPaymentSchedule`: PaymentSchedule
	fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.PatchPaymentSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentScheduleId** | **string** | Payment schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPaymentScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchPaymentSchedule** | [**PatchPaymentSchedule**](PatchPaymentSchedule.md) | payment schedule to update | 

### Return type

[**PaymentSchedule**](PaymentSchedule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

