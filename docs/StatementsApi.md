# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatement**](StatementsApi.md#GetStatement) | **Get** /statements/{statement_id} | Get a statement
[**ListStatements**](StatementsApi.md#ListStatements) | **Get** /statements | List statements

# **GetStatement**
> Statement GetStatement(ctx, statementId)
Get a statement

Get a statement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **statementId** | [**string**](.md)| The unique identifier of a statement | 

### Return type

[**Statement**](statement.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStatements**
> StatementList ListStatements(ctx, accountId)
List statements

Get list of statements

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | [**string**](.md)| The account&#x27;s unique identifier provided by Synctera | 

### Return type

[**StatementList**](statement_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

