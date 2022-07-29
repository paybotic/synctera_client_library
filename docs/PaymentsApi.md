# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTransactionOut**](PaymentsApi.md#AddTransactionOut) | **Post** /ach | Create an outgoing ACH
[**GetTransactionOut**](PaymentsApi.md#GetTransactionOut) | **Get** /ach/{transaction_id} | Get an outgoing ACH transaction
[**ListTransactionsOut**](PaymentsApi.md#ListTransactionsOut) | **Get** /ach | List outgoing ACH transactions
[**PatchTransactionOut**](PaymentsApi.md#PatchTransactionOut) | **Patch** /ach/{transaction_id} | Update outgoing ACH transaction

# **AddTransactionOut**
> OutgoingAch AddTransactionOut(ctx, body)
Create an outgoing ACH

Create an outgoing ACH 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OutgoingAchRequest**](OutgoingAchRequest.md)| Outgoing ACH | 

### Return type

[**OutgoingAch**](outgoing_ach.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionOut**
> OutgoingAch GetTransactionOut(ctx, transactionId)
Get an outgoing ACH transaction

Get a single outgoing ACH transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionId** | [**string**](.md)| Transaction ID in the ledger | 

### Return type

[**OutgoingAch**](outgoing_ach.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactionsOut**
> OutgoingAchList ListTransactionsOut(ctx, optional)
List outgoing ACH transactions

List outgoing ACH transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentsApiListTransactionsOutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiListTransactionsOutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 

### Return type

[**OutgoingAchList**](outgoing_ach_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTransactionOut**
> PatchTransactionOut(ctx, body, transactionId)
Update outgoing ACH transaction

Update outgoing ACH transaction (either status or funds availability)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OutgoingAchPatch**](OutgoingAchPatch.md)| Outgoing ACH update request | 
  **transactionId** | [**string**](.md)| Transaction ID in the ledger | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

