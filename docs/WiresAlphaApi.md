# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelWire**](WiresAlphaApi.md#CancelWire) | **Patch** /wires/{wire_id} | Cancel an outgoing wire
[**CreateWire**](WiresAlphaApi.md#CreateWire) | **Post** /wires | Send a wire
[**GetWire**](WiresAlphaApi.md#GetWire) | **Get** /wires/{wire_id} | Get a wire by id
[**ListWires**](WiresAlphaApi.md#ListWires) | **Get** /wires | List wires

# **CancelWire**
> Wire CancelWire(ctx, body, wireId)
Cancel an outgoing wire

Cancel an outgoing tranfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateTransfer**](UpdateTransfer.md)| wire to update | 
  **wireId** | [**string**](.md)| The unique identifier of a wire | 

### Return type

[**Wire**](wire.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWire**
> Wire CreateWire(ctx, body)
Send a wire

Create an outgoing wire transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WireRequest**](WireRequest.md)| Wire transfer request | 

### Return type

[**Wire**](wire.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWire**
> Wire GetWire(ctx, wireId)
Get a wire by id

Get a wire by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **wireId** | [**string**](.md)| The unique identifier of a wire | 

### Return type

[**Wire**](wire.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWires**
> WireList ListWires(ctx, optional)
List wires

Get paginated list of wires

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WiresAlphaApiListWiresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WiresAlphaApiListWiresOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 
 **status** | **optional.String**|  | 
 **customerId** | [**optional.Interface of string**](.md)|  | 
 **originatingAccountId** | [**optional.Interface of string**](.md)|  | 
 **receivingAccountId** | [**optional.Interface of string**](.md)|  | 

### Return type

[**WireList**](wire_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

