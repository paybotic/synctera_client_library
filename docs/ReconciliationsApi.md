# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReconciliation**](ReconciliationsApi.md#CreateReconciliation) | **Post** /reconciliations | Create a reconciliation
[**GetReconciliation**](ReconciliationsApi.md#GetReconciliation) | **Get** /reconciliations/{reconciliation_id} | Get reconciliation
[**ListReconciliations**](ReconciliationsApi.md#ListReconciliations) | **Get** /reconciliations | List reconciliations

# **CreateReconciliation**
> Reconciliation CreateReconciliation(ctx, optional)
Create a reconciliation

Create a new reconciliation job 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReconciliationsApiCreateReconciliationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReconciliationsApiCreateReconciliationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ReconciliationInput**](ReconciliationInput.md)| Reconciliation to perform | 

### Return type

[**Reconciliation**](reconciliation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReconciliation**
> Reconciliation GetReconciliation(ctx, reconciliationId)
Get reconciliation

Retrieves one reconciliation by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reconciliationId** | [**string**](.md)| Reconciliation id | 

### Return type

[**Reconciliation**](reconciliation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReconciliations**
> ReconciliationList ListReconciliations(ctx, optional)
List reconciliations

Retrieves paginated list of reconciliations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReconciliationsApiListReconciliationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReconciliationsApiListReconciliationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 

### Return type

[**ReconciliationList**](reconciliation_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

