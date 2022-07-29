# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentSchedule**](PaymentSchedulesApi.md#CreatePaymentSchedule) | **Post** /payment_schedules | Create a payment schedule
[**ListPaymentSchedules**](PaymentSchedulesApi.md#ListPaymentSchedules) | **Get** /payment_schedules | List payment schedules
[**ListPayments**](PaymentSchedulesApi.md#ListPayments) | **Get** /payment_schedules/payments | List payments
[**PatchPaymentSchedule**](PaymentSchedulesApi.md#PatchPaymentSchedule) | **Patch** /payment_schedules/{payment_schedule_id} | Update a payment schedule

# **CreatePaymentSchedule**
> PaymentSchedule CreatePaymentSchedule(ctx, body)
Create a payment schedule

Create a payment schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PaymentSchedule**](PaymentSchedule.md)| payment schedule to create | 

### Return type

[**PaymentSchedule**](payment_schedule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentSchedules**
> PaymentScheduleList ListPaymentSchedules(ctx, optional)
List payment schedules

Get paginated list of payment schedules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentSchedulesApiListPaymentSchedulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentSchedulesApiListPaymentSchedulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 
 **id** | [**optional.Interface of []string**](.md)| IDs. Multiple IDs can be provided as a comma-separated list. | 
 **accountId** | [**optional.Interface of []string**](.md)| Originating account IDs. Multiple IDs can be provided as a comma-separated list. | 
 **customerId** | [**optional.Interface of []string**](.md)| The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. | 

### Return type

[**PaymentScheduleList**](payment_schedule_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPayments**
> PaymentList ListPayments(ctx, optional)
List payments

Get paginated list of payments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentSchedulesApiListPaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentSchedulesApiListPaymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 
 **id** | [**optional.Interface of []string**](.md)| IDs. Multiple IDs can be provided as a comma-separated list. | 
 **scheduleId** | [**optional.Interface of []string**](.md)| Payment schedule IDs. Multiple IDs can be provided as a comma-separated list. | 
 **accountId** | [**optional.Interface of []string**](.md)| Originating account IDs. Multiple IDs can be provided as a comma-separated list. | 
 **customerId** | [**optional.Interface of []string**](.md)| The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. | 

### Return type

[**PaymentList**](payment_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPaymentSchedule**
> PaymentSchedule PatchPaymentSchedule(ctx, body, paymentScheduleId)
Update a payment schedule

Update a payment schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatchPaymentSchedule**](PatchPaymentSchedule.md)| payment schedule to update | 
  **paymentScheduleId** | [**string**](.md)| Payment schedule ID | 

### Return type

[**PaymentSchedule**](payment_schedule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

